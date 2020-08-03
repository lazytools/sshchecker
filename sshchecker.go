package sshchecker

import (
	"context"
	"errors"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type BatchOptions struct {
	UserList     []string
	PasswordList []string
	Timeout      time.Duration
	Concurrency  int
}

type BatchResult struct {
	Username string
	Password string
	Error    error
}

func BatchTrySSHLogin(ctx context.Context, addr *net.TCPAddr, opts *BatchOptions, output chan<- *BatchResult) error {
	if opts.Concurrency <= 0 {
		return errors.New("sshchecker: invalid concurrency value")
	}

	// I wouldn't typically do this, but for the sake of learning here's an
	// alternative pattern to the traditional producer-consumer worker pool.
	// This is a semaphore that limits the number of concurrent operations, and
	// behaves like a wait group at the same time.
	// Taken from: https://youtu.be/5zXAHh5tJqQ?t=1927
	sem := make(chan struct{}, opts.Concurrency)

	defer func() {
		// Wait for completion
		for i := 0; i < opts.Concurrency; i++ {
			sem <- struct{}{}
		}
	}()

	for _, username := range opts.UserList {
		for _, password := range opts.PasswordList {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sem <- struct{}{}:
			}

			go func(username, password string) {
				subctx, cancel := context.WithTimeout(ctx, opts.Timeout)
				err := TrySSHLogin(subctx, addr, username, password)
				cancel()
				output <- &BatchResult{
					Username: username,
					Password: password,
					Error:    err,
				}
				<-sem
			}(username, password)
		}
	}

	return nil
}

func TrySSHLogin(ctx context.Context, addr *net.TCPAddr, user, pass string) error {
	sshConfig := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", addr.String())
	if err != nil {
		return err
	}
	defer conn.Close()

	client, _, _, err := ssh.NewClientConn(conn, addr.String(), sshConfig)
	if err == nil {
		client.Close()
	}
	return err
}

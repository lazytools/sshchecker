package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/lazytools/sshchecker"
	"github.com/projectdiscovery/gologger"
)

func main() {
	usernamesPath := flag.String("U", "", "Text file containing list of usernames to use")
	passwordsPath := flag.String("P", "", "Text file containing list of passwords to use")
	concurrencyLevel := flag.Int("c", 20, "set the concurrency level")
	timeout := flag.Duration("t", 5*time.Second, "Connection timeout")
	showVersion := flag.Bool("version", false, "Show current program version")
	//verbose := flag.Bool("v", false, "Show verbose output")
	flag.Parse()
	showBanner()
	gologger.MaxLevel = gologger.Debug

	if *showVersion {
		gologger.Infof("Current Version: %s\n", Version)
		os.Exit(0)
	}

	if *usernamesPath == "" || *passwordsPath == "" {
		gologger.Errorf("username file path and password file path must be specified")
		flag.Usage()
		return
	}

	var options sshchecker.BatchOptions
	var err error

	options.UserList, err = parseFile(*usernamesPath)
	if err != nil {
		gologger.Fatalf("could not parse username file: %v", err)
	}
	options.PasswordList, err = parseFile(*passwordsPath)
	if err != nil {
		gologger.Fatalf("could not parse password file: %v", err)
	}

	gologger.Infof("[+] Loaded %d usernames and %d passwords", len(options.UserList), len(options.PasswordList))

	options.Timeout = *timeout
	options.Concurrency = *concurrencyLevel

	ctx := contextWithSignal(context.Background())

	processFromStdin(ctx, &options)

	gologger.Infof("[+] EOF reached")
}

func processFromStdin(ctx context.Context, options *sshchecker.BatchOptions) {
	// We use a pipe so we can close stdin upon context completion
	rd, wr := io.Pipe()
	go func() {
		_, err := io.Copy(wr, os.Stdin)
		if err != nil {
			wr.CloseWithError(err)
		} else {
			wr.CloseWithError(io.EOF)
		}
	}()
	go func() {
		<-ctx.Done()
		wr.Close()
		rd.Close()
	}()

	scn := bufio.NewScanner(rd)
	for scn.Scan() {
		rawAddr := strings.TrimSpace(scn.Text())
		if !strings.Contains(rawAddr, ":") {
			gologger.Infof("address is missing port, defaulting to port 22")
			rawAddr += ":22"
		}

		addr, err := net.ResolveTCPAddr("tcp", rawAddr)
		if err != nil {
			gologger.Errorf("[!] failed to parse address: %v", err)
			continue
		}

		gologger.Infof("[+] Now processing address: %s (resolved from %s)", addr.String(), rawAddr)
		output := make(chan *sshchecker.BatchResult)
		var batchError error

		addrCtx, cancel := context.WithCancel(ctx)
		go func() {
			batchError = sshchecker.BatchTrySSHLogin(addrCtx, addr, options, output)
			close(output)
		}()

		success := false
		for out := range output {
			if out.Error != nil {
				if !success {
					gologger.Warningf("[!] Failed to login on %s with %s:%s, error: %v",
						addr.String(), out.Username, out.Password, out.Error)
				}
				continue
			}

			if !success {
				gologger.Infof("[+] Successful login on %s with %s:%s", addr.String(), out.Username, out.Password)
				cancel()
				success = true
			}
		}
		cancel()

		if batchError != nil {
			gologger.Warningf("[!] Error while batch logging in on %s: %v", addr.String(), batchError)
		}

		if ctx.Err() != nil {
			gologger.Fatalf("[!] quitting due to context error: %v", ctx.Err())
		}
	}
}

func contextWithSignal(parent context.Context) context.Context {
	ctx, cancel := context.WithCancel(parent)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		gologger.Infof("Ctrl+C received, gracefully cancelling...")
		cancel()
	}()

	return ctx
}

func parseFile(filename string) ([]string, error) {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(d), "\n")
	i := 0
	for i < len(rows) {
		rows[i] = strings.TrimSpace(rows[i])
		if rows[i] == "" {
			rows = append(rows[:i], rows[i+1:]...)
			continue
		}
		i++
	}
	return rows, nil
}

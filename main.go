package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/projectdiscovery/gologger"
)

func main() {
	showBanner()
	//adding scanner for reading the input from terminal
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		ip := sc.Text()
		fmt.Printf("%v \n", ip)

		//		sshconfig := &gosshtool.SSHClientConfig{
		//			user:     "user",
		//			Password: "pwd",
		//			Host:     ip,
		//		}
		//		sshclient := gosshtool.NewSSHClient(sshconfig)
		//		t.log(sshclient.Host)
		//		stdout, stderr, session, err := sshclient.Cmd("pwd", nil, nil, 0)
		//		if err != nil {
		//			t.Error(err)
		//
		//		}
		//		t.Log(stdout)
		//		t.Log(stderr)
	}

}

const banner = `

   _____      __    ________              __            
  / ___/_____/ /_  / ____/ /_  ___  _____/ /_____  _____
  \__ \/ ___/ __ \/ /   / __ \/ _ \/ ___/ //_/ _ \/ ___/
 ___/ (__  ) / / / /___/ / / /  __/ /__/ ,< /  __/ /    
/____/____/_/ /_/\____/_/ /_/\___/\___/_/|_|\___/_/v1.0
                                                        

`

func showBanner() {
	gologger.Printf("%s\n", banner)
	gologger.Printf("\t\tCreated by Shazeb.\n\n")

	gologger.Labelf("Use with caution. You are responsible for your actions\n")
	gologger.Labelf("Developers assume no liability and are not responsible for any misuse or damage.\n\n")
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//adding scanner for reading the input from terminal
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		ip := sc.Text()
		fmt.Printf("%v \n", ip)
	}

}

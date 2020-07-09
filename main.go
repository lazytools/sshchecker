package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		ip := sc.Text()
		fmt.Printf("%v \n", ip)
	}

}

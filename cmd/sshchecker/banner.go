package main

import "github.com/projectdiscovery/gologger"

const banner = `
   ____  _____ __  __  _____ __  ________  ______
  / __ \/ ___// / / / / ___// / / / ____/ / ____/
 / /_/ /\__ \/ /_/ /  \__ \/ /_/ / /     / /_
/ _, _/___/ / __  /  ___/ / __  / /___  / __/
/_/ |_|/____/_/ /_/  /____/_/ /_/\____/ /_/
`

const Version = `1.0`

func showBanner() {
	gologger.Printf("%s\n", banner)
	gologger.Printf("\t\tCreated by lazytools.\n\n")

	gologger.Labelf("Use with caution. You are responsible for your actions\n")
	gologger.Labelf("Developers assume no liability and are not responsible for any misuse or damage.\n\n")
}

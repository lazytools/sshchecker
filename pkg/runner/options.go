package runner

import (
	"flag"
	"os"

	"github.com/projectdiscovery/gologger"
)

type Options struct {
	Verbose      bool
	UserList     string
	Concurrency  int
	PasswordList string
	ShowVer      bool
}

func ParseOptions() *Options {
	options := &Options{}

	flag.StringVar(&options.UserList, "U", "", "Text file containing list of usernames to use")
	flag.StringVar(&options.PasswordList, "P", "", "Text file containing list of passwords to use")
	flag.IntVar(&options.Concurrency, "c", 20, "set the concurrency level")
	flag.BoolVar(&options.ShowVer, "version", false, "Show current program version")
	flag.BoolVar(&options.Verbose, "v", false, "Show Verbose output.")
	flag.Parse()

	showBanner()

	if options.ShowVer {
		gologger.Infof("Current Version: %s\n", Version)
		os.Exit(0)
	}
	return options
}

package main

import (
	"fmt"
	"git.sevone.com/glider-guns/showrunner.git/showrunner"
	"os"
	"sort"
)

var options = map[string]func() showrunner.Command{
	"file-split": func() showrunner.Command { return new(showrunner.FileSplit) },
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments for your command")
		os.Exit(1)
	}
	command, found := options[os.Args[1]]
	if !found {
		fmt.Printf("Command %v not recognized, the choices are:\n", os.Args[1])
		opts := []string{}
		for key, _ := range options {
			opts = append(opts, key)
		}
		sort.Strings(opts)
		for _, key := range opts {
			fmt.Printf("\t%v\n", key)
		}
		os.Exit(1)
	}
	err := command().Main(os.Args[2:])
	if err != nil {
		fmt.Println("Something went wrong with your command")
		os.Exit(1)
	}

}

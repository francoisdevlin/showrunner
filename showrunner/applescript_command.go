package showrunner

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type Applescript struct {
}

func (this *Applescript) Main(args []string) error {
	f := flag.NewFlagSet("test", flag.ExitOnError)
	wordPtr := f.String("file", "foo", "This is a file to split up")
	outfile := f.String("intermediate-file", "/tmp/temp_command.applescript", "This is the intermediate file that stores the applescript")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "This is a tool to assist with the creation of training material")
		flag.PrintDefaults()
	}
	f.Parse(args)

	buff, _ := ioutil.ReadFile(*wordPtr)
	inputText := string(buff)
	lines := SplitLines(inputText)
	fileContents := lineBuilder(lines, *wordPtr, 13)
	err := ioutil.WriteFile(*outfile, []byte(fileContents), 0644)
	if err != nil {
		return err
	}
	cmd := exec.Command("/usr/bin/osascript", *outfile)
	_, execErr := cmd.Output()
	if execErr != nil {
		return execErr
	}
	return nil
}

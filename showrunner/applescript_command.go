package showrunner

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	//"path/filepath"
)

type Applescript struct {
}

func (this *Applescript) Main(args []string) error {
	f := flag.NewFlagSet("test", flag.ExitOnError)
	wordPtr := f.String("file", "foo", "This is a file to split up")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "This is a tool to assist with the creation of training material")
		flag.PrintDefaults()
	}
	f.Parse(args)

	buff, _ := ioutil.ReadFile(*wordPtr)
	inputText := string(buff)
	lines := SplitLines(inputText)
	fmt.Println(lineBuilder(lines, *wordPtr, 13))
	return nil
}

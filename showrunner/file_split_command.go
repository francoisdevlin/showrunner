package showrunner

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileSplit struct {
}

func (this *FileSplit) Main(args []string) error {
	f := flag.NewFlagSet("test", flag.ExitOnError)
	wordPtr := f.String("file", "foo", "This is a file to split up")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "This is a tool to assist with the creation of training material")
		flag.PrintDefaults()
	}
	f.Parse(args)
	fmt.Println(*wordPtr)
	fmt.Println(filepath.Base(*wordPtr))
	fmt.Println(filepath.Dir(*wordPtr))
	fmt.Println(filepath.Ext(*wordPtr))
	comment := GetLineCommentString(*wordPtr)

	buff, _ := ioutil.ReadFile(*wordPtr)
	inputText := string(buff)
	for i, temp := range BuildText(comment, SplitLines(inputText)) {
		fmt.Println(i, "\n")
		fmt.Println(temp)
	}
	return nil
}

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
	comment := GetLineCommentString(*wordPtr)

	buff, _ := ioutil.ReadFile(*wordPtr)
	inputText := string(buff)
	baseName := filepath.Base(*wordPtr)
	ext := filepath.Ext(*wordPtr)
	basic := baseName[0 : len(baseName)-len(ext)]
	//Remove the trailing -all, if it exists
	if basic[len(basic)-4:] == "-all" {
		basic = basic[:len(basic)-4]
	}
	for i, fileContents := range BuildText(comment, SplitLines(inputText)) {
		outfile := fmt.Sprintf("%s/%s-%d%s", filepath.Dir(*wordPtr), basic, i, ext)
		fmt.Println(outfile, "\n")
		err := ioutil.WriteFile(outfile, []byte(fileContents), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

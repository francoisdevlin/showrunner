package main

import (
	"flag"
	"fmt"
	"git.sevone.com/glider-guns/showrunner.git/showrunner"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	wordPtr := flag.String("file", "foo", "This is a file to split up")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "This is a tool to assist with the creation of training material")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println(*wordPtr)
	fmt.Println(filepath.Base(*wordPtr))
	fmt.Println(filepath.Dir(*wordPtr))
	fmt.Println(filepath.Ext(*wordPtr))
	comment := showrunner.GetLineCommentString(*wordPtr)

	buff, _ := ioutil.ReadFile(*wordPtr)
	inputText := string(buff)
	for i, temp := range showrunner.BuildText(comment, showrunner.SplitLines(inputText)) {
		fmt.Println(i, "\n")
		fmt.Println(temp)
	}

}

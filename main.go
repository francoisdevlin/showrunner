package main

import (
	"fmt"
	"regexp"
	//"strings"

	"git.sevone.com/glider-guns/showrunner.git/showrunner"
)

func splitLines(text string) []string {
	rp := regexp.MustCompile("[\r\n]")
	return rp.Split(text, -1)
}

func main() {
	comment := "--"
	inputText := `Hi
--split
Everyone
Isn't
--split
This
--re-split
Terrible?
--split
Awesome?
`
	fmt.Println(len(splitLines(inputText)))
	for i, temp := range showrunner.BuildText(comment, splitLines(inputText)) {
		fmt.Println(i)
		fmt.Println(temp)
	}
}

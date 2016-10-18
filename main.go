package main

import (
	"fmt"

	"git.sevone.com/glider-guns/showrunner.git/showrunner"
)

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
	for i, temp := range showrunner.BuildText(comment, showrunner.SplitLines(inputText)) {
		fmt.Println(i)
		fmt.Println(temp)
	}
}

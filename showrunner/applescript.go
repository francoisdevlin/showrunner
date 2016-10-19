package showrunner

import (
	"fmt"
)

var charCodes = map[string]string{
	"1":  "18",
	"2":  "19",
	"3":  "20",
	"4":  "21",
	"5":  "23",
	"6":  "22",
	"7":  "26",
	"8":  "28",
	"9":  "25",
	"0":  "29",
	"+":  "24 using shift down",
	"-":  "27",
	"\"": "39 using shift down",
	"`":  "50",
	".":  "47",
	"\\": "\"\\\"",
}

func printWord(word string) string {
	if len(word) == 0 {
		return ""
	}
	output := "\ttell application \"System Events\"\n"
	maxKeys := 3
	current := 0
	for _, c := range word {
		code, found := charCodes[string(c)]
		if found {
			output += fmt.Sprintf("\t\tkey code \"%s\"\n", code)
		} else {
			output += fmt.Sprintf("\t\tkeystroke \"%s\"\n", string(c))
		}
		if current < maxKeys {
			output += "\t\tdelay .02\n"
			current++
		}
	}
	output += "\tend tell\n"
	return output
}

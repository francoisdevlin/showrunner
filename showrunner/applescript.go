package showrunner

import (
	"fmt"
	"regexp"
	"strings"
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

type lineHandler interface {
	enterKeyStrokes(line string) string
}

type echoLineHandler struct{}

func (this *echoLineHandler) enterKeyStrokes(line string) string {
	return line
}

type defaultLineHandler struct {
	wordHandler        func(string) string
	includeWordComment bool
	delay              float64
}

var spaceInfo = `
	tell application "System Events" to keystroke space
	delay .1

`

func (this *defaultLineHandler) enterKeyStrokes(line string) string {
	if len(line) == 0 {
		return ""
	}
	rp := regexp.MustCompile("\\s+")
	words := rp.Split(line, -1)
	wordEntries := []string{}
	for _, word := range words {
		entry := ""
		if this.includeWordComment {
			entry += fmt.Sprintf("\t-- \"%s\"\n", word)
		}
		entry += this.wordHandler(word)
		if this.delay != 0 {
			entry += fmt.Sprintf("\tdelay %.2f\n")
		}
		wordEntries = append(wordEntries, entry)
	}
	output := strings.Join(wordEntries, spaceInfo)
	output += "\ttell application \"System Events\" to keystroke return\n"
	return output
}

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
			output += fmt.Sprintf("\t\tkey code %s\n", code)
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
			entry += fmt.Sprintf("\tdelay %.2f\n", this.delay)
		}
		wordEntries = append(wordEntries, entry)
	}
	output := strings.Join(wordEntries, spaceInfo)
	output += "\ttell application \"System Events\" to keystroke return\n"
	return output
}

func enterKeyStrokes(line string) string {
	handler := new(defaultLineHandler)
	handler.wordHandler = printWord
	handler.delay = 0.1
	return handler.enterKeyStrokes(line)
}

func waitForScript(line string) string {
	output := fmt.Sprintf("\tset w to do script \"%s\" in currentTab\n", line) +
		`
	repeat 
		delay 1
		if not busy of w then exit repeat
	end repeat

`
	return output
}

func getHeader(size int) string {
	header := `
tell application "Terminal"
	set theDesktop to POSIX path of (path to desktop as string)
	activate
	set frontWindow to window 1
	set currentTab to do script "echo 'Hello World'"
	tell application "System Events"

`
	for i := 0; i < size; i++ {
		header += "\t\tkeystroke \"+\" using {command down}\n"
	}
	header += `
		keystroke "f" using {command down, control down}
	end tell
	delay 5
	do script "clear" in currentTab
	delay 5

`
	return header
}

func lineBuilder(lines []string, filepath string, size int) string {
	output := getHeader(size)
	i := -1
	snapshot := func() string {
		i++
		return (fmt.Sprintf("\tset shellCommand to \"/usr/sbin/screencapture \" & theDesktop & \"%s-%d.png\"\n", filepath, i) +
			"\tdo shell script shellCommand\n")
	}
	output += snapshot()
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if line[0:1] == "!" {
			output += enterKeyStrokes(line[1:])
		} else {
			output += waitForScript(line)
		}
		output += "\tdelay 1\n"
		output += snapshot()
	}
	output += `
	delay 5
	tell application "System Events"
		keystroke "w" using {command down}
	end tell
end tell
`
	return output
}

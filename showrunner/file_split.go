package showrunner

import (
	//"fmt"
	"regexp"
	"strings"
)

func SplitLines(text string) []string {
	rp := regexp.MustCompile("[\r\n]")
	return rp.Split(text, -1)
}

func appendEntry(entries [][]string, current []string, isResplit bool) [][]string {

	if isResplit {
		if len(entries) > 1 {
			current = append(entries[len(entries)-2], current...)
		}
		if len(current) == 0 {
			return entries
		}
	} else {
		if len(current) == 0 {
			return entries
		}
		if len(entries) > 0 {
			current = append(entries[len(entries)-1], current...)
		}
	}
	return append(entries, current)
}

func BuildText(comment string, lines []string) []string {
	current := []string{}
	entries := [][]string{}
	isResplit := false
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == comment+"re-split" {
			entries = appendEntry(entries, current, isResplit)
			isResplit = true
			current = []string{}
		} else if trimmedLine == comment+"split" {
			entries = appendEntry(entries, current, isResplit)
			isResplit = false
			current = []string{}
		} else {
			current = append(current, line)
		}
	}
	entries = appendEntry(entries, current, isResplit)

	output := []string{}
	for _, entry := range entries {
		output = append(output, strings.Join(entry, "\n"))
	}
	return output
}

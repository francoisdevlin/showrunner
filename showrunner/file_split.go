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

func BuildText(comment string, lines []string) []string {
	current := []string{}
	entries := [][]string{}
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == comment+"re-split" {
			if len(entries) > 0 {
				current = append(entries[len(entries)-1], current...)
			}
			entries = append(entries, current)
			current = []string{}
		} else if trimmedLine == comment+"split" {
			if len(current) != 0 {
				if len(entries) > 0 {
					current = append(entries[len(entries)-1], current...)
				}
				entries = append(entries, current)
			}
			current = []string{}
		} else {
			current = append(current, line)
		}
	}
	if len(current) != 0 {
		if len(entries) > 0 {
			current = append(entries[len(entries)-1], current...)
		}
		entries = append(entries, current)
	}

	output := []string{}
	for _, entry := range entries {
		output = append(output, strings.Join(entry, "\n"))
	}
	return output
}

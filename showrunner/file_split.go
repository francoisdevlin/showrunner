package showrunner

import (
	//"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func SplitLines(text string) []string {
	rp := regexp.MustCompile("[\r\n]")
	return rp.Split(text, -1)
}

func appendEntry(entries [][]string, current []string, lastSplit int) [][]string {

	if len(entries) != lastSplit+1 {
		if len(entries) > 1 {
			current = append(entries[lastSplit], current...)
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
	lastSplit := -1
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == comment+"re-split" {
			entries = appendEntry(entries, current, lastSplit)
			current = []string{}
		} else if trimmedLine == comment+"split" {
			entries = appendEntry(entries, current, lastSplit)
			lastSplit = len(entries) - 1
			current = []string{}
		} else {
			current = append(current, line)
		}
	}
	entries = appendEntry(entries, current, lastSplit)

	output := []string{}
	for _, entry := range entries {
		output = append(output, strings.Join(entry, "\n"))
	}
	return output
}

func GetLineCommentString(filename string) string {
	extension := filepath.Ext(filename)
	return map[string]string{
		".hs":   "--",
		".go":   "//",
		".java": "//",
		".php":  "//",
	}[extension]
}

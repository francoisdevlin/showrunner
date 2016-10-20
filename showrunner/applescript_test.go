package showrunner

import (
	"testing"
)

/*

*/
func TestBasicWord(t *testing.T) {
	pairs := [][]string{
		[]string{"", ""},
		[]string{"hi", `	tell application "System Events"
		keystroke "h"
		delay .02
		keystroke "i"
		delay .02
	end tell
`},
		[]string{"2", `	tell application "System Events"
		key code "19"
		delay .02
	end tell
`},
		[]string{"baconmaster", `	tell application "System Events"
		keystroke "b"
		delay .02
		keystroke "a"
		delay .02
		keystroke "c"
		delay .02
		keystroke "o"
		keystroke "n"
		keystroke "m"
		keystroke "a"
		keystroke "s"
		keystroke "t"
		keystroke "e"
		keystroke "r"
	end tell
`},
	}
	for _, pair := range pairs {
		input := pair[0]
		expected := pair[1]
		actual := printWord(input)
		if expected != actual {
			t.Errorf("Word macro test failures\ninput:'%v'\nexpected:'%v'\nactual:'%v'", input, expected, actual)
		}
	}
}

/*

*/
func TestDefaultLineHandler(t *testing.T) {
	pairs := [][]string{
		[]string{"", ``},
		[]string{"a", `a
	tell application "System Events" to keystroke return
`},
		[]string{"2 + 2", `2

	tell application "System Events" to keystroke space
	delay .1

+

	tell application "System Events" to keystroke space
	delay .1

2
	tell application "System Events" to keystroke return
`},
	}
	handler := new(defaultLineHandler)
	handler.wordHandler = func(line string) string {
		return line + "\n"
	}
	for _, pair := range pairs {
		input := pair[0]
		expected := pair[1]
		actual := handler.enterKeyStrokes(input)
		if expected != actual {
			t.Errorf("Word macro test failures\ninput:'%v'\nexpected:'%v'\nactual:'%v'", input, expected, actual)
		}
	}
}

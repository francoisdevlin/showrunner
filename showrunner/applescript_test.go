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

package showrunner

import (
	"testing"
)

/*
If there are no split statements, there is only one entry, and it returns the entire contents of the text
*/
func TestBasicEcho(t *testing.T) {
	comment := "--"
	inputText := `Hi
Everyone
Isn't
This
Awesome?
`
	output := BuildText(comment, SplitLines(inputText))
	if len(output) != 1 {
		t.Error("The output array is the wrong size, it should be 1")
	}
	if inputText != output[0] {
		t.Error("The output body does not match the expected value")
	}
}

/*
Tests the most basic split usage
*/
func TestBasicSplit(t *testing.T) {
	comment := "--"
	inputText := `Hi
Everyone
--split
Isn't
This
--split
Awesome?
`
	output := BuildText(comment, SplitLines(inputText))
	if len(output) != 3 {
		t.Error("The output array is the wrong size, it should be 3")
	}
	if output[0] != `Hi
Everyone` {
		t.Error("The first entry does not match the expected value")
	}
	if output[1] != `Hi
Everyone
Isn't
This` {
		t.Error("The second entry does not match the expected value")
	}
	if output[2] != `Hi
Everyone
Isn't
This
Awesome?
` {
		t.Error("The thrid entry does not match the expected value")
	}
}

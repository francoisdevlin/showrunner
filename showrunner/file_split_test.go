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
		t.Error(output[0])
		t.Error(output[1])
		t.Error(output[2])
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

/*
Tests the most basic re-split usage
*/
func TestBasicResplit(t *testing.T) {
	comment := "--"
	inputText := `Hi
Everyone
--split
Isn't
This
--split
Terrible?
--re-split
Awesome?
`
	output := BuildText(comment, SplitLines(inputText))
	if len(output) != 4 {
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
Terrible?` {
		t.Error("The thrid entry does not match the expected value")
	}
	if output[3] != `Hi
Everyone
Isn't
This
Awesome?
` {
		t.Error("The fourth entry does not match the expected value")
	}
}

/*
Tests the most basic re-split usage
*/
func TestRepeatedResplit(t *testing.T) {
	comment := "--"
	inputText := `Hi
Everyone
--split
Isn't
This
--split
Terrible?
--re-split
Ugly?
--re-split
Fun?
--re-split
Awesome?
`
	output := BuildText(comment, SplitLines(inputText))
	if len(output) != 6 {
		t.Error("The output array is the wrong size, it should be 6")
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
Terrible?` {
		t.Error("The thrid entry does not match the expected value")
	}
	if output[3] != `Hi
Everyone
Isn't
This
Ugly?` {
		t.Error("The thrid entry does not match the expected value")
	}
	if output[4] != `Hi
Everyone
Isn't
This
Fun?` {
		t.Error("The thrid entry does not match the expected value")
	}
	if output[5] != `Hi
Everyone
Isn't
This
Awesome?
` {
		t.Error("The fourth entry does not match the expected value")
	}
}

/*
Tests the most basic re-split usage
*/
func TestBasicResplitWithTrailingSplit(t *testing.T) {
	comment := "--"
	inputText := `Hi
Everyone
--split
Ooops
--re-split
Isn't
This
--split
Awesome?
`
	output := BuildText(comment, SplitLines(inputText))
	if len(output) != 4 {
		t.Error("The output array is the wrong size, it should be 3")
	}
	if output[0] != `Hi
Everyone` {
		t.Error("The first entry does not match the expected value")
	}
	if output[1] != `Hi
Everyone
Ooops` {
		t.Error("The second entry does not match the expected value")
	}
	if output[2] != `Hi
Everyone
Isn't
This` {
		t.Error("The thrid entry does not match the expected value")
	}
	if output[3] != `Hi
Everyone
Isn't
This
Awesome?
` {
		t.Error("The fourth entry does not match the expected value")
	}
}

/*
Tests that splits ignore whitespace
*/
func TestKeywordsIgnoreWhitespace(t *testing.T) {
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

func TestCommentLinePrefix(t *testing.T) {
	values := map[string]string{
		"myfile.hs": "--",
	}

	for key, expectedValue := range values {
		if expectedValue != GetLineCommentString(key) {
			t.Error("Could not verify tuple", key, expectedValue, GetLineCommentString(key))
		}
	}
}

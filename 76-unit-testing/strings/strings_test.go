package strings

import "testing"

func TestReverseString(t *testing.T) {
	input := new(string)
	*input = "Victoria Secrets & Co"
	expectedOutput := "oC & sterceS airotciV"

	actualOutput := Reverse(input)
	if expectedOutput != actualOutput {
		t.Fail()
	}
}

// testing edge case with nil input string
func TestReverseStringNil(t *testing.T) {
	var expectedOutput *string
	actualOutput := Reverse(nil)
	if expectedOutput != nil || actualOutput != "" {
		t.Fail()
	}

}

func TestSrtingLength(t *testing.T) {
	input := new(string)
	*input = "Victoria Secrets & Co"

	expectedOutput := 21

	actualOutput := Length(input)

	if expectedOutput != actualOutput {
		t.Fail()
	}
}

func TestSrtingLengthNil(t *testing.T) {

	var input *string
	expectedOutput := -1

	actualOutput := Length(input)

	if (expectedOutput != actualOutput) && input == nil {
		t.Fail()
	}
}

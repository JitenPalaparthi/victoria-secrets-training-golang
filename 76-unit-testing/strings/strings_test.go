package strings

import "testing"

func TestReverseString(t *testing.T) {
	input := "Victoria Secrets & Co"
	expectedOutput := "oC & sterceS airotciV"

	actualOutput := Reverse(input)
	if expectedOutput != actualOutput {
		t.Fail()
	}
}

func TestSrtingLength(t *testing.T) {
	input := "Victoria Secrets & Co"

	expectedOutput := 21

	actualOutput := Length(input)

	if expectedOutput != actualOutput {
		t.Fail()
	}

}

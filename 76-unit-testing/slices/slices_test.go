package slices

import (
	"fmt"
	"testing"
)

func TestSecondBiggest(t *testing.T) {

	var si1 SliceInt

	si1 = si1.Fill(23, 43, 1243, 34, 34, 3, 98, 56, 45, 99, 67, 45, 65, 23, 900)

	expectedResult := 900

	actualResult, err := si1.SecondBiggest()
	fmt.Println(actualResult)
	if err != nil {
		t.Fail()
	}

	if actualResult != expectedResult {
		t.Fail()
	}
}

func TestSecondBiggest2(t *testing.T) {

	var si1 SliceInt

	si1 = si1.Fill(23, 43, 1243, 34, 34, 3, 98, 56, 45, 99, 67, 45, 65, 23, 900)

	expectedResult := 900

	actualResult, err := si1.SecondBiggest2()
	fmt.Println(actualResult)
	if err != nil {
		t.Fail()
	}

	if actualResult != expectedResult {
		t.Fail()
	}
}

func BenchmarkSecondBiggest(b *testing.B) {
	var si1 SliceInt
	for i := 1; i <= b.N; i++ {
		//si1 = si1.Fill(23, 43, 1243, 34, 34, 3, 98, 56, 45, 99, 67, 45, 65, 23, 900)
		si1 = si1.RandFill(100)
		si1.SecondBiggest()
	}
}

func BenchmarkSecondBiggest2(b *testing.B) {
	var si1 SliceInt
	for i := 1; i <= b.N; i++ {
		//si1 = si1.Fill(23, 43, 1243, 34, 34, 3, 98, 56, 45, 99, 67, 45, 65, 23, 900)
		si1 = si1.RandFill(100)
		si1.SecondBiggest2()
	}
}

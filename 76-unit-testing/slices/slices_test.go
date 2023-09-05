package slices

import (
	"fmt"
	"testing"
)

func TestInferDefaultsNil(t *testing.T) {

	var si SliceInt

	si = si.InferDefaults(10)

	if si != nil {
		t.Fail()
	}

}

func TestInferDefaults(t *testing.T) {

	var si SliceInt

	si = make([]int, 5)
	si = si.InferDefaults(10)

	for _, v := range si {
		if v != 10 {
			t.Fail()
		}
	}

}

func TestInferRandFill(t *testing.T) {

	var si SliceInt

	si = si.RandFill(10)

	if len(si) != 10 {
		t.Fail()
	}

}

func TestFill(t *testing.T) {

	var si SliceInt

	si = si.Fill(10, 20, 30)

	if si[0] != 10 || si[1] != 20 || si[2] != 30 {
		t.Fail()
	}

}

func TestMaxNil(t *testing.T) {

	var si SliceInt

	num, err := si.Max()

	if num != 0 {
		t.Fail()
	}
	if err == nil {
		t.Fail()
	}
}

func TestSecondBiggestNil(t *testing.T) {

	var si SliceInt

	num, err := si.SecondBiggest()

	if num != 0 {
		t.Fail()
	}
	if err == nil {
		t.Fail()
	}

}

func TestSecondBiggestOne(t *testing.T) {

	var si SliceInt

	si = si.Fill(10)

	num, err := si.SecondBiggest()

	if num != 10 {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}

}

func TestSecondBiggestTwo(t *testing.T) {

	var si SliceInt

	si = si.Fill(10, 20)

	num, err := si.SecondBiggest()

	if num != 10 {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
	var si1 SliceInt
	si1 = si1.Fill(20, 10)
	num1, err1 := si1.SecondBiggest()

	if num1 != 10 {
		t.Fail()
	}
	if err1 != nil {
		t.Fail()
	}

}

func TestSecondBiggest2Nil(t *testing.T) {

	var si SliceInt

	num, err := si.SecondBiggest2()

	if num != 0 {
		t.Fail()
	}
	if err == nil {
		t.Fail()
	}

}

func TestSecondBiggest2One(t *testing.T) {

	var si SliceInt

	si = si.Fill(10)

	num, err := si.SecondBiggest2()

	if num != 10 {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}

}

func TestSecondBiggest2Two(t *testing.T) {

	var si SliceInt

	si = si.Fill(10, 20)

	num, err := si.SecondBiggest2()

	if num != 10 {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
	var si1 SliceInt
	si1 = si1.Fill(20, 10)
	num1, err1 := si1.SecondBiggest2()

	if num1 != 10 {
		t.Fail()
	}
	if err1 != nil {
		t.Fail()
	}

}

func TestMaxOne(t *testing.T) {

	var si SliceInt

	si = si.Fill(10)

	num, err := si.Max()

	if num != 10 {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}

}

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

func TestSecondBiggest1_2(t *testing.T) {

	var si1 SliceInt

	si1 = si1.Fill(43, 12, 1243, 34, 34, 3, 98, 56, 45, 99, 67, 45, 65, 23, 900)

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

func TestSecondBiggest2_2(t *testing.T) {

	var si1 SliceInt

	si1 = si1.Fill(43, 23, 1243, 34, 34, 3, 98, 56, 45, 99, 67, 45, 65, 23, 900)

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

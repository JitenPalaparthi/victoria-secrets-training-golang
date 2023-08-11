package main

import "fmt"

func main() {

	var icalc ICalculator // nil interfaces can also be nil if nothing has been assigned to it
	if icalc == nil {
		fmt.Println("nil interface")
	}

	it1 := &IntType{A: 100, B: 200}
	ft1 := &Float32Type{A: 100.1, B: 200.9}
	icalc = it1
	r1 := icalc.Add()
	fmt.Println("Addition of IntType:", r1)
	icalc = ft1
	r2 := icalc.Add()
	fmt.Println("Addition of Float32Type:", r2)
	//icalc.Display() // This cannot be called becase Display is not a definition of interface
	it1.Display()

	s1 := &StringType{Str1: "Victoria", Str2: " Secrets"}
	icalc = s1
	r3 := icalc.Add()

	fmt.Println("Add of StringType:", r3)
}

type ICalculator interface {
	Add() any
	// Sub() any
	// Mul() any
	// Div() any
	//Get() any
	//Clear()
}

type IntType struct {
	A, B   int
	Result int
}

func (i *IntType) Add() any {
	i.Result = i.A + i.B
	return i.Result
}

func (i *IntType) Display() {
	fmt.Println("A:", i.A, "B:", i.B)
}

type Float32Type struct {
	A, B   float32
	Result float32
}

func (f *Float32Type) Add() any {
	f.Result = f.A + f.B
	return f.Result
}

type StringType struct {
	Str1, Str2 string
	Result     string
}

func (s *StringType) Add() any {
	s.Result = s.Str1 + s.Str2
	return s.Result
}

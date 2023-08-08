package main

import "fmt"

// gobal variables
var (
	num1 int
	num2 float32
	num3 float64
)

func main() {

	var str1 string = "Hello World"

	fmt.Println(str1)

	fmt.Println("Length of str1 ", len(str1))

	// mutating string

	str1 = "Hello Victoria Secrets & Co"

	fmt.Println("Length of str1 ", len(str1))

	num1 := 10000 // shorthand declaration. It is int and int holds 8 bytes on 64bit processors
	fmt.Println(num1)
	num1 = 20000 // mutating a num1

	num1 = 1232432324324

}

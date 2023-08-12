package main

import "fmt"

func main() {
	Greet()

	func() {
		str := "Hello Victoria Secrets & Co! USA"
		fmt.Println(str)
	}() // Exec

	func(str string) {
		fmt.Println(str)
	}("Hello Victoria Secrets & Co Europe!")

	// Question to you? What if I use for loop instead of for range
	rstr := func(str string) string {
		reverse := ""
		for _, v := range str {
			reverse = string(v) + reverse
		}
		return reverse
	}("Hello Vitoria Secrets & Co")

	fmt.Println("Reverse:", rstr)

	f1 := func(str string) string {
		reverse := ""
		for _, v := range str {
			reverse = string(v) + reverse
		}
		return reverse
	}

	rstr = f1("Hello Vitoria Secrets & Co!")
	fmt.Println("Reverse:", rstr)

	// func(str string) (v int, c int, sp int) {
	// 	// write logic here
	// 	return v, c, sp
	// }("Hello Victoria Secrets & Co!")

	// fn2 := func(str string) (v int, c int, sp int) {

	// 	// write logic here
	// 	return v, c, sp
	// }

	// func argument is an anonymous function
	sum := Calc(10, 20, func(a1, b1 int) int {
		return a1 + b1
	})
	fmt.Println("Sum:", sum)

	// func argument is a named function
	sub := Calc(10, 20, Sub)
	fmt.Println("Sub:", sub)

	fnm := func(a, b int) int {
		return a * b
	}

	// func argument is a func variable
	mul := Calc(10, 20, fnm)

	fmt.Println("Multiplication:", mul)

}

//var f1 func(string) string

func Greet() {
	fmt.Println("Hello Victoria Secrets & Co! Bangalore")
}

// What is an anonymous function
// Write an anonymous function that should return number of Vovels ,
// Consonents and Special Chars in a given string

func Calc(a, b int, fn func(a1, b1 int) int) int {
	return fn(a, b)
}

func Sub(a, b int) int {
	c := a - b
	return c
}

var fnmap map[string]func(int, int) int

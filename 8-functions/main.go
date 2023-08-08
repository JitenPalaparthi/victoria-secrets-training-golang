package main

import (
	"fmt"
	_ "time" // use blank identifier _ when not going to use the package
)

func main() {
	greet()
	greet()

	greetBy("Victoria nosecrets") // argument
	greetByM("Hey!", "Victoria")

	a1 := areaOfRect(10.23, 13.45)
	fmt.Println("Area of Rect:", a1)

	a2, p2 := Rect(10.23, 13.45)
	fmt.Println("Area of Rect:", a2, "Perimeter of Rect:", p2)

	a3, _ := Rect(10.23, 13.45)
	fmt.Println("Area of Rect:", a3)

	_, p4 := Rect(10.23, 13.45)
	fmt.Println("Perimeter of Rect:", p4)

	num1 := 100
	incr(num1)

	fmt.Println("Value of num1:", num1)

	num1 = incrR(num1)
	fmt.Println("Value of num1:", num1)

}

func greet() {
	fmt.Println("Hello Victoria Secrets")
}

func greetBy(name string) { // This is a function stack frame.
	fmt.Println("Hello", name)
}
func greetByM(msg, name string) { // This is a function stack frame.
	//func greetByM(msg string, name string) { // This is a function stack frame.
	fmt.Println(msg, name)
}

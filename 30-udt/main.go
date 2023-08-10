package main

import "fmt"

func main() {

	var num1 int = 100

	var num2 MyInt = 200

	var num3 MyInt2 = 300

	str1 := num2.ToString()
	str2 := MyInt(num1).ToString()
	str3 := MyInt(num3).ToString()

	fmt.Println("Sq of MyInt variable:", MyInt2(num2).Sq())
	fmt.Println("Sq of int variable:", MyInt2(num1).Sq())
	fmt.Println("Sq of MyInt2 variable:", num3.Sq())

	fmt.Println(str1, str2, str3)

	// int int64

	// var i1 int = 100
	// var i2 int64 = i1

	var s1 Square = 25.25
	fmt.Println("Area of Square s1:", s1.Area())
	fmt.Println("Perimeter of Square s1:", s1.Perimeter())
}

type Integer = int // Just an alias which is another name

type MyInt int // A new type which is User Defined Type
func (mi MyInt) ToString() string {
	return fmt.Sprint(mi)
}

type MyInt2 MyInt

func (mi2 MyInt2) Sq() int {
	return int(mi2 * mi2)
}

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}
func (s Square) Perimeter() float32 {
	return float32(4 * s)
}

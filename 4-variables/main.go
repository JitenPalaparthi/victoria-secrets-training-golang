package main

import "fmt"

type integer = int //typedef

func main() {

	// numbers
	// rune and byte

	var num1 integer = 100001213

	fmt.Println("Just another type if int which is integer alias of int", num1)

	var num2 rune = 'A' // if a char is given in single quote, the unicode value is assigned to the variable
	var num3 int32 = 'B'

	var num4 byte = 'A'
	var num5 uint8 = 'B'

	fmt.Println(num2, num3, num4, num5)

	var num6 rune = '的'

	fmt.Println(num6)

	fmt.Println("مرحبا بالعالم")

}

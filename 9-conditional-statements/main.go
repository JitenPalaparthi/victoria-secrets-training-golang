package main

import "fmt"

func main() {

	// check whethere a value is divisible by 2, 3 and 6
	num := 1231
	if num%2 == 0 {
		fmt.Println(num, "is divisible by 2")
	} else {
		fmt.Println(num, "is not divisible by 2")
	}

	if num := 12; num%2 == 0 && num%3 == 0 && num%6 == 0 { // can declare a variable in if statement
		num1 := 13
		fmt.Println(num, "is divisible by 2,3 and 6")
		fmt.Println(num1)
	} else {
		num1 := 14
		fmt.Println(num, "is not divisible by 2 and 3 and 6")
		fmt.Println(num1)
	}

	// if value < 50 --> Grade C

	// if value > 50 and < 100 then Grade B

	// If value >100 then Grade A

	if value := 55; value >= 0 && value < 50 {
		println("Grade C")
	} else if value > 50 && value < 100 {
		println("Grade B")
	} else {
		println("Grade A")
	}

}

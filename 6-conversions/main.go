package main

import "fmt"

func main() {

	var num1 uint8 = 38

	var num2 int = int(num1) // (int)num1

	fmt.Println(num2)

	var num3 int = 65001

	var num4 uint8 = uint8(num3)
	fmt.Println(num4)

	var str1 string = "H" // string is an array

	var num5 int = int(str1[0])

	fmt.Println(num5)

	var num6 int = 65

	var str2 string = string(num6)

	fmt.Println(str2)

	str3 := fmt.Sprint(num6) // as it is conversion from any datatype to the string

	fmt.Println(str3)

	ok := false

	str4 := fmt.Sprint(ok, num6, " Hello") // if one among the operands is string no space is addes. Hence if any space is required , it has to ve added manually

	fmt.Println(str4)

}

// 65001
// 233

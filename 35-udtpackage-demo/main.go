package main

import (
	"fmt"

	tostring "github.com/JitenPalaparthi/udt/tostring"
)

func main() {

	var num1 tostring.Cint = 100
	var num2 tostring.Cfloat = 100.25

	fmt.Println(num1.ToString())
	fmt.Println(num2.ToString())

}

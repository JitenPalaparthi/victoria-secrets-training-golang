package main

import (
	"demo/slices"
	"fmt"
)

func main() {

	var slice1 slices.SliceInt

	slice1 = slice1.Fill(10)

	fmt.Println(slice1)

}

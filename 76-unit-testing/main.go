package main

import (
	"demo/slices"
	"fmt"
)

func main() {

	var slice1 slices.SliceInt

	slice1 = slice1.Fill(10, 12, 32123, 12, 123, 123, 123, 123)

	fmt.Println(slice1)

	slice1 = slice1.InferDefaults(1)
	fmt.Println(slice1)

}

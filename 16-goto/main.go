package main

import "fmt"

func main() {

	i := 1
loop:
	if i%2 == 0 {
		goto printme
		//fmt.Println(i, "is an even number")
	}
back:
	i++
	if i <= 10 {
		goto loop
	}

printme:
	if i <= 10 {
		fmt.Println(i, "is an even number")
		goto back
	}
}

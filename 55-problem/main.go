package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := 100

	go func() {
		a = a + 100
		fmt.Println(" Go routine", a)
	}()

	b := a
	fmt.Println(b)
	runtime.Goexit()

}

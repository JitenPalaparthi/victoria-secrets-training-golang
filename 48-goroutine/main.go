package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go fmt.Println("Hello Victoria Secrets")

	go func() {
		c := add(10, 20)
		fmt.Println("Sum", c)
	}()

	go func() {
		i := 1
		for {
			fmt.Print(i, "Thread-1 ")
			i++
		}
	}()

	go func() {
		i := 1
		for {
			fmt.Print(i, " Thread-2")
			i++
		}
	}()
	go func() {
		i := 1
		for {
			fmt.Print(i, " Thread-3")
			i++
		}
	}()
	go func() {
		i := 1
		for {
			fmt.Print(i, "Thread-3 ")
			i++
		}
	}()

	fmt.Println(runtime.NumCPU())
	fmt.Println("Hello World")

	time.Sleep(time.Second * 10)

}

func add(a, b int) int {
	return a + b
}

// 1- to execute group of statemens as a go routine , need to give go infront of that group of statements.
// 2- main is also a go routine
// 3- no go routine waits for completion of execution of other go routines
// 4- order of of go routines is not guarenteed

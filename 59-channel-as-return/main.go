package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for v := range GenerateSquares(100) {
			fmt.Println("From channel-1-->", v)
		}
	}()

	go func() {
		for v := range GenerateSquares(200) {
			fmt.Println("From channel-2-->", v)
		}
	}()

	ch3 := GenerateSquares(100)
	ch4 := GenerateSquares(300)

	go func() {
		for v := range ch3 {
			fmt.Println("From ch3-->", v)
		}
	}()

	go func() {
		for v := range ch4 {
			fmt.Println("From ch4-->", v)
		}
	}()

	runtime.Goexit()
}

func GenerateSquares(num int) chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= num; i++ {
			ch <- i * i
		}
		close(ch)
	}()

	return ch
}

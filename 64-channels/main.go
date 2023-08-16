package main

import (
	"fmt"
	"strings"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := MultipleChanOut(ch1, ch2)

	go func() {
		ch1 <- "Victoria"
		ch2 <- "Secrets & Co"
	}()

	for val := range ch3 {
		fmt.Println(val)
	}

	ch5 := GenerateSquares(10)

	for val := range ch5 {
		fmt.Println(val)
	}

}

func MultipleChanOut(ch1, ch2 chan string) chan string {
	ch := make(chan string)
	go func() {

		ch <- <-ch1 + " " + <-ch2

		close(ch)
	}()

	return ch
}

func SingleToMultiple(ch1 chan string) (ch2 chan string, ch3 chan string) {
	ch2, ch3 = make(chan string), make(chan string)
	go func() {
		val := <-ch1
		vals := strings.Split(val, " ") // assuming that there is only one space
		ch2 <- vals[0]
		ch3 <- vals[1]
		close(ch2)
		close(ch3)
	}()
	return ch2, ch3
}

// func FanIn(ch ...chan int) {

// }

func GenerateSquares(num uint) chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= int(num); i++ {
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

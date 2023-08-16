package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {

		go func() {
			for i := 1; i <= 100; i++ {
				ch1 <- i
			}
			close(ch1)
		}()

		ch2 <- 200
		close(ch2)
		ch3 <- 300
		close(ch3)
	}()
	// go func() {
	// 	ch1 <- 100
	// 	ch2 <- 200
	// 	ch3 <- 300
	// }()
	// go func() {
	// 	ch1 <- 100
	// 	ch2 <- 200
	// 	ch3 <- 300
	// }()
	ch := VariadicChan(ch1, ch2, ch3)
	for val := range ch {
		fmt.Println(val)
	}

}

func VariadicChan(chs ...chan int) chan int {
	resultch := make(chan int)
	go func() {
		for _, ch := range chs {
			//go func() {
			sum := 0
			for val := range ch {
				sum += val
			}
			resultch <- sum
			//}()
		}
		close(resultch)
	}()

	return resultch
}

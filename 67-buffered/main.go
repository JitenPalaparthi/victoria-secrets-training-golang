package main

import "fmt"

func main() {
	ch := make(chan int, 2) //
	ch <- 100
	ch <- 200
	//ch <- 300
	fmt.Println(<-ch, <-ch)
	ch <- 300
	fmt.Println(<-ch, <-ch)

}

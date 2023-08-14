package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch chan int
	ch = make(chan int)
	wg := new(sync.WaitGroup)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 100; i++ {
			//fmt.Println("Sender-->")
			ch <- i * i
		}
		wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 100; i++ {
			fmt.Println("Receiver", <-ch)
		}
		wg.Done()
	}(wg)
	wg.Wait()
}

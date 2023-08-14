package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int) // Unbuffered channel has been instantiated
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		num := <-ch // receive a value from the channel
		fmt.Println(num)
		wg.Done()
	}(wg)
	//go func() {
	time.Sleep(time.Second * 5)
	ch <- 100 // Assign a value to the channel
	//}()
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

var num int
var mu *sync.Mutex
var wg *sync.WaitGroup

func main() {
	mu = new(sync.Mutex)
	wg = new(sync.WaitGroup)
	//notify := make(chan struct{})
	wg.Add(1)
	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go Incr(mu, wg)
		}
		wg.Done()
	}()

	//runtime.Goexit()
	wg.Wait()
	fmt.Println("Final value of num:", num)
}

func Incr(mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	num++
	//fmt.Println(num)
	mu.Unlock()
	wg.Done()
	//
}

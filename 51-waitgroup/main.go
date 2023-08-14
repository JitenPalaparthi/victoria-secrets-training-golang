package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)
	//wg := &sync.WaitGroup{}
	i := 1
	wg.Add(103)
	//wg.Add(2)
	go func() {
		fmt.Println("Hello World")
		wg.Done()
	}()
	go func() {
		for i = 1; i <= 111; i++ {
			//wg.Add(1)
			go func(num int) {
				if i%2 == 0 {
					fmt.Println(num, " is even")
				} else {
					fmt.Println(num, "is odd")S
				}
				//wg.Done()
			}(i)
		}
		wg.Done()
	}()
	// runtime.Goexit()

	//time.Sleep(time.Second * 5) // This is not a correct way
	wg.Wait()
}

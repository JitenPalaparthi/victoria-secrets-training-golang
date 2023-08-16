package main

import (
	"fmt"
	"runtime"
)

func main() {

	ch := make(chan bool)

	go func() {
	outer:
		for {
			select {
			case v, ok := <-ch:
				fmt.Println("Value from channel", v, "Status of the channel", ok)
				if !ok {
					break outer
				}
			}
		}
	}()
	ch <- true
	close(ch)
	runtime.Goexit()
}

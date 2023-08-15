package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := GetServer("Server-1")
	ch2 := GetServer("Server-2")
	ch3 := GetServer("Server-3")
	//timeout1 := time.After(time.Millisecond * 1)
	//timeout2 := MyTimeAfter(time.Millisecond * 1)
	count := 0
	for {
		select {
		case v1 := <-ch1:
			fmt.Println(v1)
			count++
		case v2 := <-ch2:
			fmt.Println(v2)
			count++
		case v3 := <-ch3:
			fmt.Println(v3)
			count++
		// case t1 := <-timeout1:
		// 	fmt.Println(t1, "timed out-1")
		// 	count++
		// case t2 := <-timeout2:
		// 	fmt.Println(t2, "timed out-2")
		// 	count++
		default:
			fmt.Println("Nothing happened. Just default one")

		}

		if count >= 3 {
			return
		}
	}

}

func GetServer(str string) chan string {
	ch := make(chan string)
	go func() {
		//time.Sleep(time.Millisecond * 1)
		ch <- str + "Server is running"
		close(ch)
	}()
	return ch
}

func MyTimeAfter(t time.Duration) chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(t)
		ch <- struct{}{}
	}()
	return ch
}

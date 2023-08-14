package main

import "fmt"

func main() {

	ch := make(chan int)
	notify := make(chan struct{})
	go receive(ch, notify)
	go sender(ch)
	go sender(ch)
	<-notify
	//time.Sleep(time.Second * 2)
}

func sender(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i * i
	}
	close(ch) // only sender can close a channel and once the channel is closed ..
	//cannot send data again so only once it has to be closed
}

func receive(ch <-chan int, notify chan struct{}) {
	// for val := range ch { // until the channel is closed
	// 	fmt.Println("receive-->", val)
	// }

	for {
		val, ok := <-ch
		if !ok {
			break
		} else {
			fmt.Println("receive-->", val)
		}
	}

	notify <- struct{}{}
}

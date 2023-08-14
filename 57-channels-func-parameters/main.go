package main

import "fmt"

func main() {

	ch := make(chan int) // bidirectional chann
	notify := make(chan struct{})
	go receiver(ch, notify)
	go sender(ch)
	<-notify
}

// readonly --> can only receive data from the channel but cannot write data to the channel
// writeonly  --> can send data to the channel but cannot receive that data
// bidirectional -> normal channel
func sender(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i * i
	}
}

func receiver(ch <-chan int, notify chan<- struct{}) {
	for i := 1; i <= 100; i++ {
		fmt.Println("received ", <-ch)
	}
	notify <- struct{}{}
}

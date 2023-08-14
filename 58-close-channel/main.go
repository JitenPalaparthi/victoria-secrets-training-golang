package main

import (
	"fmt"
	"sync"
)

var mu *sync.Mutex

func main() {

	ch := make(chan string)
	notify := make(chan struct{})
	isClosed := make(chan bool)
	go func() {
		isClosed <- false
	}()
	go receive(ch, notify, isClosed)
	go sender("sender-1 ", ch, isClosed)
	go sender("sender-2 ", ch, isClosed)
	<-notify
	//time.Sleep(time.Second * 2)
}

func sender(send string, ch chan<- string, isClosed <-chan bool) {
	// defer recoverme()
	for i := 1; i <= 10; i++ {
		closed := <-isClosed
		if !closed {
			ch <- send + fmt.Sprint(i*i)
		}
	}
	close(ch) // only sender can close a channel and once the channel is closed ..
	//cannot send data again so only once it has to be closed
}

func recoverme() {
	if p := recover(); p != nil {
		fmt.Println(p)
	}
}

func receive(ch <-chan string, notify chan struct{}, isClosed chan<- bool) {
	// for val := range ch { // until the channel is closed
	// 	fmt.Println("receive-->", val)
	// }
	// defer recoverme()
	for {
		val, ok := <-ch
		if !ok {
			isClosed <- true
			break
		} else {
			isClosed <- false
			fmt.Println("receive-->", val)
		}
	}

	notify <- struct{}{}
}

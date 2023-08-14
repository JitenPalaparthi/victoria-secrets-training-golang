package main

import "fmt"

// type Empty struct {
// }

// func (e *Empty) Reverse(str string) string {

// 	return ""
// }

func main() {

	ch := make(chan int)
	notify := make(chan struct{})

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i * i
		}
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println("receiver-->", <-ch)
		}

		notify <- struct{}{} // Instance of a struct; No data but an empty instance
	}()

	st := <-notify
	fmt.Println("Value received from empty struct", st)

}

// Use channel as a notifier

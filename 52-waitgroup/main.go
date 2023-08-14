package main

import (
	"custom/strings"
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	//wg.Add(1) // main is also included
	wg.Add(1)
	go func() {
		fmt.Println("Hello Vitoria Secrets & Co USA")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Hello Vitoria Secrets & Co India")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Hello Vitoria Secrets & Co Srilanka")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Hello Vitoria Secrets & Co Europe")
		wg.Done()
	}()
	wg.Add(1)
	go Greet(wg, "Hello Vitoria Secrets & Co World")

	//wg.Done() // Dobe main as well
	wg.Add(1)
	go func(wg *sync.WaitGroup, msg string) {
		rstr := strings.Reverse(msg)
		fmt.Println(rstr)
		wg.Done()
	}(wg, "Hello Victoria Secrets & Co")

	wg.Wait()
}

func Greet(wg *sync.WaitGroup, msg string) {
	fmt.Println(msg)
	wg.Done()
}

// Wrapper function with WaitGroup
func LReverse(wg *sync.WaitGroup, msg string) string {
	wg.Done()
	return strings.Reverse(msg)
}

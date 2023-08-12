package main

import "fmt"

func main() {
	defer fmt.Println("This is end of main")
	defer func() {
		defer fmt.Println("Where will I be called?")
		fmt.Println("Hello Victoria Secrets & Co USA")
	}()
	fmt.Println("Hello Victoria Secrets & Co")
	fmt.Println("Hello Victoria Secrets & Co Indian team!")
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
)

func main() {

	go fmt.Println("Hello Victoria Secrets & Co")
	go func() {
		for i := 1; i <= 1000; i++ {
			num := rand.Intn(10000)
			go IsEven(num)
		}
	}()

	runtime.Goexit()

}

func IsEven(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is Even")
	} else {
		fmt.Println(num, "is Odd")
	}
}

func Fatal(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}

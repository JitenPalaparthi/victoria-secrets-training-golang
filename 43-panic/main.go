package main

import "fmt"

func main() {
	//var num int
	//fmt.Println(10 / 0) // compiler uderstands that it is an error
	func() {
		func() {
			func() {
				//fmt.Println(10 / num)
				r := Divide(100, 0)
				fmt.Println(r)
			}()
			fmt.Println("Hey Victoria Secrets & Co!")
		}()
		fmt.Println("Thank God.. all functions executed")
	}()
	fmt.Println("Main--> I am done")
}

func Divide(num, divider int) int {
	return num / divider
}

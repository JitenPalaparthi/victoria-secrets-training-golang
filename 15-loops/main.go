package main

import "fmt"

func main() {
outer:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			for k := 1; k <= 5; k++ {
				if i == 5 {
					break outer
				}
				println("i:", i, "j:", j, "k:", k)
			}

		}
	}

	fmt.Println("It is continued ")

	for i, j := 1, 20; i <= j; i, j = i+3, j+2 {
		println("i:", i, "j:", j)
	}

}

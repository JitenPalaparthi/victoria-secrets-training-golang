package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// fmt.Println("Enter some text")
	// var str string
	// fmt.Scanln(&str)
	// fmt.Println(str)

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("Enter some data")

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	// }

	fmt.Println("Enter name")

	//var str string

	//fmt.Scanf("%s", &str)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println(name)

	var num int

	fmt.Println("Enter a number")

	fmt.Scanf("%d", &num)

	fmt.Println(num)
}

// If you are reading user input as a string which has spaces then use bufio.NewReader and read data from the keyboard

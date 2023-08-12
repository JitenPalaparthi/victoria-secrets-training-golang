package main

import (
	"demo/fileops"
	"fmt"
)

func main() {

	// bytes, err := os.ReadFile("data.txt")
	// if err != nil {
	// 	println(err)
	// }
	// fmt.Println("Count of Chars in the file:", len(bytes))

	fileName := new(string)
	*fileName = "data1.txt"
	// fileName1 := "data.txt"
	// fileops.GetFileCharCount(&fileName1)
	count, err := fileops.GetFileCharCount(fileName)
	if err != nil {
		//fmt.Println(err)
		// log error here
		// or do what ever the logic to be implemented
		fmt.Println("there seems to be a problem to read the file. Contact your admin")
	} else {
		fmt.Println("Char count:", count)
	}

}

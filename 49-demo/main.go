package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	//f, err := os.Create("demo.txt")
	//fmt.Println(err)
	//f.WriteString("Hello world")
	//os.ReadFile()

	getCh := func() chan string {
		ch := make(chan string)
		go func() {
			f, err := os.OpenFile("demo.txt", syscall.O_RDWR, 0644)
			if err != nil {
				fmt.Println(err)
			}
			//isLine := false
			line := ""
			//slice := make([]string, 0)
			for {
				b := make([]byte, 100)
				n, err := f.Read(b)
				fmt.Println(err, n)

				for _, v := range b {
					if string(v) == "\n" {
						//slice = append(slice, line)

						ch <- line
						fmt.Println("------------------------------------->")
						line = ""
					} else {
						line = line + string(v)
					}

				}
				if err != nil {
					ch <- line
					//slice = append(slice, line)
					close(ch)
					break
				}

				//fmt.Print(string(b))
			}
		}()
		return ch
	}

	//fmt.Println(slice)

	for line := range getCh() {
		fmt.Println(line)
	}

}

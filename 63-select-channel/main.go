package main

import (
	"fmt"
	"time"
)

func main() {

	// ch1 := GetSquare(10)
	// ch2 := GetCube(10)
	// count := 0
	// for {

	// 	select {
	// 	case sq := <-ch1:
	// 		fmt.Println(sq)
	// 		count++
	// 	case cbe := <-ch2:
	// 		fmt.Println(cbe)
	// 		count++
	// 	default:
	// 		fmt.Println("Nothing doing")
	// 	}
	// 	if count == 20 {
	// 		break
	// 	}
	// }

	ch3 := GetSquareT()
	ch4 := GetCubeT()
	t1 := time.After(time.Millisecond * 2)

	for {
		select {
		case v1 := <-ch3:
			fmt.Println(v1)
		case v2 := <-ch4:
			fmt.Println(v2)
		case <-t1:
			return
		default:
			fmt.Println("Doing nothing")
		}
	}

}

func GetSquare(num int) chan string {
	ch := make(chan string)

	go func() {

		for i := 1; i <= num; i++ {
			ch <- "Square Go Routine-->" + fmt.Sprint(i*i)
		}

	}()

	return ch
}

func GetCube(num int) chan string {
	ch := make(chan string)

	go func() {

		for i := 1; i <= num; i++ {
			ch <- "Cube Go Routine-->" + fmt.Sprint(i*i*i)
		}

	}()

	return ch
}

func GetSquareT() chan string {
	ch := make(chan string)

	go func() {

		for i := 1; ; i++ {
			ch <- "Square Go Routine-->" + fmt.Sprint(i*i)
		}

	}()

	return ch
}

func GetCubeT() chan string {
	ch := make(chan string)

	go func() {

		for i := 1; ; i++ {
			ch <- "Cube Go Routine-->" + fmt.Sprint(i*i*i)
		}

	}()

	return ch
}

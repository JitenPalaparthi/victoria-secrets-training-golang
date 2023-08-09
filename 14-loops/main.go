package main

import (
	"errors"
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println("\nEven number")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}

	println("\nEven numbers with continue")

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue // continue continues to the next iteration
		}
		print(i, " ")
	}

	// Check whether a number is prime number or not

	if ok, err := isPrime(0); err != nil {
		fmt.Println(err)
	} else {
		if ok {
			fmt.Println("0 is a prime number")
		} else {
			fmt.Println("0 is not")
		}
	}

	num := 1

	if ok, err := isPrime(num); err != nil {
		fmt.Println(err)
	} else {
		if ok {
			fmt.Println(num, " is a prime number")
		} else {
			fmt.Println(num, " is not a prime number")
		}
	}

	num = 29
	if ok, err := isPrime(num); err != nil {
		fmt.Println(err)
	} else {
		if ok {
			fmt.Println(num, " is a prime number")
		} else {
			fmt.Println(num, " is not a prime number")
		}
	}

	num = 32
	if ok, err := isPrime(num); err != nil {
		fmt.Println(err)
	} else {
		if ok {
			fmt.Println(num, " is a prime number")
		} else {
			fmt.Println(num, " is not a prime number")
		}
	}

	num = 123
	if ok, err := isPrime(num); err != nil {
		fmt.Println(err)
	} else {
		if ok {
			fmt.Println(num, " is a prime number")
		} else {
			fmt.Println(num, " is not a prime number")
		}
	}

	j := 1

	for {
		if j <= 10 {
			fmt.Println("Victoria Secrets-->", j)
		} else {
			break
		}
		j++
	}

}

// for loop

func isPrime(num int) (bool, error) {
	if num == 0 {
		return false, errors.New("num is zero;Hence cannot decide")
	}
	if num == 1 || num == -1 || num == 2 || num == -2 {
		return true, nil
	}
	i := 2 // initiliser is declared outside

	isPrime := true
	//var count uint8 = 0
	// for ; i <= num/2; i++ {
	// 	if num%i == 0 {
	// 		//println(num, " is not a prime number")
	// 		isPrime = false
	// 		//count++
	// 		//break
	// 		return isPrime, nil
	// 	}
	// }

	for i <= num/2 {
		if num%i == 0 {
			//println(num, " is not a prime number")
			isPrime = false
			//count++
			//break
			return isPrime, nil
		}
		i++
	}

	return isPrime, nil

}

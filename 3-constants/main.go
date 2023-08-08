package main

import (
	"fmt"
	"reflect"
)

// Usually constants are stored in data segment
// only at compile time the value to the const has to be evaluated

const PI float32 = 3.14

const SQUARE_OF_PI1 float32 = PI * PI // this is a valid const

const SQUARE_OF_PI2 = 3.14 * 3.14

var pi float32 = 3.14

//const SQUARE_OF_PI3 = pi * pi // Why it is not okay?

const (
	DAYS    = 7
	MONTHS  = 12
	HOUSRS  = 24
	MINS    = 60
	SECONDS = 60
)

func main() {

	//PI = 3.143

	fmt.Println(PI, SQUARE_OF_PI1, SQUARE_OF_PI2, reflect.TypeOf(SQUARE_OF_PI2))

}

// shorthand declartion
// type inference
// constatntss

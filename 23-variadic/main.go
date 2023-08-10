package main

import "fmt"

func main() {

	s1 := sum()                                                                                   // can pass no arguments
	s2 := sum(10, 20, 34, 34, 3, 343, 34, 334, 343, 434, 34343, 3, 43, 43, 43434, 43, 434, 34, 4) // can pass any number of arguments

	arr1 := [5]int{10, 11, 12, 13, 14}
	s3 := sum(arr1[:]...) // pass a array as a variadic argument

	slice1 := []int{10, 11, 12, 13, 14}
	s4 := sum(slice1...) // slice as a variadic argument

	fmt.Println(s1, s2, s3, s4)

	fmt.Println(slice1[0:2])
	fmt.Println(slice1[3:])
	slice2 := slice1 // 10,11,12,13,14
	// slice2 : 10,11,12,13,14
	slice2 = append(slice1[:2], slice1[3:]...) // 4
	fmt.Println(slice2)
	fmt.Println(slice1)

	//[10,11,12,13,14]
	//slice1[0:2] [-->10,11]
	//slice1[3:] 13,14

	// 10,11,13,14,14

}

// The variadic parameter must be the last parameter
// There should be only one variadic parameter for a function or method

func sum(args ...int) int {
	_sum := 0
	for _, v := range args {
		_sum += v
	}
	return _sum
}

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 100) // create a slice

	for i, _ := range slice1 {
		slice1[i] = rand.Intn(50) // assign values to the slice using rand
	}
	fmt.Println(slice1) // print the slice1

	mymap := make(map[int]uint16) // create a map as int and uint16

	for _, v := range slice1 { // iterating the slice1
		// In the map key is slice's value and value of map is count of that value in the slice
		val, ok := mymap[v] // map returns val and ok .. If ok is true that means already key exists
		if ok {
			mymap[v] = val + 1 // if key exits , val is incremented
		} else {
			mymap[v] = 1 // if key does not that means first time , then give 1 as a default value
		}
	}

	fmt.Println(mymap)
}

// func Delete(map, key) bool, error
// if map is nil return error
// if key exists and delete return true, nil
// if key does not exist return false and nil

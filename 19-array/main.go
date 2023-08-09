package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Short hand declarations of array
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{10, 11, 12, 13, 14} // The compiler gives the length during compilation

	var arr3 [5]int = arr1
	fmt.Println(arr1, arr2, arr3)

	//var arr4 [4]int = arr2 // The length of arrays are different
	var arr4 [4]int

	// copy function of array
	// This works if arr4 is smaller than or equal arr2
	for i := range arr4 {
		if len(arr2) <= len(arr4) {
			arr4[i] = arr2[i]
		}
	}

	// Multi dimensional arrays

	arr2d1 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("2D array")

	for _, arrd1 := range arr2d1 {
		for _, v := range arrd1 {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	fmt.Println("3D array")

	arr3d1 := [2][2][3]int{{{1, 2, 12}, {3, 4, 34}}, {{5, 6, 56}, {7, 8, 78}}}
	/* Outer most loop : 	{{1, 2}, {3, 4}}
							{{5, 6}, {7, 8}}

	 second outer loop: 	{1, 2}
	   					  	{3, 4}
						  	{5, 6}
					      	{7, 8}
	 inner most loop: 		1 2 3 4 5 6 7 8
	*/
	for _, arrd1 := range arr3d1 {
		for _, arrd2 := range arrd1 {
			for _, v := range arrd2 {
				fmt.Print(v, " ")
			}
		}
		fmt.Println()
	}

	arr4d1 := [2][2][2][3]int{{{{1, 2, 12}, {3, 4, 34}}, {{5, 6, 56}, {7, 8, 78}}}, {{{1, 2, 12}, {3, 4, 34}}, {{5, 6, 56}, {7, 8, 78}}}}
	fmt.Println(arr4d1)

	var arr4d2 [2][2][2][3]int

	for i1 := 0; i1 < len(arr4d2); i1++ {
		for i2 := 0; i2 < len(arr4d2[i1]); i2++ {
			for i3 := 0; i3 < len(arr4d2[i1][i2]); i3++ {
				for i4 := 0; i4 < len(arr4d2[i1][i2][i3]); i4++ {
					arr4d2[i1][i2][i3][i4] = rand.Intn(1000)
				}
			}
		}
	}
	fmt.Println("Array 4d")
	fmt.Println(arr4d2)
}

package main

import (
	"demo/activity"
	"demo/slices"
	"fmt"
)

//go:generate ls -a
func main() {

	var slice1 slices.SliceInt

	slice1 = slice1.Fill(10, 12, 32123, 12, 123, 123, 123, 123)

	fmt.Println(slice1)

	slice1 = slice1.InferDefaults(1)
	fmt.Println(slice1)

	a1 := new(activity.ActivityType)
	fmt.Println(a1.GetActivity("https://www.boredapi.com/api/activity"))

}

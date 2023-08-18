package main

import "encoding/json"

var i2 int

func main() {

	var i1 int = 100

	var ptr1 *int = &i1
	i2 = 200

	// arr1 := [3]int{10, 11, 12}

	// var iface1 any = arr1

	println(i1, i2, ptr1)
	// for _, v := range arr1 {
	// 	print(v, " ")
	// }

	s1 := new(Sample)
	s1.Name = "Victoria"
	s1.No = "123123"
	println(s1)

	// slice1 := []int{10, 12, 13, 14}
	// for i := 1; i <= 10000; i++ {
	// 	slice1 = append(slice1, 15, 16, 17, 18, 19)
	// }
	// for _, v := range slice1 {
	// 	print(v, " ")
	// }
	//println(arr1)
	//fmt.Println(i1, i2, ptr1, arr1)
	//log.Println(i1, i2, ptr1)
	println(sum(10, 20))
	println(sump(10, 20))

	bytes, _ := json.Marshal(s1)
	println(string(bytes))
	//json.Unmarshal()
	c := new(int)
	sumpi(10, 20, c)
	println(*c)
}

type Sample struct {
	No, Name string
}

func sum(a, b int) int {
	return a + b
}

func sump(a, b int) *int {
	c := a + b
	return &c
}

func sumpi(a, b int, c *int) {
	*c = a + b
}

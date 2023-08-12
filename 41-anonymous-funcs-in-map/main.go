package main

import "fmt"

func main() {

	var fnmap map[string]func(int, int) int

	fnmap = make(map[string]func(int, int) int)
	// fnmapexex := make(map[string]int)

	fnsub := func(a, b int) int {
		return a - b
	}

	fnmap["add"] = func(i1, i2 int) int {
		return i1 + i2
	}

	fnmap["sub"] = fnsub

	fnmap["multipy"] = Multiply

	fnmap["mod"] = Division{}.Mod

	// fnmapexex["addition"] = func(i1, i2 int) int {
	// 	return i1 + i2
	// }(10, 10)

	for k, fn := range fnmap {
		r := fn(100, 200)
		fmt.Println(k, ":", r)
	}

}

func Multiply(a, b int) int {
	return a * b
}

type Division struct {
	A, B int
}

func (d Division) Mod(a, b int) int {
	d.A = a
	d.B = b
	return d.A % d.B
}

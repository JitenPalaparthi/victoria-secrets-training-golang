package main

import (
	"fmt"

	shapes "github.com/JitenPalaparthi/demoshapes/shape"
)

func main() {

	r1 := shapes.Rect{L: 10.34, B: 13.45}
	a1 := r1.Area()
	r1.Display()
	fmt.Println("Area of Rect r1:", a1)

}

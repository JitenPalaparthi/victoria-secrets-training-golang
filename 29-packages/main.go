package main

import (
	"fmt"

	// "github.com/JitenPalaparthi/myshapes/shape"
	// cb "github.com/JitenPalaparthi/myshapes/shape/cuboid" // aliasing a package
	"shapes/shape"
	cb "shapes/shape/cuboid"
	_ "shapes/shape/demo"

	// "log"
	// mlog "github.com/JitenPalaparthi/log"
	rshapes "github.com/JitenPalaparthi/shapes/shape"
	rsquare "github.com/JitenPalaparthi/shapes/shape/square"
)

func main() {

	//cd.GetMessage()

	r1 := new(shape.Rect)
	r1.L = 10.123
	r1.B = 123.3

	a1 := r1.Area()
	fmt.Println("Area of r1:", a1)

	s1 := shape.Square{S: 25.24}

	a2 := s1.Area()
	p2 := s1.Perimeter()

	fmt.Println("Area of s1:", a2)
	fmt.Println("Perimeter of s1:", p2)
	c1 := &cb.Cuboid{L: 12.34, B: 12.56, H: 13.45}

	a3 := c1.Area()
	fmt.Println("Area of c1:", a3)

	a4 := rshapes.Area(10.12, 12.34)
	fmt.Println(a4)

	a5 := rsquare.Area(25.25)
	fmt.Println(a5)

	cr1 := cb.Circle{}
	fmt.Println(cr1)
	//cr1.r = 123.312
	cr1.X = 123.44
	// cr1 := cb.cirlce{R: 100.100}
	fmt.Println(cr1)
	cr1.Display()
}

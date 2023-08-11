package main

import (
	"fmt"
	"myshapes/interfaces"
	"myshapes/shape"
)

func main() {

	r1 := &shape.Rect{L: 10.1, B: 12.4}
	r2 := &shape.Rect{L: 10.3, B: 12.5}
	r3 := &shape.Rect{L: 10.8, B: 12.47}

	var s1 shape.Square = 24.9
	var s2 shape.Square = 25.1
	var s3 shape.Square = 25.5

	c1 := &shape.Cuboid{L: 10.1, B: 12.4, H: 13.4}
	c2 := &shape.Cuboid{L: 10.2, B: 12.6, H: 13.9}
	c3 := &shape.Cuboid{L: 10.8, B: 12.7, H: 13.6}

	// slice1 := make([]interfaces.IArea, 0)
	// slice2 := make([]interfaces.IPerimeter, 0)
	slice3 := make([]interfaces.IAreaPerimeter, 0)

	// slice1 = append(slice1, r1, r2, r3, s1, s2, s3, c1, c2, c3)
	// slice2 = append(slice2, r1, r2, r3, s1, s2, s3, c1, c2, c3)
	slice3 = append(slice3, r1, r2, r3, s1, s2, s3, c1, c2, c3)

	// for _, iface := range slice1 {
	// 	Area(iface)
	// }

	// for _, iface := range slice2 {
	// 	Perimeter(iface)
	// }

	for _, iface := range slice3 {
		AreaAndPetimeter(iface)
	}
}

func Area(iarea interfaces.IArea) { // interface as a parameter.It can be a form of dependency injecttion
	a := iarea.Area()
	fmt.Println("Area :", a)
}

func Perimeter(iperimeter interfaces.IPerimeter) {
	p := iperimeter.Perimeter()
	fmt.Println("Perimeter:", p)
}

func AreaAndPetimeter(iAreaPerimeter interfaces.IAreaPerimeter) {
	iAreaPerimeter.Display()
	Area(iAreaPerimeter)
	Perimeter(iAreaPerimeter)
	fmt.Println("----------------")
}

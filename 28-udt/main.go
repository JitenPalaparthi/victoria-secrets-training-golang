package main

import "fmt"

func main() {
	r1 := Rect{L: 120.43, B: 132.454}
	a1 := AreaOfRect(r1)
	fmt.Println("Area of Rect r1", a1)

	a2 := r1.Area()
	fmt.Println("Area of Rect r1", a2)
	p2 := r1.Perimeter()
	fmt.Println("Perimeter of Rect r1", p2)

	fmt.Println("Area and Perimter inside rect object", r1.A, r1.P)

	r2 := new(Rect)
	r2.L = 10.1
	r2.B = 12.4
	a3 := r2.Area()
	p3 := r2.Perimeter()
	fmt.Println("Area of Rect r2", a3)
	fmt.Println("Perimeter of Rect r2", p3)

	r3 := &Rect{L: 10.1, B: 12.4}
	a4 := r3.Area()
	p4 := r3.Perimeter()

	fmt.Println("Area of Rect r3", a4)
	fmt.Println("Perimeter of Rect r3", p4)
}

type Rect struct {
	L, B float32
	A, P float32

	// func Area(){

	// }
}

type Square struct {
	S float32
}

// To write method , need to add a receiver

func (r Rect) AreaBy() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r Rect) PerimeterBy() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

// AreaOfRect is a function
func AreaOfRect(r Rect) float32 {
	return r.L * r.B
}

func PerimeterOfRect(r Rect) float32 {
	return 2 * (r.L + r.B)
}

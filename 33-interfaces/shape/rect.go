package shape

import "fmt"

type Rect struct {
	L, B float32
}

func (r *Rect) Area() (a float32) {
	return r.L * r.B
}

func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

func (r *Rect) Display() {
	fmt.Println("Rect")
	fmt.Println("-----------------")
}

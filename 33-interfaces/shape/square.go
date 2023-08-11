package shape

import "fmt"

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return float32(4 * s)
}
func (s Square) Display() {
	fmt.Println("Square")
	fmt.Println("-----------------")
}

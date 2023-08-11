package shape

import "fmt"

type Cuboid struct {
	L, B, H float32
}

func (c *Cuboid) Area() float32 {
	return c.L * c.B * c.H
}

func (c *Cuboid) Perimeter() float32 {
	return 4 * (c.L * c.B * c.H)
}
func (c *Cuboid) Display() {
	fmt.Println("Cuboid")
	fmt.Println("-----------------")
}

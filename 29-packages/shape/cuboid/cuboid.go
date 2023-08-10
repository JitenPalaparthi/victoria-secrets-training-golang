package cuboid

import "fmt"

type Cuboid struct {
	L, B, H float32
}

func (c *Cuboid) Area() float32 {
	return c.L * c.B * c.H
}

// THis is unexported type
type circle struct {
	R float32
}

// Circle is exported
// r is not exported(unexported)
type Circle struct {
	r float32
	X float32
}

func (c Circle) Display() {
	c.display()
}

func (c Circle) display() { // unexported but can be accessed in the same package
	fmt.Println("This is a circle type")
}

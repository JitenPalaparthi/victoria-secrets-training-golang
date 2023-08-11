package interfaces

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}

type IDisplay interface {
	Display()
}

type IAreaPerimeter interface {
	IArea
	IPerimeter
	//Display()
	IDisplay
}

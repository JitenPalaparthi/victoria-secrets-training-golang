package shape

type Rect struct {
	L, B float32
	A, P float32
}

// type Square struct {
// 	S float32
// }

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

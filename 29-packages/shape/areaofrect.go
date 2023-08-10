package shape

func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

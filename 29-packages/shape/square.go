package shape

type Square struct {
	S float32
}

func (s *Square) Area() float32 {
	return s.S * s.S
}

func (s *Square) Perimeter() float32 {
	return 4 * s.S
}

package structs

// Circle represents a circle geo object
type Circle struct {
	base   float64
	height float64
}

// Area claculates the area of the circle
func (c Circle) Area() float64 {
	return c.base * c.height / 2
}

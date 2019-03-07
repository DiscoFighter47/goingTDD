package structs

// Rectangle represents a rectangle geo object
type Rectangle struct {
	width  float64
	height float64
}

// Area calculates the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

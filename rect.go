package life

// Rect ..
type Rect struct {
	Min Point
	Max Point
}

// Width ...
func (rectangle Rect) Width() int {
	return rectangle.Max.X - rectangle.Min.X + 1
}

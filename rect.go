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

// Height ...
func (rectangle Rect) Height() int {
	return rectangle.Max.Y - rectangle.Min.Y + 1
}

// WrapPoint ...
func (rectangle Rect) WrapPoint(point Point) Point {
	if point.X < rectangle.Min.X {
		point.X += rectangle.Width()
	}
	if point.X > rectangle.Max.X {
		point.X -= rectangle.Width()
	}

	if point.Y < rectangle.Min.Y {
		point.Y += rectangle.Height()
	}
	if point.Y > rectangle.Max.Y {
		point.Y -= rectangle.Height()
	}

	return point
}

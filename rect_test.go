package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectWidth(test *testing.T) {
	rectangle := Rect{
		Min: Point{X: 5, Y: 12},
		Max: Point{X: 23, Y: 42},
	}
	width := rectangle.Width()

	assert.Equal(test, 19, width)
}

func TestRectHeight(test *testing.T) {
	rectangle := Rect{
		Min: Point{X: 5, Y: 12},
		Max: Point{X: 23, Y: 42},
	}
	height := rectangle.Height()

	assert.Equal(test, 31, height)
}

func TestRectWrapPoint_withPointInsideRectangle(test *testing.T) {
	rectangle := Rect{
		Min: Point{X: 5, Y: 12},
		Max: Point{X: 23, Y: 42},
	}
	point := rectangle.WrapPoint(Point{X: 20, Y: 20})

	assert.Equal(test, Point{X: 20, Y: 20}, point)
}

func TestRectWrapPoint_withPointLessThanMinimum(test *testing.T) {
	rectangle := Rect{
		Min: Point{X: 5, Y: 12},
		Max: Point{X: 23, Y: 42},
	}
	point := rectangle.WrapPoint(Point{X: 0, Y: 0})

	assert.Equal(test, Point{X: 19, Y: 31}, point)
}

func TestRectWrapPoint_withPointGreaterThanMaximum(test *testing.T) {
	rectangle := Rect{
		Min: Point{X: 5, Y: 12},
		Max: Point{X: 23, Y: 42},
	}
	point := rectangle.WrapPoint(Point{X: 50, Y: 50})

	assert.Equal(test, Point{X: 31, Y: 19}, point)
}

func TestRectWrapPoints(test *testing.T) {
	rectangle := Rect{
		Min: Point{X: 1, Y: 2},
		Max: Point{X: 5, Y: 4},
	}
	points := []Point{
		Point{X: 3, Y: 3},
		Point{X: 0, Y: 3},
		Point{X: 6, Y: 3},
		Point{X: 3, Y: 1},
		Point{X: 3, Y: 5},
	}
	newPoints := rectangle.WrapPoints(points)

	wantNewPoints := []Point{
		Point{X: 3, Y: 3},
		Point{X: 5, Y: 3},
		Point{X: 1, Y: 3},
		Point{X: 3, Y: 4},
		Point{X: 3, Y: 2},
	}
	assert.Equal(test, wantNewPoints, newPoints)
}

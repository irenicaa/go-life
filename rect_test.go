package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectWidth(test *testing.T) {
	rect := Rect{
		Min: Point{X: 5, Y: 12},
		Max: Point{X: 23, Y: 42},
	}
	width := rect.Width()

	assert.Equal(test, 19, width)
}

func TestRectHeight(test *testing.T) {
	rect := Rect{
		Min: Point{X: 5, Y: 12},
		Max: Point{X: 23, Y: 42},
	}
	height := rect.Height()

	assert.Equal(test, 31, height)
}

func TestRectWrapPoint_withPointInsideRectangle(test *testing.T) {
	rect := Rect{
		Min: Point{X: 5, Y: 12},
		Max: Point{X: 23, Y: 42},
	}
	point := rect.WrapPoint(Point{X: 20, Y: 20})

	assert.Equal(test, Point{X: 20, Y: 20}, point)
}

func TestRectWrapPoint_withPointLessThanMinimum(test *testing.T) {
	rect := Rect{
		Min: Point{X: 5, Y: 12},
		Max: Point{X: 23, Y: 42},
	}
	point := rect.WrapPoint(Point{X: 0, Y: 0})

	assert.Equal(test, Point{X: 19, Y: 31}, point)
}

func TestRectWrapPoint_withPointGreaterThanMaximum(test *testing.T) {
	rect := Rect{
		Min: Point{X: 5, Y: 12},
		Max: Point{X: 23, Y: 42},
	}
	point := rect.WrapPoint(Point{X: 50, Y: 50})

	assert.Equal(test, Point{X: 31, Y: 19}, point)
}

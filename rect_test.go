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

package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointSetContains_withNonExistentPoint(test *testing.T) {
	points := PointSet{
		Point{X: 5, Y: 12}:  struct{}{},
		Point{X: 23, Y: 42}: struct{}{},
	}
	ok := points.Contains(Point{X: 10, Y: 10})

	assert.Equal(test, false, ok)
}

func TestPointSetContains_withExistentPoint(test *testing.T) {
	points := PointSet{
		Point{X: 5, Y: 12}:  struct{}{},
		Point{X: 23, Y: 42}: struct{}{},
	}
	ok := points.Contains(Point{X: 5, Y: 12})

	assert.Equal(test, true, ok)
}

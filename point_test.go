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

func TestPointSetAdd_withNonExistentPoint(test *testing.T) {
	points := PointSet{
		Point{X: 5, Y: 12}:  struct{}{},
		Point{X: 23, Y: 42}: struct{}{},
	}
	points.Add(Point{X: 10, Y: 10})

	wantPoints := PointSet{
		Point{X: 5, Y: 12}:  struct{}{},
		Point{X: 23, Y: 42}: struct{}{},
		Point{X: 10, Y: 10}: struct{}{},
	}
	assert.Equal(test, wantPoints, points)
}

func TestPointSetAdd_withExistentPoint(test *testing.T) {
	points := PointSet{
		Point{X: 5, Y: 12}:  struct{}{},
		Point{X: 23, Y: 42}: struct{}{},
	}
	points.Add(Point{X: 5, Y: 12})

	wantPoints := PointSet{
		Point{X: 5, Y: 12}:  struct{}{},
		Point{X: 23, Y: 42}: struct{}{},
	}
	assert.Equal(test, wantPoints, points)
}

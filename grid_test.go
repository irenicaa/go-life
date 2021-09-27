package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridToPoints(test *testing.T) {
	grid := ".O.\n..O\nOOO"
	shift := Point{X: 1, Y: 2}
	points := GridToPoints(grid, shift)

	wantPoints := PointSet{
		Point{X: 2, Y: 2}: struct{}{},
		Point{X: 3, Y: 3}: struct{}{},
		Point{X: 1, Y: 4}: struct{}{},
		Point{X: 2, Y: 4}: struct{}{},
		Point{X: 3, Y: 4}: struct{}{},
	}
	assert.Equal(test, wantPoints, points)
}

func TestGridToPoints_withComments(test *testing.T) {
	grid := "!comment #1\n!comment #2\n!\n.O.\n!..O\nOOO"
	shift := Point{X: 1, Y: 2}
	points := GridToPoints(grid, shift)

	wantPoints := PointSet{
		Point{X: 2, Y: 2}: struct{}{},
		Point{X: 1, Y: 3}: struct{}{},
		Point{X: 2, Y: 3}: struct{}{},
		Point{X: 3, Y: 3}: struct{}{},
	}
	assert.Equal(test, wantPoints, points)
}

func TestPointsToGrid(test *testing.T) {
	rectangle := Rect{
		Min: Point{X: 1, Y: 2},
		Max: Point{X: 5, Y: 4},
	}
	points := PointSet{
		Point{X: 2, Y: 2}: struct{}{},
		Point{X: 3, Y: 3}: struct{}{},
		Point{X: 1, Y: 4}: struct{}{},
		Point{X: 2, Y: 4}: struct{}{},
		Point{X: 3, Y: 4}: struct{}{},
	}
	grid := PointsToGrid(points, rectangle)

	wantGrid := ".O...\n..O..\nOOO..\n"
	assert.Equal(test, wantGrid, grid)
}

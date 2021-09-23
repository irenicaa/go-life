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
		Point{2, 2}: struct{}{},
		Point{3, 3}: struct{}{},
		Point{1, 4}: struct{}{},
		Point{2, 4}: struct{}{},
		Point{3, 4}: struct{}{},
	}
	assert.Equal(test, wantPoints, points)
}

func TestGridToPoints_withComments(test *testing.T) {
	grid := "!comment #1\n!comment #2\n!\n.O.\n!..O\nOOO"
	shift := Point{X: 1, Y: 2}
	points := GridToPoints(grid, shift)

	wantPoints := PointSet{
		Point{2, 2}: struct{}{},
		Point{1, 3}: struct{}{},
		Point{2, 3}: struct{}{},
		Point{3, 3}: struct{}{},
	}
	assert.Equal(test, wantPoints, points)
}

func TestPointsToGrid(test *testing.T) {
	rectangle := Rect{Min: Point{1, 2}, Max: Point{5, 4}}
	points := PointSet{
		Point{2, 2}: struct{}{},
		Point{3, 3}: struct{}{},
		Point{1, 4}: struct{}{},
		Point{2, 4}: struct{}{},
		Point{3, 4}: struct{}{},
	}
	grid := PointsToGrid(points, rectangle)

	wantGrid := ".O...\n..O..\nOOO..\n"
	assert.Equal(test, wantGrid, grid)
}

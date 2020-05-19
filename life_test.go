package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNeighborsForPoint(test *testing.T) {
	neighbors := NeighborsForPoint(Point{X: 1, Y: 2})

	wantNeighbors := []Point{
		Point{X: 2, Y: 3},
		Point{X: 2, Y: 2},
		Point{X: 2, Y: 1},
		Point{X: 1, Y: 3},
		Point{X: 1, Y: 1},
		Point{X: 0, Y: 3},
		Point{X: 0, Y: 2},
		Point{X: 0, Y: 1},
	}
	assert.Equal(test, wantNeighbors, neighbors)
}

func TestNeighborsForPoints(test *testing.T) {
	neighbors := NeighborsForPoints([]Point{
		Point{X: 2, Y: 3},
		Point{X: 6, Y: 1},
	})

	wantNeighbors := []Point{
		Point{X: 3, Y: 4},
		Point{X: 3, Y: 3},
		Point{X: 3, Y: 2},
		Point{X: 2, Y: 4},
		Point{X: 2, Y: 2},
		Point{X: 1, Y: 4},
		Point{X: 1, Y: 3},
		Point{X: 1, Y: 2},

		Point{X: 7, Y: 2},
		Point{X: 7, Y: 1},
		Point{X: 7, Y: 0},
		Point{X: 6, Y: 2},
		Point{X: 6, Y: 0},
		Point{X: 5, Y: 2},
		Point{X: 5, Y: 1},
		Point{X: 5, Y: 0},
	}
	assert.Equal(test, wantNeighbors, neighbors)
}

func TestCountSamePoints(test *testing.T) {
	points := []Point{
		Point{X: 2, Y: 3},
		Point{X: 2, Y: 2},
		Point{X: 2, Y: 1},
		Point{X: 0, Y: 3},
		Point{X: 1, Y: 1},
		Point{X: 0, Y: 3},
		Point{X: 1, Y: 1},
		Point{X: 0, Y: 1},
	}
	pointsCounters := CountSamePoints(points)

	wantPointsCounters := map[Point]int{
		Point{X: 1, Y: 1}: 2,
		Point{X: 0, Y: 3}: 2,
		Point{X: 0, Y: 1}: 1,
		Point{X: 2, Y: 3}: 1,
		Point{X: 2, Y: 2}: 1,
		Point{X: 2, Y: 1}: 1,
	}
	assert.Equal(test, wantPointsCounters, pointsCounters)
}

func TestPopulate(test *testing.T) {
	neighborsCounters := map[Point]int{
		Point{0, 0}: 1,
		Point{1, 0}: 2,
		Point{2, 0}: 3,
		Point{3, 0}: 2,
		Point{4, 0}: 1,

		Point{0, 1}: 1,
		Point{1, 1}: 1,
		Point{2, 1}: 2,
		Point{3, 1}: 1,
		Point{4, 1}: 1,

		Point{0, 2}: 1,
		Point{1, 2}: 2,
		Point{2, 2}: 3,
		Point{3, 2}: 2,
		Point{4, 2}: 1,
	}
	points := map[Point]struct{}{
		Point{1, 1}: struct{}{},
		Point{2, 1}: struct{}{},
		Point{3, 1}: struct{}{},
	}
	newPopulation := Populate(points, neighborsCounters)

	wantPopulation := map[Point]struct{}{
		Point{2, 0}: struct{}{},
		Point{2, 1}: struct{}{},
		Point{2, 2}: struct{}{},
	}
	assert.Equal(test, wantPopulation, newPopulation)
}

func TestPointsToGrid(test *testing.T) {
	rectangle := Rect{Min: Point{1, 2}, Max: Point{5, 4}}
	points := []Point{Point{2, 3}, Point{3, 3}, Point{4, 3}}
	grid := PointsToGrid(points, rectangle)

	wantGrid := "     \n *** \n     \n"
	assert.Equal(test, wantGrid, grid)
}

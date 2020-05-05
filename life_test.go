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
	pointCounters := CountSamePoints(points)

	wantPointCounters := map[Point]int{
		Point{X: 1, Y: 1}: 2,
		Point{X: 0, Y: 3}: 2,
		Point{X: 0, Y: 1}: 1,
		Point{X: 2, Y: 3}: 1,
		Point{X: 2, Y: 2}: 1,
		Point{X: 2, Y: 1}: 1,
	}
	assert.Equal(test, wantPointCounters, pointCounters)
}

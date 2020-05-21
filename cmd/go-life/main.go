package main

import (
	"fmt"

	life "github.com/irenicaa/go-life"
)

func main() {
	points := map[life.Point]struct{}{
		life.Point{X: 2, Y: 2}: struct{}{},
		life.Point{X: 3, Y: 3}: struct{}{},
		life.Point{X: 1, Y: 4}: struct{}{},
		life.Point{X: 2, Y: 4}: struct{}{},
		life.Point{X: 3, Y: 4}: struct{}{},
	}
	rectangle := life.Rect{
		Min: life.Point{X: 0, Y: 0},
		Max: life.Point{X: 80, Y: 24},
	}

	for {
		grid := life.PointsToGrid(points, rectangle)
		fmt.Println(grid)

		neighbors := life.NeighborsForPoints(points)
		neighbors = life.WrapPointsToRect(neighbors, rectangle)
		neighborsCounters := life.CountSamePoints(neighbors)
		points = life.PopulatePoints(points, neighborsCounters)
	}
}

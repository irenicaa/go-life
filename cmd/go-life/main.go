package main

import (
	"flag"
	"fmt"
	"time"

	life "github.com/irenicaa/go-life"
)

func main() {
	xMin := flag.Int("x-min", 0, "")
	xMax := flag.Int("x-max", 80, "")
	yMin := flag.Int("y-min", 0, "")
	yMax := flag.Int("y-max", 24, "")
	outDelay := flag.Duration("outDelay", 100*time.Millisecond, "")
	flag.Parse()

	points := map[life.Point]struct{}{
		life.Point{X: 2, Y: 2}: struct{}{},
		life.Point{X: 3, Y: 3}: struct{}{},
		life.Point{X: 1, Y: 4}: struct{}{},
		life.Point{X: 2, Y: 4}: struct{}{},
		life.Point{X: 3, Y: 4}: struct{}{},
	}
	rectangle := life.Rect{
		Min: life.Point{X: *xMin, Y: *yMin},
		Max: life.Point{X: *xMax, Y: *yMax},
	}

	for {
		grid := life.PointsToGrid(points, rectangle)
		fmt.Print(grid)

		neighbors := life.NeighborsForPoints(points)
		neighbors = life.WrapPointsToRect(neighbors, rectangle)
		neighborsCounters := life.CountSamePoints(neighbors)
		points = life.PopulatePoints(points, neighborsCounters)

		time.Sleep(*outDelay)
	}
}

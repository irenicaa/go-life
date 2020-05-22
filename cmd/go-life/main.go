package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	life "github.com/irenicaa/go-life"
)

func main() {
	gridX := flag.Int("grid-x", 0, "")
	gridY := flag.Int("grid-y", 0, "")
	xMin := flag.Int("x-min", 0, "")
	xMax := flag.Int("x-max", 80, "")
	yMin := flag.Int("y-min", 0, "")
	yMax := flag.Int("y-max", 24, "")
	outDelay := flag.Duration("outDelay", 100*time.Millisecond, "")
	flag.Parse()

	gridBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	gridShift := life.Point{X: *gridX, Y: *gridY}
	points := life.GridToPoints(string(gridBytes), gridShift)
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

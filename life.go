package life

// Point ...
type Point struct {
	X int
	Y int
}

// NeighborsForPoint ...
func NeighborsForPoint(point Point) []Point {
	neighbors := []Point{}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			neighbor := Point{X: point.X - dx, Y: point.Y - dy}
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

// NeighborsForPoints ...
func NeighborsForPoints(points []Point) []Point {
	allNeighbors := []Point{}
	for _, point := range points {
		neighbors := NeighborsForPoint(point)
		allNeighbors = append(allNeighbors, neighbors...)
	}

	return allNeighbors
}

// CountSamePoints ...
func CountSamePoints(points []Point) map[Point]int {
	pointCounters := map[Point]int{}
	for _, point := range points {
		newCounter := 0
		currentCounter, ok := pointCounters[point]
		if ok {
			newCounter = currentCounter + 1
		} else {
			newCounter = 1
		}

		pointCounters[point] = newCounter
	}

	return pointCounters
}

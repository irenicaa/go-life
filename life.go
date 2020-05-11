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
	pointsCounters := map[Point]int{}
	for _, point := range points {
		pointsCounters[point]++
	}

	return pointsCounters
}

// Populate ...
func Populate(points []Point, neighborsCounters map[Point]int) []Point {
	newPopulation := []Point{}
	for neighbor, counter := range neighborsCounters {
		switch counter {
		case 3:
			newPopulation = append(newPopulation, neighbor)
		case 2:
			for _, point := range points {
				if point == neighbor {
					newPopulation = append(newPopulation, neighbor)
					break
				}
			}
		}
	}
	return newPopulation
}

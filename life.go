package life

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
func NeighborsForPoints(points PointSet) []Point {
	allNeighbors := []Point{}
	for point := range points {
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

// PopulatePoints ...
func PopulatePoints(
	points PointSet,
	neighborsCounters map[Point]int,
) PointSet {
	newPoints := PointSet{}
	for neighbor, counter := range neighborsCounters {
		switch counter {
		case 3:
			newPoints.Add(neighbor)
		case 2:
			if points.Contains(neighbor) {
				newPoints.Add(neighbor)
			}
		}
	}

	return newPoints
}

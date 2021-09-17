package life

const (
	deadCell     = '.'
	liveCell     = 'O'
	lineBreak    = '\n'
	commentStart = '!'
)

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
func NeighborsForPoints(points map[Point]struct{}) []Point {
	allNeighbors := []Point{}
	for point := range points {
		neighbors := NeighborsForPoint(point)
		allNeighbors = append(allNeighbors, neighbors...)
	}

	return allNeighbors
}

// WrapPointsToRect ...
func WrapPointsToRect(points []Point, rectangle Rect) []Point {
	newPoints := []Point{}
	for _, point := range points {
		if point.X < rectangle.Min.X {
			point.X = point.X + rectangle.Width()
		}
		if point.X > rectangle.Max.X {
			point.X = point.X - rectangle.Width()
		}

		if point.Y < rectangle.Min.Y {
			point.Y = point.Y + rectangle.Height()
		}
		if point.Y > rectangle.Max.Y {
			point.Y = point.Y - rectangle.Height()
		}

		newPoints = append(newPoints, point)
	}

	return newPoints
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
	points map[Point]struct{},
	neighborsCounters map[Point]int,
) map[Point]struct{} {
	newPoints := map[Point]struct{}{}
	for neighbor, counter := range neighborsCounters {
		switch counter {
		case 3:
			newPoints[neighbor] = struct{}{}
		case 2:
			if _, ok := points[neighbor]; ok {
				newPoints[neighbor] = struct{}{}
			}
		}
	}

	return newPoints
}

// GridToPoints ...
func GridToPoints(grid string, shift Point) map[Point]struct{} {
	points := map[Point]struct{}{}
	x, y := 0, 0
	inComment := false
	for i := 0; i < len(grid); i, x = i+1, x+1 {
		switch grid[i] {
		case liveCell:
			if inComment {
				break
			}

			point := Point{X: x + shift.X, Y: y + shift.Y}
			points[point] = struct{}{}
		case lineBreak:
			x = -1
			if inComment {
				inComment = false
				break
			}

			y++
		case commentStart:
			inComment = true
		}
	}

	return points
}

// PointsToGrid ...
func PointsToGrid(points map[Point]struct{}, rectangle Rect) string {
	grid := ""
	for y := rectangle.Min.Y; y <= rectangle.Max.Y; y++ {
		gridLine := ""
		for x := rectangle.Min.X; x <= rectangle.Max.X; x++ {
			point := Point{X: x, Y: y}
			if _, ok := points[point]; ok {
				gridLine += string(liveCell)
			} else {
				gridLine += string(deadCell)
			}
		}

		grid += gridLine + string(lineBreak)
	}

	return grid
}

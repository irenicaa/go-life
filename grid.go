package life

const (
	deadCell     = '.'
	liveCell     = 'O'
	lineBreak    = '\n'
	commentStart = '!'
)

// GridToPoints ...
func GridToPoints(grid string, shift Point) PointSet {
	points := PointSet{}
	x, y := 0, 0
	inComment := false
	for i := 0; i < len(grid); i, x = i+1, x+1 {
		switch grid[i] {
		case liveCell:
			if inComment {
				break
			}

			point := Point{X: x + shift.X, Y: y + shift.Y}
			points.Add(point)
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
func PointsToGrid(points PointSet, rectangle Rect) string {
	grid := ""
	for y := rectangle.Min.Y; y <= rectangle.Max.Y; y++ {
		gridLine := ""
		for x := rectangle.Min.X; x <= rectangle.Max.X; x++ {
			point := Point{X: x, Y: y}
			if points.Contains(point) {
				gridLine += string(liveCell)
			} else {
				gridLine += string(deadCell)
			}
		}

		grid += gridLine + string(lineBreak)
	}

	return grid
}

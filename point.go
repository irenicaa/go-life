package life

// Point ...
type Point struct {
	X int
	Y int
}

// PointSet ...
type PointSet map[Point]struct{}

// Contains ...
func (points PointSet) Contains(point Point) bool {
	_, ok := points[point]
	return ok
}

// Add ...
func (points PointSet) Add(point Point) {
	points[point] = struct{}{}
}

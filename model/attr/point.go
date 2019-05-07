package model

import (
	"math"
)

// Point is 棋子坐标
type Point struct {
	X int
	Y int
}

func (a Point) GetManhattanDistance(b Point) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

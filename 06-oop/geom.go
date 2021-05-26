package geom

import (
	"math"
)

func Distance(x1, x2, y1, y2 float64) float64 {
	if x1 < 0 || x2 < 0 || y1 < 0 || y2 < 0 {
		return -1
	}
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

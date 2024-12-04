package wave

import (
	"math"
)

func Sample(t float64, f float64) (w float64) {
	w = math.Sin(2 * math.Pi * t / f)
	w = (w + 1) / 2

	return
}
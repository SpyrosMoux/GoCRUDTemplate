package util

import "math"

func Round(val float64) float64 {
	return math.Round(val*100) / 100
}

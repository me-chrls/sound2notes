package utils

import (
	"math"
)

func RadiansToDegrees(raidians float64) float64 {
	return raidians * (180 / math.Pi)
}
func DegreesToRadians(degrees float64) float64 {
	return degrees / (180 / math.Pi)
}

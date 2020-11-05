package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	previous := float64(0)

	for notCloseEnough(math.Abs(z - previous)) {
		previous = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func notCloseEnough(change float64) bool {
	return change != 0 && change > 0.000000000001
}

func main() {
	fmt.Println(Sqrt(2))
}

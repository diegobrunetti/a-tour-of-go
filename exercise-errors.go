package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqtr float64

func (err ErrNegativeSqtr) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(err))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqtr(x)
	}
	z := 1.0
	previous := float64(0)

	for notCloseEnough(math.Abs(z - previous)) {
		previous = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func notCloseEnough(change float64) bool {
	return change != 0 && change > 0.000000000001
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

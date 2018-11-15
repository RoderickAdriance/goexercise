package exercise

import (
	"fmt"
	"math"
)

func Swap(x, y *int) (*int, *int) {
	return y, x
}

func Fu() {
	sqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(sqrt(9))
}

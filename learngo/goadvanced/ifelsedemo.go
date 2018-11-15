package main

import (
	"fmt"
	"math"
)

func pp(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}
func main() {
	fmt.Println(
		pp(3, 2, 10),
		pp(3, 3, 20),
	)
}

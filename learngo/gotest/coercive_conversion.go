package main

import (
	"fmt"
	"math"
)

func main() {
	//i := 42
	//f := float64(i)
	//u := uint(f)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(f)
	var z int = int(f)
	fmt.Println(x, z)
}

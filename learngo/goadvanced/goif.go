package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	//if 不能写()
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, y, lim float64) float64 {
	//Pow计算x的y次方
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}

func main() {
	//fmt.Println(sqrt(-9),sqrt(math.E))
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))
}

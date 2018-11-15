package main

import (
	"fmt"
	"math"
)

func main() {
	f1, f2 := f2(), f2()
	for i := 0; i < 10; i++ {
		fmt.Println(f1(i), f2(i*-2))
	}
}

//方法返回一个 参数为int 返回值为int 的函数
func f2() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func f1() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(4, 5))
}

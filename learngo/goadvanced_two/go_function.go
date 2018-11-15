package main

import (
	"fmt"
	"math"
)

func km() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
}

//func closure 函数闭包
//闭包是一个函数值，它来自函数体的外部的变量引用。
// 函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	return func() int {
		return 5
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

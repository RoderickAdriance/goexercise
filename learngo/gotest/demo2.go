package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	//多个返回值
	a, b := swap("a", "b")
	fmt.Println(add(50, 60))
	println(a, b)
}

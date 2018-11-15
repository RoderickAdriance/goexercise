package main

import "fmt"

var i, j int = 4, 5

func split(sum int) (x, y int) {
	k := 5
	x = sum*5 + k
	y = sum - 2
	return
}

func main() {
	fmt.Println(split(10))
}

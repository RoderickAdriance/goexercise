package main

import "fmt"

type V struct {
	X int
	Y int
}

func main() {
	v := V{1, 2}
	v.X = 10
	fmt.Println(v)

	p := &v
	p.X = 1e9
	fmt.Println(v)
}

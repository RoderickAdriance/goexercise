package main

import "fmt"

//iota是常量的计数器，从0开始，组中每定义1个常量自动递增1
const (
	f = "A"
	b
	r = iota
	d
)

const (
	ee = iota
	g
	n
)

func main() {
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(g)
}

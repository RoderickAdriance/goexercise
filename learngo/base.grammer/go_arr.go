package main

import "fmt"

func main() {
	a1()
}

func a1() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

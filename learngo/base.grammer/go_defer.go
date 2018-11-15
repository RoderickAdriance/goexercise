package main

import "fmt"

func main() {
	def2()
}

func def2() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func df1() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

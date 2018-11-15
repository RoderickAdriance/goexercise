package main

import "fmt"

var q = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	ran2()
}

func ran2() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func ran1() {
	for i, v := range q {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

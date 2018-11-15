package main

import "fmt"

func main() {
s:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue s
			fmt.Println("asdf")
		}
	}
}

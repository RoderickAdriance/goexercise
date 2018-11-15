package main

import "fmt"

func main() {
	b()
}

func b() {
	for {
		fmt.Println("while loop!!")
	}

}

func a() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	for sum > 0 {
		fmt.Println(sum)
		sum -= 1
	}
}

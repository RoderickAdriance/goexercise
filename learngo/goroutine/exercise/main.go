package main

import "fmt"

func main() {
	a := 0
	point(&a)
	fmt.Println(a)

}

func point(a *int) {
	*a++
}

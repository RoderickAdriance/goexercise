package main

import "fmt"

func main() {
	var i int
	j := i //j也是int,自动推断:=
	fmt.Printf("j is type of %T", j)
}

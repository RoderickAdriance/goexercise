package main

import "fmt"

//变量定义可以包含初始值，每个变量对应一个。

var i, j int = 1, 2

func m() {
	var c, a, l = true, false, 9
	fmt.Println(c, a, l, j, i)
}

const (
	yu = "123"
	kh = len(yu)
	qq = 899
	bn = 10.6
)

func nn() {
	var z int = 67
	var mm string = string(z)
	fmt.Println(mm)
}

func main() {
	nn()
}

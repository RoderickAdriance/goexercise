package main

import "fmt"

//defer 语句会延迟函数的执行直到上层函数返回。

func d() {
	defer fmt.Println("defer invoke")
	fmt.Println("hello")
}

//延迟的函数调用被压入一个栈中。当函数返回时,会按照 '后进先出' 的顺序调用被延迟的函数调用。
func w() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	w()
}

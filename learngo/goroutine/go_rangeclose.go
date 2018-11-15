package main

import "fmt"

//发送者可以 close 一个 channel 来表示再没有值会被发送了
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//v, ok := <-ch   ok可以判断通道是否被关闭,关闭后ok 为false
func vb() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//循环 `for i := range c` 会不断从 channel 接收值，直到它被关闭
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	vb()
}

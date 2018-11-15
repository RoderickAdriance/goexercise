package main

import (
	"fmt"
	"time"
)

//channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。

//ch <- v    // 将 v 送入 channel ch。
//v := <-ch  // 从 ch 接收，并且赋值给 v。

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum //将 sum 送入c
}

func k() {
	a := []int{5, 6, -4, 6, 8, 10}
	c := make(chan int)
	go sum(a[:len(a)/2], c) //7
	go sum(a[len(a)/2:], c) //24
	x, y := <-c, <-c        //c取两次值, 就要往c中放入两次值
	fmt.Println(x, y)       //并发,不保证顺序

}

//为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：
//向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞。
func o() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(<-c)
	//fmt.Println(<-c)
}

func main() {
	o()
}

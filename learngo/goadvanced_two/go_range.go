package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func ra() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

//可以通过赋值给 _ 来忽略序号和值。
//如果只需要索引值，去掉“, value”的部分即可。
func ra2() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	//循环数组和字典都是 先索引或者key在 值
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	ra2()
}

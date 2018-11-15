package main

import "fmt"

//没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(20))
}

package main

import "fmt"

//类型 *T 是指向类型 T 的值的指针。其零值是 `nil`。 var p *int
//& 符号会生成一个指向其作用对象的指针。 i := 42  p = &i
//* 符号表示指针指向的底层的值。
func a() {
	i, j := 42, 222
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func main() {
	a()
}

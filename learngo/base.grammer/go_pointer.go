package main

import "fmt"

//类型 *T 是指向类型 T 的值的指针。其零值是 `nil`。

//& 符号会生成一个指向其作用对象的指针。
//* 符号表示指针指向的底层的值。
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

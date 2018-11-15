package main

import "fmt"

//类型 [n]T 是一个有 n 个类型为 T 的值的数组。
func l() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

//[]T 是一个元素类型为 T 的 slice。
func slice() {
	p := []int{2, 3, 4, 6, 7, 11, 23}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d]==%d\n", i, p[i])
	}
}

//slice 可以重新切片，创建一个新的 slice 值指向相同的数组。
//s[lo:hi]
func slice2() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])
}

//slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组：
func m() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := b[2:5]
	printSlice("d", d)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func slicezero() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

//func append(s []T, vs ...T) []T
//append 的第一个参数 s 是一个类型为 T 的数组，其余类型为 T 的值将会添加到 slice。
func sliceapp() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func main() {
	sliceapp()
}

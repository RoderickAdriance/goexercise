package main

import "fmt"

func main() {
	sliceAppend()
}

func sliceAppend() {
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

func sliceNil() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func s3() {
	a := make([]int, 5) // len(b)=0, cap(b)=5
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func s2() {
	p := []int{2, 3, 5, 7, 11, 13}
	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])

}

func s1() {
	p := []int{2, 3, 4, 5, 67}
	fmt.Println("p:", p)

	for i := 0; i < len(p); i++ {
		fmt.Println("p[%d] == %d\n", i, p[i])
	}
}

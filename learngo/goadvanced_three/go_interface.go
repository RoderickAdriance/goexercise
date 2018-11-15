package main

import (
	"fmt"
)

//接口类型是由一组方法定义的集合。
//接口类型的值可以存放实现这些方法的任何值。
type Abser interface {
	Fly()
}

type bird struct {
	n, a string
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func (b *bird) Fly() {
	fmt.Println(b.n + "can fly")
}

type Person struct {
	Name string
	Age  int
}

//String() 方法就是输出时调用, 相当于java toString()方法
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	//var w Writer
	//// os.Stdout 实现了 Writer
	//w = os.Stdout
	//fmt.Fprintf(w, "hello, writer\n")
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

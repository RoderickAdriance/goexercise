package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Ab() float64
}

func if1() {
	//接口类型
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vv{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser
	fmt.Println(a.Ab())
	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	//*Vertex实现了接口 Vertex并没有实现接口
	// 所以没有实现 Abser。
	//a = v

	//fmt.Println(a.Abs())
}

func main() {

}

func (f MyFloat) Ab() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vv struct {
	X, Y float64
}

func (v *Vv) Ab() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
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

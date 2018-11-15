package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//在结构体类型上定义方法
func (v *Vertex) sq() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (myf MyFloat) Abs() float64 {
	if myf < 0 {
		return float64(-myf)
	}
	return float64(myf)
}

//方法可以与命名类型或命名类型的指针关联。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	//v := &Vertex{3, 4}
	//fmt.Println(v.sq())
	//bb := MyFloat(-math.Sqrt2)
	//fmt.Println(bb.Abs())
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())

}

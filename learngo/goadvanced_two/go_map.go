package main

import "fmt"

type Position struct {
	Lat, Long float64
}

var x map[string]Position

//map 在使用之前必须用 make 而不是 new 来创建
func ma() {
	x = make(map[string]Position)
	x["site"] = Position{40, -36.8}
	fmt.Println(x["site"])
}

var v = map[string]Position{
	"a": Position{55, 99},
	"b": Position{37.4, -122},
}

//如果顶级的类型只有类型名的话，可以在文法的元素中省略键名
var q = map[string]Position{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

//同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
func revise() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	revise()
}

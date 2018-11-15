package main

import "fmt"

type Ver struct {
	Lat, Long float64
}

var m map[string]Ver

var mm = map[string]Ver{
	"a": Ver{40.1, 9},
	"b": Ver{11.21, 23.44},
}

//顶级的类型只有类型名的话，可以在文法的元素中省略键名。
var mmm = map[string]Ver{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m2()
}

func m1() {
	m = make(map[string]Ver)
	m["Bell Labs"] = Ver{9, 10}
	fmt.Println(m["Bell Labs"])
}

func m2() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	//如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
	//同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}

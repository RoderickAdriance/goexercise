package main

import "fmt"

//一个结构体（`struct`）就是一个字段的集合。

type Vertex struct {
	X, Y int
}

//结构体字段使用点号来访问。
func b() {
	vertex := Vertex{1, 6}
	fmt.Println(vertex.X)
	fmt.Println(vertex.Y)
}

//结构体字段可以通过结构体指针来访问。通过指针间接的访问是透明的。
func c() {
	vertex := Vertex{2, 5}
	v := &vertex
	//k没有取地址值不会改变
	k := vertex
	//x*10^8
	v.X = 2e8
	k.Y = 50
	fmt.Println(vertex)
}

//结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。
//使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
//特殊的前缀 & 返回一个指向结构体的指针。
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

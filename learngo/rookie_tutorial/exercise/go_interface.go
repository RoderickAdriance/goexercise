package exercise

import "fmt"

/*Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，
任何其他类型只要实现了这些方法就是实现了这个接口。*/

type Phone interface {
	call()
}

type IPhone struct {
}

func (p IPhone) call() {
	fmt.Println("IPhone 实现了call方法,所以实现了Phone接口")
}

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

//与 fmt.Stringer 类似，`error` 类型是一个内建接口：
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

//MyError 实现了 Error()interface 故可以返回
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		//输出run()的返回值error 的时候就会调用Error()方法
		fmt.Println(err)
	}
}

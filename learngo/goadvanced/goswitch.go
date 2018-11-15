package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	fmt.Println("Go runs on")
	//除非以 fallthrough 语句结束，否则分支会自动终止

	//获取操作系统环境
	switch a := runtime.GOOS; a {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
		fallthrough
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", a)
	}
}

func sw() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	//switch 不会向下穿过 ，如需使用fallthrough
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
		//fallthrough
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

//这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
func nocondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func main() {
	nocondition()
}

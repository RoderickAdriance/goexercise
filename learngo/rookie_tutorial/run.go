package main

import (
	"awesomeProject/learngo/rookie_tutorialtutorial/exercise"
	"fmt"
)

func main() {
	result, err := exercise.Sqrt(-8.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

var i int = 0

func recursion() {
	i++
	fmt.Println("递归函数", i)
	recursion()
}

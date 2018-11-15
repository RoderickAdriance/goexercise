package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//类似Java的toString方法
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	d := Person{"Danna", 18}
	a := Person{"aa", 897}
	fmt.Println(a, d)
}

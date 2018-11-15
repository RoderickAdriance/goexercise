package main

import "fmt"

//Go 只有一种循环结构——`for` 循环。
func main() {
	/*sum := 0
	for i:=0;i<10;i++{
		sum+=i
	}
	fmt.Println(sum)*/
	//跟 C 或者 Java 中一样，可以让前置、后置语句为空。
	/*sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)*/

	//for 是 Go 的 “while”
	for {
		fmt.Println("a")
	}

}

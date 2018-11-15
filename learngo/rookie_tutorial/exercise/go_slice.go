package exercise

import "fmt"

func SliceDemo() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	var numbers = make([]int, 3, 5)

	printSlice(numbers)

}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

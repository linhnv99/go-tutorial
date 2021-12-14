package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	//case1
	//fmt.Println(&a[0])
	//passArg(a...)
	//fmt.Println("Rs: ",a)

	//case2
	fmt.Println(&a[2])
	b := a[0:2]
	b = append(b, 1111)
	fmt.Println(&a[2])
	fmt.Println(&b[2])
	fmt.Println(b)
	fmt.Println(a)
}
func passArg(numbers ...int) {
	fmt.Println("Before: ",&numbers[0])
	numbers = append(numbers, 10000)
	fmt.Println("After: ", &numbers[0])
	fmt.Println("func: ", numbers)
}

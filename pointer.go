package main

import "fmt"

func main() {
	a := [5]int{1,2,3,4,5}
	//change1(&a) // không nên
	change2(a[:]) // nên
	fmt.Println(a)

}

func change1(a *[5]int) {
	a[0] = 1111
}
func change2(a []int) {
	a[0] = 1111
}


/*
	Pointer là biến lưu địa chỉ của 1 biến khác
	Note: nên sử dụng slice thay cho array làm đối số cho func
*/
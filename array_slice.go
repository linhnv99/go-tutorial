package main

import "fmt"

func main() {
	fmt.Println("============ARRAY===========")
	//var a [3]int
	//a := [...]int{1,2,3,4}
	a := [3]string{"PHP", "GO", "JAVA"}
	fmt.Println("root: ", a)
	//1. by pass value

	//b := a
	//b[0] = "ABC"
	//fmt.Println(a)
	//fmt.Println(b)

	//2. by pass reference
	//changeVal(&a)

	//3. iterating
	//3.1 for
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	//3.2 range => return key and value
	for i, v := range a {
		fmt.Printf("%d = %s \n", i, v)
	}
	fmt.Println("_ ignore field")
	for _, v := range a {
		fmt.Println("v = ", v)
	}

	fmt.Println("============Slice===========")
	s := a[0:2]
	fmt.Println(s)
	//s[0] = "ABC" // reflect to array
	//fmt.Println(a)

	//len() and cap() => tự động lới capacity khi tràn slice
	slices := make([]int, 1, 3) // crate slide with len = 1 and capacity = 3
	slices[0] = 10
	fmt.Println(slices, len(slices), cap(slices))
	slices = append(slices, 1, 2)
	fmt.Println(slices, len(slices), cap(slices))
	slices = append(slices, 4, 5)
	fmt.Println(slices, len(slices), cap(slices))

}

func changeVal(a *[3]string) {
	a[0] = "AHAHAH"
	fmt.Println("inside func", a)
}

/*
	Arrays là kiểu giả trị (không phải tham chiếu)
	_ ignore field

	- Slice same as JS
	a[start:end] : cắt từ vị trí start -> end - 1, [:] -> lấy tất cả
	Slice thì tham chiếu đến từng phần tủ của array

	len() -> số lượng phần tử trong arr
	cap() -> số lượng tham chiếu tới các phần tử trong array bắt đầu từ start

	- Slice same as Java ArrayList(default cap is 10)


	Note: Khai báo: slide k cần chỉ định size còn array thì cần
	+ array
		var arr = [3] int{1,2,3}
		arr := [3] int{1,2,3}
	+ slice
		var slice = [] int{1,2,3,4}
		slice := []int{1,2,3,4}
		slice := make(type, len, cap)
*/

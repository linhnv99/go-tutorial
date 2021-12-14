package main

import "fmt"

func main() {
	 //1.
	// a := make([]string, 1, 3)
	// a[0] = "A"
	//
	// fmt.Println(&a[0])
	//
	// fmt.Println("Append 1:")
	// a = append(a, "B")
	//
	// fmt.Println(&a[0])
	// fmt.Println(a)
	//
	//
	//fmt.Println("Append 2:")
	//a = append(a, "C", "D", "E")
	//
	//fmt.Println(&a[0])
	//fmt.Println(a)

	//2.
	a := []int{1,2,3,4}
	b := a[1:2]

	fmt.Println("Before Append: ")

	fmt.Printf("(A) => Len = %d and cap = %d \n", len(a), cap(a))
	fmt.Printf("(B) => Len = %d and cap = %d \n", len(b), cap(b))

	fmt.Println(&a[1])
	fmt.Println(&b[0])
	fmt.Println(a)
	fmt.Println(b)

	b = append(b, 12) // khi append vượt quá cap thì sẽ tạo 1 slice mới và copy sang => lúc này slice ban đầu sẽ ko bị update

	fmt.Println("After Append: ")
	fmt.Printf("(A) => Len = %d and cap = %d \n", len(a), cap(a))
	fmt.Printf("(B) => Len = %d and cap = %d \n", len(b), cap(b))
	fmt.Println(&a[1])
	fmt.Println(&b[0])
	fmt.Println(a)
	fmt.Println(b)
}
//Note: Giải thích về việc slice là dynamic
// khi vuợt qua cap ban đầu thì sẽ tự động tạo 1 slice mới với cap = cap * 2 và copy slice ban đầu -> slice mới
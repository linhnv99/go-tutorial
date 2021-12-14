package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	ID   int
	Name string
	Age  int
}


func main() {
	students := make([]Student, 0, 0)
	students = append(students, Student{1, "Linh", 22}, Student{2, "Lan", 20})

	fmt.Println("1. Add student")
	fmt.Println("2. Delete student")

	//var choice int
	//for ok := true; ok != false;{
	//	fmt.Println("> ")
	//	fmt.Scanf("%d", &choice)
	//	fmt.Println("choice", choice)
	//	switch choice {
	//	case 1:
	//		students = AddStudent(students[:])
	//		break
	//	default:
	//		fmt.Println("END")
	//		ok = false
	//		break
	//	}
	//}
}

func AddStudent(students []Student) []Student{
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input name: ")
	scanner.Scan()
	name := scanner.Text()

	var age int
	fmt.Println("Input age: ")
	fmt.Scanf("%d", &age)

	student := Student{len(students) + 1, name, age}
	return append(students, student)
}


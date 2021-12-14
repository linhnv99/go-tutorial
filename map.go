package main

import (
	"fmt"
	"reflect"
)

func main() {
	//khai bao
	//var maps map[int]string // map = nil


	languages := make(map[int]string)
	languages[0] = "PHP"
	languages[1] = "JAVA"
	languages[2] = "GOLANG"
	fmt.Println("root: ",languages)

	//check nil
	value, ok := languages[5]
	fmt.Println(value)
	fmt.Println(ok)

	//delete(map, key)
	delete(languages, 0)
	fmt.Println(languages)

	// reference type
	languages[0] = "ABC"
	fmt.Println(languages)

	//compare  (k dùng ==)
	map1 := map[int]int{
		0:0,
		1:1,
	}
	map2 := map[int]int {
		0:0,
		1:1,
	}
	fmt.Println(reflect.DeepEqual(map1, map2))


	// không đảm bảo thứ tự khi dùng for range
	//employeeSalary := map[string]int{
	//	"steve": 12000,
	//	"jamie": 15000,
	//	"mike":  9000,
	//}
	//fmt.Println("Contents of the map")
	//for key, value := range employeeSalary {
	//	fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	//}

	fmt.Println("NGHICH")
	b := []byte("I'm here")
	fmt.Println(b)
	s := string([]byte{73,39,109,32,104,101,114,101})
	fmt.Println(s)
}
/*
	Maps are reference types

*/
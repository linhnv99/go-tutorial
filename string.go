package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1,2,3,4}
	fmt.Println(reflect.TypeOf(a).Kind())


}
/*
	Strings are immutable

*/

package main

import (
	"errors"
	"fmt"
	"strconv"
)

type MyError struct {
	status int
	Err    error
}

func (e *MyError) Error() string {
	return "status: " + strconv.Itoa(e.status) + ", error: " + e.Err.Error()
}

func testChia0(a, b int) (interface{}, error) {
	if b != 0 {
		return a / b, nil
	}
	return nil, &MyError{400, errors.New("Loi chia cho 0")}
}

func main() {
	rs, error := testChia0(5, 0)
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Println("Result: ", rs)
}

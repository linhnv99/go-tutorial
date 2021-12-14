package main

import "errors"

func main() {

	a, err := chia(4, 5)
	if err != nil {

	}

}
func chia(a, b int) (int, error) {
	return 5, errors.New("Chia cho 0")
}

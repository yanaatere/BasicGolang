package main

import (
	"errors"
	"fmt"
)

/*
Handling Error
*/
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Pembagian Dengan 0")
	} else {
		return a / b, nil
	}
}
func main() {
	result, err := divide(100, 0)
	if err == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Error", err)
	}
}

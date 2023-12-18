package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func (n *validationError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation Error"}
	}

	if id != "yana" {
		return &notFoundError{"Data Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	/*
		Using If
		if err != nil {
			if validationErr, ok := err.(*validationError); ok {
				fmt.Println("Validation Error:", validationErr.Message)
			} else if notFoundErr, ok := err.(*notFoundError); ok {
				fmt.Println("Not found error:", notFoundErr.Error())
			} else {
				fmt.Println("Unknown Error", err.Error())
			}
		} else {
			fmt.Println("Suksess")
		}*/

	switch finalError := err.(type) {
	case *validationError:
		fmt.Println("Validation Error:", finalError.Error())
	case *notFoundError:
		fmt.Println("Not Found Error", finalError.Error())
	default:
		fmt.Println("Unknown Error", finalError.Error())
	}
}

package main

import "fmt"

/*
Interface Kosong dapat mengembalikan nilai apapun
Atau kalau di go versi terbaru itu pakai any
*/
func Ups(input int) interface{} {
	if input == 1 {
		return 1
	} else if input == 2 {
		return true
	} else {

		return "ups"
	}
}

func main() {
	var data = Ups(1)
	fmt.Println(data)
}

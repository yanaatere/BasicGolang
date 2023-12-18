package main

import "fmt"

/*
Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
*/
func random() any {
	return true
}
func main() {

	var result any = random()
	/*Ini cara kurang bagus karena bisa menyebabkan panic, kalau ada yang salah
	resultString := result.(string) //Manual Assertion, mungkin kalau di java ini ada casting
	fmt.Println(resultString)*/

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}

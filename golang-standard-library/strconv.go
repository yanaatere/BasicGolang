package main

import (
	"fmt"
	"strconv"
)

/*
Package Untuk Converter
Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string, atau sebaliknya
Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
https://golang.org/pkg/strconv/

*/

func main() {
	parseBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error parsing boolean", err)
	} else {
		fmt.Println("parse boolean", parseBool)
	}

	parseInt, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("error parsing boolean", err)
	} else {
		fmt.Println("parse Int", parseInt)
	}
}

package main

import (
	"fmt"
	"strings"
)

/*
Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String
Ada banyak sekali function yang bisa kita gunakan
https://golang.org/pkg/strings/

*/

func main() {

	fmt.Println(strings.Contains("Yana Andika", "Yana"))
	fmt.Println(strings.Split("Yana Andika", ""))
	fmt.Println(strings.ToLower("Yana Andika"))
	fmt.Println(strings.ToUpper("Yana Andika"))
	fmt.Println(strings.Trim(" Yana Andika ", " "))
	fmt.Println(strings.ReplaceAll("Yana Andika", "Yana", "Don"))
}

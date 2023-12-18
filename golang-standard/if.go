package main

import "fmt"

func main() {
	var name string = "Handoko"
	if name == "Yana" {
		fmt.Println("Benar Ini Yana")
	} else if name == "Handoko" {
		fmt.Println("Ini Handoko")
	}

	//Ini adalah if short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Berhasil Approve")

	}
}

package main

import "fmt"

func main() {
	name := "Yan"

	switch name {
	case "Yana":
		fmt.Println("Ini Benar Namaku", name)
	case "Handoko":
		fmt.Println("Ini Benar Namaku", name)
	default: //Kayak Else Di If
		fmt.Println("Semua Nama Diatas Salah")
	}

	//Switch Expression Short Statement
	switch length := len(name); length > 4 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Mencukupi")
	}

	//Switch Without Condition
	length := len(name)
	switch {
	case length < 3:
		fmt.Println("Kurang Dari 3")
	case length >= 3:
		fmt.Println("Lebih Dari 3")
	}

}

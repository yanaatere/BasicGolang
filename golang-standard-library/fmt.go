package main

import "fmt"

/*
https://pkg.go.dev/fmt
*/

func main() {
	firstName := "Yana"
	lastName := "Andika"

	//fmt.Println("Hello '", firstName, lastName, "'") //Biasanya kan kayak gini, tapi ini bakal menimbulkan spasi
	fmt.Printf("Hello %s %s \n", firstName, lastName)
	fmt.Printf("Hello '%s %s' \n", firstName, lastName)
}

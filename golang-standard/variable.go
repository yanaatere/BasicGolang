package main

import "fmt"

func main() {
	//Variable Apa bila tidak digunakan menjadi error

	// Deklarasi Multiple variable
	var (
		firstName = "Yana"
		lastName  = "Andika"
	)

	fmt.Println("First Name ", firstName)
	fmt.Println("Last Name ", lastName)

	//Dekalrasi awal tidak perlu memerlukan var, cukup menggunakan :=
	name := "Yana Andika"
	fmt.Println(name)

	//Dan cuma untuk di deklarasi awal
	name = "Yana Bedugul"
	fmt.Println(name)

	var age int8 = 38
	fmt.Println(age)
}

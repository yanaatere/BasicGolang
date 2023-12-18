package main

import "fmt"

/*Di golang, 1 package tidak boleh function yang sama*/

func fullName() (firstName string, middleName string, lastName string) {
	/*Untuk Variable tidak perlu di return, karena sudah anda deklarasikan diatas.
	Anda hanya perlu set valuenya saja*/
	firstName = "Yana"
	middleName = "Retere"
	lastName = "Andika"
	return
}

func main() {
	a, b, c := fullName()
	fmt.Println(a, b, c)
}

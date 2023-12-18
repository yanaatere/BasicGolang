package main

import "fmt"

func main() {
	/*Return Value Perlu ditambahkan*/
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	/*Menghiraukan return Value cukup tambahkan _*/
	one, _, _ := getFullName()
	fmt.Println(one)

}

func getFullName() (string, string, string) {
	return "Yana", "Andika", "Kanjurhan"
}

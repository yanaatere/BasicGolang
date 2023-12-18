package main

import "fmt"

func main() {
	//Buat Type Sebagai Alias
	type NoKTP string
	type Married bool

	//Ini contoh alias yang di pakai
	var noKTP NoKTP = "314062706950005"
	var marriedStatus Married = true

	fmt.Println(noKTP)
	fmt.Println(marriedStatus)

}

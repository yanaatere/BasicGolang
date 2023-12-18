package main

import "fmt"

/*
Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
Untuk melakukan ini, kita juga bisa menggunakan pointer di function
Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya

*/

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesian(address *Address) {
	address.Country = "Indonesian"
}
func main() {
	/*address := new(Address)
	ChangeCountryToIndonesian(address)
	fmt.Println("Address 1", address)*/

	address2 := Address{}
	ChangeCountryToIndonesian(&address2) //Change To Pointer

	fmt.Println("Address 2", address2)
}

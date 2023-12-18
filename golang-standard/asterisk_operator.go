package main

import "fmt"

/*
Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
Semua variable yang mengacu ke data yang sama tidak akan berubah
Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *

*/

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "West Java", "Indonesia"}
	var address2 *Address = &address1 //Ini berarti akan menjadi pointer, Nah pointer ini kayak di pemograman c. Jadi akan menghemat memory
	address2.City = "Bandung"
	fmt.Println("Address 1", address1) //Ikut Berubah
	fmt.Println("Address 2", address2) //Ikut Berubah Bandung

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}

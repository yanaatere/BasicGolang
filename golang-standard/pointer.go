package main

import "fmt"

/* Pass By Value
Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
sebenarnya yang dikirim adalah duplikasi value nya
*/

type Address struct {
	City, Province, Country string
}

func main() {
	/*address := Address{"Subang", "West Java", "Indonesia"}
	address2 := address //Berbeda dengan di java, kalau di java pasti akan me-replace si address 2 sedangkan di sini tidak teteap sama

	address2.City = "Bandung"
	fmt.Println("Address 1", address)
	fmt.Println("Address 2", address2)*/

	var address1 Address = Address{"Subang", "West Java", "Indonesia"}
	var address2 *Address = &address1 //Ini berarti akan menjadi pointer, Nah pointer ini kayak di pemograman c. Jadi akan menghemat memory

	address2.City = "Bandung"
	fmt.Println("Address 1", address1) //Ikut Berubah
	fmt.Println("Address 2", address2) //Ikut Berubah Bandung
}

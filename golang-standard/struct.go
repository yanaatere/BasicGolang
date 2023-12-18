package main

import "fmt"

/*
* Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
* Struct biasanya representasi data dalam program aplikasi yang kita buat
* Data di struct disimpan dalam field
* Sederhananya struct adalah kumpulan dari field
 */

/*Kalau Di Java Ini Ibarat Class*/
type Customer struct {
	Name, Address string // Dan ini variablenya dan selalu pascal case(Naming Selalu Paling awal)
	Age           int
}

/*Struct Method*/
func (customer Customer) sayHi(name string) { //Ini bukan lagi function tapi method
	fmt.Println("Hi", name, "My Name is ", customer.Name)
}

func (customer Customer) sayHuu() {
	fmt.Println("Say HUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUU", customer.Name)
}

func main() {
	/*Pembuatan Struct ada 3 cara*/
	var yana Customer

	//Cara Pertama
	yana.Name = "Yana"
	yana.Age = 28
	yana.Address = "Jalan Buah Batu"
	yana.sayHi("Wahid")
	yana.sayHuu()

	/*fmt.Println(yana.Name)
	fmt.Println(yana.Address)
	fmt.Println(yana.Age)

	//Cara Kedua
	wahid := Customer{Name: "Wahid", Address: "CIpenjo", Age: 28}
	fmt.Println(wahid.Name)
	fmt.Println(wahid.Address)
	fmt.Println(wahid.Age)*/

	/*Cara Ketiga Tidak Disarankan, walaupun cepat tapi sangat bergantung kepada komponen scruct*/
	yoga := Customer{"Yoga", "Metland", 28}
	yoga.sayHi("Wahid")
}

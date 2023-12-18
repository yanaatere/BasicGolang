package main

import "fmt"

/*
Berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe data tertentu,
maka secara otomatis akan dibuatkan default value nya

Tidak Semua TIpe Data Didikung, hanya interface, function, map, slice, pointer dan channel
Tipe Data NIL, yaitu tipe data kosong
*/

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = newMap("")

	/*
		Ini Pengecekkan lama
		if person["name"] == "" {
			fmt.Println("Data Kosong")
		} else {
			fmt.Println(person)
		}*/

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}

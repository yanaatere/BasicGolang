package main

import "fmt"

func main() {
	/*
		Function Slice
		len(slice) = Untuk Mendapatkan Panjang
		cap(slice) = Untuk Mendapatkan Kapasitas
		append(slice, data) = Membuat Slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas udah penuh , maka akan membuat array baru
		make([]TypeData, length, capacity) = Membuat slice baru
		copy(destination, source) = menyalin slice dari source ke destination
	*/

	/*
		Jika kamu tahu capcity of array gunakan angka
			var months = [12]
		jika tidak tahu bisa gunakan di bawah ini
		var months = [...]
	*/

	//var months = [...]string{
	//	"Januari",
	//	"February",
	//	"Maret",
	//	"April",
	//	"Mei",
	//	"Juni",
	//	"Juli",
	//	"Agustus",
	//	"September",
	//	"Oktober",
	//	"November",
	//	"Desember",
	//}

	//var slice = months[4:7]
	//fmt.Println(slice)      //Print Slice
	//fmt.Println(len(slice)) //Untuk Mendapatkan panjang Slice
	//fmt.Println(cap(slice)) //Untuk Mendapatkan Kapasitas Slice
	//
	///*Apabila kita merubah array, maka akan merubah juga slicenya*/
	//months[5] = "Diubah"
	//fmt.Println(slice)
	//
	///*Merubah slice juga pun merubah array*/
	//slice[0] = "Mei Lagi" //0 Di slice, adalah 4 di array. Karena index awal slice adalah 4
	//fmt.Println("Hasil setelah dirubah Slice 0 :", months)

	/*Start Of Append
	Tolong Command apabila ingin mencoba fungsi
	*/
	/*Append Dengan contoh tidak membuat array baru*/
	//var slice2 = months[10:] //Bila tidak ada angka sehabis : berarti sampai baris terakhir
	//fmt.Println("Hasil Slice 2 :", slice2)

	/*Append Dengan contoh  membuat array baru*/
	//var slice2 = months[2:4] //Bila tidak ada angka sehabis : berarti sampai baris terakhir
	//fmt.Println("Hasil Slice 2 :", slice2)
	//
	//var slice3 = append(slice2, "eko")
	//fmt.Println("Hasil Append Slice 3 : ", slice3)
	//
	//slice3[1] = "Bukan Desember"
	//fmt.Println("Rubah Slice 3[1] menjadi Bukan Desember : ", slice3)
	//
	//fmt.Println("Hasil Keseluruhan Slice 2 : ", slice2)
	//fmt.Println("Hasil Keseluruhan Months : ", months)
	/*End Of Append*/

	/*Make Slice*/
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Yana"
	newSlice[0] = "Andika"
	fmt.Println(newSlice)      //Print Slice
	fmt.Println(len(newSlice)) //Untuk Mendapatkan panjang Slice
	fmt.Println(cap(newSlice)) //Untuk Mendapatkan Kapasitas Slice
	/*End Of Make Slice*/

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("Hasil Copy Of Slice", copySlice)
}

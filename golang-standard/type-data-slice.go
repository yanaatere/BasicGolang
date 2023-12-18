// Tipe Data Slice Adalah Potongan Dari Array
// Slice mirip dengan array, tetapi ukuran slice bisa berubah
// slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian seluruh data di array
// Slice memiliki data yaitu pointer, length dan capacity
// Pointer adalah Penunjuk data pertama di array pada slice
// length adalah panjang dari slice
// Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

package main

import "fmt"

func main() {
	var months = [...]string{ //[...] apabila array belum tau jumlahnya berapa
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)

	//Function Slice
	fmt.Println(len(slice1)) //Hitung Panjang (Length)
	fmt.Println(cap(slice1)) //Hitung Cap

	//Panjang array month adalah 12, jika ada penambahan melebih panjang array makan akan membuat array baru
	//Seperti dibawah ini contohnya
	// var slice2 = months[10:]
	// fmt.Println(slice2)
	// var slice3 = append(slice2, "yana")
	// fmt.Println(slice3)
	// slice3[1] = "Bukan Desember"
	// fmt.Println("Slice3", slice3)
	// fmt.Println("Month", months)

	//Bandingkan Dengan Code Dibawah ini
	var slice2 = months[2:4]
	fmt.Println(slice2)
	var slice3 = append(slice2, "yana")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println("Slice3", slice3)
	fmt.Println("Month", months)

	//Make new Slice
	newSlice := make([]string, 2, 5)
	fmt.Println(newSlice)

	newSlice[0] = "Yana"
	newSlice[1] = "Andika"

	fmt.Println(newSlice)
	fmt.Println("Length ", len(newSlice))
	fmt.Println("Capacity ", cap(newSlice))
	fmt.Println("Nilai Pertama ", newSlice[0])

	//Copy Slice
	// Pastikan Ukurannya sama

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("Copy From New Slice ", copySlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Ini Array ", iniArray)
	fmt.Println("Ini Slice ", iniSlice)

}

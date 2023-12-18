package main

import "fmt"

func main() {
	//Daya Tampung Array tidak bisa bertambah setelah array dibuat
	//Kita juga perlu menentukan jumlah arraynya

	//[3] adalah jumlah maximum data array
	var names [3]string

	names[0] = "Yana"
	names[1] = "Andika"

	var nameYana = names[0]

	fmt.Println(nameYana)

	fmt.Println(names)
	fmt.Println(names[1])

	//Deklarasi Multiple Array
	var values = [3]int{
		90, 30, 50,
	}

	fmt.Println("Multiple Array", values)

	var panjangArray [10]string

	// Function array
	fmt.Println("Untuk Menampilkan Panjang Array", len(values))
	fmt.Println("Untuk Menampilkan Panjang Array", len(panjangArray))
}

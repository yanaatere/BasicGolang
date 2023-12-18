package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Yana",
		"address": "Jakarta",
	}

	/*Cara Akses Valuenya*/
	fmt.Println(person)
	fmt.Println(person["name"])

	//Penambahan Key Baru/ mengubah
	person["title"] = "Programmer"

	/*Cara Akses Valuenya*/
	fmt.Println(person)
	fmt.Println(person["title"])

	//Membuat Map Baru
	book := make(map[string]string)
	book["title"] = "Buku GO-Lang"
	book["author"] = "Yana"
	book["wrong"] = "ups"
	/*Untuk Menghapus Data di Map Key*/
	fmt.Println("Before Delete", len(book))
	delete(book, "wrong")
	fmt.Println("After Delete", len(book))
	fmt.Println(book)
}

package main

import "fmt"

func main() {
	// Constant Harus Diisi Set Variable, kalau tidak error
	// Constatnt Apa bila tidak dipakai tidak menimbulkan error
	const firstName string = "Yana"
	const lastName = "Andika"

	const (
		umur         int8 = 25
		jenisKelamin      = true
	)

	fmt.Println(jenisKelamin)

}

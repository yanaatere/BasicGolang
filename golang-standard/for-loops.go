package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan Ke :", counter)
		counter++
	}

	//For Short Statement
	for i := 0; i < 10; i++ {
		fmt.Println(i, "For Short Statement")
	}

	//Using Slice
	slice := [...]string{"Yana", "Andika", "Mantap"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//Using Hashmap
	person := make(map[string]string)
	person["Name"] = "Yana"
	person["Age"] = "25"

	//Ini adalah For Range
	for key, value := range person {
		fmt.Println("Key", key, "Value", value)
		if value == "Yana" {
			fmt.Println("Terdapat Sebuah Value bernama ", value)
		}
	}

	//Bisa Menggunakan "_" bila ada variable yang tidak dibutuhkan
	for _, value := range person {
		fmt.Println("Value", value)
		if value == "Yana" {
			fmt.Println("Terdapat Sebuah Value bernama ", value)
		}
	}

}

package main

import "fmt"

/*Untuk Mempersingkat Declaration Variable Bisa gunakan type declaration*/

type Filter func(string) string //Alias for String

func spamFilter(word string) string {
	if word == "Anjing" {
		return "..."
	} else {
		return word
	}
}

/*Before Using Type Declaration*/
/*func sayMyWordWithFilter(name string, spamFilter func(name string) string) string {
	nameFiltered := spamFilter(name)
	return "Hello, " + nameFiltered
}*/

func sayMyWordWithFilter(name string, spamFilter Filter) string {
	nameFiltered := spamFilter(name)
	return "Hello, " + nameFiltered
}

func main() {
	fmt.Println(sayMyWordWithFilter("Anjing", spamFilter))
}

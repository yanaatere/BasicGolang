package main

import "fmt"

/*Bisa Membuat Function, langsung menggunakan variable. Tanpa menggunakan nama*/

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're Blocked: ", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	/*Inilah Yang disebut anonymous function*/
	blackList := func(name string) bool {
		return name == "anjing"
	}

	/*Bisa juga seperti ini*/
	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("admin", blackList)
	registerUser("Yana", blackList)
}

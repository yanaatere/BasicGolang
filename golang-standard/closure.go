package main

import "fmt"

/*Sebuah Function berintaraksi data data yang di akses sebelumnya  */

func main() {
	name := "eko"
	counters := 0

	increment := func() {
		name = "budi"
		fmt.Println("Increment")
		counters++
		fmt.Println("name: ", name)
		fmt.Println("Counter: ", counters)
		//a := 1;
	}

	//fmt.Println(a); Inilah yang disebut closure karena a tidak bisa dikases dibawahnya, namun berbeda dengan name yang bisa dirubah
	increment()
	increment()
}

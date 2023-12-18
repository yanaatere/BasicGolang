package main

import "fmt"

func main() {

	/*
		Continue = Akan Skip
		Break = Akan Berhenti
	*/
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Println("Perulangan Ke", i)
	}
}

package main

import "fmt"

/*Sebaga First Class Citizen - Tidak dianak Tirikan*/
func getGoodBye(name string) string {
	return "Good Bye: " + name
}

func main() {
	sayGoodBye := getGoodBye /*Set Function To Another Variable atau bisa dijadikan value*/
	fmt.Println("From First Function: ", getGoodBye("Kemiskinan"))
	result := sayGoodBye("Masalah Hidup")
	fmt.Println("From Second Function: ", result)
}

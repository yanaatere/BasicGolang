package main

import (
	"fmt"
	"golang-basic/helper" //Ini cara import
)

func main() {
	result := helper.SayHello("Yana")
	fmt.Println(result)

	fmt.Println(helper.Application) //Bisa diakses
	fmt.Println(helper.version)     //Ini bakal jadi error, karena tidak bisa di akses
}

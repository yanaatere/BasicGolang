package main

import "fmt"

func loging() {
	fmt.Println("Selesai Memanggil Function")
}

func runApplication(value int) {
	/*Defer Function Adalah Function Yang bisa kita jadwalkan untuk di eksekusi setelah sebuah function selesai di eksekusi
	Defer Function akan selalu di eksekusi walaupun terjadi error di function yang di eksekusi
	Sebenernya juga bisa pake fmt.println tapi coba bandingkan memakai defer dan println.
	Defer function akan tetap di eksekusi walaupun ada error berbeda dengan println*/

	defer loging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("result: ", result)
}

func main() {
	runApplication(0)
}

package main

import "fmt"

/*
Panic Function adalah function yang biasa kita gunakan untuk menghentikan program
Panic Function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
Saat Panic function dipanggil, program akan terhenti, namun defer function akan tetap di eksekusi
*/

func endApp() {
	fmt.Println("End APP")
	/*Fungsi Recover adalah untuk mengambil message dari panic, seperti fitur try and catch dalam pemograman java*/
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error") //untuk menghentikan program
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Eko") //ini fungsinya recover
}

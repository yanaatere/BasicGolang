package main

import "fmt"

/*
Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
*/

type Man struct {
	Name string
}

func (man Man) single() {
	man.Name = "Unmarried " + man.Name
}

func (man *Man) Married() { //Ini adalah pointer dalam method
	man.Name = "Mr. " + man.Name
}
func main() {
	yana := Man{Name: "Yana"}
	yana.single()
	fmt.Println(yana) //Karena masih pass by value
	yana.Married()
	fmt.Println(yana)
}

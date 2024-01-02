package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

/*Harus Dibuat Dulu Kontraknya sebelum dapat bisa menggunakan package sort*/

type UserSlice []User

func (s UserSlice) Len() int { return len(s) }

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

/*End Of Contract*/
func main() {
	users := []User{
		{"Yana", 30},
		{"Budi", 35},
		{"Joko", 20},
		{"Adit", 25},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}

package main

import (
	"container/list"
	"fmt"
)

/*
Package container/list adalah implementasi struktur data double linked list di Go-Lang
https://golang.org/pkg/container/list/

*/

func main() {
	var data *list.List = list.New()

	data.PushBack("Yana")
	data.PushBack(12)
	data.PushBack("Andika")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) //Seharusnya Yana

	fmt.Println(head.Next().Value) //Ini bisa cara ngambil element ke dua

	for i := data.Front(); i != nil; i = i.Next() { //Ini cara Ngelihat semua element
		fmt.Println(i.Value)
	}

}

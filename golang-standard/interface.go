package main

import "fmt"

/*
Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
Sebuah interface berisikan definisi-definisi method
Biasanya interface digunakan sebagai kontrak
*/

type HasName interface { /*Interface juga adalah type data*/
	GetName() string /*GetName() adalaha sebuah method, string return valuenya*/
}

func sayMyName(name HasName) {
	fmt.Println("Hello", name.GetName())
}

type Person struct {
	name string
}

func (person Person) GetName() string {
	return person.name
}

type Animal struct {
	name string
}

func (animal Animal) GetName() string {
	return animal.name
}

func main() {
	var yana Person
	yana.name = "Yana Andika"
	sayMyName(yana)

	cat := Animal{name: "Push"}
	sayMyName(cat)
}

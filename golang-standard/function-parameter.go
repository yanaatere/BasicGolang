package main

import "fmt"

func main() {
	sayHelloTo("Yana", "Andika")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

package main

import "fmt"

func main() {
	sayHello()
}

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}
}

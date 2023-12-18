package main

import "fmt"

func main() {
	hello := getHello("Yana Andika")
	fmt.Println(hello)
}

/*
Return Value String
*/
func getHello(name string) string {
	if name == "" {
		return "hello"
	} else {
		return "Hello " + name
	}

}

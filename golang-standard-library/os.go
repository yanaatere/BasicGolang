package main

import (
	"fmt"
	"os"
)

/*

https://golang.org/pkg/os/
Args itu mengambil arguments
contohnya bisa running seperti ini
go run os.go Yana Andika
go run os.go "Yana Andika"

*/

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}

package main

import (
	"flag"
	"fmt"
)

/*
Package flag berisikan fungsionalitas untuk memparsing command line argument
https://golang.org/pkg/flag/

sebenernya juga bisa pakai args. tapi manual/
runningnya dengan cara ini
go run flag.go -username=Yana -password="rahasia" -host=123.231.23.3 -port 5505

*/

func main() {
	/*Dan value disamping adalah nilai default kalau value tidak dikirim*/
	username := flag.String("username", "root", "Database Username")
	password := flag.String("password", "root", "Database Password")
	host := flag.String("host", "localhost", "Database Host")
	port := flag.Int("port", 0, "Database port")

	flag.Parse()

	fmt.Println("Username: ", *username)
	fmt.Println("password: ", *password)
	fmt.Println("host: ", *host)
	fmt.Println("port: ", *port)

}

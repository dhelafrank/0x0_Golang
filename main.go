package main

import (
	"0x0_golang/lib"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	lib.Hello()
	lib.Variables()
}

// We can create binary files in go by simply running '~go build hello.go'
// or achitecture specific binary by running '~GOOS=windows OGARCH=amd64 go build hello.go'

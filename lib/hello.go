package lib

import "fmt"


func Hello() {
	fmt.Println("Hello World")
}

// We can create binary files in go by simply running '~go build hello.go'
// or achitecture specific binary by running '~GOOS=windows OGARCH=amd64 go build hello.go'
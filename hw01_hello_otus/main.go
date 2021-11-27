package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	helloString := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(helloString))
}

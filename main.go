package main

import (
	"fmt"
	"ksm/lexer"
)

func main() {
	l := lexer.New("Hello")
	fmt.Println(l)
}

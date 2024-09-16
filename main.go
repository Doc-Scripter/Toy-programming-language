package main

import (
	"fmt"
	"ksm/lexer"
	"ksm/token"
)

func main() {
	input := "1 = 1"
	l := lexer.New(input)

	for {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)
		if tok.Type == token.EOF {
			break
		}
	}
}
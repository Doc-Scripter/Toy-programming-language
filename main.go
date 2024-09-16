package main

import (
	"fmt"
	"os"

	"ksm/repl"
)

func main() {
	fmt.Println("Welcome to Kisumu Language :")
	repl.StartRepl(os.Stdin, os.Stdout)
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"ksm/repl"
	"os"

	"ksm/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Printf("Hello %s! This the Ksm Toy Programming language!\n", user.Username)
	fmt.Println()
	repl.StartRepl(os.Stdin, os.Stdout)
	// fmt.Println(lexer.New("Hello"))
}

	fmt.Println("Welcome to Kisumu Language :")
	repl.StartRepl(os.Stdin, os.Stdout)
}

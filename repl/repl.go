package repl

import (
	"bufio"
	"fmt"
	
	"io"

	"ksm/lexer"
	"ksm/token"
)

const REPL_PROMPT = "Enter Input: "

func StartRepl(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(REPL_PROMPT)

		scanned := scanner.Scan()
		// if it encounters the end of the input or an error.
		if !scanned {
			return
		}
		line := scanner.Text()

		L := lexer.New(line)

		for tok := L.NextToken(); tok.Type != token.EOF; tok = L.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

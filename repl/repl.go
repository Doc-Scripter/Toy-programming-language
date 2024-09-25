package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"ksm/lexer"
	"ksm/parser"
)

const REPL_PROMPT = "\033[32mEnter Input:\033[0m "

func StartRepl(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Println(REPL_PROMPT)

		scanned := scanner.Scan()
		// if it encounters the end of the input or an error.
		if !scanned {
			return
		}
		line := scanner.Text()

		L := lexer.New(line)
		p := parser.New(L)

		program := p.ParseProgram()
		checkParserErrors(p)

		for _, brnch := range program.Statements {
			fmt.Printf("%+++v\n", brnch)
		}
		fmt.Println()
	}
}

func checkParserErrors(p *parser.Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	fmt.Printf("parser has %d errors", len(errors))
	for _, msg := range errors {
		fmt.Printf("parser error: %q", msg)
	}
	os.Exit(1)
}

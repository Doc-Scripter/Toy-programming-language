package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"ksm/ast"
	"ksm/lexer"
	"ksm/parser"
)

const REPL_PROMPT = "\033[32mEnter Input:\033[0m "

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
		p := parser.New(L)

		program := p.ParseProgram()
		checkParserErrors(p)

		for _, stmt := range program.Statements {
			fmt.Println(formatStatement(stmt))
		}
		fmt.Println()
	}
}

func checkParserErrors(p *parser.Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	fmt.Printf("\033[31mparser has %d errors\033[0m\n", len(errors))
	for _, msg := range errors {
		fmt.Printf("\033[31mparser error: %q\033[0m\n", msg)
	}
	os.Exit(1)
}

func formatStatement(stmt ast.Statement) string {
	// If stmt is a known type like VarStement or ReturnStatement, formats accordingly
	switch stmt := stmt.(type) {
	case *ast.VarStatement:
		return fmt.Sprintf("Var Statement - Name: %s, value: %v", stmt.Name.Value, stmt.Value)
	case *ast.AssignmentStatement:
		return fmt.Sprintf("Assignment Statement - Name: %s, Value: %v", stmt.Name.Value, stmt.Value)
	case *ast.ReturnStatment:
		return fmt.Sprintf("Return Statement - Value: %v", stmt.ReturnValue)
	default:
		return fmt.Sprintf("Unkown Statement Type: %T", stmt)
	}
}

package parser

import (
	"testing"

	"ksm/ast"
	"ksm/lexer"
)

func TestVarStatements(t *testing.T) {
	input := `var x = 5;
	var y = 10;
	var foobar = 909090;`
	// var type = "string"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 4 statements. got= %d", len(program.Statements))
	}

	tt := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
		//{"type"},
	}

	for i, tc := range tt {
		stmt := program.Statements[i]
		if !testVarStatement(t, stmt, tc.expectedIdentifier) {
			return
		}
	}
}

func testVarStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not 'var'. got= %q", s.TokenLiteral())
		return false
	}

	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got= %T", s)
		return false
	}

	if varStmt.Name == nil {
		t.Errorf("varStmt.Name is nil")
		return false
	}
	if varStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value ot '%s'. got= %s", name, varStmt.Name.Value)
		return false
	}
	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got= %s", name, varStmt.Name)
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got= %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got= %T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

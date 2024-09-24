package parser

import (
	"ksm/ast"
	"ksm/lexer"
	"ksm/token"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

// (p *Parser)nextToken() advances both curToken and peekToken
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// (p *Parser)ParseProgram() constructs the root node of the AST, an `*ast.Program`.  
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseVarStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() *ast.Statement {
	switch p.curToken.Type {
	case token.VAR:
		return p.parseStatement()
	default:
		return nil
	}
}

// 
// () parseStatement () parses a statement
func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curToken}

	if !p.expectedPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}

	// TODO: we're skipping the expressions until
	// encounter a semicolon

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIS (t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectedPeek (t token.TokenType) bool {
	if p.peekTokenIS(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
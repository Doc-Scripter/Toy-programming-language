package ast

import "ksm/token"

/*A parser is a sotware component that takes input data (requently text) and builds
a data structure – oten some kind o parse tree, abstract syntax tree or other
hierarchical structure – giving a structural representation o the input, checking or
correct syntax in the process. */

// common interface for all AST nodes.
type Node interface {
	TokenLiteral() string
}

// represents the whole program
type Program struct {
	Statements []Statement
}

// interface representing a statement
type Statement interface {
	Node
	StatementNode()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// represents a variable declaration(let x = 5)
type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}

// represents variable names
type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string
}

// is an interface for all expressions
type Expression interface {
	Node
	ExpressionNode()
}

// represents interger values
type IntegerLiteral struct {
	Token token.Token // The token.INT token
	Value int64
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
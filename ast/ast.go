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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// interface representing a statement
type Statement interface {
	Node
	statementNode()
}

// is an interface for all expressions
type Expression interface {
	Node
	expressionNode()
}

// represents a variable declaration(var x = 5)
type VarStatement struct {
	Token token.Token // The token.VAR token
	Name  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) TokenLiteral() string {
	if vs.Token.Literal == "" {
		return "nil"
	}
	return vs.Token.Literal
}

// represents variable names
type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	if i.Token.Literal == "" {
		return "nil"
	}
	return i.Token.Literal
}

// represents interger values
// type IntegerLiteral struct {
// 	Token token.Token // The token.INT token
// 	Value int64
// }

type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

type AssignmentStatement struct {
	Name  string
	Value Expression
}

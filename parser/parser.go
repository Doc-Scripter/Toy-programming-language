package parser

import( "ksm/lexer-tokens")

type Parser struct{
	lexer *tokens.Lexer.Lexer
}

func New(l *lexer.Lexer) *Parser{
	p:= &Parser{lexer:l}
}
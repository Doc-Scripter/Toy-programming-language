package token

type TokenType string

// Token represents a token in the source code.
type Token struct {
	Type    TokenType // A type for different token categories.
	Literal string
}


// Different token types
const (
	// Special Tokens:
	ILLEGAL = "ILLEGAL" // Used when the lexer encounters a character or sequence it doesn't recognize
	EOF     = "EOF"     // "End of file" - Indicates the lexer has reached the end of the input

	// Identifiers and Literals:
	IDENTIFIER = "IDENTIFIER"
	NUMBER     = "NUMBER" // Represents integer literals: sequences of digits.
	// Operators:
	ASSIGN   = "ASSIGN" // Assignment operator, used for assigning values to variables.
	PLUS     = "PLUS"   // Used for adding numbers
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	FSLASH   = "/"

	// Comparison Operators;
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"

	// Delimiters:
	COMMA     = "," // Separates elements, such as parameters in function calls
	SEMICOLON = ";" // Marks the end of statements
	LPAREN    = "(" // Opening Parenthesis
	RPAREN    = ")" // Closing
	LBRACE    = "{" // Opening curly Braces
	RBRACE    = "}" // Closing Curly Braces

	// KeyWords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FOR      = "FOR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

// Maps keywords strings to their respective token types
var KeywordsMap = map[string]TokenType{
	"func":   FUNCTION,
	"var":    VAR,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
	"true":   TRUE,
	"false":  FALSE,
}

// Used by the lexer to determine if an identifier is a keyword or a variable
func LookupIdent(ident string) TokenType {
	if tok, ok := KeywordsMap[ident]; ok {
		return tok
	}
	return IDENTIFIER
}

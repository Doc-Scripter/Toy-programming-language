package lexer

// New() returns a new instance of the lexer with initialized state.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

// Readchar() gives us the next character and advances our position in the input string
func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // Ascii code for 'NUL', represents EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

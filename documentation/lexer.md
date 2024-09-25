# Lexer Package

## Overview
The `lexer` package is responsible for breaking down the input source code into tokens that can be used by the parser to create an Abstract Syntax Tree(AST). This process is also known as lexical analysis or tokenization.

The `Lexer` reads through the input string one character at a time and outputs individual tokens. Each token represents a meaningful component of the source code, such as an identifier, a keyword, or an operator.

## `Lexer` Structure
```go
    type Lexer struct {
        input           string
        position        int    //Current position in input (points to current char)\
        readPosition    int    // Current reading position in input (after current char)
        ch              byte   // Current character under examination
    }
```

+ `input`: The string of source code being tokenized.
+ `position`: The index of the current character in the input string.
+ `readposition`: The index of the next character in the input string.
+ `ch`: The current character being evaluated.

## Lexer Functions
1. `New(input string)` *Lexer
```go
func New(input sting) *Lexer
```

The `New` function initializes a new instance of the `Lexer`. It takes the input string (the source code) as an argument. The `readChar` method is called during initialization to load the first character into the lexer.

### Example Usage:
```go 
l := lexer.New("var x = 10;")```

```
2. `readCHar()`
```go 
func (l *Lexer) readChar()
```

`readChar` advances the lexer by one character in the input string. It moves the `position` to the current character and the `readPosition` to the next character.  If the `readPosition` exceeds the lenth of the input, the `ch` variable is set to `0`, indicating the end of file(EOF).

3. `NextToken() token.Token`
```go 
func (l *Lexer) NextToken() token.Token
```
This function is the core of the lexer, responsible for returning the next token from the input. The lexer skips any whitespace before examining the current character (`ch`) and determines what type of token to produce based on the character.

#### Token Types Handled:
+ **Operators & Symbol:**
    + `=` (assignment) and `==` (equality)
    + `+`,`-`,`*`,`/`(arithmetic operators)
    + `!`, (negation) and `!=` (inequality)
    + `<`,`>` (comparison operators)
    + `(`, `)`, `{`,`}` (grouping and scoping symbols)
    + `,`,`;`(punctuation)
 + **KeyWords:** The lexer differentiates keywords such as `var` and identifiers (variable names).
 + **Numbers:** Recognizes sequences of digits and returns them as number tokens.
 + **`EOF (End of File):** When the end of the input is reached, the lexer returns an EOF token.

 ### Example Usage:
 ```go
 tok := l.NextToken()
 fmt.Println(tok.Type, tok.Literal)
 ```

 4. `skipWhiteSpaces()`
 ```go
 func (l *Lexer) skipWhiteSpaces()
 ```
This helper functions skips over whitespace characters (`' '`,`\t`,`\n`,`\r`) by repeatedly calling `readChar()` until a non-whitespace character is found.

5. `readIdentifier() string`
```go
func (l *Lexer) readIdentifier() string
```
This function reads a sequence of letters (a valid identifier or keyword) starting from the current position. It continues advancing the lexer until a non-letter character is encounterd. The identified string is returned, and its type (identifier or keyword) is determined in `NextToken()`.

#### Example:
```go
varValue := l.readIdentifier() // could return "var" or "variableName"
```
6. `readNumber () string`
```go
func (l *Lexer) readNUmber() string
```
Similar to `readIdentifier`, this function reads a sequence of digits (a valid number) from the current position and continues advancing until a non-digit character is encountered.

#### Example:
```go
numValue := l.readNumber() // returns "10"
```

7. `peekCHar() byte`
```go
(l *Lexer) peekChar() byte
```
`peekChar` allows the lexer to look ahead one character without advancing the `position` or `readPosition`. It is used to identify two-character tokens like `==` and `!=`.

#### Example Usage:
```go
if l.peekChar() == '=' { 
    // handle ==
}
```
8. `newToken(tokenType token.TokenType, ch byte) token.Token`
```go 
func newToken(tokenType token.tokenType, ch byte) token.Token
```
A helper function that constructs and returns a `Token` struct given the token type and its literal character value.

#### Example Usage:
```go
tok := newToken(token.PLUS, '+')
```
### Helper Functions
1. `isLetter(ch byte) bool`
```go
func isLetter(ch byte) bool 
```
2. `isDigit(ch byte)bool`
```go
func isDigit(ch byte) bool 
```
This function checks if the given character is a digit (`0-9`).

### Token Type Mapping
The `NextToken()` method relies on the `token` package, which provides constants for various token types like `ASSIGN`, `PLUS`, `EQ`, `LT`, etc. These constants are defined as `token.TokenType` and represent the different components that the lexer recognizes.

### Example Usage of the Lexer
```go
package main

import (
    "fmt"
    "ksm/lexer"
    "ksm/token"
)

func main() {
    input := "var x = 5 + 10;"
    
    l := lexer.New(input)

    for tok := l.NextToken(); tok.Type != token.EOF; tok =l.NextToken() {
        fmt.Printf("%+v\n", tok)
    }
}
```

## Conclusion
The `lexer` package is responsible for converting a raw input string into a series of meaningful tokens that can be used by the parser to construct the Abstract Syntax Tree(AST). The lexer can handle various symbols, keywords, and identifiers as it scans the input character by character, skipping over whitespaces and identifying multi-character token such as `==` and `i=`.

[Back to main page](../README.md#documnentation)

[To Package REPL](/documentation/repl.md)


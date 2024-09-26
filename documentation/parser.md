# Parser Package

## Overview
The `parser` package is responsible for transforming tokens from the lexer into an Abstract Syntax Tree (AST). This forms the core of the parser's functionality in the `Kisumu` toy programming language. The parser follows a recursive descent approach, processing expressions and statements based on token types and sassociating prefix and infix parsing functions to handle different constrcts in the language.

## Package Contents
The `parser` package contains the following key components:
1. **Parser struct:** The structure which holds the state of the parsing process.
2. **Prefix and Infix Parse Functions:** These are function types that handle parsing operations for different kinds of tokens (prefix or infix).
3. **AST Construction:** The parser builds the AST by parsing tokens into statements and expressions.

## `Parser` Structure
### Fields:
+ `l`: This is the lexer from which the parser retrieves tokens.
+ `errors`: A list that stores parsing errors encountered during the parsing process.
+ `curToken` and `peekToken`: These represent the current and next tokens in the token stream.
+ `prefixParseFns` and `infixParseFns`: Maps of token types to corresponding parsing functions (prefix or infix).

### Key Methods
`New(l *lexer.Lexer) *Parser`

Thid function initializes a new parser with a given lexer, advancing the tokens twice to set the initial state for both `curToken` and `peekToken`.

`nextToken()`

Advances the current token and the peek token by fetching the next token from the lexer. This keeps the token stream in sync and is called frequently during parsing.

`ParseProgram() *ast.Program`

This method constructs the root node of the AST, which is an instance of `*ast.Program`. It loops over tokens, calling `parseStatement()` to convert each statement into  an AST node and adding it to the program's statement list.

`parseStatement() ast.Statement`

This method determines which type of statement to parse based on the current token. It can parse `var` declarations and `return` statements, but defaults to `nil` for other cases.

`parseVarStatement() *ast.VarStatement`

Parses a `return` statement. Similar to `parse
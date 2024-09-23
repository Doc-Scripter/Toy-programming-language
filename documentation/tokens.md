# Token Package

## Overview
The `token` package defines the basic building blocks for lexical analysis and parsing. It defines various token types and a `Token` structure, which represents individual components of the source code. Tokens are used by the lexer to categorize sections of the source code and by the parser to construct the Abstract Syntax Tree(AST).

Tokens fall into catergories such as operators, keywords, literals, and delimiters. The `token` package also contains a function to lookup identifiers and check if they are the keywords or variables.

## `Token` Structure
```go
type Token struct {
    Type TokenType // Defines the category/type of the token
    Literal string  // The actual value of the token as a string
}
```
+ `Type`: Represents the category of the token(e.g, `IDENTIFIER`, `NUMBER`, `ASSIGN`).
+ `Literal`: The actual string value from the source code (e.g., `"var"`,`"5"`, `"="`).

## `TokenType`
```go
type TokenTYpe string
```
`TokenType` is a string-based type used to define various types of tokens. The tokens represent different syntatic components, such as keywords, operators, or symbols.

### Token Types
1. **Special Tokens**

These tokens are used to represent non-code components.
  + `ILLEGAL`: Represents an unrecognized or invalid token encountered by the lexer.
  + `EOF`: End-og-file token, used to signal the end of the input.
  
  2. **Identifiers and Literals**

These tokens represent user-defined variables, functions names, and literal values.
  + `IDENTIFIER`: Represents variable names, function names, or other identifiers.
  + `NUMBER`: Represents numerical literals, such as integer values.
   3. **Operators**
   
   These tokens represent mathematical and logical operators.
   + `ASSIGN (`=`)`: Assignment operator, used for assigning values to variables.
   + `MINUS (`-`)`: Subtraction operator.
   + `BANG (`!`)`: Logical negation operator.
   + `ASTERISK (`*`)`: Multiplication operator.
   + `FSLASH (`/`)`: Division operator. 

   4. **Comparison Operators**
   These tokens are user for equality and relational comparisons.
   + `EQ (`==`)`: Equal operator.
   + `NOT_EQ (`!=`)`: Inequality operator.
   + `LT`(`<`)`: Less than operator.
   + `GT(`>`)`: Greater than operator.

   5. **Delimiters**

   These tokens are used for grouping code or separating elements.
   + `COMMA (`,`)`: Separates elements, such as function parameters.
   + `SEMICOLON (`;`)`: Marks the end of a statement.
   + `LPAREN(`(`)`: Left parenthesis, used for grouping expressions or function calls.
   + `RPAREN(`)`)`: Rigjt parethesis, closes an expression or function call.
   + `LBRACE(`{`)`: Left curly brace, used to define code blocks.
   + `RBRACE(`}`)`: Right curly brace, closes a code block.

   6. **Keywords**

   These are reserved words in the language that have specific syntactical meaning.
    + `FUNCTION (`func`)`: Used to define functions.
    + `VAR` (`var`)`: Used to declare variables.
    + `IF (`if`)`: Conditional statement.
    + `ELSE (`else`)`: Used in conjunction with `if` for alternate logic.
    + `RETURN (`return`)`: Used to return from functions.
    + `FOR (`for`)`: Used for loops.
    + `TRUE (`true`)`: Boolean literal.
    + `FALSE (`false`)`: Boolean literal.

## Keywords Mapping
```go
var KeywordsMap = map[string]TokenType{
    "func":     FUNCTION,
    "var":      VAR,
    "if":       IF,
    "else":     ELSE,
    "return":   RETURN,
    "for":      FOR,
    "true":     TRUE,
    "false":    FALSE,
}
```

The `KeywordsMap` provides a lookup mechanism for reserved keywords in the language. When the lexer encounters an identifier, it checks this map to determine if the identifier is a reserved keyword. If the identifier is found in the map, it returns the corresponding token type; otherwise, it defaults to the `IDENTIFIER` token type.

## Functions
1. `LookupIdent(ident string) TokenType`
```go
func LookupIdent(ident string) TokenType
```
The `LookupIdent` function is used by the lexer to determine if an identifier is a keyword or a variable. It checks the `KeywordsMap` to see if the string `ident` corresponds to a keyword. If it does, the function returns the corresponding keyword token type. If not, it returns the `IDENTIFIER` token type, treating the string as a user-defined name (e.g., a variable name).

#### Example Usage
```go
tokType := token.LookupIdent("func") // returns Function
tokType := token.LookupIdent("myVar") // returns IDENTIFIER
```
### Example Token Types and Usage

#### Identifiers and Keywords

When the lexer encounters a string of letters, it checks whether the string is a reserved keyword or a user-defined identifier. For example, For example, the string `func` is classified as the `FUNCTION` token, whereas `myFunction` would be classified as an `IDENTFIER`.

#### Operators 

Operators are classified as specific token types, such as `ASSIGN (`=`)`, `PLUS (`+`)`, and so on. These tokens are used by the parser to construct mathematical and logical expressions.

##### Delimiters

Delimeters such as parenthesese (`LPAREN`,`RPAREN`), and curly braces (`LBRACE`,`RBRACE`) help define the structure of the code, such as grouping expressions and defining blocks of code.

### Example Usage in Lexer

In the lexer, the token types are used to categorize different characters or sequences of characters from the input string. The lexer reads characters from the input and assigns them to the appropriate token tye.

for example, the lexer might convert the input:
```go
var x = 5;
```
into the following tokens:
<table>
  <thead>
    <tr>
      <th>Token Type</th>
      <th>Literal</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>VAR</td>
      <td>"let"</td>
    </tr>
    <tr>
      <td>IDENTIFIER</td>
      <td>"x"</td>
    </tr>
    <tr>
      <td>ASSIGN</td>
      <td>"="</td>
    </tr>
    <tr>
      <td>NUMBER</td>
      <td>"5"</td>
    </tr>
    <tr>
      <td>SEMICOLON</td>
      <td>";"</td>
    </tr>
  </tbody>
</table>

## Conclusion

The `token` package is a foundational part of the lexer and parser. It defines various token types, which represent different components of the source coe, such as operaors, keywords, literals, and delimiters. It also provides a looup mechanism for keywords, allowing the lexer to differentiate between user-defined identifiers and reserved keywords.

The `Token` structure, along with the predefined token types, ensures that that the lexer can effectively categorize the input into meaningful components, which the parser can use to build the Abstract Syntax Tree(AST).

[Back to Main Page](../README.md#documnentation)

[Move to Package Lexer](/documentation/lexer.md#overview)
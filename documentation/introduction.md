# Project Progress Journal

Overview of Ksm

Ksm is an interpreted toy programming language written in Go. The design draws inspiration from Go's syntax, and the project is currently focused on setting up fundamental building blocks like tokenization (lexer), parsing (AST), and token definitions.

## File Structure:
### Abstract Syntax Tree(AST)
+ **Purpose:** The code defines the basic structure of representing the Abstract Syntax Tree (AST) in Ksm. ASTs represent the structure of the course code in a tree form and help break down the code for interpration or compilation.

+ **Key Componensts:**
    + **Node Interface:** Provides a `TokenLiteral()` method, which every AST node implements to return the literal value of the token.
    + **Program Struct:** Represents the entire program, which is essentially a list of statements. The `TokenLiteral()` function returns the literal of the first statement.
    + **LetStatement Struct:** Represents a variable declareation (e.g., `let x = 5;`), containing the token for `let`, a variable name (`Identifier`), and its value (`Expression`).
    + **Expression Interface:** Represents any expression in the language, including integer literals (`IntegerLiteral`).

 ### Lexer
+ **Purpose:** Converts the input source code into a stream of tokens that will be processed into an Abstract Syntax Tree (AST).

+ **Key Components:**
    + **Lexer Struct:** Manages the input source code, tracks positions, and handles the current character (`ch`).
    + **Function New():** Initializes the lexer and reads the first character.
    + **readChar Method:** Reads the next character and convets it into a token (e.g., `=`, `+`, `(`,`{`). It also identifies keywords like `if`, `var`, or operators like `==`, `!=`.
    + **peekChar Method**: Peeks at the next character without advancing the read position.
    + **skipWhiteSpace Method:** Skips over whitespce (spaces, tabs, newlines).

### Token
 + **Purpose:** Defines the structure of a token and the types of the tokens available in Ksm.
 + **Key Components:**
    + **Token Struct:** Represents a token with its type and literal value.
    + **Token Types :** Predefined constants for various types of tokens like `ILLEGAL`,`EOF`, `IDENTIFIER`,`NUMBER`, `ASSIGN`, `PLUS`, `EQ`, `LT`, `GT`, etc.
        + **KeywordsMap:** Maps keywords (e.g., `func`, `var`, `if`) to their respective token types.

    ### Main Function
 + **Purpose:** Demonstrates the creation of the a lexer instance and token scanning.
 + **Key Components:** 
    + **main Function:** Inializes the lexer witha simple input (`"Hello"`) and prints out the the lexer instance for inspection.

    ### Parser (Work in Progress)
+ **Purpose:** Converts the stream of tokens produced by the lexer into an Abstract Syntax  Tree(AST), which can then be interpreted or compiled.
+ **Key Components:** Advances the current token to the next one in the input.
    + **parseStatement Method:** Placeholder for parsing statements. This is where various langauage constructs (like variable declaration or expressions) will be parsed and converted into AST nodes.

+ **Work in Progress:**
    We have the basic structure for a parser, which will process tokens from the lexer and convert them into an Abstract Abstract Tree(AST). The parser's job will eventually include handling expressions, statements, and control flow.

## Next Steps:
1. **Parser Enhancement:**
    + Implementation of parsing logic for expressions(arithmetic, comparison, etc.) and statements (like variable declarations, assignments).
    + Starting with simple statements like `var` declations and gradually adding support for functions, conditionals(`id/else`), loops(`for`) etc.
2. **Interpreter or Code Generation:**
    + Once parsing is complete, we will create an interpreter to evaluate the AST or a code generator to compile it to bytecode or Go code
3. **Data Structrue Implementation:**
    + Flesh out the details for the `array`, `struct`, and `hash` data structures. We'll need methods to manipulate them, such as adding/ removing elements, accessing members, etc.
4. **Expand Tests:**
    + Testing different cases(valid and invalid code) will help in catching issues early, 
    + Lexer tests are already in place.

Click [here](../README.md#documnentation) to naviagate back to the [homepage](../README.md#documnentation)

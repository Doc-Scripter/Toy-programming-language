# Project Progress Journal

Here's a brief overview of what's in the code so far

## Features of Ksm
1. **Ksm language Design:**Ksm, with a sintax inspired by Go, currently focuse on setting up its fundamental  building block: `tokenization(lexer)`, `parsing(AST)`, and `token definations`

## File Structure:
### Abstract Syntax Tree(AST)
+ **Purpose:** The code defines the basic structure ofr representing the Abstract Syntax Tree (AST) in Ksm. ASTs represent the structure of the course code in a tree form and help break dwn the code for interpration or compilation.

+ **Key Componensts:**
    + **Node Interface:** Provides a `TokenLiteral()` method, which every AST node implements to return the literal value of the token.
    + **Program Struct:** Represents the entire program, which is essentially a list of statements. The `TokenLiteral()` function returns the literal of the first statement.
    + **LetStatement Struct:** Represents a variable declareation (e.g., `let x = 5;`), containing the token for `let`, a variable name (`Identifier`), and its value (`Expression`).
    + **Expression Interface:** Represents any expression in the language, including integer literals (`IntegerLiteral`).

    ### Lexer
+ **Purpose:** Converts the stream of tokens produced by the lexer into an Abstract Syntax Tree (AST), which can  then be interpreted or compiled.

+ **Key Components:**
    + **Lexer Struct:** Manages the input source code, tracks positions, and handles the current character (`ch`).
    + **Function New():** Initializes the lexer and reads the first character.
    + **readChar Method:** Reads the next character and convets it into a token (e.g., `=`, `+`, `(`,`{`). It also identifies keywords like `if`, `var`, or operators like `==`, `!=`.
    + **peekChar Method**: Peeks at the next character without advancing the read position.
    + **skipWhiteSpace Method:** Skips over whitespce (spaces, tabs, newlines).

    ### Parser
    + **Purpose:** Converts the stream of tokens produced by the lexer into an Abstract Syntax TRee(AST), which can then be interpreted or compiled.
    + **Key Components:** Advances the current token to the next one in the input.
    + **parseStatement Method:** Placeholder for parsing statements. This is where various langauage constructs (like variable declaration or expressions) will be parsed and converted into AST nodes.
    

    
# Toy-programming-language 

## Introduction
Ksm is an interpreted Toy Programming language written in Golang.

## Features
1. It uses a syntax similar to that of the Go Programming Language.
2. It supports the following data structures, along with built-in methods and functions:

    ```
    ______________________________________________
    | Data structure |  Methods     |  Functions |
    ______________________________________________
    | string         | len()        | concat()   |
    |                | substring()  | indexOf()  |
    |                | toUpper()    | toLower()  |
    |----------------|--------------|------------|
    | boolean        |              |            |
    |----------------|--------------|------------|
    | null           |              |            |
    |----------------|--------------|------------|
    | array          | len()        | append()   |
    |                | first()      | last()     |
    |                | slice()      | contains() |
    |----------------|--------------|------------|
    | struct         |              |            |
    |----------------|--------------|------------|
    | hash           | keys()       | values()   |
    |                | get()        | set()      |
    ______________________________________________
    ```

## File Structure pf the Complete langugae interpreter

```markdown
Kisumu/
│
├── cmd/                      # Main command line entry point
│   └── kisumu.go              # Main Kisumu interpreter
│
├── lexer/                    # Lexer module for tokenizing input
│   ├── lexer.go               # Lexical analyzer logic
│   └──  lexer_test.go          # Tests for lexer 
│
├── token/                    # Token module for token definition and helper functions
│   └── token.go               # Token definitions
│
├── parser/                   # Parser module for constructing AST
│   ├── parser.go              # Recursive descent parser logic
│   └──  parser_test.go         # Tests for parser
│
├── ast/                      # Abstract Syntax Tree module for node definitions
│   ├── ast.go                 # Abstract Syntax Tree (AST) node definitions
│   └── grammar.go             # Grammar rules and parser utilities
│
├── interpreter/              # Interpreter module for evaluating AST
│   ├── interpreter.go         # Main interpreter logic
│   ├── interpreter_test.go    # Tests for interpreter
│   ├── environment.go         # Environment for storing variable bindings
│   └── eval.go                # Evaluation of expressions/statements
│
├── repl/                     # REPL (Read-Eval-Print Loop)
│   └── repl.go                # Interactive Kisumu shell for testing
│
├── types/                    # Kisumu-specific data types
│   ├── types.go               # Type definitions (number, string, boolean, etc.)
│   └── methods.go             # Built-in methods for each type
│
├── documentation/           # Documentation files
│   ├── introduction.md
│   ├── tokens.md
│   ├── lexer.md
│   ├── parser.md
│   ├── ast.md
│   ├── interpreter.md
│   ├── repl.md
│   └── types.md
│
├── go.mod                    # Go module definition
└── README.md                 # Project overview and instructions

```
## Getting Started 
To get started with the Kisumu Programming language:
1. CLone the repository:
```
git clone https://github.com/LuvDokta/Toy-programming-language.git
cd kisumu
```
2. Install dependancies using Go:
go mod tidy

## Writing Ksm Programs
Ksm programs use a syntax similar to Go. You can write programs using .ksm files. Here's an example of a simple program:
```
// Sample progrram in Kisumu 
var "Kisumu";
var number = 42;

if (number > 10) {
    print("The number is greater than 10");
}

var greeting = "Hello, " + name + "!";
print(greeting);
```

### Documnentation
Use the list of links below to navigate the different chapters of the ksm toy programming language Documentation
1. [Project Journal](/documentation/introduction.md)

### completed featers
2. [Token Module for token defination and helper functions.](/documentation/tokens.md)
3. [Lexer Module for constructing AST](/documentation/lexer.md)
4. [Abstract Syntax Tree module for node definations](/documentation/ast.md)

### Ongoing implementions
5. [Interpreter module for evaluating AST](/documentation/interpreter.md)
6. [REPL (Read-Eval-Print Loop)](/documentation/repl.md)
7. [Ksm-specific data types](/documentation/datatypes.md) 


## Contributing
Contributions are welcome! Feel free to submit issues  or pull requests to improve the language.

## Authors
1. [Stephen Odhiambo](https://github.com/steodhiambo)
2. [Clifford Otieno](https://github.com/LuvDokta)
3. [Ouma Godwin](https://github.com/garveyshah)
4. [Anne Okingo](https://github.com/Anne-Okingo)

## LIcense
This Project is licensed under the MIT License-see the LICENSE file for details.

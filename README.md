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
├── documentation/             # Main command line entry point
│   ├── Introduction.md         # Environment for storing variable bindings
│   ├── tokens.md         # Environment for storing variable bindings
│   ├── lexer.md         # Environment for storing variable bindings
│   ├── parser.md         # Environment for storing variable bindings
│   ├── ast.md         # Environment for storing variable bindings
│   ├── interpratot.md         # Environment for storing variable bindings
│   ├── repl.md         # Environment for storing variable bindings
│   └── types.md              # Main Kisumu interpreter
│
├── go.mod                    # Go module definition
└── README.md                 # Project overview and instructions

```
## Getting Started 
To get started with the Kisumu Programming language:
1. CLone the repository:
```
git clone https://github.com/yourusername/kisumu.git
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
2. [Token Module for token defination and helper functions.]()
3. [Lexer Module for constructing AST]()
4. [Abstract Syntax Tree module for node definations]()

### Ongoing implementions
5. [Interpreter module for evaluating AST]()
6. [REPL (Read-Eval-Print Loop)]()
7. [Ksm-specific data types]() 


## Contributing
Contributions are welcome! Feel free to submit issues  or pull requests to improve the language.

## Authors
1. [Stephen Odhiambo]()
2. [Clifford Otieno]()
3. [Ouma Godwin]()
4. [Anne Okingo]()

## LIcense
This Project is licensed under the MIT License-see the LICENSE file for details.

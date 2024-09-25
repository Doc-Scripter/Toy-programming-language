package evaluator

import (
	"fmt"

	"ksm/ast"
	"ksm/environment"
)

// Eval evaluates the AST node
func Eval(node ast.Node, env *environment.SymbolTable) interface{} {
	switch node := node.(type) {

	// Handling the root of the program (list of statements)
	case *ast.Program:
		return evalProgram(node, env)

	// Handling variable assignments
	case *ast.AssignStatement:
		value := Eval(node.Value, env)
		env.Set(node.Name.Value, value)
		return value

	// Handling variable declarations (let)
	case *ast.VarStatement:
		value := Eval(node.Value, env)
		env.Set(node.Name.Value, value)
		return value

	// Handling variable retrieval
	case *ast.Identifier:
		if value, exists := env.Get(node.Value); exists {
			return value
		}
		return fmt.Sprintf("Unknown identifier: %s", node.Value)

	// Handling integer literals
	case *ast.IntegerLiteral:
		return node.Value

	// Handling string literals
	case *ast.StringLiteral:
		return node.Value

	// Handling boolean literals
	case *ast.Boolean:
		return node.Value

	// Add more cases for infix expressions, block statements, if expressions, etc.
	default:
		return fmt.Sprintf("Unknown AST node type: %T", node)
	}
}

// evalProgram evaluates the statements inside the program
func evalProgram(program *ast.Program, env *environment.SymbolTable) interface{} {
	var result interface{}

	for _, statement := range program.Statements {
		result = Eval(statement, env) // Evaluate each statement

		// You can add additional logic to handle different types of statements here
	}

	return result
}

// evalPrefixExpression evaluates a prefix expression (e.g., -5, !true)
func evalPrefixExpression(operator string, right interface{}) interface{} {
	switch operator {
	case "!":
		return !isTruthy(right)
	case "-":
		if v, ok := right.(int); ok {
			return -v
		}
	}
	return fmt.Sprintf("Unknown prefix operator: %s", operator)
}

// isTruthy returns the boolean truthiness of a value
func isTruthy(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case nil:
		return false
	default:
		return true
	}
}

package evaluator

import (
	"fmt"

	"ksm/ast"
	"ksm/environment"
)

func Eval(node ast.Node, env *environment.SymbolTable) interface{} {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.VarStatement:
		return evalLetStatement(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.IntegerLiteral:
		return evalIntegerLiteral(node, env)
	case *ast.BinaryExpression:
		return evalBinaryExpression(node, env)
	// Add more cases for other node types
	default:
		fmt.Println("Unknown node type")
		return nil
	}
}

func evalProgram(program *ast.Program, env *environment.SymbolTable) interface{} {
	var result interface{}
	for _, statement := range program.Statements {
		result = Eval(statement, env)
	}
	return result
}

func evalLetStatement(node *ast.VarStatement, env *environment.SymbolTable) interface{} {
	value := Eval(node.Value, env)
	env.Set(node.Name.Value, value)
	return value
}

func evalIdentifier(node *ast.Identifier, env *environment.SymbolTable) interface{} {
	value, found := env.Get(node.Value)
	if !found {
		fmt.Printf("Variable %s not found\n", node.Value)
		return nil
	}
	return value
}

func evalIntegerLiteral(node *ast.IntegerLiteral, env *environment.SymbolTable) interface{} {
	return node.Value
}

func evalBinaryExpression(node *ast.BinaryExpression, env *environment.SymbolTable) interface{} {
	left := Eval(node.Left, env)
	right := Eval(node.Right, env)

	leftVal, okLeft := left.(int64)
	rightVal, okRight := right.(int64)

	switch node.Operator {
	case "+":
		if okLeft && okRight {
			return leftVal + rightVal
		}
	case "-":
		if okLeft && okRight {
			return leftVal - rightVal
		}
	case "*":
		if okLeft && okRight {
			return leftVal * rightVal
		}
	case "/":
		if okLeft && okRight {
			if rightVal == 0 {
				fmt.Println("Division by zero")
				return nil
			}
			return leftVal / rightVal
		}
	}
	fmt.Println("Unsupported operation or type mismatch")
	return nil
}

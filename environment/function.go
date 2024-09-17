package environment

import (
	"fmt"
	"go/ast"
)

type Environment struct {
	SymbolTable *SymbolTable // current symbol table context
}

// Create a new execution environment
func NewEnvironment(parent *SymbolTable) *Environment {
	return &Environment{
		SymbolTable: NewSymbolTable(parent),
	}
}

// Create a new execution environment
type Function struct {
	Name       string
	Parameters []string
	Body       ast.BlockStmt
}

// Function execution method
func (env *Environment) CallFunction(fn *Function, args []interface{}) (interface{}, error) {
	newEnv := NewEnvironment(env.SymbolTable) // Create an instance of the environement for the function call

	// Bind the parameters to the new environment
	for i, param := range fn.Parameters {
		newEnv.SymbolTable.Set(param, args[i])
	}

	// Execute the function body [if it exists]
	result := executeBlock(fn.Body, newEnv)
	return result, nil
}

func (env *Environment) Assign(name string, value interface{}) {
	env.SymbolTable.Set(name, value)
}

func (env *Environment) Get(name string) (interface{}, bool) {
	return env.SymbolTable.Get(name)
}

func main() {
	// Creating a global symbol table
	Scope := NewSymbolTable(nil)
	env := NewEnvironment(Scope)

	// Setting a variable
	env.Assign("x", 10)

	// Calling a function(assume we have a function defined)
	fn := &Function{Name: "Concat", Parameters: []string{"a", "b"}}
	result,_:=env.CallFunction(fn,[]interface{}{"Hello"},{"World"})

	fmt.Println("Result of a function call",result)

	//Getting the variable
	value ,found :=env.Get("x")
	if found{
		fmt.Println("Value of x :",value)//should output latest value of x
	}
}

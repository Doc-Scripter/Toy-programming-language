package environment

// Symbol table entry to hold variable names and values
type SymbolTable struct {
	vars   map[string]interface{} // interface to allow for any type
	parent *SymbolTable           // Pointer to the current scope(for nested scopes)
}

// Create the instance of the symbol Table
func NewSymbolTable(parent *SymbolTable) *SymbolTable {
	return &SymbolTable{
		vars:   make(map[string]interface{}),
		parent: parent,
	}
}

// set (add or update a variable)
func (st *SymbolTable) Set(name string, value interface{}) {
	st.vars[name] = value
}

// Get(retrieve a variable)
func (st *SymbolTable) Get(name string) (interface{}, bool) {
	if value, exists := st.vars[name]; exists {
		return value, true
	}
	if st.parent != nil {
		return st.parent.Get(name) // Look in the parent if not found
	}
	return nil, false // Not found
}

// Remove(delete a variabe)
func (st *SymbolTable) Remove(name string) {
	delete(st.vars, name)
}

// func main() {
// 	global := NewSymbolTable(nil) // intial scope
// 	// global.Set("x", 10)
// child:=NewSymbolTable(global)
// 	// value, found := global.Get("x")
// 	// if found {
// 	// 	fmt.Println("Value of x:", value) // output value of x
// 	// } else {
// 	// 	fmt.Println("Variable not found")
// 	// }
// 	// global.Set("x", 20)

// 	// value, found = global.Get("x")
// 	// if found {
// 	// 	fmt.Println("Value of x:", value) // output value of x
// 	// } else {
// 	// 	fmt.Println("Variable not found")
// 	// }
// 	// global.Remove("x")
// 	// value, found = global.Get("x")
// 	// if found {
// 	// 	fmt.Println("Value of x:", value) // output value of x
// 	// } else {
// 	// 	fmt.Println("Variable not found")
// 	// }

// 	// the next thing is function call.go.
// 	// objects.go is used to define our data structure and how they behave
// 	// for example to define that "an interger can not be added to an string directly"
// 	// x=5
// 	// x+5
// 	// x=10
// 	// Call the function
// 	exampleFunc(5, 15, child)

// 	resultValue, foundResult := global.Get("result")
// 	if foundResult {
// 		fmt.Println("Result:", resultValue) // Output: Result: 20
// 	}
// }

// func exampleFunc(a int, b int, scope *SymbolTable) int {
// 	sum := a + b
// 	scope.Set("result", sum) // Store the result in child scope
// 	return sum
// }

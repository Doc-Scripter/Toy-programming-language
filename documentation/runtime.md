# Overview
The runtime file defines a symbol table which facilitates variable storage and retrieval allowing for nested scopes through parent references

# SymbolTable Struct
holds a mapping of variable names to their values and maintains a reference to a parent SymbolTable for nested scopes

``` go
type SymbolTable struct {
    vars   map[string]interface{} 
    parent *SymbolTable 
}

```

# Methods

## 1.NewSymbolTable
Creates an instance of SymbolTable and returns a pointer to the newly created SymbolTable
```go
func NewSymbolTable(parent *SymbolTable) *SymbolTable

```
* Parameters
  * parent: A pointer to the parent SymbolTable, which can be nil for the global scope.
  * Returns: A pointer to the newly created SymbolTable.

## 2.Set
Adds or updates a variable in the symbol table.
name
``` go
func (st *SymbolTable) Set(name string, value interface{})

```
* Parameters:
  * name: The name of the variable to set.
  * value: The value to associate with the variable. It can be of any type due to the use of interface{}.


## 3.Get
Retrieves the value of a variable from the symbol table


```go
func (st *SymbolTable) Get(name string) (interface{}, bool)

```
* Parameters:

  * name: The name of the variable to find.
  * Returns: The value of the variable and a boolean to indicate if found.
## 4.Remove
Deletes a variable from the symbol table.
```go
func (st *SymbolTable) Remove(name string)

```
* Parameters:
   * name: The name of the variable to remove.


Example Usage

```go
func main() {
    // Initialize the global symbol table (initial scope)
    global := NewSymbolTable(nil)

    // Set variable x to 10
    global.Set("x", 10)

    // Retrieve and print the value of x
    if value, found := global.Get("x"); found {
        fmt.Println("Value of x:", value) // Output: Value of x: 10
    } else {
        fmt.Println("Variable not found")
    }

    // Update variable x to 20
    global.Set("x", 20)

    // Retrieve and print the updated value of x
    if value, found := global.Get("x"); found {
        fmt.Println("Value of x:", value) // Output: Value of x: 20
    } else {
        fmt.Println("Variable not found")
    }

    // Remove variable x
    global.Remove("x")

    // Attempt to retrieve x after removal
    if value, found := global.Get("x"); found {
        fmt.Println("Value of x:", value)
    } else {
        fmt.Println("Variable not found") // Output: Variable not found
    }
}

```

Conclusion
Using the Symboltable struct allows for efficient management and supports nested scopes through parent-child relationships

package environment

import (
	"testing"
)

func TestNewSymbolTable(t *testing.T) {
	// Test case 1: Creating a new symbol table with no parent
	symbolTable := NewSymbolTable(nil)
	if symbolTable == nil {
		t.Errorf("Expected non-nil symbol table, got nil")
	}
	if symbolTable.parent != nil {
		t.Errorf("Expected nil parent, got %v", symbolTable.parent)
	}
	if len(symbolTable.vars) != 0 {
		t.Errorf("Expected empty vars map, got %v", symbolTable.vars)
	}

	// Test case 2: Creating a new symbol table with a parent
	parentTable := &SymbolTable{
		vars:   make(map[string]interface{}),
		parent: nil,
	}
	symbolTableWithParent := NewSymbolTable(parentTable)
	if symbolTableWithParent == nil {
		t.Errorf("Expected non-nil symbol table, got nil")
	}
	if symbolTableWithParent.parent != parentTable {
		t.Errorf("Expected parent %v, got %v", parentTable, symbolTableWithParent.parent)
	}
	if len(symbolTableWithParent.vars) != 0 {
		t.Errorf("Expected empty vars map, got %v", symbolTableWithParent.vars)
	}
}

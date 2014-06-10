package mutan

import (
	"container/list"
	"fmt"
)

type Scope interface {
	NewVar(id string, typ varType) (Var, error)
	SetVar(Var)
	GetVar(string) Var
	Size() int

	MakeReturn(expr *SyntaxTree, gen *IntGen) *IntInstr
}

type IntGen struct {
	VarTable map[string]Var

	//locals        map[string]*Variable
	functionTable map[string]*Function
	arrayTable    map[string][]*IntInstr
	stringTable   map[string][]*IntInstr

	errors []error

	scopes *list.List

	currentStackSize int
}

func NewGen() *IntGen {
	return &IntGen{
		VarTable:      make(map[string]Var),
		functionTable: make(map[string]*Function),
		arrayTable:    make(map[string][]*IntInstr),
		stringTable:   make(map[string][]*IntInstr),
		scopes:        list.New(),
	}
}

func (self *IntGen) NewVar(id string, typ varType) (Var, error) {
	if self.VarTable[id] != nil {
		return nil, fmt.Errorf("redeclaration of '%v'", id)
	}

	var v *Variable
	switch typ {
	case varNumTy:
		v = NewNumeric(id, self.Size())
	default:
		v = NewNumeric(id, self.Size())
	}

	v.offset = self.Size()
	v.typ = typ
	self.VarTable[id] = v

	return v, nil
}

func (self *IntGen) SetVar(v Var) {
	self.VarTable[v.Id()] = v
}

func (self *IntGen) GetVar(id string) Var {
	if self.CurrentScope() != self {
		variable := self.CurrentScope().GetVar(id)
		if variable != nil {
			return variable
		}
	}

	return self.VarTable[id]
}

func (self *IntGen) Size() (size int) {
	for _, variable := range self.VarTable {
		size += variable.Size()
	}

	return
}

func (self *IntGen) CurrentScope() Scope {
	scope := self.scopes.Back()
	if scope != nil {
		return scope.Value.(Scope)
	}

	return self
}

func (self *IntGen) MakeReturn(expr *SyntaxTree, gen *IntGen) *IntInstr {
	c, err := expr.Errorf("return now allowed in global scope")
	gen.addError(err)

	return c
}

func (self *IntGen) PopScope() Scope {
	scope := self.scopes.Back()
	if scope != nil {
		self.scopes.Remove(scope)

		return scope.Value.(Scope)
	}

	return nil
}

func (gen *IntGen) PushScope(scope Scope) {
	gen.scopes.PushBack(scope)
}

func (gen *IntGen) Errors() []error {
	return gen.errors
}

func (gen *IntGen) addError(e error) {
	gen.errors = append(gen.errors, e)
}

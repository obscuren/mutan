package frontend

import "fmt"

type LocalScope struct {
	VarTable Variables
}

func NewLocalScope() *LocalScope {
	return &LocalScope{make(Variables)}
}

func (self *LocalScope) SetVar(v Var) {
	self.VarTable[v.Id()] = v
}

func (self *LocalScope) GetVar(id string) Var {
	return self.VarTable[id]
}

func (self *LocalScope) Size() (size int) {
	for _, variable := range self.VarTable {
		size += variable.Size()
	}

	return
}

func (self *LocalScope) NewVar(id string, typ varType) (Var, error) {
	if self.VarTable[id] != nil {
		return nil, fmt.Errorf("redeclaration of '%v'", id)
	}

	var v *Variable
	switch typ {
	case varNumTy:
		v = NewNumeric(id, self.Size())
		// TODO other types
	default:
		v = NewNumeric(id, self.Size())
	}

	v.offset = self.Size()
	v.typ = typ
	self.VarTable[id] = v

	return v, nil
}

func (self *LocalScope) MakeReturn(expr *SyntaxTree, gen *IntGen) *IntInstr { return nil }

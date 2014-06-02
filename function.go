package mutan

import (
	"fmt"
	"strconv"
)

type Function struct {
	CallTarget    *IntInstr
	Id            string
	ArgCount      int
	Ret           bool
	VarTable      map[string]*Variable
	frame, offset int
}

func NewFunction(id string, target *IntInstr, ac int, ret bool) *Function {
	return &Function{target, id, ac, ret, make(map[string]*Variable), 0, 0}
}

func (self *Function) StackSize() (size int) {
	for _, variable := range self.VarTable {
		size += variable.size
	}

	// Add 32 for the return location
	// Add 32 for the stack size
	size += 64

	return
}

func (self *Function) GenOffset() int {
	return self.StackSize()
}

func (self *Function) NewVariable(id string, typ varType) (*Variable, error) {
	if self.VarTable[id] != nil {
		return nil, fmt.Errorf("redeclaration of '%v'", id)
	}

	variable := &Variable{id: id, pos: self.GenOffset()}
	self.VarTable[id] = variable

	return variable, nil
}

func (self *Function) SetVariable(v *Variable) {
	self.VarTable[v.id] = v
}

func (self *Function) GetVariable(id string) *Variable {
	return self.VarTable[id]
}

func (self *Function) Call(gen *IntGen, scope Scope) *IntInstr {
	self.frame = self.StackSize() + scope.StackSize()

	setPtr := gen.addStackPtr(scope.StackSize())
	size := gen.makePush(strconv.Itoa(self.StackSize()))
	ptr1 := gen.loadStackPtr()
	offset := gen.makePush("32")
	add1 := newIntInstr(intAdd, "")
	sizeStore := newIntInstr(intMStore, "")

	pc := newIntInstr(intPc, "")
	push := gen.makePush("14")
	add := newIntInstr(intAdd, "")
	ret := gen.loadStackPtr()
	retStore := newIntInstr(intMStore, "")

	p, jmp := newJumpInstr(intJump)
	jmp.Target = self.CallTarget

	concat(setPtr, size)
	concat(size, ptr1)
	concat(ptr1, offset)
	concat(offset, add1)
	concat(add1, sizeStore)
	concat(sizeStore, pc)
	concat(pc, push)
	concat(push, add)
	concat(add, ret)
	concat(ret, retStore)
	concat(retStore, p)

	return setPtr
}

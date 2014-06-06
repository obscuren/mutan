package mutan

import (
	"fmt"
	_ "strconv"
)

type Variables map[string]Var

func (self Variables) String() (ret string) {
	for k, v := range self {
		ret += fmt.Sprintf("  %-14s : %d\n", k, v.Offset())
	}
	return
}

type ArgList []Var

type Function struct {
	CallTarget *IntInstr
	Id         string
	ArgCount   int
	Ret        bool
	//VarTable      map[string]*Variable
	VarTable      Variables
	ArgTable      ArgList
	frame, offset int
}

func NewFunction(id string, target *IntInstr, ac int, ret bool) *Function {
	return &Function{target, id, ac, ret, make(map[string]Var), nil, 0, 0}
}

func (self *Function) String() string {
	return fmt.Sprintf(`
### %s ###
Arguments: %d
Var Table:
%v
`, self.Id, self.ArgCount, self.VarTable)

}

func (self *Function) GenOffset() int {
	return self.Size()
}

func (self *Function) NewVar(id string, typ varType) (Var, error) {
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
	self.VarTable[id] = v

	return v, nil
}

func (self *Function) PushArg(v Var) {
	self.ArgTable = append(self.ArgTable, v)
}

func (self *Function) SetVar(v Var) {
	self.VarTable[v.Id()] = v
}

func (self *Function) GetVar(id string) Var {
	return self.VarTable[id]
}

func (self *Function) Size() (size int) {
	for _, variable := range self.VarTable {
		size += variable.Size()
	}

	return
}

func (self *Function) Call(args []*SyntaxTree, gen *IntGen, scope Scope) *IntInstr {
	if len(args) != self.ArgCount {
		gen.addError(fmt.Errorf("%s takes %d arguments. Received %d arguments", self.Id, self.ArgCount, len(args)))
		return newIntInstr(intIgnore, "")
	}

	self.frame = self.Size() + scope.Size()

	stackPtr := gen.loadStackPtr()
	setPtr := gen.addStackPtr(scope.Size())
	nStackPtr := gen.loadStackPtr()
	offset := gen.makePush("32")
	add1 := newIntInstr(intAdd, "")
	sizeStore := newIntInstr(intMStore, "")

	argInstr := newIntInstr(intIgnore, "")
	for i, arg := range args {
		arg := gen.MakeIntCode(arg)
		assign := gen.assignMemory(self.ArgTable[i].Offset())

		concat(argInstr, arg)
		concat(arg, assign)
	}

	pc := newIntInstr(intPc, "")
	push := gen.makePush("14")
	add := newIntInstr(intAdd, "")
	ret := gen.loadStackPtr()
	retStore := newIntInstr(intMStore, "")

	p, jmp := newJumpInstr(intJump)
	jmp.Target = self.CallTarget

	concat(stackPtr, setPtr)
	concat(setPtr, nStackPtr)
	concat(nStackPtr, offset)
	concat(offset, add1)
	concat(add1, sizeStore)
	concat(sizeStore, argInstr)
	concat(argInstr, pc)
	concat(pc, push)
	concat(push, add)
	concat(add, ret)
	concat(ret, retStore)
	concat(retStore, p)

	return stackPtr
}

func (self *Function) MakeReturn(expr *SyntaxTree, gen *IntGen) *IntInstr {
	retVal := gen.MakeIntCode(expr)

	rPos := gen.loadStackPtr()
	dup := newIntInstr(intDup, "")
	// Increment by 1 word
	offset := gen.makePush("32")
	add := newIntInstr(intAdd, "")
	sizeLoad := newIntInstr(intMLoad, "")
	// Now pop the frame off the stack
	//sub := newIntInstr(intSub, "")

	stackPtrOffset := gen.makePush("0")
	stackPtrStore := newIntInstr(intMStore, "")

	rPosLoad := newIntInstr(intMLoad, "")
	jumpBack := newIntInstr(intJump, "")

	concat(retVal, rPos)
	concat(rPos, dup)
	concat(dup, offset)
	concat(offset, add)
	concat(add, sizeLoad)
	/*
		concat(sizeLoad, sub)
		concat(sub, stackPtrOffset)
	*/
	concat(sizeLoad, stackPtrOffset)
	concat(stackPtrOffset, stackPtrStore)
	concat(stackPtrStore, rPosLoad)
	concat(rPosLoad, jumpBack)

	return retVal
}

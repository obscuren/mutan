package frontend

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
	v.typ = typ
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
		return newIntInstr(IntIgnore, "")
	}

	argsPush := newIntInstr(IntIgnore, "")
	for _, arg := range args {
		push := gen.MakeIntCode(arg)

		concat(argsPush, push)
	}

	stack := gen.PushStack()

	argsPop := newIntInstr(IntIgnore, "")
	for i, _ := range args {
		//arg := gen.MakeIntCode(arg)
		pop := gen.assignMemory(self.ArgTable[i].Offset())

		concat(argsPop, pop)
	}

	pc := newIntInstr(IntPc, "")           // 1
	push := gen.makePush("17")             // 2
	add := newIntInstr(IntAdd, "")         // 1
	ret := gen.loadStackPtr()              // 3
	offset := gen.makePush("32")           // 2
	add2 := newIntInstr(IntAdd, "")        // 1
	retStore := newIntInstr(IntMStore, "") // 1
	p, jmp := newJumpInstr(IntJump)        // 6
	jmp.Target = self.CallTarget

	cc(argsPush, stack, argsPop, pc, push, add, ret, offset, add2, retStore, p)

	return argsPush
}

func (self *Function) MakeReturn(expr *SyntaxTree, gen *IntGen) *IntInstr {
	retVal := gen.MakeIntCode(expr)

	rPos := gen.loadStackPtr()
	stack := gen.PopStack()
	/*
		dup := newIntInstr(IntDup1, "")
		sizeLoad := newIntInstr(IntMLoad, "")
		stackPtrOffset := gen.makePush("0")
		stackPtrStore := newIntInstr(IntMStore, "")
	*/

	offset := gen.makePush("32")
	add := newIntInstr(IntAdd, "")

	rPosLoad := newIntInstr(IntMLoad, "")
	jumpBack := newIntInstr(IntJump, "")

	cc(retVal, rPos, stack, offset, add, rPosLoad, jumpBack)

	return retVal
}

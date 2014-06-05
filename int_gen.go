package mutan

import (
	"container/list"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

type Scope interface {
	NewVar(id string, typ varType) (Var, error)
	SetVar(Var)
	GetVar(string) Var
	Size() int

	//NewVariable(id string, typ varType) (*Variable, error)
	//SetVariable(*Variable)
	//GetVariable(string) *Variable
	//StackSize() int
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
		VarTable: make(map[string]Var),
		//locals:        make(map[string]*Variable),
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

/*
func (self *IntGen) SetVariable(v *Variable) {
	self.locals[v.id] = v
}

func (self *IntGen) GetVariable(id string) *Variable {
	if self.CurrentScope() != self {
		variable := self.CurrentScope().GetVariable(id)
		if variable != nil {
			return variable
		}
	}

	return self.locals[id]
}

func (self *IntGen) NewVariable(id string, typ varType) (*Variable, error) {
	if self.locals[id] != nil {
		return nil, fmt.Errorf("redeclaration of '%v'", id)
	}

	variable := &Variable{id: id, pos: self.StackSize(), size: 32}
	self.locals[id] = variable

	return variable, nil
}

func (self *IntGen) StackSize() (size int) {
	for _, variable := range self.locals {
		size += variable.size
	}

	// Stack ptr
	//size += 32

	return
}
*/

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

func (gen *IntGen) findOffset(tree *SyntaxTree, offset int) (position string, err error) {
	var pos string
	variable := gen.GetVar(tree.Constant)

	if variable == nil {
		return "", fmt.Errorf("Undefined variable: %v", tree.Constant)
	} else {
		var p int
		switch variable.Type() {
		case varNumTy:
			p = variable.Offset()
		}

		/*
			if variable.typ == varArrTy {
				p = variable.pos + (offset * variable.size)
			} else {
				p = variable.pos
			}
		*/
		pos = strconv.Itoa(p)
	}

	return pos, nil
}

// Generates asm for getting a memory address
func (gen *IntGen) getMemory(tree *SyntaxTree, offset int) (push *IntInstr, err error) {
	pos, err := gen.findOffset(tree, offset)
	if err != nil {
		push = newIntInstr(intIgnore, "")
		return
	}

	push, cons := pushConstant(pos)
	load := newIntInstr(intMLoad, "")
	concat(push, cons)
	concat(cons, load)

	return
}

func makeStore(offset int) *IntInstr {
	push, cons := pushConstant(strconv.Itoa(offset))
	store := newIntInstr(intMStore, "")

	concat(push, cons)
	concat(cons, store)

	return push
}

func (gen *IntGen) assignMemory(offset int) *IntInstr {
	ptr := gen.loadStackPtr()
	push, cons := pushConstant(strconv.Itoa(offset))
	add := newIntInstr(intAdd, "")
	store := newIntInstr(intMStore, "")

	concat(ptr, push)
	concat(push, cons)
	concat(cons, add)
	concat(add, store)

	return ptr
}

// Generates asm for setting a memory address
func (gen *IntGen) setMemory(tree *SyntaxTree) (*IntInstr, error) {
	variable := gen.GetVar(tree.Constant)
	if variable == nil {
		return tree.Errorf("Undefined variable '%s'", tree.Constant)
	}

	instr := gen.assignMemory(variable.Offset())

	return instr, nil
}

// Type size table
var typeToSize = map[string]int{
	"bool":   8,
	"int":    32,
	"int8":   8,
	"int16":  16,
	"int32":  32,
	"int64":  64,
	"int256": 256,
	"byte":   8,
	"big":    256,
	"string": 8,
	"addr":   160,
}

func sizeOf(typ string) int {
	size := typeToSize[typ]

	size /= 8

	return 256 / 8
}

func (gen *IntGen) initNewNumber(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant

	scope := gen.CurrentScope()
	_, err := scope.NewVar(name, varNumTy)
	if err != nil {
		return newIntInstr(intIgnore, ""), err
	}

	return nil, nil
}

func (gen *IntGen) sizeof(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)

	if variable == nil {
		return tree.Errorf("undefined variable: '%s'", name)
	}

	push, constant := pushConstant(strconv.Itoa(variable.Size()))

	return concat(push, constant), nil
}

func (gen *IntGen) getArray(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)

	if variable == nil {
		tree.Errorf("undefined array: %v", name)
	}

	// TODO optimize if the expression in offset. If regular const (i.e. 0-9)
	// do an inline calculation instead.

	// Get the location of the variable in memory
	loc, locConst := pushConstant(strconv.Itoa(variable.Offset()))
	// Get the offset (= expression between brackets [expression])
	offset := gen.MakeIntCode(tree.Children[0])
	// Get the size of the variable in bytes
	size, sizeConst := pushConstant(strconv.Itoa(variable.Size()))
	// Multiply offset with size
	mul := newIntInstr(intMul, "")
	// Add the result to the memory location
	add := newIntInstr(intAdd, "")
	// b = a[0] // loc(a) + sizeOf(type(a)) * len(a)
	load := newIntInstr(intMLoad, "")

	concat(loc, locConst)
	concat(locConst, offset)
	concat(offset, size)
	concat(size, sizeConst)
	concat(sizeConst, mul)
	concat(mul, add)
	concat(add, load)

	return loc, nil
}

func (gen *IntGen) setArray(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)
	if variable == nil {
		return tree.Errorf("undefined array: %v", name)
	}

	// The value which we want to assign
	val := gen.MakeIntCode(tree.Children[1])

	// Get the location of the variable in memory
	loc, locConst := pushConstant(strconv.Itoa(variable.Offset()))
	gen.arrayTable[name] = append(gen.arrayTable[name], locConst)

	// Get the offset (= expression between brackets [expression])
	offset := gen.MakeIntCode(tree.Children[0])
	// Get the size of the variable in bytes
	//size, sizeConst := pushConstant(strconv.Itoa(local.varSize))
	size, sizeConst := pushConstant("32")
	// Multiply offset with size
	mul := newIntInstr(intMul, "")
	// Add the result to the memory location
	add := newIntInstr(intAdd, "")
	store := newIntInstr(intMStore, "")

	concat(val, loc)
	concat(loc, locConst)
	concat(locConst, offset)
	concat(offset, size)
	concat(size, sizeConst)
	concat(sizeConst, mul)
	concat(mul, add)
	concat(add, store)

	return val, nil
}

func (gen *IntGen) initNewArray(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)
	variable, err := gen.CurrentScope().NewVar(name, varArrTy)
	if err != nil {
		return tree.Errorf("Redeclaration of variable '%s'", name)
	}

	/*
		var size int
		switch tree.VarType {
		case "var", "bool", "int8", "int16", "int32", "int64", "int256", "big":
			size = sizeOf(tree.VarType)
		}
		size *= length
	*/

	length, _ := strconv.Atoi(tree.Size)
	variable.SetSize(32 * length)

	return newIntInstr(intIgnore, ""), nil
}

func validLhSide(variable Var, typ varType) {
	if variable != nil && variable.Type() != varUndefinedTy && variable.Type() != typ {
		panic(fmt.Sprintf("cannat assign %v to '%v' of type %v", typ, variable.Id(), variable.Type()))
	}
}

func (gen *IntGen) setVariable(tree *SyntaxTree, identifier *SyntaxTree) *IntInstr {
	var instr *IntInstr

	/*
		name := identifier.Constant
		id := gen.locals[identifier.Constant]
	*/

	variable := gen.GetVar(identifier.Constant)

	//TODO Do left hand side type checking at this point
	switch tree.Type {
	case StringTy:
		validLhSide(variable, varStrTy)

		var t Instr
		if identifier.Type == SetStoreTy {
			if len(tree.Constant) > 32 {
				panic("attempting to store string greater than 32")
			}

			return gen.makePush("0x" + hex.EncodeToString([]byte(tree.Constant)))
		} else {
			t = intMStore
			variable.SetType(varStrTy)
		}

		var length int
		instr, length = gen.stringToInstr(variable, []byte(tree.Constant), t)

		if identifier.Type != SetStoreTy {
			variable.SetSize(int(math.Max(math.Max(float64(variable.Size()), float64(length)), 32.0)))
		}

		return instr
	case ConstantTy:
		validLhSide(variable, varNumTy)

		if identifier.Type != SetStoreTy {
			variable.SetType(varNumTy)
		}

		instr = gen.MakeIntCode(tree)
	default:
		instr = gen.MakeIntCode(tree)
	}

	return instr
}

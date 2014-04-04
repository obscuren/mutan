package mutan

/*
int code generator takes care of memory allocation
generate all opcode and instructions
sets pc for each instruction afterwards

compiler transforms int code to ASM (very static)
*/

import (
	"fmt"
	"strconv"
	"strings"
)

type varType byte

const (
	varNumTy varType = iota
	varStrTy
	varArrTy
)

type Variable struct {
	typ  varType
	val  string
	pos  int
	size int
}

type Instr byte

const (
	intEqual Instr = iota
	intGt
	intLt
	intMul
	intAdd
	intDiv
	intSub
	intAssign
	intConst
	intEmpty
	intJump
	intTarget
	intPush
	intMStore
	intMLoad
	intNot
	intJumpi
	intSStore
	intSLoad
	intStop
	intOrigin
	intCaller
	intCallVal
	intCallDataLoad
	intCallDataSize
	intGasPrice

	// Asm is a special opcode. It's not malformed in anyway
	intASM
	intArray
	intCall
	intSizeof

	intIgnore
)

var instrAsString = []string{
	"equal",
	"gt",
	"lt",
	"mul",
	"add",
	"div",
	"sub",
	"assign",
	"const",
	"empty",
	"jump",
	"target",
	"push",
	"mstore",
	"mload",
	"not",
	"jmpi",
	"sstore",
	"sload",
	"stop",
	"origin",
	"caller",
	"value",
	"dataload",
	"datasize",
	"gasprice",

	"asm",
	"array",
	"call",
	"sizeof",

	"ignore",
}

func (op Instr) String() string {
	if len(instrAsString) < int(op) {
		return "Unknown"
	}

	return instrAsString[op]
}

type CodeGen struct {
	locals map[string]*Variable

	memPos   int
	lastPush *IntInstr

	errors []error
}

func NewGen() *CodeGen {
	return &CodeGen{make(map[string]*Variable), 0, nil, nil}
}

func (gen *CodeGen) addError(e error) {
	gen.errors = append(gen.errors, e)
}

type IntInstr struct {
	Code     Instr
	Constant string
	Number   int
	Next     *IntInstr
	Target   *IntInstr
	n        int
}

func NewIntInstr(code Instr, constant string) *IntInstr {
	return &IntInstr{Code: code, Constant: constant}
}

func (instr *IntInstr) setNumbers(i int) {
	num := instr
	for num != nil {
		if num.Code != intTarget && num.Code != intIgnore {
			i++
		}
		num.n = i
		num = num.Next
	}
}

// Creates an error and returns an ignore instruction
func Errorf(format string, v ...interface{}) (*IntInstr, error) {
	return NewIntInstr(intIgnore, ""), fmt.Errorf(format, v...)
}

// Generates asm for getting a memory address
func (gen *CodeGen) getMemory(name string, offset int) (push *IntInstr, err error) {
	var pos string
	local := gen.locals[name]
	if local == nil {
		err = fmt.Errorf("undefined: %v", name)
	} else {
		var p int
		if local.typ == varArrTy {
			p = local.pos + (offset * local.size)
		} else {
			p = local.pos
		}
		pos = strconv.Itoa(p)
	}

	push = NewIntInstr(intPush, "")
	cons := NewIntInstr(intConst, pos)
	load := NewIntInstr(intMLoad, "")
	concat(push, cons)
	concat(cons, load)

	return
}

// Generates asm for setting a memory address
func (gen *CodeGen) setMemory(name string) (*IntInstr, error) {
	// TODO Only accept numbers. Lnegth checking will have to
	// occur when I implement strings
	local := gen.locals[name]
	if local == nil {
		return NewIntInstr(intIgnore, ""), fmt.Errorf("Undefined variable '%s'", name)
	}
	local.val = gen.lastPush.Constant

	push := NewIntInstr(intPush, "")
	cons := NewIntInstr(intConst, strconv.Itoa(local.pos))
	store := NewIntInstr(intMStore, "")

	concat(push, cons)
	concat(cons, store)

	return push, nil
}

// Type size table
var typeToSize = map[string]int{
	"int8":   8,
	"int16":  16,
	"int32":  32,
	"int64":  64,
	"int256": 256,
	"byte":   8,
	"big":    256,
}

func sizeOf(typ string) int {

	if typ == "big" {
		return 32
	}

	size, _ := strconv.Atoi(typ[3:])
	// Everything is 256 bit for now untill poc 5 comes along
	//size /= 8
	size = 32

	return size
}

func (gen *CodeGen) initNewVar(tree *SyntaxTree) error {
	name := tree.Constant
	if gen.locals[name] != nil {
		return fmt.Errorf("Redeclaration of variable '%s'", name)
	}

	var size int
	switch t := tree.VarType; {
	case typeToSize[t] > 0:
		size = sizeOf(t)
	}

	variable := &Variable{typ: varNumTy, pos: gen.memPos, size: size}
	gen.locals[name] = variable

	gen.memPos += size

	return nil
}

func (gen *CodeGen) sizeof(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	if gen.locals[name] == nil {
		return Errorf("undefined variable: '%s'", name)
	}

	local := gen.locals[name]

	push := NewIntInstr(intPush, "")
	size := strconv.Itoa(local.size)
	constant := NewIntInstr(intConst, size)

	return concat(push, constant), nil
}

func (gen *CodeGen) getArray(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	local := gen.locals[name]
	if local == nil {
		return NewIntInstr(intIgnore, ""), fmt.Errorf("undefined: %v", name)
	}

	// TODO optimize if the expression in offset. If regular const (i.e. 0-9)
	// do an inline calculation instead.

	// Get the location of the variable in memory
	loc := NewIntInstr(intPush, "")
	locConst := NewIntInstr(intConst, strconv.Itoa(local.pos))
	// Get the offset (= expression between brackets [expression])
	offset := gen.MakeIntCode(tree.Children[0])
	// Get the size of the variable in bytes
	size := NewIntInstr(intPush, "")
	sizeConst := NewIntInstr(intConst, strconv.Itoa(local.size))
	// Multiply offset with size
	mul := NewIntInstr(intMul, "")
	// Add the result to the memory location
	add := NewIntInstr(intAdd, "")
	// b = a[0] // loc(a) + sizeOf(type(a)) * len(a)
	load := NewIntInstr(intMLoad, "")

	concat(loc, locConst)
	concat(locConst, offset)
	concat(offset, size)
	concat(size, sizeConst)
	concat(sizeConst, mul)
	concat(mul, add)
	concat(add, load)

	return loc, nil
}

func (gen *CodeGen) setArray(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	local := gen.locals[name]
	if local == nil {
		return Errorf("undefined: %v", name)
	}

	if local.typ != varArrTy {
		return Errorf("left hand not an array")
	}

	// The value which we want to assign
	val := gen.MakeIntCode(tree.Children[1])

	// Get the location of the variable in memory
	loc := NewIntInstr(intPush, "")
	locConst := NewIntInstr(intConst, strconv.Itoa(local.pos))
	// Get the offset (= expression between brackets [expression])
	offset := gen.MakeIntCode(tree.Children[0])
	// Get the size of the variable in bytes
	size := NewIntInstr(intPush, "")
	sizeConst := NewIntInstr(intConst, strconv.Itoa(local.size))
	// Multiply offset with size
	mul := NewIntInstr(intMul, "")
	// Add the result to the memory location
	add := NewIntInstr(intAdd, "")
	store := NewIntInstr(intMStore, "")

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

func (gen *CodeGen) initNewArray(tree *SyntaxTree) error {
	name := tree.Constant
	if gen.locals[name] != nil {
		return fmt.Errorf("Redeclaration of variable '%s'", name)
	}

	var size int
	switch tree.VarType {
	case "int8", "int16", "int32", "int64", "int256", "big":
		size = sizeOf(tree.VarType)
	}
	length, _ := strconv.Atoi(tree.Size)
	size *= length
	fmt.Println("array l", size)

	variable := &Variable{typ: varArrTy, pos: gen.memPos, size: size}
	gen.locals[name] = variable

	gen.memPos += size

	return nil
}

func (gen *CodeGen) setStorage(tree *SyntaxTree) *IntInstr {
	fmt.Println(tree)
	return nil
}

// Concatenate two block of code together
func concat(blk1 *IntInstr, blk2 *IntInstr) *IntInstr {
	if blk2.Code == intEmpty {
		return blk1
	}

	if blk1.Code == intEmpty {
		return blk2
	}

	search := blk1
	for search.Next != nil {
		search = search.Next
	}

	search.Next = blk2

	return blk1
}

// Recursive intermediate code generator
func (gen *CodeGen) MakeIntCode(tree *SyntaxTree) *IntInstr {
	switch tree.Type {
	case StatementListTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := gen.MakeIntCode(tree.Children[1])

		return concat(blk1, blk2)
	case AssignmentTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := gen.MakeIntCode(tree.Children[1])
		c := concat(blk1, blk2)
		// Regular assignment a = 10
		if len(tree.Children) == 2 {
			return c
		}

		// Init assignment int8 a = 20
		blk3 := gen.MakeIntCode(tree.Children[2])
		return concat(c, blk3)
	case IfThenTy:
		cond := gen.MakeIntCode(tree.Children[0])
		not := NewIntInstr(intNot, "")
		jmpi := NewIntInstr(intJumpi, "")
		then := gen.MakeIntCode(tree.Children[1])
		target := NewIntInstr(intTarget, "")
		jmpi.Target = target

		concat(cond, not)
		concat(not, jmpi)
		concat(jmpi, then)
		concat(then, target)

		return cond
	case EqualTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := gen.MakeIntCode(tree.Children[1])
		concat(blk1, blk2)

		return concat(blk1, NewIntInstr(intEqual, ""))
	case IdentifierTy:
		c, err := gen.getMemory(tree.Constant, 0)
		if err != nil {
			gen.addError(err)
		}

		return c
	case ConstantTy:
		blk1 := NewIntInstr(intPush, "")
		blk2 := NewIntInstr(intConst, tree.Constant)
		gen.lastPush = blk2

		return concat(blk1, blk2)
	case SetLocalTy:
		c, err := gen.setMemory(tree.Constant)
		if err != nil {
			gen.addError(err)
		}

		return c
	case SetStoreTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := NewIntInstr(intSStore, "")

		return concat(blk1, blk2)
	case StoreTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := NewIntInstr(intSLoad, "")

		return concat(blk1, blk2)

	case OpTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := gen.MakeIntCode(tree.Children[1])
		concat(blk1, blk2)

		var op Instr

		switch tree.Constant {
		case "==":
			op = intEqual
		case ">":
			op = intGt
		}

		return concat(blk1, NewIntInstr(op, ""))
	case StopTy:
		return NewIntInstr(intStop, "")
	case OriginTy:
		return NewIntInstr(intOrigin, "")
	case CallerTy:
		return NewIntInstr(intCaller, "")
	case CallValTy:
		return NewIntInstr(intCallVal, "")
	case CallDataLoadTy:
		return NewIntInstr(intCallDataLoad, "")
	case CallDataSizeTy:
		return NewIntInstr(intCallDataSize, "")
	case GasPriceTy:
		return NewIntInstr(intGasPrice, "")
	case NewArrayTy: // Create a new array
		err := gen.initNewArray(tree)
		if err != nil {
			gen.addError(err)
		}

		return NewIntInstr(intIgnore, "")
	case ArrayTy: // Get a value of an array
		c, err := gen.getArray(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case AssignArrayTy: // Assign a value to an element in the array
		c, err := gen.setArray(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case NewVarTy:
		err := gen.initNewVar(tree)
		if err != nil {
			gen.addError(err)
			return NewIntInstr(intIgnore, "")
		}

		return NewIntInstr(intIgnore, "")
	case ArgTy:
		next := gen.MakeIntCode(tree.Children[0])
		arg := gen.MakeIntCode(tree.Children[1])

		return concat(arg, next)
	case CallTy:
		Push := func(t *SyntaxTree) (*IntInstr, error) {
			l := gen.locals[t.Constant]
			if l == nil {
				return Errorf("undefined variable: '%s'")
			}

			pushOff := NewIntInstr(intPush, "")
			offCons := NewIntInstr(intConst, strconv.Itoa(l.size+l.pos))
			pushLoc := NewIntInstr(intPush, "")
			locCons := NewIntInstr(intConst, strconv.Itoa(l.pos))

			concat(pushOff, offCons)
			concat(offCons, pushLoc)
			concat(pushLoc, locCons)

			return pushOff, nil
		}

		arg, err := Push(tree.Children[3])
		if err != nil {
			gen.addError(err)
			return arg
		}
		ret, err := Push(tree.Children[4])
		if err != nil {
			gen.addError(err)
			return ret
		}
		sender := gen.MakeIntCode(tree.Children[0])
		value := gen.MakeIntCode(tree.Children[1])
		gas := gen.MakeIntCode(tree.Children[2])
		call := NewIntInstr(intCall, "")

		concat(ret, arg)
		concat(arg, gas)
		concat(gas, value)
		concat(value, sender)
		concat(sender, call)

		return ret
	case SizeofTy:
		c, err := gen.sizeof(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case InlineAsmTy:
		// Remove tabs
		asm := strings.Replace(tree.Constant, "\t", "", -1)
		// Remove double spaces
		asm = strings.Replace(asm, "  ", " ", -1)
		asmSlice := strings.FieldsFunc(asm, func(r rune) bool {
			switch r {
			case '\n', ' ':
				return true
			}

			return false
		})

		firstInstr := NewIntInstr(intASM, asmSlice[0])
		if len(asmSlice) > 1 {
			for _, instr := range asmSlice[1:] {
				concat(firstInstr, NewIntInstr(intASM, instr))
			}
		}

		return firstInstr
	case EmptyTy:
		return NewIntInstr(intEmpty, "")
	}

	return nil
}

func (instr *IntInstr) String() string {
	str := fmt.Sprintf("%-3d %-12v : %v\n", instr.n, instr.Code, instr.Constant)
	if instr.Next != nil {
		str += instr.Next.String()
	}

	return str
}

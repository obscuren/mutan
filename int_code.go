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

	intIgnore
)

var instrAsString = []string{
	"equal",
	"gt",
	"lt",
	"mul",
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

// Generates asm for getting a memory address
func (gen *CodeGen) getMemory(name string) (push *IntInstr, err error) {
	var pos string
	if gen.locals[name] == nil {
		err = fmt.Errorf("undefined: %v", name)
	} else {
		pos = strconv.Itoa(gen.locals[name].pos)
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

		local = &Variable{typ: varNumTy, pos: gen.memPos}
		gen.locals[name] = local
		gen.memPos += 32
	}
	local.val = gen.lastPush.Constant

	push := NewIntInstr(intPush, "")
	cons := NewIntInstr(intConst, strconv.Itoa(local.pos))
	store := NewIntInstr(intMStore, "")

	concat(push, cons)
	concat(cons, store)

	return push, nil
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
	switch tree.VarType {
	case "int8", "int16", "int32", "int64", "int256", "big":
		size = sizeOf(tree.VarType)
	}

	variable := &Variable{typ: varNumTy, pos: gen.memPos, size: size}
	gen.locals[name] = variable

	gen.memPos += size

	return nil
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

	variable := &Variable{typ: varArrTy, pos: gen.memPos, size: size}
	gen.locals[name] = variable

	gen.memPos += size

	return nil
}

func (gen *CodeGen) setArray(name, size string) error {
	if gen.locals[name] != nil {
		return fmt.Errorf("%s already initialized. Can't overwrite", name)
	}

	s, _ := strconv.Atoi(size)
	gen.locals[name] = &Variable{typ: varArrTy, pos: gen.memPos, size: s}
	gen.memPos += s

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
		c, err := gen.getMemory(tree.Constant)
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
	case NewArrayTy:
		err := gen.initNewArray(tree)
		if err != nil {
			gen.addError(err)
		}

		return NewIntInstr(intIgnore, "")
	case NewVarTy:
		err := gen.initNewVar(tree)
		if err != nil {
			gen.addError(err)
			return NewIntInstr(intIgnore, "")
		}

		return NewIntInstr(intIgnore, "")
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

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
)

type Variable struct {
	typ varType
	val string
	pos int
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
		if num.Code != intTarget {
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
func (gen *CodeGen) setMemory(name string) *IntInstr {
	// TODO Only accept numbers. Lnegth checking will have to
	// occur when I implement strings
	local := gen.locals[name]
	if local == nil {
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

	return push
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

		return concat(blk1, blk2)
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
		return gen.setMemory(tree.Constant)
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

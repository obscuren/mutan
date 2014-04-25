package mutan

/*
int code generator takes care of memory allocation
generate all opcode and instructions
sets pc for each instruction afterwards

compiler transforms int code to ASM (very static)
*/

import (
	"encoding/hex"
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
	intExp
	intAssign
	intConst
	intEmpty
	intJump
	intTarget
	intPush
	intPush20
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
	intDiff
	intPrevHash
	intTimestamp
	intCoinbase
	intGas
	intBlockNum

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
	"exp",
	"assign",
	"const",
	"empty",
	"jump",
	"target",
	"push",
	"push20",
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
	"diff",
	"prevhash",
	"timestamp",
	"oinbase",
	"gas",
	"blocknum",

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

func (gen *CodeGen) Errors() []error {
	return gen.errors
}

func (gen *CodeGen) addError(e error) {
	gen.errors = append(gen.errors, e)
}

type IntInstr struct {
	Code      Instr
	Constant  interface{}
	Number    int
	Next      *IntInstr
	Target    *IntInstr
	TargetNum *IntInstr
	size      int
	n         int
}

func NewIntInstr(code Instr, constant string) *IntInstr {
	return &IntInstr{Code: code, Constant: constant}
}

func (instr *IntInstr) setNumbers(i int) {
	num := instr
	for num != nil {
		num.n = i

		if num.Code != intTarget && num.Code != intIgnore {
			i++
			if num.Code == intConst {
				i += 31
			}
		}

		num = num.Next
	}

	// XXX Do we need a 2 pass?
	num = instr
	for num != nil {
		if num.Code == intJump || num.Code == intJumpi {
			// Set the target constant which we couldn't seet before hand
			// when the numbers weren't all set.
			num.TargetNum.Constant = num.Target.n
		}

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
	local.val = gen.lastPush.Constant.(string)

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
	"string": 8,
	"addr":   160,
}

func sizeOf(typ string) int {

	if typ == "big" {
		return 32
	} else if typ == "addr" {
		return 20
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

	variable := &Variable{typ: varArrTy, pos: gen.memPos, size: size}
	gen.locals[name] = variable

	gen.memPos += size

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

func NewJumpInstr(op Instr) (*IntInstr, *IntInstr) {
	push := NewIntInstr(intPush, "")
	cons := NewIntInstr(intConst, "")
	jump := NewIntInstr(op, "")
	jump.TargetNum = cons
	concat(push, cons)
	concat(cons, jump)

	return push, jump
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
		//jmpi := NewIntInstr(intJumpi, "")
		p, jmpi := NewJumpInstr(intJumpi)
		then := gen.MakeIntCode(tree.Children[1])
		target := NewIntInstr(intTarget, "")
		jmpi.Target = target

		concat(cond, not)
		concat(not, p)
		concat(jmpi, then)
		concat(then, target)

		return cond
	case IfThenElseTy:
		cond := gen.MakeIntCode(tree.Children[0])
		// TODO reverse "else/if" in order to get rid of the "NOT"
		not := NewIntInstr(intNot, "")
		//jump2else := NewIntInstr(intJumpi, "")
		p, jump2else := NewJumpInstr(intJumpi)
		then := gen.MakeIntCode(tree.Children[1])
		//jump2end := NewIntInstr(intJump, "")
		p2, jump2end := NewJumpInstr(intJump)
		elsetarget := NewIntInstr(intTarget, "")
		elsethen := gen.MakeIntCode(tree.Children[2])
		end := NewIntInstr(intTarget, "")

		jump2end.Target = end
		jump2else.Target = elsetarget

		concat(cond, not)
		concat(not, p)
		concat(jump2else, then)
		concat(then, p2)
		concat(jump2end, elsetarget)
		concat(elsetarget, elsethen)
		concat(elsethen, end)

		return cond
	case ForThenTy:
		// Init part
		init := gen.MakeIntCode(tree.Children[0])
		// The condition for the loop
		cond := gen.MakeIntCode(tree.Children[1])
		// Cast to not
		not := NewIntInstr(intNot, "")
		// Jump to end if statement is false
		//jmpi := NewIntInstr(intJumpi, "")
		p, jmpi := NewJumpInstr(intJumpi)
		// Body of the loop
		then := gen.MakeIntCode(tree.Children[3])
		// Iterator
		aft := gen.MakeIntCode(tree.Children[2])
		// Jump back to the start of the loop (targetBack)
		//jmp := NewIntInstr(intJump, "")
		p2, jmp := NewJumpInstr(intJump)
		// Target for the conditional jump (jmpi)
		target := NewIntInstr(intTarget, "")
		// Set targets
		jmpi.Target = target
		jmp.Target = cond

		concat(init, cond)
		concat(cond, not)
		concat(not, p)
		concat(jmpi, then)
		concat(then, aft)
		concat(aft, p2)
		concat(jmp, target)

		return init
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

		var op Instr
		switch tree.Constant {
		case "==":
			op = intEqual
		case ">":
			op = intGt
		case "<":
			op = intLt
		case "^":
			op = intExp
		case "*":
			op = intMul
		case "/":
			op = intDiv
		case "+":
			op = intAdd
		case "-":
			op = intSub
		case "++", "--":
			one := NewIntInstr(intPush, "")
			cons := NewIntInstr(intConst, "1")
			if tree.Constant == "++" {
				op = intAdd
			} else {
				op = intSub
			}
			opInstr := NewIntInstr(op, "")
			concat(blk1, one)
			concat(one, cons)
			concat(cons, opInstr)
			// Get the child
			child := tree.Children[0]
			if child.Type == IdentifierTy {
				c, err := gen.setMemory(child.Constant)
				if err != nil {
					gen.addError(err)
					return c
				}

				concat(opInstr, c)
			} else {
				// TODO?
				c, err := Errorf("++ only supported on identifiers")
				gen.addError(err)
				return c
			}

			return blk1
		}

		blk2 := gen.MakeIntCode(tree.Children[1])
		concat(blk1, blk2)

		return concat(blk1, NewIntInstr(op, ""))
	case StringTy:
		blk1 := NewIntInstr(intPush20, "")
		byts, err := hex.DecodeString(tree.Constant)
		if err != nil {
			st, e := Errorf("%v: %s", err, tree.Constant)

			gen.addError(e)

			return st
		}
		blk2 := NewIntInstr(intConst, string(byts))
		gen.lastPush = blk2

		return concat(blk1, blk2)
	case StopTy:
		return NewIntInstr(intStop, "")
	case OriginTy:
		return NewIntInstr(intOrigin, "")
	case CallerTy:
		return NewIntInstr(intCaller, "")
	case CallValTy:
		return NewIntInstr(intCallVal, "")
	case CallDataLoadTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		// XXX There's room for optimization here
		push := NewIntInstr(intPush, "")
		cons := NewIntInstr(intConst, "32")
		mul := NewIntInstr(intMul, "")
		blk2 := NewIntInstr(intCallDataLoad, "")
		concat(blk1, push)
		concat(push, cons)
		concat(cons, mul)
		concat(mul, blk2)

		return blk1
	case CallDataSizeTy:
		return NewIntInstr(intCallDataSize, "")
	case GasPriceTy:
		return NewIntInstr(intGasPrice, "")
	case DiffTy:
		return NewIntInstr(intDiff, "")
	case PrevHashTy:
		return NewIntInstr(intPrevHash, "")
	case TimestampTy:
		return NewIntInstr(intTimestamp, "")
	case CoinbaseTy:
		return NewIntInstr(intCoinbase, "")
	case GasTy:
		return NewIntInstr(intGas, "")
	case BlockNumTy:
		return NewIntInstr(intBlockNum, "")
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
			offCons := NewIntInstr(intConst, strconv.Itoa(l.size))
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

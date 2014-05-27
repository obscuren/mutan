package mutan

/*
int code generator takes care of memory allocation
generate all opcode and instructions
sets pc for each instruction afterwards

compiler transforms int code to ASM (very static)
*/

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

type varType byte

const (
	varUndefinedTy varType = iota
	varNumTy
	varStrTy
	varArrTy
)

type Variable struct {
	id      string
	typ     varType
	pos     int
	size    int
	varSize int

	instr *IntInstr
}

type CodeGen struct {
	locals      map[string]*Variable
	arrayTable  map[string][]*IntInstr
	stringTable map[string][]*IntInstr

	memPos int
	//lastPush *IntInstr

	errors []error
}

func NewGen() *CodeGen {
	return &CodeGen{make(map[string]*Variable), make(map[string][]*IntInstr), make(map[string][]*IntInstr), 0, nil}
}

func (gen *CodeGen) Errors() []error {
	return gen.errors
}

func (gen *CodeGen) addError(e error) {
	gen.errors = append(gen.errors, e)
}

// Generates asm for getting a memory address
func (gen *CodeGen) getMemory(tree *SyntaxTree, offset int) (push *IntInstr, err error) {
	var pos string
	local := gen.locals[tree.Constant]
	if local == nil {
		return tree.Errorf("Undefined variable: %v", tree.Constant)
	} else {
		var p int
		if local.typ == varArrTy {
			p = local.pos + (offset * local.size)
		} else {
			p = local.pos
		}
		pos = strconv.Itoa(p)
	}

	push, cons := pushConstant(pos)
	load := NewIntInstr(intMLoad, "")
	concat(push, cons)
	concat(cons, load)

	return
}

// Generates asm for setting a memory address
func (gen *CodeGen) setMemory(tree *SyntaxTree) (*IntInstr, error) {
	// TODO Only accept numbers. Lnegth checking will have to
	// occur when I implement strings
	local := gen.locals[tree.Constant]
	if local == nil {
		return tree.Errorf("Undefined variable '%s'", tree.Constant)
	}

	push, cons := pushConstant(strconv.Itoa(local.pos))
	store := NewIntInstr(intMStore, "")

	concat(push, cons)
	concat(cons, store)

	return push, nil
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

func (gen *CodeGen) initNewVar(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	if gen.locals[name] != nil {
		return tree.Errorf("Redeclaration of variable '%s'", name)
	}

	/*
		var size int
		switch t := tree.VarType; {
		case typeToSize[t] > 0:
			size = sizeOf(t)
		default:
			return tree.Errorf("undefined type %s", tree.VarType)
		}
	*/
	size := 32

	//variable := &Variable{typ: varNumTy, pos: gen.memPos, size: size, varSize: size}
	variable := &Variable{id: name, size: size, varSize: size}
	gen.locals[name] = variable

	//gen.memPos += size

	return nil, nil
}

func (gen *CodeGen) sizeof(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	if gen.locals[name] == nil {
		return tree.Errorf("undefined variable: '%s'", name)
	}

	local := gen.locals[name]
	push, constant := pushConstant(strconv.Itoa(local.size))

	return concat(push, constant), nil
}

func (gen *CodeGen) getArray(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	local := gen.locals[name]
	if local == nil {
		tree.Errorf("undefined array: %v", name)
	}

	// TODO optimize if the expression in offset. If regular const (i.e. 0-9)
	// do an inline calculation instead.

	// Get the location of the variable in memory
	loc, locConst := pushConstant(strconv.Itoa(local.pos))
	// Get the offset (= expression between brackets [expression])
	offset := gen.MakeIntCode(tree.Children[0])
	// Get the size of the variable in bytes
	size, sizeConst := pushConstant(strconv.Itoa(local.size))
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
		return tree.Errorf("undefined array: %v", name)
	}

	if local.typ != varArrTy {
		return tree.Errorf("left hand not an array")
	}

	// The value which we want to assign
	val := gen.MakeIntCode(tree.Children[1])

	// Get the location of the variable in memory
	loc, locConst := pushConstant(strconv.Itoa(local.pos))
	gen.arrayTable[name] = append(gen.arrayTable[name], locConst)

	// Get the offset (= expression between brackets [expression])
	offset := gen.MakeIntCode(tree.Children[0])
	// Get the size of the variable in bytes
	size, sizeConst := pushConstant(strconv.Itoa(local.varSize))
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

func (gen *CodeGen) initNewArray(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	if gen.locals[name] != nil {
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

	variable := &Variable{id: name, typ: varArrTy, size: 32 * length, varSize: 32}
	gen.locals[name] = variable

	return NewIntInstr(intIgnore, ""), nil
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
	push := NewIntInstr(intPush4, "")
	cons := NewIntInstr(intConst, "")
	cons.size = 4
	cons.Target = push

	jump := NewIntInstr(op, "")
	jump.TargetNum = cons
	concat(push, cons)
	concat(cons, jump)

	return push, jump
}

func validLhSide(variable *Variable, typ varType) {
	// ignore
	if variable == nil {
		return
	}

	if variable.typ != varUndefinedTy && variable.typ != typ {
		panic(fmt.Sprintf("cannat assign %v to '%v' of type %v", typ, variable.id, variable.typ))
	}
}

func (gen *CodeGen) setVariable(tree *SyntaxTree, identifier *SyntaxTree) *IntInstr {
	var instr *IntInstr

	name := identifier.Constant
	id := gen.locals[identifier.Constant]

	//TODO Do left hand side type checking at this point
	switch tree.Type {
	case StringTy:
		validLhSide(id, varStrTy)

		var t Instr
		if identifier.Type == SetStoreTy {
			if len(tree.Constant) > 32 {
				panic("attempting to store string greater than 32")
			}

			return gen.makePush("0x" + hex.EncodeToString([]byte(tree.Constant)))
		} else {
			t = intMStore
			gen.locals[name].typ = varStrTy
		}

		var length int
		//instr, length = gen.bytesToHexInstr(id.pos, []byte(tree.Constant))
		instr, length = gen.stringToInstr(id, []byte(tree.Constant), t)

		if identifier.Type != SetStoreTy {
			id.size = int(math.Max(math.Max(float64(id.size), float64(length)), 32.0))
		}

		return instr
	case ConstantTy:
		validLhSide(id, varNumTy)

		if identifier.Type != SetStoreTy {
			id.typ = varNumTy
		}

		instr = gen.MakeIntCode(tree)
	default:
		instr = gen.MakeIntCode(tree)
	}

	return instr
}

// Recursive intermediate code generator
func (gen *CodeGen) MakeIntCode(tree *SyntaxTree) *IntInstr {
	switch tree.Type {
	case StatementListTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := gen.MakeIntCode(tree.Children[1])

		return concat(blk1, blk2)
	case AssignmentTy:
		var blk1 *IntInstr
		if len(tree.Children) == 2 {
			blk2 := gen.MakeIntCode(tree.Children[1])
			blk1 = gen.setVariable(tree.Children[0], tree.Children[1])
			concat(blk1, blk2)
			/*
				if tree.Children[0].Type != StringTy {
					concat(blk1, blk2)
				}
			*/
		} else {
			blk2 := gen.MakeIntCode(tree.Children[1])
			blk3 := gen.MakeIntCode(tree.Children[2])
			gen.locals[tree.Children[2].Constant].instr = blk3.Next
			blk1 = gen.setVariable(tree.Children[0], tree.Children[2])
			concat(blk1, blk2)
			// In case the type is a string we do _not_ want to concat blk3
			// because it handles all the MSTORE / etc itself
			if tree.Children[0].Type != StringTy {
				concat(blk2, blk3)
			}

		}

		return blk1
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
		c, err := gen.getMemory(tree, 0)
		if err != nil {
			gen.addError(err)
		}

		return c
	case ConstantTy:
		blk1, blk2 := pushConstant(tree.Constant)
		concat(blk1, blk2)

		return blk1
	case BoolTy:
		var value string
		if tree.Constant == "true" {
			value = "1"
		} else {
			value = "0"
		}
		blk1, blk2 := pushConstant(value)
		concat(blk1, blk2)

		return blk1
	case SetLocalTy:
		c, err := gen.setMemory(tree)
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
		// TODO clean this up
		var blk1, blk2, blk3 *IntInstr
		blk1 = gen.MakeIntCode(tree.Children[0])

		var op Instr
		switch tree.Constant {
		case "==":
			op = intEqual
		case ">":
			op = intGt
		case "<":
			op = intLt
		case "*":
			op = intMul
		case "/":
			op = intDiv
		case "+":
			op = intAdd
		case "-":
			op = intSub
		case "%":
			op = intMod
		case "&":
			op = intAnd
		case "|":
			op = intOr
		case "^":
			op = intXor
		case "**":
			op = intExp
		case "<<", ">>":
			push, cons := pushConstant("2")
			blk2 := gen.MakeIntCode(tree.Children[1])
			exp := NewIntInstr(intExp, "")
			var o *IntInstr
			if tree.Constant == "<<" {
				o = NewIntInstr(intMul, "")
			} else {
				o = NewIntInstr(intDiv, "")
			}
			concat(push, cons)
			concat(cons, blk2)
			concat(blk2, exp)
			concat(exp, blk1)
			concat(blk1, o)

			return push
		case "++", "--":
			one, cons := pushConstant("1")
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
				c, err := gen.setMemory(child)
				if err != nil {
					gen.addError(err)
					return c
				}

				concat(opInstr, c)
			} else {
				// TODO?
				c, err := tree.Errorf("++ only supported on identifiers")
				gen.addError(err)
				return c
			}

			return blk1
		case ">=":
			op = intLt
			blk3 = NewIntInstr(intNot, "")
		case "<=":
			op = intGt
			blk3 = NewIntInstr(intNot, "")
		case "!=":
			op = intEqual
			blk3 = NewIntInstr(intNot, "")
		case "!":
			// Reconstruct this one (We ought to clean this code)
			blk1 = gen.MakeIntCode(tree.Children[1])
			opinstr := NewIntInstr(intNot, "")
			return concat(blk1, opinstr)

		default:
			c, err := tree.Errorf("Expected operator, got '%v'", tree.Constant)
			gen.addError(err)
			return c
		}

		blk2 = gen.MakeIntCode(tree.Children[1])
		concat(blk1, blk2)

		if blk3 != nil {
			opinstr := NewIntInstr(op, "")
			concat(blk2, opinstr)
			concat(opinstr, blk3)

			return blk1
		}

		return concat(blk1, NewIntInstr(op, ""))
		/*
			case StringTy:
				blk1 := NewIntInstr(intPush20, "")
				byts, err := hex.DecodeString(tree.Constant)
				if err != nil {
					st, e := tree.Errorf("%v: %s", err, tree.Constant)

					gen.addError(e)

					return st
				}
				blk2 := NewIntInstr(intConst, string(byts))
				//gen.lastPush = blk2

				return concat(blk1, blk2)
		*/
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
		push, cons := pushConstant("32")
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
	case BalanceTy:
		return NewIntInstr(intBalance, "")
	case GasTy:
		return NewIntInstr(intGas, "")
	case BlockNumTy:
		return NewIntInstr(intBlockNum, "")
	case NewArrayTy: // Create a new array
		c, err := gen.initNewArray(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
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
		c, err := gen.initNewVar(tree)
		if err != nil {
			gen.addError(err)

			return c
		}

		return NewIntInstr(intIgnore, "")
	case ArgTy:
		next := gen.MakeIntCode(tree.Children[0])
		arg := gen.MakeIntCode(tree.Children[1])

		return concat(arg, next)
	case CallTy:
		arg, err := gen.makeArg(tree.Children[3])
		if err != nil {
			gen.addError(err)
			return arg
		}
		ret, err := gen.makeArg(tree.Children[4])
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
	case TransactTy:
		ret := gen.pushNil()
		arg, err := gen.makeArg(tree.Children[2])
		if err != nil {
			gen.addError(err)
			return arg
		}
		sender := gen.MakeIntCode(tree.Children[0])
		value := gen.MakeIntCode(tree.Children[1])
		gas := gen.makePush("0")
		call := NewIntInstr(intCall, "")

		concat(ret, arg)
		concat(arg, gas)
		concat(gas, value)
		concat(value, sender)
		concat(sender, call)

		return ret
	case CreateTy:
		script, err := gen.makeArg(tree.Children[1])
		if err != nil {
			gen.addError(err)
			return script
		}
		val := gen.MakeIntCode(tree.Children[0])
		create := NewIntInstr(intCreate, "")

		concat(script, val)
		concat(val, create)

		return script
	case ReturnTy:
		switch tree.Children[0].Type {
		case LambdaTy:
			retVal, num := gen.compileLambda(0, tree.Children[0])
			if num != 0 {
				size := gen.makePush("0")
				offset := gen.makePush(strconv.Itoa(num))
				concat(retVal, offset)
				concat(offset, size)

				return concat(retVal, NewIntInstr(intReturn, ""))
			}

		default:
			c, err := tree.Errorf("return; only supports Lambda")
			gen.addError(err)

			return c
		}

		return NewIntInstr(intIgnore, "")
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
	case LambdaTy:
		panic("auto lambda triggered in int code gen. report this issue")
	case EmptyTy:
		return NewIntInstr(intEmpty, "")
	}

	return nil
}

// Compiles the given code and stores it at memory position given by the mem offset
// Returns the instructions necessary to handle this lambda and returns the maximum size
// e.g. if the code script was 50 bytes, it would generate 2 memory storages instructions
// for: [0-31] & [32-50] and returns (50)
func (gen *CodeGen) compileLambda(memOffset int, tree *SyntaxTree) (*IntInstr, int) {
	code, errors := Compile(strings.NewReader(tree.Constant), false)
	if len(errors) != 0 {
		gen.errors = append(gen.errors, errors...)
		return nil, 0
	}

	return gen.bytesToHexInstr(memOffset, code)

}

func (gen *CodeGen) bytesToHexInstr(memOffset int, b []byte) (*IntInstr, int) {
	ignore := NewIntInstr(intIgnore, "")
	var lastPush *IntInstr
	i := 0
	for ; i < len(b); i += 32 {
		offset := int(math.Min(float64(i+32), float64(len(b))))
		hex := hex.EncodeToString(b[i:offset])
		mem := gen.makePush(strconv.Itoa(i + memOffset))
		push := gen.makePush("0x" + hex)
		store := NewIntInstr(intMStore, "")
		concat(push, mem)
		concat(mem, store)

		if lastPush != nil {
			concat(lastPush, push)
		} else {
			concat(ignore, push)
		}

		lastPush = push
	}

	return ignore, i
}

func (gen *CodeGen) stringToInstr(id *Variable, b []byte, t Instr) (*IntInstr, int) {
	ignore := NewIntInstr(intIgnore, "")
	var lastPush *IntInstr
	i := 0
	for ; i < len(b); i += 32 {
		offset := int(math.Min(float64(i+32), float64(len(b))))
		hex := hex.EncodeToString(b[i:offset])
		mem := gen.makePush(strconv.Itoa(i))

		if t != intSStore {
			gen.stringTable[id.id] = append(gen.stringTable[id.id], mem.Next)
		}

		push := gen.makePush("0x" + hex)
		store := NewIntInstr(t, "")
		concat(push, mem)
		concat(mem, store)

		if lastPush != nil {
			concat(lastPush, push)
		} else {
			concat(ignore, push)
		}

		lastPush = push
	}

	return ignore, i
}

// XXX This is actually a range function. FIXME
func (gen *CodeGen) makeArg(t *SyntaxTree) (*IntInstr, error) {
	if t.Type == NilTy {
		return gen.pushNil(), nil
	}

	l := gen.locals[t.Constant]
	if l == nil {
		return t.Errorf("undefine variable: '%s'", t.Constant)
	}

	pushOff, offCons := pushConstant(strconv.Itoa(l.size))
	pushLoc, locCons := pushConstant(strconv.Itoa(l.pos))

	concat(pushOff, offCons)
	concat(offCons, pushLoc)
	concat(pushLoc, locCons)

	return pushOff, nil
}

func (gen *CodeGen) makePush(num string) *IntInstr {
	push, cons := pushConstant(num)
	concat(push, cons)

	return push
}

func (gen *CodeGen) pushNil() *IntInstr {
	p1 := concat(gen.makePush("0"), gen.makePush("0"))
	p2 := concat(gen.makePush("0"), gen.makePush("0"))

	return concat(p1, p2)
}

func constToPush(constant string) (*IntInstr, int) {
	num, _ := new(big.Int).SetString(constant, 0)
	numBytes := len(num.Bytes())
	if numBytes == 0 {
		numBytes = 1
	}
	return NewIntInstr(Instr(int(intPush1)-1+numBytes), ""), numBytes
}

func pushConstant(constant string) (*IntInstr, *IntInstr) {
	blk1, numBytes := constToPush(constant)
	blk2 := NewIntInstr(intConst, constant)
	blk2.size = numBytes

	return blk1, blk2
}

func numberToBytes(num interface{}, bits int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()[buf.Len()-(bits/8):]
}

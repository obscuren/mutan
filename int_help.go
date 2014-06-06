package mutan

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math"
	"math/big"
	"strconv"
	"strings"
)

/*********
 * STACK PTR HELPERS
 */

func (self *IntGen) loadStackPtr() *IntInstr {
	push := self.makePush("0")
	mload := newIntInstr(intMLoad, "")
	concat(push, mload)

	return push
}

func (self *IntGen) setStackPtr(ptr int) *IntInstr {
	self.currentStackSize = ptr

	push := self.makePush(strconv.Itoa(ptr))
	push2 := self.makePush("0")
	mstore := newIntInstr(intMStore, "")
	concat(push, push2)
	concat(push2, mstore)

	return push
}

func (self *IntGen) addStackPtr(size int) *IntInstr {
	push := self.makePush(strconv.Itoa(size))
	load := self.loadStackPtr()
	add := newIntInstr(intAdd, "")
	loc := self.makePush("0")
	store := newIntInstr(intMStore, "")

	concat(push, load)
	concat(load, add)
	concat(add, loc)
	concat(loc, store)

	return push
}

/******************
 * END
 */

// XXX This is actually a range function. FIXME
func (gen *IntGen) makeArg(t *SyntaxTree) (*IntInstr, error) {
	if t.Type == NilTy {
		return gen.pushNil(), nil
	}

	name := t.Constant
	variable := gen.CurrentScope().GetVar(name)
	if variable == nil {
		return t.Errorf("undefine variable: '%s'", t.Constant)
	}

	pushOff, offCons := pushConstant(strconv.Itoa(variable.Size()))
	pushLoc, locCons := pushConstant(strconv.Itoa(variable.Offset()))
	locCons.ConstRef = t.Constant

	concat(pushOff, offCons)
	concat(offCons, pushLoc)
	concat(pushLoc, locCons)

	return pushOff, nil
}

func (gen *IntGen) makeString(tree *SyntaxTree) *IntInstr {
	byts := []byte(tree.Constant)
	hexStr := "0x" + hex.EncodeToString(byts)

	push := newIntInstr(Instr(int(intPush1)-1+len(byts)), "")
	cons := newIntInstr(intConst, hexStr)
	cons.size = len(byts)

	concat(push, cons)

	return push
}

func (gen *IntGen) makePush(num string, v ...int) *IntInstr {
	var push, cons *IntInstr
	if len(v) > 0 {
		n, _ := strconv.Atoi(num)
		push = newIntInstr(intPush4, "")
		cons = newIntInstr(intConst, "0x"+hex.EncodeToString(numberToBytes(int32(n), v[0])))
		cons.size = v[0]
	} else {
		push, cons = pushConstant(num)
	}

	concat(push, cons)

	return push
}

func (gen *IntGen) pushNil() *IntInstr {
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
	return newIntInstr(Instr(int(intPush1)-1+numBytes), ""), numBytes
}

func pushConstant(constant string) (*IntInstr, *IntInstr) {
	blk1, numBytes := constToPush(constant)
	blk2 := newIntInstr(intConst, constant)
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

// Compiles the given code and stores it at memory position given by the mem offset
// Returns the instructions necessary to handle this lambda and returns the maximum size
// e.g. if the code script was 50 bytes, it would generate 2 memory storages instructions
// for: [0-31] & [32-50] and returns (50)
func (gen *IntGen) compileLambda(memOffset int, tree *SyntaxTree) (*IntInstr, int) {
	code, errors := Compile(strings.NewReader(tree.Constant), false)
	if len(errors) != 0 {
		gen.errors = append(gen.errors, errors...)
		return nil, 0
	}

	return gen.bytesToHexInstr(memOffset, code)

}

func (gen *IntGen) bytesToHexInstr(memOffset int, b []byte) (*IntInstr, int) {
	ignore := newIntInstr(intIgnore, "")
	var lastPush *IntInstr
	i := 0
	l := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for ; i < len(b); i += 32 {
		offset := int(math.Min(float64(i+32), float64(len(b))))
		hex := hex.EncodeToString(append(b[i:offset], l[0:32-len(b[i:offset])]...))
		mem := gen.makePush(strconv.Itoa(i + memOffset))
		push := gen.makePush("0x" + hex)
		store := newIntInstr(intMStore, "")
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

func (gen *IntGen) stringToInstr(variable Var, b []byte, t Instr) (*IntInstr, int) {
	ignore := newIntInstr(intIgnore, "")
	var lastPush *IntInstr
	i := 0
	// TODO clean this up. Strings can no longer be longer than 32 bytes
	for ; i < len(b); i += 32 {
		offset := int(math.Min(float64(i+32), float64(len(b))))
		hex := hex.EncodeToString(b[i:offset])
		mem := gen.makePush(strconv.Itoa(i))

		if t != intSStore {
			gen.stringTable[variable.Id()] = append(gen.stringTable[variable.Id()], mem.Next)
		}

		push := gen.makePush("0x" + hex)
		store := newIntInstr(t, "")
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

// Generates asm for getting a memory address
func (gen *IntGen) getMemory(tree *SyntaxTree) (*IntInstr, error) {
	variable := gen.GetVar(tree.Constant)
	if variable == nil {
		return tree.Errorf("Undefined variable: %v", tree.Constant)
	}

	offset := strconv.Itoa(variable.Offset())
	rPos := gen.loadStackPtr()
	push, cons := pushConstant(offset)
	add := newIntInstr(intAdd, "")
	load := newIntInstr(intMLoad, "")

	concat(rPos, push)
	concat(push, cons)
	concat(cons, add)
	concat(add, load)

	return rPos, nil
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

func (gen *IntGen) initNewNumber(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant

	scope := gen.CurrentScope()
	_, err := scope.NewVar(name, varUndefinedTy)
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

	length, _ := strconv.Atoi(tree.Size)
	variable.SetSize(32 * length)

	return newIntInstr(intIgnore, ""), nil
}

func (gen *IntGen) setVariable(tree *SyntaxTree, identifier *SyntaxTree) *IntInstr {
	variable := gen.GetVar(identifier.Constant)

	var instr *IntInstr
	switch tree.Type {
	case StringTy:
		if len(tree.Constant) > 32 {
			c, err := tree.Errorf("attempting to store string greater than 32 bytes")
			gen.addError(err)

			return c
		}

		var t Instr
		if identifier.Type == SetStoreTy {
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
		if identifier.Type != SetStoreTy {
			variable.SetType(varNumTy)
		}

		instr = gen.MakeIntCode(tree)
	default:
		instr = gen.MakeIntCode(tree)
	}

	return instr
}

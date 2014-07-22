package frontend

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

func __ignore() { fmt.Sprintf("") }

/*********
 * STACK PTR HELPERS
 */

func (self *IntGen) loadStackPtr() *IntInstr {
	push := self.makePush("0")
	mload := newIntInstr(IntMLoad, "")
	concat(push, mload)

	return push
}

func (self *IntGen) SetStackPtr(ptr int) *IntInstr {
	self.currentStackSize = ptr

	push := self.makePush(strconv.Itoa(ptr))
	push2 := self.makePush("0")
	mstore := newIntInstr(IntMStore, "")
	concat(push, push2)
	concat(push2, mstore)

	return push
}

func (self *IntGen) addStackPtr(size int) *IntInstr {
	push := self.makePush(strconv.Itoa(size))
	load := self.loadStackPtr()
	add := newIntInstr(IntAdd, "")
	loc := self.makePush("0")
	store := newIntInstr(IntMStore, "")

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
		return gen.Errorf(t, "undefine variable: '%s'", t.Constant)
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

	push := newIntInstr(Instr(int(IntPush1)-1+len(byts)), "")
	cons := newIntInstr(IntConst, hexStr)
	cons.size = len(byts)

	concat(push, cons)

	return push
}

func (gen *IntGen) makePush(num string, v ...int) *IntInstr {
	var push, cons *IntInstr
	if len(v) > 0 {
		n, _ := strconv.Atoi(num)
		push = newIntInstr(IntPush4, "")
		cons = newIntInstr(IntConst, "0x"+hex.EncodeToString(numberToBytes(int32(n), v[0])))
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
	num, ok := new(big.Int).SetString(constant, 0)
	if !ok {
		return newIntInstr(IntIgnore, ""), 0
	}

	numBytes := len(num.Bytes())
	if numBytes == 0 {
		numBytes = 1
	}

	return newIntInstr(Instr(int(IntPush1)-1+numBytes), ""), numBytes

}

func pushConstant(constant string) (*IntInstr, *IntInstr) {
	blk1, numBytes := constToPush(constant)
	blk2 := newIntInstr(IntConst, constant)
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

func (gen *IntGen) compileLambda(memOffset int, tree *SyntaxTree) (*IntInstr, int) {
	code, errors := Compiler.Compile(strings.NewReader(tree.Constant))
	if len(errors) != 0 {
		gen.Errors = append(gen.Errors, errors...)
		return newIntInstr(IntIgnore, ""), 0
	}

	var (
		inlineCode = InlineCode{code, nil, 0}
		length     = gen.makePush(strconv.Itoa(len(code)))
		cOffset    = newIntInstr(IntPush4, "")
		cConst     = newIntInstr(IntConst, "")
		//mOffset    = newIntInstr(IntMSize, "") // gen.makePush("0")
		mOffset  = gen.makePush("0")
		codecopy = newIntInstr(IntCodeCopy, "")
	)

	cConst.Constant = "0x" + hex.EncodeToString(numberToBytes(int32(0), 32))
	cConst.size = 4
	inlineCode.OffsetInstr = cConst
	cc(length, cOffset, cConst, mOffset, codecopy)

	gen.InlineCode = append(gen.InlineCode, inlineCode)

	return length, len(code)
}

func (gen *IntGen) bytesToHexInstr(memOffset int, b []byte) (*IntInstr, int) {
	ignore := newIntInstr(IntIgnore, "")
	var lastPush *IntInstr
	i := 0
	l := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for ; i < len(b); i += 32 {
		offset := int(math.Min(float64(i+32), float64(len(b))))
		hex := hex.EncodeToString(append(b[i:offset], l[0:32-len(b[i:offset])]...))
		mem := gen.makePush(strconv.Itoa(i + memOffset))
		push := gen.makePush("0x" + hex)
		store := newIntInstr(IntMStore, "")
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
	ignore := newIntInstr(IntIgnore, "")
	var lastPush *IntInstr
	i := 0
	// TODO clean this up. Strings can no longer be longer than 32 bytes
	for ; i < len(b); i += 32 {
		offset := int(math.Min(float64(i+32), float64(len(b))))
		hex := hex.EncodeToString(b[i:offset])
		mem := gen.makePush(strconv.Itoa(i))

		if t != IntSStore {
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

func (gen *IntGen) getMemoryAddress(tree *SyntaxTree) (*IntInstr, error) {
	variable := gen.GetVar(tree.Constant)
	if variable == nil {
		return gen.Errorf(tree, "Undefined variable: %v", tree.Constant)
	}

	return gen.getAddress(variable), nil
}

// Generates asm for getting a memory address
func (gen *IntGen) getMemory(tree *SyntaxTree) (*IntInstr, error) {
	variable := gen.GetVar(tree.Constant)
	if variable == nil {
		return gen.Errorf(tree, "Undefined variable: %v", tree.Constant)
	}

	offset := strconv.Itoa(variable.Offset())
	rPos := gen.loadStackPtr()
	push, cons := pushConstant(offset)
	add := newIntInstr(IntAdd, "")
	load := newIntInstr(IntMLoad, "")

	concat(rPos, push)
	concat(push, cons)
	concat(cons, add)
	concat(add, load)

	return rPos, nil
}

func makeStore(offset int) *IntInstr {
	push, cons := pushConstant(strconv.Itoa(offset))
	store := newIntInstr(IntMStore, "")

	concat(push, cons)
	concat(cons, store)

	return push
}

func (gen *IntGen) assignMemory(offset int) *IntInstr {
	ptr := gen.loadStackPtr()
	push, cons := pushConstant(strconv.Itoa(offset))
	add := newIntInstr(IntAdd, "")
	store := newIntInstr(IntMStore, "")

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
		return gen.Errorf(tree, "Undefined variable '%s'", tree.Constant)
	}

	instr := gen.assignMemory(variable.Offset())

	return instr, nil
}

func (gen *IntGen) initNewNumber(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant

	scope := gen.CurrentScope()
	var err error
	if tree.Ptr {
		_, err = scope.NewVar(name, varPtrTy)
	} else {
		_, err = scope.NewVar(name, varUndefinedTy)
	}

	if err != nil {
		return gen.Errorf(tree, "%v", err)
	}

	return nil, nil
}

func (gen *IntGen) sizeof(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)

	if variable == nil {
		return gen.Errorf(tree, "undefined variable: '%s'", name)
	}

	push, constant := pushConstant(strconv.Itoa(variable.Size()))

	return concat(push, constant), nil
}

func (gen *IntGen) getAddress(v Var) *IntInstr {
	// Get stack ptr
	ptr := gen.loadStackPtr()
	// Push the variable offset
	push, cons := pushConstant(strconv.Itoa(v.Offset()))
	add := newIntInstr(IntAdd, "")

	cc(ptr, push, cons, add)

	return ptr
}

func (gen *IntGen) pushDerefPtr(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)

	if variable == nil {
		return gen.Errorf(tree, "undefined array: %v", name)
	}

	ptr := gen.dereferencePtr(variable)
	loadRef := newIntInstr(IntMLoad, "")
	cc(ptr, loadRef)

	return ptr, nil
}

func (gen *IntGen) dereferencePtr(v Var) *IntInstr {
	ptr := gen.getAddress(v)
	loadPtr := newIntInstr(IntMLoad, "")

	cc(ptr, loadPtr)

	return ptr
}

func (gen *IntGen) setPtr(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)

	if variable == nil {
		return gen.Errorf(tree, "undefined array: %v", name)
	}

	ptr := gen.dereferencePtr(variable)
	store := newIntInstr(IntMStore, "")

	cc(ptr, store)

	return ptr, nil
}

func (gen *IntGen) getPtr(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)

	if variable == nil {
		return gen.Errorf(tree, "undefined array: %v", name)
	}

	// Get stack ptr
	ptr := gen.loadStackPtr()
	// Push the variable offset
	push, cons := pushConstant(strconv.Itoa(variable.Offset()))
	cc(ptr, push, cons)

	return ptr, nil

	return gen.getAddress(variable), nil
}

func (gen *IntGen) getArray(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)

	if variable == nil {
		return gen.Errorf(tree, "undefined array: %v", name)
	}

	// TODO optimize if the expression in offset. If regular const (i.e. 0-9)
	// do an inline calculation instead.

	// Dereference pointer
	ptr := gen.getAddress(variable)

	if IsPtr(variable) {
		ptr = gen.dereferencePtr(variable)
	}

	// Get the offset (= expression between brackets [expression])
	offset := gen.MakeIntCode(tree.Children[0])
	// Get the size of the variable in bytes
	size, sizeConst := pushConstant("32")
	// Multiply offset with size
	mul := newIntInstr(IntMul, "")
	// Add the result to the memory location
	iAdd := newIntInstr(IntAdd, "")
	// b = a[0] // loc(a) + sizeOf(type(a)) * len(a)
	load := newIntInstr(IntMLoad, "")

	cc(ptr, offset, size, sizeConst, mul, iAdd, load)

	return ptr, nil
}

func (gen *IntGen) setArray(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)
	if variable == nil {
		return gen.Errorf(tree, "undefined array: %v", name)
	}

	val := gen.MakeIntCode(tree.Children[1])

	ptr := gen.getAddress(variable)
	// Get the offset (= expression between brackets [expression])
	expr := gen.MakeIntCode(tree.Children[0])
	// The value which we want to assign
	size, sizeConst := pushConstant("32")
	// Multiply offset with size
	mul := newIntInstr(IntMul, "")
	// Add the result to the memory location (ptr + offset + expr)
	iAdd := newIntInstr(IntAdd, "")
	store := newIntInstr(IntMStore, "")

	cc(val, ptr, expr, size, sizeConst, mul, iAdd, store)

	return val, nil
}

func (gen *IntGen) initNewArray(tree *SyntaxTree) (*IntInstr, error) {
	name := tree.Constant
	variable := gen.CurrentScope().GetVar(name)
	variable, err := gen.CurrentScope().NewVar(name, varArrTy)
	if err != nil {
		return gen.Errorf(tree, "Redeclaration of variable '%s'", name)
	}

	length, _ := strconv.Atoi(tree.Size)
	variable.SetSize(32 * length)

	return newIntInstr(IntIgnore, ""), nil
}

func (gen *IntGen) err(err error) *IntInstr {
	gen.addError(err)

	return newIntInstr(IntIgnore, "")
}

func (gen *IntGen) setVariable(tree *SyntaxTree, identifier *SyntaxTree) *IntInstr {
	variable := gen.GetVar(identifier.Constant)

	var lhs, rhs *IntInstr
	var err error

	switch identifier.Type {
	case DerefPtrTy:
		lhs, err = gen.setPtr(identifier)
		if err != nil {
			return gen.err(err)
		}
	default:
		lhs = gen.MakeIntCode(identifier)
	}

	switch tree.Type {
	case StringTy:
		if len(tree.Constant) > 32 {
			c, err := gen.Errorf(tree, "attempting to store string greater than 32 bytes")
			gen.addError(err)

			return c
		}

		var t Instr
		if identifier.Type == SetStoreTy {
			rhs = gen.makePush("0x" + hex.EncodeToString([]byte(tree.Constant)))
			break
		} else {
			t = IntMStore
			variable.SetType(varStrTy)
		}

		var length int
		rhs, length = gen.stringToInstr(variable, []byte(tree.Constant), t)

		if identifier.Type != SetStoreTy {
			variable.SetSize(int(math.Max(math.Max(float64(variable.Size()), float64(length)), 32.0)))
		}
	case ConstantTy:
		if identifier.Type != SetStoreTy {
			variable.SetType(varNumTy)
		}

		rhs = gen.MakeIntCode(tree)
	case DerefPtrTy:
		rhs = cc(gen.MakeIntCode(tree), newIntInstr(IntMLoad, ""))
	default:
		rhs = gen.MakeIntCode(tree)
	}

	concat(rhs, lhs)

	return rhs
}

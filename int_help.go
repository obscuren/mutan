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
	mload := NewIntInstr(intMLoad, "")
	concat(push, mload)

	return push
}

func (self *IntGen) setStackPtr(ptr int) *IntInstr {
	self.currentStackSize = ptr

	push := self.makePush(strconv.Itoa(ptr))
	push2 := self.makePush("0")
	mstore := NewIntInstr(intMStore, "")
	concat(push, push2)
	concat(push2, mstore)

	return push
}

func (self *IntGen) addStackPtr(size int) *IntInstr {
	push := self.makePush(strconv.Itoa(size))
	load := self.loadStackPtr()
	add := NewIntInstr(intAdd, "")
	loc := self.makePush("0")
	store := NewIntInstr(intMStore, "")

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

	l := gen.locals[t.Constant]
	if l == nil {
		return t.Errorf("undefine variable: '%s'", t.Constant)
	}

	pushOff, offCons := pushConstant(strconv.Itoa(l.size))
	pushLoc, locCons := pushConstant(strconv.Itoa(l.pos))
	locCons.ConstRef = t.Constant

	concat(pushOff, offCons)
	concat(offCons, pushLoc)
	concat(pushLoc, locCons)

	return pushOff, nil
}

func (gen *IntGen) makePush(num string, v ...int) *IntInstr {
	var push, cons *IntInstr
	if len(v) > 0 {
		n, _ := strconv.Atoi(num)
		push = NewIntInstr(intPush4, "")
		cons = NewIntInstr(intConst, "0x"+hex.EncodeToString(numberToBytes(int32(n), v[0])))
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
	ignore := NewIntInstr(intIgnore, "")
	var lastPush *IntInstr
	i := 0
	l := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for ; i < len(b); i += 32 {
		offset := int(math.Min(float64(i+32), float64(len(b))))
		hex := hex.EncodeToString(append(b[i:offset], l[0:32-len(b[i:offset])]...))
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

func (gen *IntGen) stringToInstr(id *Variable, b []byte, t Instr) (*IntInstr, int) {
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

package mutan

import (
	"fmt"
	"io"
	"io/ioutil"
)

type Compiler struct {
	intInsrs *IntInstr

	asm []interface{}
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) add(v ...interface{}) {
	c.asm = append(c.asm, v...)
}

func (c *Compiler) Compile(instr *IntInstr) ([]interface{}, error) {
	c.asm = nil

	// The following code is very explicit. I could have used string.Upcase
	for instr != nil {
		switch instr.Code {
		case intPush:
			c.add("PUSH")
		case intPush20:
			c.add("PUSH20")
		case intConst:
			c.add(instr.Constant)
		case intEqual:
			c.add("EQ")
		case intGt:
			c.add("GT")
		case intLt:
			c.add("LT")
		case intMul:
			c.add("MUL")
		case intSub:
			c.add("SUB")
		case intDiv:
			c.add("DIV")
		case intAdd:
			c.add("ADD")
		case intAssign:
		case intEmpty:
		case intJump:
			fmt.Println("Jump with target", instr.Target.n)
		case intMStore:
			c.add("MSTORE")
		case intMLoad:
			c.add("MLOAD")

		case intNot:
			c.add("NOT")
		case intJumpi:
			c.add("PUSH")
			c.add(instr.Target.n + 1)
			c.add("JUMPI")
		case intSStore:
			c.add("SSTORE")
		case intSLoad:
			c.add("SLOAD")
		case intStop:
			c.add("STOP")
		case intOrigin:
			c.add("ORIGIN")
		case intCaller:
			c.add("CALLER")
		case intCallVal:
			c.add("CALLVALUE")
		case intCallDataLoad:
			c.add("CALLDATALOAD")
		case intCallDataSize:
			c.add("CALLDATASIZE")
		case intGasPrice:
			c.add("GASPRICE")
		case intCall:
			c.add("CALL")
		case intASM:
			c.add(instr.Constant)
		case intTarget:
			// XXX Ignore this, it's not really an actual opcode. It just helps in figuring out where to
			// jump to if a jump is set.
		}

		instr = instr.Next
	}

	return c.asm, nil
}

func Compile(source io.Reader, debug bool) (asm []interface{}, err error) {
	var buff []byte
	// Read all at once
	buff, err = ioutil.ReadAll(source)
	if err != nil {
		return
	}

	var ast *SyntaxTree
	ast, err = MakeAst(string(buff))
	if err != nil {
		return
	}

	if debug {
		fmt.Println(ast)
	}

	gen := NewGen()
	intCode := gen.MakeIntCode(ast)
	if len(gen.errors) > 0 {
		for _, genErr := range gen.errors {
			fmt.Println(genErr)
		}
		return nil, fmt.Errorf("Exited with errors\n")
	}
	intCode.setNumbers(0)

	if debug {
		fmt.Println(intCode)
	}

	comp := NewCompiler()
	asm, err = comp.Compile(intCode)

	return
}

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
		case intDup:
			c.add("DUP")
		case intPc:
			c.add("PC")
		case intSwp:
			c.add("SWAP")
		case intPush1:
			c.add("PUSH1")
		case intPush2:
			c.add("PUSH2")
		case intPush3:
			c.add("PUSH3")
		case intPush4:
			c.add("PUSH4")
		case intPush5:
			c.add("PUSH5")
		case intPush6:
			c.add("PUSH6")
		case intPush7:
			c.add("PUSH7")
		case intPush8:
			c.add("PUSH8")
		case intPush9:
			c.add("PUSH9")
		case intPush10:
			c.add("PUSH10")
		case intPush11:
			c.add("PUSH11")
		case intPush12:
			c.add("PUSH12")
		case intPush13:
			c.add("PUSH13")
		case intPush14:
			c.add("PUSH14")
		case intPush15:
			c.add("PUSH15")
		case intPush16:
			c.add("PUSH16")
		case intPush17:
			c.add("PUSH17")
		case intPush18:
			c.add("PUSH18")
		case intPush19:
			c.add("PUSH19")
		case intPush20:
			c.add("PUSH20")
		case intPush21:
			c.add("PUSH21")
		case intPush22:
			c.add("PUSH22")
		case intPush23:
			c.add("PUSH23")
		case intPush24:
			c.add("PUSH24")
		case intPush25:
			c.add("PUSH25")
		case intPush26:
			c.add("PUSH26")
		case intPush27:
			c.add("PUSH27")
		case intPush28:
			c.add("PUSH28")
		case intPush29:
			c.add("PUSH29")
		case intPush30:
			c.add("PUSH30")
		case intPush31:
			c.add("PUSH31")
		case intPush32:
			c.add("PUSH32")
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
		case intExp:
			c.add("EXP")
		case intMod:
			c.add("MOD")
		case intXor:
			c.add("XOR")
		case intOr:
			c.add("OR")
		case intAnd:
			c.add("AND")
		case intAdd:
			c.add("ADD")
		case intAssign:
		case intEmpty:
		case intMStore:
			c.add("MSTORE")
		case intMLoad:
			c.add("MLOAD")
		case intNot:
			c.add("NOT")
		case intJumpi:
			c.add("JUMPI")
		case intJump:
			c.add("JUMP")
		case intSStore:
			c.add("SSTORE")
		case intSLoad:
			c.add("SLOAD")
		case intStop:
			c.add("STOP")
		case intOrigin:
			c.add("ORIGIN")
		case intAddress:
			c.add("ADDRESS")
		case intBalance:
			c.add("BALANCE")
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
		case intDiff:
			c.add("DIFFICULTY")
		case intPrevHash:
			c.add("PREVHASH")
		case intTimestamp:
			c.add("TIMESTAMP")
		case intCoinbase:
			c.add("COINBASE")
		case intGas:
			c.add("GAS")
		case intBlockNum:
			c.add("NUMBER")
		case intCall:
			c.add("CALL")
		case intCreate:
			c.add("CREATE")
		case intReturn:
			c.add("RETURN")
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

func CompileStage(source io.Reader, debug bool) (asm []interface{}, errors []error) {
	var buff []byte
	// Read all at once
	buff, err := ioutil.ReadAll(source)
	if err != nil {
		errors = append(errors, err)
		return
	}

	s, _ := PreProcess(string(buff))

	var ast *SyntaxTree
	ast, err = MakeAst(s)
	if err != nil {
		errors = append(errors, err)
		return
	}

	if debug {
		fmt.Println(ast)
	}

	gen := NewGen()
	gen.NewVar("___stackPtr", varNumTy)
	ptr := gen.setStackPtr(0)

	intCode := concat(ptr, gen.MakeIntCode(ast))
	if len(gen.errors) > 0 {
		for _, genErr := range gen.errors {
			fmt.Println(genErr)
		}
		return nil, gen.Errors()
	}
	intCode.setNumbers(0, gen)

	if debug {
		fmt.Println(intCode)
	}

	comp := NewCompiler()
	asm, err = comp.Compile(intCode)

	return
}

func Compile(source io.Reader, debug bool) (byteCode []byte, errors []error) {
	asm, err := CompileStage(source, debug)
	if err != nil {
		return nil, err
	}

	bytes := Assemble(asm...)

	return bytes, nil
}

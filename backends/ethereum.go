package backend

import (
	"github.com/ethereum/go-ethereum/core/vm"
	. "github.com/obscuren/mutan/front"
)

type EthereumBackend struct {
	asm []interface{}
}

func NewEthereumBackend() *EthereumBackend {
	return &EthereumBackend{}
}

func (c *EthereumBackend) add(v ...interface{}) {
	c.asm = append(c.asm, v...)
}

func (c *EthereumBackend) Compile(instr *IntInstr) ([]interface{}, error) {
	c.asm = nil

	// The following code is very explicit. I could have used string.Upcase
	for instr != nil {
		switch instr.Code {
		case IntPc:
			c.add(vm.PC)
			//c.add("pc")
		case IntMulMod:
			c.add(vm.MULMOD)
			//c.add("mulmod")
		case IntPop:
			c.add(vm.POP)
			//c.add("pop")
		case IntAddMod:
			c.add(vm.ADDMOD)
			//c.add("addmod")
		case IntSwp1:
			c.add(vm.SWAP1)
			//c.add("swap1")
		case IntSwp2:
			c.add(vm.SWAP2)
			//c.add("swap2")
		case IntSwp3:
			c.add(vm.SWAP3)
			//c.add("swap3")
		case IntSwp4:
			c.add(vm.SWAP4)
			//c.add("swap4")
		case IntSwp5:
			c.add(vm.SWAP5)
			//c.add("swap5")
		case IntSwp6:
			c.add(vm.SWAP6)
			//c.add("swap6")
		case IntSwp7:
			c.add(vm.SWAP7)
			//c.add("swap7")
		case IntSwp8:
			c.add(vm.SWAP8)
			//c.add("swap8")
		case IntSwp9:
			c.add(vm.SWAP9)
			//c.add("swap9")
		case IntSwp10:
			c.add(vm.SWAP10)
			//c.add("swap10")
		case IntSwp11:
			c.add(vm.SWAP11)
			//c.add("swap11")
		case IntSwp12:
			c.add(vm.SWAP12)
			//c.add("swap12")
		case IntSwp13:
			c.add(vm.SWAP13)
			//c.add("swap13")
		case IntSwp14:
			c.add(vm.SWAP14)
			//c.add("swap14")
		case IntSwp15:
			c.add(vm.SWAP15)
			//c.add("swap15")
		case IntSwp16:
			c.add(vm.SWAP16)
			//c.add("dup16")
		case IntDup1:
			c.add(vm.DUP1)
			//c.add("dup1")
		case IntDup2:
			c.add(vm.DUP2)
			//c.add("dup2")
		case IntDup3:
			c.add(vm.DUP3)
			//c.add("dup3")
		case IntDup4:
			c.add(vm.DUP4)
			//c.add("dup4")
		case IntDup5:
			c.add(vm.DUP5)
			//c.add("dup5")
		case IntDup6:
			c.add(vm.DUP6)
			//c.add("dup6")
		case IntDup7:
			c.add(vm.DUP7)
			//c.add("dup7")
		case IntDup8:
			c.add(vm.DUP8)
			//c.add("dup8")
		case IntDup9:
			c.add(vm.DUP9)
			//c.add("dup9")
		case IntDup10:
			c.add(vm.DUP10)
			//c.add("dup10")
		case IntDup11:
			c.add(vm.DUP11)
			//c.add("dup11")
		case IntDup12:
			c.add(vm.DUP12)
			//c.add("dup12")
		case IntDup13:
			c.add(vm.DUP13)
			//c.add("dup13")
		case IntDup14:
			c.add(vm.DUP14)
			//c.add("dup14")
		case IntDup15:
			c.add(vm.DUP15)
			//c.add("dup15")
		case IntDup16:
			c.add(vm.DUP16)
			//c.add("dup16")
		case IntPush1:
			c.add(vm.PUSH1)
			//c.add("push1")
		case IntPush2:
			c.add(vm.PUSH2)
			//c.add("push2")
		case IntPush3:
			c.add(vm.PUSH3)
			//c.add("push3")
		case IntPush4:
			c.add(vm.PUSH4)
			//c.add("push4")
		case IntPush5:
			c.add(vm.PUSH5)
			//c.add("push5")
		case IntPush6:
			c.add(vm.PUSH6)
			//c.add("push6")
		case IntPush7:
			c.add(vm.PUSH7)
			//c.add("push7")
		case IntPush8:
			c.add(vm.PUSH8)
			//c.add("push8")
		case IntPush9:
			c.add(vm.PUSH9)
			//c.add("push9")
		case IntPush10:
			c.add(vm.PUSH10)
			//c.add("push10")
		case IntPush11:
			c.add(vm.PUSH11)
			//c.add("push11")
		case IntPush12:
			c.add(vm.PUSH12)
			//c.add("push12")
		case IntPush13:
			c.add(vm.PUSH13)
			//c.add("push13")
		case IntPush14:
			c.add(vm.PUSH14)
			//c.add("push14")
		case IntPush15:
			c.add(vm.PUSH15)
			//c.add("push15")
		case IntPush16:
			c.add(vm.PUSH16)
			//c.add("push16")
		case IntPush17:
			c.add(vm.PUSH17)
			//c.add("push17")
		case IntPush18:
			c.add(vm.PUSH18)
			//c.add("push18")
		case IntPush19:
			c.add(vm.PUSH19)
			//c.add("push19")
		case IntPush20:
			c.add(vm.PUSH20)
			//c.add("push20")
		case IntPush21:
			c.add(vm.PUSH21)
			//c.add("push21")
		case IntPush22:
			c.add(vm.PUSH22)
			//c.add("push22")
		case IntPush23:
			c.add(vm.PUSH23)
			//c.add("push23")
		case IntPush24:
			c.add(vm.PUSH24)
			//c.add("push24")
		case IntPush25:
			c.add(vm.PUSH25)
			//c.add("push25")
		case IntPush26:
			c.add(vm.PUSH26)
			//c.add("push26")
		case IntPush27:
			c.add(vm.PUSH27)
			//c.add("push27")
		case IntPush28:
			c.add(vm.PUSH28)
			//c.add("push28")
		case IntPush29:
			c.add(vm.PUSH29)
			//c.add("push29")
		case IntPush30:
			c.add(vm.PUSH30)
			//c.add("push30")
		case IntPush31:
			c.add(vm.PUSH31)
			//c.add("push31")
		case IntPush32:
			c.add(vm.PUSH32)
			//c.add("push32")
		case IntConst:
			c.add(instr.Constant)
		case IntEqual:
			c.add(vm.EQ)
			//c.add("eq")
		case IntGt:
			c.add(vm.GT)
			//c.add("gt")
		case IntLt:
			c.add(vm.LT)
			//c.add("lt")
		case IntMul:
			c.add(vm.MUL)
			//c.add("mul")
		case IntSub:
			c.add(vm.SUB)
			//c.add("sub")
		case IntDiv:
			c.add(vm.DIV)
			//c.add("div")
		case IntExp:
			c.add(vm.EXP)
			//c.add("exp")
		case IntMod:
			c.add(vm.MOD)
			//c.add("mod")
		case IntXor:
			c.add(vm.XOR)
			//c.add("xor")
		case IntOr:
			c.add(vm.OR)
			//c.add("or")
		case IntAnd:
			c.add(vm.AND)
			//c.add("and")
		case IntAdd:
			c.add(vm.ADD)
			//c.add("add")
		case IntSuicide:
			c.add(vm.SUICIDE)
			//c.add("suicide")
		case IntAssign:
		case IntEmpty:
		case IntMStore:
			c.add(vm.MSTORE)
			//c.add("mstore")
		case IntMLoad:
			c.add(vm.MLOAD)
			//c.add("mload")
		case IntMSize:
			c.add(vm.MSIZE)
			//c.add("msize")
		case IntNot:
			c.add(vm.NOT)
			//c.add("not")
		case IntJumpi:
			c.add(vm.JUMPI)
			//c.add("jumpi")
		case IntJump:
			c.add(vm.JUMP)
			//c.add("jump")
		case IntSStore:
			c.add(vm.SSTORE)
			//c.add("sstore")
		case IntSLoad:
			c.add(vm.SLOAD)
			//c.add("sload")
		case IntStop:
			c.add(vm.STOP)
			//c.add("stop")
		case IntOrigin:
			c.add(vm.ORIGIN)
			//c.add("origin")
		case IntAddress:
			c.add(vm.ADDRESS)
			//c.add("address")
		case IntCodeCopy:
			c.add(vm.CODECOPY)
			//c.add("codecopy")
		case IntInlineCode:
			c.add(instr.Constant)
		case IntBalance:
			c.add(vm.BALANCE)
			//c.add("balance")
		case IntCaller:
			c.add(vm.CALLER)
			//c.add("caller")
		case IntCallVal:
			c.add(vm.CALLVALUE)
			//c.add("callvalue")
		case IntCallDataLoad:
			c.add(vm.CALLDATALOAD)
			//c.add("calldataload")
		case IntCallDataSize:
			c.add(vm.CALLDATASIZE)
			//c.add("calldatasize")
		case IntGasPrice:
			c.add(vm.GASPRICE)
			//c.add("gasprice")
		case IntDiff:
			c.add(vm.DIFFICULTY)
			//c.add("difficulty")
		case IntPrevHash:
			c.add(vm.BLOCKHASH)
			//c.add("prevhash")
		case IntTimestamp:
			c.add(vm.TIMESTAMP)
			//c.add("timestamp")
		case IntCoinbase:
			c.add(vm.COINBASE)
			//c.add("coinbase")
		case IntGas:
			c.add(vm.GAS)
			//c.add("gas")
		case IntBlockNum:
			c.add(vm.NUMBER)
			//c.add("number")
		case IntCall:
			c.add(vm.CALL)
			//c.add("call")
		case IntCreate:
			c.add(vm.CREATE)
			//c.add("create")
		case IntReturn:
			c.add(vm.RETURN)
			//c.add("return")
		case IntByte:
			c.add(vm.BYTE)
			//c.add("byte")
		case IntSha3:
			c.add(vm.SHA3)
			//c.add("sha3")
		case IntASM:
			c.add(instr.Constant)
		case IntTarget:
			c.add(vm.JUMPDEST)
			//c.add("dest")
		}

		instr = instr.Next
	}

	return c.asm, nil
}

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
		case IntMulMod:
			c.add(vm.MULMOD)
		case IntPop:
			c.add(vm.POP)
		case IntAddMod:
			c.add(vm.ADDMOD)
		case IntSwp1:
			c.add(vm.SWAP1)
		case IntSwp2:
			c.add(vm.SWAP2)
		case IntSwp3:
			c.add(vm.SWAP3)
		case IntSwp4:
			c.add(vm.SWAP4)
		case IntSwp5:
			c.add(vm.SWAP5)
		case IntSwp6:
			c.add(vm.SWAP6)
		case IntSwp7:
			c.add(vm.SWAP7)
		case IntSwp8:
			c.add(vm.SWAP8)
		case IntSwp9:
			c.add(vm.SWAP9)
		case IntSwp10:
			c.add(vm.SWAP10)
		case IntSwp11:
			c.add(vm.SWAP11)
		case IntSwp12:
			c.add(vm.SWAP12)
		case IntSwp13:
			c.add(vm.SWAP13)
		case IntSwp14:
			c.add(vm.SWAP14)
		case IntSwp15:
			c.add(vm.SWAP15)
		case IntSwp16:
			c.add(vm.SWAP16)
		case IntDup1:
			c.add(vm.DUP1)
		case IntDup2:
			c.add(vm.DUP2)
		case IntDup3:
			c.add(vm.DUP3)
		case IntDup4:
			c.add(vm.DUP4)
		case IntDup5:
			c.add(vm.DUP5)
		case IntDup6:
			c.add(vm.DUP6)
		case IntDup7:
			c.add(vm.DUP7)
		case IntDup8:
			c.add(vm.DUP8)
		case IntDup9:
			c.add(vm.DUP9)
		case IntDup10:
			c.add(vm.DUP10)
		case IntDup11:
			c.add(vm.DUP11)
		case IntDup12:
			c.add(vm.DUP12)
		case IntDup13:
			c.add(vm.DUP13)
		case IntDup14:
			c.add(vm.DUP14)
		case IntDup15:
			c.add(vm.DUP15)
		case IntDup16:
			c.add(vm.DUP16)
		case IntPush1:
			c.add(vm.PUSH1)
		case IntPush2:
			c.add(vm.PUSH2)
		case IntPush3:
			c.add(vm.PUSH3)
		case IntPush4:
			c.add(vm.PUSH4)
		case IntPush5:
			c.add(vm.PUSH5)
		case IntPush6:
			c.add(vm.PUSH6)
		case IntPush7:
			c.add(vm.PUSH7)
		case IntPush8:
			c.add(vm.PUSH8)
		case IntPush9:
			c.add(vm.PUSH9)
		case IntPush10:
			c.add(vm.PUSH10)
		case IntPush11:
			c.add(vm.PUSH11)
		case IntPush12:
			c.add(vm.PUSH12)
		case IntPush13:
			c.add(vm.PUSH13)
		case IntPush14:
			c.add(vm.PUSH14)
		case IntPush15:
			c.add(vm.PUSH15)
		case IntPush16:
			c.add(vm.PUSH16)
		case IntPush17:
			c.add(vm.PUSH17)
		case IntPush18:
			c.add(vm.PUSH18)
		case IntPush19:
			c.add(vm.PUSH19)
		case IntPush20:
			c.add(vm.PUSH20)
		case IntPush21:
			c.add(vm.PUSH21)
		case IntPush22:
			c.add(vm.PUSH22)
		case IntPush23:
			c.add(vm.PUSH23)
		case IntPush24:
			c.add(vm.PUSH24)
		case IntPush25:
			c.add(vm.PUSH25)
		case IntPush26:
			c.add(vm.PUSH26)
		case IntPush27:
			c.add(vm.PUSH27)
		case IntPush28:
			c.add(vm.PUSH28)
		case IntPush29:
			c.add(vm.PUSH29)
		case IntPush30:
			c.add(vm.PUSH30)
		case IntPush31:
			c.add(vm.PUSH31)
		case IntPush32:
			c.add(vm.PUSH32)
		case IntConst:
			c.add(instr.Constant)
		case IntEqual:
			c.add(vm.EQ)
		case IntGt:
			c.add(vm.GT)
		case IntLt:
			c.add(vm.LT)
		case IntMul:
			c.add(vm.MUL)
		case IntSub:
			c.add(vm.SUB)
		case IntDiv:
			c.add(vm.DIV)
		case IntExp:
			c.add(vm.EXP)
		case IntMod:
			c.add(vm.MOD)
		case IntXor:
			c.add(vm.XOR)
		case IntOr:
			c.add(vm.OR)
		case IntAnd:
			c.add(vm.AND)
		case IntAdd:
			c.add(vm.ADD)
		case IntSuicide:
			c.add(vm.SUICIDE)
		case IntAssign:
		case IntEmpty:
		case IntMStore:
			c.add(vm.MSTORE)
		case IntMLoad:
			c.add(vm.MLOAD)
		case IntMSize:
			c.add(vm.MSIZE)
		case IntNot:
			c.add(vm.NOT)
		case IntJumpi:
			c.add(vm.JUMPI)
		case IntJump:
			c.add(vm.JUMP)
		case IntSStore:
			c.add(vm.SSTORE)
		case IntSLoad:
			c.add(vm.SLOAD)
		case IntStop:
			c.add(vm.STOP)
		case IntOrigin:
			c.add(vm.ORIGIN)
		case IntAddress:
			c.add(vm.ADDRESS)
		case IntCodeCopy:
			c.add(vm.CODECOPY)
		case IntInlineCode:
			c.add(instr.Constant)
		case IntBalance:
			c.add(vm.BALANCE)
		case IntCaller:
			c.add(vm.CALLER)
		case IntCallVal:
			c.add(vm.CALLVALUE)
		case IntCallDataLoad:
			c.add(vm.CALLDATALOAD)
		case IntCallDataSize:
			c.add(vm.CALLDATASIZE)
		case IntGasPrice:
			c.add(vm.GASPRICE)
		case IntDiff:
			c.add(vm.DIFFICULTY)
		case IntPrevHash:
			c.add(vm.BLOCKHASH)
		case IntTimestamp:
			c.add(vm.TIMESTAMP)
		case IntCoinbase:
			c.add(vm.COINBASE)
		case IntGas:
			c.add(vm.GAS)
		case IntBlockNum:
			c.add(vm.NUMBER)
		case IntCall:
			c.add(vm.CALL)
		case IntCallCode:
			c.add(vm.CALLCODE)
		case IntCreate:
			c.add(vm.CREATE)
		case IntReturn:
			c.add(vm.RETURN)
		case IntByte:
			c.add(vm.BYTE)
		case IntSha3:
			c.add(vm.SHA3)
		case IntASM:
			c.add(instr.Constant)
		case IntTarget:
			c.add(vm.JUMPDEST)
		}

		instr = instr.Next
	}

	return c.asm, nil
}

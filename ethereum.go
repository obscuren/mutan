package mutan

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
		case intDup:
			c.add("dup")
		case intPc:
			c.add("pc")
		case intSwp:
			c.add("swap")
		case intPush1:
			c.add("push1")
		case intPush2:
			c.add("push2")
		case intPush3:
			c.add("push3")
		case intPush4:
			c.add("push4")
		case intPush5:
			c.add("push5")
		case intPush6:
			c.add("push6")
		case intPush7:
			c.add("push7")
		case intPush8:
			c.add("push8")
		case intPush9:
			c.add("push9")
		case intPush10:
			c.add("push10")
		case intPush11:
			c.add("push11")
		case intPush12:
			c.add("push12")
		case intPush13:
			c.add("push13")
		case intPush14:
			c.add("push14")
		case intPush15:
			c.add("push15")
		case intPush16:
			c.add("push16")
		case intPush17:
			c.add("push17")
		case intPush18:
			c.add("push18")
		case intPush19:
			c.add("push19")
		case intPush20:
			c.add("push20")
		case intPush21:
			c.add("push21")
		case intPush22:
			c.add("push22")
		case intPush23:
			c.add("push23")
		case intPush24:
			c.add("push24")
		case intPush25:
			c.add("push25")
		case intPush26:
			c.add("push26")
		case intPush27:
			c.add("push27")
		case intPush28:
			c.add("push28")
		case intPush29:
			c.add("push29")
		case intPush30:
			c.add("push30")
		case intPush31:
			c.add("push31")
		case intPush32:
			c.add("push32")
		case intConst:
			c.add(instr.Constant)
		case intEqual:
			c.add("eq")
		case intGt:
			c.add("gt")
		case intLt:
			c.add("lt")
		case intMul:
			c.add("mul")
		case intSub:
			c.add("sub")
		case intDiv:
			c.add("div")
		case intExp:
			c.add("exp")
		case intMod:
			c.add("mod")
		case intXor:
			c.add("xor")
		case intOr:
			c.add("or")
		case intAnd:
			c.add("and")
		case intAdd:
			c.add("add")
		case intAssign:
		case intEmpty:
		case intMStore:
			c.add("mstore")
		case intMLoad:
			c.add("mload")
		case intNot:
			c.add("not")
		case intJumpi:
			c.add("jumpi")
		case intJump:
			c.add("jump")
		case intSStore:
			c.add("sstore")
		case intSLoad:
			c.add("sload")
		case intStop:
			c.add("stop")
		case intOrigin:
			c.add("origin")
		case intAddress:
			c.add("address")
		case intBalance:
			c.add("balance")
		case intCaller:
			c.add("caller")
		case intCallVal:
			c.add("callvalue")
		case intCallDataLoad:
			c.add("calldataload")
		case intCallDataSize:
			c.add("calldatasize")
		case intGasPrice:
			c.add("gasprice")
		case intDiff:
			c.add("difficulty")
		case intPrevHash:
			c.add("prevhash")
		case intTimestamp:
			c.add("timestamp")
		case intCoinbase:
			c.add("coinbase")
		case intGas:
			c.add("gas")
		case intBlockNum:
			c.add("number")
		case intCall:
			c.add("call")
		case intCreate:
			c.add("create")
		case intReturn:
			c.add("return")
		case intASM:
			c.add(instr.Constant)
		case intTarget:
		}

		instr = instr.Next
	}

	return c.asm, nil
}

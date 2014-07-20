package frontend

type Instr byte

const (
	IntEqual Instr = iota
	IntDup
	IntPc
	IntSwp
	IntGt
	IntLt
	IntMul
	IntAdd
	IntDiv
	IntSub
	IntExp
	IntMod
	IntXor
	IntOr
	IntAnd
	IntAssign
	IntConst
	IntEmpty
	IntJump
	IntTarget
	IntPush1
	IntPush2
	IntPush3
	IntPush4
	IntPush5
	IntPush6
	IntPush7
	IntPush8
	IntPush9
	IntPush10
	IntPush11
	IntPush12
	IntPush13
	IntPush14
	IntPush15
	IntPush16
	IntPush17
	IntPush18
	IntPush19
	IntPush20
	IntPush21
	IntPush22
	IntPush23
	IntPush24
	IntPush25
	IntPush26
	IntPush27
	IntPush28
	IntPush29
	IntPush30
	IntPush31
	IntPush32

	IntMStore
	IntMLoad
	IntMSize
	IntNot
	IntJumpi
	IntSStore
	IntSLoad
	IntStop
	IntOrigin
	IntAddress
	IntCodeCopy
	IntCaller
	IntCallVal
	IntCallDataLoad
	IntCallDataSize
	IntGasPrice
	IntDiff
	IntPrevHash
	IntTimestamp
	IntCoinbase
	IntBalance
	IntGas
	IntBlockNum
	IntReturn
	IntSuicide

	IntFuncDef
	IntFuncCall

	// Asm is a special opcode. It's not malformed in anyway
	IntASM
	IntArray
	IntCall
	IntCreate
	IntSizeof
	IntByte

	IntIgnore
	IntInlineCode
)

var instrAsString = []string{
	"equal",
	"dup",
	"pc",
	"swap",
	"gt",
	"lt",
	"mul",
	"add",
	"div",
	"sub",
	"exp",
	"mod",
	"xor",
	"or",
	"and",
	"assign",
	"const",
	"empty",
	"jump",
	"target",
	"push1",
	"push2",
	"push3",
	"push4",
	"push5",
	"push6",
	"push7",
	"push8",
	"push9",
	"push10",
	"push11",
	"push12",
	"push13",
	"push14",
	"push15",
	"push16",
	"push17",
	"push18",
	"push19",
	"push20",
	"push21",
	"push22",
	"push23",
	"push24",
	"push25",
	"push26",
	"push27",
	"push28",
	"push29",
	"push30",
	"push31",
	"push32",
	"mstore",
	"mload",
	"msize",
	"not",
	"jmpi",
	"sstore",
	"sload",
	"stop",
	"origin",
	"address",
	"codecopy",
	"caller",
	"value",
	"dataload",
	"datasize",
	"gasprice",
	"diff",
	"prevhash",
	"timestamp",
	"coinbase",
	"balance",
	"gas",
	"blocknum",
	"return",
	"suicide",

	"func def",
	"func call",

	"asm",
	"array",
	"call",
	"create",
	"sizeof",
	"byte",

	"ignore",
	"inline",
}

func (op Instr) String() string {
	if len(instrAsString) < int(op) {
		return "Unknown"
	}

	return instrAsString[op]
}

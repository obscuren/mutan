package mutan

type Instr byte

const (
	intEqual Instr = iota
	intGt
	intLt
	intMul
	intAdd
	intDiv
	intSub
	intExp
	intMod
	intXor
	intOr
	intAnd
	intAssign
	intConst
	intEmpty
	intJump
	intTarget
	intPush1
	intPush2
	intPush3
	intPush4
	intPush5
	intPush6
	intPush7
	intPush8
	intPush9
	intPush10
	intPush11
	intPush12
	intPush13
	intPush14
	intPush15
	intPush16
	intPush17
	intPush18
	intPush19
	intPush20
	intPush21
	intPush22
	intPush23
	intPush24
	intPush25
	intPush26
	intPush27
	intPush28
	intPush29
	intPush30
	intPush31
	intPush32

	intMStore
	intMLoad
	intNot
	intJumpi
	intSStore
	intSLoad
	intStop
	intOrigin
	intAddress
	intCaller
	intCallVal
	intCallDataLoad
	intCallDataSize
	intGasPrice
	intDiff
	intPrevHash
	intTimestamp
	intCoinbase
	intBalance
	intGas
	intBlockNum
	intReturn

	// Asm is a special opcode. It's not malformed in anyway
	intASM
	intArray
	intCall
	intCreate
	intSizeof

	intIgnore
)

var instrAsString = []string{
	"equal",
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
	"not",
	"jmpi",
	"sstore",
	"sload",
	"stop",
	"origin",
	"address",
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

	"asm",
	"array",
	"call",
	"create",
	"sizeof",

	"ignore",
}

func (op Instr) String() string {
	if len(instrAsString) < int(op) {
		return "Unknown"
	}

	return instrAsString[op]
}

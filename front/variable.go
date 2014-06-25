package frontend

type varType byte

const (
	varUndefinedTy varType = iota
	varNumTy
	varStrTy
	varArrTy
	varPtrTy
	varFuncTy
)

var varTypeStringTable = []string{
	"Undefined",
	"Number",
	"String",
	"Array",
	"Ptr",
	"Func",
}

func (self varType) String() string {
	if len(varTypeStringTable) > int(self) {
		return varTypeStringTable[int(self)]
	}

	return "unknown var type"
}

type Var interface {
	Id() string

	Type() varType
	SetType(t varType)

	Size() int
	SetSize(int)

	Offset() int
	Instr() *IntInstr
	SetInstr(*IntInstr)
}

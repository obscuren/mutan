package mutan

type varType byte

const (
	varUndefinedTy varType = iota
	varNumTy
	varStrTy
	varArrTy
	varPtrTy
	varFuncTy
)

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

/*
type Variable struct {
	id      string
	typ     varType
	pos     int
	size    int
	varSize int

	instr *IntInstr
}
*/

package mutan

type varType byte

const (
	varUndefinedTy varType = iota
	varNumTy
	varStrTy
	varArrTy
)

type Variable struct {
	id      string
	typ     varType
	pos     int
	size    int
	varSize int

	instr *IntInstr
}

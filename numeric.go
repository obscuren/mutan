package mutan

type Variable struct {
	id     string
	size   int
	typ    varType
	offset int

	instr *IntInstr

	// Amount of elements an array has
	elements int
}

func NewNumeric(id string, offset int) *Variable {
	return &Variable{id: id, typ: varNumTy, offset: offset, size: 32}
}

func NewArray(id string, offset, count int) *Variable {
	return &Variable{id: id, typ: varArrTy, offset: offset, size: 32 * count, elements: count}
}

func (self *Variable) Id() string               { return self.id }
func (self *Variable) Type() varType            { return varNumTy }
func (self *Variable) SetType(t varType)        { self.typ = t }
func (self *Variable) Offset() int              { return self.offset }
func (self *Variable) Size() int                { return self.size }
func (self *Variable) SetSize(s int)            { self.size = s }
func (self *Variable) Instr() *IntInstr         { return self.instr }
func (self *Variable) SetInstr(instr *IntInstr) { self.instr = instr }

func IsNum(v Var) bool {
	return v.Type() == varNumTy
}

func IsArray(v Var) bool {
	return v.Type() == varArrTy
}

func VarLength(v Var) int {
	return v.Size() / 32
}

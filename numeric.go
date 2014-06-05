package mutan

type Numeric struct {
	id     string
	offset int
	size   int

	instr *IntInstr
}

func NewNumeric(id string, offset int) *Numeric {
	return &Numeric{id: id, offset: offset}
}

func (self *Numeric) Id() string               { return self.id }
func (self *Numeric) Type() varType            { return varNumTy }
func (self *Numeric) Offset() int              { return self.offset }
func (self *Numeric) Size() int                { return 32 }
func (self *Numeric) Instr() *IntInstr         { return self.instr }
func (self *Numeric) SetInstr(instr *IntInstr) { self.instr = instr }

func IsNum(v Var) bool {
	_, ok := v.(*Numeric)

	return ok
}

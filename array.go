package mutan

type Array struct {
	id     string
	offset int
	// Element count
	eCount int

	instr *IntInstr
}

func NewArray(id string, offset, count int) *Array {
	return &Array{id: id, offset: offset, eCount: count}
}

func (self *Array) Id() string               { return self.id }
func (self *Array) Type() varType            { return varArrTy }
func (self *Array) Offset() int              { return self.offset }
func (self *Array) Size() int                { return self.eCount * 32 }
func (self *Array) Instr() *IntInstr         { return self.instr }
func (self *Array) SetInstr(instr *IntInstr) { self.instr = instr }

func IsArr(v Var) bool {
	_, ok := v.(*Array)

	return ok
}

package frontend

type Pointer struct {
	id     string
	offset int
	ptrTy  varType

	instr *IntInstr
}

func NewPointer(id string, offset int, ptrTy varType) *Pointer {
	return &Pointer{id: id, offset: offset, ptrTy: ptrTy}
}

func (self *Pointer) Id() string               { return self.id }
func (self *Pointer) Type() varType            { return varArrTy }
func (self *Pointer) Offset() int              { return self.offset }
func (self *Pointer) Size() int                { return 32 }
func (self *Pointer) Instr() *IntInstr         { return self.instr }
func (self *Pointer) SetInstr(instr *IntInstr) { self.instr = instr }

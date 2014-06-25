package frontend

import (
	"testing"
)

func TestCc(t *testing.T) {
	instr1 := newIntInstr(intPush1, "")
	instr2 := newIntInstr(intPush2, "")
	instr3 := newIntInstr(intPush3, "")
	instr4 := newIntInstr(intPush4, "")
	instr5 := newIntInstr(intPush5, "")
	instr6 := newIntInstr(intPush6, "")

	cc(nil, nil, instr1, instr2, nil, instr3, instr4, instr5, instr6)
}

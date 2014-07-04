package frontend

import (
	"encoding/hex"
	"fmt"
)

type IntInstr struct {
	Code      Instr
	Constant  interface{}
	ConstRef  string
	Number    int
	Next      *IntInstr
	Target    *IntInstr
	TargetNum *IntInstr
	size      int
	n         int
	variable  Variable
	LineNo    int
}

func (instr *IntInstr) String() string {
	str := fmt.Sprintf("%-3d %-12v : %v\n", instr.n, instr.Code, instr.Constant)
	if instr.Next != nil {
		str += instr.Next.String()
	}

	return str
}

func newIntInstr(code Instr, constant string) *IntInstr {
	return &IntInstr{Code: code, Constant: constant}
}

func (instr *IntInstr) SetNumbers(i int, gen *IntGen) {
	/*
		var memLoc int
		for _, variable := range gen.locals {
			variable.pos = memLoc

			switch variable.typ {
			case varArrTy:
				for _, cons := range gen.arrayTable[variable.id] {
					cons.Constant = strconv.Itoa(memLoc)
				}
			case varStrTy:
				for _, instr := range gen.stringTable[variable.id] {
					num, _ := strconv.Atoi(instr.Constant.(string))
					instr.Constant = strconv.Itoa(num + memLoc)
					variable.pos = num + memLoc
				}
			default:
				if variable.instr != nil {
					variable.instr.Constant = strconv.Itoa(memLoc)
				}
			}

			memLoc += variable.size
		}
	*/

	num := instr
	for num != nil {
		num.n = i

		if len(num.ConstRef) > 0 {
			//num.Constant = strconv.Itoa(gen.locals[num.ConstRef].pos)
		}

		if num.Code != IntTarget && num.Code != IntIgnore {
			switch num.Code {
			case IntConst:
				if num.size == 0 {
					panic("NULL")
				}
				i += num.size
			case IntInlineCode:
				//gen.InlineCode[num.Number].OffsetInstr.Constant = "0x" + hex.EncodeToString(numberToBytes(int32(i), 32))
				gen.InlineCode[num.Number].OffsetInstr.Constant = string(numberToBytes(int32(i), 32))

				i += num.size
			default:
				i++
			}
		}

		num = num.Next
	}
}

func (self *IntInstr) LinkTargets() {
	instr := self
	for instr != nil {
		if instr.Code == IntJump || instr.Code == IntJumpi {
			// Set the target constant which we couldn't set before hand
			// when the instrbers weren't all set.
			if instr.TargetNum != nil {
				instr.TargetNum.Constant = string(numberToBytes(int32(instr.Target.n), 32))
				//instr.TargetNum.Constant = "0x" + hex.EncodeToString(numberToBytes(int32(instr.Target.n), 32))
			}
		}

		instr = instr.Next
	}
}

func (self *IntInstr) LinkCode(inlineCode []InlineCode) {
	// Append stop
	cc(self, newIntInstr(IntStop, ""))

	for i, inlineCode := range inlineCode {
		hex := hex.EncodeToString(inlineCode.Code)
		cons := newIntInstr(IntInlineCode, "0x"+hex)
		cons.size = len(hex) / 2
		cons.Number = i

		cc(self, cons)
	}
}

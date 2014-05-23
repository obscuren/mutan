package mutan

import (
	"fmt"
)

type IntInstr struct {
	Code      Instr
	Constant  interface{}
	Number    int
	Next      *IntInstr
	Target    *IntInstr
	TargetNum *IntInstr
	size      int
	n         int
	variable  Variable
}

func (instr *IntInstr) String() string {
	str := fmt.Sprintf("%-3d %-12v : %v\n", instr.n, instr.Code, instr.Constant)
	if instr.Next != nil {
		str += instr.Next.String()
	}

	return str
}

func NewIntInstr(code Instr, constant string) *IntInstr {
	return &IntInstr{Code: code, Constant: constant}
}

func (instr *IntInstr) setNumbers(i int) {
	num := instr
	for num != nil {
		num.n = i

		if num.Code != intTarget && num.Code != intIgnore {

			switch num.Code {
			case intConst:
				if num.size == 0 {
					fmt.Println("TIS", num.Constant)
					panic("NULL")
				}
				i += num.size
			default:
				i++
			}
		}

		num = num.Next
	}

	num = instr
	for num != nil {
		if num.Code == intJump || num.Code == intJumpi {
			// Set the target constant which we couldn't seet before hand
			// when the numbers weren't all set.
			num.TargetNum.Constant = string(numberToBytes(int32(num.Target.n), 32))
		}

		num = num.Next
	}
}

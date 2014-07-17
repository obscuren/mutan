package mutan

import (
	"github.com/obscuren/mutan/front"
)

func (self *Compiler) Optimize(intCode *frontend.IntInstr) *frontend.IntInstr {
	/*
		var (
			instr = intCode
			prev  = new(IntInstr)

			newCode *IntInstr
		)

		for instr != nil {
			if prev.Type == IntNot && instr.Type == IntNot {
				// TODO
			}

			instr = instr.Next
		}
	*/

	return intCode
}

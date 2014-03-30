package main

import (
	"fmt"
)

type Instr byte

const (
	intEqual Instr = iota
	intAssign
	intConst
	intMemSet
	intEmpty
	intIdentifier
)

var instrAsString = []string{
	"equal",
	"assign",
	"const",
	"memset",
	"empty",
	"identifier",
}

func (op Instr) String() string {
	if len(instrAsString) < int(op) {
		return "Unknown"
	}

	return instrAsString[op]
}

type IntInstr struct {
	Code     Instr
	Constant string
	Next     *IntInstr
}

func NewIntInstr(code Instr, constant string) *IntInstr {
	return &IntInstr{Code: code, Constant: constant}
}
func (instr *IntInstr) String() string {
	str := fmt.Sprintf("%-12v : %v\n", instr.Code, instr.Constant)
	if instr.Next != nil {
		str += instr.Next.String()
	}

	return str
}

// Concatenate two block of code together
func concat(blk1 *IntInstr, blk2 *IntInstr) *IntInstr {
	search := blk1
	for search.Next != nil {
		search = search.Next
	}

	search.Next = blk2

	return blk1
}

// Recursive intermediate code generator
func MakeIntCode(tree *SyntaxTree) *IntInstr {
	switch tree.Type {
	case StatementListTy:
		blk1 := MakeIntCode(tree.Children[0])
		blk2 := MakeIntCode(tree.Children[1])
		concat(blk1, blk2)

		return blk1
	case AssignmentTy:
		blk1 := MakeIntCode(tree.Children[0])
		blk2 := MakeIntCode(tree.Children[1])
		concat(blk1, blk2)

		return blk1
	case IdentifierTy:
		return NewIntInstr(intIdentifier, tree.Constant)
	case ConstantTy:
		return NewIntInstr(intConst, tree.Constant)
	case SetLocalTy:
		return NewIntInstr(intMemSet, tree.Constant)

	case EmptyTy:
		return NewIntInstr(intEmpty, "")
	}

	return nil
}

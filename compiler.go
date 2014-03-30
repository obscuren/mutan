package main

import (
	"fmt"
)

type varType byte

const (
	varNumTy varType = iota
	varStrTy
)

type Variable struct {
	typ varType
	val string
	pos int
}

type Compiler struct {
	intInsrs *IntInstr

	asm []interface{}

	locals map[string]*Variable

	memPos   int
	lastPush *IntInstr
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) add(v ...interface{}) {
	c.asm = append(c.asm, v...)
}

// Generates asm for getting a memory address
func (c *Compiler) getMemory(name string) error {
	if c.locals[name] == nil {
		return fmt.Errorf("undefined: %v", name)
	}

	c.add("PUSH", c.locals[name].pos, "MLOAD")

	return nil
}

// Generates asm for setting a memory address
func (c *Compiler) setMemory(instr *IntInstr) {
	// TODO Only accept numbers. Lnegth checking will have to
	// occur when I implement strings
	local := c.locals[instr.Constant]
	if local == nil {
		local = &Variable{typ: varNumTy, pos: c.memPos}
		c.locals[instr.Constant] = local
		c.memPos += 32
	}
	local.val = c.lastPush.Constant

	c.add("PUSH", local.pos, "MSTORE")
}

func (c *Compiler) Compile(intCode *IntInstr) ([]interface{}, error) {
	c.asm = nil
	c.locals = make(map[string]*Variable)

	instr := intCode
	for instr != nil {
		switch instr.Code {
		case intEqual:
		case intAssign:
		case intConst:
			c.add("PUSH")
			c.add(instr.Constant)
			c.lastPush = instr
		case intMemSet:
			c.setMemory(instr)
		case intEmpty:
		case intIdentifier:
			err := c.getMemory(instr.Constant)
			if err != nil {
				return nil, err
			}
		}

		instr = instr.Next
	}

	return c.asm, nil
}

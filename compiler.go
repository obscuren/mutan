package mutan

import (
	"fmt"
	"github.com/Logiraptor/chicken/peg"
	"io"
	"os"
	"strconv"
	"strings"
)

type Atom interface{}
type List []Atom
type Number int

const (
	Regular = iota
	Array
)

type Variable struct {
	Type int
	Loc  int
	// Used by arrays
	Size int
}

type Compiler struct {
	// Last memory position will figure out where to store the next
	// variable in memory
	lastMemPos int
	// The store holds locations for each named variable in memory
	store map[string]*Variable

	parser *peg.Language

	asm []interface{}
}

// Creates and returns a new compiler
func NewCompiler() *Compiler {
	parser, err := peg.NewParser(strings.NewReader(grammer))
	if err != nil {
		panic(err)
	}

	return &Compiler{0, make(map[string]*Variable), parser, []interface{}{}}
}

func (c *Compiler) Pos() int {
	return c.lastMemPos
}

func (c *Compiler) Write(str string, v ...interface{}) {
	if len(v) > 0 {
		c.asm = append(c.asm, []interface{}{str, v[0]}...)
	} else {
		c.asm = append(c.asm, str)
	}
}

func (c *Compiler) storeSet(p *peg.ParseTree) {
	c.compile(p.Children[1])
	c.Write("SSTORE")
}

// Compile a single instruction to ASM
func (c *Compiler) compile(p *peg.ParseTree) Atom {
	switch p.Type {
	case "assign_expression":
		// Get the right hand side assignment expression
		c.compile(p.Children[2])

		if p.Children[0].Type == "store" {
			c.storeSet(p.Children[0])
		} else {
			// Resole the variables memory location.
			// If non is present allocate new memory
			// otherwise overwrite the memory
			var loc int
			if c.store[string(p.Children[0].Data)] == nil {
				loc = c.lastMemPos
				c.store[string(p.Children[0].Data)] = &Variable{Regular, c.lastMemPos, 0}

				c.lastMemPos += 32
			} else {
				loc = c.store[string(p.Children[0].Data)].Loc
			}
			c.Write("PUSH", loc)
			c.Write("MSTORE")
		}
	case "number":
		c.Write("PUSH", string(p.Data))
		n, _ := strconv.Atoi(string(p.Data))

		return Number(n)
	case "name":
		c.Write("PUSH", c.store[string(p.Data)].Loc)
		c.Write("MLOAD")
	case "store":
		c.compile(p.Children[1])
		c.Write("SLOAD")
	case "call":
		//args := p.Children[2]
		if len(p.Children[2].Children) != 4 {
			panic("Invalid amount of arguments to call")
		}

		for i := 0; i < len(p.Children[2].Children); i++ {
			c.compile(p.Children[2].Children[i])
		}
		c.Write("CALL")
	case "array":
		n, _ := strconv.Atoi(string(p.Children[2].Data))
		c.Write("PUSH", c.store[string(p.Children[0].Data)].Loc*n)
		c.Write("MLOAD")
	case "create_array":
		n, _ := strconv.Atoi(string(p.Children[2].Data))

		c.store[string(p.Children[0].Data)] = &Variable{Array, c.lastMemPos, n}
		c.lastMemPos += n * 32
	case "arithmetic_expression":
		c.compile(p.Children[0])
		c.compile(p.Children[2])

		switch string(p.Children[1].Data) {
		case "+":
			c.Write("ADD")
		case "-":
			c.Write("SUB")
		case "/":
			c.Write("DIV")
		case "*":
			c.Write("MUL")
		}
	}

	return 0
}

// Compile program
func (c *Compiler) compileProgram(p *peg.ParseTree, out chan Atom) {
	for _, child := range p.Children {
		out <- c.compile(child)
	}
	close(out)
}

func (c *Compiler) CompileFile(fn string) []interface{} {
	source, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	return c.Compile(source)
}

func (c *Compiler) Compile(reader io.Reader) []interface{} {
	tree, err := c.parser.Parse(reader)
	if err != nil {
		panic(err)
	}

	lists := make(chan Atom)

	go c.compileProgram(tree, lists)
	for a := range lists {
		fmt.Sprintln(a)
	}

	return c.asm
}

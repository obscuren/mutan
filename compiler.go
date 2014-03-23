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

type Compiler struct {
	// Last memory position will figure out where to store the next
	// variable in memory
	lastMemPos int
	// The store holds locations for each named variable in memory
	store map[string]int

	parser *peg.Language

	asm []interface{}
}

// Creates and returns a new compiler
func NewCompiler() *Compiler {
	parser, err := peg.NewParser(strings.NewReader(grammer))
	if err != nil {
		panic(err)
	}

	return &Compiler{-32, make(map[string]int), parser, []interface{}{}}
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
	case "expr":
		fmt.Println("expr")
	case "assignment":
		// Get the right hand side assignment expression
		c.compile(p.Children[2])

		if p.Children[0].Type == "store" {
			c.storeSet(p.Children[0])
		} else {
			// Set memory position for current variable
			c.lastMemPos += 32
			c.Write("PUSH", c.lastMemPos)
			c.Write("MSTORE")
			c.store[string(p.Children[0].Data)] = c.lastMemPos
		}
	case "number":
		c.Write("PUSH", string(p.Data))
		n, _ := strconv.Atoi(string(p.Data))

		return Number(n)
	case "name":
		c.Write("PUSH", c.store[string(p.Data)])
		c.Write("MLOAD")
	case "store":
		c.compile(p.Children[1])
		c.Write("SLOAD")
	case "arithmetic":
		c.compile(p.Children[0])
		c.compile(p.Children[2])

		switch string(p.Children[1].Data) {
		case "+":
			c.Write("ADD")
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

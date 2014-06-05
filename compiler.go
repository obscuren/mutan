package mutan

import (
	"fmt"
	"io"
	"io/ioutil"
)

type CompilerBackend interface {
	Compile(instr *IntInstr) ([]interface{}, error)
}

type Compiler struct {
	intInsrs *IntInstr

	Backend CompilerBackend
}

func NewCompiler(backend CompilerBackend) *Compiler {
	return &Compiler{Backend: backend}
}

func (self *Compiler) Compile(instr *IntInstr) ([]interface{}, error) {
	return self.Backend.Compile(instr)
}

func CompileStage(source io.Reader, debug bool) (asm []interface{}, errors []error) {
	var buff []byte
	// Read all at once
	buff, err := ioutil.ReadAll(source)
	if err != nil {
		errors = append(errors, err)
		return
	}

	s, _ := PreProcess(string(buff))

	var ast *SyntaxTree
	ast, err = MakeAst(s)
	if err != nil {
		errors = append(errors, err)
		return
	}

	if debug {
		fmt.Println(ast)
	}

	gen := NewGen()
	gen.NewVar("___stackPtr", varNumTy)
	ptr := gen.setStackPtr(0)

	intCode := concat(ptr, gen.MakeIntCode(ast))
	if len(gen.errors) > 0 {
		for _, genErr := range gen.errors {
			fmt.Println(genErr)
		}
		return nil, gen.Errors()
	}
	intCode.setNumbers(0, gen)
	intCode.linkTargets()

	if debug {
		fmt.Println(intCode)
	}

	comp := NewCompiler(NewEthereumBackend())
	asm, err = comp.Compile(intCode)

	return
}

func Compile(source io.Reader, debug bool) (byteCode []byte, errors []error) {
	asm, err := CompileStage(source, debug)
	if err != nil {
		return nil, err
	}

	bytes := Assemble(asm...)

	return bytes, nil
}

package mutan

import (
	"fmt"
	"github.com/obscuren/mutan/front"
	"io"
	"io/ioutil"
)

type CompilerBackend interface {
	Compile(*frontend.IntInstr) ([]interface{}, error)
}

type Compiler struct {
	intInsrs *frontend.IntInstr

	Backend CompilerBackend

	Debug bool
}

func NewCompiler(backend CompilerBackend) *Compiler {
	c := &Compiler{Backend: backend}
	frontend.Compiler = c

	return c
}

func (self *Compiler) PreProcessorStage(str string) (string, error) {
	return frontend.PreProcess(str)
}

func (self *Compiler) CompileStage(code string) (asm []interface{}, errors []error) {
	ast, err := frontend.MakeAst(code)
	if err != nil {
		errors = append(errors, err)
		return
	}

	if self.Debug {
		fmt.Println(ast)
	}

	gen := frontend.NewGen()
	gen.NewVar("___stackPtr", 1)
	ptr := gen.SetStackPtr(0)

	intCode := frontend.Concat(ptr, gen.MakeIntCode(ast))
	if len(gen.Errors) > 0 {
		for _, genErr := range gen.Errors {
			fmt.Println(genErr)
		}
		return nil, gen.Errors
	}
	intCode.SetNumbers(0, gen)
	intCode.LinkTargets()

	if self.Debug {
		fmt.Println(intCode)
	}

	asm, err = self.Backend.Compile(intCode)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return
}

func (self *Compiler) AssemblerStage(asm ...interface{}) ([]byte, error) {
	return frontend.Assemble(asm...)
}

func (self *Compiler) ReadAll(source io.Reader) (string, error) {
	var buff []byte
	// Read all at once
	buff, err := ioutil.ReadAll(source)
	if err != nil {
		return "", err
	}

	return string(buff), nil
}

func (self *Compiler) Compile(source io.Reader) (bytecode []byte, errors []error) {
	code, err := self.ReadAll(source)
	if err != nil {
		errors = append(errors, err)
		return
	}

	code, err = self.PreProcessorStage(code)
	if err != nil {
		errors = append(errors, err)
		return
	}

	var asm []interface{}
	asm, errors = self.CompileStage(code)
	if errors != nil {
		return
	}

	bytecode, err = self.AssemblerStage(asm...)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return
}

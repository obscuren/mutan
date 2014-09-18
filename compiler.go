package mutan

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/obscuren/mutan/front"
)

const (
	Version = "0.5"
)

type CompilerBackend interface {
	Compile(*frontend.IntInstr) ([]interface{}, error)
}

type Compiler struct {
	intInsrs *frontend.IntInstr

	Backend CompilerBackend

	Debug bool

	Silent bool

	OptimizeCode bool
}

func NewCompiler(backend CompilerBackend) *Compiler {
	c := &Compiler{Backend: backend}
	frontend.Compiler = c

	return c
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

func (self *Compiler) PreProcessorStage(str string) (string, error) {
	return frontend.PreProcess(str)
}

func (self *Compiler) CompileStage(instrs *frontend.IntInstr) (asm []interface{}, err error) {
	asm, err = self.Backend.Compile(instrs)
	if err != nil {
		return
	}

	return
}

func (self *Compiler) AssemblerStage(asm ...interface{}) ([]byte, error) {
	return frontend.Assemble(asm...)
}

func (self *Compiler) IntermediateStage(code string) (intCode *frontend.IntInstr, errors []error) {
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

	intCode = frontend.Concat(ptr, gen.MakeIntCode(ast))
	if len(gen.Errors) > 0 && !self.Silent {
		for _, genErr := range gen.Errors {
			fmt.Println(genErr)
		}
		return nil, gen.Errors
	}

	if self.OptimizeCode {
		intCode = self.Optimize(intCode)
	}

	intCode.LinkCode(gen.InlineCode)
	intCode.SetNumbers(0, gen)
	intCode.LinkTargets()

	if self.Debug {
		fmt.Println(intCode)
	}

	return
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

	var intCode *frontend.IntInstr
	intCode, errors = self.IntermediateStage(code)
	if errors != nil {
		return
	}

	var asm []interface{}
	asm, err = self.CompileStage(intCode)
	if err != nil {
		errors = append(errors, err)
		return
	}

	bytecode, err = self.AssemblerStage(asm...)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return
}

func (self *Compiler) Intermediate(source io.Reader) (intCode *frontend.IntInstr, errors []error) {
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

	intCode, errors = self.IntermediateStage(code)
	if err != nil {
		return
	}

	return
}

func (self *Compiler) Assemble(source io.Reader) (asm []interface{}, errors []error) {
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

	var intCode *frontend.IntInstr
	intCode, errors = self.IntermediateStage(code)
	if errors != nil {
		return
	}

	asm, err = self.CompileStage(intCode)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return
}

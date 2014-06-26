package frontend

import (
	"io"
)

type CompilerFrontend interface {
	Compile(io.Reader) ([]byte, []error)
	Intermediate(io.Reader) (*IntInstr, []error)
}

var Compiler CompilerFrontend

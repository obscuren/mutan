package frontend

import (
	"io"
)

type CompilerFrontend interface {
	Compile(io.Reader) ([]byte, []error)
}

var Compiler CompilerFrontend

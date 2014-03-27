package mutan

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompiler(t *testing.T) {
	a := NewCompiler().Compile(strings.NewReader(`
	a = 10
	store[10] = a
	`))
	//[100] = a
	//c = [100]

	fmt.Println(a)
}

func TestArrayAllocation(t *testing.T) {
	compiler := NewCompiler()
	compiler.Compile(strings.NewReader("a = array(20)"))

	if compiler.Pos() != 20*32 {
		t.Error("Expected mem length to be", 21*32, "got", compiler.Pos())
	}
}

func TestArrayAccassing(t *testing.T) {
	compiler := NewCompiler()
	a := compiler.Compile(strings.NewReader(`
	a = array(20)
	a[1] = 10
	a[ 2] = 20
	b = a[1]
	a222 = 2
	`))

	fmt.Println(a)
}

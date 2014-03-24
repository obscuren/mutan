package mutan

import (
	_ "fmt"
	"strings"
	"testing"
)

func TestCompiler(t *testing.T) {
	NewCompiler().Compile(strings.NewReader(`
	a = 20
	c = 30
	b = 10
	d = a + b
	d = a - b
	d = a * b
	d = a / b
	call(1, 2, 3, 4)
	store[0]
	a[1]
	f = array(10)
	`))
	//[100] = a
	//c = [100]

	//fmt.Println(a)
}

func TestArrayAllocation(t *testing.T) {
	compiler := NewCompiler()
	compiler.Compile(strings.NewReader("a = array(20)"))

	if compiler.Pos() != 21*32 {
		t.Error("Expected mem length to be", 20*32, "got", compiler.Pos())
	}
}

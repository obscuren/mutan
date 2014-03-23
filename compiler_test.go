package mutan

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompiler(t *testing.T) {
	a := NewCompiler().Compile(strings.NewReader(`
	a = 20
	c = 30
	d = a + b
	`))
	//[100] = a
	//c = [100]

	fmt.Println(a)
}

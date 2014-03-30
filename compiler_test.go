package mutan

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompiler(t *testing.T) {
	/*
		ast := MakeAst(`a = 0
		if a == a {
			if b == b {
				a = 2
			}
			a = 3
		}
		`)
	*/
	_, err := Compile(strings.NewReader(`
        a = 100000
	b = a
	b = a
	`), true)

	if err != nil {
		fmt.Println(err)
	}
}

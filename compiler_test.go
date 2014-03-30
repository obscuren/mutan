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
	if a == b {
		b = 10000
		if b == a {
			c = 10
		}
	}

	store[0] = 10
	a = store[0]
	store[a] = 300
	`), true)

	if err != nil {
		fmt.Println(err)
	}
}

package mutan

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompiler(t *testing.T) {
	_, err := Compile(strings.NewReader(`
        a = 100000
	b = a
	if a == b {
		b = 10000
		if b == a {
			c = 10
		}

		asm(
			PUSH 0
			PUSH 1
			MLOAD
		)
	}

	store[0] = 10
	a = store[0]
	store[a] = 300
	e = store[a]

	asm(
		PUSH 1
		PUSH 2
		PUSH 3
	)

	`), true)

	if err != nil {
		fmt.Println(err)
	}
}

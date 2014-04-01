package mutan

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompiler(t *testing.T) {
	_, err := Compile(strings.NewReader(`
	big[2] c

	`), true)

	if err != nil {
		fmt.Println(err)
	}
}

/*
func TestCompiler(t *testing.T) {
	_, err := Compile(strings.NewReader(`
        int8 a
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

	if a > b {
		a = 1
	}
	if a < b {
		test = 100000
	}

	exit()

	a = this.caller()
	b = this.dataSize()
	c = this.value()
	d = this.dataLoad()
	e = this.gasPrice()

	arrayTest = array(20)
	`), true)

	if err != nil {
		fmt.Println(err)
	}
}
*/

package mutan

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompiler(t *testing.T) {
	asm, err := Compile(strings.NewReader(`
	int32 a
	int32 b

	if 200 < this.Value() {
		a = 20
		b = 30
	}
	`), true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(asm)
}

func TestArray(t *testing.T) {
	_, err := Compile(strings.NewReader(`
	int8[10] a
	int8[10] b
	addr address = "a46df28529eb8aa8b8c025b0b413c5f4b688352f"
	Call(address, 15, 10, a, b)
	// comment int8 b = 20
	`), true)

	if err != nil {
		fmt.Println(err)
	}
}

func TestString(t *testing.T) {
	ast, err := Compile(strings.NewReader(`
	      big t 
	      t = 10 ^ 20
	      store[100] = 10^20 - 10 * 20
	`), true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestPlusPlus(t *testing.T) {
	ast, err := Compile(strings.NewReader(`
	      big i = 0
	      i++
	`), true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestElse(t *testing.T) {
	ast, err := Compile(strings.NewReader(`
	      if 20 < 10 {
		      store[1001] = 10^20
	      } else {
		      store[1001] = 10^50
	      }
	`), true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestFor(t *testing.T) {
	ast, err := Compile(strings.NewReader(`
	      int32 b = 10
	      // regular for
	      for int32 i = 0; i < 10; i++ {
		      big t = 10
	      }
	`), true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestHex(t *testing.T) {
	ast, err := Compile(strings.NewReader(`
		addr address = 0xa46df28529eb8aa8b8c025b0b413c5f4b688352f
	`), true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestData(t *testing.T) {
	ast, err := Compile(strings.NewReader(`
		big t = this.Data(2)
		// OR
		asm (
			PUSH 64
			CALLDATALOAD
		)
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

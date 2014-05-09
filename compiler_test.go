package mutan

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompiler(t *testing.T) {
	asm, err := CompileStage(strings.NewReader(`
	int32 a
	int32 b

	if 200 < this.value() {
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
	_, err := CompileStage(strings.NewReader(`
	int8[10] a
	int8[10] b
	addr address = 0xa46df28529eb8aa8b8c025b0b413c5f4b688352f
	call(address, 0, 15, a, b)
	// comment int8 b = 20
	`), true)

	if err != nil {
		fmt.Println(err)
	}
}

func TestString(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	      big t 
	      t = 10 ^ 20
	      this.store[100] = 10^20 - 10 * 20
	`), true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestPlusPlus(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	      big i = 0
	      i++
	`), true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestElse(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	      if 20 < 10 {
		      this.store[1001] = 10^20
	      } else {
		      this.store[1001] = 10^50
	      }
	`), true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestFor(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
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
	ast, err := CompileStage(strings.NewReader(`
		addr address = 0xa46df28529eb8aa8b8c025b0b413c5f4b688352f
	`), true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestData(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		big t = this.data[2]
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

func TestReturn(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		return 1000
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestTransact(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		transact(0xaa1adef765cd, 100, nil)
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestCreate(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		big a = 0xdeadbeef
		big res = create(10000, a)
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestOpt(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		int8[2] test
		test[0] = 1000000000000000000000000000000000000000000
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestBalance(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		big a = this.balance()
		if a > 10 {
			big b = 10
		}
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestFunction(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		// import register from stdlib
		// // Future feature :-)
		// func main() {
		//     register("MyContractName")
		// }
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestOps(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		int8 test = 10 >= 20
		int8 itst2 = 10 <= 20
		int8 test3 = 10 != 20
		int test4 = 10 % 1 == 0

		test4 = test & test4
		test4 = test4 | test3
		test4 = test4 ^ test3
		test4 = test4 ** 2
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestBoolean(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		bool t = true
		bool f = false
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestNot(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		if !this.data[0] {
		}
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

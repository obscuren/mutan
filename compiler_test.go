package mutan

import (
	"fmt"
	"strings"
	"testing"
)

/*
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

func TestShift(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		int8 test = 8 << 2
		test = 256 >> 3
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestLambda(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	a := "hello world"
	a = "hello world hello world hello world hello world hello world hello world hello world hello world hello world"
	b := 20

	var[2] c
	c[0] = 10
	c[2] = 100

	d := "hello"

	return lambda {
		to := this.data[0]
		from := this.origin()
		value := this.data[1]

		if this.store[from] >= value {
			this.store[from] = this.store[from] - value
			this.store[to] = this.store[to] + value
		}
	}
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestString(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	a := "hello world"
	b := 10
	this.store[1] = "hello"

	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestReturn(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		return 1000
		return this.store[this.data[0]]
		var a = 10
		var b = 20
		return b
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestSome(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = this.address()
	var in = "jeff"
	var out
	var res = call(0xa4976648142a1e624c27dca4e9b1a6d8195f660c, 0, 0, in, out)
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}
*/

func TestCc(t *testing.T) {
	instr1 := newIntInstr(intPush1, "")
	instr2 := newIntInstr(intPush2, "")
	instr3 := newIntInstr(intPush3, "")
	instr4 := newIntInstr(intPush4, "")
	instr5 := newIntInstr(intPush5, "")
	instr6 := newIntInstr(intPush6, "")

	cc(nil, nil, instr1, instr2, nil, instr3, instr4, instr5, instr6)
}

func TestFuncDef(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	func two() {
		return 2
	}

	func one() {
		var b = two()

		return b
	}

	var a = one()
	var b = "hello"

	exit a

	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestFuncArgs(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	func fn(var a, var b) {
		return b
	}

	var a = fn(9, 5)
	exit a
	`), true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

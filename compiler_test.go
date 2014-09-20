package mutan

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/obscuren/mutan/backends"
)

func CompileStage(reader io.Reader, debug bool) (asm []interface{}, errors []error) {
	compiler := NewCompiler(backend.NewEthereumBackend())

	ic, errors := compiler.Intermediate(reader)
	if errors != nil {
		return
	}

	var err error
	asm, err = compiler.CompileStage(ic)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return
}

func TestCompiler(t *testing.T) {
	asm, err := CompileStage(strings.NewReader(`
	var a
	var b

	if 200 < tx.value() {
		a = 20
		b = 30
	}
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(asm)
	}
}

func TestArray(t *testing.T) {
	asm, err := CompileStage(strings.NewReader(`
	var[10] a
	var[10] b
	var address = 0xa46df28529eb8aa8b8c025b0b413c5f4b688352f
	call(address, 0, 15, a, b)
	// comment int8 b = 20
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(asm)
	}
}

func TestPlusPlus(t *testing.T) {
	asm, err := CompileStage(strings.NewReader(`
	      var i = 0
	      i++
	`), false)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(asm)
	}
}

func TestElse(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	      if 20 < 10 {
		      contract.storage[1001] = 10^20
	      } else {
		      contract.storage[1001] = 10^50
	      }
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestHex(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var address = 0xa46df28529eb8aa8b8c025b0b413c5f4b688352f
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestData(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var t = message.data[2]
		// OR
		asm {
			PUSH 64
			CALLDATALOAD
		}
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestCreate(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	a := 0xdeadbeef
	res := create(10000, a)
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestOpt(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var[2] test
		test[0] = 1000000000000000000000000000000000000000000
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestOps(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var test = 10 >= 20
		var itst2 = 10 <= 20
		var test3 = 10 != 20
		var test4 = 10 % 1 == 0

		test4 = test & test4
		test4 = test4 | test3
		test4 = test4 ^ test3
		test4 = test4 ** 2
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestBoolean(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var t = true
		var f = false
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestNot(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		if !message.data[0] {
		}
	`), false)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}

}

func TestShift(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var test = 8 << 2
		test = 256 >> 3
	`), false)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}

}

func TestReturn(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		exit 1000
		exit contract.storage[message.data[0]]
		var a = 10
		var b = 20
		exit b
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestSome(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = contract.address()
	var in = "jeff"
	var out
	var res = call(0xa4976648142a1e624c27dca4e9b1a6d8195f660c, 0, 0, in, out)
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestFuncDef(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	func two() {
		exit 2
	}

	func one() {
		var b = two()

		return b
	}

	var a = one()
	var b = "hello"

	exit a

	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestFuncArgs(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	func fn(var a, var b) {
		var[10] c
		c[2] = b
		return c[1]
	}

	var a = fn("test", 5)
	exit a
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestPointers(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = 10
	var *b = &a
	*b = 9
	var c = *b
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

/*
func TestAsm(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	func sha3(var *a, var s) var {
		m_push(a + s)
		m_push(a)
		asm {
			sha3      ; perform sha3
		}

		return m_pop() // Returns outcome of the SHA3
	}

	var a = 5
	var b = sha3(&a, 1)
	var *c = &a
	*c = *c * *c


	exit b
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}

}
*/

func TestWhile(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	i := 0
	for i < 10 {
		i++
	}
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}

}

func TestTransact(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		transact(0xaa1adef765cd, 1000, 100, nil)
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestStatementList(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = 10
	var b = (a * 2 + 2)
	`), true)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestArrayToPointer(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	func fn(var *a, var i) var {
		return a[i]
	}

	var[2] a
	a[0] = 1
	a[10] = 2

	var e = fn(&a, 1)
	`), true)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestImport(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	import "mutan/test.mu"
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}

}

func TestLocalScope(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = 10

	// Test the scope
	{
		var b = 20
	}
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}

}

func TestSyntaxErr(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	contract.call(
	`), false)

	if err == nil {
		t.Error("expected error")
	}

	fmt.Println(ast)
}

func TestBalance(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var a = balance(0xaabbccddeeff)
		if a > 10 {
			var b = 10
		}
	`), false)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}

}

func TestInlineCompile(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var addr = create(0, compile {
		return compile {
			return 10
		}
	})

	return compile {
		return 10
	}
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestByte(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = 100
	var b = 1
	if byte(a, b) != 1 {
		stop()
	}

	message.data

	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestSh3Single(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = 100
	var s = sha3(a, 32)
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestSuicide(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	suicide(tx.sender())
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestString(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var i = "test"
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestElseIf(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = 1
	if a == 0 {
		return 1
	} else if a == 1 {
		return 2
	} else if a == 2 {
		return 3
	} else {
		return 4
	}
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestAnd(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	return 0 && 1
	`), true)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestOr(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	return 1 || 0
	`), true)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestPrint(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	print("hello world")
	print(1)
	`), true)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestScope(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var x = 10

	{
		var x = 20
		return x
	}

	return x
	`), true)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

func TestFunc(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
return compile {
	to := message.data[0]
	from := message.caller()
	value := message.data[1]

	if contract.storage[from] >= value {
		contract.storage[from] = contract.storage[from] - value
		contract.storage[to] = contract.storage[to] + value
	}
}

	`), true)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}

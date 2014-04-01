mutan
=====

Mutan is in its design state. Features:

* Lexer
* Parser (=> AST)
* Int code generator
* Compiler

```go
asm, err := mutan.Compile(strings.NewReader(`
	int32 a = 10
	big b = 10
	if a == b {
		if b == b {
			a = 2
		}
		a = 3
	}

	big[10] arrays
	big b = this.caller()
`), true)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(asm)
}
```

Lexer is heavily inspired by Rob Pike's idea about lexers.

Mutan is a compiler & language for the [ethereum](http://ethereum.org) project.
Eventually it will move to the main [repo](https://github.com/ethereum).

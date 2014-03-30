mutan
=====

Mutan is in its design state. Features:

* Lexer
* Parser (=> AST)
* Int code generator
* Compiler

```golang
asm, err := mutan.Compile(strings.NewReader(`
	a = 10
	b = 10
	if a == b {
		if b == b {
			a = 2
		}
		a = 3
	}
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

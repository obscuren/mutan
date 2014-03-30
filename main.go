package main

import (
	"fmt"
)

func MakeAst(source string) *SyntaxTree {
	lex := lexer("test", source)
	yyParse(lex)

	return Tree
}

func main() {
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
	ast := MakeAst(`
        a = 100000
	b = a
	b = a
	`)
	fmt.Println(ast)
	intCode := MakeIntCode(ast)
	fmt.Println(intCode)

	comp := NewCompiler()
	asm, err := comp.Compile(intCode)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(asm)
	}
}

package main

import (
	"fmt"
)

func main() {
	lex := lexer("test", `a = 0
	if a == a {
		a = 2
	}
	`)
	yyParse(lex)
	fmt.Println(Tree)
}

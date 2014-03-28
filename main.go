package main

import (
	"fmt"
)

func main() {
	lex := lexer("test", "a = b")
	yyParse(lex)
	fmt.Println(Tree)
}

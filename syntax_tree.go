package mutan

import (
	"fmt"
)

type AstType byte

const (
	InvalidTy AstType = iota
	EmptyTy
	StatementListTy
	AssignmentTy
	IdentifierTy
	ConstantTy
	SetLocalTy
	BlockTy
	IfThenTy
	EqualTy
	StoreTy
	SetStoreTy
	InlineAsmTy
)

var astAsString = []string{
	"Invalid",
	"Empty",
	"Statement list",
	"Assignment",
	"Identifier",
	"Constant",
	"Set local",
	"Block",
	"If then",
	"Equal",
	"Store",
	"Set store",
	"Inline asm",
}

func (ast AstType) String() string {
	if len(astAsString) < int(ast) {
		return "Unknown"
	}

	return astAsString[ast]
}

// Basic tree node for the AST
type SyntaxTree struct {
	Type     AstType
	Children []*SyntaxTree
	Constant string
	VarType  string
}

func NewNode(typ AstType, v ...*SyntaxTree) *SyntaxTree {
	node := &SyntaxTree{Type: typ}
	node.Children = make([]*SyntaxTree, len(v))
	for i, child := range v {
		node.Children[i] = child
	}

	return node
}

func (p *SyntaxTree) prettyPrint(indent string) string {
	resp := fmt.Sprintln(indent, p.Type)
	resp += fmt.Sprintf("%s %q\n", indent, p.Constant)
	for _, child := range p.Children {
		resp += child.prettyPrint(indent + " |")
	}
	return resp
}

func (p *SyntaxTree) String() string {
	return p.prettyPrint("")
}

func MakeAst(source string) (*SyntaxTree, error) {
	lex := lexer("mutan", source)
	yyParse(lex)

	return Tree, lex.err
}

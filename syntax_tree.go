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
	BoolTy
	SetLocalTy
	BlockTy
	IfThenTy
	IfThenElseTy
	ForThenTy
	EqualTy
	GtTy
	LtTy
	OpTy
	StoreTy
	SetStoreTy
	InlineAsmTy
	StopTy
	ReturnTy

	OriginTy
	AddressTy
	CallerTy
	CallValTy
	CallDataLoadTy
	CallDataSizeTy
	GasPriceTy
	DiffTy
	PrevHashTy
	TimestampTy
	CoinbaseTy
	BalanceTy
	GasTy
	BlockNumTy

	NewArrayTy
	NewVarTy
	ArrayTy
	AssignArrayTy
	CallTy
	TransactTy
	CreateTy
	ArgTy
	SizeofTy
	StringTy

	LambdaTy

	FuncDefTy
	FuncCallTy

	NilTy
)

var astAsString = []string{
	"Invalid",
	"Empty",
	"Statement list",
	"Assignment",
	"Identifier",
	"Constant",
	"Boolean",
	"Set local",
	"Block",
	"If then",
	"If then else",
	"For then",
	"Equal",
	"Gt",
	"Lt",
	"Operator",
	"Store",
	"Set store",
	"Inline asm",
	"Stop",
	"Return",

	"Origin",
	"Address",
	"Caller",
	"Value",
	"Data load",
	"Data size",
	"Gas price",
	"Difficulty",
	"Previous hash",
	"Timestamp",
	"Coinbase",
	"Balance",
	"Gas",
	"Block number",

	"Create array",
	"New var",
	"Array",
	"Set array",
	"Call",
	"Transact",
	"Create",
	"Arg",
	"Sizeof",
	"String",

	"Lambda",
	"Func def",
	"Func call",

	"Nil",
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
	Size     string
	Lineno   int
}

func NewNode(typ AstType, v ...*SyntaxTree) *SyntaxTree {
	// Lineno is a global variable for ease of use
	node := &SyntaxTree{Type: typ, Lineno: Lineno}
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

// Creates an error and returns an ignore instruction
func (p *SyntaxTree) Errorf(format string, v ...interface{}) (*IntInstr, error) {
	return NewIntInstr(intIgnore, ""), fmt.Errorf("line %d: "+format, append([]interface{}{p.Lineno}, v...)...)
}

func MakeAst(source string) (*SyntaxTree, error) {
	lex := lexer("mutan", source)
	yyParse(lex)

	return Tree, lex.err
}

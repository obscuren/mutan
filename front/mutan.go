//line front/mutan.y:2
package frontend

import __yyfmt__ "fmt"

//line front/mutan.y:2
import (
	_ "fmt"
)

var Tree *SyntaxTree

// Helper function to turn a tree in to a regular list.
// Especially handy when parsing argument lists
func makeSlice(tree *SyntaxTree) (ret []*SyntaxTree) {
	if tree != nil && tree.Type != EmptyTy {
		ret = append(ret, tree)
		for _, i := range tree.Children {
			ret = append(ret, makeSlice(i)...)
		}
		tree.Children = nil
	}

	return
}

func makeArgs(tree *SyntaxTree, reverse bool) (ret []*SyntaxTree) {
	l := makeSlice(tree)
	// Quick reverse
	if reverse {
		for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
			l[i], l[j] = l[j], l[i]
		}
	}

	for _, s := range l {
		if s.Type != ArgTy {
			ret = append(ret, s)
		}
	}

	return
}

//line front/mutan.y:44
type yySymType struct {
	yys   int
	num   int
	str   string
	tnode *SyntaxTree
	check bool
}

const BLOCK = 57346
const TX = 57347
const CONTRACT = 57348
const CALL_S = 57349
const ADDR = 57350
const ORIGIN = 57351
const CALLER = 57352
const CALLVAL = 57353
const CALLDATALOAD = 57354
const CALLDATASIZE = 57355
const GASPRICE = 57356
const CALL = 57357
const SIZEOF = 57358
const EXIT = 57359
const CREATE = 57360
const BALANCE = 57361
const SHA3 = 57362
const DIFFICULTY = 57363
const PREVHASH = 57364
const TIMESTAMP = 57365
const BLOCKNUM = 57366
const COINBASE = 57367
const GAS = 57368
const ADDRESS = 57369
const BYTE = 57370
const PUSH = 57371
const POP = 57372
const TRANSACT = 57373
const STORE = 57374
const SUICIDE = 57375
const ASSIGN = 57376
const EQUAL = 57377
const END_STMT = 57378
const NIL = 57379
const VAR_ASSIGN = 57380
const LAMBDA = 57381
const COLON = 57382
const RETURN = 57383
const IF = 57384
const ELSE = 57385
const FOR = 57386
const LEFT_BRACES = 57387
const RIGHT_BRACES = 57388
const LEFT_BRACKET = 57389
const RIGHT_BRACKET = 57390
const ASM = 57391
const LEFT_PAR = 57392
const RIGHT_PAR = 57393
const STOP = 57394
const VAR = 57395
const FUNC = 57396
const FUNC_CALL = 57397
const IMPORT = 57398
const DOT = 57399
const ARRAY = 57400
const COMMA = 57401
const QUOTE = 57402
const ID = 57403
const NUMBER = 57404
const INLINE_ASM = 57405
const OP = 57406
const DOP = 57407
const STR = 57408
const BOOLEAN = 57409
const CODE = 57410
const oper = 57411
const AND = 57412
const MUL = 57413

var yyToknames = []string{
	"BLOCK",
	"TX",
	"CONTRACT",
	"CALL_S",
	"ADDR",
	"ORIGIN",
	"CALLER",
	"CALLVAL",
	"CALLDATALOAD",
	"CALLDATASIZE",
	"GASPRICE",
	"CALL",
	"SIZEOF",
	"EXIT",
	"CREATE",
	"BALANCE",
	"SHA3",
	"DIFFICULTY",
	"PREVHASH",
	"TIMESTAMP",
	"BLOCKNUM",
	"COINBASE",
	"GAS",
	"ADDRESS",
	"BYTE",
	"PUSH",
	"POP",
	"TRANSACT",
	"STORE",
	"SUICIDE",
	"ASSIGN",
	"EQUAL",
	"END_STMT",
	"NIL",
	"VAR_ASSIGN",
	"LAMBDA",
	"COLON",
	"RETURN",
	"IF",
	"ELSE",
	"FOR",
	"LEFT_BRACES",
	"RIGHT_BRACES",
	"LEFT_BRACKET",
	"RIGHT_BRACKET",
	"ASM",
	"LEFT_PAR",
	"RIGHT_PAR",
	"STOP",
	"VAR",
	"FUNC",
	"FUNC_CALL",
	"IMPORT",
	"DOT",
	"ARRAY",
	"COMMA",
	"QUOTE",
	"ID",
	"NUMBER",
	"INLINE_ASM",
	"OP",
	"DOP",
	"STR",
	"BOOLEAN",
	"CODE",
	"oper",
	"AND",
	"MUL",
}
var yyStatenames = []string{}

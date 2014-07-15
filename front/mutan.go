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

const ASSIGN = 57346
const EQUAL = 57347
const IF = 57348
const ELSE = 57349
const FOR = 57350
const LEFT_BRACES = 57351
const RIGHT_BRACES = 57352
const STORE = 57353
const LEFT_BRACKET = 57354
const RIGHT_BRACKET = 57355
const ASM = 57356
const LEFT_PAR = 57357
const RIGHT_PAR = 57358
const STOP = 57359
const ADDR = 57360
const ORIGIN = 57361
const CALLER = 57362
const CALLVAL = 57363
const CALLDATALOAD = 57364
const CALLDATASIZE = 57365
const GASPRICE = 57366
const DOT = 57367
const THIS = 57368
const ARRAY = 57369
const CALL = 57370
const COMMA = 57371
const SIZEOF = 57372
const QUOTE = 57373
const END_STMT = 57374
const EXIT = 57375
const CREATE = 57376
const TRANSACT = 57377
const NIL = 57378
const BALANCE = 57379
const VAR_ASSIGN = 57380
const LAMBDA = 57381
const COLON = 57382
const ADDRESS = 57383
const RETURN = 57384
const PUSH = 57385
const POP = 57386
const BYTE = 57387
const DIFFICULTY = 57388
const PREVHASH = 57389
const TIMESTAMP = 57390
const BLOCKNUM = 57391
const COINBASE = 57392
const GAS = 57393
const VAR = 57394
const FUNC = 57395
const FUNC_CALL = 57396
const IMPORT = 57397
const ID = 57398
const NUMBER = 57399
const INLINE_ASM = 57400
const OP = 57401
const DOP = 57402
const STR = 57403
const BOOLEAN = 57404
const CODE = 57405
const oper = 57406
const AND = 57407
const MUL = 57408

var yyToknames = []string{
	"ASSIGN",
	"EQUAL",
	"IF",
	"ELSE",
	"FOR",
	"LEFT_BRACES",
	"RIGHT_BRACES",
	"STORE",
	"LEFT_BRACKET",
	"RIGHT_BRACKET",
	"ASM",
	"LEFT_PAR",
	"RIGHT_PAR",
	"STOP",
	"ADDR",
	"ORIGIN",
	"CALLER",
	"CALLVAL",
	"CALLDATALOAD",
	"CALLDATASIZE",
	"GASPRICE",
	"DOT",
	"THIS",
	"ARRAY",
	"CALL",
	"COMMA",
	"SIZEOF",
	"QUOTE",
	"END_STMT",
	"EXIT",
	"CREATE",
	"TRANSACT",
	"NIL",
	"BALANCE",
	"VAR_ASSIGN",
	"LAMBDA",
	"COLON",
	"ADDRESS",
	"RETURN",
	"PUSH",
	"POP",
	"BYTE",
	"DIFFICULTY",
	"PREVHASH",
	"TIMESTAMP",
	"BLOCKNUM",
	"COINBASE",
	"GAS",
	"VAR",
	"FUNC",
	"FUNC_CALL",
	"IMPORT",
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

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line front/mutan.y:308

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 92
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 782

var yyAct = []int{

	32, 4, 172, 26, 204, 102, 2, 133, 33, 21,
	36, 74, 56, 93, 73, 19, 86, 49, 203, 44,
	59, 37, 70, 40, 35, 106, 16, 39, 38, 34,
	210, 19, 45, 65, 50, 15, 41, 42, 43, 139,
	51, 52, 195, 174, 150, 24, 114, 50, 105, 20,
	28, 72, 196, 51, 52, 29, 71, 46, 27, 25,
	194, 50, 19, 50, 99, 166, 68, 51, 52, 51,
	52, 169, 104, 107, 50, 19, 69, 111, 111, 111,
	51, 52, 115, 35, 34, 50, 109, 112, 113, 95,
	229, 51, 52, 19, 99, 221, 137, 220, 61, 173,
	136, 167, 138, 142, 227, 19, 62, 19, 202, 64,
	201, 147, 94, 146, 19, 145, 83, 22, 232, 230,
	206, 191, 189, 188, 187, 186, 185, 184, 182, 181,
	180, 179, 149, 148, 63, 116, 108, 103, 163, 92,
	162, 66, 64, 13, 161, 160, 111, 111, 159, 19,
	158, 19, 157, 156, 154, 175, 176, 183, 153, 54,
	57, 58, 152, 111, 151, 85, 192, 67, 19, 82,
	81, 199, 190, 80, 79, 78, 77, 19, 76, 75,
	207, 205, 19, 165, 144, 140, 164, 155, 92, 84,
	135, 30, 212, 211, 88, 89, 90, 91, 209, 208,
	47, 48, 111, 111, 96, 97, 98, 198, 215, 171,
	100, 213, 214, 217, 60, 216, 53, 222, 218, 219,
	193, 111, 111, 23, 226, 228, 117, 101, 19, 19,
	111, 19, 19, 231, 134, 197, 6, 17, 14, 18,
	10, 225, 118, 31, 141, 8, 33, 5, 36, 12,
	3, 1, 0, 0, 0, 0, 0, 44, 0, 37,
	0, 40, 35, 11, 16, 39, 38, 34, 0, 0,
	45, 0, 0, 15, 41, 42, 43, 0, 0, 0,
	0, 0, 0, 24, 7, 0, 9, 20, 28, 0,
	0, 177, 0, 29, 178, 0, 27, 25, 17, 0,
	18, 10, 224, 0, 0, 0, 8, 33, 0, 36,
	0, 0, 0, 0, 0, 200, 0, 0, 44, 0,
	37, 0, 40, 35, 11, 16, 39, 38, 34, 0,
	0, 45, 0, 0, 15, 41, 42, 43, 0, 0,
	0, 0, 0, 0, 24, 7, 0, 9, 20, 28,
	0, 0, 0, 0, 29, 0, 0, 27, 25, 17,
	0, 18, 10, 223, 0, 0, 0, 8, 33, 0,
	36, 0, 0, 0, 0, 0, 0, 0, 0, 44,
	0, 37, 0, 40, 35, 11, 16, 39, 38, 34,
	0, 0, 45, 0, 0, 15, 41, 42, 43, 0,
	0, 0, 0, 0, 0, 24, 7, 0, 9, 20,
	28, 0, 0, 0, 0, 29, 0, 0, 27, 25,
	17, 0, 18, 10, 170, 0, 0, 0, 8, 33,
	0, 36, 0, 0, 0, 0, 0, 0, 0, 0,
	44, 0, 37, 0, 40, 35, 11, 16, 39, 38,
	34, 0, 0, 45, 0, 0, 15, 41, 42, 43,
	0, 0, 0, 0, 0, 0, 24, 7, 0, 9,
	20, 28, 0, 0, 0, 0, 29, 0, 0, 27,
	25, 17, 0, 18, 10, 168, 0, 0, 0, 8,
	33, 0, 36, 0, 0, 0, 0, 0, 0, 0,
	0, 44, 0, 37, 0, 40, 35, 11, 16, 39,
	38, 34, 0, 0, 45, 0, 0, 15, 41, 42,
	43, 0, 0, 0, 0, 0, 0, 24, 7, 0,
	9, 20, 28, 0, 0, 0, 0, 29, 0, 0,
	27, 25, 17, 0, 18, 10, 87, 0, 0, 0,
	8, 33, 0, 36, 0, 0, 0, 0, 0, 0,
	0, 0, 44, 0, 37, 0, 40, 35, 11, 16,
	39, 38, 34, 0, 0, 45, 0, 0, 15, 41,
	42, 43, 0, 0, 0, 0, 0, 0, 24, 7,
	0, 9, 20, 28, 0, 0, 0, 17, 29, 18,
	10, 27, 25, 0, 0, 8, 33, 0, 36, 0,
	0, 0, 0, 0, 0, 0, 0, 44, 0, 37,
	0, 40, 35, 11, 16, 39, 38, 34, 0, 0,
	45, 0, 0, 15, 41, 42, 43, 0, 0, 0,
	0, 0, 0, 24, 7, 0, 9, 20, 28, 33,
	143, 36, 0, 29, 0, 0, 27, 25, 0, 0,
	44, 0, 37, 0, 40, 35, 0, 16, 39, 38,
	34, 0, 0, 45, 0, 0, 15, 41, 42, 43,
	0, 0, 33, 0, 36, 0, 24, 0, 0, 0,
	20, 28, 0, 44, 0, 37, 29, 40, 35, 27,
	25, 39, 38, 34, 0, 0, 45, 0, 36, 0,
	41, 42, 43, 0, 0, 0, 0, 44, 0, 37,
	0, 40, 35, 55, 28, 39, 38, 34, 0, 29,
	45, 0, 27, 25, 41, 42, 43, 0, 0, 0,
	0, 132, 0, 0, 0, 0, 0, 110, 28, 119,
	121, 122, 123, 29, 127, 0, 27, 25, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 130, 0, 0,
	0, 120, 0, 0, 0, 0, 124, 125, 126, 128,
	129, 131,
}
var yyPact = []int{

	-1000, -1000, 591, -1000, -1000, -1000, -1000, 1, 191, 52,
	-1000, -1000, -1000, -25, 212, 667, 667, 667, -7, 210,
	94, -1000, -1000, 667, 10, 0, -1000, -5, -1000, -1000,
	-1000, -1000, -46, -1000, -1000, -50, 164, 163, 161, 160,
	159, 158, 155, 154, 91, 177, 150, -42, -1000, 536,
	667, 667, 667, 667, -25, 127, -1000, -25, 4, 80,
	667, 667, -7, 206, -1000, 121, 121, -25, -1000, -8,
	-32, -1000, -1000, -1000, 42, 120, 691, 691, 691, -10,
	-7, 119, 667, 730, -56, -1000, 180, -1000, -25, -25,
	-25, -25, -7, -1000, -7, -1000, -25, -25, 26, 172,
	667, 634, -1000, -1000, -1000, -1000, 171, -1000, -1000, 86,
	176, -1000, 84, 82, 117, 116, -1000, 15, -1000, 149,
	147, 143, 139, 175, 138, 137, 135, 133, 130, 129,
	125, 123, 174, 170, 49, -1000, 475, 39, 414, 205,
	-1000, -25, 70, -1000, -13, 691, 691, 667, -1000, -1000,
	667, 115, 114, 113, 112, -7, 111, 110, 109, 108,
	107, 106, 691, 105, -7, -1000, 8, -14, 200, -7,
	-1000, 667, -1000, -1000, -1000, 81, 79, 2, -12, -1000,
	-1000, -1000, -1000, 168, -1000, -1000, -1000, -1000, -1000, -1000,
	104, -1000, 167, 190, -1000, 70, -26, -1000, 184, 183,
	-25, 691, 691, -1000, -1000, -1000, -1000, 204, -1000, -1000,
	70, -1000, -1000, 68, 66, -7, 353, -1000, 292, 231,
	48, 48, -1000, -1000, -1000, -1000, 61, -1000, 103, 48,
	-1000, 102, -1000,
}
var yyPgo = []int{

	0, 251, 6, 250, 1, 249, 143, 9, 247, 117,
	243, 242, 238, 238, 2, 0, 191, 236, 235, 3,
	234, 227, 12, 223, 5, 220,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 20, 20, 20, 25, 25, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 13, 13, 14,
	14, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 8, 18, 18,
	17, 17, 4, 4, 4, 4, 4, 4, 21, 21,
	5, 5, 5, 5, 5, 12, 12, 12, 6, 6,
	6, 6, 6, 9, 9, 9, 9, 7, 7, 7,
	7, 7, 7, 7, 7, 22, 19, 19, 15, 16,
	23, 24,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 1, 9, 4, 2,
	3, 1, 4, 5, 0, 1, 0, 3, 12, 10,
	6, 4, 4, 3, 6, 3, 4, 3, 0, 1,
	0, 3, 3, 3, 3, 4, 1, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 6, 6, 4, 0,
	9, 5, 1, 1, 1, 2, 2, 0, 3, 0,
	3, 3, 6, 3, 4, 2, 3, 5, 1, 1,
	3, 3, 4, 2, 3, 3, 3, 1, 1, 2,
	1, 4, 1, 1, 1, 2, 1, 1, 1, 3,
	1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -8, -17, 53, 14, 55,
	9, 32, -5, -6, -12, 42, 33, 6, 8, -22,
	56, -7, -9, -23, 52, 66, -19, 65, 57, 62,
	-16, -10, -15, 15, 36, 31, 17, 28, 35, 34,
	30, 43, 44, 45, 26, 39, 56, 9, -16, -2,
	59, 65, 66, 4, -6, 56, -22, -6, -6, -4,
	4, 4, 12, 40, 15, -7, -9, -6, 56, 66,
	12, 56, 56, 60, 61, 15, 15, 15, 15, 15,
	15, 15, 15, 25, 12, 15, 58, 10, -6, -6,
	-6, -6, 12, 9, 32, 9, -6, -6, -6, -4,
	4, -21, -24, 16, -24, 56, 57, 31, 16, -7,
	56, -15, -7, -7, 56, -4, 16, -6, -11, 19,
	41, 20, 21, 22, 46, 47, 48, 24, 49, 50,
	37, 51, 11, 63, -20, 10, -2, -4, -2, 13,
	13, -6, -4, 16, 13, 29, 29, 29, 16, 16,
	29, 15, 15, 15, 15, 12, 15, 15, 15, 15,
	15, 15, 15, 15, 12, 13, 16, 52, 10, 32,
	10, 4, -14, 29, 56, -7, -7, -6, -6, 16,
	16, 16, 16, -4, 16, 16, 16, 16, 16, 16,
	-7, 16, -4, -25, 52, 56, 66, -18, 7, -4,
	-6, 29, 29, 16, 16, 13, 16, 13, 9, -14,
	56, 9, 9, -7, -7, 4, -2, -14, -2, -2,
	29, 29, -4, 10, 10, 10, -19, 56, -19, 29,
	16, -19, 16,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 5, 6, 0, 0, 0,
	3, 11, 52, 53, 54, 0, 0, 0, 57, 78,
	88, 68, 69, 0, 0, 0, 77, 0, 80, 82,
	83, 84, 86, 90, 87, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 9, 0,
	0, 0, 0, 0, 55, 88, 78, 56, 0, 0,
	0, 0, 57, 0, 59, 68, 69, 0, 65, 0,
	0, 85, 79, 73, 0, 0, 0, 0, 0, 0,
	57, 0, 0, 0, 0, 14, 0, 10, 74, 75,
	76, 63, 57, 3, 57, 3, 60, 61, 0, 0,
	0, 57, 70, 91, 71, 66, 0, 89, 17, 0,
	88, 86, 0, 0, 0, 0, 23, 0, 25, 0,
	0, 0, 0, 36, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 8, 0, 0, 0, 0,
	81, 64, 30, 72, 0, 0, 0, 0, 21, 22,
	0, 0, 0, 0, 0, 57, 0, 0, 0, 0,
	0, 0, 0, 0, 57, 26, 16, 0, 49, 57,
	51, 0, 58, 29, 67, 0, 0, 0, 0, 31,
	32, 33, 34, 0, 37, 38, 39, 40, 41, 42,
	0, 44, 0, 0, 15, 30, 0, 47, 0, 0,
	62, 0, 0, 20, 24, 35, 43, 45, 3, 12,
	30, 3, 3, 0, 0, 57, 0, 13, 0, 0,
	0, 0, 46, 7, 48, 50, 0, 88, 0, 0,
	19, 0, 18,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line front/mutan.y:65
		{
			Tree = yyS[yypt-0].tnode
		}
	case 2:
		//line front/mutan.y:69
		{
			yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode)
		}
	case 3:
		//line front/mutan.y:70
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 4:
		//line front/mutan.y:75
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 5:
		//line front/mutan.y:76
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 6:
		//line front/mutan.y:77
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 7:
		//line front/mutan.y:79
		{
			yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-7].str
			yyVAL.tnode.HasRet = yyS[yypt-3].check
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-5].tnode, true)
		}
	case 8:
		//line front/mutan.y:85
		{
			yyVAL.tnode = NewNode(InlineAsmTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 9:
		//line front/mutan.y:86
		{
			yyVAL.tnode = NewNode(ImportTy)
			yyVAL.tnode.Constant = yyS[yypt-0].tnode.Constant
		}
	case 10:
		//line front/mutan.y:87
		{
			yyVAL.tnode = NewNode(ScopeTy, yyS[yypt-1].tnode)
		}
	case 11:
		//line front/mutan.y:88
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 12:
		//line front/mutan.y:92
		{
			yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 13:
		//line front/mutan.y:93
		{
			yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-4].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
			yyVAL.tnode.Ptr = true
		}
	case 14:
		//line front/mutan.y:94
		{
			yyVAL.tnode = nil
		}
	case 15:
		//line front/mutan.y:99
		{
			yyVAL.check = true
		}
	case 16:
		//line front/mutan.y:100
		{
			yyVAL.check = false
		}
	case 17:
		//line front/mutan.y:105
		{
			yyVAL.tnode = NewNode(StopTy)
		}
	case 18:
		//line front/mutan.y:108
		{
			yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 19:
		//line front/mutan.y:112
		{
			yyVAL.tnode = NewNode(TransactTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 20:
		//line front/mutan.y:115
		{
			yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 21:
		//line front/mutan.y:116
		{
			yyVAL.tnode = NewNode(SizeofTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 22:
		//line front/mutan.y:117
		{
			yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode)
		}
	case 23:
		//line front/mutan.y:118
		{
			yyVAL.tnode = NewNode(PopTy)
		}
	case 24:
		//line front/mutan.y:119
		{
			yyVAL.tnode = NewNode(ByteTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 25:
		//line front/mutan.y:120
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 26:
		//line front/mutan.y:121
		{
			yyVAL.tnode = NewNode(LambdaTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 27:
		//line front/mutan.y:125
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 28:
		//line front/mutan.y:126
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 29:
		//line front/mutan.y:130
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 30:
		//line front/mutan.y:131
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 31:
		//line front/mutan.y:136
		{
			yyVAL.tnode = NewNode(OriginTy)
		}
	case 32:
		//line front/mutan.y:137
		{
			yyVAL.tnode = NewNode(AddressTy)
		}
	case 33:
		//line front/mutan.y:138
		{
			yyVAL.tnode = NewNode(CallerTy)
		}
	case 34:
		//line front/mutan.y:139
		{
			yyVAL.tnode = NewNode(CallValTy)
		}
	case 35:
		//line front/mutan.y:140
		{
			yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode)
		}
	case 36:
		//line front/mutan.y:141
		{
			yyVAL.tnode = NewNode(CallDataSizeTy)
		}
	case 37:
		//line front/mutan.y:142
		{
			yyVAL.tnode = NewNode(DiffTy)
		}
	case 38:
		//line front/mutan.y:143
		{
			yyVAL.tnode = NewNode(PrevHashTy)
		}
	case 39:
		//line front/mutan.y:144
		{
			yyVAL.tnode = NewNode(TimestampTy)
		}
	case 40:
		//line front/mutan.y:145
		{
			yyVAL.tnode = NewNode(GasPriceTy)
		}
	case 41:
		//line front/mutan.y:146
		{
			yyVAL.tnode = NewNode(BlockNumTy)
		}
	case 42:
		//line front/mutan.y:147
		{
			yyVAL.tnode = NewNode(CoinbaseTy)
		}
	case 43:
		//line front/mutan.y:148
		{
			yyVAL.tnode = NewNode(BalanceTy, yyS[yypt-1].tnode)
		}
	case 44:
		//line front/mutan.y:149
		{
			yyVAL.tnode = NewNode(GasTy)
		}
	case 45:
		//line front/mutan.y:150
		{
			yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode)
		}
	case 46:
		//line front/mutan.y:152
		{
			node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 47:
		//line front/mutan.y:160
		{
			if yyS[yypt-0].tnode == nil {
				yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
			} else {
				yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			}
		}
	case 48:
		//line front/mutan.y:170
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 49:
		//line front/mutan.y:173
		{
			yyVAL.tnode = nil
		}
	case 50:
		//line front/mutan.y:178
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 51:
		//line front/mutan.y:182
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 52:
		//line front/mutan.y:188
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 53:
		//line front/mutan.y:189
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 54:
		//line front/mutan.y:190
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 55:
		//line front/mutan.y:191
		{
			yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode)
		}
	case 56:
		//line front/mutan.y:192
		{
			yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode)
		}
	case 57:
		//line front/mutan.y:193
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 58:
		//line front/mutan.y:197
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 59:
		//line front/mutan.y:198
		{
			yyVAL.tnode = nil
		}
	case 60:
		//line front/mutan.y:203
		{
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode)
		}
	case 61:
		//line front/mutan.y:207
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-2].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 62:
		//line front/mutan.y:213
		{
			yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-5].str
		}
	case 63:
		//line front/mutan.y:217
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-2].tnode.Constant
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		}
	case 64:
		//line front/mutan.y:223
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
			varNode := NewNode(NewVarTy)
			varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		}
	case 65:
		//line front/mutan.y:233
		{
			yyVAL.tnode = NewNode(NewVarTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 66:
		//line front/mutan.y:238
		{
			yyVAL.tnode = NewNode(NewVarTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
			yyVAL.tnode.Ptr = true
		}
	case 67:
		//line front/mutan.y:244
		{
			yyVAL.tnode = NewNode(NewArrayTy)
			yyVAL.tnode.Size = yyS[yypt-2].str
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 68:
		//line front/mutan.y:252
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 69:
		//line front/mutan.y:253
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 70:
		//line front/mutan.y:254
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 71:
		//line front/mutan.y:255
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 72:
		//line front/mutan.y:257
		{
			yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
		}
	case 73:
		//line front/mutan.y:266
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 74:
		//line front/mutan.y:268
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 75:
		//line front/mutan.y:269
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 76:
		//line front/mutan.y:270
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 77:
		//line front/mutan.y:274
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 78:
		//line front/mutan.y:275
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 79:
		//line front/mutan.y:276
		{
			yyVAL.tnode = NewNode(RefTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 80:
		//line front/mutan.y:277
		{
			yyVAL.tnode = NewNode(ConstantTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 81:
		//line front/mutan.y:278
		{
			yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
		}
	case 82:
		//line front/mutan.y:279
		{
			yyVAL.tnode = NewNode(BoolTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 83:
		//line front/mutan.y:280
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 84:
		//line front/mutan.y:281
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 85:
		//line front/mutan.y:285
		{
			yyVAL.tnode = NewNode(DerefPtrTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 86:
		//line front/mutan.y:289
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 87:
		//line front/mutan.y:290
		{
			yyVAL.tnode = NewNode(NilTy)
		}
	case 88:
		//line front/mutan.y:294
		{
			yyVAL.tnode = NewNode(IdentifierTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 89:
		//line front/mutan.y:298
		{
			yyVAL.tnode = NewNode(StringTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 90:
		//line front/mutan.y:302
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 91:
		//line front/mutan.y:306
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	}
	goto yystack /* stack new state and value */
}

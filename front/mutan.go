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

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line front/mutan.y:335

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 98
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 951

var yyAct = []int{

	32, 4, 194, 26, 116, 217, 2, 48, 49, 50,
	51, 63, 82, 153, 19, 218, 80, 56, 38, 41,
	66, 40, 45, 46, 227, 100, 226, 120, 225, 232,
	19, 44, 42, 43, 39, 159, 47, 57, 21, 57,
	35, 57, 52, 58, 59, 58, 59, 58, 59, 170,
	196, 57, 128, 34, 57, 37, 173, 58, 59, 119,
	58, 59, 72, 36, 62, 28, 121, 33, 19, 57,
	29, 113, 107, 27, 25, 58, 59, 57, 118, 35,
	77, 19, 79, 58, 59, 125, 125, 125, 250, 78,
	129, 57, 125, 125, 75, 53, 133, 58, 59, 36,
	19, 97, 243, 134, 76, 242, 195, 224, 113, 223,
	157, 172, 167, 166, 156, 165, 158, 162, 19, 96,
	19, 13, 95, 123, 126, 127, 94, 19, 216, 188,
	132, 189, 253, 251, 68, 214, 212, 61, 64, 65,
	70, 210, 209, 208, 207, 74, 206, 69, 205, 204,
	71, 185, 203, 202, 171, 81, 169, 168, 130, 122,
	117, 106, 186, 184, 71, 182, 125, 125, 19, 181,
	19, 180, 179, 178, 177, 176, 175, 174, 99, 102,
	103, 104, 105, 93, 92, 211, 91, 213, 90, 110,
	111, 112, 89, 221, 88, 19, 87, 19, 86, 85,
	84, 83, 229, 19, 197, 198, 228, 187, 164, 160,
	183, 106, 131, 98, 155, 135, 234, 108, 233, 230,
	231, 54, 220, 191, 125, 125, 109, 237, 193, 114,
	67, 60, 150, 147, 151, 239, 161, 238, 148, 244,
	240, 241, 22, 125, 125, 215, 248, 249, 152, 19,
	19, 125, 19, 19, 252, 138, 139, 137, 140, 141,
	23, 115, 235, 236, 154, 143, 73, 145, 30, 219,
	144, 48, 49, 50, 51, 6, 14, 31, 55, 5,
	149, 142, 38, 41, 16, 40, 45, 46, 146, 199,
	136, 12, 200, 3, 201, 44, 42, 43, 39, 1,
	47, 0, 0, 11, 35, 0, 52, 0, 15, 17,
	0, 18, 10, 247, 0, 222, 8, 34, 0, 37,
	24, 7, 0, 9, 0, 0, 0, 36, 20, 28,
	0, 33, 0, 0, 29, 0, 0, 27, 25, 48,
	49, 50, 51, 0, 0, 0, 0, 0, 0, 0,
	38, 41, 16, 40, 45, 46, 0, 0, 0, 0,
	0, 0, 0, 44, 42, 43, 39, 0, 47, 0,
	0, 11, 35, 0, 52, 0, 15, 17, 0, 18,
	10, 246, 0, 0, 8, 34, 0, 37, 24, 7,
	0, 9, 0, 0, 0, 36, 20, 28, 0, 33,
	0, 0, 29, 0, 0, 27, 25, 48, 49, 50,
	51, 0, 0, 0, 0, 0, 0, 0, 38, 41,
	16, 40, 45, 46, 0, 0, 0, 0, 0, 0,
	0, 44, 42, 43, 39, 0, 47, 0, 0, 11,
	35, 0, 52, 0, 15, 17, 0, 18, 10, 245,
	0, 0, 8, 34, 0, 37, 24, 7, 0, 9,
	0, 0, 0, 36, 20, 28, 0, 33, 0, 0,
	29, 0, 0, 27, 25, 48, 49, 50, 51, 0,
	0, 0, 0, 0, 0, 0, 38, 41, 16, 40,
	45, 46, 0, 0, 0, 0, 0, 0, 0, 44,
	42, 43, 39, 0, 47, 0, 0, 11, 35, 0,
	52, 0, 15, 17, 0, 18, 10, 192, 0, 0,
	8, 34, 0, 37, 24, 7, 0, 9, 0, 0,
	0, 36, 20, 28, 0, 33, 0, 0, 29, 0,
	0, 27, 25, 48, 49, 50, 51, 0, 0, 0,
	0, 0, 0, 0, 38, 41, 16, 40, 45, 46,
	0, 0, 0, 0, 0, 0, 0, 44, 42, 43,
	39, 0, 47, 0, 0, 11, 35, 0, 52, 0,
	15, 17, 0, 18, 10, 190, 0, 0, 8, 34,
	0, 37, 24, 7, 0, 9, 0, 0, 0, 36,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 48, 49, 50, 51, 0, 0, 0, 0, 0,
	0, 0, 38, 41, 16, 40, 45, 46, 0, 0,
	0, 0, 0, 0, 0, 44, 42, 43, 39, 0,
	47, 0, 0, 11, 35, 0, 52, 0, 15, 17,
	0, 18, 10, 101, 0, 0, 8, 34, 0, 37,
	24, 7, 0, 9, 0, 0, 0, 36, 20, 28,
	0, 33, 0, 0, 29, 0, 0, 27, 25, 48,
	49, 50, 51, 0, 0, 0, 0, 0, 0, 0,
	38, 41, 16, 40, 45, 46, 0, 0, 0, 0,
	0, 0, 0, 44, 42, 43, 39, 0, 47, 0,
	0, 11, 35, 0, 52, 0, 15, 17, 0, 18,
	10, 0, 0, 0, 8, 34, 0, 37, 24, 7,
	0, 9, 0, 0, 0, 36, 20, 28, 0, 33,
	0, 0, 29, 0, 0, 27, 25, 48, 49, 50,
	51, 0, 0, 0, 0, 0, 0, 0, 38, 41,
	16, 40, 45, 46, 0, 0, 0, 0, 0, 0,
	0, 44, 42, 43, 39, 0, 47, 0, 0, 0,
	35, 0, 52, 0, 15, 0, 0, 0, 0, 0,
	0, 0, 0, 34, 163, 37, 24, 0, 0, 0,
	0, 0, 0, 36, 20, 28, 0, 33, 0, 0,
	29, 0, 0, 27, 25, 48, 49, 50, 51, 0,
	0, 0, 0, 0, 0, 0, 38, 41, 16, 40,
	45, 46, 0, 0, 0, 0, 0, 0, 0, 44,
	42, 43, 39, 0, 47, 0, 0, 0, 35, 0,
	52, 0, 15, 0, 0, 0, 0, 0, 0, 0,
	0, 34, 0, 37, 24, 0, 0, 0, 0, 0,
	0, 36, 20, 28, 0, 33, 0, 0, 29, 0,
	0, 27, 25, 48, 49, 50, 51, 0, 0, 0,
	0, 0, 0, 0, 38, 41, 0, 40, 45, 46,
	0, 0, 0, 0, 0, 0, 0, 44, 42, 43,
	39, 0, 47, 0, 0, 0, 35, 0, 52, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 37, 0, 0, 0, 0, 0, 0, 0, 36,
	124, 28, 0, 0, 0, 0, 29, 0, 0, 27,
	25,
}
var yyPact = []int{

	-1000, -1000, 675, -1000, -1000, -1000, -1000, 34, 176, 39,
	-1000, -1000, -1000, 13, 197, 3, 3, 3, 811, 196,
	100, -1000, -1000, 3, 33, 28, -1000, 21, -1000, -1000,
	-1000, -1000, -49, 3, -1000, -1000, -54, 151, 150, 149,
	148, 146, 144, 142, 138, 136, 134, 133, 69, 65,
	62, 44, 166, 128, -38, -1000, 607, 3, 3, 3,
	3, 13, 114, -1000, 13, 27, 181, 3, 3, 811,
	195, -1000, 109, 109, 13, -1000, -2, -35, -1000, -1000,
	-1000, 13, 6, 108, 879, 879, 879, -9, 811, 107,
	3, 879, 42, 3, 234, 256, 206, 222, -55, -1000,
	168, -1000, 13, 13, 13, 13, 811, -1000, 811, -1000,
	13, 13, -13, 161, 3, 743, -1000, -1000, -1000, -1000,
	160, -1000, -1000, 56, 164, -1000, 54, 53, 106, 105,
	-1000, -10, 103, 52, -1000, 5, -1000, 127, 126, 125,
	124, 123, -1000, 122, 121, 119, -1000, 115, 163, -1000,
	113, 104, 112, 159, 78, -1000, 539, 187, 471, 194,
	-1000, 13, 47, -1000, -11, 879, 879, 3, -1000, -1000,
	3, -1000, 3, -1000, 102, 101, 98, 97, 95, 93,
	92, 91, 90, 811, 85, 811, 84, -1000, 75, -56,
	179, 811, -1000, 3, -1000, -1000, -1000, 50, 48, -23,
	-25, -27, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 158, -1000, 154, -1000, 174, -1000, 47, -32, -1000,
	173, 171, 13, 879, 879, -1000, -1000, -1000, 193, -1000,
	-1000, -1000, 47, -1000, -1000, 46, 43, 811, 403, -1000,
	335, 267, 42, 42, -1000, -1000, -1000, -1000, 29, 82,
	42, -1000, 81, -1000,
}
var yyPgo = []int{

	0, 299, 6, 293, 1, 291, 121, 38, 290, 288,
	281, 280, 279, 242, 277, 276, 276, 2, 0, 268,
	275, 269, 3, 264, 261, 11, 260, 4, 245,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 23, 23, 23, 28, 28, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 16, 16, 17, 17, 9, 9, 9,
	11, 11, 11, 11, 8, 8, 8, 8, 8, 10,
	10, 10, 12, 21, 21, 20, 20, 4, 4, 4,
	4, 4, 4, 24, 24, 5, 5, 5, 5, 5,
	15, 15, 15, 6, 6, 6, 6, 6, 13, 13,
	13, 13, 13, 7, 7, 7, 7, 7, 7, 7,
	7, 25, 22, 22, 18, 19, 26, 27,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 1, 9, 4, 2,
	3, 1, 4, 5, 0, 1, 0, 3, 12, 10,
	6, 4, 4, 3, 6, 4, 6, 4, 3, 3,
	3, 3, 4, 3, 0, 1, 0, 3, 4, 6,
	3, 4, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 6, 4, 0, 9, 5, 1, 1, 1,
	2, 2, 0, 3, 0, 3, 3, 6, 3, 4,
	2, 3, 5, 1, 1, 3, 3, 4, 2, 3,
	3, 3, 2, 1, 1, 2, 1, 4, 1, 1,
	1, 2, 1, 1, 1, 3, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -12, -20, 54, 49, 56,
	45, 36, -5, -6, -15, 41, 17, 42, 44, -25,
	61, -7, -13, -26, 53, 71, -22, 70, 62, 67,
	-19, -14, -18, 64, 50, 37, 60, 52, 15, 31,
	18, 16, 29, 30, 28, 19, 20, 33, 4, 5,
	6, 7, 39, 61, 45, -19, -2, 64, 70, 71,
	34, -6, 61, -25, -6, -6, -4, 34, 34, 47,
	40, 50, -7, -13, -6, 61, 71, 47, 61, 61,
	65, -6, 66, 50, 50, 50, 50, 50, 50, 50,
	50, 50, 50, 50, 57, 57, 57, 57, 47, 50,
	63, 46, -6, -6, -6, -6, 47, 45, 36, 45,
	-6, -6, -6, -4, 34, -24, -27, 51, -27, 61,
	62, 60, 51, -7, 61, -18, -7, -7, 61, -4,
	51, -6, -7, -22, 61, -6, -8, 23, 21, 22,
	24, 25, -10, 9, 14, 11, -9, 27, 32, -11,
	10, 12, 26, 68, -23, 46, -2, -4, -2, 48,
	48, -6, -4, 51, 48, 59, 59, 59, 51, 51,
	59, 51, 59, 51, 50, 50, 50, 50, 50, 50,
	50, 50, 50, 47, 50, 47, 50, 48, 51, 53,
	46, 36, 46, 34, -17, 59, 61, -7, -7, -6,
	-6, -6, 51, 51, 51, 51, 51, 51, 51, 51,
	51, -4, 51, -4, 51, -28, 53, 61, 71, -21,
	43, -4, -6, 59, 59, 51, 51, 51, 48, 48,
	45, -17, 61, 45, 45, -7, -7, 34, -2, -17,
	-2, -2, 59, 59, -4, 46, 46, 46, -22, -22,
	59, 51, -22, 51,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 5, 6, 0, 0, 0,
	3, 11, 57, 58, 59, 0, 0, 0, 62, 84,
	94, 73, 74, 0, 0, 0, 83, 0, 86, 88,
	89, 90, 92, 0, 96, 93, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 9, 0, 0, 0, 0,
	0, 60, 94, 84, 61, 0, 0, 0, 0, 62,
	0, 64, 73, 74, 0, 70, 0, 0, 91, 85,
	78, 82, 0, 0, 0, 0, 0, 0, 62, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 14,
	0, 10, 79, 80, 81, 68, 62, 3, 62, 3,
	65, 66, 0, 0, 0, 62, 75, 97, 76, 71,
	0, 95, 17, 0, 94, 92, 0, 0, 0, 0,
	23, 0, 0, 0, 94, 0, 28, 0, 0, 0,
	0, 0, 29, 0, 0, 0, 30, 0, 0, 31,
	0, 42, 0, 0, 0, 8, 0, 0, 0, 0,
	87, 69, 36, 77, 0, 0, 0, 0, 21, 22,
	0, 25, 0, 27, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 62, 0, 62, 0, 32, 16, 0,
	54, 62, 56, 0, 63, 35, 72, 0, 0, 0,
	0, 0, 44, 45, 46, 47, 48, 49, 50, 51,
	37, 0, 40, 0, 43, 0, 15, 36, 0, 52,
	0, 0, 67, 0, 0, 20, 24, 26, 38, 41,
	3, 12, 36, 3, 3, 0, 0, 62, 0, 13,
	0, 0, 0, 0, 39, 7, 53, 55, 0, 0,
	0, 19, 0, 18,
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
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
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
		//line front/mutan.y:76
		{
			Tree = yyS[yypt-0].tnode
		}
	case 2:
		//line front/mutan.y:80
		{
			yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode)
		}
	case 3:
		//line front/mutan.y:81
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 4:
		//line front/mutan.y:86
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 5:
		//line front/mutan.y:87
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 6:
		//line front/mutan.y:88
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 7:
		//line front/mutan.y:90
		{
			yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-7].str
			yyVAL.tnode.HasRet = yyS[yypt-3].check
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-5].tnode, true)
		}
	case 8:
		//line front/mutan.y:96
		{
			yyVAL.tnode = NewNode(InlineAsmTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 9:
		//line front/mutan.y:97
		{
			yyVAL.tnode = NewNode(ImportTy)
			yyVAL.tnode.Constant = yyS[yypt-0].tnode.Constant
		}
	case 10:
		//line front/mutan.y:98
		{
			yyVAL.tnode = NewNode(ScopeTy, yyS[yypt-1].tnode)
		}
	case 11:
		//line front/mutan.y:99
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 12:
		//line front/mutan.y:103
		{
			yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 13:
		//line front/mutan.y:104
		{
			yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-4].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
			yyVAL.tnode.Ptr = true
		}
	case 14:
		//line front/mutan.y:105
		{
			yyVAL.tnode = nil
		}
	case 15:
		//line front/mutan.y:110
		{
			yyVAL.check = true
		}
	case 16:
		//line front/mutan.y:111
		{
			yyVAL.check = false
		}
	case 17:
		//line front/mutan.y:116
		{
			yyVAL.tnode = NewNode(StopTy)
		}
	case 18:
		//line front/mutan.y:119
		{
			yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 19:
		//line front/mutan.y:123
		{
			yyVAL.tnode = NewNode(TransactTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 20:
		//line front/mutan.y:126
		{
			yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 21:
		//line front/mutan.y:127
		{
			yyVAL.tnode = NewNode(SizeofTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 22:
		//line front/mutan.y:128
		{
			yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode)
		}
	case 23:
		//line front/mutan.y:129
		{
			yyVAL.tnode = NewNode(PopTy)
		}
	case 24:
		//line front/mutan.y:130
		{
			yyVAL.tnode = NewNode(ByteTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 25:
		//line front/mutan.y:131
		{
			yyVAL.tnode = NewNode(BalanceTy, yyS[yypt-1].tnode)
		}
	case 26:
		//line front/mutan.y:132
		{
			yyVAL.tnode = NewNode(Sha3Ty, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 27:
		//line front/mutan.y:133
		{
			yyVAL.tnode = NewNode(SuicideTy, yyS[yypt-1].tnode)
		}
	case 28:
		//line front/mutan.y:134
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 29:
		//line front/mutan.y:135
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 30:
		//line front/mutan.y:136
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 31:
		//line front/mutan.y:137
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 32:
		//line front/mutan.y:138
		{
			yyVAL.tnode = NewNode(LambdaTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 33:
		//line front/mutan.y:142
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 34:
		//line front/mutan.y:143
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 35:
		//line front/mutan.y:147
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 36:
		//line front/mutan.y:148
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 37:
		//line front/mutan.y:153
		{
			yyVAL.tnode = NewNode(AddressTy)
		}
	case 38:
		//line front/mutan.y:154
		{
			yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode)
		}
	case 39:
		//line front/mutan.y:156
		{
			node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 40:
		//line front/mutan.y:163
		{
			yyVAL.tnode = NewNode(CallerTy)
		}
	case 41:
		//line front/mutan.y:164
		{
			yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode)
		}
	case 42:
		//line front/mutan.y:165
		{
			yyVAL.tnode = NewNode(CallDataSizeTy)
		}
	case 43:
		//line front/mutan.y:166
		{
			yyVAL.tnode = NewNode(GasTy)
		}
	case 44:
		//line front/mutan.y:171
		{
			yyVAL.tnode = NewNode(TimestampTy)
		}
	case 45:
		//line front/mutan.y:172
		{
			yyVAL.tnode = NewNode(DiffTy)
		}
	case 46:
		//line front/mutan.y:173
		{
			yyVAL.tnode = NewNode(PrevHashTy)
		}
	case 47:
		//line front/mutan.y:174
		{
			yyVAL.tnode = NewNode(BlockNumTy)
		}
	case 48:
		//line front/mutan.y:175
		{
			yyVAL.tnode = NewNode(CoinbaseTy)
		}
	case 49:
		//line front/mutan.y:179
		{
			yyVAL.tnode = NewNode(OriginTy)
		}
	case 50:
		//line front/mutan.y:180
		{
			yyVAL.tnode = NewNode(GasPriceTy)
		}
	case 51:
		//line front/mutan.y:181
		{
			yyVAL.tnode = NewNode(CallValTy)
		}
	case 52:
		//line front/mutan.y:186
		{
			if yyS[yypt-0].tnode == nil {
				yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
			} else {
				yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			}
		}
	case 53:
		//line front/mutan.y:196
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 54:
		//line front/mutan.y:199
		{
			yyVAL.tnode = nil
		}
	case 55:
		//line front/mutan.y:204
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 56:
		//line front/mutan.y:208
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 57:
		//line front/mutan.y:214
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 58:
		//line front/mutan.y:215
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 59:
		//line front/mutan.y:216
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 60:
		//line front/mutan.y:217
		{
			yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode)
		}
	case 61:
		//line front/mutan.y:218
		{
			yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode)
		}
	case 62:
		//line front/mutan.y:219
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 63:
		//line front/mutan.y:223
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 64:
		//line front/mutan.y:224
		{
			yyVAL.tnode = nil
		}
	case 65:
		//line front/mutan.y:229
		{
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode)
		}
	case 66:
		//line front/mutan.y:233
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-2].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 67:
		//line front/mutan.y:239
		{
			yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-5].str
		}
	case 68:
		//line front/mutan.y:243
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-2].tnode.Constant
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		}
	case 69:
		//line front/mutan.y:249
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
			varNode := NewNode(NewVarTy)
			varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		}
	case 70:
		//line front/mutan.y:259
		{
			yyVAL.tnode = NewNode(NewVarTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 71:
		//line front/mutan.y:264
		{
			yyVAL.tnode = NewNode(NewVarTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
			yyVAL.tnode.Ptr = true
		}
	case 72:
		//line front/mutan.y:270
		{
			yyVAL.tnode = NewNode(NewArrayTy)
			yyVAL.tnode.Size = yyS[yypt-2].str
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 73:
		//line front/mutan.y:278
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 74:
		//line front/mutan.y:279
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 75:
		//line front/mutan.y:280
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 76:
		//line front/mutan.y:281
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 77:
		//line front/mutan.y:283
		{
			yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
		}
	case 78:
		//line front/mutan.y:292
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 79:
		//line front/mutan.y:294
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 80:
		//line front/mutan.y:295
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 81:
		//line front/mutan.y:296
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 82:
		//line front/mutan.y:297
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 83:
		//line front/mutan.y:301
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 84:
		//line front/mutan.y:302
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 85:
		//line front/mutan.y:303
		{
			yyVAL.tnode = NewNode(RefTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 86:
		//line front/mutan.y:304
		{
			yyVAL.tnode = NewNode(ConstantTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 87:
		//line front/mutan.y:305
		{
			yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
		}
	case 88:
		//line front/mutan.y:306
		{
			yyVAL.tnode = NewNode(BoolTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 89:
		//line front/mutan.y:307
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 90:
		//line front/mutan.y:308
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 91:
		//line front/mutan.y:312
		{
			yyVAL.tnode = NewNode(DerefPtrTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 92:
		//line front/mutan.y:316
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 93:
		//line front/mutan.y:317
		{
			yyVAL.tnode = NewNode(NilTy)
		}
	case 94:
		//line front/mutan.y:321
		{
			yyVAL.tnode = NewNode(IdentifierTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 95:
		//line front/mutan.y:325
		{
			yyVAL.tnode = NewNode(StringTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 96:
		//line front/mutan.y:329
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 97:
		//line front/mutan.y:333
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	}
	goto yystack /* stack new state and value */
}

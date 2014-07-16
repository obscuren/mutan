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
const ASSIGN = 57375
const EQUAL = 57376
const END_STMT = 57377
const NIL = 57378
const VAR_ASSIGN = 57379
const LAMBDA = 57380
const COLON = 57381
const RETURN = 57382
const IF = 57383
const ELSE = 57384
const FOR = 57385
const LEFT_BRACES = 57386
const RIGHT_BRACES = 57387
const LEFT_BRACKET = 57388
const RIGHT_BRACKET = 57389
const ASM = 57390
const LEFT_PAR = 57391
const RIGHT_PAR = 57392
const STOP = 57393
const VAR = 57394
const FUNC = 57395
const FUNC_CALL = 57396
const IMPORT = 57397
const DOT = 57398
const ARRAY = 57399
const COMMA = 57400
const QUOTE = 57401
const ID = 57402
const NUMBER = 57403
const INLINE_ASM = 57404
const OP = 57405
const DOP = 57406
const STR = 57407
const BOOLEAN = 57408
const CODE = 57409
const oper = 57410
const AND = 57411
const MUL = 57412

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

//line front/mutan.y:333

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 97
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 934

var yyAct = []int{

	32, 4, 190, 26, 150, 81, 2, 47, 48, 49,
	50, 62, 79, 98, 19, 114, 223, 55, 38, 41,
	65, 40, 45, 46, 105, 222, 118, 221, 228, 56,
	19, 44, 42, 43, 39, 57, 58, 21, 56, 35,
	56, 51, 192, 56, 57, 58, 57, 58, 167, 57,
	58, 56, 34, 56, 37, 156, 126, 57, 58, 57,
	58, 71, 36, 61, 28, 213, 33, 19, 119, 29,
	111, 56, 27, 25, 35, 214, 76, 57, 58, 117,
	19, 78, 77, 52, 123, 123, 123, 36, 116, 127,
	74, 123, 123, 246, 239, 131, 95, 238, 132, 19,
	75, 191, 220, 219, 169, 164, 111, 163, 154, 162,
	94, 93, 153, 92, 155, 159, 19, 13, 19, 212,
	249, 121, 124, 125, 184, 19, 185, 247, 130, 210,
	67, 208, 206, 60, 63, 64, 69, 205, 204, 203,
	202, 73, 201, 68, 200, 199, 70, 181, 198, 168,
	166, 80, 165, 128, 120, 115, 104, 182, 180, 70,
	178, 177, 176, 123, 123, 19, 175, 19, 174, 173,
	172, 171, 170, 97, 100, 101, 102, 103, 91, 90,
	89, 207, 88, 209, 108, 109, 110, 87, 86, 217,
	85, 19, 84, 19, 83, 82, 225, 224, 183, 19,
	193, 194, 161, 157, 179, 104, 96, 129, 152, 106,
	230, 229, 226, 53, 216, 187, 227, 233, 107, 189,
	123, 123, 144, 112, 66, 59, 211, 145, 30, 23,
	158, 235, 22, 234, 113, 240, 236, 237, 54, 123,
	123, 151, 244, 245, 215, 19, 19, 123, 19, 19,
	248, 135, 136, 134, 137, 138, 72, 231, 232, 6,
	14, 31, 5, 146, 47, 48, 49, 50, 147, 139,
	148, 143, 133, 12, 3, 38, 41, 16, 40, 45,
	46, 140, 195, 142, 149, 196, 141, 197, 44, 42,
	43, 39, 1, 0, 0, 11, 35, 0, 51, 0,
	15, 17, 0, 18, 10, 243, 0, 218, 8, 34,
	0, 37, 24, 7, 0, 9, 0, 0, 0, 36,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 47, 48, 49, 50, 0, 0, 0, 0, 0,
	0, 0, 38, 41, 16, 40, 45, 46, 0, 0,
	0, 0, 0, 0, 0, 44, 42, 43, 39, 0,
	0, 0, 11, 35, 0, 51, 0, 15, 17, 0,
	18, 10, 242, 0, 0, 8, 34, 0, 37, 24,
	7, 0, 9, 0, 0, 0, 36, 20, 28, 0,
	33, 0, 0, 29, 0, 0, 27, 25, 47, 48,
	49, 50, 0, 0, 0, 0, 0, 0, 0, 38,
	41, 16, 40, 45, 46, 0, 0, 0, 0, 0,
	0, 0, 44, 42, 43, 39, 0, 0, 0, 11,
	35, 0, 51, 0, 15, 17, 0, 18, 10, 241,
	0, 0, 8, 34, 0, 37, 24, 7, 0, 9,
	0, 0, 0, 36, 20, 28, 0, 33, 0, 0,
	29, 0, 0, 27, 25, 47, 48, 49, 50, 0,
	0, 0, 0, 0, 0, 0, 38, 41, 16, 40,
	45, 46, 0, 0, 0, 0, 0, 0, 0, 44,
	42, 43, 39, 0, 0, 0, 11, 35, 0, 51,
	0, 15, 17, 0, 18, 10, 188, 0, 0, 8,
	34, 0, 37, 24, 7, 0, 9, 0, 0, 0,
	36, 20, 28, 0, 33, 0, 0, 29, 0, 0,
	27, 25, 47, 48, 49, 50, 0, 0, 0, 0,
	0, 0, 0, 38, 41, 16, 40, 45, 46, 0,
	0, 0, 0, 0, 0, 0, 44, 42, 43, 39,
	0, 0, 0, 11, 35, 0, 51, 0, 15, 17,
	0, 18, 10, 186, 0, 0, 8, 34, 0, 37,
	24, 7, 0, 9, 0, 0, 0, 36, 20, 28,
	0, 33, 0, 0, 29, 0, 0, 27, 25, 47,
	48, 49, 50, 0, 0, 0, 0, 0, 0, 0,
	38, 41, 16, 40, 45, 46, 0, 0, 0, 0,
	0, 0, 0, 44, 42, 43, 39, 0, 0, 0,
	11, 35, 0, 51, 0, 15, 17, 0, 18, 10,
	99, 0, 0, 8, 34, 0, 37, 24, 7, 0,
	9, 0, 0, 0, 36, 20, 28, 0, 33, 0,
	0, 29, 0, 0, 27, 25, 47, 48, 49, 50,
	0, 0, 0, 0, 0, 0, 0, 38, 41, 16,
	40, 45, 46, 0, 0, 0, 0, 0, 0, 0,
	44, 42, 43, 39, 0, 0, 0, 11, 35, 0,
	51, 0, 15, 17, 0, 18, 10, 0, 0, 0,
	8, 34, 0, 37, 24, 7, 0, 9, 0, 0,
	0, 36, 20, 28, 0, 33, 0, 0, 29, 0,
	0, 27, 25, 47, 48, 49, 50, 0, 0, 0,
	0, 0, 0, 0, 38, 41, 16, 40, 45, 46,
	0, 0, 0, 0, 0, 0, 0, 44, 42, 43,
	39, 0, 0, 0, 0, 35, 0, 51, 0, 15,
	0, 0, 0, 0, 0, 0, 0, 0, 34, 160,
	37, 24, 0, 0, 0, 0, 0, 0, 36, 20,
	28, 0, 33, 0, 0, 29, 0, 0, 27, 25,
	47, 48, 49, 50, 0, 0, 0, 0, 0, 0,
	0, 38, 41, 16, 40, 45, 46, 0, 0, 0,
	0, 0, 0, 0, 44, 42, 43, 39, 0, 0,
	0, 0, 35, 0, 51, 0, 15, 0, 0, 0,
	0, 0, 0, 0, 0, 34, 0, 37, 24, 0,
	0, 0, 0, 0, 0, 36, 20, 28, 0, 33,
	0, 0, 29, 0, 0, 27, 25, 47, 48, 49,
	50, 0, 0, 0, 0, 0, 0, 0, 38, 41,
	0, 40, 45, 46, 0, 0, 0, 0, 0, 0,
	0, 44, 42, 43, 39, 0, 0, 0, 0, 35,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 37, 0, 0, 0, 0, 0,
	0, 0, 36, 122, 28, 0, 0, 0, 0, 29,
	0, 0, 27, 25,
}
var yyPact = []int{

	-1000, -1000, 662, -1000, -1000, -1000, -1000, 23, 169, 28,
	-1000, -1000, -1000, -12, 192, 3, 3, 3, 796, 191,
	97, -1000, -1000, 3, 30, 22, -1000, 21, -1000, -1000,
	-1000, -1000, -52, 3, -1000, -1000, -60, 146, 145, 143,
	141, 139, 138, 133, 131, 130, 129, 57, 55, 54,
	40, 160, 124, -49, -1000, 595, 3, 3, 3, 3,
	-12, 110, -1000, -12, -20, 174, 3, 3, 796, 190,
	-1000, 105, 105, -12, -1000, 19, -35, -1000, -1000, -1000,
	-12, 9, 104, 863, 863, 863, -4, 796, 103, 3,
	863, 38, 230, 272, 195, 258, -63, -1000, 163, -1000,
	-12, -12, -12, -12, 796, -1000, 796, -1000, -12, -12,
	8, 156, 3, 729, -1000, -1000, -1000, -1000, 155, -1000,
	-1000, 51, 159, -1000, 49, 47, 102, 100, -1000, -10,
	99, 46, -1000, -1000, 123, 122, 121, 120, 119, -1000,
	117, 113, 112, -1000, 111, 158, -1000, 109, 101, 108,
	151, 74, -1000, 528, 180, 461, 186, -1000, -12, 43,
	-1000, -18, 863, 863, 3, -1000, -1000, 3, -1000, 3,
	98, 95, 94, 92, 90, 89, 88, 87, 82, 796,
	81, 796, 79, -1000, 67, 5, 172, 796, -1000, 3,
	-1000, -1000, -1000, 45, 44, -23, -25, -34, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 150, -1000, 149,
	-1000, 168, -1000, 43, -32, -1000, 167, 166, -12, 863,
	863, -1000, -1000, -1000, 184, -1000, -1000, -1000, 43, -1000,
	-1000, 39, 36, 796, 394, -1000, 327, 260, 38, 38,
	-1000, -1000, -1000, -1000, 35, 77, 38, -1000, 70, -1000,
}
var yyPgo = []int{

	0, 292, 6, 274, 1, 273, 117, 37, 272, 271,
	269, 263, 262, 232, 261, 260, 260, 2, 0, 228,
	259, 244, 3, 241, 234, 11, 229, 15, 226,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 23, 23, 23, 28, 28, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 16, 16, 17, 17, 9, 9, 9, 11,
	11, 11, 11, 8, 8, 8, 8, 8, 10, 10,
	10, 12, 21, 21, 20, 20, 4, 4, 4, 4,
	4, 4, 24, 24, 5, 5, 5, 5, 5, 15,
	15, 15, 6, 6, 6, 6, 6, 13, 13, 13,
	13, 13, 7, 7, 7, 7, 7, 7, 7, 7,
	25, 22, 22, 18, 19, 26, 27,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 1, 9, 4, 2,
	3, 1, 4, 5, 0, 1, 0, 3, 12, 10,
	6, 4, 4, 3, 6, 4, 6, 3, 3, 3,
	3, 4, 3, 0, 1, 0, 3, 4, 6, 3,
	4, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 6, 4, 0, 9, 5, 1, 1, 1, 2,
	2, 0, 3, 0, 3, 3, 6, 3, 4, 2,
	3, 5, 1, 1, 3, 3, 4, 2, 3, 3,
	3, 2, 1, 1, 2, 1, 4, 1, 1, 1,
	2, 1, 1, 1, 3, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -12, -20, 53, 48, 55,
	44, 35, -5, -6, -15, 40, 17, 41, 43, -25,
	60, -7, -13, -26, 52, 70, -22, 69, 61, 66,
	-19, -14, -18, 63, 49, 36, 59, 51, 15, 31,
	18, 16, 29, 30, 28, 19, 20, 4, 5, 6,
	7, 38, 60, 44, -19, -2, 63, 69, 70, 33,
	-6, 60, -25, -6, -6, -4, 33, 33, 46, 39,
	49, -7, -13, -6, 60, 70, 46, 60, 60, 64,
	-6, 65, 49, 49, 49, 49, 49, 49, 49, 49,
	49, 49, 56, 56, 56, 56, 46, 49, 62, 45,
	-6, -6, -6, -6, 46, 44, 35, 44, -6, -6,
	-6, -4, 33, -24, -27, 50, -27, 60, 61, 59,
	50, -7, 60, -18, -7, -7, 60, -4, 50, -6,
	-7, -22, 60, -8, 23, 21, 22, 24, 25, -10,
	9, 14, 11, -9, 27, 32, -11, 10, 12, 26,
	67, -23, 45, -2, -4, -2, 47, 47, -6, -4,
	50, 47, 58, 58, 58, 50, 50, 58, 50, 58,
	49, 49, 49, 49, 49, 49, 49, 49, 49, 46,
	49, 46, 49, 47, 50, 52, 45, 35, 45, 33,
	-17, 58, 60, -7, -7, -6, -6, -6, 50, 50,
	50, 50, 50, 50, 50, 50, 50, -4, 50, -4,
	50, -28, 52, 60, 70, -21, 42, -4, -6, 58,
	58, 50, 50, 50, 47, 47, 44, -17, 60, 44,
	44, -7, -7, 33, -2, -17, -2, -2, 58, 58,
	-4, 45, 45, 45, -22, -22, 58, 50, -22, 50,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 5, 6, 0, 0, 0,
	3, 11, 56, 57, 58, 0, 0, 0, 61, 83,
	93, 72, 73, 0, 0, 0, 82, 0, 85, 87,
	88, 89, 91, 0, 95, 92, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0, 0, 0,
	59, 93, 83, 60, 0, 0, 0, 0, 61, 0,
	63, 72, 73, 0, 69, 0, 0, 90, 84, 77,
	81, 0, 0, 0, 0, 0, 0, 61, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 14, 0, 10,
	78, 79, 80, 67, 61, 3, 61, 3, 64, 65,
	0, 0, 0, 61, 74, 96, 75, 70, 0, 94,
	17, 0, 93, 91, 0, 0, 0, 0, 23, 0,
	0, 0, 93, 27, 0, 0, 0, 0, 0, 28,
	0, 0, 0, 29, 0, 0, 30, 0, 41, 0,
	0, 0, 8, 0, 0, 0, 0, 86, 68, 35,
	76, 0, 0, 0, 0, 21, 22, 0, 25, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 61,
	0, 61, 0, 31, 16, 0, 53, 61, 55, 0,
	62, 34, 71, 0, 0, 0, 0, 0, 43, 44,
	45, 46, 47, 48, 49, 50, 36, 0, 39, 0,
	42, 0, 15, 35, 0, 51, 0, 0, 66, 0,
	0, 20, 24, 26, 37, 40, 3, 12, 35, 3,
	3, 0, 0, 61, 0, 13, 0, 0, 0, 0,
	38, 7, 52, 54, 0, 0, 0, 19, 0, 18,
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
	62, 63, 64, 65, 66, 67, 68, 69, 70,
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
		//line front/mutan.y:75
		{
			Tree = yyS[yypt-0].tnode
		}
	case 2:
		//line front/mutan.y:79
		{
			yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode)
		}
	case 3:
		//line front/mutan.y:80
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 4:
		//line front/mutan.y:85
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 5:
		//line front/mutan.y:86
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 6:
		//line front/mutan.y:87
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 7:
		//line front/mutan.y:89
		{
			yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-7].str
			yyVAL.tnode.HasRet = yyS[yypt-3].check
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-5].tnode, true)
		}
	case 8:
		//line front/mutan.y:95
		{
			yyVAL.tnode = NewNode(InlineAsmTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 9:
		//line front/mutan.y:96
		{
			yyVAL.tnode = NewNode(ImportTy)
			yyVAL.tnode.Constant = yyS[yypt-0].tnode.Constant
		}
	case 10:
		//line front/mutan.y:97
		{
			yyVAL.tnode = NewNode(ScopeTy, yyS[yypt-1].tnode)
		}
	case 11:
		//line front/mutan.y:98
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 12:
		//line front/mutan.y:102
		{
			yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 13:
		//line front/mutan.y:103
		{
			yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-4].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
			yyVAL.tnode.Ptr = true
		}
	case 14:
		//line front/mutan.y:104
		{
			yyVAL.tnode = nil
		}
	case 15:
		//line front/mutan.y:109
		{
			yyVAL.check = true
		}
	case 16:
		//line front/mutan.y:110
		{
			yyVAL.check = false
		}
	case 17:
		//line front/mutan.y:115
		{
			yyVAL.tnode = NewNode(StopTy)
		}
	case 18:
		//line front/mutan.y:118
		{
			yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 19:
		//line front/mutan.y:122
		{
			yyVAL.tnode = NewNode(TransactTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 20:
		//line front/mutan.y:125
		{
			yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 21:
		//line front/mutan.y:126
		{
			yyVAL.tnode = NewNode(SizeofTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 22:
		//line front/mutan.y:127
		{
			yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode)
		}
	case 23:
		//line front/mutan.y:128
		{
			yyVAL.tnode = NewNode(PopTy)
		}
	case 24:
		//line front/mutan.y:129
		{
			yyVAL.tnode = NewNode(ByteTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 25:
		//line front/mutan.y:130
		{
			yyVAL.tnode = NewNode(BalanceTy, yyS[yypt-1].tnode)
		}
	case 26:
		//line front/mutan.y:131
		{
			yyVAL.tnode = NewNode(Sha3Ty, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 27:
		//line front/mutan.y:132
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 28:
		//line front/mutan.y:133
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 29:
		//line front/mutan.y:134
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 30:
		//line front/mutan.y:135
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 31:
		//line front/mutan.y:136
		{
			yyVAL.tnode = NewNode(LambdaTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 32:
		//line front/mutan.y:140
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 33:
		//line front/mutan.y:141
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 34:
		//line front/mutan.y:145
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 35:
		//line front/mutan.y:146
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 36:
		//line front/mutan.y:151
		{
			yyVAL.tnode = NewNode(AddressTy)
		}
	case 37:
		//line front/mutan.y:152
		{
			yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode)
		}
	case 38:
		//line front/mutan.y:154
		{
			node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 39:
		//line front/mutan.y:161
		{
			yyVAL.tnode = NewNode(CallerTy)
		}
	case 40:
		//line front/mutan.y:162
		{
			yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode)
		}
	case 41:
		//line front/mutan.y:163
		{
			yyVAL.tnode = NewNode(CallDataSizeTy)
		}
	case 42:
		//line front/mutan.y:164
		{
			yyVAL.tnode = NewNode(GasTy)
		}
	case 43:
		//line front/mutan.y:169
		{
			yyVAL.tnode = NewNode(TimestampTy)
		}
	case 44:
		//line front/mutan.y:170
		{
			yyVAL.tnode = NewNode(DiffTy)
		}
	case 45:
		//line front/mutan.y:171
		{
			yyVAL.tnode = NewNode(PrevHashTy)
		}
	case 46:
		//line front/mutan.y:172
		{
			yyVAL.tnode = NewNode(BlockNumTy)
		}
	case 47:
		//line front/mutan.y:173
		{
			yyVAL.tnode = NewNode(CoinbaseTy)
		}
	case 48:
		//line front/mutan.y:177
		{
			yyVAL.tnode = NewNode(OriginTy)
		}
	case 49:
		//line front/mutan.y:178
		{
			yyVAL.tnode = NewNode(GasPriceTy)
		}
	case 50:
		//line front/mutan.y:179
		{
			yyVAL.tnode = NewNode(CallValTy)
		}
	case 51:
		//line front/mutan.y:184
		{
			if yyS[yypt-0].tnode == nil {
				yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
			} else {
				yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			}
		}
	case 52:
		//line front/mutan.y:194
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 53:
		//line front/mutan.y:197
		{
			yyVAL.tnode = nil
		}
	case 54:
		//line front/mutan.y:202
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 55:
		//line front/mutan.y:206
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 56:
		//line front/mutan.y:212
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 57:
		//line front/mutan.y:213
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 58:
		//line front/mutan.y:214
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 59:
		//line front/mutan.y:215
		{
			yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode)
		}
	case 60:
		//line front/mutan.y:216
		{
			yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode)
		}
	case 61:
		//line front/mutan.y:217
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 62:
		//line front/mutan.y:221
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 63:
		//line front/mutan.y:222
		{
			yyVAL.tnode = nil
		}
	case 64:
		//line front/mutan.y:227
		{
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode)
		}
	case 65:
		//line front/mutan.y:231
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-2].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 66:
		//line front/mutan.y:237
		{
			yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-5].str
		}
	case 67:
		//line front/mutan.y:241
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-2].tnode.Constant
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		}
	case 68:
		//line front/mutan.y:247
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
			varNode := NewNode(NewVarTy)
			varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		}
	case 69:
		//line front/mutan.y:257
		{
			yyVAL.tnode = NewNode(NewVarTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 70:
		//line front/mutan.y:262
		{
			yyVAL.tnode = NewNode(NewVarTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
			yyVAL.tnode.Ptr = true
		}
	case 71:
		//line front/mutan.y:268
		{
			yyVAL.tnode = NewNode(NewArrayTy)
			yyVAL.tnode.Size = yyS[yypt-2].str
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 72:
		//line front/mutan.y:276
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 73:
		//line front/mutan.y:277
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 74:
		//line front/mutan.y:278
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 75:
		//line front/mutan.y:279
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 76:
		//line front/mutan.y:281
		{
			yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
		}
	case 77:
		//line front/mutan.y:290
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 78:
		//line front/mutan.y:292
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 79:
		//line front/mutan.y:293
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 80:
		//line front/mutan.y:294
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 81:
		//line front/mutan.y:295
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 82:
		//line front/mutan.y:299
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 83:
		//line front/mutan.y:300
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 84:
		//line front/mutan.y:301
		{
			yyVAL.tnode = NewNode(RefTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 85:
		//line front/mutan.y:302
		{
			yyVAL.tnode = NewNode(ConstantTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 86:
		//line front/mutan.y:303
		{
			yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
		}
	case 87:
		//line front/mutan.y:304
		{
			yyVAL.tnode = NewNode(BoolTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 88:
		//line front/mutan.y:305
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 89:
		//line front/mutan.y:306
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 90:
		//line front/mutan.y:310
		{
			yyVAL.tnode = NewNode(DerefPtrTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 91:
		//line front/mutan.y:314
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 92:
		//line front/mutan.y:315
		{
			yyVAL.tnode = NewNode(NilTy)
		}
	case 93:
		//line front/mutan.y:319
		{
			yyVAL.tnode = NewNode(IdentifierTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 94:
		//line front/mutan.y:323
		{
			yyVAL.tnode = NewNode(StringTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 95:
		//line front/mutan.y:327
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 96:
		//line front/mutan.y:331
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	}
	goto yystack /* stack new state and value */
}

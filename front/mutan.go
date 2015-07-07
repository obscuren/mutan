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
const MSG = 57349
const ADDR = 57350
const ORIGIN = 57351
const CALLER = 57352
const CALLVAL = 57353
const CALLDATALOAD = 57354
const CALLDATASIZE = 57355
const GASPRICE = 57356
const CALL = 57357
const CALLCODE = 57358
const SIZEOF = 57359
const EXIT = 57360
const CREATE = 57361
const BALANCE = 57362
const SHA3 = 57363
const DIFFICULTY = 57364
const PREVHASH = 57365
const TIMESTAMP = 57366
const BLOCKNUM = 57367
const COINBASE = 57368
const GAS = 57369
const ADDRESS = 57370
const BYTE = 57371
const PUSH = 57372
const POP = 57373
const TRANSACT = 57374
const STORE = 57375
const SUICIDE = 57376
const ASSIGN = 57377
const EQUAL = 57378
const END_STMT = 57379
const NIL = 57380
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
const CONST = 57396
const FUNC = 57397
const FUNC_CALL = 57398
const IMPORT = 57399
const DOT = 57400
const ARRAY = 57401
const COMMA = 57402
const QUOTE = 57403
const PRINT = 57404
const ID = 57405
const NUMBER = 57406
const INLINE_ASM = 57407
const OP = 57408
const DOP = 57409
const STR = 57410
const BOOLEAN = 57411
const CODE = 57412
const oper = 57413
const AND = 57414
const MUL = 57415

var yyToknames = []string{
	"BLOCK",
	"TX",
	"CONTRACT",
	"MSG",
	"ADDR",
	"ORIGIN",
	"CALLER",
	"CALLVAL",
	"CALLDATALOAD",
	"CALLDATASIZE",
	"GASPRICE",
	"CALL",
	"CALLCODE",
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
	"CONST",
	"FUNC",
	"FUNC_CALL",
	"IMPORT",
	"DOT",
	"ARRAY",
	"COMMA",
	"QUOTE",
	"PRINT",
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

//line front/mutan.y:353

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 10,
	60, 67,
	-2, 3,
}

const yyNprod = 102
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1191

var yyAct = []int{

	67, 32, 2, 19, 26, 231, 166, 13, 124, 264,
	60, 240, 162, 58, 86, 84, 61, 62, 106, 19,
	59, 128, 182, 64, 68, 69, 60, 4, 60, 239,
	60, 78, 61, 62, 61, 62, 61, 62, 229, 245,
	238, 85, 200, 207, 60, 137, 70, 35, 230, 21,
	61, 62, 271, 127, 129, 60, 115, 60, 83, 19,
	19, 61, 62, 61, 62, 82, 55, 185, 110, 111,
	112, 113, 143, 76, 19, 36, 228, 60, 270, 118,
	119, 120, 60, 61, 62, 260, 126, 108, 61, 62,
	133, 133, 133, 133, 19, 259, 123, 81, 133, 133,
	258, 121, 142, 140, 171, 167, 144, 237, 236, 235,
	102, 184, 163, 79, 179, 19, 101, 19, 168, 178,
	170, 138, 60, 80, 19, 177, 176, 100, 61, 62,
	173, 99, 201, 278, 202, 277, 272, 226, 131, 134,
	135, 136, 121, 72, 169, 224, 141, 222, 74, 221,
	220, 108, 219, 218, 217, 73, 216, 215, 75, 232,
	214, 183, 181, 180, 139, 130, 125, 114, 198, 19,
	75, 19, 196, 194, 193, 192, 191, 190, 133, 133,
	133, 189, 188, 187, 186, 105, 104, 211, 98, 97,
	212, 96, 213, 95, 94, 93, 19, 92, 19, 91,
	90, 89, 88, 87, 242, 19, 241, 199, 175, 172,
	197, 195, 114, 103, 234, 247, 165, 248, 246, 243,
	56, 116, 204, 223, 252, 225, 208, 209, 210, 117,
	206, 122, 233, 71, 63, 156, 244, 133, 133, 133,
	157, 147, 148, 146, 149, 150, 253, 22, 227, 255,
	23, 257, 254, 19, 19, 256, 19, 159, 19, 160,
	133, 133, 133, 266, 267, 268, 164, 269, 30, 6,
	19, 77, 133, 133, 161, 274, 275, 14, 57, 276,
	261, 49, 50, 51, 52, 249, 250, 251, 31, 5,
	158, 151, 38, 39, 42, 16, 41, 46, 47, 152,
	155, 154, 145, 12, 153, 3, 45, 43, 44, 40,
	1, 48, 0, 0, 11, 35, 53, 0, 15, 17,
	0, 18, 10, 273, 0, 0, 8, 34, 0, 37,
	24, 0, 7, 0, 9, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 11, 35, 53, 0, 15, 17,
	0, 18, 10, 265, 0, 0, 8, 34, 0, 37,
	24, 0, 7, 0, 9, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 11, 35, 53, 0, 15, 17,
	0, 18, 10, 263, 0, 0, 8, 34, 0, 37,
	24, 0, 7, 0, 9, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 11, 35, 53, 0, 15, 17,
	0, 18, 10, 262, 0, 0, 8, 34, 0, 37,
	24, 0, 7, 0, 9, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 11, 35, 53, 0, 15, 17,
	0, 18, 10, 205, 0, 0, 8, 34, 0, 37,
	24, 0, 7, 0, 9, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 11, 35, 53, 0, 15, 17,
	0, 18, 10, 203, 0, 0, 8, 34, 0, 37,
	24, 0, 7, 0, 9, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 11, 35, 53, 0, 15, 17,
	0, 18, 10, 107, 0, 0, 8, 34, 0, 37,
	24, 0, 7, 0, 9, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 11, 35, 53, 0, 15, 17,
	0, 18, 10, 0, 0, 0, 8, 34, 0, 37,
	24, 0, 7, 0, 9, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 0, 35, 53, 0, 15, 0,
	0, 0, 65, 0, 0, 0, 0, 34, 174, 37,
	24, 0, 0, 0, 0, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 0, 35, 53, 0, 15, 0,
	0, 0, 65, 109, 0, 0, 0, 34, 0, 37,
	24, 0, 0, 0, 0, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 0, 35, 53, 0, 15, 0,
	0, 0, 65, 0, 0, 0, 0, 34, 0, 37,
	24, 0, 0, 0, 0, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 0, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 0, 35, 53, 0, 0, 0,
	0, 0, 65, 0, 0, 0, 0, 34, 0, 37,
	0, 0, 0, 0, 0, 0, 0, 0, 36, 54,
	66, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 0, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 0, 35, 53, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 37,
	0, 0, 0, 0, 0, 0, 0, 0, 36, 54,
	132, 28, 0, 0, 0, 0, 29, 0, 0, 27,
	25,
}
var yyPact = []int{

	-1000, -1000, 767, -1000, -1000, -1000, -1000, 3, 175, 14,
	-1000, -1000, -1000, -56, 199, 1047, 1047, 1047, 977, 198,
	108, -1000, -1000, 1047, 50, 2, -1000, -5, -1000, -1000,
	-1000, -1000, -52, 1047, -1000, -1000, -54, 153, 152, 151,
	150, 149, 147, 145, 144, 143, 141, 139, 138, 73,
	69, 58, 52, 166, 136, 135, -47, -1000, 697, 907,
	1047, 1047, 1047, 1047, -56, -1000, 120, -1000, -56, 11,
	184, 1047, 1047, 977, 196, -1000, 115, 115, -56, -1000,
	-10, -43, -1000, -1000, -1000, -56, -7, 114, 1117, 1117,
	1117, 1117, -18, 977, 113, 1047, 1117, 9, 1047, 219,
	290, 207, 247, -58, 1047, -1000, 170, -1000, 45, -1000,
	-56, -56, -56, -56, 977, -1000, 977, -1000, -56, -56,
	56, 161, 1047, 837, -1000, -1000, -1000, -1000, 160, -1000,
	-1000, 66, 165, -1000, 65, 59, 54, 112, 111, -1000,
	-38, 110, 51, -1000, 16, -1000, 134, 133, 132, 131,
	127, -1000, 126, 125, 124, -1000, 123, 164, -1000, 122,
	163, 118, 159, -9, 81, -1000, -1000, -1000, 627, 185,
	557, 195, -1000, -56, -1000, -20, 1117, 1117, 1117, 1047,
	-1000, -1000, 1047, -1000, 1047, -1000, 109, 106, 105, 103,
	102, 101, 99, 98, 96, 977, 94, 977, 86, -1000,
	-1000, 23, -25, 116, 977, -1000, 1047, -1000, 49, 48,
	47, -11, -22, -40, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 158, -1000, 156, -1000, 174, -1000, 45,
	-24, -1000, 173, 172, -56, 1117, 1117, 1117, -1000, -1000,
	-1000, 189, -1000, -1000, -1000, 45, -1000, 1047, -1000, 40,
	35, 25, 977, 487, -1000, 417, -36, 347, 9, 9,
	9, -1000, -1000, -1000, -1000, -1000, 18, -8, 85, 277,
	9, 9, -1000, 116, 84, 82, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 310, 2, 305, 27, 303, 7, 49, 302, 300,
	291, 290, 289, 247, 288, 277, 277, 6, 1, 268,
	269, 5, 4, 266, 20, 0, 250, 8, 248,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 23, 23, 23, 28, 28, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 16, 16, 17, 17, 9,
	9, 9, 11, 11, 11, 11, 8, 8, 8, 8,
	8, 10, 10, 10, 12, 21, 21, 21, 20, 20,
	4, 4, 4, 4, 4, 4, 24, 24, 5, 5,
	5, 5, 5, 15, 15, 15, 6, 6, 6, 6,
	6, 6, 13, 13, 13, 13, 13, 7, 7, 7,
	7, 7, 7, 7, 7, 25, 22, 22, 18, 19,
	26, 27,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 1, 9, 4, 2,
	3, 1, 4, 5, 0, 1, 0, 3, 12, 12,
	10, 6, 4, 4, 3, 6, 4, 6, 4, 3,
	3, 3, 3, 4, 4, 3, 0, 1, 0, 3,
	4, 6, 3, 4, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 6, 4, 7, 0, 9, 5,
	1, 1, 1, 2, 2, 0, 3, 0, 3, 3,
	6, 3, 4, 2, 3, 5, 1, 1, 3, 3,
	3, 4, 2, 3, 3, 3, 2, 1, 1, 2,
	1, 4, 1, 1, 1, 2, 1, 1, 1, 3,
	1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -12, -20, 55, 49, 57,
	45, 37, -5, -6, -15, 41, 18, 42, 44, -25,
	63, -7, -13, -26, 53, 73, -22, 72, 64, 69,
	-19, -14, -18, 66, 50, 38, 61, 52, 15, 16,
	32, 19, 17, 30, 31, 29, 20, 21, 34, 4,
	5, 6, 7, 39, 62, 63, 45, -19, -2, -24,
	66, 72, 73, 35, -6, 45, 63, -25, -6, -6,
	-4, 35, 35, 47, 40, 50, -7, -13, -6, 63,
	73, 47, 63, 63, 67, -6, 68, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 50, 50, 58,
	58, 58, 58, 47, 50, 50, 65, 46, -4, 46,
	-6, -6, -6, -6, 47, 45, 37, 45, -6, -6,
	-6, -4, 35, -24, -27, 51, -27, 63, 64, 61,
	51, -7, 63, -18, -7, -7, -7, 63, -4, 51,
	-6, -7, -22, 63, -6, -8, 24, 22, 23, 25,
	26, -10, 9, 14, 11, -9, 28, 33, -11, 10,
	12, 27, 70, -6, -23, 46, -17, 60, -2, -4,
	-2, 48, 48, -6, 51, 48, 60, 60, 60, 60,
	51, 51, 60, 51, 60, 51, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 47, 50, 47, 50, 48,
	51, 51, 53, 46, 37, 46, 35, 63, -7, -7,
	-7, -6, -6, -6, 51, 51, 51, 51, 51, 51,
	51, 51, 51, -4, 51, -4, 51, -28, 53, 63,
	73, -21, 43, -4, -6, 60, 60, 60, 51, 51,
	51, 48, 48, 45, -17, 63, 45, 42, 45, -7,
	-7, -7, 35, -2, -17, -2, -6, -2, 60, 60,
	60, -4, 46, 46, 45, 46, -22, -22, -22, -2,
	60, 60, 51, 46, -22, -22, -21, 51, 51,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 5, 6, 0, 0, 0,
	-2, 11, 60, 61, 62, 0, 0, 0, 65, 88,
	98, 76, 77, 0, 0, 0, 87, 0, 90, 92,
	93, 94, 96, 0, 100, 97, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 9, 0, 65,
	0, 0, 0, 0, 63, 67, 98, 88, 64, 0,
	0, 0, 0, 65, 0, 67, 76, 77, 0, 73,
	0, 0, 95, 89, 82, 86, 0, 0, 0, 0,
	0, 0, 0, 65, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 14, 0, 10, 38, 80,
	83, 84, 85, 71, 65, 3, 65, 3, 68, 69,
	0, 0, 0, 65, 78, 101, 79, 74, 0, 99,
	17, 0, 98, 96, 0, 0, 0, 0, 0, 24,
	0, 0, 0, 98, 0, 29, 0, 0, 0, 0,
	0, 30, 0, 0, 0, 31, 0, 0, 32, 0,
	44, 0, 0, 0, 0, 8, 66, 37, 0, 0,
	0, 0, 91, 72, 81, 0, 0, 0, 0, 0,
	22, 23, 0, 26, 0, 28, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 65, 0, 33,
	34, 16, 0, 57, 0, 59, 0, 75, 0, 0,
	0, 0, 0, 0, 46, 47, 48, 49, 50, 51,
	52, 53, 39, 0, 42, 0, 45, 0, 15, 38,
	0, 54, 0, 0, 70, 0, 0, 0, 21, 25,
	27, 40, 43, 3, 12, 38, 3, 0, 3, 0,
	0, 0, 65, 0, 13, 0, 0, 0, 0, 0,
	0, 41, 7, 55, 3, 58, 0, 0, 0, 0,
	0, 0, 20, 57, 0, 0, 56, 18, 19,
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
	72, 73,
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
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-5].tnode, false)
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
			yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 20:
		//line front/mutan.y:127
		{
			yyVAL.tnode = NewNode(TransactTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 21:
		//line front/mutan.y:130
		{
			yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 22:
		//line front/mutan.y:131
		{
			yyVAL.tnode = NewNode(SizeofTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 23:
		//line front/mutan.y:132
		{
			yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode)
		}
	case 24:
		//line front/mutan.y:133
		{
			yyVAL.tnode = NewNode(PopTy)
		}
	case 25:
		//line front/mutan.y:134
		{
			yyVAL.tnode = NewNode(ByteTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 26:
		//line front/mutan.y:135
		{
			yyVAL.tnode = NewNode(BalanceTy, yyS[yypt-1].tnode)
		}
	case 27:
		//line front/mutan.y:136
		{
			yyVAL.tnode = NewNode(Sha3Ty, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 28:
		//line front/mutan.y:137
		{
			yyVAL.tnode = NewNode(SuicideTy, yyS[yypt-1].tnode)
		}
	case 29:
		//line front/mutan.y:138
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 30:
		//line front/mutan.y:139
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 31:
		//line front/mutan.y:140
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 32:
		//line front/mutan.y:141
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 33:
		//line front/mutan.y:142
		{
			yyVAL.tnode = NewNode(LambdaTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 34:
		//line front/mutan.y:143
		{
			yyVAL.tnode = NewNode(PrintTy, yyS[yypt-1].tnode)
		}
	case 35:
		//line front/mutan.y:147
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 36:
		//line front/mutan.y:148
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 37:
		//line front/mutan.y:152
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 38:
		//line front/mutan.y:153
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 39:
		//line front/mutan.y:158
		{
			yyVAL.tnode = NewNode(AddressTy)
		}
	case 40:
		//line front/mutan.y:159
		{
			yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode)
		}
	case 41:
		//line front/mutan.y:161
		{
			node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 42:
		//line front/mutan.y:168
		{
			yyVAL.tnode = NewNode(CallerTy)
		}
	case 43:
		//line front/mutan.y:169
		{
			yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode)
		}
	case 44:
		//line front/mutan.y:170
		{
			yyVAL.tnode = NewNode(CallDataSizeTy)
		}
	case 45:
		//line front/mutan.y:171
		{
			yyVAL.tnode = NewNode(GasTy)
		}
	case 46:
		//line front/mutan.y:176
		{
			yyVAL.tnode = NewNode(TimestampTy)
		}
	case 47:
		//line front/mutan.y:177
		{
			yyVAL.tnode = NewNode(DiffTy)
		}
	case 48:
		//line front/mutan.y:178
		{
			yyVAL.tnode = NewNode(PrevHashTy)
		}
	case 49:
		//line front/mutan.y:179
		{
			yyVAL.tnode = NewNode(BlockNumTy)
		}
	case 50:
		//line front/mutan.y:180
		{
			yyVAL.tnode = NewNode(CoinbaseTy)
		}
	case 51:
		//line front/mutan.y:184
		{
			yyVAL.tnode = NewNode(OriginTy)
		}
	case 52:
		//line front/mutan.y:185
		{
			yyVAL.tnode = NewNode(GasPriceTy)
		}
	case 53:
		//line front/mutan.y:186
		{
			yyVAL.tnode = NewNode(CallValTy)
		}
	case 54:
		//line front/mutan.y:191
		{
			if yyS[yypt-0].tnode == nil {
				yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
			} else {
				yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			}
		}
	case 55:
		//line front/mutan.y:201
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 56:
		//line front/mutan.y:205
		{
			if yyS[yypt-0].tnode == nil {
				yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
			} else {
				yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			}
		}
	case 57:
		//line front/mutan.y:212
		{
			yyVAL.tnode = nil
		}
	case 58:
		//line front/mutan.y:217
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 59:
		//line front/mutan.y:221
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 60:
		//line front/mutan.y:227
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 61:
		//line front/mutan.y:228
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 62:
		//line front/mutan.y:229
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 63:
		//line front/mutan.y:230
		{
			yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode)
		}
	case 64:
		//line front/mutan.y:231
		{
			yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode)
		}
	case 65:
		//line front/mutan.y:232
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 66:
		//line front/mutan.y:236
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 67:
		//line front/mutan.y:237
		{
			yyVAL.tnode = nil
		}
	case 68:
		//line front/mutan.y:242
		{
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode)
		}
	case 69:
		//line front/mutan.y:246
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-2].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 70:
		//line front/mutan.y:252
		{
			yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-5].str
		}
	case 71:
		//line front/mutan.y:256
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-2].tnode.Constant
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		}
	case 72:
		//line front/mutan.y:262
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
			varNode := NewNode(NewVarTy)
			varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		}
	case 73:
		//line front/mutan.y:272
		{
			yyVAL.tnode = NewNode(NewVarTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 74:
		//line front/mutan.y:277
		{
			yyVAL.tnode = NewNode(NewVarTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
			yyVAL.tnode.Ptr = true
		}
	case 75:
		//line front/mutan.y:283
		{
			yyVAL.tnode = NewNode(NewArrayTy)
			yyVAL.tnode.Size = yyS[yypt-2].str
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 76:
		//line front/mutan.y:291
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 77:
		//line front/mutan.y:292
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 78:
		//line front/mutan.y:293
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 79:
		//line front/mutan.y:294
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 80:
		//line front/mutan.y:296
		{
			yyVAL.tnode = NewNode(InitListTy)
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
		}
	case 81:
		//line front/mutan.y:301
		{
			yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
		}
	case 82:
		//line front/mutan.y:310
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 83:
		//line front/mutan.y:312
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 84:
		//line front/mutan.y:313
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 85:
		//line front/mutan.y:314
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 86:
		//line front/mutan.y:315
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 87:
		//line front/mutan.y:319
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 88:
		//line front/mutan.y:320
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 89:
		//line front/mutan.y:321
		{
			yyVAL.tnode = NewNode(RefTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 90:
		//line front/mutan.y:322
		{
			yyVAL.tnode = NewNode(ConstantTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 91:
		//line front/mutan.y:323
		{
			yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
		}
	case 92:
		//line front/mutan.y:324
		{
			yyVAL.tnode = NewNode(BoolTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 93:
		//line front/mutan.y:325
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 94:
		//line front/mutan.y:326
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 95:
		//line front/mutan.y:330
		{
			yyVAL.tnode = NewNode(DerefPtrTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 96:
		//line front/mutan.y:334
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 97:
		//line front/mutan.y:335
		{
			yyVAL.tnode = NewNode(NilTy)
		}
	case 98:
		//line front/mutan.y:339
		{
			yyVAL.tnode = NewNode(IdentifierTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 99:
		//line front/mutan.y:343
		{
			yyVAL.tnode = NewNode(StringTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 100:
		//line front/mutan.y:347
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 101:
		//line front/mutan.y:351
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	}
	goto yystack /* stack new state and value */
}

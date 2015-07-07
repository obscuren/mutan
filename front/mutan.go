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
	-1, 9,
	60, 67,
	-2, 3,
}

const yyNprod = 102
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1191

var yyAct = []int{

	67, 32, 2, 19, 26, 231, 165, 12, 264, 124,
	59, 240, 57, 162, 86, 84, 60, 61, 229, 19,
	58, 112, 239, 64, 68, 69, 59, 4, 230, 59,
	238, 78, 60, 61, 128, 60, 61, 59, 245, 200,
	81, 85, 185, 60, 61, 59, 70, 35, 207, 21,
	137, 60, 61, 127, 59, 83, 79, 59, 19, 19,
	60, 61, 82, 60, 61, 55, 80, 109, 110, 111,
	182, 113, 143, 76, 19, 129, 59, 36, 271, 118,
	119, 120, 60, 61, 270, 260, 107, 126, 259, 171,
	133, 133, 133, 133, 19, 115, 123, 258, 133, 133,
	166, 121, 142, 140, 237, 236, 144, 59, 235, 184,
	102, 179, 163, 60, 61, 19, 59, 19, 168, 178,
	170, 138, 60, 61, 19, 177, 176, 101, 100, 99,
	173, 201, 228, 202, 278, 277, 272, 226, 131, 134,
	135, 136, 121, 72, 169, 224, 141, 222, 74, 221,
	220, 107, 219, 218, 217, 73, 216, 215, 75, 232,
	214, 183, 181, 180, 139, 130, 125, 114, 198, 19,
	75, 19, 196, 194, 193, 192, 191, 190, 133, 133,
	133, 189, 188, 187, 186, 105, 104, 211, 98, 97,
	212, 96, 213, 95, 94, 93, 19, 92, 19, 91,
	90, 89, 88, 87, 242, 19, 241, 199, 175, 172,
	197, 195, 114, 103, 234, 247, 167, 248, 246, 243,
	62, 116, 204, 223, 252, 225, 208, 209, 210, 117,
	206, 122, 233, 71, 63, 156, 244, 133, 133, 133,
	157, 147, 148, 146, 149, 150, 253, 22, 227, 255,
	23, 257, 254, 19, 19, 256, 19, 159, 19, 160,
	133, 133, 133, 266, 267, 268, 164, 269, 30, 6,
	19, 77, 133, 133, 161, 274, 275, 56, 14, 276,
	261, 49, 50, 51, 52, 249, 250, 251, 31, 5,
	158, 151, 38, 39, 42, 16, 41, 46, 47, 152,
	155, 154, 145, 11, 153, 3, 45, 43, 44, 40,
	1, 48, 0, 0, 10, 35, 53, 0, 15, 17,
	0, 18, 9, 273, 0, 0, 13, 34, 0, 37,
	24, 0, 7, 0, 8, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 10, 35, 53, 0, 15, 17,
	0, 18, 9, 265, 0, 0, 13, 34, 0, 37,
	24, 0, 7, 0, 8, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 10, 35, 53, 0, 15, 17,
	0, 18, 9, 263, 0, 0, 13, 34, 0, 37,
	24, 0, 7, 0, 8, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 10, 35, 53, 0, 15, 17,
	0, 18, 9, 262, 0, 0, 13, 34, 0, 37,
	24, 0, 7, 0, 8, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 10, 35, 53, 0, 15, 17,
	0, 18, 9, 205, 0, 0, 13, 34, 0, 37,
	24, 0, 7, 0, 8, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 10, 35, 53, 0, 15, 17,
	0, 18, 9, 203, 0, 0, 13, 34, 0, 37,
	24, 0, 7, 0, 8, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 10, 35, 53, 0, 15, 17,
	0, 18, 9, 106, 0, 0, 13, 34, 0, 37,
	24, 0, 7, 0, 8, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 10, 35, 53, 0, 15, 17,
	0, 18, 9, 0, 0, 0, 13, 34, 0, 37,
	24, 0, 7, 0, 8, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 0, 35, 53, 0, 15, 0,
	0, 0, 65, 0, 0, 0, 13, 34, 174, 37,
	24, 0, 0, 0, 0, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 0, 35, 53, 0, 15, 0,
	0, 0, 65, 108, 0, 0, 13, 34, 0, 37,
	24, 0, 0, 0, 0, 0, 0, 0, 36, 54,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 49, 50, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 42, 16, 41, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 45, 43, 44, 40,
	0, 48, 0, 0, 0, 35, 53, 0, 15, 0,
	0, 0, 65, 0, 0, 0, 13, 34, 0, 37,
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

	-1000, -1000, 767, -1000, -1000, -1000, -1000, 2, 16, -1000,
	-1000, -1000, -56, 175, 199, 1047, 1047, 1047, 977, 198,
	108, -1000, -1000, 1047, -7, -1, -1000, -8, -1000, -1000,
	-1000, -1000, -52, 1047, -1000, -1000, -54, 153, 152, 151,
	150, 149, 147, 145, 144, 143, 141, 139, 138, 71,
	70, 69, 52, 166, 136, 135, -1000, 697, 907, 1047,
	1047, 1047, -44, 1047, -56, -1000, 120, -1000, -56, 50,
	184, 1047, 1047, 977, 196, -1000, 115, 115, -56, -1000,
	-10, -30, -1000, -1000, -1000, -56, 14, 114, 1117, 1117,
	1117, 1117, -13, 977, 113, 1047, 1117, 9, 1047, 219,
	290, 207, 247, -57, 1047, -1000, -1000, 40, -1000, -56,
	-56, -56, 170, -56, 977, -1000, 977, -1000, -56, -56,
	41, 161, 1047, 837, -1000, -1000, -1000, -1000, 160, -1000,
	-1000, 66, 165, -1000, 65, 59, 51, 112, 111, -1000,
	10, 110, 49, -1000, -9, -1000, 134, 133, 132, 131,
	127, -1000, 126, 125, 124, -1000, 123, 164, -1000, 122,
	163, 118, 159, -12, 80, -1000, -1000, -1000, 627, 185,
	557, 195, -1000, -56, -1000, -15, 1117, 1117, 1117, 1047,
	-1000, -1000, 1047, -1000, 1047, -1000, 109, 106, 105, 103,
	102, 101, 99, 98, 96, 977, 94, 977, 86, -1000,
	-1000, 79, -45, 116, 977, -1000, 1047, -1000, 48, 45,
	44, -21, -29, -40, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 158, -1000, 156, -1000, 174, -1000, 40,
	-25, -1000, 173, 172, -56, 1117, 1117, 1117, -1000, -1000,
	-1000, 189, -1000, -1000, -1000, 40, -1000, 1047, -1000, 37,
	28, 25, 977, 487, -1000, 417, -37, 347, 9, 9,
	9, -1000, -1000, -1000, -1000, -1000, 24, 18, 85, 277,
	9, 9, -1000, 116, 84, 83, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 310, 2, 305, 27, 303, 7, 49, 302, 300,
	291, 290, 289, 247, 288, 278, 278, 6, 1, 268,
	269, 5, 4, 266, 20, 0, 250, 9, 248,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 23, 23, 23, 28, 28, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 16, 16, 17, 17, 9, 9,
	9, 11, 11, 11, 11, 8, 8, 8, 8, 8,
	10, 10, 10, 12, 21, 21, 21, 20, 20, 4,
	4, 4, 4, 4, 4, 4, 24, 24, 5, 5,
	5, 5, 5, 15, 15, 15, 6, 6, 6, 6,
	6, 6, 13, 13, 13, 13, 13, 7, 7, 7,
	7, 7, 7, 7, 7, 25, 22, 22, 18, 19,
	26, 27,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 1, 9, 2, 3,
	1, 4, 5, 0, 1, 0, 3, 12, 12, 10,
	6, 4, 4, 3, 6, 4, 6, 4, 3, 3,
	3, 3, 4, 4, 3, 0, 1, 0, 3, 4,
	6, 3, 4, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 6, 4, 7, 0, 9, 5, 1,
	1, 4, 1, 2, 2, 0, 3, 0, 3, 3,
	6, 3, 4, 2, 3, 5, 1, 1, 3, 3,
	3, 4, 2, 3, 3, 3, 2, 1, 1, 2,
	1, 4, 1, 1, 1, 2, 1, 1, 1, 3,
	1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -12, -20, 55, 57, 45,
	37, -5, -6, 49, -15, 41, 18, 42, 44, -25,
	63, -7, -13, -26, 53, 73, -22, 72, 64, 69,
	-19, -14, -18, 66, 50, 38, 61, 52, 15, 16,
	32, 19, 17, 30, 31, 29, 20, 21, 34, 4,
	5, 6, 7, 39, 62, 63, -19, -2, -24, 66,
	72, 73, 45, 35, -6, 45, 63, -25, -6, -6,
	-4, 35, 35, 47, 40, 50, -7, -13, -6, 63,
	73, 47, 63, 63, 67, -6, 68, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 50, 50, 58,
	58, 58, 58, 47, 50, 50, 46, -4, 46, -6,
	-6, -6, 65, -6, 47, 45, 37, 45, -6, -6,
	-6, -4, 35, -24, -27, 51, -27, 63, 64, 61,
	51, -7, 63, -18, -7, -7, -7, 63, -4, 51,
	-6, -7, -22, 63, -6, -8, 24, 22, 23, 25,
	26, -10, 9, 14, 11, -9, 28, 33, -11, 10,
	12, 27, 70, -6, -23, -17, 60, 46, -2, -4,
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

	3, -2, 1, 2, 4, 5, 6, 0, 0, -2,
	10, 59, 60, 0, 62, 0, 0, 0, 65, 88,
	98, 76, 77, 0, 0, 0, 87, 0, 90, 92,
	93, 94, 96, 0, 100, 97, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 8, 0, 65, 0,
	0, 0, 0, 0, 63, 67, 98, 88, 64, 0,
	0, 0, 0, 65, 0, 67, 76, 77, 0, 73,
	0, 0, 95, 89, 82, 86, 0, 0, 0, 0,
	0, 0, 0, 65, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 13, 9, 37, 80, 83,
	84, 85, 0, 71, 65, 3, 65, 3, 68, 69,
	0, 0, 0, 65, 78, 101, 79, 74, 0, 99,
	16, 0, 98, 96, 0, 0, 0, 0, 0, 23,
	0, 0, 0, 98, 0, 28, 0, 0, 0, 0,
	0, 29, 0, 0, 0, 30, 0, 0, 31, 0,
	43, 0, 0, 0, 0, 66, 36, 61, 0, 0,
	0, 0, 91, 72, 81, 0, 0, 0, 0, 0,
	21, 22, 0, 25, 0, 27, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 65, 0, 32,
	33, 15, 0, 56, 0, 58, 0, 75, 0, 0,
	0, 0, 0, 0, 45, 46, 47, 48, 49, 50,
	51, 52, 38, 0, 41, 0, 44, 0, 14, 37,
	0, 53, 0, 0, 70, 0, 0, 0, 20, 24,
	26, 39, 42, 3, 11, 37, 3, 0, 3, 0,
	0, 0, 65, 0, 12, 0, 0, 0, 0, 0,
	0, 40, 7, 54, 3, 57, 0, 0, 0, 0,
	0, 0, 19, 56, 0, 0, 55, 17, 18,
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
			yyVAL.tnode = NewNode(ImportTy)
			yyVAL.tnode.Constant = yyS[yypt-0].tnode.Constant
		}
	case 9:
		//line front/mutan.y:97
		{
			yyVAL.tnode = NewNode(ScopeTy, yyS[yypt-1].tnode)
		}
	case 10:
		//line front/mutan.y:98
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 11:
		//line front/mutan.y:102
		{
			yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 12:
		//line front/mutan.y:103
		{
			yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-4].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
			yyVAL.tnode.Ptr = true
		}
	case 13:
		//line front/mutan.y:104
		{
			yyVAL.tnode = nil
		}
	case 14:
		//line front/mutan.y:109
		{
			yyVAL.check = true
		}
	case 15:
		//line front/mutan.y:110
		{
			yyVAL.check = false
		}
	case 16:
		//line front/mutan.y:115
		{
			yyVAL.tnode = NewNode(StopTy)
		}
	case 17:
		//line front/mutan.y:118
		{
			yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 18:
		//line front/mutan.y:122
		{
			yyVAL.tnode = NewNode(CallCodeTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 19:
		//line front/mutan.y:126
		{
			yyVAL.tnode = NewNode(TransactTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 20:
		//line front/mutan.y:129
		{
			yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 21:
		//line front/mutan.y:130
		{
			yyVAL.tnode = NewNode(SizeofTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 22:
		//line front/mutan.y:131
		{
			yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode)
		}
	case 23:
		//line front/mutan.y:132
		{
			yyVAL.tnode = NewNode(PopTy)
		}
	case 24:
		//line front/mutan.y:133
		{
			yyVAL.tnode = NewNode(ByteTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 25:
		//line front/mutan.y:134
		{
			yyVAL.tnode = NewNode(BalanceTy, yyS[yypt-1].tnode)
		}
	case 26:
		//line front/mutan.y:135
		{
			yyVAL.tnode = NewNode(Sha3Ty, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 27:
		//line front/mutan.y:136
		{
			yyVAL.tnode = NewNode(SuicideTy, yyS[yypt-1].tnode)
		}
	case 28:
		//line front/mutan.y:137
		{
			yyVAL.tnode = yyS[yypt-0].tnode
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
			yyVAL.tnode = NewNode(LambdaTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 33:
		//line front/mutan.y:142
		{
			yyVAL.tnode = NewNode(PrintTy, yyS[yypt-1].tnode)
		}
	case 34:
		//line front/mutan.y:146
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 35:
		//line front/mutan.y:147
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 36:
		//line front/mutan.y:151
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 37:
		//line front/mutan.y:152
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 38:
		//line front/mutan.y:157
		{
			yyVAL.tnode = NewNode(AddressTy)
		}
	case 39:
		//line front/mutan.y:158
		{
			yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode)
		}
	case 40:
		//line front/mutan.y:160
		{
			node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 41:
		//line front/mutan.y:167
		{
			yyVAL.tnode = NewNode(CallerTy)
		}
	case 42:
		//line front/mutan.y:168
		{
			yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode)
		}
	case 43:
		//line front/mutan.y:169
		{
			yyVAL.tnode = NewNode(CallDataSizeTy)
		}
	case 44:
		//line front/mutan.y:170
		{
			yyVAL.tnode = NewNode(GasTy)
		}
	case 45:
		//line front/mutan.y:175
		{
			yyVAL.tnode = NewNode(TimestampTy)
		}
	case 46:
		//line front/mutan.y:176
		{
			yyVAL.tnode = NewNode(DiffTy)
		}
	case 47:
		//line front/mutan.y:177
		{
			yyVAL.tnode = NewNode(PrevHashTy)
		}
	case 48:
		//line front/mutan.y:178
		{
			yyVAL.tnode = NewNode(BlockNumTy)
		}
	case 49:
		//line front/mutan.y:179
		{
			yyVAL.tnode = NewNode(CoinbaseTy)
		}
	case 50:
		//line front/mutan.y:183
		{
			yyVAL.tnode = NewNode(OriginTy)
		}
	case 51:
		//line front/mutan.y:184
		{
			yyVAL.tnode = NewNode(GasPriceTy)
		}
	case 52:
		//line front/mutan.y:185
		{
			yyVAL.tnode = NewNode(CallValTy)
		}
	case 53:
		//line front/mutan.y:190
		{
			if yyS[yypt-0].tnode == nil {
				yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
			} else {
				yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			}
		}
	case 54:
		//line front/mutan.y:200
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 55:
		//line front/mutan.y:204
		{
			if yyS[yypt-0].tnode == nil {
				yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
			} else {
				yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			}
		}
	case 56:
		//line front/mutan.y:211
		{
			yyVAL.tnode = nil
		}
	case 57:
		//line front/mutan.y:216
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 58:
		//line front/mutan.y:220
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 59:
		//line front/mutan.y:226
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 60:
		//line front/mutan.y:227
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 61:
		//line front/mutan.y:228
		{
			yyVAL.tnode = NewNode(InlineAsmTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
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

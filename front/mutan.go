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
const PRINT = 57403
const ID = 57404
const NUMBER = 57405
const INLINE_ASM = 57406
const OP = 57407
const DOP = 57408
const STR = 57409
const BOOLEAN = 57410
const CODE = 57411
const oper = 57412
const AND = 57413
const MUL = 57414

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

//line front/mutan.y:349

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 10,
	59, 66,
	-2, 3,
}

const yyNprod = 101
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1172

var yyAct = []int{

	66, 32, 2, 19, 26, 226, 122, 256, 13, 234,
	58, 159, 163, 57, 59, 85, 83, 233, 104, 19,
	60, 61, 126, 59, 63, 67, 68, 59, 4, 60,
	61, 59, 77, 60, 61, 232, 224, 60, 61, 178,
	196, 35, 84, 181, 239, 59, 225, 69, 203, 59,
	21, 60, 61, 134, 59, 60, 61, 59, 19, 19,
	60, 61, 125, 60, 61, 82, 140, 81, 108, 109,
	110, 111, 54, 19, 75, 127, 36, 261, 252, 116,
	117, 118, 251, 124, 164, 121, 168, 106, 231, 131,
	131, 131, 19, 230, 180, 175, 131, 131, 174, 113,
	139, 119, 137, 59, 173, 141, 80, 100, 99, 60,
	61, 160, 98, 19, 97, 19, 165, 223, 167, 59,
	135, 78, 19, 266, 262, 60, 61, 221, 197, 170,
	198, 79, 219, 217, 216, 215, 71, 214, 129, 132,
	133, 119, 73, 166, 213, 138, 212, 211, 210, 72,
	106, 209, 74, 227, 179, 177, 176, 136, 128, 123,
	112, 194, 192, 74, 190, 189, 19, 188, 19, 187,
	186, 185, 184, 183, 182, 131, 131, 103, 102, 96,
	95, 94, 93, 92, 206, 91, 90, 207, 89, 208,
	88, 87, 19, 86, 19, 236, 235, 195, 172, 169,
	193, 19, 191, 112, 101, 241, 162, 114, 240, 242,
	237, 229, 55, 200, 245, 202, 115, 120, 70, 62,
	218, 153, 220, 222, 204, 205, 154, 23, 156, 228,
	157, 22, 131, 131, 149, 161, 151, 238, 30, 150,
	246, 6, 14, 248, 158, 250, 19, 19, 56, 19,
	249, 19, 247, 131, 131, 76, 258, 259, 31, 260,
	5, 19, 155, 131, 148, 152, 264, 142, 12, 265,
	3, 1, 0, 0, 253, 48, 49, 50, 51, 0,
	0, 243, 244, 0, 0, 0, 38, 41, 16, 40,
	45, 46, 0, 144, 145, 143, 146, 147, 0, 44,
	42, 43, 39, 0, 47, 0, 0, 11, 35, 0,
	52, 0, 15, 17, 0, 18, 10, 263, 0, 0,
	8, 34, 0, 37, 24, 7, 0, 9, 0, 0,
	0, 36, 53, 20, 28, 0, 33, 0, 0, 29,
	0, 0, 27, 25, 48, 49, 50, 51, 0, 0,
	0, 0, 0, 0, 0, 38, 41, 16, 40, 45,
	46, 0, 0, 0, 0, 0, 0, 0, 44, 42,
	43, 39, 0, 47, 0, 0, 11, 35, 0, 52,
	0, 15, 17, 0, 18, 10, 257, 0, 0, 8,
	34, 0, 37, 24, 7, 0, 9, 0, 0, 0,
	36, 53, 20, 28, 0, 33, 0, 0, 29, 0,
	0, 27, 25, 48, 49, 50, 51, 0, 0, 0,
	0, 0, 0, 0, 38, 41, 16, 40, 45, 46,
	0, 0, 0, 0, 0, 0, 0, 44, 42, 43,
	39, 0, 47, 0, 0, 11, 35, 0, 52, 0,
	15, 17, 0, 18, 10, 255, 0, 0, 8, 34,
	0, 37, 24, 7, 0, 9, 0, 0, 0, 36,
	53, 20, 28, 0, 33, 0, 0, 29, 0, 0,
	27, 25, 48, 49, 50, 51, 0, 0, 0, 0,
	0, 0, 0, 38, 41, 16, 40, 45, 46, 0,
	0, 0, 0, 0, 0, 0, 44, 42, 43, 39,
	0, 47, 0, 0, 11, 35, 0, 52, 0, 15,
	17, 0, 18, 10, 254, 0, 0, 8, 34, 0,
	37, 24, 7, 0, 9, 0, 0, 0, 36, 53,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 48, 49, 50, 51, 0, 0, 0, 0, 0,
	0, 0, 38, 41, 16, 40, 45, 46, 0, 0,
	0, 0, 0, 0, 0, 44, 42, 43, 39, 0,
	47, 0, 0, 11, 35, 0, 52, 0, 15, 17,
	0, 18, 10, 201, 0, 0, 8, 34, 0, 37,
	24, 7, 0, 9, 0, 0, 0, 36, 53, 20,
	28, 0, 33, 0, 0, 29, 0, 0, 27, 25,
	48, 49, 50, 51, 0, 0, 0, 0, 0, 0,
	0, 38, 41, 16, 40, 45, 46, 0, 0, 0,
	0, 0, 0, 0, 44, 42, 43, 39, 0, 47,
	0, 0, 11, 35, 0, 52, 0, 15, 17, 0,
	18, 10, 199, 0, 0, 8, 34, 0, 37, 24,
	7, 0, 9, 0, 0, 0, 36, 53, 20, 28,
	0, 33, 0, 0, 29, 0, 0, 27, 25, 48,
	49, 50, 51, 0, 0, 0, 0, 0, 0, 0,
	38, 41, 16, 40, 45, 46, 0, 0, 0, 0,
	0, 0, 0, 44, 42, 43, 39, 0, 47, 0,
	0, 11, 35, 0, 52, 0, 15, 17, 0, 18,
	10, 105, 0, 0, 8, 34, 0, 37, 24, 7,
	0, 9, 0, 0, 0, 36, 53, 20, 28, 0,
	33, 0, 0, 29, 0, 0, 27, 25, 48, 49,
	50, 51, 0, 0, 0, 0, 0, 0, 0, 38,
	41, 16, 40, 45, 46, 0, 0, 0, 0, 0,
	0, 0, 44, 42, 43, 39, 0, 47, 0, 0,
	11, 35, 0, 52, 0, 15, 17, 0, 18, 10,
	0, 0, 0, 8, 34, 0, 37, 24, 7, 0,
	9, 0, 0, 0, 36, 53, 20, 28, 0, 33,
	0, 0, 29, 0, 0, 27, 25, 48, 49, 50,
	51, 0, 0, 0, 0, 0, 0, 0, 38, 41,
	16, 40, 45, 46, 0, 0, 0, 0, 0, 0,
	0, 44, 42, 43, 39, 0, 47, 0, 0, 0,
	35, 0, 52, 0, 15, 0, 0, 0, 64, 0,
	0, 0, 0, 34, 171, 37, 24, 0, 0, 0,
	0, 0, 0, 36, 53, 20, 28, 0, 33, 0,
	0, 29, 0, 0, 27, 25, 48, 49, 50, 51,
	0, 0, 0, 0, 0, 0, 0, 38, 41, 16,
	40, 45, 46, 0, 0, 0, 0, 0, 0, 0,
	44, 42, 43, 39, 0, 47, 0, 0, 0, 35,
	0, 52, 0, 15, 0, 0, 0, 64, 107, 0,
	0, 0, 34, 0, 37, 24, 0, 0, 0, 0,
	0, 0, 36, 53, 20, 28, 0, 33, 0, 0,
	29, 0, 0, 27, 25, 48, 49, 50, 51, 0,
	0, 0, 0, 0, 0, 0, 38, 41, 16, 40,
	45, 46, 0, 0, 0, 0, 0, 0, 0, 44,
	42, 43, 39, 0, 47, 0, 0, 0, 35, 0,
	52, 0, 15, 0, 0, 0, 64, 0, 0, 0,
	0, 34, 0, 37, 24, 0, 0, 0, 0, 0,
	0, 36, 53, 20, 28, 0, 33, 0, 0, 29,
	0, 0, 27, 25, 48, 49, 50, 51, 0, 0,
	0, 0, 0, 0, 0, 38, 41, 0, 40, 45,
	46, 0, 0, 0, 0, 0, 0, 0, 44, 42,
	43, 39, 0, 47, 0, 0, 0, 35, 0, 52,
	0, 0, 0, 0, 0, 64, 0, 0, 0, 0,
	34, 0, 37, 0, 0, 0, 0, 0, 0, 0,
	36, 53, 65, 28, 0, 33, 0, 0, 29, 0,
	0, 27, 25, 48, 49, 50, 51, 0, 0, 0,
	0, 0, 0, 0, 38, 41, 0, 40, 45, 46,
	0, 0, 0, 0, 0, 0, 0, 44, 42, 43,
	39, 0, 47, 0, 0, 0, 35, 0, 52, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 37, 0, 0, 0, 0, 0, 0, 0, 36,
	53, 130, 28, 0, 0, 0, 0, 29, 0, 0,
	27, 25,
}
var yyPact = []int{

	-1000, -1000, 754, -1000, -1000, -1000, -1000, 10, 167, 16,
	-1000, -1000, -1000, -51, 185, 1030, 1030, 1030, 961, 184,
	102, -1000, -1000, 1030, 59, 5, -1000, 3, -1000, -1000,
	-1000, -1000, -50, 1030, -1000, -1000, -52, 143, 141, 140,
	138, 136, 135, 133, 132, 131, 130, 129, 57, 55,
	51, 50, 157, 128, 127, -46, -1000, 685, 892, 1030,
	1030, 1030, 1030, -51, -1000, 113, -1000, -51, 54, 171,
	1030, 1030, 961, 183, -1000, 108, 108, -51, -1000, 0,
	-41, -1000, -1000, -1000, -51, 15, 107, 1099, 1099, 1099,
	-9, 961, 106, 1030, 1099, 4, 1030, 272, 225, 194,
	218, -58, 1030, -1000, 160, -1000, 25, -1000, -51, -51,
	-51, -51, 961, -1000, 961, -1000, -51, -51, 38, 151,
	1030, 823, -1000, -1000, -1000, -1000, 150, -1000, -1000, 45,
	156, -1000, 39, 36, 105, 104, -1000, -20, 103, 35,
	-1000, -8, -1000, 124, 123, 122, 121, 120, -1000, 119,
	117, 115, -1000, 114, 155, -1000, 112, 153, 111, 149,
	-11, 77, -1000, -1000, -1000, 616, 177, 547, 181, -1000,
	-51, -1000, -14, 1099, 1099, 1030, -1000, -1000, 1030, -1000,
	1030, -1000, 100, 97, 96, 95, 93, 86, 84, 83,
	82, 961, 81, 961, 76, -1000, -1000, 64, -26, 110,
	961, -1000, 1030, -1000, 34, 29, -16, -34, -42, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 148, -1000,
	147, -1000, 165, -1000, 25, -18, -1000, 163, 164, -51,
	1099, 1099, -1000, -1000, -1000, 180, -1000, -1000, -1000, 25,
	-1000, 1030, -1000, 23, 19, 961, 478, -1000, 409, -38,
	340, 4, 4, -1000, -1000, -1000, -1000, -1000, 18, 73,
	271, 4, -1000, 110, 72, -1000, -1000,
}
var yyPgo = []int{

	0, 271, 2, 270, 28, 268, 8, 50, 267, 265,
	264, 262, 260, 231, 258, 242, 242, 12, 1, 238,
	241, 5, 4, 235, 10, 0, 227, 6, 223,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 23, 23, 23, 28, 28, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 16, 16, 17, 17, 9, 9,
	9, 11, 11, 11, 11, 8, 8, 8, 8, 8,
	10, 10, 10, 12, 21, 21, 21, 20, 20, 4,
	4, 4, 4, 4, 4, 24, 24, 5, 5, 5,
	5, 5, 15, 15, 15, 6, 6, 6, 6, 6,
	6, 13, 13, 13, 13, 13, 7, 7, 7, 7,
	7, 7, 7, 7, 25, 22, 22, 18, 19, 26,
	27,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 1, 9, 4, 2,
	3, 1, 4, 5, 0, 1, 0, 3, 12, 10,
	6, 4, 4, 3, 6, 4, 6, 4, 3, 3,
	3, 3, 4, 4, 3, 0, 1, 0, 3, 4,
	6, 3, 4, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 6, 4, 7, 0, 9, 5, 1,
	1, 1, 2, 2, 0, 3, 0, 3, 3, 6,
	3, 4, 2, 3, 5, 1, 1, 3, 3, 3,
	4, 2, 3, 3, 3, 2, 1, 1, 2, 1,
	4, 1, 1, 1, 2, 1, 1, 1, 3, 1,
	1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -12, -20, 54, 49, 56,
	45, 36, -5, -6, -15, 41, 17, 42, 44, -25,
	62, -7, -13, -26, 53, 72, -22, 71, 63, 68,
	-19, -14, -18, 65, 50, 37, 60, 52, 15, 31,
	18, 16, 29, 30, 28, 19, 20, 33, 4, 5,
	6, 7, 39, 61, 62, 45, -19, -2, -24, 65,
	71, 72, 34, -6, 45, 62, -25, -6, -6, -4,
	34, 34, 47, 40, 50, -7, -13, -6, 62, 72,
	47, 62, 62, 66, -6, 67, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 57, 57, 57,
	57, 47, 50, 50, 64, 46, -4, 46, -6, -6,
	-6, -6, 47, 45, 36, 45, -6, -6, -6, -4,
	34, -24, -27, 51, -27, 62, 63, 60, 51, -7,
	62, -18, -7, -7, 62, -4, 51, -6, -7, -22,
	62, -6, -8, 23, 21, 22, 24, 25, -10, 9,
	14, 11, -9, 27, 32, -11, 10, 12, 26, 69,
	-6, -23, 46, -17, 59, -2, -4, -2, 48, 48,
	-6, 51, 48, 59, 59, 59, 51, 51, 59, 51,
	59, 51, 50, 50, 50, 50, 50, 50, 50, 50,
	50, 47, 50, 47, 50, 48, 51, 51, 53, 46,
	36, 46, 34, 62, -7, -7, -6, -6, -6, 51,
	51, 51, 51, 51, 51, 51, 51, 51, -4, 51,
	-4, 51, -28, 53, 62, 72, -21, 43, -4, -6,
	59, 59, 51, 51, 51, 48, 48, 45, -17, 62,
	45, 42, 45, -7, -7, 34, -2, -17, -2, -6,
	-2, 59, 59, -4, 46, 46, 45, 46, -22, -22,
	-2, 59, 51, 46, -22, -21, 51,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 5, 6, 0, 0, 0,
	-2, 11, 59, 60, 61, 0, 0, 0, 64, 87,
	97, 75, 76, 0, 0, 0, 86, 0, 89, 91,
	92, 93, 95, 0, 99, 96, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 9, 0, 64, 0,
	0, 0, 0, 62, 66, 97, 87, 63, 0, 0,
	0, 0, 64, 0, 66, 75, 76, 0, 72, 0,
	0, 94, 88, 81, 85, 0, 0, 0, 0, 0,
	0, 64, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 14, 0, 10, 37, 79, 82, 83,
	84, 70, 64, 3, 64, 3, 67, 68, 0, 0,
	0, 64, 77, 100, 78, 73, 0, 98, 17, 0,
	97, 95, 0, 0, 0, 0, 23, 0, 0, 0,
	97, 0, 28, 0, 0, 0, 0, 0, 29, 0,
	0, 0, 30, 0, 0, 31, 0, 43, 0, 0,
	0, 0, 8, 65, 36, 0, 0, 0, 0, 90,
	71, 80, 0, 0, 0, 0, 21, 22, 0, 25,
	0, 27, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 64, 0, 32, 33, 16, 0, 56,
	0, 58, 0, 74, 0, 0, 0, 0, 0, 45,
	46, 47, 48, 49, 50, 51, 52, 38, 0, 41,
	0, 44, 0, 15, 37, 0, 53, 0, 0, 69,
	0, 0, 20, 24, 26, 39, 42, 3, 12, 37,
	3, 0, 3, 0, 0, 64, 0, 13, 0, 0,
	0, 0, 0, 40, 7, 54, 3, 57, 0, 0,
	0, 0, 19, 56, 0, 55, 18,
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
	72,
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
		//line front/mutan.y:139
		{
			yyVAL.tnode = NewNode(PrintTy, yyS[yypt-1].tnode)
		}
	case 34:
		//line front/mutan.y:143
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 35:
		//line front/mutan.y:144
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 36:
		//line front/mutan.y:148
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 37:
		//line front/mutan.y:149
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 38:
		//line front/mutan.y:154
		{
			yyVAL.tnode = NewNode(AddressTy)
		}
	case 39:
		//line front/mutan.y:155
		{
			yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode)
		}
	case 40:
		//line front/mutan.y:157
		{
			node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 41:
		//line front/mutan.y:164
		{
			yyVAL.tnode = NewNode(CallerTy)
		}
	case 42:
		//line front/mutan.y:165
		{
			yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode)
		}
	case 43:
		//line front/mutan.y:166
		{
			yyVAL.tnode = NewNode(CallDataSizeTy)
		}
	case 44:
		//line front/mutan.y:167
		{
			yyVAL.tnode = NewNode(GasTy)
		}
	case 45:
		//line front/mutan.y:172
		{
			yyVAL.tnode = NewNode(TimestampTy)
		}
	case 46:
		//line front/mutan.y:173
		{
			yyVAL.tnode = NewNode(DiffTy)
		}
	case 47:
		//line front/mutan.y:174
		{
			yyVAL.tnode = NewNode(PrevHashTy)
		}
	case 48:
		//line front/mutan.y:175
		{
			yyVAL.tnode = NewNode(BlockNumTy)
		}
	case 49:
		//line front/mutan.y:176
		{
			yyVAL.tnode = NewNode(CoinbaseTy)
		}
	case 50:
		//line front/mutan.y:180
		{
			yyVAL.tnode = NewNode(OriginTy)
		}
	case 51:
		//line front/mutan.y:181
		{
			yyVAL.tnode = NewNode(GasPriceTy)
		}
	case 52:
		//line front/mutan.y:182
		{
			yyVAL.tnode = NewNode(CallValTy)
		}
	case 53:
		//line front/mutan.y:187
		{
			if yyS[yypt-0].tnode == nil {
				yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
			} else {
				yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			}
		}
	case 54:
		//line front/mutan.y:197
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 55:
		//line front/mutan.y:201
		{
			if yyS[yypt-0].tnode == nil {
				yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
			} else {
				yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			}
		}
	case 56:
		//line front/mutan.y:208
		{
			yyVAL.tnode = nil
		}
	case 57:
		//line front/mutan.y:213
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 58:
		//line front/mutan.y:217
		{
			yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		}
	case 59:
		//line front/mutan.y:223
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 60:
		//line front/mutan.y:224
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 61:
		//line front/mutan.y:225
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 62:
		//line front/mutan.y:226
		{
			yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode)
		}
	case 63:
		//line front/mutan.y:227
		{
			yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode)
		}
	case 64:
		//line front/mutan.y:228
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 65:
		//line front/mutan.y:232
		{
			yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode)
		}
	case 66:
		//line front/mutan.y:233
		{
			yyVAL.tnode = nil
		}
	case 67:
		//line front/mutan.y:238
		{
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode)
		}
	case 68:
		//line front/mutan.y:242
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-2].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		}
	case 69:
		//line front/mutan.y:248
		{
			yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-5].str
		}
	case 70:
		//line front/mutan.y:252
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-2].tnode.Constant
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		}
	case 71:
		//line front/mutan.y:258
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
			varNode := NewNode(NewVarTy)
			varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		}
	case 72:
		//line front/mutan.y:268
		{
			yyVAL.tnode = NewNode(NewVarTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 73:
		//line front/mutan.y:273
		{
			yyVAL.tnode = NewNode(NewVarTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
			yyVAL.tnode.Ptr = true
		}
	case 74:
		//line front/mutan.y:279
		{
			yyVAL.tnode = NewNode(NewArrayTy)
			yyVAL.tnode.Size = yyS[yypt-2].str
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 75:
		//line front/mutan.y:287
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 76:
		//line front/mutan.y:288
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 77:
		//line front/mutan.y:289
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 78:
		//line front/mutan.y:290
		{
			yyVAL.tnode = yyS[yypt-1].tnode
		}
	case 79:
		//line front/mutan.y:292
		{
			yyVAL.tnode = NewNode(InitListTy)
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
		}
	case 80:
		//line front/mutan.y:297
		{
			yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
		}
	case 81:
		//line front/mutan.y:306
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 82:
		//line front/mutan.y:308
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 83:
		//line front/mutan.y:309
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 84:
		//line front/mutan.y:310
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 85:
		//line front/mutan.y:311
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 86:
		//line front/mutan.y:315
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 87:
		//line front/mutan.y:316
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 88:
		//line front/mutan.y:317
		{
			yyVAL.tnode = NewNode(RefTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 89:
		//line front/mutan.y:318
		{
			yyVAL.tnode = NewNode(ConstantTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 90:
		//line front/mutan.y:319
		{
			yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
		}
	case 91:
		//line front/mutan.y:320
		{
			yyVAL.tnode = NewNode(BoolTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 92:
		//line front/mutan.y:321
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 93:
		//line front/mutan.y:322
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 94:
		//line front/mutan.y:326
		{
			yyVAL.tnode = NewNode(DerefPtrTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 95:
		//line front/mutan.y:330
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 96:
		//line front/mutan.y:331
		{
			yyVAL.tnode = NewNode(NilTy)
		}
	case 97:
		//line front/mutan.y:335
		{
			yyVAL.tnode = NewNode(IdentifierTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 98:
		//line front/mutan.y:339
		{
			yyVAL.tnode = NewNode(StringTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 99:
		//line front/mutan.y:343
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 100:
		//line front/mutan.y:347
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	}
	goto yystack /* stack new state and value */
}

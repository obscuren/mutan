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

//line front/mutan.y:344

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 100
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1104

var yyAct = []int{

	64, 32, 2, 19, 26, 223, 4, 253, 231, 118,
	198, 58, 155, 57, 83, 81, 230, 59, 60, 19,
	102, 229, 58, 35, 191, 67, 176, 58, 59, 60,
	58, 122, 123, 59, 60, 58, 59, 60, 58, 21,
	58, 59, 60, 173, 59, 60, 59, 60, 136, 58,
	162, 78, 109, 236, 221, 59, 60, 200, 19, 130,
	121, 80, 79, 73, 222, 54, 76, 58, 36, 258,
	249, 19, 58, 59, 60, 248, 77, 115, 59, 60,
	199, 228, 227, 175, 120, 170, 169, 127, 127, 127,
	19, 168, 98, 97, 127, 127, 131, 96, 135, 95,
	192, 220, 193, 263, 259, 13, 218, 216, 214, 19,
	213, 19, 159, 212, 161, 115, 211, 160, 19, 210,
	209, 62, 65, 66, 165, 125, 128, 129, 69, 75,
	208, 207, 134, 206, 71, 174, 172, 171, 132, 82,
	124, 70, 119, 108, 72, 189, 72, 187, 185, 184,
	183, 182, 181, 180, 179, 178, 177, 101, 100, 94,
	19, 93, 19, 92, 104, 105, 106, 107, 91, 90,
	127, 127, 89, 88, 112, 113, 114, 87, 86, 85,
	84, 233, 232, 190, 167, 163, 188, 19, 186, 19,
	108, 99, 158, 215, 110, 217, 19, 133, 238, 239,
	137, 237, 225, 111, 234, 55, 156, 224, 201, 202,
	195, 242, 197, 116, 68, 61, 149, 152, 219, 153,
	23, 150, 164, 140, 141, 139, 142, 143, 22, 127,
	127, 30, 235, 154, 145, 117, 147, 243, 157, 146,
	245, 56, 247, 19, 19, 6, 19, 244, 19, 250,
	127, 127, 74, 255, 256, 14, 257, 31, 19, 5,
	127, 151, 144, 261, 148, 138, 262, 240, 241, 12,
	3, 1, 48, 49, 50, 51, 203, 0, 0, 204,
	0, 205, 0, 38, 41, 16, 40, 45, 46, 0,
	0, 0, 0, 0, 0, 0, 44, 42, 43, 39,
	0, 47, 0, 226, 11, 35, 0, 52, 0, 15,
	17, 0, 18, 10, 260, 0, 0, 8, 34, 0,
	37, 24, 7, 0, 9, 0, 0, 0, 36, 53,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 0, 0, 0, 246, 48, 49, 50, 51, 0,
	0, 0, 0, 0, 0, 0, 38, 41, 16, 40,
	45, 46, 0, 0, 0, 0, 0, 0, 0, 44,
	42, 43, 39, 0, 47, 0, 0, 11, 35, 0,
	52, 0, 15, 17, 0, 18, 10, 254, 0, 0,
	8, 34, 0, 37, 24, 7, 0, 9, 0, 0,
	0, 36, 53, 20, 28, 0, 33, 0, 0, 29,
	0, 0, 27, 25, 48, 49, 50, 51, 0, 0,
	0, 0, 0, 0, 0, 38, 41, 16, 40, 45,
	46, 0, 0, 0, 0, 0, 0, 0, 44, 42,
	43, 39, 0, 47, 0, 0, 11, 35, 0, 52,
	0, 15, 17, 0, 18, 10, 252, 0, 0, 8,
	34, 0, 37, 24, 7, 0, 9, 0, 0, 0,
	36, 53, 20, 28, 0, 33, 0, 0, 29, 0,
	0, 27, 25, 48, 49, 50, 51, 0, 0, 0,
	0, 0, 0, 0, 38, 41, 16, 40, 45, 46,
	0, 0, 0, 0, 0, 0, 0, 44, 42, 43,
	39, 0, 47, 0, 0, 11, 35, 0, 52, 0,
	15, 17, 0, 18, 10, 251, 0, 0, 8, 34,
	0, 37, 24, 7, 0, 9, 0, 0, 0, 36,
	53, 20, 28, 0, 33, 0, 0, 29, 0, 0,
	27, 25, 48, 49, 50, 51, 0, 0, 0, 0,
	0, 0, 0, 38, 41, 16, 40, 45, 46, 0,
	0, 0, 0, 0, 0, 0, 44, 42, 43, 39,
	0, 47, 0, 0, 11, 35, 0, 52, 0, 15,
	17, 0, 18, 10, 196, 0, 0, 8, 34, 0,
	37, 24, 7, 0, 9, 0, 0, 0, 36, 53,
	20, 28, 0, 33, 0, 0, 29, 0, 0, 27,
	25, 48, 49, 50, 51, 0, 0, 0, 0, 0,
	0, 0, 38, 41, 16, 40, 45, 46, 0, 0,
	0, 0, 0, 0, 0, 44, 42, 43, 39, 0,
	47, 0, 0, 11, 35, 0, 52, 0, 15, 17,
	0, 18, 10, 194, 0, 0, 8, 34, 0, 37,
	24, 7, 0, 9, 0, 0, 0, 36, 53, 20,
	28, 0, 33, 0, 0, 29, 0, 0, 27, 25,
	48, 49, 50, 51, 0, 0, 0, 0, 0, 0,
	0, 38, 41, 16, 40, 45, 46, 0, 0, 0,
	0, 0, 0, 0, 44, 42, 43, 39, 0, 47,
	0, 0, 11, 35, 0, 52, 0, 15, 17, 0,
	18, 10, 103, 0, 0, 8, 34, 0, 37, 24,
	7, 0, 9, 0, 0, 0, 36, 53, 20, 28,
	0, 33, 0, 0, 29, 0, 0, 27, 25, 48,
	49, 50, 51, 0, 0, 0, 0, 0, 0, 0,
	38, 41, 16, 40, 45, 46, 0, 0, 0, 0,
	0, 0, 0, 44, 42, 43, 39, 0, 47, 0,
	0, 11, 35, 0, 52, 0, 15, 17, 0, 18,
	10, 0, 0, 0, 8, 34, 0, 37, 24, 7,
	0, 9, 0, 0, 0, 36, 53, 20, 28, 0,
	33, 0, 0, 29, 0, 0, 27, 25, 48, 49,
	50, 51, 0, 0, 0, 0, 0, 0, 0, 38,
	41, 16, 40, 45, 46, 0, 0, 0, 0, 0,
	0, 0, 44, 42, 43, 39, 0, 47, 0, 0,
	0, 35, 0, 52, 0, 15, 0, 0, 0, 0,
	0, 0, 0, 0, 34, 166, 37, 24, 0, 0,
	0, 0, 0, 0, 36, 53, 20, 28, 0, 33,
	0, 0, 29, 0, 0, 27, 25, 48, 49, 50,
	51, 0, 0, 0, 0, 0, 0, 0, 38, 41,
	16, 40, 45, 46, 0, 0, 0, 0, 0, 0,
	0, 44, 42, 43, 39, 0, 47, 0, 0, 0,
	35, 0, 52, 0, 15, 0, 0, 0, 0, 0,
	0, 0, 0, 34, 0, 37, 24, 0, 0, 0,
	0, 0, 0, 36, 53, 20, 28, 0, 33, 0,
	0, 29, 0, 0, 27, 25, 48, 49, 50, 51,
	0, 0, 0, 0, 0, 0, 0, 38, 41, 0,
	40, 45, 46, 0, 0, 0, 0, 0, 0, 0,
	44, 42, 43, 39, 0, 47, 0, 0, 0, 35,
	0, 52, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 34, 0, 37, 0, 0, 0, 0, 0,
	0, 0, 36, 53, 63, 28, 0, 33, 0, 0,
	29, 0, 0, 27, 25, 48, 49, 50, 51, 0,
	0, 0, 0, 0, 0, 0, 38, 41, 0, 40,
	45, 46, 0, 0, 0, 0, 0, 0, 0, 44,
	42, 43, 39, 0, 47, 0, 0, 0, 35, 0,
	52, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 37, 0, 0, 0, 0, 0, 0,
	0, 36, 53, 126, 28, 0, 0, 0, 0, 29,
	0, 0, 27, 25,
}
var yyPact = []int{

	-1000, -1000, 755, -1000, -1000, -1000, -1000, 3, 160, 8,
	-1000, -1000, -1000, -54, 181, 962, 962, 962, 893, 180,
	94, -1000, -1000, 962, 4, 0, -1000, -1, -1000, -1000,
	-1000, -1000, -51, 962, -1000, -1000, -53, 130, 129, 128,
	127, 123, 122, 119, 118, 113, 111, 109, 42, 40,
	36, 35, 144, 108, 107, -44, -1000, 686, 962, 962,
	962, 962, -54, 96, -1000, -54, 7, 158, 962, 962,
	893, 179, -1000, 91, 91, -54, -1000, -2, -32, -1000,
	-1000, -1000, -54, -28, 89, 1031, 1031, 1031, -3, 893,
	87, 962, 1031, -14, 962, 202, 225, 189, 207, -57,
	962, -1000, 146, -1000, -54, -54, -54, -54, 893, -1000,
	893, -1000, -54, -54, 2, 137, 962, 824, -1000, -1000,
	-1000, -1000, 136, -1000, -1000, 32, 143, -1000, 27, 26,
	86, 85, -1000, -16, 84, 24, -1000, -25, -1000, 106,
	105, 104, 103, 102, -1000, 101, 100, 99, -1000, 98,
	141, -1000, 97, 139, 95, 135, -27, 49, -1000, 617,
	174, 548, 178, -1000, -54, 21, -1000, -5, 1031, 1031,
	962, -1000, -1000, 962, -1000, 962, -1000, 82, 80, 79,
	69, 68, 65, 62, 59, 57, 893, 56, 893, 55,
	-1000, -1000, 48, -8, 164, 893, -1000, 962, -1000, -1000,
	-1000, 23, 22, -30, -35, -43, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 134, -1000, 133, -1000, 159,
	-1000, 21, -9, -1000, 156, 154, -54, 1031, 1031, -1000,
	-1000, -1000, 177, -1000, -1000, -1000, 21, -1000, 962, -1000,
	16, 11, 893, 479, -1000, 410, -38, 341, -14, -14,
	-1000, -1000, -1000, -1000, -1000, 10, 53, 268, -14, -1000,
	164, 52, -1000, -1000,
}
var yyPgo = []int{

	0, 271, 2, 270, 6, 269, 105, 39, 265, 264,
	262, 261, 259, 228, 257, 255, 255, 10, 1, 231,
	245, 5, 4, 238, 235, 0, 220, 9, 218,
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
	13, 13, 13, 13, 13, 7, 7, 7, 7, 7,
	7, 7, 7, 25, 22, 22, 18, 19, 26, 27,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 1, 9, 4, 2,
	3, 1, 4, 5, 0, 1, 0, 3, 12, 10,
	6, 4, 4, 3, 6, 4, 6, 4, 3, 3,
	3, 3, 4, 4, 3, 0, 1, 0, 3, 4,
	6, 3, 4, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 6, 4, 7, 0, 9, 5, 1,
	1, 1, 2, 2, 0, 3, 0, 3, 3, 6,
	3, 4, 2, 3, 5, 1, 1, 3, 3, 4,
	2, 3, 3, 3, 2, 1, 1, 2, 1, 4,
	1, 1, 1, 2, 1, 1, 1, 3, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -12, -20, 54, 49, 56,
	45, 36, -5, -6, -15, 41, 17, 42, 44, -25,
	62, -7, -13, -26, 53, 72, -22, 71, 63, 68,
	-19, -14, -18, 65, 50, 37, 60, 52, 15, 31,
	18, 16, 29, 30, 28, 19, 20, 33, 4, 5,
	6, 7, 39, 61, 62, 45, -19, -2, 65, 71,
	72, 34, -6, 62, -25, -6, -6, -4, 34, 34,
	47, 40, 50, -7, -13, -6, 62, 72, 47, 62,
	62, 66, -6, 67, 50, 50, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 57, 57, 57, 57, 47,
	50, 50, 64, 46, -6, -6, -6, -6, 47, 45,
	36, 45, -6, -6, -6, -4, 34, -24, -27, 51,
	-27, 62, 63, 60, 51, -7, 62, -18, -7, -7,
	62, -4, 51, -6, -7, -22, 62, -6, -8, 23,
	21, 22, 24, 25, -10, 9, 14, 11, -9, 27,
	32, -11, 10, 12, 26, 69, -6, -23, 46, -2,
	-4, -2, 48, 48, -6, -4, 51, 48, 59, 59,
	59, 51, 51, 59, 51, 59, 51, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 47, 50, 47, 50,
	48, 51, 51, 53, 46, 36, 46, 34, -17, 59,
	62, -7, -7, -6, -6, -6, 51, 51, 51, 51,
	51, 51, 51, 51, 51, -4, 51, -4, 51, -28,
	53, 62, 72, -21, 43, -4, -6, 59, 59, 51,
	51, 51, 48, 48, 45, -17, 62, 45, 42, 45,
	-7, -7, 34, -2, -17, -2, -6, -2, 59, 59,
	-4, 46, 46, 45, 46, -22, -22, -2, 59, 51,
	46, -22, -21, 51,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 5, 6, 0, 0, 0,
	3, 11, 59, 60, 61, 0, 0, 0, 64, 86,
	96, 75, 76, 0, 0, 0, 85, 0, 88, 90,
	91, 92, 94, 0, 98, 95, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 62, 96, 86, 63, 0, 0, 0, 0,
	64, 0, 66, 75, 76, 0, 72, 0, 0, 93,
	87, 80, 84, 0, 0, 0, 0, 0, 0, 64,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 14, 0, 10, 81, 82, 83, 70, 64, 3,
	64, 3, 67, 68, 0, 0, 0, 64, 77, 99,
	78, 73, 0, 97, 17, 0, 96, 94, 0, 0,
	0, 0, 23, 0, 0, 0, 96, 0, 28, 0,
	0, 0, 0, 0, 29, 0, 0, 0, 30, 0,
	0, 31, 0, 43, 0, 0, 0, 0, 8, 0,
	0, 0, 0, 89, 71, 37, 79, 0, 0, 0,
	0, 21, 22, 0, 25, 0, 27, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 64, 0,
	32, 33, 16, 0, 56, 64, 58, 0, 65, 36,
	74, 0, 0, 0, 0, 0, 45, 46, 47, 48,
	49, 50, 51, 52, 38, 0, 41, 0, 44, 0,
	15, 37, 0, 53, 0, 0, 69, 0, 0, 20,
	24, 26, 39, 42, 3, 12, 37, 3, 0, 3,
	0, 0, 64, 0, 13, 0, 0, 0, 0, 0,
	40, 7, 54, 3, 57, 0, 0, 0, 0, 19,
	56, 0, 55, 18,
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
			yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
			yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
		}
	case 80:
		//line front/mutan.y:301
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 81:
		//line front/mutan.y:303
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 82:
		//line front/mutan.y:304
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 83:
		//line front/mutan.y:305
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 84:
		//line front/mutan.y:306
		{
			yyVAL.tnode = NewNode(OpTy, yyS[yypt-0].tnode)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 85:
		//line front/mutan.y:310
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 86:
		//line front/mutan.y:311
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 87:
		//line front/mutan.y:312
		{
			yyVAL.tnode = NewNode(RefTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 88:
		//line front/mutan.y:313
		{
			yyVAL.tnode = NewNode(ConstantTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 89:
		//line front/mutan.y:314
		{
			yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode)
			yyVAL.tnode.Constant = yyS[yypt-3].str
		}
	case 90:
		//line front/mutan.y:315
		{
			yyVAL.tnode = NewNode(BoolTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 91:
		//line front/mutan.y:316
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 92:
		//line front/mutan.y:317
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 93:
		//line front/mutan.y:321
		{
			yyVAL.tnode = NewNode(DerefPtrTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 94:
		//line front/mutan.y:325
		{
			yyVAL.tnode = yyS[yypt-0].tnode
		}
	case 95:
		//line front/mutan.y:326
		{
			yyVAL.tnode = NewNode(NilTy)
		}
	case 96:
		//line front/mutan.y:330
		{
			yyVAL.tnode = NewNode(IdentifierTy)
			yyVAL.tnode.Constant = yyS[yypt-0].str
		}
	case 97:
		//line front/mutan.y:334
		{
			yyVAL.tnode = NewNode(StringTy)
			yyVAL.tnode.Constant = yyS[yypt-1].str
		}
	case 98:
		//line front/mutan.y:338
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	case 99:
		//line front/mutan.y:342
		{
			yyVAL.tnode = NewNode(EmptyTy)
		}
	}
	goto yystack /* stack new state and value */
}

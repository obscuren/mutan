
//line front/mutan.y:2
package frontend
import __yyfmt__ "fmt"
//line front/mutan.y:2
		
import (
_	"fmt"
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
	yys int
	num int
	str string
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
const DIFFICULTY = 57387
const PREVHASH = 57388
const TIMESTAMP = 57389
const BLOCKNUM = 57390
const COINBASE = 57391
const GAS = 57392
const VAR = 57393
const FUNC = 57394
const FUNC_CALL = 57395
const IMPORT = 57396
const ID = 57397
const NUMBER = 57398
const INLINE_ASM = 57399
const OP = 57400
const DOP = 57401
const STR = 57402
const BOOLEAN = 57403
const CODE = 57404
const oper = 57405
const AND = 57406
const MUL = 57407

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

//line front/mutan.y:307



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 91
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 788

var yyAct = []int{

	32, 2, 26, 4, 170, 131, 194, 13, 100, 73,
	137, 72, 48, 21, 84, 91, 195, 55, 104, 34,
	19, 49, 58, 53, 56, 57, 207, 50, 51, 69,
	172, 66, 112, 103, 71, 70, 19, 64, 176, 45,
	164, 193, 167, 105, 35, 60, 228, 225, 218, 217,
	171, 201, 200, 61, 145, 49, 63, 86, 87, 88,
	89, 50, 51, 93, 49, 97, 19, 94, 95, 96,
	50, 51, 67, 144, 102, 165, 109, 109, 109, 19,
	143, 62, 68, 113, 81, 22, 92, 226, 202, 107,
	110, 111, 190, 134, 97, 136, 135, 19, 189, 188,
	187, 186, 185, 140, 184, 183, 139, 182, 19, 65,
	19, 180, 179, 178, 177, 147, 146, 19, 114, 106,
	101, 90, 161, 160, 63, 130, 159, 158, 157, 133,
	156, 155, 154, 116, 118, 119, 120, 121, 125, 153,
	151, 150, 149, 148, 109, 109, 109, 83, 175, 80,
	79, 128, 19, 78, 19, 117, 181, 173, 174, 122,
	123, 124, 126, 127, 129, 77, 191, 76, 75, 74,
	19, 198, 204, 203, 163, 142, 138, 199, 162, 152,
	19, 90, 82, 30, 209, 19, 208, 205, 46, 197,
	212, 169, 98, 47, 59, 52, 192, 23, 99, 206,
	132, 109, 109, 196, 6, 14, 115, 213, 31, 5,
	215, 216, 214, 12, 210, 211, 219, 3, 109, 109,
	223, 224, 1, 0, 0, 0, 109, 0, 227, 0,
	19, 19, 0, 19, 19, 17, 0, 18, 10, 222,
	0, 0, 0, 8, 33, 0, 36, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 0, 37, 0, 40,
	35, 11, 16, 39, 38, 34, 0, 0, 44, 0,
	0, 15, 41, 42, 0, 0, 0, 0, 0, 0,
	24, 7, 0, 9, 20, 28, 0, 0, 0, 0,
	29, 0, 0, 27, 25, 17, 0, 18, 10, 221,
	0, 0, 0, 8, 33, 0, 36, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 0, 37, 0, 40,
	35, 11, 16, 39, 38, 34, 0, 0, 44, 0,
	0, 15, 41, 42, 0, 0, 0, 0, 0, 0,
	24, 7, 0, 9, 20, 28, 0, 0, 0, 0,
	29, 0, 0, 27, 25, 17, 0, 18, 10, 220,
	0, 0, 0, 8, 33, 0, 36, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 0, 37, 0, 40,
	35, 11, 16, 39, 38, 34, 0, 0, 44, 0,
	0, 15, 41, 42, 0, 0, 0, 0, 0, 0,
	24, 7, 0, 9, 20, 28, 0, 0, 0, 0,
	29, 0, 0, 27, 25, 17, 0, 18, 10, 168,
	0, 0, 0, 8, 33, 0, 36, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 0, 37, 0, 40,
	35, 11, 16, 39, 38, 34, 0, 0, 44, 0,
	0, 15, 41, 42, 0, 0, 0, 0, 0, 0,
	24, 7, 0, 9, 20, 28, 0, 0, 0, 0,
	29, 0, 0, 27, 25, 17, 0, 18, 10, 166,
	0, 0, 0, 8, 33, 0, 36, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 0, 37, 0, 40,
	35, 11, 16, 39, 38, 34, 0, 0, 44, 0,
	0, 15, 41, 42, 0, 0, 0, 0, 0, 0,
	24, 7, 0, 9, 20, 28, 0, 0, 0, 0,
	29, 0, 0, 27, 25, 17, 0, 18, 10, 85,
	0, 0, 0, 8, 33, 0, 36, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 0, 37, 0, 40,
	35, 11, 16, 39, 38, 34, 0, 0, 44, 0,
	0, 15, 41, 42, 0, 0, 0, 0, 0, 0,
	24, 7, 0, 9, 20, 28, 0, 0, 0, 17,
	29, 18, 10, 27, 25, 0, 0, 8, 33, 0,
	36, 0, 0, 0, 0, 0, 0, 0, 0, 43,
	0, 37, 0, 40, 35, 11, 16, 39, 38, 34,
	0, 0, 44, 0, 0, 15, 41, 42, 0, 0,
	0, 0, 0, 0, 24, 7, 0, 9, 20, 28,
	33, 141, 36, 0, 29, 0, 0, 27, 25, 0,
	0, 43, 0, 37, 0, 40, 35, 0, 16, 39,
	38, 34, 0, 0, 44, 0, 0, 15, 41, 42,
	0, 0, 0, 0, 0, 0, 24, 0, 0, 0,
	20, 28, 33, 0, 36, 0, 29, 0, 0, 27,
	25, 0, 0, 43, 0, 37, 0, 40, 35, 0,
	16, 39, 38, 34, 0, 0, 44, 0, 0, 15,
	41, 42, 0, 0, 33, 0, 36, 0, 24, 0,
	0, 0, 20, 28, 0, 43, 0, 37, 29, 40,
	35, 27, 25, 39, 38, 34, 0, 0, 44, 36,
	0, 0, 41, 42, 0, 0, 0, 0, 43, 0,
	37, 0, 40, 35, 54, 28, 39, 38, 34, 0,
	29, 44, 0, 27, 25, 41, 42, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 108, 28, 0,
	0, 0, 0, 29, 0, 0, 27, 25,
}
var yyPact = []int{

	-1000, -1000, 583, -1000, -1000, -1000, -1000, -16, 179, 13,
	-1000, -1000, -1000, -37, 191, 699, 699, 699, 667, 190,
	41, -1000, -1000, 699, 17, -20, -1000, -21, -1000, -1000,
	-1000, -1000, -48, -1000, -1000, -51, 154, 153, 152, 150,
	138, 135, 134, 59, 170, 132, -43, -1000, 529, 699,
	699, 699, 699, -37, 109, -1000, -37, 6, 54, 699,
	699, 667, 188, -1000, 104, 104, -37, -1000, -22, -38,
	-1000, -1000, -1000, 12, 103, 722, 722, 722, -23, 667,
	102, 114, -57, -1000, 119, -1000, -37, -37, -37, -37,
	667, -1000, 667, -1000, -37, -37, -3, 163, 699, 625,
	-1000, -1000, -1000, -1000, 162, -1000, -1000, 51, 169, -1000,
	44, 25, 100, 99, -1000, -1000, 128, 127, 126, 125,
	167, 124, 117, 116, 115, 113, 112, 111, 108, 107,
	166, 161, 24, -1000, 469, 10, 409, 187, -1000, -37,
	21, -1000, -25, 722, 722, -17, -1000, -1000, 98, 97,
	96, 95, 667, 91, 89, 88, 86, 85, 84, 83,
	82, 76, 667, -1000, -10, -49, 182, 667, -1000, 699,
	-1000, -1000, -1000, 23, 22, 72, -1000, -1000, -1000, -1000,
	-1000, 160, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 159, 178, -1000, 21, -29, -1000, 177, 175, -37,
	722, 722, -1000, -1000, 186, -1000, -1000, 21, -1000, -1000,
	20, 19, 667, 349, -1000, 289, 229, -17, -17, -1000,
	-1000, -1000, -1000, 18, 71, -17, -1000, 30, -1000,
}
var yyPgo = []int{

	0, 222, 1, 217, 3, 213, 7, 13, 209, 85,
	208, 206, 205, 205, 4, 0, 183, 204, 203, 2,
	200, 198, 17, 197, 8, 196,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 20, 20, 20, 25, 25, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 13, 13, 14, 14,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 8, 18, 18, 17,
	17, 4, 4, 4, 4, 4, 4, 21, 21, 5,
	5, 5, 5, 5, 12, 12, 12, 6, 6, 6,
	6, 6, 9, 9, 9, 9, 7, 7, 7, 7,
	7, 7, 7, 7, 22, 19, 19, 15, 16, 23,
	24,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 1, 9, 4, 2,
	3, 1, 4, 5, 0, 1, 0, 3, 12, 10,
	6, 4, 4, 3, 3, 4, 3, 0, 1, 0,
	3, 3, 3, 3, 4, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 6, 6, 4, 0, 9,
	5, 1, 1, 1, 2, 2, 0, 3, 0, 3,
	3, 6, 3, 4, 2, 3, 5, 1, 1, 3,
	3, 4, 2, 3, 3, 3, 1, 1, 2, 1,
	4, 1, 1, 1, 2, 1, 1, 1, 3, 1,
	1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -8, -17, 52, 14, 54,
	9, 32, -5, -6, -12, 42, 33, 6, 8, -22,
	55, -7, -9, -23, 51, 65, -19, 64, 56, 61,
	-16, -10, -15, 15, 36, 31, 17, 28, 35, 34,
	30, 43, 44, 26, 39, 55, 9, -16, -2, 58,
	64, 65, 4, -6, 55, -22, -6, -6, -4, 4,
	4, 12, 40, 15, -7, -9, -6, 55, 65, 12,
	55, 55, 59, 60, 15, 15, 15, 15, 15, 15,
	15, 25, 12, 15, 57, 10, -6, -6, -6, -6,
	12, 9, 32, 9, -6, -6, -6, -4, 4, -21,
	-24, 16, -24, 55, 56, 31, 16, -7, 55, -15,
	-7, -7, 55, -4, 16, -11, 19, 41, 20, 21,
	22, 23, 45, 46, 47, 24, 48, 49, 37, 50,
	11, 62, -20, 10, -2, -4, -2, 13, 13, -6,
	-4, 16, 13, 29, 29, 29, 16, 16, 15, 15,
	15, 15, 12, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 12, 13, 16, 51, 10, 32, 10, 4,
	-14, 29, 55, -7, -7, -19, 55, 16, 16, 16,
	16, -4, 16, 16, 16, 16, 16, 16, 16, 16,
	16, -4, -25, 51, 55, 65, -18, 7, -4, -6,
	29, 29, 16, 13, 13, 9, -14, 55, 9, 9,
	-7, -7, 4, -2, -14, -2, -2, 29, 29, -4,
	10, 10, 10, -19, -19, 29, 16, -19, 16,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 5, 6, 0, 0, 0,
	3, 11, 51, 52, 53, 0, 0, 0, 56, 77,
	87, 67, 68, 0, 0, 0, 76, 0, 79, 81,
	82, 83, 85, 89, 86, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 9, 0, 0,
	0, 0, 0, 54, 87, 77, 55, 0, 0, 0,
	0, 56, 0, 58, 67, 68, 0, 64, 0, 0,
	84, 78, 72, 0, 0, 0, 0, 0, 0, 56,
	0, 0, 0, 14, 0, 10, 73, 74, 75, 62,
	56, 3, 56, 3, 59, 60, 0, 0, 0, 56,
	69, 90, 70, 65, 0, 88, 17, 0, 87, 85,
	0, 0, 0, 0, 23, 24, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 8, 0, 0, 0, 0, 80, 63,
	29, 71, 0, 0, 0, 0, 21, 22, 0, 0,
	0, 0, 56, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 56, 25, 16, 0, 48, 56, 50, 0,
	57, 28, 66, 0, 0, 0, 87, 30, 31, 32,
	33, 0, 35, 36, 37, 38, 39, 40, 41, 42,
	43, 0, 0, 15, 29, 0, 46, 0, 0, 61,
	0, 0, 20, 34, 44, 3, 12, 29, 3, 3,
	0, 0, 56, 0, 13, 0, 0, 0, 0, 45,
	7, 47, 49, 0, 0, 0, 19, 0, 18,
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
	62, 63, 64, 65,
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
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line front/mutan.y:69
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode) }
	case 3:
		//line front/mutan.y:70
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line front/mutan.y:75
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 5:
		//line front/mutan.y:76
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 6:
		//line front/mutan.y:77
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 7:
		//line front/mutan.y:79
		{
				yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode);
				yyVAL.tnode.Constant = yyS[yypt-7].str
				yyVAL.tnode.HasRet = yyS[yypt-3].check
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-5].tnode, true)
			}
	case 8:
		//line front/mutan.y:85
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 9:
		//line front/mutan.y:86
		{ yyVAL.tnode = NewNode(ImportTy); yyVAL.tnode.Constant = yyS[yypt-0].tnode.Constant }
	case 10:
		//line front/mutan.y:87
		{ yyVAL.tnode = NewNode(ScopeTy, yyS[yypt-1].tnode) }
	case 11:
		//line front/mutan.y:88
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 12:
		//line front/mutan.y:92
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 13:
		//line front/mutan.y:93
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-4].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str; yyVAL.tnode.Ptr = true }
	case 14:
		//line front/mutan.y:94
		{ yyVAL.tnode = nil }
	case 15:
		//line front/mutan.y:99
		{ yyVAL.check = true }
	case 16:
		//line front/mutan.y:100
		{ yyVAL.check = false }
	case 17:
		//line front/mutan.y:105
		{ yyVAL.tnode = NewNode(StopTy) }
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
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 21:
		//line front/mutan.y:116
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 22:
		//line front/mutan.y:117
		{ yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode) }
	case 23:
		//line front/mutan.y:118
		{ yyVAL.tnode = NewNode(PopTy) }
	case 24:
		//line front/mutan.y:119
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 25:
		//line front/mutan.y:120
		{ yyVAL.tnode = NewNode(LambdaTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 26:
		//line front/mutan.y:124
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 27:
		//line front/mutan.y:125
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 28:
		//line front/mutan.y:129
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 29:
		//line front/mutan.y:130
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 30:
		//line front/mutan.y:135
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 31:
		//line front/mutan.y:136
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 32:
		//line front/mutan.y:137
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 33:
		//line front/mutan.y:138
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 34:
		//line front/mutan.y:139
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 35:
		//line front/mutan.y:140
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 36:
		//line front/mutan.y:141
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 37:
		//line front/mutan.y:142
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 38:
		//line front/mutan.y:143
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 39:
		//line front/mutan.y:144
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 40:
		//line front/mutan.y:145
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 41:
		//line front/mutan.y:146
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 42:
		//line front/mutan.y:147
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 43:
		//line front/mutan.y:148
		{ yyVAL.tnode = NewNode(GasTy) }
	case 44:
		//line front/mutan.y:149
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 45:
		//line front/mutan.y:151
		{
				node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 46:
		//line front/mutan.y:159
		{
				if yyS[yypt-0].tnode == nil {
					yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
				} else {
					yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
				}
			}
	case 47:
		//line front/mutan.y:169
		{
				yyVAL.tnode = yyS[yypt-1].tnode
			}
	case 48:
		//line front/mutan.y:172
		{ yyVAL.tnode = nil }
	case 49:
		//line front/mutan.y:177
		{
				yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
			}
	case 50:
		//line front/mutan.y:181
		{
				yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
			}
	case 51:
		//line front/mutan.y:187
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 52:
		//line front/mutan.y:188
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 53:
		//line front/mutan.y:189
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 54:
		//line front/mutan.y:190
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 55:
		//line front/mutan.y:191
		{ yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode) }
	case 56:
		//line front/mutan.y:192
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 57:
		//line front/mutan.y:196
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode);}
	case 58:
		//line front/mutan.y:197
		{ yyVAL.tnode = nil }
	case 59:
		//line front/mutan.y:202
		{
		      		yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode)
			}
	case 60:
		//line front/mutan.y:206
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 61:
		//line front/mutan.y:212
		{
				yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
			}
	case 62:
		//line front/mutan.y:216
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].tnode.Constant
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
			}
	case 63:
		//line front/mutan.y:222
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-3].str
				varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
			}
	case 64:
		//line front/mutan.y:232
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 65:
		//line front/mutan.y:237
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
				yyVAL.tnode.Ptr = true
			}
	case 66:
		//line front/mutan.y:243
		{
				yyVAL.tnode = NewNode(NewArrayTy)
				yyVAL.tnode.Size = yyS[yypt-2].str
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 67:
		//line front/mutan.y:251
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 68:
		//line front/mutan.y:252
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 69:
		//line front/mutan.y:253
		{ yyVAL.tnode = yyS[yypt-1].tnode }
	case 70:
		//line front/mutan.y:254
		{ yyVAL.tnode = yyS[yypt-1].tnode }
	case 71:
		//line front/mutan.y:256
		{
				yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
				yyVAL.tnode.Constant = yyS[yypt-3].str
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
			}
	case 72:
		//line front/mutan.y:265
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 73:
		//line front/mutan.y:267
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 74:
		//line front/mutan.y:268
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 75:
		//line front/mutan.y:269
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 76:
		//line front/mutan.y:273
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 77:
		//line front/mutan.y:274
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 78:
		//line front/mutan.y:275
		{ yyVAL.tnode = NewNode(RefTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 79:
		//line front/mutan.y:276
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 80:
		//line front/mutan.y:277
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 81:
		//line front/mutan.y:278
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 82:
		//line front/mutan.y:279
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 83:
		//line front/mutan.y:280
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 84:
		//line front/mutan.y:284
		{ yyVAL.tnode = NewNode(DerefPtrTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 85:
		//line front/mutan.y:288
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 86:
		//line front/mutan.y:289
		{ yyVAL.tnode = NewNode(NilTy) }
	case 87:
		//line front/mutan.y:293
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 88:
		//line front/mutan.y:297
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 89:
		//line front/mutan.y:301
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 90:
		//line front/mutan.y:305
		{ yyVAL.tnode = NewNode(EmptyTy) }
	}
	goto yystack /* stack new state and value */
}

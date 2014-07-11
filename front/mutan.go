
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

const yyLast = 756

var yyAct = []int{

	32, 170, 2, 26, 4, 201, 137, 69, 193, 100,
	49, 131, 55, 48, 73, 19, 50, 51, 194, 72,
	13, 84, 34, 58, 104, 207, 172, 112, 103, 71,
	91, 19, 21, 70, 45, 192, 53, 56, 57, 93,
	167, 224, 105, 35, 66, 226, 218, 49, 164, 217,
	67, 49, 171, 50, 51, 200, 64, 50, 51, 161,
	68, 19, 92, 199, 145, 144, 97, 143, 81, 22,
	86, 87, 88, 89, 19, 102, 109, 109, 109, 49,
	94, 95, 96, 165, 113, 50, 51, 229, 227, 203,
	60, 189, 19, 65, 134, 97, 136, 135, 61, 187,
	186, 63, 160, 19, 140, 19, 185, 184, 107, 110,
	111, 183, 19, 182, 33, 141, 36, 181, 179, 139,
	178, 177, 176, 147, 146, 43, 62, 37, 114, 40,
	35, 106, 16, 39, 38, 34, 101, 90, 44, 159,
	63, 15, 41, 42, 109, 109, 158, 19, 157, 19,
	24, 156, 155, 154, 20, 28, 153, 180, 151, 150,
	29, 109, 149, 27, 25, 19, 175, 190, 148, 83,
	80, 79, 197, 78, 77, 19, 173, 174, 76, 75,
	19, 74, 204, 202, 163, 142, 138, 162, 152, 90,
	198, 82, 133, 188, 30, 206, 209, 208, 205, 46,
	109, 109, 196, 212, 47, 169, 98, 59, 213, 214,
	52, 215, 216, 191, 23, 99, 132, 219, 109, 109,
	195, 223, 225, 6, 14, 19, 19, 109, 19, 19,
	228, 115, 210, 211, 17, 31, 18, 10, 222, 5,
	12, 3, 8, 33, 1, 36, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 0, 37, 0, 40, 35,
	11, 16, 39, 38, 34, 0, 0, 44, 0, 0,
	15, 41, 42, 0, 0, 0, 0, 0, 0, 24,
	7, 0, 9, 20, 28, 0, 0, 0, 0, 29,
	0, 0, 27, 25, 17, 0, 18, 10, 221, 0,
	0, 0, 8, 33, 0, 36, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 0, 37, 0, 40, 35,
	11, 16, 39, 38, 34, 0, 0, 44, 0, 0,
	15, 41, 42, 0, 0, 0, 0, 0, 0, 24,
	7, 0, 9, 20, 28, 0, 0, 0, 0, 29,
	0, 0, 27, 25, 17, 0, 18, 10, 220, 0,
	0, 0, 8, 33, 0, 36, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 0, 37, 0, 40, 35,
	11, 16, 39, 38, 34, 0, 0, 44, 0, 0,
	15, 41, 42, 0, 0, 0, 0, 0, 0, 24,
	7, 0, 9, 20, 28, 0, 0, 0, 0, 29,
	0, 0, 27, 25, 17, 0, 18, 10, 168, 0,
	0, 0, 8, 33, 0, 36, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 0, 37, 0, 40, 35,
	11, 16, 39, 38, 34, 0, 0, 44, 0, 0,
	15, 41, 42, 0, 0, 0, 0, 0, 0, 24,
	7, 0, 9, 20, 28, 0, 0, 0, 0, 29,
	0, 0, 27, 25, 17, 0, 18, 10, 166, 0,
	0, 0, 8, 33, 0, 36, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 0, 37, 0, 40, 35,
	11, 16, 39, 38, 34, 0, 0, 44, 0, 0,
	15, 41, 42, 0, 0, 0, 0, 0, 0, 24,
	7, 0, 9, 20, 28, 0, 0, 0, 0, 29,
	0, 0, 27, 25, 17, 0, 18, 10, 85, 0,
	0, 0, 8, 33, 0, 36, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 0, 37, 0, 40, 35,
	11, 16, 39, 38, 34, 0, 0, 44, 0, 0,
	15, 41, 42, 0, 0, 0, 0, 0, 0, 24,
	7, 0, 9, 20, 28, 0, 0, 0, 17, 29,
	18, 10, 27, 25, 0, 0, 8, 33, 0, 36,
	0, 0, 0, 0, 0, 0, 0, 0, 43, 0,
	37, 0, 40, 35, 11, 16, 39, 38, 34, 0,
	0, 44, 0, 0, 15, 41, 42, 0, 0, 0,
	0, 0, 0, 24, 7, 0, 9, 20, 28, 33,
	0, 36, 0, 29, 0, 0, 27, 25, 0, 0,
	43, 0, 37, 0, 40, 35, 0, 16, 39, 38,
	34, 0, 0, 44, 0, 0, 15, 41, 42, 0,
	0, 33, 0, 36, 0, 24, 0, 0, 0, 20,
	28, 0, 43, 0, 37, 29, 40, 35, 27, 25,
	39, 38, 34, 0, 0, 44, 36, 0, 0, 41,
	42, 0, 0, 0, 0, 43, 0, 37, 0, 40,
	35, 54, 28, 39, 38, 34, 130, 29, 44, 0,
	27, 25, 41, 42, 116, 118, 119, 120, 121, 125,
	0, 0, 0, 0, 108, 28, 0, 0, 0, 0,
	29, 0, 128, 27, 25, 0, 117, 0, 0, 0,
	122, 123, 124, 126, 127, 129,
}
var yyPact = []int{

	-1000, -1000, 582, -1000, -1000, -1000, -1000, -21, 190, 12,
	-1000, -1000, -1000, -48, 206, 656, 656, 656, 624, 203,
	86, -1000, -1000, 656, -5, -22, -1000, -26, -1000, -1000,
	-1000, -1000, -40, -1000, -1000, -46, 166, 164, 163, 159,
	158, 156, 155, 43, 179, 154, -36, -1000, 528, 656,
	656, 656, 656, -48, 125, -1000, -48, 21, 30, 656,
	656, 624, 202, -1000, 120, 120, -48, -1000, -27, -32,
	-1000, -1000, -1000, 11, 115, 679, 679, 679, -28, 624,
	112, 705, -51, -1000, 182, -1000, -48, -48, -48, -48,
	624, -1000, 624, -1000, -48, -48, -7, 173, 656, 99,
	-1000, -1000, -1000, -1000, 172, -1000, -1000, 38, 177, -1000,
	36, 35, 108, 107, -1000, -1000, 153, 147, 144, 143,
	176, 141, 138, 137, 136, 133, 131, 124, 87, 44,
	175, 171, 32, -1000, 468, 8, 408, 201, -1000, -48,
	23, -1000, -29, 679, 679, 656, -1000, -1000, 106, 105,
	104, 102, 624, 101, 97, 95, 91, 90, 84, 83,
	679, 75, 624, -1000, -16, -47, 195, 624, -1000, 656,
	-1000, -1000, -1000, 34, 26, -11, -1000, -1000, -1000, -1000,
	170, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 73, -1000,
	169, 189, -1000, 23, -30, -1000, 188, 187, -48, 679,
	679, -1000, -1000, -1000, 199, -1000, -1000, 23, -1000, -1000,
	20, 17, 624, 348, -1000, 288, 228, -14, -14, -1000,
	-1000, -1000, -1000, 16, -1000, 72, -14, -1000, 71, -1000,
}
var yyPgo = []int{

	0, 244, 2, 241, 4, 240, 20, 32, 239, 69,
	235, 231, 224, 224, 1, 0, 194, 223, 220, 3,
	216, 215, 12, 214, 9, 213,
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
	3, 3, 4, 3, 4, 6, 6, 4, 0, 9,
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
	-14, 29, 55, -7, -7, -6, 16, 16, 16, 16,
	-4, 16, 16, 16, 16, 16, 16, 16, -7, 16,
	-4, -25, 51, 55, 65, -18, 7, -4, -6, 29,
	29, 16, 13, 16, 13, 9, -14, 55, 9, 9,
	-7, -7, 4, -2, -14, -2, -2, 29, 29, -4,
	10, 10, 10, -19, 55, -19, 29, 16, -19, 16,
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
	57, 28, 66, 0, 0, 0, 30, 31, 32, 33,
	0, 35, 36, 37, 38, 39, 40, 41, 0, 43,
	0, 0, 15, 29, 0, 46, 0, 0, 61, 0,
	0, 20, 34, 42, 44, 3, 12, 29, 3, 3,
	0, 0, 56, 0, 13, 0, 0, 0, 0, 45,
	7, 47, 49, 0, 87, 0, 0, 19, 0, 18,
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
		{ yyVAL.tnode = NewNode(BalanceTy, yyS[yypt-1].tnode) }
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


//line mutan.y:2
package mutan
import __yyfmt__ "fmt"
//line mutan.y:2
		
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


//line mutan.y:44
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

//line mutan.y:312



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

const yyLast = 686

var yyAct = []int{

	20, 29, 4, 163, 2, 14, 59, 22, 72, 192,
	12, 82, 60, 61, 98, 73, 58, 84, 103, 193,
	35, 56, 57, 25, 206, 170, 110, 102, 69, 49,
	46, 191, 167, 104, 66, 91, 36, 65, 223, 174,
	161, 216, 164, 200, 199, 51, 21, 143, 142, 141,
	64, 70, 81, 52, 86, 87, 50, 159, 90, 225,
	217, 71, 201, 188, 187, 96, 97, 66, 66, 66,
	93, 93, 93, 68, 186, 162, 107, 107, 107, 51,
	100, 53, 111, 66, 66, 66, 185, 52, 133, 184,
	183, 135, 3, 137, 136, 182, 138, 181, 180, 106,
	108, 109, 178, 177, 139, 176, 92, 94, 95, 54,
	55, 34, 132, 37, 175, 53, 145, 144, 112, 105,
	99, 158, 44, 157, 38, 156, 41, 36, 155, 17,
	40, 39, 35, 154, 153, 152, 151, 149, 16, 42,
	43, 148, 107, 107, 107, 173, 147, 28, 146, 66,
	66, 15, 30, 179, 83, 80, 79, 31, 78, 77,
	13, 27, 76, 189, 75, 171, 172, 74, 203, 202,
	198, 194, 169, 140, 134, 129, 93, 160, 150, 101,
	45, 32, 131, 208, 207, 204, 89, 47, 197, 211,
	165, 88, 48, 63, 62, 190, 205, 26, 85, 130,
	107, 107, 210, 196, 7, 24, 23, 66, 113, 212,
	213, 33, 214, 215, 218, 6, 1, 107, 222, 0,
	0, 0, 0, 209, 107, 224, 18, 0, 19, 0,
	221, 0, 0, 0, 9, 34, 0, 37, 0, 0,
	0, 0, 0, 0, 0, 0, 44, 0, 38, 0,
	41, 36, 11, 17, 40, 39, 35, 0, 0, 5,
	0, 0, 16, 42, 43, 0, 0, 0, 0, 0,
	0, 28, 8, 0, 10, 15, 30, 0, 18, 0,
	19, 31, 220, 0, 13, 27, 9, 34, 0, 37,
	0, 0, 0, 0, 0, 0, 0, 0, 44, 0,
	38, 0, 41, 36, 11, 17, 40, 39, 35, 0,
	0, 5, 0, 0, 16, 42, 43, 0, 0, 0,
	0, 0, 0, 28, 8, 0, 10, 15, 30, 0,
	18, 0, 19, 31, 219, 0, 13, 27, 9, 34,
	0, 37, 0, 0, 0, 0, 0, 0, 0, 0,
	44, 0, 38, 0, 41, 36, 11, 17, 40, 39,
	35, 0, 0, 5, 0, 0, 16, 42, 43, 0,
	0, 0, 0, 0, 0, 28, 8, 0, 10, 15,
	30, 0, 18, 0, 19, 31, 168, 0, 13, 27,
	9, 34, 0, 37, 0, 0, 0, 0, 0, 0,
	0, 0, 44, 0, 38, 0, 41, 36, 11, 17,
	40, 39, 35, 0, 0, 5, 0, 0, 16, 42,
	43, 0, 0, 0, 0, 0, 0, 28, 8, 0,
	10, 15, 30, 0, 18, 0, 19, 31, 166, 0,
	13, 27, 9, 34, 0, 37, 0, 0, 0, 0,
	0, 0, 0, 0, 44, 0, 38, 0, 41, 36,
	11, 17, 40, 39, 35, 0, 0, 5, 0, 0,
	16, 42, 43, 0, 0, 0, 0, 0, 0, 28,
	8, 0, 10, 15, 30, 0, 18, 0, 19, 31,
	0, 0, 13, 27, 9, 34, 0, 37, 0, 0,
	0, 0, 0, 0, 0, 0, 44, 0, 38, 0,
	41, 36, 11, 17, 40, 39, 35, 0, 0, 5,
	0, 0, 16, 42, 43, 0, 0, 0, 0, 0,
	0, 28, 8, 0, 10, 15, 30, 34, 0, 37,
	0, 31, 0, 0, 13, 27, 0, 0, 44, 0,
	38, 0, 41, 36, 0, 17, 40, 39, 35, 0,
	0, 0, 0, 0, 16, 42, 43, 0, 0, 34,
	0, 37, 0, 28, 0, 0, 0, 15, 30, 0,
	44, 0, 38, 31, 41, 36, 13, 27, 40, 39,
	35, 0, 0, 0, 0, 0, 0, 42, 43, 0,
	0, 34, 0, 37, 0, 28, 0, 0, 0, 195,
	30, 0, 44, 0, 38, 31, 41, 36, 0, 27,
	40, 39, 35, 0, 0, 0, 37, 0, 0, 42,
	43, 0, 0, 0, 0, 44, 0, 38, 0, 41,
	36, 67, 30, 40, 39, 35, 128, 31, 0, 0,
	0, 27, 42, 43, 114, 116, 117, 118, 119, 123,
	0, 0, 0, 0, 67, 30, 0, 0, 0, 0,
	31, 0, 126, 0, 27, 0, 115, 0, 0, 0,
	120, 121, 122, 124, 125, 127,
}
var yyPact = []int{

	-1000, -1000, 480, -1000, -1000, 168, -1000, -1000, -25, 178,
	5, -1000, -1000, -26, -1000, 41, 480, 480, 522, 522,
	-43, -52, 190, 189, -1000, -1000, 586, -27, -4, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -45, 152, 149, 147,
	144, 143, 141, 140, 27, -51, 139, -40, -1000, -1000,
	-1000, 522, 522, 187, -1000, -1000, 177, 26, -1000, 586,
	586, 586, 522, 522, 104, 104, -1000, 167, -52, -1000,
	-1000, -28, -38, 2, 103, 609, 609, 609, -29, 522,
	102, 635, 162, -1000, 172, 96, -1000, 161, 522, -1000,
	522, -1000, -52, -1000, -52, -52, -1000, -1000, -1000, -1000,
	-1000, 522, -1000, 160, -1000, -1000, 20, -1000, 19, 18,
	101, 100, -1000, -1000, 133, 131, 126, 122, 166, 121,
	120, 119, 118, 113, 110, 108, 106, 42, 165, -1000,
	24, -1000, -1000, 13, 186, -1000, 428, 0, 376, 159,
	-30, 609, 609, -16, -1000, -1000, 98, 89, 87, 86,
	522, 82, 81, 79, 74, 73, 70, 58, 48, 47,
	522, -20, -46, -1000, -1000, 554, 181, 522, -1000, -1000,
	-1000, 15, 14, 46, -1000, -1000, -1000, -1000, -1000, 156,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 155,
	176, -1000, 13, -31, -1000, 75, -1000, 175, 174, 609,
	-16, -1000, -1000, 185, -1000, -1000, 13, -1000, -1000, 12,
	44, 522, 324, -1000, 272, 220, -16, -1000, -1000, -1000,
	-1000, -1000, 9, -16, 43, -1000,
}
var yyPgo = []int{

	0, 216, 4, 92, 2, 5, 46, 23, 215, 10,
	211, 208, 206, 205, 205, 3, 0, 181, 204, 203,
	1, 199, 198, 7, 197, 14, 195,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 21, 21, 21, 26, 26, 10, 10, 10,
	10, 10, 10, 10, 10, 14, 14, 15, 15, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 8, 19, 19, 18, 18,
	4, 4, 4, 4, 4, 4, 4, 22, 22, 5,
	5, 5, 5, 5, 5, 5, 5, 12, 12, 13,
	6, 6, 6, 6, 9, 9, 9, 9, 7, 7,
	7, 7, 7, 7, 7, 23, 20, 20, 16, 17,
	24, 25,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 9, 4,
	2, 1, 4, 5, 0, 1, 0, 3, 12, 8,
	6, 4, 4, 3, 3, 3, 0, 1, 0, 3,
	3, 3, 3, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 6, 6, 4, 0, 9, 5,
	1, 2, 1, 4, 2, 2, 0, 3, 0, 3,
	3, 6, 3, 4, 1, 1, 1, 2, 3, 5,
	1, 1, 3, 3, 2, 3, 3, 3, 1, 1,
	1, 4, 1, 1, 1, 2, 1, 1, 1, 3,
	1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -18, 52, 14,
	54, 32, -9, 64, -5, 55, 42, 33, 6, 8,
	-16, -6, -23, -12, -13, -7, -24, 65, 51, -20,
	56, 61, -17, -10, 15, 36, 31, 17, 28, 35,
	34, 30, 43, 44, 26, 12, 55, 9, -17, 55,
	15, 4, 12, 40, -3, -3, -4, -4, 59, 58,
	64, 65, 4, 4, -7, -9, -23, 55, -6, 55,
	55, 65, 12, 60, 15, 15, 15, 15, 15, 15,
	15, 25, 62, 15, 57, -22, -4, -4, 4, 9,
	32, 9, -6, -9, -6, -6, -4, -4, -25, 16,
	-25, 12, 55, 56, 31, 16, -7, -16, -7, -7,
	55, -4, 16, -11, 19, 41, 20, 21, 22, 23,
	45, 46, 47, 24, 48, 49, 37, 50, 11, 13,
	-21, 10, 16, -4, 13, -4, -2, -4, -2, -4,
	13, 29, 29, 29, 16, 16, 15, 15, 15, 15,
	12, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	12, 16, 51, -15, 29, 4, 10, 32, 10, 13,
	55, -7, -7, -20, 55, 16, 16, 16, 16, -4,
	16, 16, 16, 16, 16, 16, 16, 16, 16, -4,
	-26, 51, 55, 65, -5, 55, -19, 7, -4, 29,
	29, 16, 13, 13, 9, -15, 55, 9, 9, -7,
	-20, 4, -2, -15, -2, -2, 29, 16, -4, 10,
	10, 10, -20, 29, -20, 16,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 0, 6, 7, 0, 0,
	0, 11, 50, 0, 52, 88, 56, 56, 56, 56,
	86, 66, 79, 64, 65, 70, 0, 0, 0, 78,
	80, 82, 83, 84, 90, 87, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 10, 51,
	58, 56, 56, 0, 54, 55, 0, 0, 74, 0,
	0, 0, 56, 56, 70, 71, 79, 88, 0, 85,
	67, 0, 0, 0, 0, 0, 0, 0, 0, 56,
	0, 0, 0, 14, 0, 56, 60, 0, 56, 3,
	56, 3, 75, 71, 76, 77, 59, 62, 72, 91,
	73, 56, 68, 0, 89, 17, 0, 86, 0, 0,
	0, 0, 23, 24, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 5,
	0, 9, 53, 28, 81, 63, 0, 0, 0, 0,
	0, 0, 0, 0, 21, 22, 0, 0, 0, 0,
	56, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	56, 16, 0, 57, 27, 0, 47, 56, 49, 81,
	69, 0, 0, 0, 88, 29, 30, 31, 32, 0,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 0,
	0, 15, 28, 0, 61, 88, 45, 0, 0, 0,
	0, 20, 33, 43, 3, 12, 28, 3, 3, 0,
	0, 56, 0, 13, 0, 0, 0, 19, 44, 8,
	46, 48, 0, 0, 0, 18,
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
		//line mutan.y:65
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line mutan.y:69
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode) }
	case 3:
		//line mutan.y:70
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line mutan.y:75
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 5:
		//line mutan.y:76
		{ yyVAL.tnode = NewNode(LambdaTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 6:
		//line mutan.y:77
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 7:
		//line mutan.y:78
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 8:
		//line mutan.y:80
		{
				yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode);
				yyVAL.tnode.Constant = yyS[yypt-7].str
				yyVAL.tnode.HasRet = yyS[yypt-3].check
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-5].tnode, true)
			}
	case 9:
		//line mutan.y:86
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 10:
		//line mutan.y:87
		{ yyVAL.tnode = NewNode(ImportTy, yyS[yypt-0].tnode) }
	case 11:
		//line mutan.y:88
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 12:
		//line mutan.y:92
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 13:
		//line mutan.y:93
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-4].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str; yyVAL.tnode.Ptr = true }
	case 14:
		//line mutan.y:94
		{ yyVAL.tnode = nil }
	case 15:
		//line mutan.y:99
		{ yyVAL.check = true }
	case 16:
		//line mutan.y:100
		{ yyVAL.check = false }
	case 17:
		//line mutan.y:105
		{ yyVAL.tnode = NewNode(StopTy) }
	case 18:
		//line mutan.y:108
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 19:
		//line mutan.y:112
		{
			  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 20:
		//line mutan.y:115
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 21:
		//line mutan.y:116
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 22:
		//line mutan.y:117
		{ yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode) }
	case 23:
		//line mutan.y:118
		{ yyVAL.tnode = NewNode(PopTy) }
	case 24:
		//line mutan.y:119
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 25:
		//line mutan.y:123
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 26:
		//line mutan.y:124
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 27:
		//line mutan.y:128
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 28:
		//line mutan.y:129
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 29:
		//line mutan.y:134
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 30:
		//line mutan.y:135
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 31:
		//line mutan.y:136
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 32:
		//line mutan.y:137
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 33:
		//line mutan.y:138
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 34:
		//line mutan.y:139
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 35:
		//line mutan.y:140
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 36:
		//line mutan.y:141
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 37:
		//line mutan.y:142
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 38:
		//line mutan.y:143
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 39:
		//line mutan.y:144
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 40:
		//line mutan.y:145
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 41:
		//line mutan.y:146
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 42:
		//line mutan.y:147
		{ yyVAL.tnode = NewNode(GasTy) }
	case 43:
		//line mutan.y:148
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 44:
		//line mutan.y:150
		{
				node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 45:
		//line mutan.y:158
		{
				if yyS[yypt-0].tnode == nil {
					yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
				} else {
					yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
				}
			}
	case 46:
		//line mutan.y:168
		{
				yyVAL.tnode = yyS[yypt-1].tnode
			}
	case 47:
		//line mutan.y:171
		{ yyVAL.tnode = nil }
	case 48:
		//line mutan.y:176
		{
				yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
			}
	case 49:
		//line mutan.y:180
		{
				yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
			}
	case 50:
		//line mutan.y:186
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 51:
		//line mutan.y:187
		{ yyVAL.tnode = NewNode(RefTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 52:
		//line mutan.y:188
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 53:
		//line mutan.y:190
		{
				yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
				yyVAL.tnode.Constant = yyS[yypt-3].str
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
			}
	case 54:
		//line mutan.y:195
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 55:
		//line mutan.y:196
		{ yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode) }
	case 56:
		//line mutan.y:197
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 57:
		//line mutan.y:201
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode);}
	case 58:
		//line mutan.y:202
		{ yyVAL.tnode = nil }
	case 59:
		//line mutan.y:207
		{
				node := yyS[yypt-2].tnode
		      		yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 60:
		//line mutan.y:212
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 61:
		//line mutan.y:218
		{
				yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
			}
	case 62:
		//line mutan.y:222
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].tnode.Constant
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
			}
	case 63:
		//line mutan.y:228
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-3].str
				varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
			}
	case 64:
		//line mutan.y:234
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 65:
		//line mutan.y:235
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 66:
		//line mutan.y:236
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 67:
		//line mutan.y:241
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 68:
		//line mutan.y:246
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
				yyVAL.tnode.Ptr = true
			}
	case 69:
		//line mutan.y:255
		{
				yyVAL.tnode = NewNode(NewArrayTy)
				yyVAL.tnode.Size = yyS[yypt-2].str
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 70:
		//line mutan.y:263
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 71:
		//line mutan.y:264
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 72:
		//line mutan.y:265
		{ yyVAL.tnode = yyS[yypt-1].tnode }
	case 73:
		//line mutan.y:266
		{ yyVAL.tnode = yyS[yypt-1].tnode }
	case 74:
		//line mutan.y:271
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 75:
		//line mutan.y:273
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 76:
		//line mutan.y:274
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 77:
		//line mutan.y:275
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 78:
		//line mutan.y:279
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 79:
		//line mutan.y:280
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 80:
		//line mutan.y:281
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 81:
		//line mutan.y:282
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 82:
		//line mutan.y:283
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 83:
		//line mutan.y:284
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 84:
		//line mutan.y:285
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 85:
		//line mutan.y:289
		{ yyVAL.tnode = NewNode(DerefPtrTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 86:
		//line mutan.y:293
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 87:
		//line mutan.y:294
		{ yyVAL.tnode = NewNode(NilTy) }
	case 88:
		//line mutan.y:298
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 89:
		//line mutan.y:302
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 90:
		//line mutan.y:306
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 91:
		//line mutan.y:310
		{ yyVAL.tnode = NewNode(EmptyTy) }
	}
	goto yystack /* stack new state and value */
}


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

//line mutan.y:307



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 90
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 670

var yyAct = []int{

	32, 26, 4, 162, 133, 190, 81, 48, 72, 20,
	71, 2, 21, 49, 50, 191, 97, 13, 83, 34,
	102, 58, 59, 203, 168, 109, 101, 70, 69, 45,
	68, 189, 166, 63, 103, 35, 61, 220, 172, 160,
	213, 65, 163, 197, 196, 142, 141, 80, 140, 48,
	222, 214, 198, 186, 53, 49, 50, 91, 63, 63,
	63, 63, 54, 63, 185, 52, 84, 85, 86, 87,
	63, 89, 90, 66, 161, 106, 106, 106, 96, 99,
	22, 110, 3, 67, 63, 63, 63, 105, 107, 108,
	55, 132, 95, 184, 183, 182, 181, 137, 180, 56,
	57, 179, 63, 91, 62, 136, 178, 138, 176, 175,
	135, 174, 173, 144, 143, 94, 111, 33, 158, 36,
	104, 98, 157, 156, 155, 154, 153, 152, 43, 151,
	37, 150, 40, 35, 148, 147, 39, 38, 34, 146,
	145, 106, 106, 106, 171, 41, 42, 82, 79, 78,
	63, 63, 177, 169, 170, 77, 76, 64, 28, 75,
	74, 73, 187, 29, 200, 199, 27, 25, 139, 195,
	134, 128, 159, 149, 63, 100, 44, 30, 130, 205,
	204, 201, 192, 93, 46, 194, 208, 164, 47, 92,
	60, 51, 188, 23, 202, 88, 129, 106, 106, 207,
	193, 7, 14, 112, 31, 6, 63, 210, 12, 206,
	1, 215, 0, 209, 106, 219, 211, 212, 0, 0,
	0, 106, 221, 18, 0, 19, 0, 218, 0, 0,
	0, 9, 33, 0, 36, 0, 0, 0, 0, 0,
	0, 0, 0, 43, 0, 37, 0, 40, 35, 11,
	17, 39, 38, 34, 0, 0, 5, 0, 0, 16,
	41, 42, 0, 0, 0, 0, 0, 0, 24, 8,
	0, 10, 15, 28, 0, 18, 0, 19, 29, 217,
	0, 27, 25, 9, 33, 0, 36, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 0, 37, 0, 40,
	35, 11, 17, 39, 38, 34, 0, 0, 5, 0,
	0, 16, 41, 42, 0, 0, 0, 0, 0, 0,
	24, 8, 0, 10, 15, 28, 0, 18, 0, 19,
	29, 216, 0, 27, 25, 9, 33, 0, 36, 0,
	0, 0, 0, 0, 0, 0, 0, 43, 0, 37,
	0, 40, 35, 11, 17, 39, 38, 34, 0, 0,
	5, 0, 0, 16, 41, 42, 0, 0, 0, 0,
	0, 0, 24, 8, 0, 10, 15, 28, 0, 18,
	0, 19, 29, 167, 0, 27, 25, 9, 33, 0,
	36, 0, 0, 0, 0, 0, 0, 0, 0, 43,
	0, 37, 0, 40, 35, 11, 17, 39, 38, 34,
	0, 0, 5, 0, 0, 16, 41, 42, 0, 0,
	0, 0, 0, 0, 24, 8, 0, 10, 15, 28,
	0, 18, 0, 19, 29, 165, 0, 27, 25, 9,
	33, 0, 36, 0, 0, 0, 0, 0, 0, 0,
	0, 43, 0, 37, 0, 40, 35, 11, 17, 39,
	38, 34, 0, 0, 5, 0, 0, 16, 41, 42,
	0, 0, 0, 0, 0, 0, 24, 8, 0, 10,
	15, 28, 0, 18, 0, 19, 29, 0, 0, 27,
	25, 9, 33, 0, 36, 0, 0, 0, 0, 0,
	0, 0, 0, 43, 0, 37, 0, 40, 35, 11,
	17, 39, 38, 34, 0, 0, 5, 0, 0, 16,
	41, 42, 0, 0, 0, 0, 0, 0, 24, 8,
	0, 10, 15, 28, 33, 131, 36, 0, 29, 0,
	0, 27, 25, 0, 0, 43, 0, 37, 0, 40,
	35, 0, 17, 39, 38, 34, 0, 0, 0, 0,
	0, 16, 41, 42, 0, 0, 0, 0, 0, 0,
	24, 0, 0, 0, 15, 28, 33, 0, 36, 0,
	29, 0, 0, 27, 25, 0, 0, 43, 0, 37,
	0, 40, 35, 0, 17, 39, 38, 34, 0, 0,
	0, 0, 0, 16, 41, 42, 0, 0, 0, 0,
	36, 0, 24, 0, 0, 0, 15, 28, 0, 43,
	0, 37, 29, 40, 35, 27, 25, 39, 38, 34,
	127, 0, 0, 0, 0, 0, 41, 42, 113, 115,
	116, 117, 118, 122, 0, 0, 0, 0, 64, 28,
	0, 0, 0, 0, 29, 0, 125, 27, 25, 0,
	114, 0, 0, 0, 119, 120, 121, 123, 124, 126,
}
var yyPact = []int{

	-1000, -1000, 477, -1000, -1000, 164, -1000, -1000, -26, 175,
	4, -1000, -1000, -51, 187, 50, 477, 477, 561, 561,
	186, -1000, -1000, 102, 18, -27, -1000, -28, -1000, -1000,
	-1000, -1000, -49, -1000, -1000, -52, 146, 145, 144, 141,
	140, 134, 133, 22, -56, 132, -39, -1000, 102, 102,
	102, 102, -1000, 102, 561, 185, -1000, -1000, 174, 83,
	102, 105, 105, -1000, 163, -51, -1000, -29, -36, -1000,
	-1000, -1000, 3, 104, 593, 593, 593, -30, 561, 100,
	619, 158, -1000, 168, -51, -51, -51, -51, 519, -51,
	-9, 157, 102, -1000, 561, -1000, -51, -1000, -1000, -1000,
	561, -1000, 155, -1000, -1000, 19, -1000, 17, 16, 98,
	97, -1000, -1000, 125, 124, 120, 119, 161, 116, 114,
	112, 111, 110, 109, 108, 107, 103, 160, -1000, 23,
	-1000, -1000, 13, 183, -1000, -51, 425, 0, 373, -31,
	593, 593, -17, -1000, -1000, 96, 95, 93, 92, 561,
	90, 85, 82, 80, 79, 78, 77, 48, 37, 561,
	-20, -50, -1000, -1000, 102, 178, 561, -1000, -1000, 15,
	14, 36, -1000, -1000, -1000, -1000, -1000, 152, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 151, 172, -1000,
	13, -32, -51, -1000, 171, 170, 593, -17, -1000, -1000,
	182, -1000, -1000, 13, -1000, -1000, 11, 35, 561, 321,
	-1000, 269, 217, -17, -1000, -1000, -1000, -1000, -1000, 8,
	-17, 34, -1000,
}
var yyPgo = []int{

	0, 210, 11, 82, 2, 208, 17, 12, 205, 80,
	204, 203, 202, 202, 3, 0, 177, 201, 200, 1,
	196, 195, 9, 193, 16, 192,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 20, 20, 20, 25, 25, 10, 10, 10,
	10, 10, 10, 10, 10, 13, 13, 14, 14, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 8, 18, 18, 17, 17,
	4, 4, 4, 4, 4, 4, 4, 21, 21, 5,
	5, 5, 5, 5, 12, 12, 12, 6, 6, 6,
	6, 9, 9, 9, 9, 7, 7, 7, 7, 7,
	7, 7, 7, 22, 19, 19, 15, 16, 23, 24,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 9, 4,
	2, 1, 4, 5, 0, 1, 0, 3, 12, 8,
	6, 4, 4, 3, 3, 3, 0, 1, 0, 3,
	3, 3, 3, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 6, 6, 4, 0, 9, 5,
	1, 1, 1, 4, 2, 2, 0, 3, 0, 3,
	3, 6, 3, 4, 2, 3, 5, 1, 1, 3,
	3, 2, 3, 3, 3, 1, 1, 2, 1, 4,
	1, 1, 1, 2, 1, 1, 1, 3, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -17, 52, 14,
	54, 32, -5, -6, -12, 55, 42, 33, 6, 8,
	-22, -7, -9, -23, 51, 65, -19, 64, 56, 61,
	-16, -10, -15, 15, 36, 31, 17, 28, 35, 34,
	30, 43, 44, 26, 12, 55, 9, -16, 58, 64,
	65, 4, 15, 4, 12, 40, -3, -3, -4, -4,
	4, -7, -9, -22, 55, -6, 55, 65, 12, 55,
	55, 59, 60, 15, 15, 15, 15, 15, 15, 15,
	25, 62, 15, 57, -6, -6, -6, -6, -21, -6,
	-6, -4, 4, 9, 32, 9, -6, -24, 16, -24,
	12, 55, 56, 31, 16, -7, -15, -7, -7, 55,
	-4, 16, -11, 19, 41, 20, 21, 22, 23, 45,
	46, 47, 24, 48, 49, 37, 50, 11, 13, -20,
	10, 16, -4, 13, 13, -6, -2, -4, -2, 13,
	29, 29, 29, 16, 16, 15, 15, 15, 15, 12,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 12,
	16, 51, -14, 29, 4, 10, 32, 10, 55, -7,
	-7, -19, 55, 16, 16, 16, 16, -4, 16, 16,
	16, 16, 16, 16, 16, 16, 16, -4, -25, 51,
	55, 65, -6, -18, 7, -4, 29, 29, 16, 13,
	13, 9, -14, 55, 9, 9, -7, -19, 4, -2,
	-14, -2, -2, 29, 16, -4, 10, 10, 10, -19,
	29, -19, 16,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 0, 6, 7, 0, 0,
	0, 11, 50, 51, 52, 86, 56, 56, 56, 56,
	76, 67, 68, 0, 0, 0, 75, 0, 78, 80,
	81, 82, 84, 88, 85, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 10, 0, 0,
	0, 0, 58, 0, 56, 0, 54, 55, 0, 0,
	0, 67, 68, 76, 86, 0, 64, 0, 0, 83,
	77, 71, 0, 0, 0, 0, 0, 0, 56, 0,
	0, 0, 14, 0, 72, 73, 74, 62, 56, 60,
	0, 0, 0, 3, 56, 3, 59, 69, 89, 70,
	56, 65, 0, 87, 17, 0, 84, 0, 0, 0,
	0, 23, 24, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 5, 0,
	9, 53, 28, 0, 79, 63, 0, 0, 0, 0,
	0, 0, 0, 21, 22, 0, 0, 0, 0, 56,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 56,
	16, 0, 57, 27, 0, 47, 56, 49, 66, 0,
	0, 0, 86, 29, 30, 31, 32, 0, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 0, 0, 15,
	28, 0, 61, 45, 0, 0, 0, 0, 20, 33,
	43, 3, 12, 28, 3, 3, 0, 0, 56, 0,
	13, 0, 0, 0, 19, 44, 8, 46, 48, 0,
	0, 0, 18,
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
		{ yyVAL.tnode = yyS[yypt-0].tnode }
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
		//line mutan.y:238
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 65:
		//line mutan.y:243
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
				yyVAL.tnode.Ptr = true
			}
	case 66:
		//line mutan.y:249
		{
				yyVAL.tnode = NewNode(NewArrayTy)
				yyVAL.tnode.Size = yyS[yypt-2].str
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 67:
		//line mutan.y:257
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 68:
		//line mutan.y:258
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 69:
		//line mutan.y:259
		{ yyVAL.tnode = yyS[yypt-1].tnode }
	case 70:
		//line mutan.y:260
		{ yyVAL.tnode = yyS[yypt-1].tnode }
	case 71:
		//line mutan.y:265
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 72:
		//line mutan.y:267
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 73:
		//line mutan.y:268
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 74:
		//line mutan.y:269
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 75:
		//line mutan.y:273
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 76:
		//line mutan.y:274
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 77:
		//line mutan.y:275
		{ yyVAL.tnode = NewNode(RefTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 78:
		//line mutan.y:276
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 79:
		//line mutan.y:277
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 80:
		//line mutan.y:278
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 81:
		//line mutan.y:279
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 82:
		//line mutan.y:280
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 83:
		//line mutan.y:284
		{ yyVAL.tnode = NewNode(DerefPtrTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 84:
		//line mutan.y:288
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 85:
		//line mutan.y:289
		{ yyVAL.tnode = NewNode(NilTy) }
	case 86:
		//line mutan.y:293
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 87:
		//line mutan.y:297
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 88:
		//line mutan.y:301
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 89:
		//line mutan.y:305
		{ yyVAL.tnode = NewNode(EmptyTy) }
	}
	goto yystack /* stack new state and value */
}

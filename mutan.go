
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

const yyLast = 645

var yyAct = []int{

	32, 26, 4, 167, 2, 48, 191, 19, 98, 135,
	81, 49, 50, 72, 71, 83, 192, 34, 102, 13,
	56, 57, 21, 54, 204, 169, 110, 101, 70, 69,
	45, 54, 91, 190, 164, 52, 173, 161, 103, 35,
	221, 214, 68, 65, 168, 198, 63, 197, 143, 142,
	80, 141, 223, 215, 48, 90, 54, 54, 54, 54,
	49, 50, 22, 95, 199, 187, 54, 54, 84, 85,
	86, 87, 162, 100, 186, 107, 107, 107, 92, 93,
	94, 111, 54, 54, 54, 66, 64, 185, 3, 184,
	183, 95, 182, 133, 132, 67, 134, 105, 108, 109,
	138, 181, 180, 179, 54, 55, 177, 176, 175, 174,
	59, 33, 139, 36, 145, 144, 137, 112, 60, 104,
	99, 62, 43, 159, 37, 88, 40, 35, 62, 16,
	39, 38, 34, 158, 157, 156, 155, 154, 15, 41,
	42, 153, 107, 107, 107, 172, 61, 24, 152, 54,
	54, 20, 28, 178, 151, 149, 148, 29, 147, 146,
	27, 25, 82, 188, 170, 171, 79, 195, 78, 77,
	76, 75, 74, 73, 54, 201, 200, 140, 136, 129,
	160, 150, 88, 44, 131, 30, 196, 206, 205, 202,
	89, 46, 194, 209, 166, 203, 47, 96, 107, 107,
	208, 58, 51, 189, 23, 54, 97, 210, 211, 130,
	212, 213, 216, 193, 7, 107, 220, 17, 14, 18,
	207, 219, 107, 222, 113, 9, 33, 31, 36, 6,
	12, 1, 0, 0, 0, 0, 0, 43, 0, 37,
	0, 40, 35, 11, 16, 39, 38, 34, 0, 0,
	5, 0, 0, 15, 41, 42, 0, 0, 0, 0,
	0, 0, 24, 8, 0, 10, 20, 28, 0, 17,
	0, 18, 29, 218, 0, 27, 25, 9, 33, 0,
	36, 0, 0, 0, 0, 0, 0, 0, 0, 43,
	0, 37, 0, 40, 35, 11, 16, 39, 38, 34,
	0, 0, 5, 0, 0, 15, 41, 42, 0, 0,
	0, 0, 0, 0, 24, 8, 0, 10, 20, 28,
	0, 17, 0, 18, 29, 217, 0, 27, 25, 9,
	33, 0, 36, 0, 0, 0, 0, 0, 0, 0,
	0, 43, 0, 37, 0, 40, 35, 11, 16, 39,
	38, 34, 0, 0, 5, 0, 0, 15, 41, 42,
	0, 0, 0, 0, 0, 0, 24, 8, 0, 10,
	20, 28, 0, 17, 0, 18, 29, 165, 0, 27,
	25, 9, 33, 0, 36, 0, 0, 0, 0, 0,
	0, 0, 0, 43, 0, 37, 0, 40, 35, 11,
	16, 39, 38, 34, 0, 0, 5, 0, 0, 15,
	41, 42, 0, 0, 0, 0, 0, 0, 24, 8,
	0, 10, 20, 28, 0, 17, 0, 18, 29, 163,
	0, 27, 25, 9, 33, 0, 36, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 0, 37, 0, 40,
	35, 11, 16, 39, 38, 34, 0, 0, 5, 0,
	0, 15, 41, 42, 0, 0, 0, 0, 0, 0,
	24, 8, 0, 10, 20, 28, 0, 17, 0, 18,
	29, 0, 0, 27, 25, 9, 33, 0, 36, 0,
	0, 0, 0, 0, 0, 0, 0, 43, 0, 37,
	0, 40, 35, 11, 16, 39, 38, 34, 0, 0,
	5, 0, 0, 15, 41, 42, 0, 0, 0, 0,
	0, 0, 24, 8, 0, 10, 20, 28, 33, 0,
	36, 0, 29, 0, 0, 27, 25, 0, 0, 43,
	0, 37, 0, 40, 35, 0, 16, 39, 38, 34,
	0, 0, 0, 0, 0, 15, 41, 42, 0, 0,
	33, 0, 36, 0, 24, 0, 0, 0, 20, 28,
	0, 43, 0, 37, 29, 40, 35, 27, 25, 39,
	38, 34, 0, 0, 0, 36, 0, 0, 41, 42,
	0, 0, 0, 0, 43, 0, 37, 0, 40, 35,
	53, 28, 39, 38, 34, 128, 29, 0, 0, 27,
	25, 41, 42, 114, 116, 117, 118, 119, 123, 0,
	0, 0, 0, 106, 28, 0, 0, 0, 0, 29,
	0, 126, 27, 25, 0, 115, 0, 0, 0, 120,
	121, 122, 124, 125, 127,
}
var yyPact = []int{

	-1000, -1000, 471, -1000, -1000, 171, -1000, -1000, -25, 182,
	8, -1000, -1000, -53, 198, 545, 471, 513, 513, 197,
	106, -1000, -1000, 545, 30, -26, -1000, -27, -1000, -1000,
	-1000, -1000, -45, -1000, -1000, -47, 158, 157, 156, 155,
	154, 153, 151, 25, -52, 147, -42, -1000, 545, 545,
	545, 545, -53, 113, -1000, -1000, 181, 23, 545, 545,
	513, 193, -1000, 104, 104, -53, -1000, -28, -38, -1000,
	-1000, -1000, 7, 103, 568, 568, 568, -29, 513, 101,
	594, 166, -1000, 174, -53, -53, -53, -53, 513, -1000,
	513, -1000, -53, -53, -4, 165, 545, 96, -1000, -1000,
	-1000, -1000, 164, -1000, -1000, 22, 170, -1000, 20, 19,
	99, 98, -1000, -1000, 144, 143, 141, 140, 169, 139,
	133, 126, 122, 121, 120, 119, 118, 108, 168, -1000,
	21, -1000, 419, 2, 367, 190, -1000, -53, 15, -1000,
	-30, 568, 568, -19, -1000, -1000, 93, 92, 91, 90,
	513, 87, 86, 85, 76, 74, 73, 71, 58, 49,
	513, -18, -49, 185, 513, -1000, 545, -1000, -1000, -1000,
	18, 16, 48, -1000, -1000, -1000, -1000, -1000, 163, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 162, 180,
	-1000, 15, -31, -1000, 179, 178, -53, 568, -19, -1000,
	-1000, 189, -1000, -1000, 15, -1000, -1000, 12, 37, 513,
	315, -1000, 263, 211, -19, -1000, -1000, -1000, -1000, -1000,
	11, -19, 36, -1000,
}
var yyPgo = []int{

	0, 231, 4, 88, 2, 230, 19, 22, 229, 62,
	227, 224, 218, 218, 3, 0, 185, 214, 213, 1,
	209, 206, 7, 204, 8, 203,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 20, 20, 20, 25, 25, 10, 10, 10,
	10, 10, 10, 10, 10, 13, 13, 14, 14, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 8, 18, 18, 17, 17,
	4, 4, 4, 4, 4, 4, 21, 21, 5, 5,
	5, 5, 5, 12, 12, 12, 6, 6, 6, 6,
	6, 9, 9, 9, 9, 7, 7, 7, 7, 7,
	7, 7, 7, 22, 19, 19, 15, 16, 23, 24,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 9, 4,
	2, 1, 4, 5, 0, 1, 0, 3, 12, 8,
	6, 4, 4, 3, 3, 3, 0, 1, 0, 3,
	3, 3, 3, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 6, 6, 4, 0, 9, 5,
	1, 1, 1, 2, 2, 0, 3, 0, 3, 3,
	6, 3, 4, 2, 3, 5, 1, 1, 3, 3,
	4, 2, 3, 3, 3, 1, 1, 2, 1, 4,
	1, 1, 1, 2, 1, 1, 1, 3, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -17, 52, 14,
	54, 32, -5, -6, -12, 42, 33, 6, 8, -22,
	55, -7, -9, -23, 51, 65, -19, 64, 56, 61,
	-16, -10, -15, 15, 36, 31, 17, 28, 35, 34,
	30, 43, 44, 26, 12, 55, 9, -16, 58, 64,
	65, 4, -6, 55, -22, -3, -4, -4, 4, 4,
	12, 40, 15, -7, -9, -6, 55, 65, 12, 55,
	55, 59, 60, 15, 15, 15, 15, 15, 15, 15,
	25, 62, 15, 57, -6, -6, -6, -6, 12, 9,
	32, 9, -6, -6, -6, -4, 4, -21, -24, 16,
	-24, 55, 56, 31, 16, -7, 55, -15, -7, -7,
	55, -4, 16, -11, 19, 41, 20, 21, 22, 23,
	45, 46, 47, 24, 48, 49, 37, 50, 11, 13,
	-20, 10, -2, -4, -2, 13, 13, -6, -4, 16,
	13, 29, 29, 29, 16, 16, 15, 15, 15, 15,
	12, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	12, 16, 51, 10, 32, 10, 4, -14, 29, 55,
	-7, -7, -19, 55, 16, 16, 16, 16, -4, 16,
	16, 16, 16, 16, 16, 16, 16, 16, -4, -25,
	51, 55, 65, -18, 7, -4, -6, 29, 29, 16,
	13, 13, 9, -14, 55, 9, 9, -7, -19, 4,
	-2, -14, -2, -2, 29, 16, -4, 10, 10, 10,
	-19, 29, -19, 16,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 0, 6, 7, 0, 0,
	0, 11, 50, 51, 52, 0, 55, 55, 55, 76,
	86, 66, 67, 0, 0, 0, 75, 0, 78, 80,
	81, 82, 84, 88, 85, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 10, 0, 0,
	0, 0, 53, 86, 76, 54, 0, 0, 0, 0,
	55, 0, 57, 66, 67, 0, 63, 0, 0, 83,
	77, 71, 0, 0, 0, 0, 0, 0, 55, 0,
	0, 0, 14, 0, 72, 73, 74, 61, 55, 3,
	55, 3, 58, 59, 0, 0, 0, 55, 68, 89,
	69, 64, 0, 87, 17, 0, 86, 84, 0, 0,
	0, 0, 23, 24, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 5,
	0, 9, 0, 0, 0, 0, 79, 62, 28, 70,
	0, 0, 0, 0, 21, 22, 0, 0, 0, 0,
	55, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	55, 16, 0, 47, 55, 49, 0, 56, 27, 65,
	0, 0, 0, 86, 29, 30, 31, 32, 0, 34,
	35, 36, 37, 38, 39, 40, 41, 42, 0, 0,
	15, 28, 0, 45, 0, 0, 60, 0, 0, 20,
	33, 43, 3, 12, 28, 3, 3, 0, 0, 55,
	0, 13, 0, 0, 0, 19, 44, 8, 46, 48,
	0, 0, 0, 18,
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
		//line mutan.y:189
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 54:
		//line mutan.y:190
		{ yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode) }
	case 55:
		//line mutan.y:191
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 56:
		//line mutan.y:195
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode);}
	case 57:
		//line mutan.y:196
		{ yyVAL.tnode = nil }
	case 58:
		//line mutan.y:201
		{
				node := yyS[yypt-2].tnode
		      		yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 59:
		//line mutan.y:206
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 60:
		//line mutan.y:212
		{
				yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
			}
	case 61:
		//line mutan.y:216
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].tnode.Constant
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
			}
	case 62:
		//line mutan.y:222
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-3].str
				varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
			}
	case 63:
		//line mutan.y:232
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 64:
		//line mutan.y:237
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
				yyVAL.tnode.Ptr = true
			}
	case 65:
		//line mutan.y:243
		{
				yyVAL.tnode = NewNode(NewArrayTy)
				yyVAL.tnode.Size = yyS[yypt-2].str
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 66:
		//line mutan.y:251
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 67:
		//line mutan.y:252
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 68:
		//line mutan.y:253
		{ yyVAL.tnode = yyS[yypt-1].tnode }
	case 69:
		//line mutan.y:254
		{ yyVAL.tnode = yyS[yypt-1].tnode }
	case 70:
		//line mutan.y:256
		{
				yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
				yyVAL.tnode.Constant = yyS[yypt-3].str
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
			}
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

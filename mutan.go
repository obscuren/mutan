
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
const ID = 57396
const NUMBER = 57397
const INLINE_ASM = 57398
const OP = 57399
const DOP = 57400
const STR = 57401
const BOOLEAN = 57402
const CODE = 57403
const oper = 57404
const AND = 57405
const MUL = 57406

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
	-1, 24,
	57, 62,
	63, 62,
	64, 62,
	-2, 71,
}

const yyNprod = 86
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 534

var yyAct = []int{

	19, 27, 4, 152, 2, 13, 62, 26, 55, 72,
	32, 63, 54, 24, 56, 57, 74, 120, 34, 91,
	52, 53, 182, 160, 98, 90, 58, 41, 164, 35,
	45, 38, 33, 43, 16, 37, 36, 32, 181, 92,
	157, 214, 150, 15, 39, 40, 47, 207, 60, 153,
	76, 77, 25, 192, 48, 14, 28, 46, 61, 191,
	71, 29, 89, 156, 12, 21, 95, 95, 95, 83,
	83, 83, 99, 94, 96, 97, 34, 151, 121, 132,
	131, 123, 49, 125, 124, 41, 126, 35, 3, 38,
	33, 128, 16, 37, 36, 32, 130, 81, 34, 216,
	20, 15, 39, 40, 50, 51, 208, 41, 11, 35,
	25, 38, 33, 14, 28, 37, 36, 32, 47, 29,
	80, 193, 12, 21, 39, 40, 48, 178, 177, 176,
	159, 95, 95, 95, 163, 85, 28, 116, 161, 162,
	175, 29, 169, 174, 173, 102, 104, 105, 106, 107,
	111, 172, 179, 171, 49, 95, 82, 86, 87, 188,
	183, 170, 189, 114, 84, 84, 84, 103, 185, 168,
	167, 108, 109, 110, 112, 113, 115, 166, 165, 134,
	133, 119, 100, 93, 148, 147, 197, 146, 145, 144,
	143, 142, 95, 95, 202, 141, 140, 138, 137, 201,
	136, 204, 135, 205, 206, 73, 209, 70, 95, 213,
	17, 69, 18, 68, 212, 95, 215, 67, 9, 195,
	66, 34, 65, 64, 44, 194, 190, 129, 122, 117,
	41, 149, 35, 139, 38, 33, 10, 16, 37, 36,
	32, 127, 42, 5, 199, 198, 15, 39, 40, 17,
	196, 18, 203, 211, 79, 25, 8, 9, 14, 28,
	34, 187, 154, 88, 29, 78, 59, 12, 21, 41,
	180, 35, 75, 38, 33, 10, 16, 37, 36, 32,
	118, 186, 5, 7, 30, 15, 39, 40, 17, 23,
	18, 22, 210, 101, 25, 8, 9, 14, 28, 34,
	31, 6, 1, 29, 0, 0, 12, 21, 41, 0,
	35, 0, 38, 33, 10, 16, 37, 36, 32, 0,
	0, 5, 0, 0, 15, 39, 40, 17, 0, 18,
	0, 200, 0, 25, 8, 9, 14, 28, 34, 0,
	0, 0, 29, 0, 0, 12, 21, 41, 0, 35,
	0, 38, 33, 10, 16, 37, 36, 32, 0, 0,
	5, 0, 0, 15, 39, 40, 17, 0, 18, 0,
	158, 0, 25, 8, 9, 14, 28, 34, 0, 0,
	0, 29, 0, 0, 12, 21, 41, 0, 35, 0,
	38, 33, 10, 16, 37, 36, 32, 0, 0, 5,
	0, 0, 15, 39, 40, 17, 0, 18, 0, 155,
	0, 25, 8, 9, 14, 28, 34, 0, 0, 0,
	29, 0, 0, 12, 21, 41, 0, 35, 0, 38,
	33, 10, 16, 37, 36, 32, 0, 0, 5, 0,
	0, 15, 39, 40, 17, 0, 18, 0, 0, 0,
	25, 8, 9, 14, 28, 34, 0, 0, 0, 29,
	0, 0, 12, 21, 41, 0, 35, 0, 38, 33,
	10, 16, 37, 36, 32, 0, 0, 5, 0, 0,
	15, 39, 40, 0, 0, 0, 34, 0, 0, 25,
	8, 0, 14, 28, 0, 41, 0, 35, 29, 38,
	33, 12, 21, 37, 36, 32, 0, 0, 0, 0,
	0, 0, 39, 40, 0, 0, 0, 0, 0, 0,
	25, 0, 0, 184, 28, 0, 0, 0, 0, 29,
	0, 0, 0, 21,
}
var yyPact = []int{

	-1000, -1000, 438, -1000, -1000, 230, -1000, -1000, -21, 209,
	-1000, -1000, -24, -1000, 42, 438, 438, 59, 59, -46,
	-49, -28, 262, -1000, -1000, -6, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -48, 208, 207, 205, 202, 198, 196,
	192, 35, -52, 190, -40, -1000, -1000, 59, 59, 261,
	-1000, -1000, 245, 88, -1000, 81, 81, 81, 259, 59,
	-1000, -29, -36, 8, 167, 81, 81, 81, -30, 59,
	166, 126, 216, -1000, 165, 1, -1000, 215, 59, -1000,
	59, -1000, -49, -1000, -1000, 229, -49, -49, 59, -1000,
	-1000, 214, -1000, -1000, 67, -1000, 51, 50, 164, 163,
	-1000, -1000, 187, 185, 183, 182, 221, 181, 180, 176,
	175, 174, 173, 172, 170, 169, 219, -1000, 26, -1000,
	-1000, 20, 258, -1000, 399, 31, 360, 59, -1000, -31,
	81, 81, -26, -1000, -1000, 162, 161, 154, 153, 59,
	145, 137, 135, 128, 127, 124, 113, 112, 111, 59,
	-13, -32, -1000, -1000, 469, 254, 59, -1000, -1000, 213,
	-1000, 30, 24, 105, -1000, -1000, -1000, -1000, -1000, 212,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 206,
	241, -1000, 20, -1000, 114, -1000, -1000, 236, 235, 321,
	-1000, 81, -26, -1000, -1000, 248, -1000, -1000, -1000, -1000,
	-1000, 18, 90, 59, 282, 243, 204, -26, -1000, -1000,
	-1000, -1000, -1000, 12, -26, 83, -1000,
}
var yyPgo = []int{

	0, 302, 4, 88, 2, 5, 13, 7, 301, 108,
	300, 293, 291, 289, 289, 3, 0, 284, 283, 281,
	1, 100, 280, 272, 270,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 22, 22, 24, 24, 10, 10, 10, 10, 10,
	10, 10, 10, 14, 14, 15, 15, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 8, 19, 19, 18, 18, 18, 4,
	4, 4, 4, 4, 4, 4, 23, 23, 9, 9,
	9, 9, 21, 21, 5, 5, 5, 5, 5, 5,
	5, 5, 12, 12, 13, 6, 7, 7, 7, 7,
	7, 7, 20, 20, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 9, 4,
	1, 4, 0, 1, 0, 3, 12, 8, 6, 4,
	4, 3, 3, 3, 0, 1, 0, 3, 3, 3,
	3, 4, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 6, 6, 4, 0, 9, 7, 5, 1,
	2, 1, 4, 2, 2, 0, 3, 0, 2, 3,
	3, 3, 1, 1, 4, 3, 6, 3, 4, 1,
	1, 1, 2, 3, 5, 1, 1, 1, 4, 1,
	1, 1, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -18, 52, 14,
	32, -9, 63, -5, 54, 42, 33, 6, 8, -16,
	-21, 64, -12, -13, -6, 51, -7, -20, 55, 60,
	-17, -10, 36, 31, 17, 28, 35, 34, 30, 43,
	44, 26, 12, 54, 15, 54, 15, 4, 12, 40,
	-3, -3, -4, -4, 58, 57, 63, 64, 54, 4,
	54, 64, 12, 59, 15, 15, 15, 15, 15, 15,
	15, 25, 61, 15, 56, -23, -4, -4, 4, 9,
	32, 9, -21, -6, -9, 54, -21, -21, 4, -4,
	54, 55, 31, 16, -7, -16, -7, -7, 54, -4,
	16, -11, 19, 41, 20, 21, 22, 23, 45, 46,
	47, 24, 48, 49, 37, 50, 11, 13, -22, 16,
	16, -4, 13, -4, -2, -4, -2, 12, -4, 13,
	29, 29, 29, 16, 16, 15, 15, 15, 15, 12,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 12,
	16, 51, -15, 29, 4, 10, 32, 9, 10, -4,
	54, -7, -7, -20, 54, 16, 16, 16, 16, -4,
	16, 16, 16, 16, 16, 16, 16, 16, 16, -4,
	-24, 51, 54, -5, 54, -6, -19, 7, -4, -2,
	13, 29, 29, 16, 13, 13, 9, -15, 9, 9,
	10, -7, -20, 4, -2, -2, -2, 29, 16, -4,
	10, 10, 10, -20, 29, -20, 16,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 0, 6, 7, 0, 0,
	10, 49, 0, 51, 84, 55, 55, 55, 55, 82,
	0, 0, 69, 70, -2, 0, 75, 76, 77, 79,
	80, 81, 83, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 50, 57, 55, 55, 0,
	53, 54, 0, 0, 58, 0, 0, 0, 0, 55,
	72, 0, 0, 0, 0, 0, 0, 0, 0, 55,
	0, 0, 0, 12, 0, 55, 65, 0, 55, 3,
	55, 3, 59, 62, 63, 84, 60, 61, 55, 67,
	73, 0, 85, 15, 0, 82, 0, 0, 0, 0,
	21, 22, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 5, 0, 9,
	52, 26, 78, 68, 0, 0, 0, 55, 64, 0,
	0, 0, 0, 19, 20, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 55,
	14, 0, 56, 25, 0, 45, 55, 3, 48, 0,
	74, 0, 0, 0, 84, 27, 28, 29, 30, 0,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 0,
	0, 13, 26, 66, 84, 71, 43, 0, 0, 0,
	78, 0, 0, 18, 31, 41, 3, 11, 3, 3,
	47, 0, 0, 55, 0, 0, 0, 0, 17, 42,
	8, 44, 46, 0, 0, 0, 16,
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
	62, 63, 64,
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
		//line mutan.y:64
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line mutan.y:68
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode) }
	case 3:
		//line mutan.y:69
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line mutan.y:73
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 5:
		//line mutan.y:74
		{ yyVAL.tnode = NewNode(LambdaTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 6:
		//line mutan.y:75
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 7:
		//line mutan.y:76
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 8:
		//line mutan.y:78
		{
				yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode);
				yyVAL.tnode.Constant = yyS[yypt-7].str
				yyVAL.tnode.HasRet = yyS[yypt-3].check
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-5].tnode, true)
			}
	case 9:
		//line mutan.y:84
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 10:
		//line mutan.y:85
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 11:
		//line mutan.y:89
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 12:
		//line mutan.y:90
		{ yyVAL.tnode = nil }
	case 13:
		//line mutan.y:95
		{ yyVAL.check = true }
	case 14:
		//line mutan.y:96
		{ yyVAL.check = false }
	case 15:
		//line mutan.y:101
		{ yyVAL.tnode = NewNode(StopTy) }
	case 16:
		//line mutan.y:104
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 17:
		//line mutan.y:108
		{
			  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 18:
		//line mutan.y:111
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 19:
		//line mutan.y:112
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 20:
		//line mutan.y:113
		{ yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode) }
	case 21:
		//line mutan.y:114
		{ yyVAL.tnode = NewNode(PopTy) }
	case 22:
		//line mutan.y:115
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 23:
		//line mutan.y:119
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 24:
		//line mutan.y:120
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 25:
		//line mutan.y:124
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 26:
		//line mutan.y:125
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 27:
		//line mutan.y:130
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 28:
		//line mutan.y:131
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 29:
		//line mutan.y:132
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 30:
		//line mutan.y:133
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 31:
		//line mutan.y:134
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 32:
		//line mutan.y:135
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 33:
		//line mutan.y:136
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 34:
		//line mutan.y:137
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 35:
		//line mutan.y:138
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 36:
		//line mutan.y:139
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 37:
		//line mutan.y:140
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 38:
		//line mutan.y:141
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 39:
		//line mutan.y:142
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 40:
		//line mutan.y:143
		{ yyVAL.tnode = NewNode(GasTy) }
	case 41:
		//line mutan.y:144
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 42:
		//line mutan.y:146
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 43:
		//line mutan.y:154
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 44:
		//line mutan.y:164
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 45:
		//line mutan.y:167
		{ yyVAL.tnode = nil }
	case 46:
		//line mutan.y:172
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 47:
		//line mutan.y:177
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 48:
		//line mutan.y:182
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 49:
		//line mutan.y:188
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 50:
		//line mutan.y:189
		{ yyVAL.tnode = NewNode(RefTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 51:
		//line mutan.y:190
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 52:
		//line mutan.y:192
		{
				yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
				yyVAL.tnode.Constant = yyS[yypt-3].str
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
			}
	case 53:
		//line mutan.y:197
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 54:
		//line mutan.y:198
		{ yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode) }
	case 55:
		//line mutan.y:199
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 56:
		//line mutan.y:203
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode);}
	case 57:
		//line mutan.y:204
		{ yyVAL.tnode = nil }
	case 58:
		//line mutan.y:209
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 59:
		//line mutan.y:211
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 60:
		//line mutan.y:212
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 61:
		//line mutan.y:213
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 62:
		//line mutan.y:218
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 63:
		//line mutan.y:219
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 64:
		//line mutan.y:224
		{
				node := NewNode(SetPtrTy)
				node.Constant = yyS[yypt-2].str
		      		yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 65:
		//line mutan.y:230
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 66:
		//line mutan.y:236
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 67:
		//line mutan.y:240
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 68:
		//line mutan.y:246
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
			varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		  }
	case 69:
		//line mutan.y:252
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 70:
		//line mutan.y:253
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 71:
		//line mutan.y:254
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 72:
		//line mutan.y:259
		{
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		  }
	case 73:
		//line mutan.y:264
		{
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		  }
	case 74:
		//line mutan.y:272
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      //$$.VarType = $1
	      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
	
	}
	case 75:
		//line mutan.y:282
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 76:
		//line mutan.y:286
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 77:
		//line mutan.y:287
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 78:
		//line mutan.y:288
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 79:
		//line mutan.y:289
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 80:
		//line mutan.y:290
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 81:
		//line mutan.y:291
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 82:
		//line mutan.y:295
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 83:
		//line mutan.y:296
		{ yyVAL.tnode = NewNode(NilTy) }
	case 84:
		//line mutan.y:300
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 85:
		//line mutan.y:304
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

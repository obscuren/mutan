
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

//line mutan.y:314



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 24,
	57, 64,
	63, 64,
	64, 64,
	-2, 74,
}

const yyNprod = 89
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 566

var yyAct = []int{

	20, 28, 153, 2, 27, 122, 35, 24, 14, 13,
	4, 74, 64, 65, 56, 42, 183, 36, 21, 39,
	34, 57, 17, 38, 37, 33, 184, 58, 59, 54,
	55, 16, 40, 41, 76, 93, 201, 33, 161, 100,
	26, 92, 61, 15, 29, 46, 11, 44, 151, 30,
	182, 49, 12, 25, 62, 165, 94, 158, 77, 50,
	79, 80, 48, 219, 63, 86, 86, 86, 97, 97,
	97, 91, 96, 98, 99, 212, 85, 89, 90, 154,
	157, 195, 101, 152, 35, 84, 126, 51, 128, 123,
	194, 133, 125, 42, 127, 36, 73, 39, 34, 132,
	17, 38, 37, 33, 87, 87, 87, 49, 83, 16,
	40, 41, 131, 221, 213, 50, 196, 179, 26, 198,
	178, 15, 29, 177, 3, 176, 175, 30, 174, 173,
	12, 25, 97, 97, 97, 164, 162, 163, 172, 171,
	160, 52, 53, 51, 169, 168, 167, 166, 135, 134,
	102, 170, 95, 149, 148, 147, 97, 146, 145, 144,
	143, 180, 192, 188, 186, 185, 142, 141, 191, 139,
	138, 137, 136, 75, 72, 71, 70, 69, 68, 67,
	66, 197, 193, 130, 124, 119, 200, 150, 140, 129,
	43, 121, 203, 202, 199, 97, 97, 206, 82, 205,
	45, 190, 207, 208, 209, 47, 210, 211, 18, 155,
	19, 81, 217, 97, 218, 60, 9, 181, 214, 35,
	97, 220, 78, 120, 189, 7, 31, 23, 42, 22,
	36, 103, 39, 34, 10, 17, 38, 37, 33, 32,
	6, 5, 1, 0, 16, 40, 41, 18, 0, 19,
	0, 216, 0, 26, 8, 9, 15, 29, 35, 0,
	0, 0, 30, 0, 0, 12, 25, 42, 0, 36,
	0, 39, 34, 10, 17, 38, 37, 33, 0, 0,
	5, 0, 0, 16, 40, 41, 18, 0, 19, 0,
	215, 0, 26, 8, 9, 15, 29, 35, 0, 0,
	0, 30, 0, 0, 12, 25, 42, 0, 36, 0,
	39, 34, 10, 17, 38, 37, 33, 0, 0, 5,
	0, 0, 16, 40, 41, 18, 0, 19, 0, 204,
	0, 26, 8, 9, 15, 29, 35, 0, 0, 0,
	30, 0, 0, 12, 25, 42, 0, 36, 0, 39,
	34, 10, 17, 38, 37, 33, 0, 0, 5, 0,
	0, 16, 40, 41, 18, 0, 19, 0, 159, 0,
	26, 8, 9, 15, 29, 35, 0, 0, 0, 30,
	0, 0, 12, 25, 42, 0, 36, 0, 39, 34,
	10, 17, 38, 37, 33, 0, 0, 5, 0, 0,
	16, 40, 41, 18, 0, 19, 0, 156, 0, 26,
	8, 9, 15, 29, 35, 0, 0, 0, 30, 0,
	0, 12, 25, 42, 0, 36, 0, 39, 34, 10,
	17, 38, 37, 33, 0, 0, 5, 0, 0, 16,
	40, 41, 18, 0, 19, 0, 0, 0, 26, 8,
	9, 15, 29, 35, 0, 0, 0, 30, 0, 0,
	12, 25, 42, 0, 36, 0, 39, 34, 10, 17,
	38, 37, 33, 0, 0, 5, 0, 0, 16, 40,
	41, 0, 0, 0, 35, 0, 0, 26, 8, 0,
	15, 29, 0, 42, 0, 36, 30, 39, 34, 12,
	25, 38, 37, 33, 0, 0, 35, 0, 0, 0,
	40, 41, 0, 0, 0, 42, 0, 36, 26, 39,
	34, 187, 29, 38, 37, 33, 118, 30, 0, 0,
	0, 25, 40, 41, 104, 106, 107, 108, 109, 113,
	0, 0, 0, 88, 29, 0, 0, 0, 0, 30,
	0, 0, 116, 0, 0, 0, 105, 0, 0, 0,
	110, 111, 112, 114, 115, 117,
}
var yyPact = []int{

	-1000, -1000, 436, -1000, -1000, 178, -1000, -1000, -7, 191,
	-1000, -1000, -9, -1000, 201, 47, 436, 436, 67, 67,
	-44, -36, 211, -1000, -1000, -12, 0, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -46, 165, 164, 163, 162, 161,
	160, 159, 71, -50, 158, -22, -1000, 67, -1000, 67,
	67, 207, -1000, -1000, 189, 76, -1000, 489, 489, 489,
	67, -1000, -1000, -13, -20, 25, 136, 489, 489, 489,
	-15, 67, 134, 515, 172, -1000, 181, -1000, -11, -1000,
	171, 67, -1000, 67, -1000, -36, -1000, -1000, 177, -36,
	-36, -1000, -1000, 170, -1000, -1000, 83, -1000, 70, 62,
	133, 132, -1000, -1000, 157, 156, 155, 154, 176, 152,
	151, 145, 144, 143, 142, 140, 139, 138, 175, -1000,
	32, -1000, -1000, 50, 205, -1000, 397, 48, 358, 67,
	-16, 489, 489, 1, -1000, -1000, 131, 130, 129, 128,
	67, 123, 122, 113, 112, 110, 109, 107, 104, 101,
	67, -1, -38, -1000, -1000, 467, 194, 67, -1000, -1000,
	169, -1000, 61, 52, 100, -1000, -1000, -1000, -1000, -1000,
	168, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	106, 185, -1000, 50, -18, -1000, 201, 103, -1000, -1000,
	184, 183, 319, -1000, 489, 1, -1000, -1000, 198, -1000,
	-1000, 50, -1000, -1000, -1000, 46, 98, 67, 280, -1000,
	241, 202, 1, -1000, -1000, -1000, -1000, -1000, 34, 1,
	97, -1000,
}
var yyPgo = []int{

	0, 242, 3, 124, 10, 9, 7, 4, 240, 46,
	239, 231, 229, 227, 227, 2, 0, 226, 225, 224,
	1, 18, 223, 222, 8, 217,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 22, 22, 22, 25, 25, 10, 10, 10, 10,
	10, 10, 10, 10, 14, 14, 15, 15, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 8, 19, 19, 18, 18, 18,
	4, 4, 4, 4, 4, 4, 4, 4, 23, 23,
	9, 9, 9, 9, 21, 21, 24, 5, 5, 5,
	5, 5, 5, 5, 5, 12, 12, 13, 6, 7,
	7, 7, 7, 7, 7, 20, 20, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 9, 4,
	1, 4, 5, 0, 1, 0, 3, 12, 8, 6,
	4, 4, 3, 3, 3, 0, 1, 0, 3, 3,
	3, 3, 4, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 6, 6, 4, 0, 9, 7, 5,
	1, 2, 1, 1, 4, 2, 2, 0, 3, 0,
	2, 3, 3, 3, 1, 1, 2, 3, 3, 6,
	3, 4, 1, 1, 1, 2, 3, 5, 1, 1,
	1, 4, 1, 1, 1, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -18, 52, 14,
	32, -9, 63, -5, -24, 54, 42, 33, 6, 8,
	-16, -21, -12, -13, -6, 64, 51, -7, -20, 55,
	60, -17, -10, 36, 31, 17, 28, 35, 34, 30,
	43, 44, 26, 12, 54, 9, 54, 4, 15, 4,
	12, 40, -3, -3, -4, -4, 58, 57, 63, 64,
	4, 54, 54, 64, 12, 59, 15, 15, 15, 15,
	15, 15, 15, 25, 61, 15, 56, -4, -23, -4,
	-4, 4, 9, 32, 9, -21, -6, -9, 54, -21,
	-21, -4, 54, 55, 31, 16, -7, -16, -7, -7,
	54, -4, 16, -11, 19, 41, 20, 21, 22, 23,
	45, 46, 47, 24, 48, 49, 37, 50, 11, 13,
	-22, 10, 16, -4, 13, -4, -2, -4, -2, 12,
	13, 29, 29, 29, 16, 16, 15, 15, 15, 15,
	12, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	12, 16, 51, -15, 29, 4, 10, 32, 9, 10,
	-4, 54, -7, -7, -20, 54, 16, 16, 16, 16,
	-4, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	-4, -25, 51, 54, 64, -5, -24, 54, -6, -19,
	7, -4, -2, 13, 29, 29, 16, 13, 13, 9,
	-15, 54, 9, 9, 10, -7, -20, 4, -2, -15,
	-2, -2, 29, 16, -4, 10, 10, 10, -20, 29,
	-20, 16,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 0, 6, 7, 0, 0,
	10, 50, 0, 52, 53, 87, 57, 57, 57, 57,
	85, 0, 72, 73, -2, 0, 0, 78, 79, 80,
	82, 83, 84, 86, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 51, 57, 59, 57,
	57, 0, 55, 56, 0, 0, 60, 0, 0, 0,
	57, 66, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 57, 0, 0, 0, 13, 0, 67, 57, 68,
	0, 57, 3, 57, 3, 61, 64, 65, 87, 62,
	63, 70, 76, 0, 88, 16, 0, 85, 0, 0,
	0, 0, 22, 23, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 5,
	0, 9, 54, 27, 81, 71, 0, 0, 0, 57,
	0, 0, 0, 0, 20, 21, 0, 0, 0, 0,
	57, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	57, 15, 0, 58, 26, 0, 46, 57, 3, 49,
	0, 77, 0, 0, 0, 87, 28, 29, 30, 31,
	0, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	0, 0, 14, 27, 0, 69, 0, 87, 74, 44,
	0, 0, 0, 81, 0, 0, 19, 32, 42, 3,
	11, 27, 3, 3, 48, 0, 0, 57, 0, 12,
	0, 0, 0, 18, 43, 8, 45, 47, 0, 0,
	0, 17,
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
		//line mutan.y:65
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line mutan.y:69
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode) }
	case 3:
		//line mutan.y:70
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line mutan.y:74
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 5:
		//line mutan.y:75
		{ yyVAL.tnode = NewNode(LambdaTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 6:
		//line mutan.y:76
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 7:
		//line mutan.y:77
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 8:
		//line mutan.y:79
		{
				yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode);
				yyVAL.tnode.Constant = yyS[yypt-7].str
				yyVAL.tnode.HasRet = yyS[yypt-3].check
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-5].tnode, true)
			}
	case 9:
		//line mutan.y:85
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 10:
		//line mutan.y:86
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 11:
		//line mutan.y:90
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 12:
		//line mutan.y:91
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-4].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str; yyVAL.tnode.Ptr = true }
	case 13:
		//line mutan.y:92
		{ yyVAL.tnode = nil }
	case 14:
		//line mutan.y:97
		{ yyVAL.check = true }
	case 15:
		//line mutan.y:98
		{ yyVAL.check = false }
	case 16:
		//line mutan.y:103
		{ yyVAL.tnode = NewNode(StopTy) }
	case 17:
		//line mutan.y:106
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 18:
		//line mutan.y:110
		{
			  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 19:
		//line mutan.y:113
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 20:
		//line mutan.y:114
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 21:
		//line mutan.y:115
		{ yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode) }
	case 22:
		//line mutan.y:116
		{ yyVAL.tnode = NewNode(PopTy) }
	case 23:
		//line mutan.y:117
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 24:
		//line mutan.y:121
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 25:
		//line mutan.y:122
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 26:
		//line mutan.y:126
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 27:
		//line mutan.y:127
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 28:
		//line mutan.y:132
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 29:
		//line mutan.y:133
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 30:
		//line mutan.y:134
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 31:
		//line mutan.y:135
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 32:
		//line mutan.y:136
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 33:
		//line mutan.y:137
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 34:
		//line mutan.y:138
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 35:
		//line mutan.y:139
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 36:
		//line mutan.y:140
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 37:
		//line mutan.y:141
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 38:
		//line mutan.y:142
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 39:
		//line mutan.y:143
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 40:
		//line mutan.y:144
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 41:
		//line mutan.y:145
		{ yyVAL.tnode = NewNode(GasTy) }
	case 42:
		//line mutan.y:146
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 43:
		//line mutan.y:148
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 44:
		//line mutan.y:156
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 45:
		//line mutan.y:166
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 46:
		//line mutan.y:169
		{ yyVAL.tnode = nil }
	case 47:
		//line mutan.y:174
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 48:
		//line mutan.y:179
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 49:
		//line mutan.y:184
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 50:
		//line mutan.y:190
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 51:
		//line mutan.y:191
		{ yyVAL.tnode = NewNode(RefTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 52:
		//line mutan.y:192
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 53:
		//line mutan.y:193
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 54:
		//line mutan.y:195
		{
				yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
				yyVAL.tnode.Constant = yyS[yypt-3].str
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
			}
	case 55:
		//line mutan.y:200
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 56:
		//line mutan.y:201
		{ yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode) }
	case 57:
		//line mutan.y:202
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 58:
		//line mutan.y:206
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode);}
	case 59:
		//line mutan.y:207
		{ yyVAL.tnode = nil }
	case 60:
		//line mutan.y:212
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 61:
		//line mutan.y:214
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 62:
		//line mutan.y:215
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 63:
		//line mutan.y:216
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 64:
		//line mutan.y:221
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 65:
		//line mutan.y:222
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 66:
		//line mutan.y:226
		{ yyVAL.tnode = NewNode(DerefPtrTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 67:
		//line mutan.y:231
		{
				node := yyS[yypt-2].tnode
		      		yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 68:
		//line mutan.y:236
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 69:
		//line mutan.y:242
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 70:
		//line mutan.y:246
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 71:
		//line mutan.y:252
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
			varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		  }
	case 72:
		//line mutan.y:258
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 73:
		//line mutan.y:259
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 74:
		//line mutan.y:260
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 75:
		//line mutan.y:265
		{
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		  }
	case 76:
		//line mutan.y:270
		{
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      yyVAL.tnode.Ptr = true
		  }
	case 77:
		//line mutan.y:279
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      //$$.VarType = $1
	      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
	
	}
	case 78:
		//line mutan.y:289
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 79:
		//line mutan.y:293
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 80:
		//line mutan.y:294
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 81:
		//line mutan.y:295
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 82:
		//line mutan.y:296
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 83:
		//line mutan.y:297
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 84:
		//line mutan.y:298
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 85:
		//line mutan.y:302
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 86:
		//line mutan.y:303
		{ yyVAL.tnode = NewNode(NilTy) }
	case 87:
		//line mutan.y:307
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 88:
		//line mutan.y:311
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

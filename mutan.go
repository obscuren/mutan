
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
	-1, 25,
	58, 63,
	64, 63,
	65, 63,
	-2, 72,
}

const yyNprod = 89
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 605

var yyAct = []int{

	20, 29, 2, 14, 76, 22, 156, 58, 25, 125,
	36, 4, 28, 59, 60, 185, 67, 57, 78, 43,
	21, 37, 96, 40, 35, 186, 17, 39, 38, 34,
	55, 56, 34, 201, 163, 16, 41, 42, 103, 95,
	63, 12, 48, 45, 27, 184, 66, 160, 15, 30,
	97, 167, 154, 35, 31, 85, 218, 13, 26, 211,
	157, 195, 80, 81, 89, 89, 89, 87, 87, 87,
	100, 100, 100, 93, 94, 89, 89, 89, 84, 86,
	91, 92, 99, 101, 102, 104, 129, 155, 131, 64,
	194, 126, 136, 135, 128, 36, 130, 134, 75, 65,
	88, 88, 88, 220, 43, 212, 37, 3, 40, 35,
	50, 17, 39, 38, 34, 50, 196, 181, 51, 180,
	16, 41, 42, 51, 53, 54, 49, 124, 179, 27,
	178, 177, 176, 15, 30, 100, 100, 100, 166, 31,
	89, 89, 13, 26, 162, 175, 52, 164, 165, 174,
	173, 52, 171, 170, 169, 172, 168, 138, 137, 100,
	105, 98, 187, 152, 151, 182, 150, 189, 149, 148,
	147, 146, 192, 145, 144, 142, 141, 140, 139, 77,
	74, 73, 72, 71, 70, 69, 68, 198, 197, 193,
	133, 127, 200, 122, 153, 100, 100, 205, 143, 132,
	89, 44, 207, 32, 203, 209, 210, 204, 208, 202,
	199, 83, 100, 217, 47, 46, 191, 206, 213, 100,
	219, 18, 158, 19, 82, 216, 62, 61, 183, 9,
	79, 123, 36, 190, 7, 24, 23, 106, 33, 6,
	1, 43, 0, 37, 0, 40, 35, 11, 17, 39,
	38, 34, 121, 0, 5, 0, 0, 16, 41, 42,
	107, 109, 110, 111, 112, 116, 27, 8, 0, 10,
	15, 30, 0, 18, 0, 19, 31, 215, 119, 13,
	26, 9, 108, 0, 36, 0, 113, 114, 115, 117,
	118, 120, 0, 43, 0, 37, 0, 40, 35, 11,
	17, 39, 38, 34, 0, 0, 5, 0, 0, 16,
	41, 42, 0, 0, 0, 0, 0, 0, 27, 8,
	0, 10, 15, 30, 0, 18, 0, 19, 31, 214,
	0, 13, 26, 9, 0, 0, 36, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 0, 37, 0, 40,
	35, 11, 17, 39, 38, 34, 0, 0, 5, 0,
	0, 16, 41, 42, 0, 0, 0, 0, 0, 0,
	27, 8, 0, 10, 15, 30, 0, 18, 0, 19,
	31, 161, 0, 13, 26, 9, 0, 0, 36, 0,
	0, 0, 0, 0, 0, 0, 0, 43, 0, 37,
	0, 40, 35, 11, 17, 39, 38, 34, 0, 0,
	5, 0, 0, 16, 41, 42, 0, 0, 0, 0,
	0, 0, 27, 8, 0, 10, 15, 30, 0, 18,
	0, 19, 31, 159, 0, 13, 26, 9, 0, 0,
	36, 0, 0, 0, 0, 0, 0, 0, 0, 43,
	0, 37, 0, 40, 35, 11, 17, 39, 38, 34,
	0, 0, 5, 0, 0, 16, 41, 42, 0, 0,
	0, 0, 0, 0, 27, 8, 0, 10, 15, 30,
	0, 18, 0, 19, 31, 0, 0, 13, 26, 9,
	0, 0, 36, 0, 0, 0, 0, 0, 0, 0,
	0, 43, 0, 37, 0, 40, 35, 11, 17, 39,
	38, 34, 0, 0, 5, 0, 0, 16, 41, 42,
	0, 0, 0, 0, 36, 0, 27, 8, 0, 10,
	15, 30, 0, 43, 0, 37, 31, 40, 35, 13,
	26, 39, 38, 34, 0, 0, 0, 0, 0, 0,
	41, 42, 0, 0, 0, 0, 36, 0, 27, 0,
	0, 0, 188, 30, 0, 43, 0, 37, 31, 40,
	35, 0, 26, 39, 38, 34, 0, 0, 0, 0,
	0, 0, 41, 42, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 90, 30, 0, 0, 0, 0,
	31, 0, 0, 0, 26,
}
var yyPact = []int{

	-1000, -1000, 475, -1000, -1000, 189, -1000, -1000, -12, 206,
	22, -1000, -1000, -13, -1000, 111, 475, 475, 78, 78,
	-42, -51, 223, 222, -1000, -1000, -15, 34, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -44, 171, 170, 169, 168,
	167, 166, 165, 73, -58, 164, -39, -1000, -1000, -1000,
	78, 78, 220, -1000, -1000, 202, 46, -1000, 539, 539,
	539, 78, 78, -1000, -1000, -16, -34, 19, 145, 539,
	539, 539, -17, 78, 144, 241, 180, -1000, 117, -7,
	-1000, 178, 78, -1000, 78, -1000, -51, -1000, -1000, -1000,
	187, -51, -51, -1000, -1000, -1000, 177, -1000, -1000, 68,
	-1000, 64, 63, 142, 141, -1000, -1000, 163, 162, 161,
	160, 186, 159, 158, 156, 155, 154, 153, 151, 149,
	148, 182, -1000, 36, -1000, -1000, 31, 218, -1000, 423,
	15, 371, 78, -21, 539, 539, -4, -1000, -1000, 140,
	138, 137, 136, 78, 134, 133, 129, 116, 115, 114,
	112, 103, 101, 78, -6, -40, -1000, -1000, 507, 209,
	78, -1000, 176, -1000, 61, 32, 100, -1000, -1000, -1000,
	-1000, -1000, 175, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 174, 201, -1000, 31, -22, -1000, 106, -1000,
	-1000, 200, 195, -1000, 539, -4, -1000, -1000, 213, -1000,
	-1000, 31, -1000, -1000, 30, 89, 78, 319, -1000, 267,
	215, -4, -1000, -1000, -1000, -1000, -1000, 27, -4, 87,
	-1000,
}
var yyPgo = []int{

	0, 240, 2, 107, 11, 3, 8, 12, 239, 41,
	238, 237, 236, 235, 235, 6, 0, 203, 234, 233,
	1, 20, 231, 230, 5, 228,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 22, 22, 22, 25, 25, 10, 10, 10,
	10, 10, 10, 10, 10, 14, 14, 15, 15, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 8, 19, 19, 18, 18,
	4, 4, 4, 4, 4, 4, 4, 23, 23, 9,
	9, 9, 9, 21, 21, 5, 5, 5, 5, 5,
	5, 5, 5, 12, 12, 13, 6, 7, 7, 7,
	7, 7, 7, 7, 24, 20, 20, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 9, 4,
	2, 1, 4, 5, 0, 1, 0, 3, 12, 8,
	6, 4, 4, 3, 3, 3, 0, 1, 0, 3,
	3, 3, 3, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 6, 6, 4, 0, 9, 5,
	1, 2, 1, 4, 2, 2, 0, 3, 0, 2,
	3, 3, 3, 1, 1, 3, 3, 6, 3, 4,
	1, 1, 1, 2, 3, 5, 1, 1, 1, 1,
	4, 1, 1, 1, 2, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -18, 52, 14,
	54, 32, -9, 64, -5, 55, 42, 33, 6, 8,
	-16, -21, -24, -12, -13, -6, 65, 51, -7, -20,
	56, 61, -17, -10, 36, 31, 17, 28, 35, 34,
	30, 43, 44, 26, 12, 55, 9, -17, 55, 15,
	4, 12, 40, -3, -3, -4, -4, 59, 58, 64,
	65, 4, 4, 55, 55, 65, 12, 60, 15, 15,
	15, 15, 15, 15, 15, 25, 62, 15, 57, -23,
	-4, -4, 4, 9, 32, 9, -21, -6, -9, -24,
	55, -21, -21, -4, -4, 55, 56, 31, 16, -7,
	-16, -7, -7, 55, -4, 16, -11, 19, 41, 20,
	21, 22, 23, 45, 46, 47, 24, 48, 49, 37,
	50, 11, 13, -22, 10, 16, -4, 13, -4, -2,
	-4, -2, 12, 13, 29, 29, 29, 16, 16, 15,
	15, 15, 15, 12, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 12, 16, 51, -15, 29, 4, 10,
	32, 10, -4, 55, -7, -7, -20, 55, 16, 16,
	16, 16, -4, 16, 16, 16, 16, 16, 16, 16,
	16, 16, -4, -25, 51, 55, 65, -5, 55, -6,
	-19, 7, -4, 13, 29, 29, 16, 13, 13, 9,
	-15, 55, 9, 9, -7, -20, 4, -2, -15, -2,
	-2, 29, 16, -4, 10, 10, 10, -20, 29, -20,
	16,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 0, 6, 7, 0, 0,
	0, 11, 50, 0, 52, 87, 56, 56, 56, 56,
	85, 0, 78, 70, 71, -2, 0, 0, 76, 77,
	79, 81, 82, 83, 86, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 10, 51, 58,
	56, 56, 0, 54, 55, 0, 0, 59, 0, 0,
	0, 56, 56, 84, 73, 0, 0, 0, 0, 0,
	0, 0, 0, 56, 0, 0, 0, 14, 0, 56,
	66, 0, 56, 3, 56, 3, 60, 63, 64, 78,
	87, 61, 62, 65, 68, 74, 0, 88, 17, 0,
	85, 0, 0, 0, 0, 23, 24, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 5, 0, 9, 53, 28, 80, 69, 0,
	0, 0, 56, 0, 0, 0, 0, 21, 22, 0,
	0, 0, 0, 56, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 56, 16, 0, 57, 27, 0, 47,
	56, 49, 0, 75, 0, 0, 0, 87, 29, 30,
	31, 32, 0, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 0, 0, 15, 28, 0, 67, 87, 72,
	45, 0, 0, 80, 0, 0, 20, 33, 43, 3,
	12, 28, 3, 3, 0, 0, 56, 0, 13, 0,
	0, 0, 19, 44, 8, 46, 48, 0, 0, 0,
	18,
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
		{ yyVAL.tnode = NewNode(ImportTy, yyS[yypt-0].tnode) }
	case 11:
		//line mutan.y:87
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 12:
		//line mutan.y:91
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 13:
		//line mutan.y:92
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-4].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str; yyVAL.tnode.Ptr = true }
	case 14:
		//line mutan.y:93
		{ yyVAL.tnode = nil }
	case 15:
		//line mutan.y:98
		{ yyVAL.check = true }
	case 16:
		//line mutan.y:99
		{ yyVAL.check = false }
	case 17:
		//line mutan.y:104
		{ yyVAL.tnode = NewNode(StopTy) }
	case 18:
		//line mutan.y:107
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 19:
		//line mutan.y:111
		{
			  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 20:
		//line mutan.y:114
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 21:
		//line mutan.y:115
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 22:
		//line mutan.y:116
		{ yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode) }
	case 23:
		//line mutan.y:117
		{ yyVAL.tnode = NewNode(PopTy) }
	case 24:
		//line mutan.y:118
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 25:
		//line mutan.y:122
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 26:
		//line mutan.y:123
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 27:
		//line mutan.y:127
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 28:
		//line mutan.y:128
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 29:
		//line mutan.y:133
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 30:
		//line mutan.y:134
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 31:
		//line mutan.y:135
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 32:
		//line mutan.y:136
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 33:
		//line mutan.y:137
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 34:
		//line mutan.y:138
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 35:
		//line mutan.y:139
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 36:
		//line mutan.y:140
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 37:
		//line mutan.y:141
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 38:
		//line mutan.y:142
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 39:
		//line mutan.y:143
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 40:
		//line mutan.y:144
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 41:
		//line mutan.y:145
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 42:
		//line mutan.y:146
		{ yyVAL.tnode = NewNode(GasTy) }
	case 43:
		//line mutan.y:147
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 44:
		//line mutan.y:149
		{
				node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 45:
		//line mutan.y:157
		{
				if yyS[yypt-0].tnode == nil {
					yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
				} else {
					yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
				}
			}
	case 46:
		//line mutan.y:167
		{
				yyVAL.tnode = yyS[yypt-1].tnode
			}
	case 47:
		//line mutan.y:170
		{ yyVAL.tnode = nil }
	case 48:
		//line mutan.y:175
		{
				yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
			}
	case 49:
		//line mutan.y:179
		{
				yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
			}
	case 50:
		//line mutan.y:185
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 51:
		//line mutan.y:186
		{ yyVAL.tnode = NewNode(RefTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 52:
		//line mutan.y:187
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 53:
		//line mutan.y:189
		{
				yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
				yyVAL.tnode.Constant = yyS[yypt-3].str
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
			}
	case 54:
		//line mutan.y:194
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 55:
		//line mutan.y:195
		{ yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode) }
	case 56:
		//line mutan.y:196
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 57:
		//line mutan.y:200
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode);}
	case 58:
		//line mutan.y:201
		{ yyVAL.tnode = nil }
	case 59:
		//line mutan.y:206
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 60:
		//line mutan.y:208
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 61:
		//line mutan.y:209
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 62:
		//line mutan.y:210
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 63:
		//line mutan.y:215
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 64:
		//line mutan.y:216
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 65:
		//line mutan.y:221
		{
				node := yyS[yypt-2].tnode
		      		yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 66:
		//line mutan.y:226
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 67:
		//line mutan.y:232
		{
				yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
			}
	case 68:
		//line mutan.y:236
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].tnode.Constant
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
			}
	case 69:
		//line mutan.y:242
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-3].str
				varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
			}
	case 70:
		//line mutan.y:248
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 71:
		//line mutan.y:249
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 72:
		//line mutan.y:250
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 73:
		//line mutan.y:255
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 74:
		//line mutan.y:260
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
				yyVAL.tnode.Ptr = true
			}
	case 75:
		//line mutan.y:269
		{
				yyVAL.tnode = NewNode(NewArrayTy)
				yyVAL.tnode.Size = yyS[yypt-2].str
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 76:
		//line mutan.y:277
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 77:
		//line mutan.y:281
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 78:
		//line mutan.y:282
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 79:
		//line mutan.y:283
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 80:
		//line mutan.y:284
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 81:
		//line mutan.y:285
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 82:
		//line mutan.y:286
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 83:
		//line mutan.y:287
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 84:
		//line mutan.y:291
		{ yyVAL.tnode = NewNode(DerefPtrTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 85:
		//line mutan.y:295
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 86:
		//line mutan.y:296
		{ yyVAL.tnode = NewNode(NilTy) }
	case 87:
		//line mutan.y:300
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 88:
		//line mutan.y:304
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

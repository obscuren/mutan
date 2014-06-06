
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
const DIFFICULTY = 57385
const PREVHASH = 57386
const TIMESTAMP = 57387
const BLOCKNUM = 57388
const COINBASE = 57389
const GAS = 57390
const VAR = 57391
const FUNC = 57392
const FUNC_CALL = 57393
const ID = 57394
const NUMBER = 57395
const INLINE_ASM = 57396
const OP = 57397
const DOP = 57398
const TYPE = 57399
const STR = 57400
const BOOLEAN = 57401
const CODE = 57402

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
	"TYPE",
	"STR",
	"BOOLEAN",
	"CODE",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line mutan.y:291



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 52,
	-1, 62,
	55, 47,
	56, 47,
	-2, 58,
}

const yyNprod = 79
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 475

var yyAct = []int{

	23, 4, 132, 2, 22, 12, 64, 174, 133, 137,
	53, 102, 31, 38, 37, 173, 66, 178, 48, 49,
	73, 36, 29, 32, 71, 35, 30, 75, 15, 34,
	33, 29, 136, 52, 38, 37, 162, 14, 144, 170,
	63, 139, 82, 72, 21, 68, 69, 13, 24, 38,
	37, 104, 74, 40, 25, 38, 37, 38, 37, 130,
	78, 80, 81, 38, 37, 161, 38, 37, 11, 103,
	38, 37, 105, 51, 107, 106, 16, 108, 17, 20,
	191, 38, 37, 76, 9, 193, 186, 31, 133, 171,
	169, 43, 131, 38, 37, 113, 36, 112, 32, 44,
	35, 30, 10, 15, 34, 33, 29, 62, 110, 5,
	59, 195, 14, 141, 143, 140, 187, 142, 61, 21,
	8, 149, 13, 24, 16, 43, 17, 45, 190, 25,
	172, 159, 9, 44, 128, 31, 42, 127, 167, 158,
	163, 168, 157, 3, 36, 156, 32, 155, 35, 30,
	10, 15, 34, 33, 29, 154, 153, 5, 46, 47,
	14, 45, 152, 151, 150, 176, 148, 21, 8, 147,
	13, 24, 181, 146, 180, 145, 114, 25, 101, 183,
	77, 184, 185, 126, 188, 125, 124, 192, 16, 123,
	17, 122, 189, 121, 194, 120, 9, 109, 118, 31,
	117, 116, 115, 65, 58, 57, 56, 55, 36, 54,
	32, 41, 35, 30, 10, 15, 34, 33, 29, 99,
	129, 5, 119, 111, 14, 16, 39, 17, 182, 179,
	177, 21, 8, 9, 13, 24, 31, 175, 166, 134,
	70, 25, 50, 160, 67, 36, 100, 32, 60, 35,
	30, 10, 15, 34, 33, 29, 165, 7, 5, 26,
	28, 14, 16, 19, 17, 18, 138, 83, 21, 8,
	9, 13, 24, 31, 27, 6, 1, 0, 25, 0,
	0, 0, 36, 0, 32, 0, 35, 30, 10, 15,
	34, 33, 29, 0, 0, 5, 0, 0, 14, 16,
	0, 17, 0, 135, 0, 21, 8, 9, 13, 24,
	31, 0, 0, 0, 0, 25, 0, 0, 0, 36,
	0, 32, 0, 35, 30, 10, 15, 34, 33, 29,
	0, 0, 5, 0, 0, 14, 16, 0, 17, 0,
	0, 0, 21, 8, 9, 13, 24, 31, 0, 0,
	0, 0, 25, 0, 0, 0, 36, 0, 32, 0,
	35, 30, 10, 15, 34, 33, 29, 0, 0, 5,
	0, 0, 14, 0, 0, 0, 0, 31, 0, 21,
	8, 0, 13, 24, 0, 0, 36, 0, 32, 25,
	35, 30, 0, 15, 34, 33, 29, 31, 0, 0,
	0, 0, 14, 0, 0, 0, 36, 0, 32, 21,
	35, 30, 13, 24, 34, 33, 29, 31, 0, 25,
	0, 0, 0, 0, 0, 0, 36, 0, 32, 21,
	35, 30, 164, 24, 34, 33, 29, 98, 0, 25,
	0, 0, 0, 0, 0, 84, 86, 87, 88, 89,
	93, 0, 79, 24, 0, 0, 0, 0, 0, 25,
	0, 0, 0, 96, 0, 0, 0, 85, 0, 90,
	91, 92, 94, 95, 97,
}
var yyPact = []int{

	-1000, -1000, 330, -1000, -42, 214, -1000, -1000, 1, 196,
	-1000, -1000, -1000, 121, 330, 330, 360, 360, 238, -1000,
	-1000, 21, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-48, 194, 192, 191, 190, 189, 85, -1000, 360, -54,
	188, -38, -1000, 360, 360, 236, -1000, -1000, 15, 11,
	360, -1000, -26, 52, 164, 400, 400, 400, -10, 426,
	-1000, -1000, -1000, -42, 206, -1000, 162, -5, -42, 38,
	360, -1000, 360, -1000, -42, 184, -1000, -1000, 79, 211,
	68, 66, 160, -1000, 187, 186, 185, 183, 210, 180,
	178, 176, 174, 171, 170, 168, 122, 119, 208, -1000,
	43, -1000, -1000, -21, 235, -42, 293, 0, 256, -11,
	400, 360, 400, -14, -1000, 159, 157, 153, 150, 360,
	148, 147, 146, 140, 139, 131, 129, 126, 123, 360,
	16, -16, -1000, -1000, 380, 231, 360, -1000, -1000, -1000,
	61, 26, 60, 114, -1000, -1000, -1000, -1000, -1000, 2,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -6,
	228, -1000, 59, -1000, 87, -1000, 221, 8, 219, 400,
	-1000, -14, -1000, -1000, 224, -1000, -1000, -1000, -1000, -1000,
	57, 100, 360, 182, 118, 70, -14, -1000, -42, -1000,
	-1000, -1000, 56, -14, 95, -1000,
}
var yyPgo = []int{

	0, 276, 3, 143, 1, 5, 79, 4, 275, 68,
	274, 267, 265, 263, 263, 2, 260, 259, 257, 256,
	0, 248, 246, 244, 243,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 22, 22, 24, 24, 10, 10, 10, 10, 10,
	10, 14, 14, 15, 15, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 8, 19, 19, 18, 18, 18, 4, 4, 4,
	4, 4, 4, 23, 23, 9, 9, 21, 21, 5,
	5, 5, 5, 5, 5, 5, 12, 13, 6, 7,
	7, 7, 7, 7, 7, 20, 20, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 9, 4,
	1, 4, 0, 1, 0, 3, 12, 8, 6, 4,
	3, 3, 0, 1, 0, 3, 3, 3, 3, 4,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	6, 6, 4, 0, 9, 7, 5, 1, 1, 4,
	2, 2, 0, 3, 0, 2, 3, 1, 1, 3,
	6, 3, 4, 1, 1, 1, 2, 5, 1, 1,
	1, 4, 1, 1, 1, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -18, 50, 14,
	32, -9, -5, 52, 42, 33, 6, 8, -12, -13,
	-6, 49, -7, -20, 53, 59, -17, -10, -16, 36,
	31, 17, 28, 35, 34, 30, 26, 56, 55, 12,
	52, 15, 15, 4, 12, 40, -3, -3, -4, -4,
	4, 52, 12, 58, 15, 15, 15, 15, 15, 25,
	-21, -6, -9, -4, 60, 15, 54, -23, -4, -4,
	4, 9, 32, 9, -4, 53, 31, 16, -7, 52,
	-7, -7, 52, -11, 19, 41, 20, 21, 22, 23,
	43, 44, 45, 24, 46, 47, 37, 48, 11, 13,
	-22, 16, 16, -4, 13, -4, -2, -4, -2, 13,
	29, 12, 29, 29, 16, 15, 15, 15, 15, 12,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 12,
	16, 49, -15, 29, 4, 10, 32, 9, 10, 52,
	-7, -4, -7, -20, 52, 16, 16, 16, 16, -4,
	16, 16, 16, 16, 16, 16, 16, 16, 16, -4,
	-24, 49, 52, -5, 52, -19, 7, -4, -2, 29,
	13, 29, 16, 13, 13, 9, -15, 9, 9, 10,
	-7, -20, 4, -2, -2, -2, 29, 16, -4, 10,
	10, 10, -20, 29, -20, 16,
}
var yyDef = []int{

	3, -2, -2, 2, 4, 0, 6, 7, 0, 0,
	10, 47, 48, 77, 52, 52, 52, 52, 63, 64,
	65, 0, 68, 69, 70, 72, 73, 74, 75, 76,
	0, 0, 0, 0, 0, 0, 0, 55, 52, 0,
	0, 0, 54, 52, 52, 0, 50, 51, 0, 0,
	52, 66, 0, 0, 0, 0, 0, 0, 0, 0,
	56, 57, -2, 0, 0, 12, 0, 52, 59, 0,
	52, 3, 52, 3, 61, 0, 78, 15, 0, 77,
	0, 0, 0, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 5,
	0, 9, 49, 24, 71, 62, 52, 0, 52, 0,
	0, 52, 0, 0, 19, 0, 0, 0, 0, 52,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 52,
	14, 0, 53, 23, 0, 43, 52, 3, 46, 67,
	0, 0, 0, 0, 77, 25, 26, 27, 28, 0,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 0,
	0, 13, 24, 60, 77, 41, 0, 0, 52, 0,
	71, 0, 18, 29, 39, 3, 11, 3, 3, 45,
	0, 0, 52, 52, 52, 52, 0, 17, 40, 8,
	42, 44, 0, 0, 0, 16,
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
	52, 53, 54, 55, 56, 57, 58, 59, 60,
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
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 21:
		//line mutan.y:117
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 22:
		//line mutan.y:118
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 23:
		//line mutan.y:122
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 24:
		//line mutan.y:123
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 25:
		//line mutan.y:128
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 26:
		//line mutan.y:129
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 27:
		//line mutan.y:130
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 28:
		//line mutan.y:131
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 29:
		//line mutan.y:132
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 30:
		//line mutan.y:133
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 31:
		//line mutan.y:134
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 32:
		//line mutan.y:135
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 33:
		//line mutan.y:136
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 34:
		//line mutan.y:137
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 35:
		//line mutan.y:138
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 36:
		//line mutan.y:139
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 37:
		//line mutan.y:140
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 38:
		//line mutan.y:141
		{ yyVAL.tnode = NewNode(GasTy) }
	case 39:
		//line mutan.y:142
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 40:
		//line mutan.y:144
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 41:
		//line mutan.y:152
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 42:
		//line mutan.y:162
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 43:
		//line mutan.y:165
		{ yyVAL.tnode = nil }
	case 44:
		//line mutan.y:170
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 45:
		//line mutan.y:175
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 46:
		//line mutan.y:180
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 47:
		//line mutan.y:186
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 48:
		//line mutan.y:187
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 49:
		//line mutan.y:189
		{
				yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
				yyVAL.tnode.Constant = yyS[yypt-3].str
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
			}
	case 50:
		//line mutan.y:194
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 51:
		//line mutan.y:195
		{ yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode) }
	case 52:
		//line mutan.y:196
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 53:
		//line mutan.y:200
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode);}
	case 54:
		//line mutan.y:201
		{ yyVAL.tnode = nil }
	case 55:
		//line mutan.y:206
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 56:
		//line mutan.y:208
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 57:
		//line mutan.y:212
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 58:
		//line mutan.y:213
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 59:
		//line mutan.y:218
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 60:
		//line mutan.y:224
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 61:
		//line mutan.y:228
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 62:
		//line mutan.y:234
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
			varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		  }
	case 63:
		//line mutan.y:240
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 64:
		//line mutan.y:241
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 65:
		//line mutan.y:242
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 66:
		//line mutan.y:247
		{
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      //$$.VarType = $1
	  }
	case 67:
		//line mutan.y:256
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      //$$.VarType = $1
	      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
	
	}
	case 68:
		//line mutan.y:266
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 69:
		//line mutan.y:270
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 70:
		//line mutan.y:271
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 71:
		//line mutan.y:272
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 72:
		//line mutan.y:273
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 73:
		//line mutan.y:274
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 74:
		//line mutan.y:275
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 75:
		//line mutan.y:279
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 76:
		//line mutan.y:280
		{ yyVAL.tnode = NewNode(NilTy) }
	case 77:
		//line mutan.y:284
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 78:
		//line mutan.y:288
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

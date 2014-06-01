
//line mutan.y:2
package mutan
import __yyfmt__ "fmt"
//line mutan.y:2
		
import (
	_"fmt"
)

var Tree *SyntaxTree


//line mutan.y:12
type yySymType struct {
	yys int
	num int
	str string
	tnode *SyntaxTree
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
const RETURN = 57375
const CREATE = 57376
const TRANSACT = 57377
const NIL = 57378
const BALANCE = 57379
const VAR_ASSIGN = 57380
const LAMBDA = 57381
const COLON = 57382
const ADDRESS = 57383
const DIFFICULTY = 57384
const PREVHASH = 57385
const TIMESTAMP = 57386
const BLOCKNUM = 57387
const COINBASE = 57388
const GAS = 57389
const VAR = 57390
const FUNC = 57391
const FUNC_CALL = 57392
const ID = 57393
const NUMBER = 57394
const INLINE_ASM = 57395
const OP = 57396
const DOP = 57397
const TYPE = 57398
const STR = 57399
const BOOLEAN = 57400
const CODE = 57401

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
	"RETURN",
	"CREATE",
	"TRANSACT",
	"NIL",
	"BALANCE",
	"VAR_ASSIGN",
	"LAMBDA",
	"COLON",
	"ADDRESS",
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

//line mutan.y:227



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 46,
	-1, 59,
	54, 43,
	55, 43,
	-2, 50,
}

const yyNprod = 71
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 452

var yyAct = []int{

	22, 4, 12, 61, 21, 50, 37, 36, 63, 2,
	30, 71, 130, 78, 39, 28, 42, 43, 11, 35,
	127, 31, 19, 34, 29, 72, 13, 33, 32, 28,
	135, 49, 178, 172, 159, 157, 107, 106, 104, 60,
	56, 20, 165, 126, 16, 23, 67, 68, 162, 70,
	180, 24, 161, 14, 173, 15, 59, 74, 76, 77,
	58, 9, 158, 122, 30, 37, 36, 99, 66, 160,
	48, 102, 64, 35, 98, 31, 100, 34, 29, 10,
	13, 33, 32, 28, 44, 101, 5, 37, 36, 37,
	36, 65, 45, 37, 36, 20, 8, 149, 16, 23,
	148, 147, 146, 37, 36, 24, 145, 132, 134, 131,
	144, 133, 143, 37, 36, 140, 142, 37, 36, 141,
	46, 139, 138, 137, 136, 150, 37, 36, 154, 108,
	97, 96, 156, 73, 151, 121, 120, 155, 119, 14,
	118, 15, 117, 176, 116, 115, 114, 9, 103, 112,
	30, 111, 110, 109, 62, 55, 54, 53, 52, 35,
	168, 31, 167, 34, 29, 10, 13, 33, 32, 28,
	51, 174, 5, 177, 170, 171, 40, 95, 3, 179,
	123, 20, 8, 113, 16, 23, 14, 105, 15, 38,
	175, 24, 41, 164, 9, 124, 153, 30, 169, 129,
	69, 47, 57, 152, 7, 25, 35, 27, 31, 18,
	34, 29, 10, 13, 33, 32, 28, 17, 79, 5,
	26, 6, 14, 1, 15, 0, 166, 0, 20, 8,
	9, 16, 23, 30, 0, 0, 0, 0, 24, 0,
	0, 0, 35, 0, 31, 0, 34, 29, 10, 13,
	33, 32, 28, 0, 0, 5, 0, 0, 14, 0,
	15, 0, 163, 0, 20, 8, 9, 16, 23, 30,
	0, 0, 0, 0, 24, 0, 0, 0, 35, 0,
	31, 0, 34, 29, 10, 13, 33, 32, 28, 0,
	0, 5, 0, 0, 14, 0, 15, 0, 128, 0,
	20, 8, 9, 16, 23, 30, 0, 0, 0, 0,
	24, 0, 0, 0, 35, 0, 31, 0, 34, 29,
	10, 13, 33, 32, 28, 0, 0, 5, 0, 0,
	14, 0, 15, 0, 125, 0, 20, 8, 9, 16,
	23, 30, 0, 0, 0, 0, 24, 0, 0, 0,
	35, 0, 31, 0, 34, 29, 10, 13, 33, 32,
	28, 0, 0, 5, 0, 0, 0, 0, 0, 0,
	30, 0, 20, 8, 0, 16, 23, 0, 0, 35,
	0, 31, 24, 34, 29, 0, 0, 33, 32, 28,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 30,
	0, 20, 0, 0, 16, 23, 0, 0, 35, 0,
	31, 24, 34, 29, 0, 94, 33, 32, 28, 0,
	0, 0, 0, 80, 82, 83, 84, 85, 89, 0,
	0, 0, 0, 75, 23, 0, 0, 0, 0, 0,
	24, 92, 0, 0, 0, 81, 86, 87, 88, 90,
	91, 93,
}
var yyPact = []int{

	-1000, -1000, 47, -1000, -48, 177, -1000, -1000, -37, 161,
	-1000, -1000, -1000, 47, -7, -7, 80, 197, -1000, -1000,
	19, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -52,
	155, 143, 142, 141, 140, 15, -1000, -7, -56, 139,
	-45, -1000, 63, 59, -7, -7, 196, -7, -1000, -41,
	-6, 117, 382, 382, 382, -38, 404, -1000, -1000, -1000,
	-48, 164, 115, 114, -1000, -7, -1000, -48, 72, -7,
	-48, 135, -1000, -1000, 9, 175, 8, 7, 113, -1000,
	138, 137, 136, 134, 171, 131, 130, 129, 127, 125,
	123, 121, 120, 48, 168, -1000, 186, -1000, 324, 11,
	288, 195, -48, -39, 382, -7, 382, -21, -1000, 108,
	107, 106, 105, -7, 103, 100, 96, 94, 90, 86,
	85, 84, 81, -7, -1000, 189, -7, -1000, -1000, 353,
	-1000, 6, 49, 5, 53, -1000, -1000, -1000, -1000, -1000,
	39, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	35, 252, -1000, 184, 33, 216, -1000, 382, -1000, -21,
	-1000, -1000, 194, -1000, -1000, -1000, -1000, 4, 38, -7,
	180, 133, -21, -1000, -48, -1000, -1000, 3, -21, 34,
	-1000,
}
var yyPgo = []int{

	0, 223, 9, 178, 1, 2, 22, 4, 221, 18,
	220, 218, 217, 209, 209, 209, 207, 205, 204, 203,
	0, 202,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 10, 10, 10, 10, 10, 10, 14, 14, 15,
	15, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 8, 19, 19,
	18, 18, 18, 4, 4, 4, 4, 9, 9, 21,
	21, 5, 5, 5, 5, 5, 5, 5, 12, 13,
	6, 7, 7, 7, 7, 7, 7, 20, 20, 16,
	17,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 7, 4,
	1, 3, 12, 8, 6, 4, 3, 3, 0, 1,
	0, 3, 3, 3, 3, 4, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 6, 6, 4, 0,
	9, 7, 5, 1, 1, 2, 0, 2, 3, 1,
	1, 3, 6, 3, 4, 1, 1, 1, 2, 5,
	1, 1, 1, 4, 1, 1, 1, 1, 1, 1,
	3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -18, 49, 14,
	32, -9, -5, 33, 6, 8, 51, -12, -13, -6,
	48, -7, -20, 52, 58, -17, -10, -16, 36, 31,
	17, 28, 35, 34, 30, 26, 55, 54, 12, 51,
	15, -3, -4, -4, 4, 12, 40, 4, 51, 12,
	57, 15, 15, 15, 15, 15, 25, -21, -6, -9,
	-4, 59, 15, 53, 9, 32, 9, -4, -4, 4,
	-4, 52, 31, 16, -7, 51, -7, -7, 51, -11,
	19, 41, 20, 21, 22, 23, 42, 43, 44, 24,
	45, 46, 37, 47, 11, 13, 16, 16, -2, -4,
	-2, 13, -4, 13, 29, 12, 29, 29, 16, 15,
	15, 15, 15, 12, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 12, 9, 10, 32, 9, 10, 4,
	51, -7, -4, -7, -20, 51, 16, 16, 16, 16,
	-4, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	-4, -2, -19, 7, -4, -2, -5, 29, 13, 29,
	16, 13, 13, 10, 9, 9, 10, -7, -20, 4,
	-2, -2, 29, 16, -4, 10, 10, -20, 29, -20,
	16,
}
var yyDef = []int{

	3, -2, -2, 2, 4, 0, 6, 7, 0, 0,
	10, 43, 44, 46, 46, 46, 69, 55, 56, 57,
	0, 60, 61, 62, 64, 65, 66, 67, 68, 0,
	0, 0, 0, 0, 0, 0, 47, 46, 0, 0,
	0, 45, 0, 0, 46, 46, 0, 46, 58, 0,
	0, 0, 0, 0, 0, 0, 0, 48, 49, -2,
	0, 0, 0, 0, 3, 46, 3, 51, 0, 46,
	53, 0, 70, 11, 0, 69, 0, 0, 0, 16,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 5, 0, 9, 46, 0,
	46, 63, 54, 0, 0, 46, 0, 0, 15, 0,
	0, 0, 0, 46, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 46, 3, 39, 46, 3, 42, 0,
	59, 0, 0, 0, 0, 69, 21, 22, 23, 24,
	0, 26, 27, 28, 29, 30, 31, 32, 33, 34,
	0, 46, 37, 0, 0, 46, 52, 0, 63, 0,
	14, 25, 35, 8, 3, 3, 41, 0, 0, 46,
	46, 46, 0, 13, 36, 38, 40, 0, 0, 0,
	12,
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
	52, 53, 54, 55, 56, 57, 58, 59,
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
		//line mutan.y:30
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line mutan.y:34
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode) }
	case 3:
		//line mutan.y:35
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line mutan.y:39
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 5:
		//line mutan.y:40
		{ yyVAL.tnode = NewNode(LambdaTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 6:
		//line mutan.y:41
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 7:
		//line mutan.y:42
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 8:
		//line mutan.y:43
		{ yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str }
	case 9:
		//line mutan.y:44
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 10:
		//line mutan.y:45
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 11:
		//line mutan.y:50
		{ yyVAL.tnode = NewNode(StopTy) }
	case 12:
		//line mutan.y:53
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 13:
		//line mutan.y:57
		{
		  	  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
	          }
	case 14:
		//line mutan.y:60
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 15:
		//line mutan.y:61
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 16:
		//line mutan.y:62
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 17:
		//line mutan.y:66
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 18:
		//line mutan.y:67
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 19:
		//line mutan.y:71
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 20:
		//line mutan.y:72
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 21:
		//line mutan.y:76
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 22:
		//line mutan.y:77
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 23:
		//line mutan.y:78
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 24:
		//line mutan.y:79
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 25:
		//line mutan.y:80
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 26:
		//line mutan.y:81
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 27:
		//line mutan.y:82
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 28:
		//line mutan.y:83
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 29:
		//line mutan.y:84
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 30:
		//line mutan.y:85
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 31:
		//line mutan.y:86
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 32:
		//line mutan.y:87
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 33:
		//line mutan.y:88
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 34:
		//line mutan.y:89
		{ yyVAL.tnode = NewNode(GasTy) }
	case 35:
		//line mutan.y:90
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 36:
		//line mutan.y:92
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 37:
		//line mutan.y:100
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 38:
		//line mutan.y:110
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 39:
		//line mutan.y:113
		{ yyVAL.tnode = nil }
	case 40:
		//line mutan.y:118
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 41:
		//line mutan.y:123
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 42:
		//line mutan.y:128
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 43:
		//line mutan.y:134
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 44:
		//line mutan.y:135
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 45:
		//line mutan.y:136
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 46:
		//line mutan.y:137
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 47:
		//line mutan.y:142
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 48:
		//line mutan.y:144
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 49:
		//line mutan.y:148
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 50:
		//line mutan.y:149
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 51:
		//line mutan.y:154
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 52:
		//line mutan.y:160
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 53:
		//line mutan.y:164
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 54:
		//line mutan.y:170
		{
		  	node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
		  	varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		  }
	case 55:
		//line mutan.y:176
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 56:
		//line mutan.y:177
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 57:
		//line mutan.y:178
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 58:
		//line mutan.y:183
		{
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      //$$.VarType = $1
	  }
	case 59:
		//line mutan.y:192
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      //$$.VarType = $1
	      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 60:
		//line mutan.y:202
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 61:
		//line mutan.y:206
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 62:
		//line mutan.y:207
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 63:
		//line mutan.y:208
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 64:
		//line mutan.y:209
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 65:
		//line mutan.y:210
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 66:
		//line mutan.y:211
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 67:
		//line mutan.y:215
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 68:
		//line mutan.y:216
		{ yyVAL.tnode = NewNode(NilTy) }
	case 69:
		//line mutan.y:220
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 70:
		//line mutan.y:224
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

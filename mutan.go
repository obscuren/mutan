
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

//line mutan.y:241



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 51,
	-1, 62,
	55, 46,
	56, 46,
	-2, 55,
}

const yyNprod = 76
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 471

var yyAct = []int{

	23, 4, 12, 2, 22, 64, 53, 38, 37, 66,
	75, 29, 31, 52, 136, 82, 40, 128, 48, 49,
	11, 36, 174, 32, 76, 35, 30, 141, 15, 34,
	33, 29, 188, 182, 167, 43, 165, 14, 111, 110,
	63, 71, 59, 44, 21, 68, 69, 13, 24, 108,
	134, 20, 74, 51, 25, 43, 16, 130, 17, 62,
	78, 80, 81, 44, 9, 170, 42, 31, 38, 37,
	73, 45, 103, 133, 105, 104, 36, 106, 32, 190,
	35, 30, 10, 15, 34, 33, 29, 38, 37, 5,
	61, 45, 14, 72, 169, 183, 38, 37, 168, 21,
	8, 155, 13, 24, 154, 153, 152, 38, 37, 25,
	3, 138, 140, 137, 151, 139, 38, 37, 166, 146,
	16, 126, 17, 150, 186, 46, 47, 149, 9, 156,
	125, 31, 157, 102, 159, 163, 38, 37, 164, 148,
	36, 147, 32, 145, 35, 30, 10, 15, 34, 33,
	29, 144, 143, 5, 142, 112, 14, 101, 100, 77,
	38, 37, 172, 21, 8, 67, 13, 24, 177, 124,
	176, 123, 122, 25, 121, 38, 37, 180, 181, 120,
	184, 119, 16, 187, 17, 118, 185, 116, 115, 189,
	9, 107, 114, 31, 113, 65, 58, 57, 56, 55,
	54, 41, 36, 99, 32, 127, 35, 30, 10, 15,
	34, 33, 29, 117, 109, 5, 39, 173, 14, 16,
	158, 17, 162, 179, 178, 21, 8, 9, 13, 24,
	31, 131, 70, 50, 129, 25, 60, 161, 7, 36,
	26, 32, 28, 35, 30, 10, 15, 34, 33, 29,
	19, 18, 5, 83, 27, 14, 16, 6, 17, 1,
	175, 0, 21, 8, 9, 13, 24, 31, 0, 0,
	0, 0, 25, 0, 0, 0, 36, 0, 32, 0,
	35, 30, 10, 15, 34, 33, 29, 0, 0, 5,
	0, 0, 14, 16, 0, 17, 0, 171, 0, 21,
	8, 9, 13, 24, 31, 0, 0, 0, 0, 25,
	0, 0, 0, 36, 0, 32, 0, 35, 30, 10,
	15, 34, 33, 29, 0, 0, 5, 0, 0, 14,
	16, 0, 17, 0, 135, 0, 21, 8, 9, 13,
	24, 31, 0, 0, 0, 0, 25, 0, 0, 0,
	36, 0, 32, 0, 35, 30, 10, 15, 34, 33,
	29, 0, 0, 5, 0, 0, 14, 16, 0, 17,
	0, 132, 0, 21, 8, 9, 13, 24, 31, 0,
	0, 0, 0, 25, 0, 0, 0, 36, 0, 32,
	0, 35, 30, 10, 15, 34, 33, 29, 0, 0,
	5, 0, 0, 14, 0, 0, 0, 0, 31, 0,
	21, 8, 0, 13, 24, 0, 0, 36, 0, 32,
	25, 35, 30, 98, 0, 34, 33, 29, 31, 0,
	0, 84, 86, 87, 88, 89, 93, 36, 0, 32,
	21, 35, 30, 160, 24, 34, 33, 29, 0, 96,
	25, 0, 0, 85, 0, 90, 91, 92, 94, 95,
	97, 0, 0, 79, 24, 0, 0, 0, 0, 0,
	25,
}
var yyPact = []int{

	-1000, -1000, 50, -1000, -48, 204, -1000, -1000, -36, 186,
	-1000, -1000, -1000, 51, 50, 50, -5, -5, 229, -1000,
	-1000, 1, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-52, 185, 184, 183, 182, 181, 17, -1000, -5, -55,
	180, -45, 149, -5, -5, 228, -1000, -1000, 32, 61,
	-5, -1000, -43, -7, 143, 411, 411, 411, -37, 412,
	-1000, -1000, -1000, -48, 190, 142, 141, -1000, -48, 120,
	-5, -1000, -5, -1000, -48, 178, -1000, -1000, 20, 202,
	10, 9, 139, -1000, 179, 177, 173, 172, 201, 170,
	166, 164, 159, 157, 156, 154, 115, 106, 193, -1000,
	8, -1000, 227, -48, 361, 41, 324, -38, 411, -5,
	411, -25, -1000, 138, 136, 135, 127, -5, 125, 123,
	111, 107, 98, 90, 89, 88, 85, -5, -1000, 211,
	-1000, 391, 215, -5, -1000, -1000, -1000, 7, 105, 5,
	82, -1000, -1000, -1000, -1000, -1000, 81, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 52, 287, -1000, -1000,
	31, -1000, 208, 13, 250, 411, -1000, -25, -1000, -1000,
	220, -1000, 213, -1000, -1000, -1000, 4, 79, -5, -1000,
	176, 114, -25, -1000, -48, -1000, -1000, 3, -25, 63,
	-1000,
}
var yyPgo = []int{

	0, 259, 3, 110, 1, 2, 51, 4, 257, 20,
	254, 253, 251, 250, 250, 250, 242, 240, 238, 237,
	0, 236, 234,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 22, 22, 10, 10, 10, 10, 10, 10,
	14, 14, 15, 15, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	8, 19, 19, 18, 18, 18, 4, 4, 4, 4,
	4, 4, 9, 9, 21, 21, 5, 5, 5, 5,
	5, 5, 5, 12, 13, 6, 7, 7, 7, 7,
	7, 7, 20, 20, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 7, 8,
	4, 1, 1, 0, 3, 12, 8, 6, 4, 3,
	3, 0, 1, 0, 3, 3, 3, 3, 4, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 6,
	6, 4, 0, 9, 7, 5, 1, 1, 3, 2,
	2, 0, 2, 3, 1, 1, 3, 6, 3, 4,
	1, 1, 1, 2, 5, 1, 1, 1, 4, 1,
	1, 1, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -18, 50, 14,
	32, -9, -5, 52, 42, 33, 6, 8, -12, -13,
	-6, 49, -7, -20, 53, 59, -17, -10, -16, 36,
	31, 17, 28, 35, 34, 30, 26, 56, 55, 12,
	52, 15, 15, 4, 12, 40, -3, -3, -4, -4,
	4, 52, 12, 58, 15, 15, 15, 15, 15, 25,
	-21, -6, -9, -4, 60, 15, 54, 16, -4, -4,
	4, 9, 32, 9, -4, 53, 31, 16, -7, 52,
	-7, -7, 52, -11, 19, 41, 20, 21, 22, 23,
	43, 44, 45, 24, 46, 47, 37, 48, 11, 13,
	16, 16, 13, -4, -2, -4, -2, 13, 29, 12,
	29, 29, 16, 15, 15, 15, 15, 12, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 12, 9, -22,
	49, 4, 10, 32, 9, 10, 52, -7, -4, -7,
	-20, 52, 16, 16, 16, 16, -4, 16, 16, 16,
	16, 16, 16, 16, 16, 16, -4, -2, 9, -5,
	52, -19, 7, -4, -2, 29, 13, 29, 16, 13,
	13, 10, -2, 9, 9, 10, -7, -20, 4, 10,
	-2, -2, 29, 16, -4, 10, 10, -20, 29, -20,
	16,
}
var yyDef = []int{

	3, -2, -2, 2, 4, 0, 6, 7, 0, 0,
	11, 46, 47, 74, 51, 51, 51, 51, 60, 61,
	62, 0, 65, 66, 67, 69, 70, 71, 72, 73,
	0, 0, 0, 0, 0, 0, 0, 52, 51, 0,
	0, 0, 0, 51, 51, 0, 49, 50, 0, 0,
	51, 63, 0, 0, 0, 0, 0, 0, 0, 0,
	53, 54, -2, 0, 0, 0, 0, 48, 56, 0,
	51, 3, 51, 3, 58, 0, 75, 14, 0, 74,
	0, 0, 0, 19, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 5,
	0, 10, 68, 59, 51, 0, 51, 0, 0, 51,
	0, 0, 18, 0, 0, 0, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 51, 3, 0,
	12, 0, 42, 51, 3, 45, 64, 0, 0, 0,
	0, 74, 24, 25, 26, 27, 0, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 0, 51, 3, 57,
	74, 40, 0, 0, 51, 0, 68, 0, 17, 28,
	38, 8, 51, 3, 3, 44, 0, 0, 51, 9,
	51, 51, 0, 16, 39, 41, 43, 0, 0, 0,
	15,
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
		//line mutan.y:32
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line mutan.y:36
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode) }
	case 3:
		//line mutan.y:37
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line mutan.y:41
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 5:
		//line mutan.y:42
		{ yyVAL.tnode = NewNode(LambdaTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 6:
		//line mutan.y:43
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 7:
		//line mutan.y:44
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 8:
		//line mutan.y:45
		{ yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str }
	case 9:
		//line mutan.y:47
		{
				yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode);
				yyVAL.tnode.Constant = yyS[yypt-6].str
				yyVAL.tnode.HasRet = yyS[yypt-3].check
			}
	case 10:
		//line mutan.y:52
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 11:
		//line mutan.y:53
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 12:
		//line mutan.y:57
		{ yyVAL.check = true }
	case 13:
		//line mutan.y:58
		{ yyVAL.check = false }
	case 14:
		//line mutan.y:62
		{ yyVAL.tnode = NewNode(StopTy) }
	case 15:
		//line mutan.y:65
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 16:
		//line mutan.y:69
		{
			  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 17:
		//line mutan.y:72
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 18:
		//line mutan.y:73
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 19:
		//line mutan.y:74
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 20:
		//line mutan.y:78
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 21:
		//line mutan.y:79
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 22:
		//line mutan.y:83
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 23:
		//line mutan.y:84
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 24:
		//line mutan.y:88
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 25:
		//line mutan.y:89
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 26:
		//line mutan.y:90
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 27:
		//line mutan.y:91
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 28:
		//line mutan.y:92
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 29:
		//line mutan.y:93
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 30:
		//line mutan.y:94
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 31:
		//line mutan.y:95
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 32:
		//line mutan.y:96
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 33:
		//line mutan.y:97
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 34:
		//line mutan.y:98
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 35:
		//line mutan.y:99
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 36:
		//line mutan.y:100
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 37:
		//line mutan.y:101
		{ yyVAL.tnode = NewNode(GasTy) }
	case 38:
		//line mutan.y:102
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 39:
		//line mutan.y:104
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 40:
		//line mutan.y:112
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 41:
		//line mutan.y:122
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 42:
		//line mutan.y:125
		{ yyVAL.tnode = nil }
	case 43:
		//line mutan.y:130
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 44:
		//line mutan.y:135
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 45:
		//line mutan.y:140
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 46:
		//line mutan.y:146
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 47:
		//line mutan.y:147
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 48:
		//line mutan.y:148
		{ yyVAL.tnode = NewNode(FuncCallTy); yyVAL.tnode.Constant = yyS[yypt-2].str }
	case 49:
		//line mutan.y:149
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 50:
		//line mutan.y:150
		{ yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode) }
	case 51:
		//line mutan.y:151
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 52:
		//line mutan.y:156
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 53:
		//line mutan.y:158
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 54:
		//line mutan.y:162
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 55:
		//line mutan.y:163
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 56:
		//line mutan.y:168
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 57:
		//line mutan.y:174
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 58:
		//line mutan.y:178
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 59:
		//line mutan.y:184
		{
			node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
			varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		  }
	case 60:
		//line mutan.y:190
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 61:
		//line mutan.y:191
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 62:
		//line mutan.y:192
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 63:
		//line mutan.y:197
		{
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      //$$.VarType = $1
	  }
	case 64:
		//line mutan.y:206
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      //$$.VarType = $1
	      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
	
	}
	case 65:
		//line mutan.y:216
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 66:
		//line mutan.y:220
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 67:
		//line mutan.y:221
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 68:
		//line mutan.y:222
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 69:
		//line mutan.y:223
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 70:
		//line mutan.y:224
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 71:
		//line mutan.y:225
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 72:
		//line mutan.y:229
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 73:
		//line mutan.y:230
		{ yyVAL.tnode = NewNode(NilTy) }
	case 74:
		//line mutan.y:234
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 75:
		//line mutan.y:238
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

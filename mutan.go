
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

//line mutan.y:239



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 50,
	-1, 61,
	54, 47,
	55, 47,
	-2, 54,
}

const yyNprod = 75
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 478

var yyAct = []int{

	22, 4, 13, 2, 21, 63, 52, 168, 65, 133,
	74, 72, 30, 37, 36, 127, 135, 46, 48, 172,
	28, 35, 81, 31, 51, 34, 29, 167, 14, 33,
	32, 28, 132, 39, 71, 140, 164, 12, 101, 62,
	75, 58, 70, 20, 67, 68, 47, 23, 37, 36,
	186, 73, 188, 24, 37, 36, 37, 36, 42, 77,
	79, 80, 129, 50, 37, 36, 43, 180, 37, 36,
	165, 102, 30, 104, 103, 61, 105, 37, 36, 37,
	36, 35, 163, 31, 42, 34, 29, 37, 36, 33,
	32, 28, 43, 110, 44, 41, 109, 15, 19, 16,
	107, 184, 181, 20, 166, 9, 47, 23, 30, 154,
	137, 139, 136, 24, 138, 153, 152, 35, 145, 31,
	44, 34, 29, 11, 14, 33, 32, 28, 155, 151,
	5, 156, 150, 158, 161, 149, 60, 162, 148, 20,
	8, 147, 10, 23, 146, 144, 143, 142, 141, 24,
	111, 100, 99, 76, 66, 125, 124, 123, 122, 121,
	120, 170, 119, 118, 117, 115, 175, 114, 174, 15,
	113, 16, 112, 183, 64, 178, 179, 9, 182, 57,
	30, 185, 56, 55, 54, 53, 40, 187, 3, 35,
	106, 31, 98, 34, 29, 11, 14, 33, 32, 28,
	126, 116, 5, 45, 108, 15, 38, 16, 176, 177,
	171, 20, 8, 9, 10, 23, 30, 157, 160, 130,
	69, 24, 49, 128, 59, 35, 159, 31, 7, 34,
	29, 11, 14, 33, 32, 28, 25, 27, 5, 18,
	17, 15, 82, 16, 26, 173, 6, 20, 8, 9,
	10, 23, 30, 1, 0, 0, 0, 24, 0, 0,
	0, 35, 0, 31, 0, 34, 29, 11, 14, 33,
	32, 28, 0, 0, 5, 0, 0, 15, 0, 16,
	0, 169, 0, 20, 8, 9, 10, 23, 30, 0,
	0, 0, 0, 24, 0, 0, 0, 35, 0, 31,
	0, 34, 29, 11, 14, 33, 32, 28, 0, 0,
	5, 0, 0, 15, 0, 16, 0, 134, 0, 20,
	8, 9, 10, 23, 30, 0, 0, 0, 0, 24,
	0, 0, 0, 35, 0, 31, 0, 34, 29, 11,
	14, 33, 32, 28, 0, 0, 5, 0, 0, 15,
	0, 16, 0, 131, 0, 20, 8, 9, 10, 23,
	30, 0, 0, 0, 0, 24, 0, 0, 0, 35,
	0, 31, 0, 34, 29, 11, 14, 33, 32, 28,
	0, 0, 5, 0, 0, 15, 0, 16, 0, 0,
	0, 20, 8, 9, 10, 23, 30, 0, 0, 0,
	0, 24, 0, 0, 0, 35, 0, 31, 0, 34,
	29, 11, 14, 33, 32, 28, 0, 0, 5, 0,
	0, 0, 0, 0, 0, 30, 0, 20, 8, 0,
	10, 23, 0, 0, 35, 0, 31, 24, 34, 29,
	0, 97, 33, 32, 28, 0, 0, 0, 0, 83,
	85, 86, 87, 88, 92, 0, 0, 0, 0, 78,
	23, 0, 0, 0, 0, 0, 24, 95, 0, 0,
	0, 84, 89, 90, 91, 93, 94, 96,
}
var yyPact = []int{

	-1000, -1000, 379, -1000, -41, 194, -1000, -1000, -18, 171,
	80, -1000, -1000, -1000, 379, -5, -5, 218, -1000, -1000,
	12, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -51,
	170, 169, 168, 167, 164, 16, -1000, -5, -54, 159,
	-45, 138, -5, -5, 216, -1000, 33, 54, 2, -5,
	-1000, -42, 9, 137, 408, 408, 408, -29, 430, -1000,
	-1000, -1000, -41, 179, 136, 135, -1000, -41, 25, -5,
	-1000, -5, -1000, -41, 177, -1000, -1000, 71, 192, 67,
	64, 134, -1000, 157, 155, 152, 150, 189, 149, 148,
	147, 145, 144, 143, 142, 141, 140, 188, -1000, 6,
	-1000, 215, -41, 343, 0, 307, -35, 408, -5, 408,
	-16, -1000, 132, 131, 130, 129, -5, 128, 125, 122,
	119, 116, 113, 100, 99, 93, -5, -1000, 208, -1000,
	55, 211, -5, -1000, -1000, -1000, 53, 23, 41, 88,
	-1000, -1000, -1000, -1000, -1000, 14, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -6, 271, -1000, -1000, -1000,
	201, 10, 235, 408, -1000, -16, -1000, -1000, 204, -1000,
	199, -1000, -1000, -1000, 38, 86, -5, -1000, 163, 91,
	-16, -1000, -41, -1000, -1000, 21, -16, 36, -1000,
}
var yyPgo = []int{

	0, 253, 3, 188, 1, 2, 98, 4, 246, 37,
	244, 242, 240, 239, 239, 239, 237, 236, 228, 226,
	0, 224, 223,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 22, 22, 10, 10, 10, 10, 10,
	10, 14, 14, 15, 15, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 8, 19, 19, 18, 18, 18, 4, 4, 4,
	4, 9, 9, 21, 21, 5, 5, 5, 5, 5,
	5, 5, 12, 13, 6, 7, 7, 7, 7, 7,
	7, 20, 20, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 7, 8,
	4, 3, 1, 1, 0, 3, 12, 8, 6, 4,
	3, 3, 0, 1, 0, 3, 3, 3, 3, 4,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	6, 6, 4, 0, 9, 7, 5, 1, 1, 2,
	0, 2, 3, 1, 1, 3, 6, 3, 4, 1,
	1, 1, 2, 5, 1, 1, 1, 4, 1, 1,
	1, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -18, 49, 14,
	51, 32, -9, -5, 33, 6, 8, -12, -13, -6,
	48, -7, -20, 52, 58, -17, -10, -16, 36, 31,
	17, 28, 35, 34, 30, 26, 55, 54, 12, 51,
	15, 15, 4, 12, 40, -3, -4, 51, -4, 4,
	51, 12, 57, 15, 15, 15, 15, 15, 25, -21,
	-6, -9, -4, 59, 15, 53, 16, -4, -4, 4,
	9, 32, 9, -4, 52, 31, 16, -7, 51, -7,
	-7, 51, -11, 19, 41, 20, 21, 22, 23, 42,
	43, 44, 24, 45, 46, 37, 47, 11, 13, 16,
	16, 13, -4, -2, -4, -2, 13, 29, 12, 29,
	29, 16, 15, 15, 15, 15, 12, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 12, 9, -22, 56,
	4, 10, 32, 9, 10, 51, -7, -4, -7, -20,
	51, 16, 16, 16, 16, -4, 16, 16, 16, 16,
	16, 16, 16, 16, 16, -4, -2, 9, -5, -19,
	7, -4, -2, 29, 13, 29, 16, 13, 13, 10,
	-2, 9, 9, 10, -7, -20, 4, 10, -2, -2,
	29, 16, -4, 10, 10, -20, 29, -20, 16,
}
var yyDef = []int{

	3, -2, -2, 2, 4, 0, 6, 7, 0, 0,
	73, 12, 47, 48, 50, 50, 50, 59, 60, 61,
	0, 64, 65, 66, 68, 69, 70, 71, 72, 0,
	0, 0, 0, 0, 0, 0, 51, 50, 0, 0,
	0, 0, 50, 50, 0, 49, 0, 73, 0, 50,
	62, 0, 0, 0, 0, 0, 0, 0, 0, 52,
	53, -2, 0, 0, 0, 0, 11, 55, 0, 50,
	3, 50, 3, 57, 0, 74, 15, 0, 73, 0,
	0, 0, 20, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 5, 0,
	10, 67, 58, 50, 0, 50, 0, 0, 50, 0,
	0, 19, 0, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 50, 3, 0, 13,
	0, 43, 50, 3, 46, 63, 0, 0, 0, 0,
	73, 25, 26, 27, 28, 0, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 0, 50, 3, 56, 41,
	0, 0, 50, 0, 67, 0, 18, 29, 39, 8,
	50, 3, 3, 45, 0, 0, 50, 9, 50, 50,
	0, 17, 40, 42, 44, 0, 0, 0, 16,
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
		//line mutan.y:46
		{
	        yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode);
	        yyVAL.tnode.Constant = yyS[yypt-6].str
	        yyVAL.tnode.HasRet = yyS[yypt-3].check
	      }
	case 10:
		//line mutan.y:51
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 11:
		//line mutan.y:52
		{ yyVAL.tnode = NewNode(FuncCallTy); yyVAL.tnode.Constant = yyS[yypt-2].str }
	case 12:
		//line mutan.y:53
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 13:
		//line mutan.y:57
		{ yyVAL.check = true }
	case 14:
		//line mutan.y:58
		{ yyVAL.check = false }
	case 15:
		//line mutan.y:62
		{ yyVAL.tnode = NewNode(StopTy) }
	case 16:
		//line mutan.y:65
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 17:
		//line mutan.y:69
		{
		  	  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
	          }
	case 18:
		//line mutan.y:72
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 19:
		//line mutan.y:73
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 20:
		//line mutan.y:74
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 21:
		//line mutan.y:78
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 22:
		//line mutan.y:79
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 23:
		//line mutan.y:83
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 24:
		//line mutan.y:84
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 25:
		//line mutan.y:88
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 26:
		//line mutan.y:89
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 27:
		//line mutan.y:90
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 28:
		//line mutan.y:91
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 29:
		//line mutan.y:92
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 30:
		//line mutan.y:93
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 31:
		//line mutan.y:94
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 32:
		//line mutan.y:95
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 33:
		//line mutan.y:96
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 34:
		//line mutan.y:97
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 35:
		//line mutan.y:98
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 36:
		//line mutan.y:99
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 37:
		//line mutan.y:100
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 38:
		//line mutan.y:101
		{ yyVAL.tnode = NewNode(GasTy) }
	case 39:
		//line mutan.y:102
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 40:
		//line mutan.y:104
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 41:
		//line mutan.y:112
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 42:
		//line mutan.y:122
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 43:
		//line mutan.y:125
		{ yyVAL.tnode = nil }
	case 44:
		//line mutan.y:130
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 45:
		//line mutan.y:135
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 46:
		//line mutan.y:140
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 47:
		//line mutan.y:146
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 48:
		//line mutan.y:147
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 49:
		//line mutan.y:148
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 50:
		//line mutan.y:149
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 51:
		//line mutan.y:154
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 52:
		//line mutan.y:156
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 53:
		//line mutan.y:160
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 54:
		//line mutan.y:161
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 55:
		//line mutan.y:166
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 56:
		//line mutan.y:172
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 57:
		//line mutan.y:176
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 58:
		//line mutan.y:182
		{
		  	node := NewNode(SetLocalTy)
			node.Constant = yyS[yypt-3].str
		  	varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
			yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
		  }
	case 59:
		//line mutan.y:188
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 60:
		//line mutan.y:189
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 61:
		//line mutan.y:190
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 62:
		//line mutan.y:195
		{
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      //$$.VarType = $1
	  }
	case 63:
		//line mutan.y:204
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      //$$.VarType = $1
	      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 64:
		//line mutan.y:214
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 65:
		//line mutan.y:218
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 66:
		//line mutan.y:219
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 67:
		//line mutan.y:220
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 68:
		//line mutan.y:221
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 69:
		//line mutan.y:222
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 70:
		//line mutan.y:223
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 71:
		//line mutan.y:227
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 72:
		//line mutan.y:228
		{ yyVAL.tnode = NewNode(NilTy) }
	case 73:
		//line mutan.y:232
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 74:
		//line mutan.y:236
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

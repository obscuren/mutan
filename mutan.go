
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
const DIFFICULTY = 57379
const PREVHASH = 57380
const TIMESTAMP = 57381
const BLOCKNUM = 57382
const COINBASE = 57383
const GAS = 57384
const ID = 57385
const NUMBER = 57386
const INLINE_ASM = 57387
const OP = 57388
const DOP = 57389
const TYPE = 57390
const STR = 57391

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
	"DIFFICULTY",
	"PREVHASH",
	"TIMESTAMP",
	"BLOCKNUM",
	"COINBASE",
	"GAS",
	"ID",
	"NUMBER",
	"INLINE_ASM",
	"OP",
	"DOP",
	"TYPE",
	"STR",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line mutan.y:208



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 41,
}

const yyNprod = 62
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 315

var yyAct = []int{

	5, 20, 2, 10, 19, 134, 27, 43, 33, 32,
	50, 57, 35, 36, 37, 31, 25, 28, 77, 30,
	26, 42, 11, 138, 29, 25, 65, 66, 67, 68,
	69, 73, 14, 21, 49, 106, 103, 18, 53, 54,
	55, 56, 33, 32, 70, 71, 72, 74, 75, 76,
	60, 62, 41, 80, 79, 63, 81, 3, 12, 102,
	13, 52, 146, 58, 132, 148, 8, 142, 130, 27,
	128, 131, 86, 33, 32, 33, 32, 84, 31, 48,
	28, 129, 30, 26, 51, 11, 108, 29, 25, 107,
	82, 109, 113, 83, 150, 14, 21, 33, 32, 27,
	18, 122, 143, 125, 33, 32, 126, 121, 31, 127,
	28, 120, 30, 26, 33, 32, 119, 29, 25, 118,
	117, 33, 32, 33, 32, 14, 21, 12, 116, 13,
	18, 145, 137, 136, 115, 8, 140, 141, 27, 114,
	144, 112, 111, 110, 147, 87, 78, 31, 59, 28,
	149, 30, 26, 99, 11, 98, 29, 25, 97, 12,
	96, 13, 95, 135, 14, 21, 94, 8, 139, 18,
	27, 93, 92, 90, 89, 88, 47, 46, 45, 31,
	44, 28, 34, 30, 26, 100, 11, 91, 29, 25,
	38, 12, 85, 13, 124, 104, 14, 21, 39, 8,
	133, 18, 27, 105, 40, 123, 7, 22, 24, 16,
	15, 31, 64, 28, 23, 30, 26, 9, 11, 6,
	29, 25, 17, 12, 4, 13, 1, 101, 14, 21,
	0, 8, 0, 18, 27, 0, 0, 0, 0, 0,
	0, 0, 0, 31, 0, 28, 0, 30, 26, 0,
	11, 0, 29, 25, 0, 12, 0, 13, 0, 0,
	14, 21, 0, 8, 0, 18, 27, 0, 0, 0,
	0, 0, 0, 0, 0, 31, 0, 28, 0, 30,
	26, 0, 11, 0, 29, 25, 0, 27, 0, 0,
	0, 0, 14, 21, 0, 0, 31, 18, 28, 0,
	30, 26, 0, 0, 0, 29, 25, 0, 0, 0,
	0, 0, 0, 61, 21,
}
var yyPact = []int{

	25, -1000, 249, -1000, -1000, -38, -1000, -1000, 167, -1000,
	-1000, -11, -11, -11, 186, 200, -1000, -1000, 9, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -42, 165, 163, 162,
	161, 54, -1000, -11, -35, -38, 75, 29, -11, -11,
	-11, -1000, -33, 32, 132, 270, 270, 12, 7, -38,
	130, 25, -11, 25, -38, 77, -38, 80, -1000, -1000,
	48, 180, 43, 129, -1000, 160, 159, 158, 175, 157,
	156, 151, 147, 145, 143, 140, 138, 173, -1000, 217,
	27, 185, 199, -8, 270, -11, 270, -1000, 127, 126,
	125, -11, 123, 118, 112, 104, 103, 100, 95, 91,
	-11, 187, -11, 25, -1000, 82, -1000, 41, 68, 39,
	-1000, -1000, -1000, 58, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 51, -1000, 191, -4, 153, -1000, 270, -1000,
	-20, -1000, 164, 25, 25, -1000, 38, 86, -1000, -11,
	121, 52, -20, -1000, -38, -1000, -1000, 36, -20, 78,
	-1000,
}
var yyPgo = []int{

	0, 226, 2, 224, 0, 3, 222, 4, 219, 217,
	214, 212, 210, 209, 209, 209, 208, 207, 206, 205,
	1,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 10,
	10, 10, 10, 10, 14, 14, 15, 15, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 8, 19, 19, 18, 18, 18, 4, 4,
	4, 4, 9, 9, 5, 5, 5, 5, 5, 5,
	12, 13, 6, 7, 7, 7, 7, 7, 20, 20,
	16, 17,
}
var yyR2 = []int{

	0, 1, 2, 1, 0, 1, 1, 1, 4, 3,
	12, 8, 4, 3, 3, 0, 1, 0, 3, 3,
	3, 4, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 6, 6, 4, 0, 9, 7, 5, 1, 1,
	2, 0, 2, 3, 3, 6, 3, 1, 1, 1,
	2, 5, 1, 1, 1, 4, 1, 1, 1, 1,
	1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 32, -3, -4, -8, -18, 14, -9,
	-5, 33, 6, 8, 43, -12, -13, -6, 48, -7,
	-20, 44, -17, -10, -16, 36, 31, 17, 28, 35,
	30, 26, 47, 46, 15, -4, -4, -4, 4, 12,
	4, 43, 12, 49, 15, 15, 15, 15, 25, -4,
	45, 9, 32, 9, -4, -4, -4, 44, 31, 16,
	-7, 43, -7, 43, -11, 19, 20, 21, 22, 23,
	37, 38, 39, 24, 40, 41, 42, 11, 16, -2,
	-4, -2, 13, 13, 29, 12, 29, 16, 15, 15,
	15, 12, 15, 15, 15, 15, 15, 15, 15, 15,
	12, 10, 32, 9, 10, 4, 43, -7, -4, -7,
	16, 16, 16, -4, 16, 16, 16, 16, 16, 16,
	16, 16, -4, -19, 7, -4, -2, -5, 29, 13,
	29, 13, 13, 9, 9, 10, -7, -20, 43, 4,
	-2, -2, 29, 16, -4, 10, 10, -20, 29, -20,
	16,
}
var yyDef = []int{

	4, -2, -2, 3, 2, 5, 6, 7, 0, 38,
	39, 41, 41, 41, 60, 47, 48, 49, 0, 52,
	53, 54, 56, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 42, 41, 0, 40, 0, 0, 41, 41,
	41, 50, 0, 0, 0, 0, 0, 0, 0, 43,
	0, 4, 41, 4, 44, 0, 46, 0, 61, 9,
	0, 60, 0, 0, 13, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 8, 41,
	0, 41, 55, 0, 0, 41, 0, 12, 0, 0,
	0, 41, 0, 0, 0, 0, 0, 0, 0, 0,
	41, 34, 41, 4, 37, 0, 51, 0, 0, 0,
	18, 19, 20, 0, 22, 23, 24, 25, 26, 27,
	28, 29, 0, 32, 0, 0, 41, 45, 0, 55,
	0, 21, 30, 4, 4, 36, 0, 0, 60, 41,
	41, 41, 0, 11, 31, 33, 35, 0, 0, 0,
	10,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49,
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
		//line mutan.y:36
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 5:
		//line mutan.y:40
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 6:
		//line mutan.y:41
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 7:
		//line mutan.y:42
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 8:
		//line mutan.y:43
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 9:
		//line mutan.y:48
		{ yyVAL.tnode = NewNode(StopTy) }
	case 10:
		//line mutan.y:51
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 11:
		//line mutan.y:55
		{
		  	  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
	          }
	case 12:
		//line mutan.y:58
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 13:
		//line mutan.y:59
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 14:
		//line mutan.y:63
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 15:
		//line mutan.y:64
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 16:
		//line mutan.y:68
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 17:
		//line mutan.y:69
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 18:
		//line mutan.y:73
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 19:
		//line mutan.y:74
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 20:
		//line mutan.y:75
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 21:
		//line mutan.y:76
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 22:
		//line mutan.y:77
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 23:
		//line mutan.y:78
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 24:
		//line mutan.y:79
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 25:
		//line mutan.y:80
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 26:
		//line mutan.y:81
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 27:
		//line mutan.y:82
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 28:
		//line mutan.y:83
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 29:
		//line mutan.y:84
		{ yyVAL.tnode = NewNode(GasTy) }
	case 30:
		//line mutan.y:85
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 31:
		//line mutan.y:87
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 32:
		//line mutan.y:95
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 33:
		//line mutan.y:105
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 34:
		//line mutan.y:108
		{ yyVAL.tnode = nil }
	case 35:
		//line mutan.y:113
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 36:
		//line mutan.y:118
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 37:
		//line mutan.y:123
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 38:
		//line mutan.y:129
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 39:
		//line mutan.y:130
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 40:
		//line mutan.y:131
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 41:
		//line mutan.y:132
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 42:
		//line mutan.y:136
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 43:
		//line mutan.y:137
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 44:
		//line mutan.y:143
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 45:
		//line mutan.y:149
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 46:
		//line mutan.y:153
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 47:
		//line mutan.y:158
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 48:
		//line mutan.y:159
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 49:
		//line mutan.y:160
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 50:
		//line mutan.y:165
		{
		
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      yyVAL.tnode.VarType = yyS[yypt-1].str
		  }
	case 51:
		//line mutan.y:174
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      yyVAL.tnode.VarType = yyS[yypt-4].str
		      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 52:
		//line mutan.y:184
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 53:
		//line mutan.y:188
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 54:
		//line mutan.y:189
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 55:
		//line mutan.y:190
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 56:
		//line mutan.y:191
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 57:
		//line mutan.y:192
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 58:
		//line mutan.y:196
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 59:
		//line mutan.y:197
		{ yyVAL.tnode = NewNode(NilTy) }
	case 60:
		//line mutan.y:201
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 61:
		//line mutan.y:205
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

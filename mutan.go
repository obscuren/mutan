
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
const DIFFICULTY = 57376
const PREVHASH = 57377
const TIMESTAMP = 57378
const BLOCKNUM = 57379
const COINBASE = 57380
const GAS = 57381
const ID = 57382
const NUMBER = 57383
const INLINE_ASM = 57384
const OP = 57385
const DOP = 57386
const TYPE = 57387
const STR = 57388

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

//line mutan.y:199



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 40,
}

const yyNprod = 59
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 287

var yyAct = []int{

	5, 20, 2, 19, 10, 40, 25, 30, 29, 46,
	53, 39, 32, 33, 34, 28, 137, 26, 100, 27,
	24, 58, 11, 3, 54, 138, 132, 121, 140, 14,
	21, 45, 79, 44, 18, 114, 50, 51, 52, 38,
	12, 113, 13, 112, 135, 25, 56, 111, 8, 75,
	74, 25, 76, 110, 28, 109, 26, 97, 27, 24,
	28, 108, 26, 107, 27, 24, 105, 11, 14, 21,
	104, 124, 103, 18, 14, 21, 81, 73, 55, 18,
	96, 102, 126, 101, 123, 12, 106, 13, 94, 134,
	93, 30, 29, 8, 78, 115, 25, 118, 92, 91,
	119, 30, 29, 49, 120, 28, 90, 26, 89, 27,
	24, 47, 11, 122, 30, 29, 30, 29, 77, 14,
	21, 88, 87, 86, 18, 128, 48, 84, 130, 131,
	133, 12, 83, 13, 136, 127, 82, 30, 29, 8,
	139, 43, 25, 30, 29, 30, 29, 42, 30, 29,
	41, 28, 31, 26, 85, 27, 24, 80, 11, 35,
	12, 125, 13, 117, 98, 14, 21, 36, 8, 129,
	18, 25, 99, 37, 116, 7, 22, 16, 15, 59,
	28, 23, 26, 9, 27, 24, 6, 11, 17, 12,
	4, 13, 1, 95, 14, 21, 0, 8, 0, 18,
	25, 0, 0, 0, 0, 0, 0, 0, 0, 28,
	12, 26, 13, 27, 24, 0, 11, 0, 8, 0,
	0, 25, 0, 14, 21, 0, 0, 0, 18, 0,
	28, 0, 26, 0, 27, 24, 0, 11, 0, 0,
	0, 0, 72, 0, 14, 21, 0, 0, 0, 18,
	60, 61, 62, 63, 64, 68, 0, 0, 0, 0,
	0, 0, 25, 0, 0, 65, 66, 67, 69, 70,
	71, 28, 0, 26, 0, 27, 24, 0, 0, 0,
	0, 0, 0, 0, 0, 57, 21,
}
var yyPact = []int{

	-9, -1000, 204, -1000, -1000, -36, -1000, -1000, 137, -1000,
	-1000, -11, -11, -11, 155, 169, -1000, -1000, -1, -1000,
	-1000, -1000, -1000, -1000, -41, 135, 132, 126, 8, -1000,
	-11, -33, -36, 102, 94, -11, -11, -11, -1000, -31,
	-7, 62, 245, -19, 231, -36, 61, -9, -11, -9,
	-36, 105, -36, 81, -1000, -1000, 3, 145, 60, -1000,
	121, 117, 112, 142, 108, 107, 106, 93, 91, 84,
	83, 75, 76, -1000, 183, 48, 154, 168, -22, 245,
	-11, -1000, 56, 54, 50, -11, 47, 45, 39, 37,
	31, 27, 25, 19, -11, 156, -11, -9, -1000, 28,
	-1000, -2, 100, -1000, -1000, -1000, 71, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 58, -1000, 152, 73, 125,
	-1000, 245, -1000, -1000, 165, -9, -9, -1000, -3, -11,
	79, 34, -24, -36, -1000, -1000, -4, -1000, -24, 12,
	-1000,
}
var yyPgo = []int{

	0, 192, 2, 190, 0, 4, 188, 3, 186, 183,
	181, 179, 178, 177, 177, 177, 1, 176, 175, 174,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 10,
	10, 10, 10, 14, 14, 15, 15, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 8, 19, 19, 18, 18, 18, 4, 4, 4,
	4, 9, 9, 5, 5, 5, 5, 5, 5, 12,
	13, 6, 7, 7, 7, 7, 7, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 1, 0, 1, 1, 1, 4, 3,
	12, 4, 3, 3, 0, 1, 0, 3, 3, 3,
	4, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	6, 6, 4, 0, 9, 7, 5, 1, 1, 2,
	0, 2, 3, 3, 6, 3, 1, 1, 1, 2,
	5, 1, 1, 1, 4, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 32, -3, -4, -8, -18, 14, -9,
	-5, 33, 6, 8, 40, -12, -13, -6, 45, -7,
	-16, 41, -17, -10, 31, 17, 28, 30, 26, 44,
	43, 15, -4, -4, -4, 4, 12, 4, 40, 12,
	46, 15, 15, 15, 25, -4, 42, 9, 32, 9,
	-4, -4, -4, 41, 31, 16, -7, 40, 40, -11,
	19, 20, 21, 22, 23, 34, 35, 36, 24, 37,
	38, 39, 11, 16, -2, -4, -2, 13, 13, 29,
	12, 16, 15, 15, 15, 12, 15, 15, 15, 15,
	15, 15, 15, 15, 12, 10, 32, 9, 10, 4,
	40, -7, -4, 16, 16, 16, -4, 16, 16, 16,
	16, 16, 16, 16, 16, -4, -19, 7, -4, -2,
	-5, 29, 13, 13, 13, 9, 9, 10, -7, 4,
	-2, -2, 29, -4, 10, 10, -16, 40, 29, -16,
	16,
}
var yyDef = []int{

	4, -2, -2, 3, 2, 5, 6, 7, 0, 37,
	38, 40, 40, 40, 57, 46, 47, 48, 0, 51,
	52, 53, 55, 56, 0, 0, 0, 0, 0, 41,
	40, 0, 39, 0, 0, 40, 40, 40, 49, 0,
	0, 0, 0, 0, 0, 42, 0, 4, 40, 4,
	43, 0, 45, 0, 58, 9, 0, 57, 0, 12,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 8, 40, 0, 40, 54, 0, 0,
	40, 11, 0, 0, 0, 40, 0, 0, 0, 0,
	0, 0, 0, 0, 40, 33, 40, 4, 36, 0,
	50, 0, 0, 17, 18, 19, 0, 21, 22, 23,
	24, 25, 26, 27, 28, 0, 31, 0, 0, 40,
	44, 0, 54, 20, 29, 4, 4, 35, 0, 40,
	40, 40, 0, 30, 32, 34, 0, 57, 0, 0,
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
	42, 43, 44, 45, 46,
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
		//line mutan.y:54
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 12:
		//line mutan.y:55
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 13:
		//line mutan.y:59
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 14:
		//line mutan.y:60
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 15:
		//line mutan.y:64
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 16:
		//line mutan.y:65
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 17:
		//line mutan.y:69
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 18:
		//line mutan.y:70
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 19:
		//line mutan.y:71
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 20:
		//line mutan.y:72
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 21:
		//line mutan.y:73
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 22:
		//line mutan.y:74
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 23:
		//line mutan.y:75
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 24:
		//line mutan.y:76
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 25:
		//line mutan.y:77
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 26:
		//line mutan.y:78
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 27:
		//line mutan.y:79
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 28:
		//line mutan.y:80
		{ yyVAL.tnode = NewNode(GasTy) }
	case 29:
		//line mutan.y:81
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 30:
		//line mutan.y:83
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 31:
		//line mutan.y:91
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 32:
		//line mutan.y:101
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 33:
		//line mutan.y:104
		{ yyVAL.tnode = nil }
	case 34:
		//line mutan.y:109
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 35:
		//line mutan.y:114
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 36:
		//line mutan.y:119
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 37:
		//line mutan.y:125
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 38:
		//line mutan.y:126
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 39:
		//line mutan.y:127
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 40:
		//line mutan.y:128
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 41:
		//line mutan.y:132
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 42:
		//line mutan.y:133
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 43:
		//line mutan.y:139
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 44:
		//line mutan.y:145
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 45:
		//line mutan.y:149
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 46:
		//line mutan.y:154
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 47:
		//line mutan.y:155
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 48:
		//line mutan.y:156
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 49:
		//line mutan.y:161
		{
		
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      yyVAL.tnode.VarType = yyS[yypt-1].str
		  }
	case 50:
		//line mutan.y:170
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      yyVAL.tnode.VarType = yyS[yypt-4].str
		      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 51:
		//line mutan.y:180
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 52:
		//line mutan.y:184
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 53:
		//line mutan.y:185
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 54:
		//line mutan.y:186
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 55:
		//line mutan.y:187
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 56:
		//line mutan.y:188
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 57:
		//line mutan.y:192
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 58:
		//line mutan.y:196
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

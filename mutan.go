
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

//line mutan.y:209



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 42,
}

const yyNprod = 63
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 320

var yyAct = []int{

	20, 5, 10, 2, 19, 44, 139, 27, 34, 33,
	52, 59, 110, 36, 37, 38, 32, 66, 28, 80,
	31, 26, 60, 11, 30, 29, 25, 68, 69, 70,
	71, 72, 76, 14, 21, 141, 51, 138, 18, 34,
	33, 56, 57, 58, 43, 73, 74, 75, 77, 78,
	79, 62, 64, 65, 25, 3, 83, 82, 12, 84,
	13, 115, 152, 154, 148, 136, 8, 134, 135, 27,
	34, 33, 34, 33, 90, 42, 85, 89, 32, 87,
	28, 50, 31, 26, 156, 11, 30, 29, 25, 53,
	112, 114, 111, 149, 113, 14, 21, 119, 137, 127,
	18, 34, 33, 103, 27, 126, 128, 125, 131, 34,
	33, 132, 133, 32, 124, 28, 123, 31, 26, 122,
	121, 30, 29, 25, 120, 118, 34, 33, 117, 116,
	14, 21, 12, 91, 13, 18, 151, 144, 81, 143,
	8, 61, 107, 27, 146, 147, 102, 150, 140, 153,
	101, 100, 32, 99, 28, 155, 31, 26, 98, 11,
	30, 29, 25, 97, 12, 106, 13, 96, 142, 14,
	21, 94, 8, 86, 18, 27, 55, 93, 92, 34,
	33, 49, 145, 48, 32, 47, 28, 46, 31, 26,
	45, 11, 30, 29, 25, 35, 12, 39, 13, 54,
	108, 14, 21, 104, 8, 40, 18, 27, 95, 88,
	130, 109, 41, 34, 33, 129, 32, 7, 28, 22,
	31, 26, 24, 11, 30, 29, 25, 16, 12, 15,
	13, 67, 105, 14, 21, 23, 8, 9, 18, 27,
	6, 17, 4, 1, 0, 0, 0, 0, 32, 0,
	28, 0, 31, 26, 0, 11, 30, 29, 25, 0,
	12, 0, 13, 0, 0, 14, 21, 0, 8, 0,
	18, 27, 0, 0, 0, 0, 0, 0, 0, 0,
	32, 0, 28, 0, 31, 26, 0, 11, 30, 29,
	25, 0, 27, 0, 0, 0, 0, 14, 21, 0,
	0, 32, 18, 28, 0, 31, 26, 0, 0, 30,
	29, 25, 0, 0, 0, 0, 0, 0, 63, 21,
}
var yyPact = []int{

	23, -1000, 254, -1000, -1000, -38, -1000, -1000, 180, -1000,
	-1000, -10, -10, -10, 193, 208, -1000, -1000, 32, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -44, 175, 172, 170,
	168, 166, 56, -1000, -10, -35, -38, 80, 167, -10,
	-10, -10, -1000, -33, -9, 125, 275, 275, 275, -26,
	8, -38, 122, 23, -10, 23, -38, 63, -38, 160,
	-1000, -1000, 50, 197, 48, 45, 117, -1000, 163, 162,
	156, 196, 152, 148, 143, 138, 136, 135, 131, 88,
	191, -1000, 222, 133, 190, 207, -31, 275, -10, 275,
	18, -1000, 113, 112, 109, -10, 108, 104, 103, 100,
	98, 91, 89, 83, -10, 203, -10, 23, -1000, 87,
	-1000, 38, 55, 36, 82, -1000, -1000, -1000, -1000, 24,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -7, -1000,
	139, 26, 158, -1000, 275, -1000, 18, -1000, -1000, 178,
	23, 23, -1000, 35, 77, -10, 126, 52, 18, -1000,
	-38, -1000, -1000, 34, 18, 68, -1000,
}
var yyPgo = []int{

	0, 243, 3, 242, 1, 2, 241, 4, 240, 237,
	235, 231, 229, 227, 227, 227, 222, 219, 217, 215,
	0,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 10,
	10, 10, 10, 10, 10, 14, 14, 15, 15, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 8, 19, 19, 18, 18, 18, 4,
	4, 4, 4, 9, 9, 5, 5, 5, 5, 5,
	5, 12, 13, 6, 7, 7, 7, 7, 7, 20,
	20, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 1, 0, 1, 1, 1, 4, 3,
	12, 8, 6, 4, 3, 3, 0, 1, 0, 3,
	3, 3, 4, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 6, 6, 4, 0, 9, 7, 5, 1,
	1, 2, 0, 2, 3, 3, 6, 3, 1, 1,
	1, 2, 5, 1, 1, 1, 4, 1, 1, 1,
	1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 32, -3, -4, -8, -18, 14, -9,
	-5, 33, 6, 8, 43, -12, -13, -6, 48, -7,
	-20, 44, -17, -10, -16, 36, 31, 17, 28, 35,
	34, 30, 26, 47, 46, 15, -4, -4, -4, 4,
	12, 4, 43, 12, 49, 15, 15, 15, 15, 15,
	25, -4, 45, 9, 32, 9, -4, -4, -4, 44,
	31, 16, -7, 43, -7, -7, 43, -11, 19, 20,
	21, 22, 23, 37, 38, 39, 24, 40, 41, 42,
	11, 16, -2, -4, -2, 13, 13, 29, 12, 29,
	29, 16, 15, 15, 15, 12, 15, 15, 15, 15,
	15, 15, 15, 15, 12, 10, 32, 9, 10, 4,
	43, -7, -4, -7, -20, 43, 16, 16, 16, -4,
	16, 16, 16, 16, 16, 16, 16, 16, -4, -19,
	7, -4, -2, -5, 29, 13, 29, 16, 13, 13,
	9, 9, 10, -7, -20, 4, -2, -2, 29, 16,
	-4, 10, 10, -20, 29, -20, 16,
}
var yyDef = []int{

	4, -2, -2, 3, 2, 5, 6, 7, 0, 39,
	40, 42, 42, 42, 61, 48, 49, 50, 0, 53,
	54, 55, 57, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 0, 43, 42, 0, 41, 0, 0, 42,
	42, 42, 51, 0, 0, 0, 0, 0, 0, 0,
	0, 44, 0, 4, 42, 4, 45, 0, 47, 0,
	62, 9, 0, 61, 0, 0, 0, 14, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 8, 42, 0, 42, 56, 0, 0, 42, 0,
	0, 13, 0, 0, 0, 42, 0, 0, 0, 0,
	0, 0, 0, 0, 42, 35, 42, 4, 38, 0,
	52, 0, 0, 0, 0, 61, 19, 20, 21, 0,
	23, 24, 25, 26, 27, 28, 29, 30, 0, 33,
	0, 0, 42, 46, 0, 56, 0, 12, 22, 31,
	4, 4, 37, 0, 0, 42, 42, 42, 0, 11,
	32, 34, 36, 0, 0, 0, 10,
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
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 13:
		//line mutan.y:59
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 14:
		//line mutan.y:60
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 15:
		//line mutan.y:64
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 16:
		//line mutan.y:65
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 17:
		//line mutan.y:69
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 18:
		//line mutan.y:70
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 19:
		//line mutan.y:74
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 20:
		//line mutan.y:75
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 21:
		//line mutan.y:76
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 22:
		//line mutan.y:77
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 23:
		//line mutan.y:78
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 24:
		//line mutan.y:79
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 25:
		//line mutan.y:80
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 26:
		//line mutan.y:81
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 27:
		//line mutan.y:82
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 28:
		//line mutan.y:83
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 29:
		//line mutan.y:84
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 30:
		//line mutan.y:85
		{ yyVAL.tnode = NewNode(GasTy) }
	case 31:
		//line mutan.y:86
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 32:
		//line mutan.y:88
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 33:
		//line mutan.y:96
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 34:
		//line mutan.y:106
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 35:
		//line mutan.y:109
		{ yyVAL.tnode = nil }
	case 36:
		//line mutan.y:114
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 37:
		//line mutan.y:119
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 38:
		//line mutan.y:124
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 39:
		//line mutan.y:130
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 40:
		//line mutan.y:131
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 41:
		//line mutan.y:132
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 42:
		//line mutan.y:133
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 43:
		//line mutan.y:137
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 44:
		//line mutan.y:138
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 45:
		//line mutan.y:144
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 46:
		//line mutan.y:150
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 47:
		//line mutan.y:154
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 48:
		//line mutan.y:159
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 49:
		//line mutan.y:160
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 50:
		//line mutan.y:161
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 51:
		//line mutan.y:166
		{
		
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      yyVAL.tnode.VarType = yyS[yypt-1].str
		  }
	case 52:
		//line mutan.y:175
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      yyVAL.tnode.VarType = yyS[yypt-4].str
		      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 53:
		//line mutan.y:185
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 54:
		//line mutan.y:189
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 55:
		//line mutan.y:190
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 56:
		//line mutan.y:191
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 57:
		//line mutan.y:192
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 58:
		//line mutan.y:193
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 59:
		//line mutan.y:197
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 60:
		//line mutan.y:198
		{ yyVAL.tnode = NewNode(NilTy) }
	case 61:
		//line mutan.y:202
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 62:
		//line mutan.y:206
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

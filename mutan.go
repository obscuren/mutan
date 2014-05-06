
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
const DIFFICULTY = 57380
const PREVHASH = 57381
const TIMESTAMP = 57382
const BLOCKNUM = 57383
const COINBASE = 57384
const GAS = 57385
const ID = 57386
const NUMBER = 57387
const INLINE_ASM = 57388
const OP = 57389
const DOP = 57390
const TYPE = 57391
const STR = 57392

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

//line mutan.y:210



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 43,
}

const yyNprod = 64
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 343

var yyAct = []int{

	20, 5, 10, 2, 19, 44, 109, 27, 55, 52,
	144, 34, 33, 36, 37, 38, 32, 59, 28, 25,
	31, 26, 43, 11, 30, 29, 25, 117, 112, 108,
	142, 54, 66, 141, 14, 21, 51, 138, 53, 18,
	86, 56, 57, 58, 34, 33, 34, 33, 34, 33,
	3, 62, 64, 65, 42, 157, 84, 83, 12, 85,
	13, 60, 155, 151, 34, 33, 8, 34, 33, 27,
	139, 34, 33, 137, 34, 33, 34, 33, 32, 91,
	28, 50, 31, 26, 90, 11, 30, 29, 25, 88,
	159, 114, 116, 113, 152, 115, 14, 21, 121, 140,
	130, 18, 129, 128, 127, 105, 27, 126, 131, 125,
	134, 124, 123, 135, 136, 32, 122, 28, 120, 31,
	26, 119, 118, 30, 29, 25, 92, 82, 61, 104,
	103, 102, 101, 14, 21, 12, 100, 13, 18, 154,
	147, 99, 146, 8, 87, 98, 27, 149, 150, 97,
	153, 95, 156, 94, 93, 32, 49, 28, 158, 31,
	26, 48, 11, 30, 29, 25, 47, 46, 12, 45,
	13, 35, 145, 14, 21, 39, 8, 106, 18, 27,
	96, 89, 143, 40, 133, 148, 111, 41, 32, 132,
	28, 7, 31, 26, 22, 11, 30, 29, 25, 24,
	16, 12, 15, 13, 67, 110, 14, 21, 23, 8,
	9, 18, 27, 6, 17, 4, 1, 0, 0, 0,
	0, 32, 0, 28, 0, 31, 26, 0, 11, 30,
	29, 25, 0, 0, 12, 0, 13, 0, 107, 14,
	21, 0, 8, 0, 18, 27, 0, 0, 0, 0,
	0, 0, 0, 0, 32, 0, 28, 0, 31, 26,
	0, 11, 30, 29, 25, 0, 0, 12, 0, 13,
	0, 0, 14, 21, 0, 8, 0, 18, 27, 0,
	0, 0, 0, 0, 0, 0, 0, 32, 0, 28,
	81, 31, 26, 0, 11, 30, 29, 25, 68, 69,
	70, 71, 72, 76, 0, 14, 21, 0, 0, 0,
	18, 0, 0, 0, 27, 0, 79, 73, 74, 75,
	77, 78, 80, 32, 0, 28, 0, 31, 26, 0,
	0, 30, 29, 25, 0, 0, 0, 0, 0, 0,
	0, 63, 21,
}
var yyPact = []int{

	18, -1000, 261, -1000, -1000, -36, -1000, -1000, 156, -1000,
	-1000, -10, -10, -10, 171, 183, -1000, -1000, 10, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -45, 154, 152, 151,
	146, 141, 56, -1000, -10, -37, -36, 29, -1, -10,
	-10, -10, -1000, -28, 30, 112, 297, 297, 297, -12,
	279, -36, 111, 18, -10, 18, -36, 27, -36, 131,
	-1000, -1000, 60, 169, 55, 50, 110, -1000, 139, 138,
	136, 168, 134, 130, 126, 121, 117, 116, 115, 114,
	90, 165, -1000, 228, -3, 195, 182, -16, 297, -10,
	297, -17, -1000, 106, 105, 102, -10, 100, 96, 95,
	93, 91, 88, 87, 86, 84, -10, 177, -10, 18,
	-1000, 89, -1000, 44, 24, 41, 83, -1000, -1000, -1000,
	-1000, 20, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 17, -1000, 173, 1, 162, -1000, 297, -1000, -17,
	-1000, -1000, 181, 18, 18, -1000, 34, 78, -10, 129,
	52, -17, -1000, -36, -1000, -1000, 26, -17, 74, -1000,
}
var yyPgo = []int{

	0, 216, 3, 215, 1, 2, 214, 4, 213, 210,
	208, 204, 202, 200, 200, 200, 199, 194, 191, 189,
	0,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 10,
	10, 10, 10, 10, 10, 14, 14, 15, 15, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 8, 19, 19, 18, 18, 18,
	4, 4, 4, 4, 9, 9, 5, 5, 5, 5,
	5, 5, 12, 13, 6, 7, 7, 7, 7, 7,
	20, 20, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 1, 0, 1, 1, 1, 4, 3,
	12, 8, 6, 4, 3, 3, 0, 1, 0, 3,
	3, 3, 4, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 6, 6, 4, 0, 9, 7, 5,
	1, 1, 2, 0, 2, 3, 3, 6, 3, 1,
	1, 1, 2, 5, 1, 1, 1, 4, 1, 1,
	1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 32, -3, -4, -8, -18, 14, -9,
	-5, 33, 6, 8, 44, -12, -13, -6, 49, -7,
	-20, 45, -17, -10, -16, 36, 31, 17, 28, 35,
	34, 30, 26, 48, 47, 15, -4, -4, -4, 4,
	12, 4, 44, 12, 50, 15, 15, 15, 15, 15,
	25, -4, 46, 9, 32, 9, -4, -4, -4, 45,
	31, 16, -7, 44, -7, -7, 44, -11, 19, 20,
	21, 22, 23, 38, 39, 40, 24, 41, 42, 37,
	43, 11, 16, -2, -4, -2, 13, 13, 29, 12,
	29, 29, 16, 15, 15, 15, 12, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 12, 10, 32, 9,
	10, 4, 44, -7, -4, -7, -20, 44, 16, 16,
	16, -4, 16, 16, 16, 16, 16, 16, 16, 16,
	16, -4, -19, 7, -4, -2, -5, 29, 13, 29,
	16, 13, 13, 9, 9, 10, -7, -20, 4, -2,
	-2, 29, 16, -4, 10, 10, -20, 29, -20, 16,
}
var yyDef = []int{

	4, -2, -2, 3, 2, 5, 6, 7, 0, 40,
	41, 43, 43, 43, 62, 49, 50, 51, 0, 54,
	55, 56, 58, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 0, 44, 43, 0, 42, 0, 0, 43,
	43, 43, 52, 0, 0, 0, 0, 0, 0, 0,
	0, 45, 0, 4, 43, 4, 46, 0, 48, 0,
	63, 9, 0, 62, 0, 0, 0, 14, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 8, 43, 0, 43, 57, 0, 0, 43,
	0, 0, 13, 0, 0, 0, 43, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 43, 36, 43, 4,
	39, 0, 53, 0, 0, 0, 0, 62, 19, 20,
	21, 0, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 0, 34, 0, 0, 43, 47, 0, 57, 0,
	12, 22, 32, 4, 4, 38, 0, 0, 43, 43,
	43, 0, 11, 33, 35, 37, 0, 0, 0, 10,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50,
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
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 31:
		//line mutan.y:86
		{ yyVAL.tnode = NewNode(GasTy) }
	case 32:
		//line mutan.y:87
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 33:
		//line mutan.y:89
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 34:
		//line mutan.y:97
		{
		      if yyS[yypt-0].tnode == nil {
			    yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
		      } else {
			    yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
		      }
		  }
	case 35:
		//line mutan.y:107
		{
		      yyVAL.tnode = yyS[yypt-1].tnode
		  }
	case 36:
		//line mutan.y:110
		{ yyVAL.tnode = nil }
	case 37:
		//line mutan.y:115
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 38:
		//line mutan.y:120
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 39:
		//line mutan.y:125
		{
			  yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 40:
		//line mutan.y:131
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 41:
		//line mutan.y:132
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 42:
		//line mutan.y:133
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 43:
		//line mutan.y:134
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 44:
		//line mutan.y:138
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 45:
		//line mutan.y:139
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 46:
		//line mutan.y:145
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 47:
		//line mutan.y:151
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 48:
		//line mutan.y:155
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 49:
		//line mutan.y:160
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 50:
		//line mutan.y:161
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 51:
		//line mutan.y:162
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 52:
		//line mutan.y:167
		{
		
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      yyVAL.tnode.VarType = yyS[yypt-1].str
		  }
	case 53:
		//line mutan.y:176
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      yyVAL.tnode.VarType = yyS[yypt-4].str
		      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 54:
		//line mutan.y:186
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 55:
		//line mutan.y:190
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 56:
		//line mutan.y:191
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 57:
		//line mutan.y:192
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 58:
		//line mutan.y:193
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 59:
		//line mutan.y:194
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 60:
		//line mutan.y:198
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 61:
		//line mutan.y:199
		{ yyVAL.tnode = NewNode(NilTy) }
	case 62:
		//line mutan.y:203
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 63:
		//line mutan.y:207
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

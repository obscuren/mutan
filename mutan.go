
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
const DIFFICULTY = 57381
const PREVHASH = 57382
const TIMESTAMP = 57383
const BLOCKNUM = 57384
const COINBASE = 57385
const GAS = 57386
const ID = 57387
const NUMBER = 57388
const INLINE_ASM = 57389
const OP = 57390
const DOP = 57391
const TYPE = 57392
const STR = 57393
const BOOLEAN = 57394

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
	"BOOLEAN",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line mutan.y:218



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 43,
	-1, 54,
	48, 40,
	49, 40,
	-2, 47,
}

const yyNprod = 67
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 362

var yyAct = []int{

	20, 5, 10, 2, 19, 45, 35, 34, 56, 26,
	28, 63, 116, 37, 38, 39, 70, 113, 121, 33,
	3, 29, 64, 32, 27, 59, 44, 31, 30, 26,
	9, 161, 17, 155, 143, 141, 95, 55, 67, 21,
	112, 94, 60, 61, 62, 22, 92, 12, 58, 13,
	51, 159, 66, 68, 69, 8, 35, 34, 28, 43,
	88, 87, 163, 89, 35, 34, 54, 33, 53, 29,
	156, 32, 27, 144, 11, 31, 30, 26, 134, 133,
	146, 12, 132, 13, 131, 158, 14, 21, 130, 8,
	148, 18, 28, 22, 129, 118, 120, 117, 145, 119,
	142, 33, 125, 29, 128, 32, 27, 127, 11, 31,
	30, 26, 135, 90, 138, 35, 34, 139, 140, 126,
	14, 21, 124, 123, 122, 18, 96, 22, 86, 35,
	34, 57, 65, 35, 34, 35, 34, 91, 109, 12,
	108, 13, 107, 149, 151, 106, 150, 8, 35, 34,
	28, 153, 154, 105, 157, 104, 160, 103, 102, 33,
	101, 29, 162, 32, 27, 99, 11, 31, 30, 26,
	35, 34, 98, 12, 97, 13, 50, 114, 14, 21,
	49, 8, 40, 18, 28, 22, 48, 47, 46, 36,
	41, 110, 100, 33, 93, 29, 147, 32, 27, 137,
	11, 31, 30, 26, 152, 52, 115, 12, 42, 13,
	136, 111, 14, 21, 7, 8, 23, 18, 28, 22,
	25, 16, 15, 71, 24, 6, 4, 33, 1, 29,
	0, 32, 27, 0, 11, 31, 30, 26, 0, 0,
	0, 12, 0, 13, 0, 0, 14, 21, 0, 8,
	0, 18, 28, 22, 0, 0, 0, 0, 0, 0,
	0, 33, 0, 29, 0, 32, 27, 0, 11, 31,
	30, 26, 0, 0, 0, 28, 0, 0, 0, 0,
	14, 21, 0, 0, 33, 18, 29, 22, 32, 27,
	0, 11, 31, 30, 26, 0, 0, 0, 28, 0,
	0, 0, 0, 14, 21, 0, 0, 33, 18, 29,
	22, 32, 27, 0, 0, 31, 30, 26, 0, 0,
	0, 0, 0, 0, 0, 0, 14, 21, 85, 0,
	0, 18, 0, 22, 0, 0, 72, 73, 74, 75,
	76, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 83, 0, 77, 78, 79, 81,
	82, 84,
}
var yyPact = []int{

	-12, -1000, 235, -1000, -1000, -42, -1000, -1000, 174, -1000,
	-1000, 258, 258, 258, 178, 204, -1000, -1000, 14, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -46, 173, 172,
	171, 165, 161, 25, -1000, 258, -39, -42, 122, 16,
	258, 258, 258, -1000, -35, -9, 116, -7, -7, -7,
	-29, 317, -1000, -1000, -1000, -42, 112, -12, 258, -12,
	-42, 100, -42, 124, -1000, -1000, 17, 182, 12, 7,
	110, -1000, 159, 157, 150, 180, 145, 143, 142, 140,
	138, 130, 127, 125, 123, 179, -1000, 201, 8, 167,
	202, -33, -7, 258, -7, -27, -1000, 108, 107, 106,
	258, 103, 91, 88, 78, 72, 68, 66, 63, 62,
	258, 192, 258, -12, -1000, 281, -1000, 6, 87, 5,
	57, -1000, -1000, -1000, -1000, 85, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 67, -1000, 187, 81, 133,
	-1000, -7, -1000, -27, -1000, -1000, 200, -12, -12, -1000,
	4, 54, 258, 75, 41, -27, -1000, -42, -1000, -1000,
	2, -27, 46, -1000,
}
var yyPgo = []int{

	0, 228, 3, 226, 1, 2, 32, 4, 225, 30,
	224, 223, 222, 221, 221, 221, 220, 216, 214, 210,
	0, 205,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 10,
	10, 10, 10, 10, 10, 14, 14, 15, 15, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 8, 19, 19, 18, 18, 18,
	4, 4, 4, 4, 9, 9, 21, 21, 5, 5,
	5, 5, 5, 5, 12, 13, 6, 7, 7, 7,
	7, 7, 7, 20, 20, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 1, 0, 1, 1, 1, 4, 3,
	12, 8, 6, 4, 3, 3, 0, 1, 0, 3,
	3, 3, 4, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 6, 6, 4, 0, 9, 7, 5,
	1, 1, 2, 0, 2, 3, 1, 1, 3, 6,
	3, 1, 1, 1, 2, 5, 1, 1, 1, 4,
	1, 1, 1, 1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 32, -3, -4, -8, -18, 14, -9,
	-5, 33, 6, 8, 45, -12, -13, -6, 50, -7,
	-20, 46, 52, -17, -10, -16, 36, 31, 17, 28,
	35, 34, 30, 26, 49, 48, 15, -4, -4, -4,
	4, 12, 4, 45, 12, 51, 15, 15, 15, 15,
	15, 25, -21, -6, -9, -4, 47, 9, 32, 9,
	-4, -4, -4, 46, 31, 16, -7, 45, -7, -7,
	45, -11, 19, 20, 21, 22, 23, 39, 40, 41,
	24, 42, 43, 37, 44, 11, 16, -2, -4, -2,
	13, 13, 29, 12, 29, 29, 16, 15, 15, 15,
	12, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	12, 10, 32, 9, 10, 4, 45, -7, -4, -7,
	-20, 45, 16, 16, 16, -4, 16, 16, 16, 16,
	16, 16, 16, 16, 16, -4, -19, 7, -4, -2,
	-5, 29, 13, 29, 16, 13, 13, 9, 9, 10,
	-7, -20, 4, -2, -2, 29, 16, -4, 10, 10,
	-20, 29, -20, 16,
}
var yyDef = []int{

	4, -2, -2, 3, 2, 5, 6, 7, 0, 40,
	41, 43, 43, 43, 65, 51, 52, 53, 0, 56,
	57, 58, 60, 61, 62, 63, 64, 0, 0, 0,
	0, 0, 0, 0, 44, 43, 0, 42, 0, 0,
	43, 43, 43, 54, 0, 0, 0, 0, 0, 0,
	0, 0, 45, 46, -2, 0, 0, 4, 43, 4,
	48, 0, 50, 0, 66, 9, 0, 65, 0, 0,
	0, 14, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 8, 43, 0, 43,
	59, 0, 0, 43, 0, 0, 13, 0, 0, 0,
	43, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	43, 36, 43, 4, 39, 0, 55, 0, 0, 0,
	0, 65, 19, 20, 21, 0, 23, 24, 25, 26,
	27, 28, 29, 30, 31, 0, 34, 0, 0, 43,
	49, 0, 59, 0, 12, 22, 32, 4, 4, 38,
	0, 0, 43, 43, 43, 0, 11, 33, 35, 37,
	0, 0, 0, 10,
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
	52,
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
		//line mutan.y:139
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 45:
		//line mutan.y:141
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 46:
		//line mutan.y:145
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 47:
		//line mutan.y:146
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 48:
		//line mutan.y:151
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 49:
		//line mutan.y:157
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 50:
		//line mutan.y:161
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 51:
		//line mutan.y:166
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 52:
		//line mutan.y:167
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 53:
		//line mutan.y:168
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 54:
		//line mutan.y:173
		{
		
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      yyVAL.tnode.VarType = yyS[yypt-1].str
		  }
	case 55:
		//line mutan.y:183
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      yyVAL.tnode.VarType = yyS[yypt-4].str
		      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 56:
		//line mutan.y:193
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 57:
		//line mutan.y:197
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 58:
		//line mutan.y:198
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 59:
		//line mutan.y:199
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 60:
		//line mutan.y:200
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 61:
		//line mutan.y:201
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 62:
		//line mutan.y:202
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 63:
		//line mutan.y:206
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 64:
		//line mutan.y:207
		{ yyVAL.tnode = NewNode(NilTy) }
	case 65:
		//line mutan.y:211
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 66:
		//line mutan.y:215
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

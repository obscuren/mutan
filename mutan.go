
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
const LEFT_BRACES = 57349
const RIGHT_BRACES = 57350
const STORE = 57351
const LEFT_BRACKET = 57352
const RIGHT_BRACKET = 57353
const ASM = 57354
const LEFT_PAR = 57355
const RIGHT_PAR = 57356
const STOP = 57357
const ADDR = 57358
const ORIGIN = 57359
const CALLER = 57360
const CALLVAL = 57361
const CALLDATALOAD = 57362
const CALLDATASIZE = 57363
const GASPRICE = 57364
const DOT = 57365
const THIS = 57366
const ARRAY = 57367
const CALL = 57368
const COMMA = 57369
const SIZEOF = 57370
const QUOTE = 57371
const ID = 57372
const NUMBER = 57373
const INLINE_ASM = 57374
const OP = 57375
const TYPE = 57376
const STR = 57377

var yyToknames = []string{
	"ASSIGN",
	"EQUAL",
	"IF",
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
	"ID",
	"NUMBER",
	"INLINE_ASM",
	"OP",
	"TYPE",
	"STR",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line mutan.y:155



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	33, 24,
	-2, 1,
}

const yyNprod = 44
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 131

var yyAct = []int{

	17, 16, 4, 35, 8, 2, 93, 26, 92, 64,
	41, 47, 28, 97, 79, 63, 53, 48, 9, 98,
	76, 13, 100, 42, 6, 34, 95, 22, 26, 40,
	26, 26, 43, 44, 45, 46, 25, 26, 23, 50,
	24, 21, 10, 18, 91, 33, 15, 66, 62, 26,
	39, 88, 87, 86, 85, 9, 84, 83, 13, 69,
	61, 6, 49, 75, 22, 74, 73, 72, 80, 71,
	81, 82, 70, 25, 38, 23, 37, 24, 21, 10,
	18, 90, 89, 15, 36, 13, 27, 78, 65, 29,
	68, 22, 67, 94, 32, 30, 96, 52, 77, 99,
	25, 31, 23, 22, 24, 21, 10, 18, 19, 12,
	15, 11, 25, 54, 23, 20, 24, 21, 51, 18,
	55, 56, 57, 58, 59, 60, 7, 5, 14, 3,
	1,
}
var yyPact = []int{

	-1000, -1000, 49, -1000, -26, -1000, 73, -1000, -1000, 76,
	85, 97, -1000, 84, -1000, 15, -1000, -1000, -1000, -1000,
	-1000, -32, 71, 63, 61, 27, 76, -22, 16, 76,
	76, 76, 76, -1000, -20, -12, 48, 88, -14, 103,
	-26, 46, -1000, -26, 4, -26, -2, 77, -1000, -1000,
	20, 82, 80, 45, -1000, 59, 56, 54, 53, 52,
	50, -1000, 12, 94, 83, -16, 88, 76, 76, -1000,
	43, 42, 40, 39, 38, 37, -1000, 76, 76, -1000,
	17, -3, -5, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-26, 88, -1000, -1000, -1, -17, -8, -1000, -17, 8,
	-1000,
}
var yyPgo = []int{

	0, 130, 5, 129, 2, 4, 128, 1, 127, 126,
	115, 113, 111, 109, 109, 109, 0, 108,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 10, 10, 10,
	10, 14, 14, 15, 15, 11, 11, 11, 11, 11,
	11, 8, 4, 4, 4, 9, 5, 5, 5, 5,
	5, 5, 5, 12, 13, 6, 7, 7, 7, 7,
	7, 7, 16, 17,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 4, 3, 12, 4,
	3, 3, 0, 1, 0, 3, 3, 3, 3, 3,
	3, 5, 1, 1, 0, 3, 3, 6, 3, 1,
	1, 6, 1, 2, 5, 1, 1, 1, 4, 4,
	1, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -8, 12, -9, -5, 6,
	30, -12, -13, 9, -6, 34, -7, -16, 31, -17,
	-10, 29, 15, 26, 28, 24, 33, 13, -4, 4,
	10, 4, 10, 30, 10, 35, 13, 13, 13, 23,
	-4, 32, 7, -4, -4, -4, -4, 31, 29, 14,
	-7, 30, 9, 30, -11, 17, 18, 19, 20, 21,
	22, 14, -2, 11, 11, 11, 27, 10, 10, 14,
	13, 13, 13, 13, 13, 13, 8, 4, 4, 30,
	-7, -4, -4, 14, 14, 14, 14, 14, 14, -5,
	-4, 27, 11, 11, -7, 27, -16, 30, 27, -16,
	14,
}
var yyDef = []int{

	3, -2, -2, 2, 4, 5, 0, 22, 23, 24,
	42, 29, 30, 0, 32, 0, 35, 36, 37, 40,
	41, 0, 0, 0, 0, 0, 24, 0, 0, 24,
	24, 24, 24, 33, 0, 0, 0, 0, 0, 0,
	25, 0, 3, 26, 0, 28, 0, 0, 43, 7,
	0, 42, 0, 0, 10, 0, 0, 0, 0, 0,
	0, 6, 24, 38, 39, 0, 0, 24, 24, 9,
	0, 0, 0, 0, 0, 0, 21, 0, 24, 34,
	0, 0, 0, 15, 16, 17, 18, 19, 20, 27,
	31, 0, 38, 39, 0, 0, 0, 42, 0, 0,
	8,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35,
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
		//line mutan.y:27
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line mutan.y:31
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode) }
	case 3:
		//line mutan.y:32
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line mutan.y:36
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 5:
		//line mutan.y:37
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 6:
		//line mutan.y:38
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 7:
		//line mutan.y:43
		{ yyVAL.tnode = NewNode(StopTy) }
	case 8:
		//line mutan.y:46
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 9:
		//line mutan.y:49
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 10:
		//line mutan.y:50
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 11:
		//line mutan.y:54
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 12:
		//line mutan.y:55
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 13:
		//line mutan.y:59
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 14:
		//line mutan.y:60
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 15:
		//line mutan.y:64
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 16:
		//line mutan.y:65
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 17:
		//line mutan.y:66
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 18:
		//line mutan.y:67
		{ yyVAL.tnode = NewNode(CallDataLoadTy) }
	case 19:
		//line mutan.y:68
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 20:
		//line mutan.y:69
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 21:
		//line mutan.y:73
		{ yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 22:
		//line mutan.y:77
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 23:
		//line mutan.y:78
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 24:
		//line mutan.y:79
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 25:
		//line mutan.y:83
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 26:
		//line mutan.y:89
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 27:
		//line mutan.y:95
		{
		      yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
		  }
	case 28:
		//line mutan.y:99
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 29:
		//line mutan.y:104
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 30:
		//line mutan.y:105
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 31:
		//line mutan.y:107
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 32:
		//line mutan.y:111
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 33:
		//line mutan.y:116
		{
		
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      yyVAL.tnode.VarType = yyS[yypt-1].str
		  }
	case 34:
		//line mutan.y:125
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      yyVAL.tnode.VarType = yyS[yypt-4].str
		      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 35:
		//line mutan.y:135
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 36:
		//line mutan.y:139
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 37:
		//line mutan.y:140
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 38:
		//line mutan.y:141
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 39:
		//line mutan.y:142
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 40:
		//line mutan.y:143
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 41:
		//line mutan.y:144
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 42:
		//line mutan.y:148
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 43:
		//line mutan.y:152
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}

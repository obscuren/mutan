
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
const ID = 57368
const NUMBER = 57369
const INLINE_ASM = 57370
const OP = 57371
const TYPE = 57372

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
	"ID",
	"NUMBER",
	"INLINE_ASM",
	"OP",
	"TYPE",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line mutan.y:124



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	29, 18,
	-2, 1,
}

const yyNprod = 33
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 81

var yyAct = []int{

	8, 21, 2, 9, 48, 56, 13, 32, 37, 6,
	9, 58, 19, 13, 30, 33, 6, 28, 55, 19,
	64, 20, 21, 10, 17, 34, 35, 15, 20, 13,
	10, 17, 4, 27, 15, 19, 47, 21, 63, 62,
	61, 60, 23, 59, 20, 46, 10, 17, 38, 54,
	15, 53, 52, 51, 31, 50, 49, 29, 65, 36,
	40, 41, 42, 43, 44, 45, 22, 26, 57, 25,
	24, 12, 11, 39, 18, 7, 5, 16, 14, 3,
	1,
}
var yyPact = []int{

	-1000, -1000, 4, -1000, -28, -1000, 53, -1000, -1000, 20,
	66, 65, -1000, 57, -1000, 7, -1000, -1000, -1000, 44,
	-9, 20, -21, 8, 20, 20, 20, -1000, -19, 34,
	43, -28, 31, -1000, -1000, -1000, -7, 45, -1000, -1000,
	42, 40, 39, 38, 36, 5, -1000, -3, 64, -15,
	29, 27, 26, 25, 24, 6, -1000, 20, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 80, 2, 79, 32, 0, 78, 77, 76, 75,
	74, 73, 72, 71,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 10, 10, 11,
	11, 11, 11, 11, 11, 8, 4, 4, 4, 9,
	5, 5, 5, 5, 5, 5, 12, 13, 6, 7,
	7, 7, 7,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 4, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 1, 1, 0, 3,
	3, 3, 1, 1, 6, 1, 2, 5, 1, 1,
	1, 4, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -8, 12, -9, -5, 6,
	26, -12, -13, 9, -6, 30, -7, 27, -10, 15,
	24, 29, 13, -4, 4, 4, 10, 26, 10, 13,
	23, -4, 28, 7, -5, -5, -4, 27, 14, -11,
	17, 18, 19, 20, 21, 22, 14, -2, 11, 11,
	13, 13, 13, 13, 13, 13, 8, 4, 26, 14,
	14, 14, 14, 14, 14, -5,
}
var yyDef = []int{

	3, -2, -2, 2, 4, 5, 0, 16, 17, 18,
	29, 22, 23, 0, 25, 0, 28, 30, 32, 0,
	0, 18, 0, 0, 0, 0, 18, 26, 0, 0,
	0, 19, 0, 3, 20, 21, 0, 0, 7, 8,
	0, 0, 0, 0, 0, 0, 6, 18, 31, 0,
	0, 0, 0, 0, 0, 0, 15, 0, 27, 9,
	10, 11, 12, 13, 14, 24,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30,
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
		//line mutan.y:44
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 9:
		//line mutan.y:48
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 10:
		//line mutan.y:49
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 11:
		//line mutan.y:50
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 12:
		//line mutan.y:51
		{ yyVAL.tnode = NewNode(CallDataLoadTy) }
	case 13:
		//line mutan.y:52
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 14:
		//line mutan.y:53
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 15:
		//line mutan.y:57
		{ yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 16:
		//line mutan.y:61
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 17:
		//line mutan.y:62
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 18:
		//line mutan.y:63
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 19:
		//line mutan.y:67
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 20:
		//line mutan.y:72
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 21:
		//line mutan.y:78
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].tnode.Constant
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
		  }
	case 22:
		//line mutan.y:83
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 23:
		//line mutan.y:84
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 24:
		//line mutan.y:86
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 25:
		//line mutan.y:90
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 26:
		//line mutan.y:95
		{
		
		      yyVAL.tnode = NewNode(NewVarTy)
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      yyVAL.tnode.VarType = yyS[yypt-1].str
		  }
	case 27:
		//line mutan.y:104
		{
		      yyVAL.tnode = NewNode(NewArrayTy)
		      yyVAL.tnode.VarType = yyS[yypt-4].str
		      yyVAL.tnode.Size = yyS[yypt-2].str
		      yyVAL.tnode.Constant = yyS[yypt-0].str
		      
		  }
	case 28:
		//line mutan.y:114
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 29:
		//line mutan.y:118
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 30:
		//line mutan.y:119
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 31:
		//line mutan.y:120
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 32:
		//line mutan.y:121
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	}
	goto yystack /* stack new state and value */
}
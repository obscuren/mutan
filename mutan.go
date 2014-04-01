
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
const ID = 57358
const NUMBER = 57359
const INLINE_ASM = 57360
const OP = 57361

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
	"ID",
	"NUMBER",
	"INLINE_ASM",
	"OP",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line mutan.y:83



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 20
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 45

var yyAct = []int{

	10, 2, 5, 31, 17, 24, 11, 29, 32, 13,
	25, 17, 7, 22, 20, 8, 12, 16, 19, 18,
	23, 11, 27, 26, 13, 28, 33, 7, 30, 13,
	8, 12, 16, 21, 34, 17, 12, 16, 4, 9,
	6, 15, 14, 3, 1,
}
var yyPact = []int{

	-1000, -1000, 15, -1000, -1000, -15, -1000, 6, 5, -1000,
	-1000, 20, 29, 3, -1000, -1000, -1000, 20, -13, -4,
	16, 20, 20, -15, -7, -1000, -1000, -1000, -8, -1000,
	0, 22, -1000, 20, -1000,
}
var yyPgo = []int{

	0, 44, 1, 43, 2, 0, 42, 41, 40, 39,
	38,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 10, 8,
	4, 4, 9, 5, 5, 5, 6, 7, 7, 7,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 1, 1, 4, 3, 5,
	1, 1, 3, 3, 6, 1, 1, 1, 1, 4,
}
var yyChk = []int{

	-1000, -1, -2, -3, -10, -4, -8, 12, 15, -9,
	-5, 6, 16, 9, -6, -7, 17, 19, 13, 13,
	-4, 4, 10, -4, 18, 14, 7, -5, -4, 14,
	-2, 11, 8, 4, -5,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 5, 6, 0, 0, 10,
	11, 0, 17, 0, 15, 16, 18, 0, 0, 0,
	0, 0, 0, 12, 0, 8, 3, 13, 0, 7,
	0, 19, 9, 0, 14,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19,
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
		//line mutan.y:26
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line mutan.y:30
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode) }
	case 3:
		//line mutan.y:31
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line mutan.y:35
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 5:
		//line mutan.y:36
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 6:
		//line mutan.y:37
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 7:
		//line mutan.y:38
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 8:
		//line mutan.y:42
		{ yyVAL.tnode = NewNode(StopTy) }
	case 9:
		//line mutan.y:46
		{ yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 10:
		//line mutan.y:50
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 11:
		//line mutan.y:51
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 12:
		//line mutan.y:55
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 13:
		//line mutan.y:60
		{
		      node := NewNode(SetLocalTy)
		      node.Constant = yyS[yypt-2].str
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 14:
		//line mutan.y:66
		{
		      node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
		      yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
		  }
	case 15:
		//line mutan.y:70
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 16:
		//line mutan.y:74
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 17:
		//line mutan.y:78
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 18:
		//line mutan.y:79
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 19:
		//line mutan.y:80
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	}
	goto yystack /* stack new state and value */
}

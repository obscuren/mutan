
//line mutan.y:2
package mutan
import __yyfmt__ "fmt"
//line mutan.y:2
		
import (
_	"fmt"
)

var Tree *SyntaxTree

// Helper function to turn a tree in to a regular list.
// Especially handy when parsing argument lists
func makeSlice(tree *SyntaxTree) (ret []*SyntaxTree) {
	if tree != nil && tree.Type != EmptyTy {
		ret = append(ret, tree)
		for _, i := range tree.Children {
			ret = append(ret, makeSlice(i)...)
		}
		tree.Children = nil
	} 

	return
}

func makeArgs(tree *SyntaxTree, reverse bool) (ret []*SyntaxTree) {
	l := makeSlice(tree)
	// Quick reverse
	if reverse {
		for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
			l[i], l[j] = l[j], l[i]
		}
	}

	for _, s := range l {
		if s.Type != ArgTy {
			ret = append(ret, s)
		}
	}

	return
}


//line mutan.y:44
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
const PUSH = 57385
const POP = 57386
const DIFFICULTY = 57387
const PREVHASH = 57388
const TIMESTAMP = 57389
const BLOCKNUM = 57390
const COINBASE = 57391
const GAS = 57392
const VAR = 57393
const FUNC = 57394
const FUNC_CALL = 57395
const IMPORT = 57396
const ID = 57397
const NUMBER = 57398
const INLINE_ASM = 57399
const OP = 57400
const DOP = 57401
const STR = 57402
const BOOLEAN = 57403
const CODE = 57404
const oper = 57405
const AND = 57406
const MUL = 57407

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
	"PUSH",
	"POP",
	"DIFFICULTY",
	"PREVHASH",
	"TIMESTAMP",
	"BLOCKNUM",
	"COINBASE",
	"GAS",
	"VAR",
	"FUNC",
	"FUNC_CALL",
	"IMPORT",
	"ID",
	"NUMBER",
	"INLINE_ASM",
	"OP",
	"DOP",
	"STR",
	"BOOLEAN",
	"CODE",
	"oper",
	"AND",
	"MUL",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line mutan.y:313



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 6,
	-1, 28,
	58, 66,
	64, 66,
	65, 66,
	-2, 75,
}

const yyNprod = 92
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 353

var yyAct = []int{

	23, 32, 2, 17, 80, 25, 160, 62, 28, 129,
	39, 7, 31, 63, 64, 189, 71, 61, 82, 46,
	37, 40, 100, 43, 38, 190, 20, 42, 41, 37,
	205, 167, 107, 59, 60, 19, 44, 45, 24, 171,
	70, 15, 99, 67, 30, 52, 49, 158, 18, 33,
	188, 164, 101, 38, 34, 54, 222, 16, 29, 215,
	224, 89, 161, 55, 199, 198, 84, 85, 93, 93,
	93, 91, 91, 91, 104, 104, 104, 97, 98, 93,
	93, 93, 159, 68, 88, 140, 103, 105, 106, 108,
	133, 56, 135, 69, 139, 130, 138, 79, 132, 39,
	134, 90, 95, 96, 92, 92, 92, 216, 46, 200,
	40, 6, 43, 38, 185, 20, 42, 41, 37, 54,
	4, 5, 184, 220, 19, 44, 45, 55, 4, 5,
	53, 57, 58, 30, 183, 182, 181, 18, 33, 104,
	104, 104, 170, 34, 93, 93, 16, 29, 166, 180,
	179, 168, 169, 178, 177, 56, 219, 175, 174, 176,
	173, 4, 5, 104, 172, 142, 191, 218, 141, 186,
	109, 193, 4, 5, 165, 102, 196, 163, 156, 4,
	5, 155, 4, 5, 154, 153, 152, 151, 150, 149,
	148, 146, 145, 144, 143, 81, 204, 78, 77, 104,
	104, 209, 76, 75, 93, 74, 211, 73, 72, 213,
	214, 208, 212, 35, 202, 201, 104, 221, 21, 197,
	22, 137, 217, 104, 223, 157, 12, 51, 131, 39,
	126, 147, 136, 48, 128, 207, 206, 203, 46, 87,
	40, 50, 43, 38, 14, 20, 42, 41, 37, 195,
	3, 8, 210, 162, 19, 44, 45, 47, 86, 66,
	65, 39, 187, 30, 11, 83, 13, 18, 33, 127,
	46, 194, 40, 34, 43, 38, 16, 29, 42, 41,
	37, 10, 27, 26, 110, 36, 9, 44, 45, 1,
	0, 0, 0, 39, 0, 30, 0, 0, 0, 192,
	33, 0, 46, 0, 40, 34, 43, 38, 0, 29,
	42, 41, 37, 125, 0, 0, 0, 0, 0, 44,
	45, 111, 113, 114, 115, 116, 120, 0, 0, 0,
	0, 94, 33, 0, 0, 0, 0, 34, 0, 123,
	0, 29, 0, 112, 0, 0, 0, 117, 118, 119,
	121, 122, 124,
}
var yyPact = []int{

	-1000, -1000, 105, 212, -1000, -1000, 105, -1000, 221, -1000,
	-1000, -9, 232, 22, -1000, -1000, -10, -1000, 115, 212,
	212, 82, 82, -42, -51, 256, 255, -1000, -1000, -12,
	28, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -44, 193,
	192, 190, 188, 187, 183, 182, 72, -1000, -58, 180,
	-39, -1000, -1000, -1000, 82, 82, 254, -1000, -1000, 230,
	52, -1000, 276, 276, 276, 82, 82, -1000, -1000, -13,
	-34, 21, 159, 276, 276, 276, -23, 82, 154, 302,
	217, -1000, 224, -7, -1000, 215, 82, -1000, 82, -1000,
	-51, -1000, -1000, -1000, 220, -51, -51, -1000, -1000, -1000,
	208, -1000, -1000, 67, -1000, 65, 56, 152, 149, -1000,
	-1000, 179, 178, 177, 176, 219, 175, 174, 173, 172,
	171, 170, 169, 166, 163, 213, -1000, 31, -1000, -1000,
	33, 249, -1000, 167, 19, 164, 82, -24, 276, 276,
	-16, -1000, -1000, 148, 144, 142, 141, 82, 138, 137,
	134, 133, 120, 119, 118, 106, 98, 82, -1, -40,
	-1000, -1000, 244, 242, 82, -1000, 206, -1000, 36, 35,
	93, -1000, -1000, -1000, -1000, -1000, 202, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 201, 228, -1000, 33,
	-25, -1000, 51, -1000, -1000, 227, 226, -1000, 276, -16,
	-1000, -1000, 248, -1000, -1000, 33, -1000, -1000, 30, 91,
	82, 157, -1000, 146, 113, -16, -1000, -1000, -1000, -1000,
	-1000, 27, -16, 44, -1000,
}
var yyPgo = []int{

	0, 289, 2, 111, 11, 3, 8, 12, 286, 41,
	285, 284, 283, 282, 282, 6, 0, 213, 281, 271,
	1, 38, 269, 265, 5, 250, 262,
}
var yyR1 = []int{

	0, 1, 2, 2, 25, 25, 25, 3, 3, 3,
	3, 3, 3, 3, 3, 22, 22, 22, 26, 26,
	10, 10, 10, 10, 10, 10, 10, 10, 14, 14,
	15, 15, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 8, 19,
	19, 18, 18, 4, 4, 4, 4, 4, 4, 4,
	23, 23, 9, 9, 9, 9, 21, 21, 5, 5,
	5, 5, 5, 5, 5, 5, 12, 12, 13, 6,
	7, 7, 7, 7, 7, 7, 7, 24, 20, 20,
	16, 17,
}
var yyR2 = []int{

	0, 1, 4, 0, 1, 1, 0, 1, 4, 1,
	1, 9, 4, 2, 1, 4, 5, 0, 1, 0,
	3, 12, 8, 6, 4, 4, 3, 3, 3, 0,
	1, 0, 3, 3, 3, 3, 4, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 6, 6, 4,
	0, 9, 5, 1, 2, 1, 4, 2, 2, 0,
	3, 0, 2, 3, 3, 3, 1, 1, 3, 3,
	6, 3, 4, 1, 1, 1, 2, 3, 5, 1,
	1, 1, 1, 4, 1, 1, 1, 2, 1, 1,
	1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -25, 15, 16, -3, -4, 39, -8,
	-18, 52, 14, 54, 32, -9, 64, -5, 55, 42,
	33, 6, 8, -16, -21, -24, -12, -13, -6, 65,
	51, -7, -20, 56, 61, -17, -10, 36, 31, 17,
	28, 35, 34, 30, 43, 44, 26, -25, 12, 55,
	9, -17, 55, 15, 4, 12, 40, -3, -3, -4,
	-4, 59, 58, 64, 65, 4, 4, 55, 55, 65,
	12, 60, 15, 15, 15, 15, 15, 15, 15, 25,
	62, 15, 57, -23, -4, -4, 4, 9, 32, 9,
	-21, -6, -9, -24, 55, -21, -21, -4, -4, 55,
	56, 31, 16, -7, -16, -7, -7, 55, -4, 16,
	-11, 19, 41, 20, 21, 22, 23, 45, 46, 47,
	24, 48, 49, 37, 50, 11, 13, -22, 10, 16,
	-4, 13, -4, -2, -4, -2, 12, 13, 29, 29,
	29, 16, 16, 15, 15, 15, 15, 12, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 12, 16, 51,
	-15, 29, 4, 10, 32, 10, -4, 55, -7, -7,
	-20, 55, 16, 16, 16, 16, -4, 16, 16, 16,
	16, 16, 16, 16, 16, 16, -4, -26, 51, 55,
	65, -5, 55, -6, -19, 7, -4, 13, 29, 29,
	16, 13, 13, 9, -15, 55, 9, 9, -7, -20,
	4, -2, -15, -2, -2, 29, 16, -4, 10, 10,
	10, -20, 29, -20, 16,
}
var yyDef = []int{

	3, -2, -2, 59, 4, 5, 6, 7, 0, 9,
	10, 0, 0, 0, 14, 53, 0, 55, 90, 59,
	59, 59, 59, 88, 0, 81, 73, 74, -2, 0,
	0, 79, 80, 82, 84, 85, 86, 89, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 2, 0, 0,
	0, 13, 54, 61, 59, 59, 0, 57, 58, 0,
	0, 62, 0, 0, 0, 59, 59, 87, 76, 0,
	0, 0, 0, 0, 0, 0, 0, 59, 0, 0,
	0, 17, 0, 59, 69, 0, 59, 3, 59, 3,
	63, 66, 67, 81, 90, 64, 65, 68, 71, 77,
	0, 91, 20, 0, 88, 0, 0, 0, 0, 26,
	27, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 8, 0, 12, 56,
	31, 83, 72, 6, 0, 6, 59, 0, 0, 0,
	0, 24, 25, 0, 0, 0, 0, 59, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 59, 19, 0,
	60, 30, 0, 50, 59, 52, 0, 78, 0, 0,
	0, 90, 32, 33, 34, 35, 0, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 0, 0, 18, 31,
	0, 70, 90, 75, 48, 0, 0, 83, 0, 0,
	23, 36, 46, 3, 15, 31, 3, 3, 0, 0,
	59, 6, 16, 6, 6, 0, 22, 47, 11, 49,
	51, 0, 0, 0, 21,
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
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65,
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
		//line mutan.y:65
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line mutan.y:69
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 3:
		//line mutan.y:70
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line mutan.y:74
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 5:
		//line mutan.y:75
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 6:
		//line mutan.y:76
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 7:
		//line mutan.y:80
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 8:
		//line mutan.y:81
		{ yyVAL.tnode = NewNode(LambdaTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 9:
		//line mutan.y:82
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 10:
		//line mutan.y:83
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 11:
		//line mutan.y:85
		{
				yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode);
				yyVAL.tnode.Constant = yyS[yypt-7].str
				yyVAL.tnode.HasRet = yyS[yypt-3].check
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-5].tnode, true)
			}
	case 12:
		//line mutan.y:91
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 13:
		//line mutan.y:92
		{ yyVAL.tnode = NewNode(ImportTy, yyS[yypt-0].tnode) }
	case 14:
		//line mutan.y:93
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 15:
		//line mutan.y:97
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 16:
		//line mutan.y:98
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-4].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str; yyVAL.tnode.Ptr = true }
	case 17:
		//line mutan.y:99
		{ yyVAL.tnode = nil }
	case 18:
		//line mutan.y:104
		{ yyVAL.check = true }
	case 19:
		//line mutan.y:105
		{ yyVAL.check = false }
	case 20:
		//line mutan.y:110
		{ yyVAL.tnode = NewNode(StopTy) }
	case 21:
		//line mutan.y:113
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 22:
		//line mutan.y:117
		{
			  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 23:
		//line mutan.y:120
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 24:
		//line mutan.y:121
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 25:
		//line mutan.y:122
		{ yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode) }
	case 26:
		//line mutan.y:123
		{ yyVAL.tnode = NewNode(PopTy) }
	case 27:
		//line mutan.y:124
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 28:
		//line mutan.y:128
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 29:
		//line mutan.y:129
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 30:
		//line mutan.y:133
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 31:
		//line mutan.y:134
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 32:
		//line mutan.y:139
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 33:
		//line mutan.y:140
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 34:
		//line mutan.y:141
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 35:
		//line mutan.y:142
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 36:
		//line mutan.y:143
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 37:
		//line mutan.y:144
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 38:
		//line mutan.y:145
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 39:
		//line mutan.y:146
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 40:
		//line mutan.y:147
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 41:
		//line mutan.y:148
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 42:
		//line mutan.y:149
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 43:
		//line mutan.y:150
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 44:
		//line mutan.y:151
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 45:
		//line mutan.y:152
		{ yyVAL.tnode = NewNode(GasTy) }
	case 46:
		//line mutan.y:153
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 47:
		//line mutan.y:155
		{
				node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 48:
		//line mutan.y:163
		{
				if yyS[yypt-0].tnode == nil {
					yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
				} else {
					yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
				}
			}
	case 49:
		//line mutan.y:173
		{
				yyVAL.tnode = yyS[yypt-1].tnode
			}
	case 50:
		//line mutan.y:176
		{ yyVAL.tnode = nil }
	case 51:
		//line mutan.y:181
		{
				yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
			}
	case 52:
		//line mutan.y:185
		{
				yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
			}
	case 53:
		//line mutan.y:191
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 54:
		//line mutan.y:192
		{ yyVAL.tnode = NewNode(RefTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 55:
		//line mutan.y:193
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 56:
		//line mutan.y:195
		{
				yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
				yyVAL.tnode.Constant = yyS[yypt-3].str
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
			}
	case 57:
		//line mutan.y:200
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 58:
		//line mutan.y:201
		{ yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode) }
	case 59:
		//line mutan.y:202
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 60:
		//line mutan.y:206
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode);}
	case 61:
		//line mutan.y:207
		{ yyVAL.tnode = nil }
	case 62:
		//line mutan.y:212
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 63:
		//line mutan.y:214
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 64:
		//line mutan.y:215
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 65:
		//line mutan.y:216
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 66:
		//line mutan.y:221
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 67:
		//line mutan.y:222
		{ yyVAL.tnode = yyS[yypt-0].tnode; }
	case 68:
		//line mutan.y:227
		{
				node := yyS[yypt-2].tnode
		      		yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 69:
		//line mutan.y:232
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 70:
		//line mutan.y:238
		{
				yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
			}
	case 71:
		//line mutan.y:242
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].tnode.Constant
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
			}
	case 72:
		//line mutan.y:248
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-3].str
				varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
			}
	case 73:
		//line mutan.y:254
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 74:
		//line mutan.y:255
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 75:
		//line mutan.y:256
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 76:
		//line mutan.y:261
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 77:
		//line mutan.y:266
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
				yyVAL.tnode.Ptr = true
			}
	case 78:
		//line mutan.y:275
		{
				yyVAL.tnode = NewNode(NewArrayTy)
				yyVAL.tnode.Size = yyS[yypt-2].str
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 79:
		//line mutan.y:283
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 80:
		//line mutan.y:287
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 81:
		//line mutan.y:288
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 82:
		//line mutan.y:289
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 83:
		//line mutan.y:290
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 84:
		//line mutan.y:291
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 85:
		//line mutan.y:292
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 86:
		//line mutan.y:293
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 87:
		//line mutan.y:297
		{ yyVAL.tnode = NewNode(DerefPtrTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 88:
		//line mutan.y:301
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 89:
		//line mutan.y:302
		{ yyVAL.tnode = NewNode(NilTy) }
	case 90:
		//line mutan.y:306
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 91:
		//line mutan.y:310
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	}
	goto yystack /* stack new state and value */
}


//line front/mutan.y:2
package frontend
import __yyfmt__ "fmt"
//line front/mutan.y:2
		
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


//line front/mutan.y:44
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

//line front/mutan.y:307



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 91
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 749

var yyAct = []int{

	33, 27, 4, 170, 2, 138, 50, 20, 101, 70,
	194, 83, 51, 52, 74, 73, 49, 85, 35, 14,
	195, 105, 59, 22, 56, 207, 56, 172, 113, 104,
	72, 71, 56, 46, 193, 167, 54, 176, 58, 164,
	106, 94, 36, 224, 67, 217, 171, 201, 65, 92,
	50, 200, 68, 146, 145, 61, 51, 52, 56, 56,
	56, 56, 69, 62, 93, 98, 64, 144, 56, 56,
	87, 88, 89, 90, 165, 103, 226, 110, 110, 110,
	95, 96, 97, 114, 56, 56, 56, 82, 23, 218,
	3, 63, 202, 190, 98, 189, 136, 135, 50, 137,
	108, 111, 112, 141, 51, 52, 188, 56, 57, 187,
	186, 185, 184, 66, 34, 142, 37, 183, 182, 140,
	180, 179, 178, 177, 148, 44, 147, 38, 115, 41,
	36, 107, 17, 40, 39, 35, 102, 91, 162, 161,
	64, 16, 42, 43, 160, 110, 110, 110, 175, 159,
	25, 158, 56, 56, 21, 29, 181, 157, 156, 155,
	30, 154, 152, 28, 26, 151, 191, 150, 173, 174,
	198, 149, 84, 81, 80, 79, 78, 56, 77, 76,
	75, 204, 203, 143, 139, 132, 163, 153, 91, 199,
	45, 134, 209, 31, 208, 205, 47, 197, 206, 212,
	169, 110, 110, 211, 48, 99, 60, 53, 56, 192,
	213, 214, 24, 215, 216, 219, 100, 133, 110, 223,
	196, 7, 15, 116, 210, 110, 225, 18, 32, 19,
	11, 222, 6, 13, 1, 9, 34, 0, 37, 0,
	0, 0, 0, 0, 0, 0, 0, 44, 0, 38,
	0, 41, 36, 12, 17, 40, 39, 35, 0, 0,
	5, 0, 0, 16, 42, 43, 0, 0, 0, 0,
	0, 0, 25, 8, 0, 10, 21, 29, 0, 0,
	0, 0, 30, 0, 0, 28, 26, 18, 0, 19,
	11, 221, 0, 0, 0, 9, 34, 0, 37, 0,
	0, 0, 0, 0, 0, 0, 0, 44, 0, 38,
	0, 41, 36, 12, 17, 40, 39, 35, 0, 0,
	5, 0, 0, 16, 42, 43, 0, 0, 0, 0,
	0, 0, 25, 8, 0, 10, 21, 29, 0, 0,
	0, 0, 30, 0, 0, 28, 26, 18, 0, 19,
	11, 220, 0, 0, 0, 9, 34, 0, 37, 0,
	0, 0, 0, 0, 0, 0, 0, 44, 0, 38,
	0, 41, 36, 12, 17, 40, 39, 35, 0, 0,
	5, 0, 0, 16, 42, 43, 0, 0, 0, 0,
	0, 0, 25, 8, 0, 10, 21, 29, 0, 0,
	0, 0, 30, 0, 0, 28, 26, 18, 0, 19,
	11, 168, 0, 0, 0, 9, 34, 0, 37, 0,
	0, 0, 0, 0, 0, 0, 0, 44, 0, 38,
	0, 41, 36, 12, 17, 40, 39, 35, 0, 0,
	5, 0, 0, 16, 42, 43, 0, 0, 0, 0,
	0, 0, 25, 8, 0, 10, 21, 29, 0, 0,
	0, 0, 30, 0, 0, 28, 26, 18, 0, 19,
	11, 166, 0, 0, 0, 9, 34, 0, 37, 0,
	0, 0, 0, 0, 0, 0, 0, 44, 0, 38,
	0, 41, 36, 12, 17, 40, 39, 35, 0, 0,
	5, 0, 0, 16, 42, 43, 0, 0, 0, 0,
	0, 0, 25, 8, 0, 10, 21, 29, 0, 0,
	0, 0, 30, 0, 0, 28, 26, 18, 0, 19,
	11, 86, 0, 0, 0, 9, 34, 0, 37, 0,
	0, 0, 0, 0, 0, 0, 0, 44, 0, 38,
	0, 41, 36, 12, 17, 40, 39, 35, 0, 0,
	5, 0, 0, 16, 42, 43, 0, 0, 0, 0,
	0, 0, 25, 8, 0, 10, 21, 29, 0, 0,
	0, 18, 30, 19, 11, 28, 26, 0, 0, 9,
	34, 0, 37, 0, 0, 0, 0, 0, 0, 0,
	0, 44, 0, 38, 0, 41, 36, 12, 17, 40,
	39, 35, 0, 0, 5, 0, 0, 16, 42, 43,
	0, 0, 0, 0, 0, 0, 25, 8, 0, 10,
	21, 29, 34, 0, 37, 0, 30, 0, 0, 28,
	26, 0, 0, 44, 0, 38, 0, 41, 36, 0,
	17, 40, 39, 35, 0, 0, 0, 0, 0, 16,
	42, 43, 0, 0, 34, 0, 37, 0, 25, 0,
	0, 0, 21, 29, 0, 44, 0, 38, 30, 41,
	36, 28, 26, 40, 39, 35, 0, 0, 0, 37,
	0, 0, 42, 43, 0, 0, 0, 0, 44, 0,
	38, 0, 41, 36, 55, 29, 40, 39, 35, 131,
	30, 0, 0, 28, 26, 42, 43, 117, 119, 120,
	121, 122, 126, 0, 0, 0, 0, 109, 29, 0,
	0, 0, 0, 30, 0, 129, 28, 26, 0, 118,
	0, 0, 0, 123, 124, 125, 127, 128, 130,
}
var yyPact = []int{

	-1000, -1000, 575, -1000, -1000, 178, -1000, -1000, -22, 187,
	11, -1000, -1000, -1000, -52, 203, 649, 575, 649, 617,
	202, 51, -1000, -1000, 649, -3, -24, -1000, -25, -1000,
	-1000, -1000, -1000, -44, -1000, -1000, -46, 165, 164, 163,
	161, 160, 159, 158, 62, -51, 157, -40, -1000, 521,
	649, 649, 649, 649, -52, 125, -1000, -1000, 40, 32,
	649, 649, 617, 201, -1000, 120, 120, -52, -1000, -26,
	-35, -1000, -1000, -1000, 9, 115, 672, 672, 672, -27,
	617, 112, 698, 172, -1000, 181, -1000, -52, -52, -52,
	-52, 617, -1000, 617, -1000, -52, -52, -8, 171, 649,
	99, -1000, -1000, -1000, -1000, 170, -1000, -1000, 38, 176,
	-1000, 25, 24, 110, 108, -1000, -1000, 156, 152, 150,
	147, 175, 146, 144, 143, 142, 136, 134, 129, 124,
	123, 174, -1000, 23, -1000, 461, 3, 401, 196, -1000,
	-52, 17, -1000, -28, 672, 672, -18, -1000, -1000, 107,
	106, 105, 104, 617, 102, 101, 96, 95, 94, 93,
	90, 79, 77, 617, -17, -45, 190, 617, -1000, 649,
	-1000, -1000, -1000, 22, 18, 76, -1000, -1000, -1000, -1000,
	-1000, 169, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 168, 186, -1000, 17, -30, -1000, 185, 183, -52,
	672, -18, -1000, -1000, 195, -1000, -1000, 17, -1000, -1000,
	16, 73, 617, 341, -1000, 281, 221, -18, -1000, -1000,
	-1000, -1000, -1000, 14, -18, 60, -1000,
}
var yyPgo = []int{

	0, 234, 4, 90, 2, 233, 19, 23, 232, 88,
	228, 223, 222, 222, 3, 0, 193, 221, 220, 1,
	217, 216, 7, 212, 8, 209,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 20, 20, 20, 25, 25, 10, 10,
	10, 10, 10, 10, 10, 10, 13, 13, 14, 14,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 8, 18, 18, 17,
	17, 4, 4, 4, 4, 4, 4, 21, 21, 5,
	5, 5, 5, 5, 12, 12, 12, 6, 6, 6,
	6, 6, 9, 9, 9, 9, 7, 7, 7, 7,
	7, 7, 7, 7, 22, 19, 19, 15, 16, 23,
	24,
}
var yyR2 = []int{

	0, 1, 2, 0, 1, 4, 1, 1, 9, 4,
	2, 3, 1, 4, 5, 0, 1, 0, 3, 12,
	8, 6, 4, 4, 3, 3, 3, 0, 1, 0,
	3, 3, 3, 3, 4, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 6, 6, 4, 0, 9,
	5, 1, 1, 1, 2, 2, 0, 3, 0, 3,
	3, 6, 3, 4, 2, 3, 5, 1, 1, 3,
	3, 4, 2, 3, 3, 3, 1, 1, 2, 1,
	4, 1, 1, 1, 2, 1, 1, 1, 3, 1,
	1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 39, -8, -17, 52, 14,
	54, 9, 32, -5, -6, -12, 42, 33, 6, 8,
	-22, 55, -7, -9, -23, 51, 65, -19, 64, 56,
	61, -16, -10, -15, 15, 36, 31, 17, 28, 35,
	34, 30, 43, 44, 26, 12, 55, 9, -16, -2,
	58, 64, 65, 4, -6, 55, -22, -3, -6, -4,
	4, 4, 12, 40, 15, -7, -9, -6, 55, 65,
	12, 55, 55, 59, 60, 15, 15, 15, 15, 15,
	15, 15, 25, 62, 15, 57, 10, -6, -6, -6,
	-6, 12, 9, 32, 9, -6, -6, -6, -4, 4,
	-21, -24, 16, -24, 55, 56, 31, 16, -7, 55,
	-15, -7, -7, 55, -4, 16, -11, 19, 41, 20,
	21, 22, 23, 45, 46, 47, 24, 48, 49, 37,
	50, 11, 13, -20, 10, -2, -4, -2, 13, 13,
	-6, -4, 16, 13, 29, 29, 29, 16, 16, 15,
	15, 15, 15, 12, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 12, 16, 51, 10, 32, 10, 4,
	-14, 29, 55, -7, -7, -19, 55, 16, 16, 16,
	16, -4, 16, 16, 16, 16, 16, 16, 16, 16,
	16, -4, -25, 51, 55, 65, -18, 7, -4, -6,
	29, 29, 16, 13, 13, 9, -14, 55, 9, 9,
	-7, -19, 4, -2, -14, -2, -2, 29, 16, -4,
	10, 10, 10, -19, 29, -19, 16,
}
var yyDef = []int{

	3, -2, 1, 2, 4, 0, 6, 7, 0, 0,
	0, 3, 12, 51, 52, 53, 0, 56, 0, 56,
	77, 87, 67, 68, 0, 0, 0, 76, 0, 79,
	81, 82, 83, 85, 89, 86, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 10, 0,
	0, 0, 0, 0, 54, 87, 77, 55, 0, 0,
	0, 0, 56, 0, 58, 67, 68, 0, 64, 0,
	0, 84, 78, 72, 0, 0, 0, 0, 0, 0,
	56, 0, 0, 0, 15, 0, 11, 73, 74, 75,
	62, 56, 3, 56, 3, 59, 60, 0, 0, 0,
	56, 69, 90, 70, 65, 0, 88, 18, 0, 87,
	85, 0, 0, 0, 0, 24, 25, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 5, 0, 9, 0, 0, 0, 0, 80,
	63, 29, 71, 0, 0, 0, 0, 22, 23, 0,
	0, 0, 0, 56, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 56, 17, 0, 48, 56, 50, 0,
	57, 28, 66, 0, 0, 0, 87, 30, 31, 32,
	33, 0, 35, 36, 37, 38, 39, 40, 41, 42,
	43, 0, 0, 16, 29, 0, 46, 0, 0, 61,
	0, 0, 21, 34, 44, 3, 13, 29, 3, 3,
	0, 0, 56, 0, 14, 0, 0, 0, 20, 45,
	8, 47, 49, 0, 0, 0, 19,
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
		//line front/mutan.y:65
		{ Tree = yyS[yypt-0].tnode }
	case 2:
		//line front/mutan.y:69
		{ yyVAL.tnode = NewNode(StatementListTy, yyS[yypt-1].tnode, yyS[yypt-0].tnode) }
	case 3:
		//line front/mutan.y:70
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 4:
		//line front/mutan.y:75
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 5:
		//line front/mutan.y:76
		{ yyVAL.tnode = NewNode(LambdaTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 6:
		//line front/mutan.y:77
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 7:
		//line front/mutan.y:78
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 8:
		//line front/mutan.y:80
		{
				yyVAL.tnode = NewNode(FuncDefTy, yyS[yypt-1].tnode);
				yyVAL.tnode.Constant = yyS[yypt-7].str
				yyVAL.tnode.HasRet = yyS[yypt-3].check
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-5].tnode, true)
			}
	case 9:
		//line front/mutan.y:86
		{ yyVAL.tnode = NewNode(InlineAsmTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 10:
		//line front/mutan.y:87
		{ yyVAL.tnode = NewNode(ImportTy); yyVAL.tnode.Constant = yyS[yypt-0].tnode.Constant }
	case 11:
		//line front/mutan.y:88
		{ yyVAL.tnode = NewNode(ScopeTy, yyS[yypt-1].tnode) }
	case 12:
		//line front/mutan.y:89
		{ yyVAL.tnode = NewNode(EmptyTy); }
	case 13:
		//line front/mutan.y:93
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-3].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 14:
		//line front/mutan.y:94
		{ yyVAL.tnode = NewNode(NewVarTy, yyS[yypt-4].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str; yyVAL.tnode.Ptr = true }
	case 15:
		//line front/mutan.y:95
		{ yyVAL.tnode = nil }
	case 16:
		//line front/mutan.y:100
		{ yyVAL.check = true }
	case 17:
		//line front/mutan.y:101
		{ yyVAL.check = false }
	case 18:
		//line front/mutan.y:106
		{ yyVAL.tnode = NewNode(StopTy) }
	case 19:
		//line front/mutan.y:109
		{
			  yyVAL.tnode = NewNode(CallTy, yyS[yypt-9].tnode, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 20:
		//line front/mutan.y:113
		{
			  yyVAL.tnode = NewNode(TransactTy, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
		  }
	case 21:
		//line front/mutan.y:116
		{ yyVAL.tnode = NewNode(CreateTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode) }
	case 22:
		//line front/mutan.y:117
		{ yyVAL.tnode = NewNode(SizeofTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 23:
		//line front/mutan.y:118
		{ yyVAL.tnode = NewNode(PushTy, yyS[yypt-1].tnode) }
	case 24:
		//line front/mutan.y:119
		{ yyVAL.tnode = NewNode(PopTy) }
	case 25:
		//line front/mutan.y:120
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 26:
		//line front/mutan.y:124
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode) }
	case 27:
		//line front/mutan.y:125
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 28:
		//line front/mutan.y:129
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 29:
		//line front/mutan.y:130
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 30:
		//line front/mutan.y:135
		{ yyVAL.tnode = NewNode(OriginTy) }
	case 31:
		//line front/mutan.y:136
		{ yyVAL.tnode = NewNode(AddressTy) }
	case 32:
		//line front/mutan.y:137
		{ yyVAL.tnode = NewNode(CallerTy) }
	case 33:
		//line front/mutan.y:138
		{ yyVAL.tnode = NewNode(CallValTy) }
	case 34:
		//line front/mutan.y:139
		{ yyVAL.tnode = NewNode(CallDataLoadTy, yyS[yypt-1].tnode) }
	case 35:
		//line front/mutan.y:140
		{ yyVAL.tnode = NewNode(CallDataSizeTy) }
	case 36:
		//line front/mutan.y:141
		{ yyVAL.tnode = NewNode(DiffTy) }
	case 37:
		//line front/mutan.y:142
		{ yyVAL.tnode = NewNode(PrevHashTy) }
	case 38:
		//line front/mutan.y:143
		{ yyVAL.tnode = NewNode(TimestampTy) }
	case 39:
		//line front/mutan.y:144
		{ yyVAL.tnode = NewNode(GasPriceTy) }
	case 40:
		//line front/mutan.y:145
		{ yyVAL.tnode = NewNode(BlockNumTy) }
	case 41:
		//line front/mutan.y:146
		{ yyVAL.tnode = NewNode(CoinbaseTy) }
	case 42:
		//line front/mutan.y:147
		{ yyVAL.tnode = NewNode(BalanceTy) }
	case 43:
		//line front/mutan.y:148
		{ yyVAL.tnode = NewNode(GasTy) }
	case 44:
		//line front/mutan.y:149
		{ yyVAL.tnode = NewNode(StoreTy, yyS[yypt-1].tnode) }
	case 45:
		//line front/mutan.y:151
		{
				node := NewNode(SetStoreTy, yyS[yypt-3].tnode)
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 46:
		//line front/mutan.y:159
		{
				if yyS[yypt-0].tnode == nil {
					yyVAL.tnode = NewNode(IfThenTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode)
				} else {
					yyVAL.tnode = NewNode(IfThenElseTy, yyS[yypt-4].tnode, yyS[yypt-2].tnode, yyS[yypt-0].tnode)
				}
			}
	case 47:
		//line front/mutan.y:169
		{
				yyVAL.tnode = yyS[yypt-1].tnode
			}
	case 48:
		//line front/mutan.y:172
		{ yyVAL.tnode = nil }
	case 49:
		//line front/mutan.y:177
		{
				yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-7].tnode, yyS[yypt-5].tnode, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
			}
	case 50:
		//line front/mutan.y:181
		{
				yyVAL.tnode = NewNode(ForThenTy, yyS[yypt-3].tnode, yyS[yypt-1].tnode)
			}
	case 51:
		//line front/mutan.y:187
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 52:
		//line front/mutan.y:188
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 53:
		//line front/mutan.y:189
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 54:
		//line front/mutan.y:190
		{ yyVAL.tnode = NewNode(ReturnTy, yyS[yypt-0].tnode) }
	case 55:
		//line front/mutan.y:191
		{ yyVAL.tnode = NewNode(ExitTy, yyS[yypt-0].tnode) }
	case 56:
		//line front/mutan.y:192
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 57:
		//line front/mutan.y:196
		{ yyVAL.tnode = NewNode(ArgTy, yyS[yypt-2].tnode, yyS[yypt-1].tnode);}
	case 58:
		//line front/mutan.y:197
		{ yyVAL.tnode = nil }
	case 59:
		//line front/mutan.y:202
		{
		      		yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode)
			}
	case 60:
		//line front/mutan.y:206
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, node)
			}
	case 61:
		//line front/mutan.y:212
		{
				yyVAL.tnode = NewNode(AssignArrayTy, yyS[yypt-3].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-5].str
			}
	case 62:
		//line front/mutan.y:216
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-2].tnode.Constant
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, yyS[yypt-2].tnode, node)
			}
	case 63:
		//line front/mutan.y:222
		{
				node := NewNode(SetLocalTy)
				node.Constant = yyS[yypt-3].str
				varNode := NewNode(NewVarTy); varNode.Constant = yyS[yypt-3].str
				yyVAL.tnode = NewNode(AssignmentTy, yyS[yypt-0].tnode, varNode, node)
			}
	case 64:
		//line front/mutan.y:232
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 65:
		//line front/mutan.y:237
		{
				yyVAL.tnode = NewNode(NewVarTy)
				yyVAL.tnode.Constant = yyS[yypt-0].str
				yyVAL.tnode.Ptr = true
			}
	case 66:
		//line front/mutan.y:243
		{
				yyVAL.tnode = NewNode(NewArrayTy)
				yyVAL.tnode.Size = yyS[yypt-2].str
				yyVAL.tnode.Constant = yyS[yypt-0].str
			}
	case 67:
		//line front/mutan.y:251
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 68:
		//line front/mutan.y:252
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 69:
		//line front/mutan.y:253
		{ yyVAL.tnode = yyS[yypt-1].tnode }
	case 70:
		//line front/mutan.y:254
		{ yyVAL.tnode = yyS[yypt-1].tnode }
	case 71:
		//line front/mutan.y:256
		{
				yyVAL.tnode = NewNode(FuncCallTy, yyS[yypt-1].tnode)
				yyVAL.tnode.Constant = yyS[yypt-3].str
				yyVAL.tnode.ArgList = makeArgs(yyS[yypt-1].tnode, false)
			}
	case 72:
		//line front/mutan.y:265
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 73:
		//line front/mutan.y:267
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 74:
		//line front/mutan.y:268
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 75:
		//line front/mutan.y:269
		{ yyVAL.tnode = NewNode(OpTy, yyS[yypt-2].tnode, yyS[yypt-0].tnode); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 76:
		//line front/mutan.y:273
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 77:
		//line front/mutan.y:274
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 78:
		//line front/mutan.y:275
		{ yyVAL.tnode = NewNode(RefTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 79:
		//line front/mutan.y:276
		{ yyVAL.tnode = NewNode(ConstantTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 80:
		//line front/mutan.y:277
		{ yyVAL.tnode = NewNode(ArrayTy, yyS[yypt-1].tnode); yyVAL.tnode.Constant = yyS[yypt-3].str }
	case 81:
		//line front/mutan.y:278
		{ yyVAL.tnode = NewNode(BoolTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 82:
		//line front/mutan.y:279
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 83:
		//line front/mutan.y:280
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 84:
		//line front/mutan.y:284
		{ yyVAL.tnode = NewNode(DerefPtrTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 85:
		//line front/mutan.y:288
		{ yyVAL.tnode = yyS[yypt-0].tnode }
	case 86:
		//line front/mutan.y:289
		{ yyVAL.tnode = NewNode(NilTy) }
	case 87:
		//line front/mutan.y:293
		{ yyVAL.tnode = NewNode(IdentifierTy); yyVAL.tnode.Constant = yyS[yypt-0].str }
	case 88:
		//line front/mutan.y:297
		{ yyVAL.tnode = NewNode(StringTy); yyVAL.tnode.Constant = yyS[yypt-1].str }
	case 89:
		//line front/mutan.y:301
		{ yyVAL.tnode = NewNode(EmptyTy) }
	case 90:
		//line front/mutan.y:305
		{ yyVAL.tnode = NewNode(EmptyTy) }
	}
	goto yystack /* stack new state and value */
}

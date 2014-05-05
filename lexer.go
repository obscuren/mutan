package mutan

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// This lexer is heavily inspired on Rob Pike's lexer

func isAlphaNumeric(t rune) bool {
	return unicode.IsLetter(t)
}

func isSpace(t rune) bool {
	return unicode.IsSpace(t)
}

func isNumber(t rune) bool {
	return unicode.IsNumber(t)
}

func isOperator(t rune) bool {
	return strings.IndexRune("=+-/*><^", t) >= 0
}

const eof = -1

type stateFn func(*Lexer) stateFn
type itemType int

const (
	itemEof          itemType = 0
	itemIdentifier            = ID
	itemEndStatement          = END_STMT
	itemNumber                = NUMBER
	itemAssign                = ASSIGN
	itemOp                    = OP
	itemDop                   = DOP /* Double op ++ -- */
	itemLeftBracket           = LEFT_BRACKET
	itemRightBracket          = RIGHT_BRACKET
	itemLeftBrace             = LEFT_BRACES
	itemRightBrace            = RIGHT_BRACES
	itemEqual                 = EQUAL
	itemIf                    = IF
	itemElse                  = ELSE
	itemFor                   = FOR
	itemStore                 = STORE
	itemReturn                = RETURN
	itemAsm                   = ASM
	itemInlineAsm             = INLINE_ASM
	itemLeftPar               = LEFT_PAR
	itemRightPar              = RIGHT_PAR
	itemStop                  = STOP
	itemAddr                  = ADDR
	itemOrigin                = ORIGIN
	itemCaller                = CALLER
	itemCallVal               = CALLVAL
	itemCallDataLoad          = CALLDATALOAD
	itemCallDataSize          = CALLDATASIZE
	itemDifficulty            = DIFFICULTY
	itemPrevHash              = PREVHASH
	itemTimeStamp             = TIMESTAMP
	itemGasPrice              = GASPRICE
	itemBlockNum              = BLOCKNUM
	itemCoinbase              = COINBASE
	itemGas                   = GAS
	itemDot                   = DOT
	itemThis                  = THIS
	itemArray                 = ARRAY
	itemVarType               = TYPE
	itemComma                 = COMMA
	itemCall                  = CALL
	itemTransact              = TRANSACT
	itemCreate                = CREATE
	itemSizeof                = SIZEOF
	itemQuote                 = QUOTE
	itemStr                   = STR
	itemNil                   = NIL
)

type item struct {
	typ itemType
	val string
}

func lexStatement(l *Lexer) stateFn {
	acceptance := Alpha
	l.acceptRun(acceptance)

	if l.accept("1234567890") {
		acceptance += "1234567890"
	}
	l.acceptRun(acceptance)

	switch l.blob() {
	case "nil":
		l.emit(itemNil)
	case "if":
		l.emit(itemIf)
	case "else":
		l.emit(itemElse)
	case "for":
		l.emit(itemFor)
	case "asm":
		l.emit(itemAsm)

		return lexInsideAsm
	case "exit":
		l.emit(itemStop)
	case "this":
		l.emit(itemThis)

		return lexClosureScope
	case "array":
		l.emit(itemArray)

		return lexArray
	case "int8", "int16", "int32", "int64", "int256", "big", "string", "addr":
		l.emit(itemVarType)
	case "call":
		l.emit(itemCall)
	case "create":
		l.emit(itemCreate)
	case "transact":
		l.emit(itemTransact)
	case "return":
		l.emit(itemReturn)
	case "sizeof":
		l.emit(itemSizeof)
	default:
		l.emit(itemIdentifier)
	}

	return lexText
}

const Numbers = "1234567890"
const Alpha = "abcdefghijklmnopqrstuwvxyzABCDEFGHIJKLMNOPQRSTUWVXYZ"

func lexCall(l *Lexer) stateFn {
	if !l.accept("(") {
		l.err = fmt.Errorf("Expected '('")
		return nil
	}
	l.emit(itemLeftPar)

	return lexText
}

func lexArray(l *Lexer) stateFn {
	if !l.accept("(") {
		l.err = fmt.Errorf("Exepcted '('")
		return nil
	}

	l.emit(itemLeftPar)

	if !l.accept(Numbers) {
		l.err = fmt.Errorf("Expected number")
		return nil
	}

	l.acceptRun(Numbers)

	l.emit(itemNumber)

	if !l.accept(")") {
		l.err = fmt.Errorf("Expected ')'")
		return nil
	}

	l.emit(itemRightPar)

	return lexText
}

func lexClosureScope(l *Lexer) stateFn {
	if !l.accept(".") {
		l.err = fmt.Errorf("Expected '.'")

		return nil
	}

	l.emit(itemDot)

	acceptance := "abcdefghijklmnopqrstuwvxyzABCDEFGHIJKLMNOPQRSTUWVXYZ"
	l.acceptRun(acceptance)
	switch l.blob() {
	case "caller":
		l.emit(itemCaller)
	case "origin":
		l.emit(itemOrigin)
	case "value":
		l.emit(itemCallVal)
	case "dataLoad":
		l.emit(itemCallDataLoad)
	case "data":
		l.emit(itemCallDataLoad)
	case "dataSize":
		l.emit(itemCallDataSize)
	case "gasPrice":
		l.emit(itemGasPrice)
	case "diff":
		l.emit(itemDifficulty)
	case "prevHash":
		l.emit(itemPrevHash)
	case "time":
		l.emit(itemTimeStamp)
	case "number":
		l.emit(itemBlockNum)
	case "coinbase":
		l.emit(itemCoinbase)
	case "gas":
		l.emit(itemGas)
	case "store":
		l.emit(itemStore)
	default:
		l.err = fmt.Errorf("Undefined '%s'", l.blob())

		return nil
	}

	return lexText
}

func lexNumber(l *Lexer) stateFn {
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}

	l.acceptRun(digits)

	l.emit(itemNumber)

	return lexText
}

func lexInsideAsm(l *Lexer) stateFn {
out:
	for {
		switch r := l.next(); {
		case isSpace(r):
			l.ignore()
		case r == '(':
			l.emit(itemLeftPar)
		case r == ')':
			l.backup()

			break out
		default:
			l.acceptRunUntill(')')

			l.emit(itemInlineAsm)
		}
	}

	return lexText
}

/* TODO
func lexInsidePar(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
		case isComma(r):
			l.ignore()
		case isAlphaNumeric(r):
			l.backup()

			return lexStatement
		}
	}
}
*/

func lexOperator(l *Lexer) stateFn {
	// The only special case there is, assignment

	acceptance := "="
	if !l.accept("=") {
		acceptance += "-/*+><^"
	}
	l.acceptRun(acceptance)

	switch l.blob() {
	case "=":
		l.emit(itemAssign)
	case "++", "--":
		l.emit(itemDop)
	default:
		l.emit(itemOp)
	}

	return lexText
}

func lexInsideString(l *Lexer) stateFn {
	l.acceptRunUntill('"')
	l.emit(itemStr)

	l.next()
	l.emit(itemQuote)

	return lexText
}

func lexComment(l *Lexer) stateFn {
	l.acceptRunUntill('\n')

	return lexText
}

var Lineno int

// Lex text attempts to identify the current state that *might*
// be and calls the appropriate lexing method. The lexing method
// should then take care of anything that is current (even validating)
func lexText(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == '\n':
			l.ignore()

			Lineno++
		case isSpace(r): // Check whether this is a space (which we ignore)
			l.ignore()
		case isAlphaNumeric(r): // Check if it's alpha numeric (var, if, else etc)
			l.backup()

			return lexStatement
		case isNumber(r): // Check if it's a number (constant)
			l.backup()

			return lexNumber
		case r == '{': // Block check
			l.emit(itemLeftBrace)
		case r == '}':
			l.emit(itemRightBrace)
		case r == '[':
			l.emit(itemLeftBracket)
		case r == ']':
			l.emit(itemRightBracket)
		case r == '(':
			l.emit(itemLeftPar)
		case r == ')':
			l.emit(itemRightPar)
		case r == '.':
			l.emit(itemDot)
		case r == ',':
			l.emit(itemComma)
		case r == '"':
			l.emit(itemQuote)

			return lexInsideString
		case r == '/', r == '#':
			return lexComment
		case r == ';':
			l.emit(itemEndStatement)
		case isOperator(r):
			l.backup()

			return lexOperator
		default:
			return nil
		}
	}

	return nil
}

type Lexer struct {
	name   string
	input  string
	start  int
	pos    int
	width  int
	state  stateFn
	items  chan item
	err    error
	lineno int
}

func lexer(name, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		state: lexText,
		items: make(chan item, 5),
	}

	return l
}

// Grabs the current blob of text
func (l *Lexer) blob() string {
	return l.input[l.start:l.pos]
}

// Emits a new item on to item channel for processing
func (l *Lexer) emit(t itemType) {
	l.items <- item{t, l.blob()}
	l.start = l.pos
}

// Accepts checks whether the given input matches the next rune
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}

	l.backup()

	return false
}

// Continues *eating* the next rune until no longer valid
func (l *Lexer) acceptRegexp(valid string) bool {
	if MatchRegexp(valid, []byte(string(l.next()))) {
		return true
	}
	l.backup()

	return false
}

func MatchRegexp(reg string, str []byte) bool {
	ok, _ := regexp.Match(reg, str)
	return ok
}

func (l *Lexer) acceptRunRegexp(valid string) {
	for r := l.next(); MatchRegexp(valid, []byte(string(r))); {
	}
	l.backup()
}

func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *Lexer) acceptRunUntill(until rune) {
	// Continues running until a rune is found
	for strings.IndexRune(string(until), l.next()) == -1 {
	}
	l.backup()
}

// Grabs the next item of the channel and returns it, or nil if we're done
func (l *Lexer) nextItem() item {
	for {
		select {
		case item := <-l.items:
			return item
		default:
			if l.state == nil {
				return item{}
			}

			l.state = l.state(l)
		}
	}
	panic("not reached")
}

// Takes the next rune and returns it or returns EOF
func (l *Lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0

		return eof
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width

	return rune
}

// Look ahead
func (l *Lexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

// Backup a previous *next*
func (l *Lexer) backup() {
	l.pos -= l.width
}

// Ignore the current rune
func (l *Lexer) ignore() {
	l.start = l.pos
}

// yacc's lexing method
func (l *Lexer) Lex(lval *yySymType) int {
	item := l.nextItem()
	lval.str = item.val

	return int(item.typ)
}

func (l *Lexer) Error(s string) {
	l.err = fmt.Errorf("%s", s)
}

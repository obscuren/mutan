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

const eof = -1

type stateFn func(*Lexer) stateFn
type itemType int

const (
	itemEof        itemType = 0
	itemIdentifier          = ID
	itemNumber              = NUMBER
	itemAssign              = ASSIGN
	itemLeftBracket
	itemRightBracket
	itemLeftBrace  = LEFT_BRACES
	itemRightBrace = RIGHT_BRACES
	itemEqual      = EQUAL
	itemIf         = IF
)

type item struct {
	typ itemType
	val string
}

func lexStatement(l *Lexer) stateFn {
	acceptance := "abcdefghijklmnopqrstuwvxyzABCDEFGHIJKLMNOPQRSTUWVXYZ"
	l.accept(acceptance)

	if l.accept("1234567890") {
		acceptance += "1234567890"
	}
	l.acceptRun(acceptance)

	switch l.blob() {
	case "if":
		l.emit(itemIf)
	default:
		l.emit(itemIdentifier)
	}

	return lexText
}

func lexNumber(l *Lexer) stateFn {
	digits := "1234567890"
	l.acceptRun(digits)

	l.emit(itemNumber)

	return lexText
}

// Lex text attempts to identify the current state that *might*
// be and calls the appropriate lexing method. The lexing method
// should then take care of anything that is current (even validating)
func lexText(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
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
		case r == '=': // TODO turn this in to an operator check
			if l.peek() == '=' {
				l.next()
				l.emit(itemEqual)
			} else {
				l.emit(itemAssign)
			}
		default:
			return nil
		}
	}

	return nil
}

type Lexer struct {
	name  string
	input string
	start int
	pos   int
	width int
	state stateFn
	items chan item
	err   bool
}

func lexer(name, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		state: lexText,
		items: make(chan item, 2),
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
	l.err = true

	fmt.Println(s)
}

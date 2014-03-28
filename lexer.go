package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func isAlphaNumeric(t rune) bool {
	return unicode.IsLetter(t)
}

func isSpace(t rune) bool {
	return unicode.IsSpace(t)
}

const eof = -1

type stateFn func(*Lexer) stateFn
type itemType int

const (
	itemEof        itemType = 0
	itemIdentifier          = ID
	itemLeftBracket
	itemRightBracket
)

type item struct {
	typ itemType
	val string
}

func lexIdentifier(l *Lexer) stateFn {
	l.acceptRun("abcdefghijklmnopqrstuwvxyzABCDEFGHIJKLMNOPQRSTUWVXYZ")
	l.emit(itemIdentifier)

	return lexIdentifyState
}

func lexIdentifyState(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
		case isSpace(r):
			l.ignore()
		case isAlphaNumeric(r):
			l.backup()

			return lexIdentifier
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
}

func (l *Lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}

	l.backup()

	return false
}

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

func (l *Lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0

		return eof
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width

	return rune
}

func (l *Lexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

func (l *Lexer) backup() {
	l.pos -= l.width
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

func (l *Lexer) Lex(lval *yySymType) int {
	item := l.nextItem()
	lval.str = item.val

	return int(item.typ)
}

func lexer(name, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		state: lexIdentifyState,
		items: make(chan item, 2),
	}

	return l
}

func (l *Lexer) Error(s string) {
	fmt.Printf(s)
}

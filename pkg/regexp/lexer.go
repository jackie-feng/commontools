package regexp

import (
	"log"
	"unicode/utf8"
)

const eof = 0

type Lexer struct {
	line []byte
	peek rune
	ast  Node
}

var _ YoYoLexer = (*Lexer)(nil)

func (l *Lexer) Lex(lval *YoYoSymType) int {
	for {
		c := l.next()
		s := string(c)
		switch c {
		case eof:
			return eof
		case '*', '.', '(', ')', '|', '\\':
			return int(c)
		case ' ', '\t', '\n', '\r':
			log.Printf("unrecognized character %q %s", c, s)
		default:
			lval.char = RUNE(c)
			return CHAR
		}
	}
}

func (l Lexer) Error(s string) {
	log.Printf("parse error: %s", s)
}

// Return the next rune for the lexer.
func (x *Lexer) next() rune {
	if len(x.line) == 0 {
		return eof
	}
	c, size := utf8.DecodeRune(x.line)
	x.line = x.line[size:]
	if c == utf8.RuneError {
		return YoYoErrCode
	}
	return c
}

func NewLexer(line []byte) *Lexer {
	return &Lexer{line: line}
}

func SetResult(yylex interface{}, res interface{}) {
	yylex.(*Lexer).ast = res
}

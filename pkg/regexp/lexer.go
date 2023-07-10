package regexp

import (
	"bufio"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func reg() {
	in := bufio.NewReader(os.Stdin)
	exprParse := YoYoNewParser()
	for {
		if _, err := os.Stdout.WriteString("> "); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}

		exprLex := NewLexer(line)
		exprParse.Parse(exprLex)
	}
}

const eof = 0

type lexer struct {
	line []byte
	peek rune
	res  interface{}
}

func (l *lexer) Lex(lval *YoYoSymType) int {
	for {
		c := l.next()
		switch c {
		case YoYoEofCode:
			return 0
		case '*', '.', '(', ')', '|', '\\':
			return int(c)
		case ' ', '\t', '\n', '\r':
			log.Printf("unrecognized character %q", c)
		default:
			lval.char = c
			return CHAR
		}
	}
}

var Result interface{}

func SetResult(res interface{}) {
	Result = res
}

func (l lexer) Error(s string) {
	log.Printf("parse error: %s", s)
}

// Return the next rune for the lexer.
func (x *lexer) next() rune {
	if len(x.line) == 0 {
		return YoYoEofCode
	}
	c, size := utf8.DecodeRune(x.line)
	x.line = x.line[size:]
	if c == utf8.RuneError {
		return YoYoErrCode
	}
	return c
}

func NewLexer(line []byte) YoYoLexer {
	return &lexer{line: line}
}

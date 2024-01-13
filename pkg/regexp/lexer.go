package regexp

import (
	"log"
	"unicode/utf8"
)

const eof = 0

var _ YoYoLexer = (*Lexer)(nil)

type Lexer struct {
	line []byte
	peek rune
	ast  Node
}

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

func (l *Lexer) NFA() *NFA {
	nfa := NewNFA()
	start := nfa.newState()
	end := walkNFA(l.ast, start, nfa)
	nfa.startState = start
	nfa.endState = end
	return nfa
}

const Epsilon = RUNE(0)
const AnyRUNE = RUNE(-1)

func walkNFA(node Node, start NFAState, nfa *NFA) NFAState {
	switch node.(type) {
	case RUNE:
		end := nfa.newState()
		nfa.add(start, end, node.(RUNE))
		return end
	case AnyChar:
		end := nfa.newState()
		nfa.add(start, end, AnyRUNE)
		return end
	case RepeatedNode:
		subStart := nfa.newState()
		subEnd := walkNFA(node.(RepeatedNode).Val, subStart, nfa)
		end := nfa.newState()
		nfa.add(start, subStart, Epsilon)
		nfa.add(subEnd, subStart, Epsilon)
		nfa.add(subEnd, end, Epsilon)
		nfa.add(start, end, Epsilon)
		return end
	case AndNode:
		next := walkNFA(node.(AndNode).Left, start, nfa)
		end := walkNFA(node.(AndNode).Right, next, nfa)
		return end
	case OrNode:
		subStart1 := nfa.newState()
		subStart2 := nfa.newState()
		subEnd1 := walkNFA(node.(OrNode).Left, subStart1, nfa)
		subEnd2 := walkNFA(node.(OrNode).Right, subStart2, nfa)
		end := nfa.newState()
		nfa.add(start, subStart1, Epsilon)
		nfa.add(start, subStart2, Epsilon)
		nfa.add(subEnd1, end, Epsilon)
		nfa.add(subEnd2, end, Epsilon)
		return end
	}
	return 0
}

package regexp

import (
	"fmt"
	"io"
)

type Node interface{}

type RUNE rune

func (n RUNE) String() string {
	if n == AnyRUNE {
		return "Any"
	}
	if n == Epsilon {
		return "Epsilon"
	}

	return fmt.Sprintf("Rune<%c>", rune(n))
}

type AnyChar int

func (n AnyChar) String() string {
	return fmt.Sprintf("AnyChar")
}

func (n AnyChar) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if w, ok := s.Width(); ok {
			fmt.Fprintf(s, "%*sAnyChar", w, "")
			return
		}
		fmt.Fprintf(s, n.String())
	case 's':
		io.WriteString(s, n.String())
	case 'q':
		fmt.Fprintf(s, n.String())
	}
}

func (n RUNE) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if w, ok := s.Width(); ok {
			fmt.Fprintf(s, "%*sRune<%c>", w, "", rune(n))
			return
		}
		fmt.Fprintf(s, n.String())
	case 's':
		io.WriteString(s, n.String())
	case 'q':
		fmt.Fprintf(s, n.String())
	}
}

type AndNode struct {
	Left  Node
	Right Node
}

func (n AndNode) String() string {
	return fmt.Sprintf("And<%v, %v>", n.Left, n.Right)
}

func (n AndNode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if w, ok := s.Width(); ok {
			fmt.Fprintf(s, "%*sAnd<\n%*v,\n%*v\n%*s>", w, "", w+2, n.Left, w+2, n.Right, w, "")
			return
		}
		fmt.Fprintf(s, n.String())
	case 's':
		io.WriteString(s, n.String())
	case 'q':
		fmt.Fprintf(s, n.String())
	}
}

type OrNode struct {
	Left  Node
	Right Node
}

func (n OrNode) String() string {
	return fmt.Sprintf("Or<\n\t%v, \n\t%v>", n.Left, n.Right)
}

func (n OrNode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if w, ok := s.Width(); ok {
			fmt.Fprintf(s, "%*sOr<\n%*v,\n%*v\n%*s>", w, "", w+2, n.Left, w+2, n.Right, w, "")
			return
		}
		fmt.Fprintf(s, n.String())
	case 's':
		io.WriteString(s, n.String())
	case 'q':
		fmt.Fprintf(s, n.String())
	}
}

type RepeatedNode struct {
	Val Node
}

func (n RepeatedNode) String() string {
	return fmt.Sprintf("Repeated<%v>", n.Val)
}

func (n RepeatedNode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if w, ok := s.Width(); ok {
			fmt.Fprintf(s, "%*sRepeated<\n%*v\n%*s>", w, "", w+2, n.Val, w, "")
			return
		}
		fmt.Fprintf(s, n.String())
	case 's':
		io.WriteString(s, n.String())
	case 'q':
		fmt.Fprintf(s, n.String())
	}
}

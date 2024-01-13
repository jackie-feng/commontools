package regexp

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type NFAState int

type NFA struct {
	Debug bool
	paths map[NFAState]map[RUNE][]NFAState
	state NFAState

	startState NFAState
	endState   NFAState
}

func NewNFA() *NFA {
	return &NFA{
		paths: make(map[NFAState]map[RUNE][]NFAState),
	}
}

func (n *NFA) newState() NFAState {
	n.state += 1
	return n.state
}

func (n *NFA) add(cur, next NFAState, r RUNE) {
	ps, ok := n.paths[cur]
	if !ok {
		ps = make(map[RUNE][]NFAState)
	}
	ps[r] = append(ps[r], next)
	n.paths[cur] = ps
}

func (n *NFA) Match(str string) bool {
	s := &scanner{line: []byte(str)}

	if n.Debug {
		fmt.Println(n.startState, n.endState)
		fmt.Println(n.paths)
	}
	state := n.startState
	states := n.epsilonClosure(map[NFAState]struct{}{state: {}})
	for {
		r, err := s.next()
		if err != nil {
			return false
		}
		if r == eof {
			break
		}
		newStates := map[NFAState]struct{}{}
		if n.Debug {
			fmt.Printf("cur %c, states: %v\n", r, states)
		}
		for cur, _ := range states {
			for next, _ := range n.nextStates(cur, RUNE(r)) {
				newStates[next] = struct{}{}
			}
		}
		states = n.epsilonClosure(newStates)
	}
	if n.Debug {
		fmt.Println("end states: ", states)
	}

	if _, ok := states[n.endState]; ok {
		return true
	}

	return false
}

func (n *NFA) epsilonClosure(states map[NFAState]struct{}) map[NFAState]struct{} {
	stacks := map[NFAState]struct{}{}
	for state, _ := range states {
		stacks[state] = struct{}{}
	}

	for len(stacks) > 0 {
		t := NFAState(0)
		for st, _ := range stacks {
			t = st
			delete(stacks, st)
			break
		}
		ps, ok := n.paths[t]
		if !ok {
			continue
		}
		for _, t2 := range ps[Epsilon] {
			if _, ok := states[t2]; !ok {
				states[t2] = struct{}{}
				stacks[t2] = struct{}{}
			}
		}
	}
	return states
}

func (n *NFA) nextStates(cur NFAState, r RUNE) map[NFAState]struct{} {
	states := make(map[NFAState]struct{})
	ps, ok := n.paths[cur]
	if !ok {
		return states
	}
	for _, next := range ps[r] {
		states[next] = struct{}{}
	}
	for _, next := range ps[AnyRUNE] {
		states[next] = struct{}{}
	}
	return states
}

type scanner struct {
	line []byte
}

var errInvalidRune = errors.New("invalid rune")

func (s *scanner) next() (rune, error) {
	if len(s.line) == 0 {
		return eof, nil
	}
	c, size := utf8.DecodeRune(s.line)
	s.line = s.line[size:]
	if c == utf8.RuneError {
		return 0, errInvalidRune
	}
	return c, nil
}

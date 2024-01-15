package regexp

import (
	"fmt"
)

type DFAState int

type DFA struct {
	Debug bool
	paths map[DFAState]map[RUNE]DFAState
	state DFAState

	startState DFAState
	endState   map[DFAState]struct{}
}

func NewDFA() *DFA {
	return &DFA{
		paths:    make(map[DFAState]map[RUNE]DFAState),
		endState: make(map[DFAState]struct{}),
	}
}
func (n *DFA) newState() DFAState {
	n.state += 1
	return n.state
}

func (n *DFA) add(cur, next DFAState, r RUNE) {
	ps, ok := n.paths[cur]
	if !ok {
		ps = make(map[RUNE]DFAState)
	}
	ps[r] = next
	n.paths[cur] = ps
}

func (n *DFA) nextState(cur DFAState, r RUNE) DFAState {
	ps, ok := n.paths[cur]
	if !ok {
		return -1
	}
	return ps[r]
}

func (n *DFA) Match(str string) bool {
	s := &scanner{line: []byte(str)}

	if n.Debug {
		fmt.Printf("start: %d, end: %d \n", n.startState, n.endState)
		fmt.Printf("paths: %d \n", n.paths)
	}
	state := n.startState
	for {
		r, err := s.next()
		if err != nil {
			return false
		}
		if r == eof {
			break
		}
		if n.Debug {
			fmt.Printf("cur %c, state: %v\n", r, state)
		}
		newState := n.nextState(state, RUNE(r))
		if newState <= 0 {
			return false
		}
		state = newState
	}
	if n.Debug {
		fmt.Println("end states: ", state)
	}

	_, ok := n.endState[state]
	return ok
}

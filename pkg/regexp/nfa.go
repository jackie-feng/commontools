package regexp

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
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

type NFAStates map[NFAState]struct{}

func (n NFAStates) key() string {
	states := []NFAState{}
	for s, _ := range n {
		states = append(states, s)
	}
	sort.Slice(states, func(i, j int) bool {
		return states[i] < states[j]
	})
	strs := make([]string, 0, len(states))
	for _, s := range states {
		strs = append(strs, strconv.Itoa(int(s)))
	}
	return strings.Join(strs, ",")
}

func (n *NFA) ToDFA() *DFA {
	d := NewDFA()
	s1 := d.newState()
	d.startState = s1
	states1 := n.epsilonClosure(NFAStates{n.startState: {}})
	states1Key := states1.key()
	statesMap := map[string]DFAState{states1Key: s1}

	unsearchSet := map[string]NFAStates{states1.key(): states1}
	searchedSet := map[string]struct{}{}
	for len(unsearchSet) > 0 {
		var key string
		var states = NFAStates{}
		for key, states = range unsearchSet {
			break
		}
		delete(unsearchSet, key)
		if _, ok := searchedSet[key]; ok {
			continue
		}
		searchedSet[key] = struct{}{}
		receiveChars := n.receivers(states)
		for r, _ := range receiveChars {
			newStates := n.epsilonClosure(n.move(states, r))
			newStatesKey := newStates.key()

			if _, ok := unsearchSet[newStatesKey]; !ok {
				unsearchSet[newStatesKey] = newStates
			}
			if _, ok := statesMap[newStatesKey]; !ok {
				s2 := d.newState()
				statesMap[newStatesKey] = s2
			}
			d.add(statesMap[key], statesMap[newStatesKey], r)
			if _, ok := newStates[n.endState]; ok {
				d.endState[statesMap[newStatesKey]] = struct{}{}
			}
		}
	}

	return d
}

func (n *NFA) receivers(states NFAStates) map[RUNE]struct{} {
	receiveChars := map[RUNE]struct{}{}
	for s, _ := range states {
		ps, ok := n.paths[s]
		if !ok {
			continue
		}
		for r, _ := range ps {
			if r == Epsilon {
				continue
			}
			receiveChars[r] = struct{}{}
		}
	}
	return receiveChars
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
		fmt.Printf("start: %d, end: %d \n", n.startState, n.endState)
		fmt.Printf("paths: %d \n", n.paths)
	}
	state := n.startState
	states := n.epsilonClosure(NFAStates{state: {}})
	for {
		r, err := s.next()
		if err != nil {
			return false
		}
		if r == eof {
			break
		}
		if n.Debug {
			fmt.Printf("cur %c, states: %v\n", r, states)
		}
		newStates := n.move(states, RUNE(r))
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

func (n *NFA) epsilonClosure(curStates NFAStates) NFAStates {
	stacks := NFAStates{}
	for state, _ := range curStates {
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
			if _, ok := curStates[t2]; !ok {
				curStates[t2] = struct{}{}
				stacks[t2] = struct{}{}
			}
		}
	}
	return curStates
}

// 对 states 状态集, 输入 r 字符后, 可以到达的状态集
func (n *NFA) move(states NFAStates, r RUNE) NFAStates {
	newStates := NFAStates{}
	for cur, _ := range states {
		for next, _ := range n.nextStates(cur, r) {
			newStates[next] = struct{}{}
		}
	}
	return newStates
}

// 当前 cur 状态, 输入 r 字符后, 可以到达的状态集
func (n *NFA) nextStates(cur NFAState, r RUNE) NFAStates {
	states := make(NFAStates)
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

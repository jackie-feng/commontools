package regexp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNFA_ToDFA(t *testing.T) {
	// 构建（a|b）*abb 的 NFA
	n := NewNFA()
	n.startState = 0
	n.add(0, 1, Epsilon)
	n.add(1, 2, Epsilon)
	n.add(1, 4, Epsilon)
	n.add(2, 3, 'a')
	n.add(4, 5, 'b')
	n.add(3, 6, Epsilon)
	n.add(5, 6, Epsilon)
	n.add(6, 1, Epsilon)
	n.add(6, 7, Epsilon)
	n.add(0, 7, Epsilon)
	n.add(7, 8, 'a')
	n.add(8, 9, 'b')
	n.add(9, 10, 'b')
	n.endState = 10

	// A: {0,1,2,4,7}
	// B: {1,2,3,4,6,7,8}
	// C: {1,2,4,5,6,7}
	// D: {1,2,4,5,6,7,9}
	// E: {1,2,4,5,6,7,10}

	d := n.ToDFA()
	A := d.startState
	B := d.nextState(A, 'a')
	C := d.nextState(A, 'b')
	D := d.nextState(B, 'b')
	E := d.nextState(D, 'b')

	assert.Equal(t, B, d.nextState(B, 'a'))
	assert.Equal(t, B, d.nextState(C, 'a'))
	assert.Equal(t, C, d.nextState(C, 'b'))
	assert.Equal(t, B, d.nextState(D, 'a'))
	assert.Equal(t, B, d.nextState(E, 'a'))
	assert.Equal(t, C, d.nextState(E, 'b'))
}

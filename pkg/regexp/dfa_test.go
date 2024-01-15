package regexp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFA_Match(t *testing.T) {
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

	d := n.ToDFA()
	assert.Equal(t, true, d.Match("abb"))
	assert.Equal(t, true, d.Match("aabb"))
	assert.Equal(t, true, d.Match("babb"))
	assert.Equal(t, true, d.Match("abababb"))
}

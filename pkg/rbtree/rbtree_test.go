package rbtree

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type IntIndex int

func (i IntIndex) Compare(x Index) int8 {
	v := x.(IntIndex)
	if i == x {
		return 0
	} else if i > v {
		return 1
	} else {
		return -1
	}
}

func TestRBTree_print(t *testing.T) {
	tree := &RBTree{}

	n1 := &RBNode{value: IntIndex(1)}
	n2 := &RBNode{value: IntIndex(2)}
	n3 := &RBNode{value: IntIndex(3)}
	n4 := &RBNode{value: IntIndex(4)}
	n5 := &RBNode{value: IntIndex(5)}
	n6 := &RBNode{value: IntIndex(6)}
	n7 := &RBNode{value: IntIndex(7)}
	n8 := &RBNode{value: IntIndex(8)}
	n9 := &RBNode{value: IntIndex(9)}
	n10 := &RBNode{value: IntIndex(10)}
	n11 := &RBNode{value: IntIndex(11)}
	n12 := &RBNode{value: IntIndex(12)}
	n13 := &RBNode{value: IntIndex(13)}
	n14 := &RBNode{value: IntIndex(14)}
	n15 := &RBNode{value: IntIndex(15)}

	n8.left = n4
	n8.right = n12
	n4.left = n2
	n4.right = n6
	n12.left = n10
	n12.right = n14
	n2.left = n1
	n2.right = n3
	n6.left = n5
	n6.right = n7
	n10.left = n9
	n10.right = n11
	n14.left = n13
	n14.right = n15

	tree.root = n8
	// tree.print()
}

func TestRBTree_valid(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		tree := &RBTree{}
		n1 := &RBNode{value: IntIndex(1)}
		n2 := &RBNode{value: IntIndex(2), isblack: true}
		n3 := &RBNode{value: IntIndex(3)}
		n4 := &RBNode{value: IntIndex(4)}
		n5 := &RBNode{value: IntIndex(5)}
		n6 := &RBNode{value: IntIndex(6), isblack: true}
		n7 := &RBNode{value: IntIndex(7)}
		n8 := &RBNode{value: IntIndex(8), isblack: true}
		n9 := &RBNode{value: IntIndex(9)}
		n10 := &RBNode{value: IntIndex(10), isblack: true}
		n11 := &RBNode{value: IntIndex(11)}
		n12 := &RBNode{value: IntIndex(12)}
		n13 := &RBNode{value: IntIndex(13)}
		n14 := &RBNode{value: IntIndex(14), isblack: true}
		n15 := &RBNode{value: IntIndex(15)}

		n8.left = n4
		n4.parent = n8
		n8.right = n12
		n12.parent = n8
		n4.left = n2
		n2.parent = n4
		n4.right = n6
		n6.parent = n4
		n12.left = n10
		n10.parent = n12
		n12.right = n14
		n14.parent = n12
		n2.left = n1
		n1.parent = n2
		n2.right = n3
		n3.parent = n2
		n6.left = n5
		n5.parent = n6
		n6.right = n7
		n7.parent = n6
		n10.left = n9
		n9.parent = n10
		n10.right = n11
		n11.parent = n10
		n14.left = n13
		n13.parent = n14
		n14.right = n15
		n15.parent = n14

		tree.root = n8
		assert.Equal(t, true, tree.valid())
	})
	t.Run("case2", func(t *testing.T) {
		tree := &RBTree{}
		n4 := &RBNode{value: IntIndex(4)}
		n5 := &RBNode{value: IntIndex(5)}
		n6 := &RBNode{value: IntIndex(6), isblack: true}
		n7 := &RBNode{value: IntIndex(7)}
		n8 := &RBNode{value: IntIndex(8), isblack: true}
		n9 := &RBNode{value: IntIndex(9)}
		n10 := &RBNode{value: IntIndex(10), isblack: true}
		n11 := &RBNode{value: IntIndex(11)}
		n12 := &RBNode{value: IntIndex(12)}
		n13 := &RBNode{value: IntIndex(13)}
		n14 := &RBNode{value: IntIndex(14), isblack: true}
		n15 := &RBNode{value: IntIndex(15)}

		n8.left = n4
		n8.right = n12
		n4.right = n6
		n12.left = n10
		n12.right = n14
		n6.left = n5
		n6.right = n7
		n10.left = n9
		n10.right = n11
		n14.left = n13
		n14.right = n15

		tree.root = n8
		assert.Equal(t, false, tree.valid())
	})
	t.Run("case3", func(t *testing.T) {
		tree := &RBTree{}
		n1 := &RBNode{value: IntIndex(1)}
		n2 := &RBNode{value: IntIndex(2), isblack: true}
		n3 := &RBNode{value: IntIndex(3)}
		n4 := &RBNode{value: IntIndex(4)}
		n5 := &RBNode{value: IntIndex(5)}
		n6 := &RBNode{value: IntIndex(6), isblack: true}
		n7 := &RBNode{value: IntIndex(7)}
		n8 := &RBNode{value: IntIndex(8), isblack: true}
		n9 := &RBNode{value: IntIndex(9)}
		n10 := &RBNode{value: IntIndex(10), isblack: true}
		n11 := &RBNode{value: IntIndex(11)}
		n12 := &RBNode{value: IntIndex(12)}
		n13 := &RBNode{value: IntIndex(13)}
		n14 := &RBNode{value: IntIndex(14), isblack: true}
		n15 := &RBNode{value: IntIndex(15)}
		n16 := &RBNode{value: IntIndex(16)}

		n8.left = n4
		n8.right = n12
		n4.left = n2
		n4.right = n6
		n12.left = n10
		n12.right = n14
		n2.left = n1
		n2.right = n3
		n6.left = n5
		n6.right = n7
		n10.left = n9
		n10.right = n11
		n14.left = n13
		n14.right = n15
		n15.right = n16

		tree.root = n8
		assert.Equal(t, false, tree.valid())
	})
}

func TestRBtree_findNode(t *testing.T) {
	tree := &RBTree{}
	n1 := &RBNode{value: IntIndex(1)}
	n2 := &RBNode{value: IntIndex(2), isblack: true}
	n3 := &RBNode{value: IntIndex(3)}
	n4 := &RBNode{value: IntIndex(4)}
	n5 := &RBNode{value: IntIndex(5)}
	n6 := &RBNode{value: IntIndex(6), isblack: false}
	n7 := &RBNode{value: IntIndex(7)}
	n8 := &RBNode{value: IntIndex(8), isblack: true}
	n9 := &RBNode{value: IntIndex(9)}
	n10 := &RBNode{value: IntIndex(10), isblack: true}
	n11 := &RBNode{value: IntIndex(11)}
	n12 := &RBNode{value: IntIndex(12)}
	n13 := &RBNode{value: IntIndex(13)}
	n14 := &RBNode{value: IntIndex(14), isblack: true}
	n15 := &RBNode{value: IntIndex(15)}

	n8.left = n4
	n8.right = n12
	n4.left = n2
	n4.right = n6
	n12.left = n10
	n12.right = n14
	n2.left = n1
	n2.right = n3
	n6.left = n5
	n6.right = n7
	n10.left = n9
	n10.right = n11
	n14.left = n13
	n14.right = n15
	tree.root = n8
	v1 := tree.findNode(tree.root, IntIndex(1))
	assert.Equal(t, n1, v1)
	v9 := tree.findNode(tree.root, IntIndex(9))
	assert.Equal(t, n9, v9)
	v10 := tree.findNode(tree.root, IntIndex(10))
	assert.Equal(t, n10, v10)
	v17 := tree.findNode(tree.root, IntIndex(17))
	assert.Nil(t, v17)
}

func TestRBTree_Add(t *testing.T) {
	tree := &RBTree{}
	tree.Add(&RBNode{value: IntIndex(1)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(2)})
	//tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(7)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(8)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(9)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(4)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(12)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(5)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(6)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(13)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(3)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(14)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(15)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(10)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(11)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(16)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
	tree.Add(&RBNode{value: IntIndex(2)})
	// tree.print()
	assert.Equal(t, true, tree.valid())
}

func TestRBTree_Remove(t *testing.T) {
	tree := &RBTree{}
	tree.Add(&RBNode{value: IntIndex(1)})
	tree.Add(&RBNode{value: IntIndex(2)})
	tree.Add(&RBNode{value: IntIndex(7)})
	tree.Add(&RBNode{value: IntIndex(8)})
	tree.Add(&RBNode{value: IntIndex(9)})
	tree.Add(&RBNode{value: IntIndex(4)})
	tree.Add(&RBNode{value: IntIndex(12)})
	tree.Add(&RBNode{value: IntIndex(5)})
	tree.Add(&RBNode{value: IntIndex(6)})
	tree.Add(&RBNode{value: IntIndex(13)})
	tree.Add(&RBNode{value: IntIndex(3)})
	tree.Add(&RBNode{value: IntIndex(14)})
	tree.Add(&RBNode{value: IntIndex(15)})
	tree.Add(&RBNode{value: IntIndex(10)})
	tree.Add(&RBNode{value: IntIndex(11)})
	tree.Add(&RBNode{value: IntIndex(16)})
	// tree.Add(&RBNode{value: 2})
	assert.Equal(t, true, tree.valid())

	// tree.print()
	tree.Remove(IntIndex(1))
	assert.Equal(t, true, tree.valid())
	// tree.print()
	tree.Remove(IntIndex(10))
	assert.Equal(t, true, tree.valid())
	// tree.print()
	tree.Remove(IntIndex(13))
	assert.Equal(t, true, tree.valid())
	// tree.print()
	tree.Remove(IntIndex(2))
	assert.Equal(t, true, tree.valid())
	// tree.print()
	tree.Remove(IntIndex(7))
	assert.Equal(t, true, tree.valid())
	// tree.print()
	tree.Remove(IntIndex(8))
	assert.Equal(t, true, tree.valid())
	// tree.print()
	tree.Remove(IntIndex(3))
	assert.Equal(t, true, tree.valid())
	// tree.print()
	tree.Remove(IntIndex(5))
	assert.Equal(t, true, tree.valid())

	for i := 20; i < 50; i++ {
		tree.Add(&RBNode{value: IntIndex(i)})
	}
	assert.Equal(t, true, tree.valid())
	for i := 0; i < 10; i++ {
		x := rand.Intn(30) + 20
		tree.Remove(IntIndex(x))
		assert.Equal(t, true, tree.valid())
	}

	t.Run("addNodeWeight case1", func(t *testing.T) {
		tree := &RBTree{}
		n1 := &RBNode{value: IntIndex(1), isblack: true}
		n2 := &RBNode{value: IntIndex(2), isblack: true}
		n3 := &RBNode{value: IntIndex(3), isblack: true}
		n4 := &RBNode{value: IntIndex(4), isblack: false}
		n5 := &RBNode{value: IntIndex(5), isblack: true}
		tree.root = n2
		n2.left = n1
		n2.right = n4
		n4.left = n3
		n4.right = n5
		n1.parent = n2
		n1.parent = n2
		n4.parent = n2
		n3.parent = n4
		n5.parent = n4
		assert.Equal(t, true, tree.valid())
		tree.Remove(n1.value)
		assert.Equal(t, true, tree.valid())
	})
}

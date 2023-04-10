package lcs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSequence(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		text1 := "abcde"
		text2 := "ace"

		res := Sequence(len(text1), len(text2), func(i, j int) bool {
			return text1[i] == text2[j]
		})

		assert.Equal(t, int(3), len(res))
		assert.Equal(t, 0, res[0].Idx1)
		assert.Equal(t, 0, res[0].Idx2)
		assert.Equal(t, 2, res[1].Idx1)
		assert.Equal(t, 1, res[1].Idx2)
		assert.Equal(t, 4, res[2].Idx1)
		assert.Equal(t, 2, res[2].Idx2)
	})

	t.Run("case2", func(t *testing.T) {
		l1 := []int{0, 2, 4, 5, 5, 9, 12, 12, 13, 15, 17, 19, 23, 24, 24, 24, 26, 26, 26, 28}
		l2 := []int{6, 6, 8, 12, 13, 17, 18, 19, 22, 22, 24, 24, 26, 27, 29}

		res := Sequence(len(l1), len(l2), func(i, j int) bool {
			return l1[i] == l2[j]
		})
		assert.Equal(t, 7, len(res))
		assert.Equal(t, Pair{6, 3}, res[0])
		assert.Equal(t, Pair{8, 4}, res[1])
		assert.Equal(t, Pair{10, 5}, res[2])
		assert.Equal(t, Pair{11, 7}, res[3])
		assert.Equal(t, Pair{13, 10}, res[4])
		assert.Equal(t, Pair{14, 11}, res[5])
		assert.Equal(t, Pair{16, 12}, res[6])
	})
}

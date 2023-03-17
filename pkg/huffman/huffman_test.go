package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHuffman_Encode(t *testing.T) {
	a := "we will we will r u"
	a = "ABBBCCCCCCCCDDDDDDEE"
	m := map[byte]int{}
	for i := 0; i < len(a); i++ {
		m[a[i]] = m[a[i]] + 1
	}

	chars := make([]Char, 0)
	for k, v := range m {
		chars = append(chars, Char{Value: k, Weight: v})
	}

	sort.Slice(chars, func(i, j int) bool {
		return chars[i].Weight <= chars[j].Weight
	})

	t.Logf("%v", chars)
	h := NewHuffman(chars)

	t.Logf("%v", h.codes)
	code := h.Encode(a)

	t.Logf("%v, len: %d", code, len(code))
	assert.Equal(t, "11101101101100000000010101010101011111111", code)

	str := h.Decode(code)
	assert.Equal(t, a, str)
}

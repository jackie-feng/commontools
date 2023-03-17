package main

import (
	"fmt"
	"sort"
	"strings"
	"unsafe"
)

type node struct {
	weight int
	value  byte
	left   *node
	right  *node
}

type Huffman struct {
	nodes []*node
	codes map[byte]string
}

type Char struct {
	Value  byte
	Weight int
}

func (c Char) String() string {
	return fmt.Sprintf("<%c:%d>", c.Value, c.Weight)
}

func NewHuffman(chars []Char) *Huffman {
	nodes := make([]*node, len(chars))
	for i := 0; i < len(chars); i++ {
		nodes[i] = &node{weight: chars[i].Weight, value: chars[i].Value}
	}
	h := &Huffman{nodes: nodes, codes: map[byte]string{}}
	h.generateTree()
	return h
}

func (h *Huffman) generateTree() {
	for {
		if h.merge() {
			break
		}
	}
	h.walk(h.nodes[0], h.codes, "")
}

func (h *Huffman) Encode(str string) string {
	b := strings.Builder{}

	for i := 0; i < len(str); i++ {
		c := str[i]
		code, ok := h.codes[c]
		if !ok {
			panic("invalid idx")
		}
		b.WriteString(code)
	}
	return b.String()
}

func (h *Huffman) Decode(str string) string {
	root := h.nodes[0]
	cur := root
	res := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c == '0' {
			cur = cur.left
		} else {
			cur = cur.right
		}
		if cur.left == nil {
			res = append(res, cur.value)
			cur = root
		}
	}
	return *(*string)(unsafe.Pointer(&res))
}

func (h *Huffman) walk(node *node, res map[byte]string, prefix string) {
	if node.left == nil && node.right == nil {
		res[node.value] = prefix
		return
	}

	h.walk(node.left, res, strings.Join([]string{prefix, "0"}, ""))
	h.walk(node.right, res, strings.Join([]string{prefix, "1"}, ""))
}

func (h *Huffman) merge() (done bool) {
	if len(h.nodes) <= 1 {
		return true
	}

	n := &node{
		left:   h.nodes[0],
		right:  h.nodes[1],
		weight: h.nodes[0].weight + h.nodes[1].weight,
	}
	if len(h.nodes) == 2 {
		h.nodes = []*node{n}
		return true
	}

	idx := sort.Search(len(h.nodes)-2, func(i int) bool {
		return h.nodes[i+2].weight > n.weight
	})

	copy(h.nodes[:idx], h.nodes[2:2+idx])
	h.nodes[idx] = n
	copy(h.nodes[idx+1:], h.nodes[2+idx:])
	h.nodes = h.nodes[:len(h.nodes)-1]
	return false
}

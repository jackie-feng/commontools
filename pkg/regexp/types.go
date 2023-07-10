package regexp

type Node interface{}

type AndNode struct {
	Left  Node
	Right rune
}

type OrNode struct {
	Left  Node
	Right Node
}

type RepeatedNode struct {
	Val Node
}

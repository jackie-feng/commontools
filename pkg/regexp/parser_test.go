package regexp

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var c rune = 'a'
	var x rune = '😊'
	fmt.Println(x)
	var i int = 98
	i1 := int(c)
	fmt.Println("'a' convert to", i1)
	c1 := rune(i)
	fmt.Println("98 convert to", string(c1))

	//string to rune
	for _, char := range []rune("世界你好") {
		fmt.Println(string(char))
	}
}

func TestParser(t *testing.T) {
	YoYoErrorVerbose = true
	//YoYoDebug = 5
	line := "abc(a|bb*)cc*"
	exprLex := NewLexer([]byte(line))
	exprParser := YoYoNewParser()
	exprParser.Parse(exprLex)
	t.Logf("%#v", exprLex.ast)
}

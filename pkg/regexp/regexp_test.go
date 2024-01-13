package regexp

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexp(t *testing.T) {
	reg, _ := regexp.Compile("abc(a|bb*)cc*")
	assert.Equal(t, true, reg.MatchString("abcac"))
	assert.Equal(t, true, reg.MatchString("abcbc"))
	assert.Equal(t, true, reg.MatchString("abcbcccc"))
	assert.Equal(t, true, reg.MatchString("abcbbbbcccc"))

	reg, _ = regexp.Compile("a(bc)*(ac|bcd*)*cc*")
	assert.Equal(t, true, reg.MatchString("ac"))
	assert.Equal(t, true, reg.MatchString("abcc"))
	assert.Equal(t, true, reg.MatchString("abcbcbcc"))
	assert.Equal(t, true, reg.MatchString("abcbcbcdc"))
	assert.Equal(t, true, reg.MatchString("abcacbcacbcbcdbcbcc"))
}

func BenchmarkName(b *testing.B) {
	patten := "a(b|d)*c"
	str := "abdbdbc"

	reg, _ := regexp.Compile(patten)
	reg2 := Compile(patten)
	b.Run("regexp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reg.Match([]byte(str))
		}
	})

	b.Run("regexp2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reg2.Match(str)
		}
	})
}

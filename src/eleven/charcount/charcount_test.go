package charcount

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		r    *strings.Reader
		want count
	}{
		{strings.NewReader("abcd"), count{map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1},
			[utf8.UTFMax + 1]int{0, 4, 0, 0, 0}, 0}},
	}

	for i, test := range tests {
		got := CharCount(*test.r)
		if !equal(got, test.want) {
			t.Errorf("%d. got %v, want %v", i, got, test.want)
		}
	}
}

func equal(c1, c2 count) bool {
	if len(c1.counts) != len(c2.counts) {
		return false
	}

	for k, v1 := range c1.counts {
		v2, ok := c2.counts[k]
		if !ok {
			return false
		}
		if v1 != v2 {
			return false
		}
	}

	if c1.utflen != c2.utflen {
		return false
	}

	if c1.invalid != c2.invalid {
		return false
	}
	return true
}

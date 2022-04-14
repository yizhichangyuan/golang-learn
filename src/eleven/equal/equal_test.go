package equal

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"x||x||x", "||", 3},
	}

	for i, test := range tests {
		if got := len(strings.Split(test.s, test.sep)); !equal(test.want, got) {
			t.Errorf("%d. Split(%q, %q) got %d, want %d\n", i, test.s, test.sep, got, test.want)
		}
	}
}

func equal(x, y int) bool {
	if x != y {
		return false
	}
	return true
}

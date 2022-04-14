package charcount

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type count struct {
	counts  map[rune]int         // counts of Unicode characters
	utflen  [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid int                  // count of invalid UTF-8 characters
}

func CharCount(in strings.Reader) count {
	c := count{
		counts: make(map[rune]int),
	}
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			c.invalid++
			continue
		}
		c.counts[r]++
		c.utflen[n]++
	}
	return c
}

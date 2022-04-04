package main

import (
	"fmt"
	"math/cmplx"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g C", c)
}

func main() {
	//var c Celsius
	//c.String()
	fmt.Println(1i * 1i)
	var x = complex(1, 2)
	fmt.Println(x)
	fmt.Println(cmplx.Sqrt(x))

	s := "abc"
	fmt.Println(s[0:2])
	// return byte value, not a char
	fmt.Println(s[0])
	fmt.Println(s + string(s[0]))

	// s contains chinese and eng with unicode, use utf.RuneCountInString to get how many unicode char
	s = "hello, 世界"
	fmt.Println(len(s))                    // len can only get byte number : 13 byte
	fmt.Println(utf8.RuneCountInString(s)) // 9 unicode
	// len(s) return byte length, not character length
	// s[8:] cannot get '界'， must decodeRune to get
	fmt.Println(s[8:])

	// decode utf8 to get Chinese Character
	for i := 0; i < len(s); {
		// return character and byte length of the character
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// i mean the byte index of character in s
	// r mean character
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	s = "世界"
	fmt.Println(s)
	// utf8 transfer to unicode
	r := []rune(s)
	fmt.Println(r)
	// unicode transfer to utf8
	fmt.Println(string(r))

	c := 'c'
	fmt.Println(unicode.IsLower(rune(c)))

	s = "abc a"
	b := []byte(s)
	fmt.Println(b)
	fmt.Println(strings.Contains(s, "a"))
	fmt.Println(strings.Count(s, "a"))
	fmt.Println(strings.Fields(s))

	var n int
	number, _ := fmt.Scanf("%d", &n)
	fmt.Println(n)
	fmt.Println(number)
}

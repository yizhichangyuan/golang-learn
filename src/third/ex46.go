package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func removeEmpty(bytes []byte) []byte {
	end := len(bytes)
	for i := 0; i < end; {
		r, size := utf8.DecodeRune(bytes[i:])
		if unicode.IsSpace(r) {
			second, size2 := utf8.DecodeRune(bytes[i+size:])
			if unicode.IsSpace(second) {
				copy(bytes[i:], bytes[i+size:])
				end -= size
			} else {
				i += size2
			}
		} else {
			i += size
		}
	}
	return bytes[:end]
}

func main() {
	bytes := []byte("  a  c  ")
	bytes = removeEmpty(bytes)
	fmt.Printf("%s", bytes)
}

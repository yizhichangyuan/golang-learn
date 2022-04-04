package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseBytes(bytes []byte) {
	for i := 0; i < len(bytes); {
		_, size := utf8.DecodeRune(bytes[i:])
		reverseB(bytes[i : i+size])
		i += size
	}
	reverseB(bytes)
}

func reverseB(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func main() {
	bytes := []byte("一 二 三")
	reverseBytes(bytes)
	fmt.Printf("%s", bytes)
}

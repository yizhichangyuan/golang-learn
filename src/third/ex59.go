package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCount := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// set scan mode to split empty-space word
	//input.Split(bufio.ScanWords)
	for input.Scan() {
		wordCount[input.Text()]++
	}
	for word, count := range wordCount {
		fmt.Printf("word: %s, count: %d\n", word, count)
	}
}

package main

import (
	"bytes"
	"fmt"
)

// 当 uint 为 unitSize 位时，unitSize 个 1111111... 右移 63 位变成 1，32 再左移 1 位变成 unitSize
// 当 uint 位 32 位时，32 个 1111111... 右移 63 位变成 0, 32 再左移 0 位 变成 32
var unitSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/unitSize, x%unitSize
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/unitSize, x%unitSize
	// for word > len(s.words)-1
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith union all elements of the set t into the set s
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= t.words[i]
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}"
func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < unitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", unitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for j := 0; j < unitSize; j++ {
			if word&(1<<j) != 0 {
				count++
			}
		}
	}
	return count
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}
	word, bit := x/unitSize, x%unitSize
	// &^ removes all bit position which equals 1 in both a b
	s.words[word] &^= 1 << bit
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	copyWords := IntSet{}
	// make enough space to place all elements of s
	copyWords.words = make([]uint, s.Len())
	copy(copyWords.words, s.words)
	return &copyWords
}

func (s *IntSet) AddAll(ints ...int) {
	for _, v := range ints {
		word, bit := v/unitSize, v%unitSize
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	minLen := s.Len()
	if s.Len() > t.Len() {
		minLen = t.Len()
	}
	for i := 0; i < minLen; i++ {
		s.words[i] &= t.words[i]
	}
	// very important
	s.words = s.words[:minLen]
}

// s 有， t 没有，就是将s中有t也有的进行清除，也就是将s和t中都为1的进行清除
func (s *IntSet) DifferenceWith(t *IntSet) {
	minLen := s.Len()
	if s.Len() > t.Len() {
		minLen = t.Len()
	}
	for i := 0; i < minLen; i++ {
		s.words[i] &^= t.words[i]
	}

	// 将s.words中末尾为0的进行清除
	for len(s.words) != 0 && s.words[len(s.words)-1] == 0 {
		s.words = s.words[:len(s.words)-1]
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	minLen := s.Len()
	if s.Len() > t.Len() {
		minLen = t.Len()
	}
	for i := 0; i < minLen; i++ {
		s.words[i] ^= t.words[i]
	}
	for len(s.words) != 0 && s.words[len(s.words)-1] == 0 {
		s.words = s.words[:len(s.words)-1]
	}
}

func (s *IntSet) Elems() []uint {
	var result []uint
	for i, word := range s.words {
		for j := 0; j < unitSize; j++ {
			if word&(1<<j) != 0 {
				result = append(result, uint(unitSize*i+j))
			}
		}
	}
	return result
}

func main() {
	var s IntSet
	s.Add(1)
	s.Add(144)
	s.Add(9)
	fmt.Println(s.String())
	fmt.Println(&s)

	x := s.Copy()
	fmt.Println(x.Len())
	s.IntersectWith(x)
	x.words = x.words[:x.Len()-1]
	s.IntersectWith(x)
	fmt.Println(s.String())
}

package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	//if s.Len() % 2 == 0 {
	//	center1 := (s.Len() - 1) / 2
	//	center2 := s.Len() / 2
	//	if !(!s.Less(center1, center2) && !s.Less(center2, center1)){
	//		return false
	//	}
	//	for k := 0; k <= center1; k++ {
	//		if !(!s.Less(center1-k, center2+k) && !s.Less(center2+k, center1-k)){
	//			return false
	//		}
	//	}
	//} else {
	//	center := (s.Len() - 1) / 2
	//	for k := 0; k <= center; k++ {
	//		if !(!s.Less(center-k, center+k) && !s.Less(center+k, center-k)){
	//			return false
	//		}
	//	}
	//}
	//return true
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

type runes []rune

func (r runes) Len() int {
	return len(r)
}

func (r runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r runes) Less(i, j int) bool {
	return r[i] < r[j]
}

func main() {
	s := "adccde"
	r := runes(s)
	fmt.Println(IsPalindrome(r))
}

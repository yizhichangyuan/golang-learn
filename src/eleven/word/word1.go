package word

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var letters []rune = make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

//func main() {
//	// for range 迭代的字符串s中的每个rune，index对应的是字节，所以可能不对应
//	for i, r := range "été" {
//		fmt.Println(i, string(r))
//	}
//}

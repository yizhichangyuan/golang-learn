package main

import "fmt"

func removeDuplicate(strings []string) []string {
	end := len(strings)
	for i := 0; i+1 < end; {
		if strings[i] == strings[i+1] {
			start := i
			i++
			end = end - 1
			for i+1 < end && strings[i] == strings[i+1] {
				i++
				end--
			}
			copy(strings[start+1:], strings[i+1:])
			strings = strings[:end]
			i = start // 很重要，指针回到原来
		} else {
			i++
		}
	}
	return strings[:end]
}

func main() {
	a := []string{"a", "a", "a", " ", " ", "c", "c"}
	a = removeDuplicate(a)
	fmt.Println(a)
}

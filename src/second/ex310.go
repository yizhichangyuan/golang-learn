package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	if len(s) == 0 {
		return s
	}
	str := []byte(s)
	var buf bytes.Buffer
	sign := str[0]
	if sign == '+' || sign == '-' {
		buf.WriteByte(sign)
		str = str[1:]
	}

	last := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			last = str[i:]
			str = str[:i]
		}
	}
	for i := 0; i < len(str); i++ {
		if (len(str)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(str[i])
	}
	buf.WriteString(string(last))
	return buf.String()
}

func main() {
	fmt.Println(comma("1232.32"))
}

package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buffer bytes.Buffer
	buffer.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%d", v)
	}
	buffer.WriteByte(']')
	return buffer.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}

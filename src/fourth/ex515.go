package main

import (
	"fmt"
	"strings"
)

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("must have at least one number")
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func Join(sep string, vals ...string) string {
	var builder strings.Builder
	builder.WriteString(vals[0])
	for _, val := range vals[1:] {
		builder.WriteString(sep + val)
	}
	return builder.String()
}

func main() {
	fmt.Println(Join("/", "a"))
}

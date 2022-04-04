package main

import (
	"fmt"
	"strconv"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

// appendValues appends the element of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// add value to tree, follow left value smaller than point value.库尔导出
func add(t *tree, value int) *tree {
	if t == nil {
		// equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var builder strings.Builder
	if t == nil {
		return ""
	}

	builder.WriteString(t.left.String())
	builder.WriteString(strconv.Itoa(t.value))
	builder.WriteString(t.right.String())
	return builder.String()
}

func main() {
	var values []int
	fmt.Println(values[:0] == nil)
	v := sort([]int{1, 3, 4, 5})
	fmt.Println(v.String())
}

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	var attrs []map[string]string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
			// 即使该ele没有attr元素，也要放入一个进入attrs中，是为了pop时attr和对应ele一起出栈
			attr := make(map[string]string)
			for _, a := range tok.Attr {
				attr[a.Name.Local] = a.Value
			}
			attrs = append(attrs, attr)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
			attrs = attrs[:len(attrs)-1]
		case xml.CharData:
			if containsAll(toSlice(stack, attrs), os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(toSlice(stack, attrs), " "), tok)
			}
		}
	}
}

func toSlice(stack []string, attrs []map[string]string) []string {
	var unionSlice []string
	for _, eleName := range stack {
		unionSlice = append(unionSlice, eleName)
	}
	for _, attr := range attrs {
		for name, value := range attr {
			unionSlice = append(unionSlice, fmt.Sprintf("%s=%s", name, value))
		}
	}
	return unionSlice
}

// slice相当于结构体，结构体中由三个成员：指向底层数组的指针、长度len、容量cap
// 所以作为函数参数变量进行值拷贝的时候，指针（一个新拷贝的地址）、长度（值拷贝）、容量（值拷贝）
// 所以可以通过值拷贝的指针进行修改底层数组某些元素的值从而改变外部slice
// 但是不能修改外部slice的长度和容量，因为长度和容量是值拷贝
// 所以slice在函数内部最多只能进行修改原slice元素值，但是进行包含长度以及容量的任何改变都不会影响到外部的slice
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
			x = x[1:]
		} else {
			x = x[1:]
		}
	}
	return false
}

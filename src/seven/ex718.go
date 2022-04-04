package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Node interface {
	String() string
}

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	decoder := xml.NewDecoder(os.Stdin)
	node, err := parse(decoder)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", node)
}

func (e *Element) String() string {
	var attrs, children string
	for _, attr := range e.Attr {
		attrs += fmt.Sprintf(" %s=%q", attr.Name.Local, attr.Value)
	}
	for _, child := range e.Children {
		children += child.String()
	}
	return fmt.Sprintf("<%s %s>%s</%s>",
		e.Type.Local, attrs, children, e.Type.Local)
}

func parse(decoder *xml.Decoder) (Node, error) {
	var stack []*Element
	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, fmt.Errorf("parse fail")
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			ele := Element{tok.Name, tok.Attr, []Node{}}
			// 如果该结点确实属于某个结点的孩子结点，那么说明该父节点还没有遇到EndElement
			// 因此stack存放的都是还没有遇到EndElement的结点，遇到EndElement都踢出
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, &ele)
			}
			stack = append(stack, &ele)
		case xml.EndElement:
			if len(stack) == 1 {
				return stack[0], nil
			} else if len(stack) == 0 {
				return nil, fmt.Errorf("unexcpected tag closing")
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			// 可能存在有些CharData没有父节点，也就是没有包含任何标签的
			if len(stack) > 0 {
				c := CharData(string(tok))
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, c)
			}
		}
	}
}

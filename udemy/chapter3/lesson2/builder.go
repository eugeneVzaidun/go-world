package main

import (
	"fmt"
	"strings"
)
const (
	indentSize = 2
)

type HTMLElement struct {
	name, text string
	elements []*HTMLElement
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize * indent)	
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize * (indent + 1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, e := range e.elements {
		sb.WriteString(e.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HTMLBuilder struct {
	rootName string
	root *HTMLElement
}

func NewHTMLBuilder(rootName string) *HTMLBuilder {
	return &HTMLBuilder{rootName, &HTMLElement{rootName, "", []*HTMLElement{}}}
}

func (b *HTMLBuilder) String() string {
	return b.root.String()
}	

func (b *HTMLBuilder) AddChild(childName, childText string) {
	b.root.elements = append(b.root.elements, &HTMLElement{childName, childText, []*HTMLElement{}})
}

func (b *HTMLBuilder) AddChildFluent(childName, childText string) *HTMLBuilder {
	b.AddChild(childName, childText)
	return b
}


func main() {
	sb := strings.Builder{}
	hello := "Hello"
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	works := []string{"a", "b", "c"}
	sb.Reset()
	// <ul><li>a</li><li>b</li><li>c</li></ul>
	sb.WriteString("<ul>")
	for _, w := range works {
		sb.WriteString("<li>")
		sb.WriteString(w)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")	
	fmt.Println(sb.String())
	b := NewHTMLBuilder("ul")
	b.AddChildFluent("li", "a").AddChildFluent("li", "b").AddChildFluent("li", "c").AddChildFluent("li", "d").AddChildFluent("li", "e")
	fmt.Println(b.String())

}

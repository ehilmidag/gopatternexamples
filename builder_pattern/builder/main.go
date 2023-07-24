package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElementType struct {
	elementName string
	text        string
	elements    []HtmlElementType
}

func (e *HtmlElementType) String() string {
	return e.string(0)
}

func (e *HtmlElementType) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.elementName))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n",
		i, e.elementName))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElementType
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName,
		HtmlElementType{
			elementName: rootName,
			text:        "",
			elements:    []HtmlElementType{},
		},
	}
}
func (b *HtmlBuilder) String() string {
	return b.root.String()
}
func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElementType{childName, childText, []HtmlElementType{}}
	b.root.elements = append(b.root.elements, e)
}
func main() {
	hello := "hello user"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	// her seferinde görüldüğü gibi amelelik yapıyoruz  sürekli <ul> <li> leri elle yazıyoruz
	//bunları structa toplasak her şey çok daha rahat olurdu
	words := []string{"hello", "stalker"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, s := range words {
		sb.WriteString("<li>")
		sb.WriteString(s)
		sb.WriteString("</li>")

	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	b := NewHtmlBuilder("ul")
	b.AddChild("li", "holy")
	b.AddChild("li", "moly")
	fmt.Println(b.String())

}

package main

import (
	"fmt"
	"strings"
)

type Visitor interface {
	VisitPlainText(*PlainText)
	VisitBoldText(*BoldText)
}

type Element interface {
	Accept(Visitor)
}

type PlainText struct {
	Text string
}

func (pt *PlainText) Accept(v Visitor) {
	v.VisitPlainText(pt)
}

type BoldText struct {
	Text string
}

func (bt *BoldText) Accept(v Visitor) {
	v.VisitBoldText(bt)
}

type WordCountVisitor struct {
	WordCount int
}

func (wcv *WordCountVisitor) VisitPlainText(pt *PlainText) {
	wcv.WordCount += len(strings.Fields(pt.Text))
}

func (wcv *WordCountVisitor) VisitBoldText(bt *BoldText) {
	wcv.WordCount += len(strings.Fields(bt.Text))
}

type Document struct {
	Elements []Element
}

func (d *Document) Add(e Element) {
	d.Elements = append(d.Elements, e)
}

func (d *Document) Accept(v Visitor) {
	for _, e := range d.Elements {
		e.Accept(v)
	}
}

func main() {
	doc := &Document{}
	doc.Add(&PlainText{Text: "This is a simple text."})
	doc.Add(&BoldText{Text: "This is a bold text."})
	wcv := &WordCountVisitor{}
	doc.Accept(wcv)
	fmt.Printf("Total word count: %d\n", wcv.WordCount)
}

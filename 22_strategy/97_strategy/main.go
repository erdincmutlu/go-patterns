package main

import (
	"fmt"
	"strings"
)

// Print a list of textual items. Printing using different formats

type OutputFormat int

const (
	Markdown OutputFormat = 1
	Html     OutputFormat = 2
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

// * one
// * two

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {
}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {
}

func (m *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" * " + item + "\n")
}

type HtmlListStrategy struct{}

func (h *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("/<ul>\n")
}

func (h *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("  <li>" + item + "</li>\n")
}

type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{strings.Builder{}, listStrategy}
}

func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
	switch fmt {
	case Markdown:
		t.listStrategy = &MarkdownListStrategy{}
	case Html:
		t.listStrategy = &HtmlListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	s := t.listStrategy
	s.Start(&t.builder)
	for _, item := range items {
		s.AddListItem(&t.builder, item)
	}
	s.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

func main() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)

	tp.Reset()
	tp.SetOutputFormat(Html)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)
}

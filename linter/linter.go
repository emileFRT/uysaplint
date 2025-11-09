package linter

import (
	"bytes"
	"strings"

	"mvdan.cc/sh/v3/syntax"
)

type Violation struct {
	Line, Col int
	Rule      string
	Msg       string
	Severity  string
	Fixed     bool
}

type Linter struct {
	Content    string
	File       *syntax.File
	Violations []Violation
	Disabled   map[string]bool
	Modified   bool

	Checkers map[string]func(syntax.Node)      // rule -> check function
	Fixers   map[string]func(syntax.Node) bool // rule -> fix function
}

func NewLinter(content string, disabled []string) (*Linter, error) {
	d := map[string]bool{}
	for _, r := range disabled {
		d[r] = true
	}

	parser := syntax.NewParser(syntax.KeepComments(true), syntax.Variant(syntax.LangBash))
	file, err := parser.Parse(strings.NewReader(content), "")
	if err != nil {
		return nil, err
	}

	l := &Linter{
		Content:    content,
		File:       file,
		Disabled:   d,
		Violations: make([]Violation, 0),
	}

	return l, nil
}

func (l *Linter) AddViolation(pos syntax.Pos, rule, msg, severity string, fixed bool) {
	if l.Disabled[rule] {
		return
	}
	l.Violations = append(l.Violations, Violation{
		Line: int(pos.Line()), Col: int(pos.Col()),
		Rule: rule, Msg: msg, Severity: severity, Fixed: fixed,
	})
}

func (l *Linter) GetContent() (string, error) {
	if !l.Modified {
		return l.Content, nil
	}
	var buf bytes.Buffer
	printer := syntax.NewPrinter()
	if err := printer.Print(&buf, l.File); err != nil {
		return "", err
	}
	return buf.String(), nil
}

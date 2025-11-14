package linter

import (
	"mvdan.cc/sh/v3/syntax"
)

type Violation struct {
	Line, Col int
	Rule      string
	Msg       string
	Severity  string
	Fixed     bool
}

type Linter interface {
	AddLintRule(name string, check Checker)
	AddFix(ruleName string, fix Fixer)
	DeleteRule(ruleName string)

	AddViolation(pos syntax.Pos, rule, msg, severity string, fixed bool)
	GetViolations() []Violation
	SetContent(string)
	GetContent() string

	Lint(modify bool)
}

type Checker func(Linter, syntax.Node)
type Fixer func(Linter, syntax.Node)

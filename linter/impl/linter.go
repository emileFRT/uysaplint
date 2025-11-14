package impl

import (
	"fmt"
	"strings"

	"github.com/emileFRT/ysaplint/linter"
	"mvdan.cc/sh/v3/syntax"
)

type LinterImpl struct {
	content    string
	file       *syntax.File
	violations []linter.Violation
	disabled   map[string]struct{}

	checkers map[string]func(syntax.Node) // rule -> check function
	fixers   map[string]func(syntax.Node) // rule -> fix function
}

func NewLinter(content string, disabled []string) (linter.Linter, error) {
	d := map[string]struct{}{}
	for _, r := range disabled {
		d[r] = struct{}{}
	}

	parser := syntax.NewParser(syntax.KeepComments(true), syntax.Variant(syntax.LangBash))
	file, err := parser.Parse(strings.NewReader(content), "")
	if err != nil {
		return nil, err
	}

	l := &LinterImpl{
		content:    content,
		violations: make([]linter.Violation, 0),

		file:     file, // TODO: remove?
		disabled: d,
		checkers: make(map[string]func(syntax.Node), 0),
		fixers:   make(map[string]func(syntax.Node), 0),
	}

	return l, nil
}

func (l *LinterImpl) AddViolation(pos syntax.Pos, rule, msg, severity string, fixed bool) {
	if _, disabled := l.disabled[rule]; disabled {
		return
	}
	l.violations = append(l.violations, linter.Violation{
		Line: int(pos.Line()), Col: int(pos.Col()),
		Rule: rule, Msg: msg, Severity: severity, Fixed: fixed,
	})
}

func (l *LinterImpl) GetViolations() []linter.Violation {
	return l.violations
}

func (l *LinterImpl) AddLintRule(name string, check linter.Checker) {
	l.checkers[name] = func(n syntax.Node) { check(l, n) }
}

func (l *LinterImpl) AddFix(ruleName string, fix linter.Fixer) {
	l.fixers[ruleName] = func(n syntax.Node) { fix(l, n) }
}

func (l *LinterImpl) GetContent() string {
	return l.content
}

func (l *LinterImpl) SetContent(c string) {
	var err error

	l.content = c
	parser := syntax.NewParser(syntax.KeepComments(true), syntax.Variant(syntax.LangBash))
	l.file, err = parser.Parse(strings.NewReader(l.content), "")
	if err != nil {
		panic(fmt.Errorf("Tried to set script content into a unparsable one, please check your fix: %v", err)) // a fix break bash
	}
}

func (l *LinterImpl) DeleteRule(ruleName string) {
	delete(l.checkers, ruleName)
	delete(l.fixers, ruleName)
}

package ysaplint

import (
	"github.com/emileFRT/ysaplint/linter"
	"github.com/emileFRT/ysaplint/linter/rules"
)

func registerRules(l linter.Linter) {
	for name, check := range rules.Checkers {
		l.AddLintRule(name, check)
	}
	for name, fix := range rules.Fixers {
		l.AddFix(name, fix)
	}
}

func FormatAll(l linter.Linter) {
	registerRules(l)

	for _, fix := range rules.NonWalkFixers {
		fix(l)
	}

	l.Lint(true)
}

func LintAll(l linter.Linter) {
	registerRules(l)

	for _, check := range rules.NonWalkChecker {
		check(l)
	}

	l.Lint(false)
}

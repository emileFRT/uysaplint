package rules

import (
	"unofficial-ysap-fmt/linter"

	"mvdan.cc/sh/v3/syntax"
)

// Backticks
func CheckBackticks(l *linter.Linter, node syntax.Node) {
	cs, ok := node.(*syntax.CmdSubst)
	if ok && cs.Backquotes {
		pos := cs.Pos()
		l.AddViolation(pos,
			RuleBackticks,
			"Use $(...) for command substitution instead of backticks",
			"error",
			false,
		)
	}
}

func FixBackticks(l *linter.Linter, node syntax.Node) bool {
	cs, ok := node.(*syntax.CmdSubst)
	if !ok || !cs.Backquotes {
		return false
	}
	cs.Backquotes = false
	l.AddViolation(cs.Pos(), RuleBackticks, "Converted backticks to $()", "error", true)
	return true
}

package rules

import (
	"unofficial-ysap-fmt/linter"

	"mvdan.cc/sh/v3/syntax"
)

func CheckTestCmd(l *linter.Linter, node syntax.Node) {
	ce, ok := node.(*syntax.CallExpr)
	if ok && len(ce.Args) > 0 && isLit(ce.Args[0], "[") {
		l.AddViolation(ce.Pos(), RuleTestCmd, "Use [[ ... ]] for conditional testing, not [ ... ] or test", "error", false)
	}
}

package rules

import (
	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

func CheckLet(l linter.Linter, node syntax.Node) {
	ce, ok := node.(*syntax.CallExpr)
	if ok && len(ce.Args) >= 2 && isLit(ce.Args[0], "let") {
		l.AddViolation(ce.Pos(), RuleLet, "Use ((...)) or $((...)) instead of let command", "error", false)
	}
}

package rules

import (
	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

func CheckNoEval(l linter.Linter, node syntax.Node) {
	ce, ok := node.(*syntax.CallExpr)
	if ok && len(ce.Args) > 0 && isLit(ce.Args[0], "eval") {
		l.AddViolation(ce.Pos(), RuleNoEval, "Never use eval - it's unsafe. Refactor using arrays or indirect expansion", "error", false)
	}
}

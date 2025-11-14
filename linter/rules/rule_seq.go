package rules

import (
	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

// Seq Usage
func CheckSeq(l linter.Linter, node syntax.Node) {
	ce, ok := node.(*syntax.CallExpr)
	if ok && len(ce.Args) > 0 && isLit(ce.Args[0], "seq") {
		l.AddViolation(ce.Pos(), RuleSeq, "Use bash builtins for sequences: {1..5} or ((i=0; i<n; i++))", "error", false)
	}
}

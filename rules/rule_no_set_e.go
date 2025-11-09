package rules

import (
	"strings"

	"unofficial-ysap-fmt/linter"

	"mvdan.cc/sh/v3/syntax"
)

func CheckNoSetE(l *linter.Linter, node syntax.Node) {
	ce, ok := node.(*syntax.CallExpr)
	if !ok || len(ce.Args) < 2 || !isLit(ce.Args[0], "set") {
		return
	}
	for _, arg := range ce.Args[1:] {
		if val := getLitVal(arg); strings.HasPrefix(val, "-") && strings.Contains(val, "e") {
			l.AddViolation(arg.Pos(), RuleNoSetE, "Avoid 'set -e' - handle errors explicitly. Remove if safe", "warning", false)
			return
		}
	}
}

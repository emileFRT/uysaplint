package rules

import (
	"unofficial-ysap-fmt/linter"

	"mvdan.cc/sh/v3/syntax"
)

func CheckParsingLs(l *linter.Linter, node syntax.Node) {
	fc, ok := node.(*syntax.ForClause)
	if !ok {
		return
	}

	// Check if it's a word iteration loop (for x in ...)
	wordIter, ok := fc.Loop.(*syntax.WordIter)
	if !ok || len(wordIter.Items) == 0 {
		return
	}

	for _, item := range wordIter.Items {
		if hasCmdSubst(item, "ls") {
			l.AddViolation(fc.Pos(), RuleParsingLs, "Never parse ls output. Use 'for f in *; do' instead", "error", false)
			return
		}
	}
}

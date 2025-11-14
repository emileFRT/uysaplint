package rules

import (
	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

func CheckBlockStmt(l linter.Linter, node syntax.Node) {
	switch n := node.(type) {
	case *syntax.IfClause:
		if hasNewline(n.Cond, n.ThenPos) {
			l.AddViolation(n.ThenPos, RuleBlockStmt, "'then' should be on same line as 'if': if condition; then", "error", false)
		}
	case *syntax.WhileClause:
		if hasNewline(n.Cond, n.DoPos) {
			l.AddViolation(n.DoPos, RuleBlockStmt, "'do' should be on same line as 'while': while condition; do", "error", false)
		}
	case *syntax.ForClause:
		if hasNewlineFor(n, n.DoPos) {
			l.AddViolation(n.DoPos, RuleBlockStmt, "'do' should be on same line as 'for': for ...; do", "error", false)
		}
	}
}

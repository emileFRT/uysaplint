package rules

import (
	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

// Semicolons
func CheckSemicolon(l linter.Linter, node syntax.Node) {
	stmt, ok := node.(*syntax.Stmt)
	if !ok || !stmt.Semicolon.IsValid() || isControlStruct(stmt) {
		return
	}
	l.AddViolation(stmt.Semicolon, RuleSemicolon, "Avoid semicolons unless required in control statements", "warning", false)
}

func FixSemicolon(l linter.Linter, node syntax.Node) {
	stmt, ok := node.(*syntax.Stmt)
	if !ok || !stmt.Semicolon.IsValid() || isControlStruct(stmt) {
		return
	}
	stmt.Semicolon = syntax.Pos{}
	l.AddViolation(stmt.Pos(), RuleSemicolon, "Removed unnecessary semicolon", "warning", true)
}

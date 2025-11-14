package rules

import (
	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

// Function Keyword
func CheckFunctionKw(l linter.Linter, node syntax.Node) {
	fd, ok := node.(*syntax.FuncDecl)
	if ok && fd.RsrvWord {
		// Use function name position for more precise error location
		pos := fd.Pos()
		if fd.Name != nil {
			pos = fd.Name.Pos()
		}
		l.AddViolation(pos, RuleFunctionKw, "Don't use 'function' keyword, use 'name() { ... }' instead", "error", false)
	}
}

func FixFunctionKw(l linter.Linter, node syntax.Node) {
	fd, ok := node.(*syntax.FuncDecl)
	if !ok || !fd.RsrvWord {
		return
	}
	fd.RsrvWord = false
	l.AddViolation(fd.Pos(), RuleFunctionKw, "Converted to name() syntax", "error", true)
}

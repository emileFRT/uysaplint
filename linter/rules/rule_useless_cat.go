package rules

import (
	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

func CheckUselessCat(l linter.Linter, node syntax.Node) {
	bc, ok := node.(*syntax.BinaryCmd)
	if !ok || bc.Op != syntax.Pipe {
		return
	}
	// x is left statement
	if !ok || bc.X.Cmd == nil {
		return
	}

	leftCall, ok := bc.X.Cmd.(*syntax.CallExpr)
	if !ok || len(leftCall.Args) != 2 || !isLit(leftCall.Args[0], "cat") {
		return
	}

	l.AddViolation(leftCall.Pos(), RuleUselessCat, "Useless use of cat. Use '< file' or pass filename directly", "warning", false)
}

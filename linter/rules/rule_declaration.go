package rules

import (
	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

// Declarations
func CheckDeclaration(l linter.Linter, node syntax.Node) {
	dc, ok := node.(*syntax.DeclClause)
	if !ok {
		return
	}
	if dc.Variant.Value == "readonly" {
		l.AddViolation(dc.Pos(), RuleDeclaration, "Don't use 'readonly'. Use simple assignment", "warning", false)
	} else if dc.Variant.Value == "declare" {
		for _, arg := range dc.Args {
			if isLit(arg, "-i") {
				l.AddViolation(dc.Pos(), RuleDeclaration, "Don't use 'declare -i'. Use simple assignment", "warning", false)
				return
			}
		}
	}
}

func FixDeclaration(l linter.Linter, node syntax.Node) {
	dc, ok := node.(*syntax.DeclClause)
	if !ok || dc.Variant.Value != "declare" {
		return
	}
	// TODO: size preshot
	var newArgs []*syntax.Assign
	modified := false
	for _, arg := range dc.Args {
		if !isLit(arg, "-i") {
			newArgs = append(newArgs, arg)
		} else {
			modified = true
		}
	}
	if !modified {
		return
	}
	dc.Args = newArgs
	l.AddViolation(dc.Pos(), RuleDeclaration, "Removed 'declare -i'", "warning", true)
}

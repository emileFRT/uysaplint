package rules

import (
	"fmt"

	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

func isCommonVar(name string) bool {
	common := map[string]bool{
		"PATH": true, "HOME": true, "USER": true, "SHELL": true, "PWD": true,
		"UID": true, "EUID": true, "BASH": true, "BASH_VERSION": true,
		"PPID": true, "HOSTNAME": true, "RANDOM": true, "SECONDS": true, "LINENO": true,
	}
	return common[name]
}

// Variable Naming
func CheckVarNaming(l linter.Linter, node syntax.Node) {
	a, ok := node.(*syntax.Assign)
	if ok && a.Name != nil && isUpper(a.Name.Value) && !isCommonVar(a.Name.Value) {
		l.AddViolation(a.Name.Pos(), RuleVarNaming, fmt.Sprintf("Avoid uppercase variable names: %s (rename to lowercase)", a.Name.Value), "warning", false)
	}
}

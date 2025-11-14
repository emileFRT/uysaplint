package rules

import (
	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

// TODO: fix that fct
func isQuotedWord(w *syntax.Word) bool {
	// Conservative: only flag clearly unquoted cases
	return false
}

// Unquoted Variables
func CheckUnquotedVar(l linter.Linter, node syntax.Node) {
	w, ok := node.(*syntax.Word)
	if !ok || isQuotedWord(w) || isSimpleLiteral(w) {
		return
	}
	for _, part := range w.Parts {
		if pe, ok := part.(*syntax.ParamExp); ok && shouldQuoteParam(pe) {
			l.AddViolation(pe.Pos(), RuleUnquotedVar, "Variable should be quoted: \"$var\"", "warning", false)
			return
		}
	}
}

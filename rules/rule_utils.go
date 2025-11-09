package rules

import "mvdan.cc/sh/v3/syntax"

func isControlStruct(stmt *syntax.Stmt) bool {
	switch stmt.Cmd.(type) {
	case *syntax.IfClause, *syntax.WhileClause, *syntax.ForClause:
		return true
	}
	return false
}

func isLit(n syntax.Node, val string) bool {
	w, ok := n.(*syntax.Word)
	if !ok || len(w.Parts) != 1 {
		return false
	}
	lit, ok := w.Parts[0].(*syntax.Lit)
	return ok && (val == "" || lit.Value == val)
}

func getLitVal(n syntax.Node) string {
	w, ok := n.(*syntax.Word)
	if !ok {
		return ""
	}
	for _, p := range w.Parts {
		if lit, ok := p.(*syntax.Lit); ok {
			return lit.Value
		}
	}
	return ""
}

func hasCmdSubst(w syntax.Node, cmd string) bool {
	word, ok := w.(*syntax.Word)
	if !ok {
		return false
	}
	for _, p := range word.Parts {
		cs, ok := p.(*syntax.CmdSubst)
		if !ok {
			continue
		}
		for _, stmt := range cs.Stmts {
			ce, ok := stmt.Cmd.(*syntax.CallExpr)
			if ok && len(ce.Args) > 0 && isLit(ce.Args[0], cmd) {
				return true
			}
		}
	}
	return false
}

func hasNewline(stmts []*syntax.Stmt, pos syntax.Pos) bool {
	return len(stmts) > 0 && pos.IsValid() && pos.Line() > stmts[len(stmts)-1].End().Line()
}
func hasNewlineFor(fc *syntax.ForClause, pos syntax.Pos) bool {
	if !pos.IsValid() {
		return false
	}

	// Check WordIter (for x in ...)
	if wordIter, ok := fc.Loop.(*syntax.WordIter); ok && len(wordIter.Items) > 0 {
		lastItem := wordIter.Items[len(wordIter.Items)-1]
		return pos.Line() > lastItem.End().Line()
	}

	// Check C-style for (for ((...;...;...)))
	if cFor, ok := fc.Loop.(*syntax.CStyleLoop); ok {
		if cFor.Cond != nil {
			return pos.Line() > cFor.Cond.End().Line()
		}
		return pos.Line() > cFor.Pos().Line()
	}

	// Default case
	return pos.Line() > fc.Pos().Line()
}

func isSimpleLiteral(w *syntax.Word) bool {
	return len(w.Parts) == 1 && isLit(w.Parts[0], "")
}

func shouldQuoteParam(pe *syntax.ParamExp) bool {
	return pe.Short && !pe.Length && pe.Param.Value != ""
}

func isUpper(s string) bool {
	if s == "" || s[0] < 'A' || s[0] > 'Z' {
		return false
	}
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			return false
		}
	}
	return true
}

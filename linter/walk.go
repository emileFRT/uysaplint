package linter

import "mvdan.cc/sh/v3/syntax"

func (l *Linter) WalkRules(rules []string, applyFixes bool) {
	syntax.Walk(l.File, func(node syntax.Node) bool {
		if node == nil {
			return true
		}
		for _, rule := range rules {
			if l.Disabled[rule] {
				continue
			}
			if applyFixes {
				if fix := l.Fixers[rule]; fix != nil && fix(node) {
					l.Modified = true
				}
			} else if check := l.Checkers[rule]; check != nil {
				check(node)
			}
		}
		return true
	})
}

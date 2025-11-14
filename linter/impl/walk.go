package impl

import (
	"mvdan.cc/sh/v3/syntax"
)

func (l *LinterImpl) Lint(applyFixes bool) {
	syntax.Walk(l.file, func(node syntax.Node) bool {
		if node == nil {
			return true
		}
		for rule, _ := range l.checkers {
			if applyFixes {
				if fix, ok := l.fixers[rule]; ok {
					fix(node)
				}
			} else if check, ok := l.checkers[rule]; ok {
				check(node)
			}
		}
		return true
	})
}

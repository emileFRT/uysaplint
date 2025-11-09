package unofficialysapfmt

import (
	"unofficial-ysap-fmt/linter"
	"unofficial-ysap-fmt/rules"
)

func FormatAll(l *linter.Linter) {
	l.WalkRules(rules.All, true)
}

func LintAll(l *linter.Linter) {
	l.WalkRules(rules.All, false)
}

package rules

import (
	"strings"

	"unofficial-ysap-fmt/linter"

	"mvdan.cc/sh/v3/syntax"
)

func CheckBlanklines(l *linter.Linter, node syntax.Node) {
	if l.Disabled["blanklines"] {
		return
	}
	lines := strings.Split(l.Content, "\n")
	consecutive := 0
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			consecutive++
			if consecutive > 1 {
				l.AddViolation(syntax.NewPos(1, uint(i), 1), RuleBlanklines, "No more than 1 blank line in a row", "warning", false)
			}
		} else {
			consecutive = 0
		}
	}
}

func FixBlankLines(l *linter.Linter) {
	if l.Disabled[RuleBlanklines] {
		return
	}
	lines := strings.Split(l.Content, "\n")
	var newLines []string
	consecutive := 0
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			consecutive++
			if consecutive <= 1 {
				newLines = append(newLines, line)
			} else {
				l.AddViolation(syntax.NewPos(1, uint(i), 1), RuleBlanklines, "Removed excess blank line", "warning", true)
				l.Modified = true
			}
		} else {
			consecutive = 0
			newLines = append(newLines, line)
		}
	}
	if l.Modified {
		l.Content = strings.Join(newLines, "\n")
		parser := syntax.NewParser(syntax.KeepComments(true), syntax.Variant(syntax.LangBash))
		l.File, _ = parser.Parse(strings.NewReader(l.Content), "")
	}
}

package rules

import (
	"strings"

	"github.com/emileFRT/ysaplint/linter"

	"mvdan.cc/sh/v3/syntax"
)

func CheckBlanklines(l linter.Linter) {
	lines := strings.Split(l.GetContent(), "\n")
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

func FixBlankLines(l linter.Linter) {
	var modified bool

	lines := strings.Split(l.GetContent(), "\n")
	var newLines []string
	consecutive := 0
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			consecutive++
			if consecutive <= 1 {
				newLines = append(newLines, line)
			} else {
				l.AddViolation(syntax.NewPos(1, uint(i), 1), RuleBlanklines, "Removed excess blank line", "warning", true)
				modified = true
			}
		} else {
			consecutive = 0
			newLines = append(newLines, line)
		}
	}
	if modified {
		l.SetContent(strings.Join(newLines, "\n"))
	}
}

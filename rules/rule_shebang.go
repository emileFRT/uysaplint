package rules

import (
	"strings"

	"unofficial-ysap-fmt/linter"

	"mvdan.cc/sh/v3/syntax"
)

func CheckShebang(l *linter.Linter, node syntax.Node) {
	if l.Disabled[RuleShebang] {
		return
	}
	lines := strings.Split(l.Content, "\n")
	if len(lines) > 0 && strings.HasPrefix(lines[0], "#!") {
		if !strings.Contains(lines[0], "#!/usr/bin/env bash") && !strings.Contains(lines[0], "#!/bin/bash") {
			l.AddViolation(syntax.Pos{}, RuleShebang, "Use '#!/usr/bin/env bash' for portability", "warning", false)
		}
	} else if len(lines) > 0 {
		l.AddViolation(syntax.Pos{}, RuleShebang, "Missing shebang, should be #!/usr/bin/env bash", "warning", false)
	}
}

func FixShebang(l *linter.Linter, node syntax.Node) bool {
	if l.Disabled[RuleShebang] {
		return false
	}
	lines := strings.Split(l.Content, "\n")
	if len(lines) == 0 {
		return false
	}
	modified := false
	if strings.HasPrefix(lines[0], "#!") {
		if !strings.Contains(lines[0], "#!/usr/bin/env bash") {
			lines[0] = "#!/usr/bin/env bash"
			modified = true
		}
	} else {
		lines = append([]string{"#!/usr/bin/env bash"}, lines...)
		modified = true
	}
	if modified {
		l.Content = strings.Join(lines, "\n")
		return true
	}
	return false
}

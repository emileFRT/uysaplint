package run

import (
	"fmt"
	"io"
	"os"

	"unofficial-ysap-fmt/linter"
)

func readInput(args []string) (string, string, error) {
	if len(args) == 0 {
		content, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", "", fmt.Errorf("error reading stdin: %v", err)
		}
		return string(content), "<stdin>", nil
	}
	content, err := os.ReadFile(args[0])
	if err != nil {
		return "", "", fmt.Errorf("error reading file: %v", err)
	}
	return string(content), args[0], nil
}

func printViolations(v []linter.Violation, filename string) bool {
	if len(v) == 0 {
		return false
	}
	fmt.Fprintln(os.Stderr, "[unofficial-ysap-fmt] Found violations:")
	hasErrors := false
	for _, v := range v {
		if !v.Fixed {
			fmt.Fprintf(os.Stderr, "  %s:%d:%d: [%s] %s (%s)\n",
				filename, v.Line, v.Col, v.Severity, v.Msg, v.Rule)
			if v.Severity == "error" {
				hasErrors = true
			}
		}
	}
	return hasErrors
}

package run

import (
	"fmt"
	"io"
	"os"

	"github.com/emileFRT/ysaplint/linter"
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

func hasError(v []linter.Violation) bool {
	hasErrors := false

	for _, v := range v {
		if !v.Fixed && v.Severity == "error" {
			hasErrors = true
		}
	}

	return hasErrors
}

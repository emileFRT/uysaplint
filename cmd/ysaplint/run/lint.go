package run

import (
	"fmt"
	"os"

	"github.com/emileFRT/ysaplint"
	"github.com/emileFRT/ysaplint/linter/impl"

	"github.com/spf13/cobra"
)

func Lint(cmd *cobra.Command, args []string) error {
	outFile := os.Stdout
	content, filename, err := readInput(args)
	if err != nil {
		return err
	}

	disabled, _ := cmd.Flags().GetStringSlice("disable")
	l, err := impl.NewLinter(content, disabled)
	if err != nil {
		return fmt.Errorf("parse error: %v", err)
	}

	ysaplint.LintAll(l)

	for i, v := range l.GetViolations() {
		if i == 0 {
			fmt.Fprintln(outFile, "[unofficial-ysap-fmt] Found violations:")
		}
		fmt.Fprintf(outFile, "  %s:%d:%d: [%s] %s (%s)\n",
			filename, v.Line, v.Col, v.Severity, v.Msg, v.Rule)
		if i == len(l.GetViolations())-1 {
			fmt.Fprintf(outFile, "\n%d violation(s)\n", len(l.GetViolations()))
		}
	}

	hasErrors := false
	for _, v := range l.GetViolations() {
		if v.Severity == "error" {
			hasErrors = true
			break
		}
	}
	if hasErrors {
		os.Exit(1)
	}
	return nil
}

package run

import (
	"fmt"
	"os"
	unofficialysapfmt "unofficial-ysap-fmt"

	"unofficial-ysap-fmt/linter"

	"github.com/spf13/cobra"
)

func Lint(cmd *cobra.Command, args []string) error {
	content, filename, err := readInput(args)
	if err != nil {
		return err
	}

	disabled, _ := cmd.Flags().GetStringSlice("disable")
	l, err := linter.NewLinter(content, disabled)
	if err != nil {
		return fmt.Errorf("parse error: %v", err)
	}

	unofficialysapfmt.LintAll(l)

	if len(l.Violations) > 0 {
		fmt.Fprintln(os.Stderr, "[unofficial-ysap-fmt] Found violations:")
		for _, v := range l.Violations {
			fmt.Fprintf(os.Stderr, "  %s:%d:%d: [%s] %s (%s)\n",
				filename, v.Line, v.Col, v.Severity, v.Msg, v.Rule)
		}
		fmt.Fprintf(os.Stderr, "\n%d violation(s)\n", len(l.Violations))
	}

	hasErrors := false
	for _, v := range l.Violations {
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

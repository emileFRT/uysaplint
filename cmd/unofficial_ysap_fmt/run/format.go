package run

import (
	"fmt"
	"os"
	unofficialysapfmt "unofficial-ysap-fmt"

	"unofficial-ysap-fmt/linter"

	"github.com/spf13/cobra"
)

func Format(cmd *cobra.Command, args []string) error {
	content, filename, err := readInput(args)
	if err != nil {
		return err
	}

	inPlace, _ := cmd.Flags().GetBool("inplace")
	useShfmt, _ := cmd.Flags().GetBool("shfmt")
	disabled, _ := cmd.Flags().GetStringSlice("disable")

	if inPlace && filename == "<stdin>" {
		return fmt.Errorf("cannot use -i flag with stdin")
	}

	l, err := linter.NewLinter(content, disabled)
	if err != nil {
		return fmt.Errorf("parse error: %v", err)
	}

	unofficialysapfmt.FormatAll(l)
	hasErrors := printViolations(l.Violations, filename)
	output, err := l.GetContent()
	if err != nil {
		return fmt.Errorf("print error: %v", err)
	}

	if useShfmt && !l.Disabled["shfmt"] {
		if formatted, err := runShfmt(output, cmd); err != nil {
			fmt.Fprintf(os.Stderr, "[shfmt] %v (continuing without shfmt)\n", err)
		} else {
			output = formatted
		}
	}

	if inPlace {
		return os.WriteFile(filename, []byte(output), 0644)
	}
	fmt.Print(output)

	if hasErrors {
		os.Exit(1)
	}
	return nil
}

package run

import (
	"fmt"
	"os"

	"github.com/emileFRT/ysaplint"
	"github.com/emileFRT/ysaplint/linter/impl"

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

	l, err := impl.NewLinter(content, disabled)
	if err != nil {
		return fmt.Errorf("parse error: %v", err)
	}

	ysaplint.FormatAll(l)
	hasErrors := hasError(l.GetViolations())
	if useShfmt {
		if formatted, err := runShfmt(l.GetContent(), cmd); err != nil {
			fmt.Fprintf(os.Stderr, "[shfmt] %v (continuing without shfmt)\n", err)
		} else {
			l.SetContent(formatted)
		}
	}

	if inPlace {
		return os.WriteFile(filename, []byte(l.GetContent()), 0644)
	}

	fmt.Print(l.GetContent())
	if hasErrors {
		os.Exit(1)
	}
	return nil
}

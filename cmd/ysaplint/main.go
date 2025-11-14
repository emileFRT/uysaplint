package main

import (
	"fmt"
	"os"

	"github.com/emileFRT/ysaplint/cmd/ysaplint/run"

	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt [file]",
	Short: "Best effort formatter, very limitted, also *call* shfmt",
	Long:  "Reads from stdin if no file specified. Outputs to stdout unless -i is used.",
	Args:  cobra.MaximumNArgs(1),
	RunE:  run.Format,
}

var lintCmd = &cobra.Command{
	Use:   "github.com/emileFRT/ysaplint [file]",
	Short: "Bash linter and formatter enforcing ysap.sh style guide",
	Long:  "Reads from stdin if no file specified. Outputs to stdout",
	Args:  cobra.MaximumNArgs(1),
	RunE:  run.Lint,
}

func init() {
	// TODO: add full flags passing, need version fixing too
	fmtCmd.Flags().BoolP("inplace", "i", false, "Edit file in place") // TODO: check behaviour on
	fmtCmd.Flags().BoolP("shfmt", "s", false, "Run shfmt formatter after our ysap fixes")
	fmtCmd.Flags().StringSliceP("disable", "d", []string{}, "Disable rules (comma-separated)")
	fmtCmd.Flags().Bool("binary-next-line", false, "[shfmt] Binary ops at start of next line")
	fmtCmd.Flags().Bool("no-switch-case-indent", false, "[shfmt] Don't indent switch cases")
	fmtCmd.Flags().Bool("no-space-redirects", false, "[shfmt] Don't space after redirects")
	fmtCmd.Flags().Bool("func-next-line", false, "[shfmt] Function brace on next line")

	lintCmd.Flags().StringSliceP("disable", "d", []string{}, "Disable rules (comma-separated)")
	lintCmd.AddCommand(fmtCmd)
}

func main() {
	if err := lintCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

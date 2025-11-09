package main

import (
	"fmt"
	"os"
	"unofficial-ysap-fmt/cmd/unofficial_ysap_fmt/run"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github.com/emileFRT/unofficial-ysap-fmt [file]",
	Short: "Bash linter and formatter enforcing ysap.sh style guide",
	Long:  "Reads from stdin if no file specified. Outputs to stdout unless -i is used.",
	Args:  cobra.MaximumNArgs(1),
	RunE:  run.Format,
}

var lintCmd = &cobra.Command{
	Use:   "lint [file]",
	Short: "Lint only (no fixes)",
	Args:  cobra.MaximumNArgs(1),
	RunE:  run.Lint,
}

func init() {
	rootCmd.Flags().BoolP("inplace", "i", false, "Edit file in place")
	rootCmd.Flags().BoolP("shfmt", "s", false, "Run shfmt formatter after fixes")
	rootCmd.Flags().StringSliceP("disable", "d", []string{}, "Disable rules (comma-separated)")
	rootCmd.Flags().Bool("binary-next-line", false, "[shfmt] Binary ops at start of next line")
	rootCmd.Flags().Bool("no-switch-case-indent", false, "[shfmt] Don't indent switch cases")
	rootCmd.Flags().Bool("no-space-redirects", false, "[shfmt] Don't space after redirects")
	rootCmd.Flags().Bool("func-next-line", false, "[shfmt] Function brace on next line")

	lintCmd.Flags().StringSliceP("disable", "d", []string{}, "Disable rules (comma-separated)")
	rootCmd.AddCommand(lintCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

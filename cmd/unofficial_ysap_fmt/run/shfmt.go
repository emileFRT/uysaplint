package run

import (
	"bytes"
	"strings"

	"github.com/spf13/cobra"
	"mvdan.cc/sh/v3/syntax"
)

type shfmtOptions struct {
	binaryNextLine, switchCaseIndent, spaceRedirects, funcNextLine bool
}

func runShfmt(content string, cmd *cobra.Command) (string, error) {
	opts := shfmtOptions{
		binaryNextLine:   getBool(cmd, "binary-next-line"),
		switchCaseIndent: !getBool(cmd, "no-switch-case-indent"),
		spaceRedirects:   !getBool(cmd, "no-space-redirects"),
		funcNextLine:     getBool(cmd, "func-next-line"),
	}

	parser := syntax.NewParser(syntax.KeepComments(true), syntax.Variant(syntax.LangBash))
	file, err := parser.Parse(strings.NewReader(content), "")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	printer := syntax.NewPrinter(
		syntax.BinaryNextLine(opts.binaryNextLine),
		syntax.SwitchCaseIndent(opts.switchCaseIndent),
		syntax.SpaceRedirects(opts.spaceRedirects),
		syntax.FunctionNextLine(opts.funcNextLine),
	)
	if err := printer.Print(&buf, file); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func getBool(cmd *cobra.Command, name string) bool {
	v, _ := cmd.Flags().GetBool(name)
	return v
}

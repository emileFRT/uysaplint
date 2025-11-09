package rules

import (
	"maps"
	"slices"
	"unofficial-ysap-fmt/linter"

	"mvdan.cc/sh/v3/syntax"
)

type Checker func(*linter.Linter, syntax.Node)
type Fixer func(*linter.Linter, syntax.Node) bool

var Checkers = map[string]Checker{
	RuleShebang:     CheckShebang,
	RuleSemicolon:   CheckSemicolon,
	RuleFunctionKw:  CheckFunctionKw,
	RuleTestCmd:     CheckTestCmd,
	RuleSeq:         CheckSeq,
	RuleBackticks:   CheckBackticks,
	RuleLet:         CheckLet,
	RuleParsingLs:   CheckParsingLs,
	RuleUnquotedVar: CheckUnquotedVar,
	RuleUselessCat:  CheckUselessCat,
	RuleNoEval:      CheckNoEval,
	RuleNoSetE:      CheckNoSetE,
	RuleBlockStmt:   CheckBlockStmt,
	RuleBlanklines:  CheckBlanklines,
	RuleVarNaming:   CheckVarNaming,
	RuleDeclaration: CheckDeclaration,
}

var Fixers = map[string]Fixer{
	RuleShebang:     FixShebang,
	RuleSemicolon:   FixSemicolon,
	RuleFunctionKw:  FixFunctionKw,
	RuleBackticks:   FixBackticks,
	RuleDeclaration: FixDeclaration,
}

var All = slices.Collect(maps.Keys(Checkers))

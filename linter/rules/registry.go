package rules

import (
	"github.com/emileFRT/ysaplint/linter"
)

var Checkers = map[string]linter.Checker{
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
	RuleVarNaming:   CheckVarNaming,
	RuleDeclaration: CheckDeclaration,
}

var Fixers = map[string]linter.Fixer{
	RuleSemicolon:   FixSemicolon,
	RuleFunctionKw:  FixFunctionKw,
	RuleBackticks:   FixBackticks,
	RuleDeclaration: FixDeclaration,
}

var NonWalkChecker = map[string]func(linter.Linter){
	RuleShebang:    CheckShebang,
	RuleBlanklines: CheckBlanklines,
}

var NonWalkFixers = map[string]func(linter.Linter){
	RuleShebang:    FixShebang,
	RuleBlanklines: FixBlankLines,
}

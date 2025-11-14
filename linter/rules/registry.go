package rules

import (
	"maps"
	"slices"

	"github.com/emileFRT/ysaplint/linter"
)

var Checkers = map[string]linter.Checker{
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

var Fixers = map[string]linter.Fixer{
	RuleShebang:     FixShebang,
	RuleSemicolon:   FixSemicolon,
	RuleFunctionKw:  FixFunctionKw,
	RuleBackticks:   FixBackticks,
	RuleDeclaration: FixDeclaration,
}

var All = slices.Collect(maps.Keys(Checkers))

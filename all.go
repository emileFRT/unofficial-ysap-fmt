package unofficialysapfmt

import (
	"github.com/emileFRT/unofficial-ysap-fmt/linter"
	"github.com/emileFRT/unofficial-ysap-fmt/rules"
)

func FormatAll(l *linter.Linter) {
	l.WalkRules(rules.All, true)
}

func LintAll(l *linter.Linter) {
	l.WalkRules(rules.All, false)
}

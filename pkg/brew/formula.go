package brew

import (
	"github.com/bjarn/sheepdog/pkg/command"
	"strings"
)

func FormulaIsInstalled(formula string) bool {
	out, err := command.Brew("list", "|", "grep", formula).Output()

	if err != nil {
		return false
	}

	return strings.Contains(string(out), formula)
}

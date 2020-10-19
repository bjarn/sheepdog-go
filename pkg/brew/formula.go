package brew

import (
	"github.com/bjarn/sheepdog/pkg/command"
	"strings"
)

func FormulaIsInstalled(formula string) bool {
	out, err := command.Brew("list", "--formula").Output()


	if err != nil {
		if !strings.Contains(err.Error(), "exit status 1") {
			return false
		}
	}

	return strings.Contains(string(out), formula)
}

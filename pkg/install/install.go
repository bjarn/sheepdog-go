package install

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

var qs = []*survey.Question{
	{
		Name: "domain",
		Prompt: &survey.Input{
			Message: "Enter a domain",
			Default: "test",
			Help:    "The domain will be used to access your projects. E.g. myproject.test",
		},
	},
	{
		Name: "database",
		Prompt: &survey.Select{
			Message: "Choose a database:",
			Options: []string{"mysql@8.0", "mysql@5.7", "mariadb"},
		},
	},
	{
		Name: "apps",
		Prompt: &survey.MultiSelect{
			Message: "Optional services:",
			Options: []string{"Elasticsearch", "Redis", "MailHog"},
		},
	},
}

func Run() {
	fmt.Printf("✨ Thanks for using Sheepdog! Let's get you started quickly.\n\n")

	answers := struct {
		Domain           string
		Database         string
		OptionalServices []string `survey:"apps"`
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		return
	}

	fmt.Printf("\n✨ Successfully installed Sheepdog! ✅\n")
}

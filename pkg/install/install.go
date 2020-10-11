package install

import (
	"fmt"
	"github.com/bjarn/sheepdog/internal/templates/stubs"
	"github.com/bjarn/sheepdog/utils"
	"os"
	"os/user"
	"text/template"

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
		Name: "optionalServices",
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
		OptionalServices []string `survey:"optionalServices"`
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		return
	}

	// Configure required services
	configureNginx()

	// Loop through the optional services and install them if they've been selected
	for _, optionalService := range answers.OptionalServices {
		if optionalService == "Elasticsearch" {
			configureElasticsearchNginxConf(answers.Domain)
		}
	}

	fmt.Printf("\n\n✨ Successfully installed Sheepdog! ✅\n")
}

func configureNginx() {
	fmt.Printf("\n👉 Configuring Nginx... ")
	file, err := os.Create("/usr/local/etc/nginx/nginx.conf")

	if err != nil {
		panic(err)
	}

	// Create other required nginx directories
	err = os.MkdirAll("/usr/local/etc/nginx/sheepdog/apps", 0755)

	if err != nil {
		panic(err)
	}

	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	data := struct {
		Username        string
		SheepdogHomeDir string
	}{
		currentUser.Username,
		utils.SheepdogHomeDir(),
	}

	tmpl := template.Must(template.New("nginx-tmpl").Parse(stubs.NginxConfTemplate))
	err = tmpl.Execute(file, data)

	_ = file.Chmod(0644)

	fmt.Print("Done")
}

// #################
// Configure apps
// #################

// Configure Elasticsearch Virtual Host Config
func configureElasticsearchNginxConf(domain string) {
	fmt.Printf("\n👉 Configuring Elasticsearch... ")
	file, err := os.Create("/usr/local/etc/nginx/sheepdog/apps/elasticsearch.conf")

	if err != nil {
		panic(err)
	}

	data := struct {
		Domain string
	}{
		domain,
	}

	tmpl := template.Must(template.New("elasticsearch-nginx-tmpl").Parse(stubs.ElasticsearchNginxConf))
	err = tmpl.Execute(file, data)

	_ = file.Chmod(0644)

	fmt.Print("Done")
}

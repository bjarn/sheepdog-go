package install

import (
	"fmt"
	"github.com/bjarn/sheepdog/internal/templates/stubs"
	"github.com/bjarn/sheepdog/pkg/command"
	"github.com/bjarn/sheepdog/utils"
	"os"
	"os/user"
	"strings"
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
		Name: "phpVersions",
		Prompt: &survey.MultiSelect{
			Message: "Choose one or more PHP versions:",
			Options: []string{"7.4", "7.3", "7.2", "7.1"},
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
			Options: []string{"Redis", "Elasticsearch", "MailHog"},
		},
	},
}

// Perform the survey on the user and install (when still needed) and configure the services included in
// the Sheepdog suite.
func Run() {
	fmt.Printf("✨ Thanks for using Sheepdog! Let's get you started quickly.\n\n")

	answers := struct {
		Domain           string
		PhpVersions      []string
		Database         string
		OptionalServices []string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		return
	}

	// Configure required services
	configureNginx()
	installPhpFpm(answers.PhpVersions)

	// Loop through the optional services and install them if they've been selected
	for _, optionalService := range answers.OptionalServices {
		if optionalService == "Elasticsearch" {
			configureElasticsearchNginxConf(answers.Domain)
		}
		if optionalService == "MailHog" {
			configureMailhogNginxConf(answers.Domain)
		}
	}

	fmt.Printf("\n\n✨ Successfully installed Sheepdog! ✅\n")
}

// Configure Nginx
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

func installPhpFpm(phpVersions []string) {
	fmt.Printf("\n👉 Installing php-fpm version(s): " + strings.Join(phpVersions, ", ") + "... ")
	for _, phpVersion := range phpVersions {
		err := command.Brew("install", "php@" + phpVersion).Run()
		if err != nil {
			// Brew throws exit status 1 as warning, just go on...
			if !strings.Contains(err.Error(), "exit status 1") {
				panic(err)
			}
		}
	}
	fmt.Print("Done")
}

// ##############################
// Configure optional services
// ##############################

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

// Configure MailHog Virtual Host Config
func configureMailHogNginxConf(domain string) {
	fmt.Printf("\n👉 Configuring MailHog... ")
	file, err := os.Create("/usr/local/etc/nginx/sheepdog/apps/mailhog.conf")

	if err != nil {
		panic(err)
	}

	data := struct {
		Domain string
	}{
		domain,
	}

	tmpl := template.Must(template.New("mailhog-nginx-tmpl").Parse(stubs.MailhogNginxConf))
	err = tmpl.Execute(file, data)

	_ = file.Chmod(0644)

	fmt.Print("Done")
}

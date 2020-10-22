package install

import (
	"fmt"
	"github.com/bjarn/sheepdog/internal/templates/stubs"
	"github.com/bjarn/sheepdog/pkg/brew"
	"github.com/bjarn/sheepdog/pkg/command"
	"github.com/bjarn/sheepdog/pkg/service"
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
			Options: []string{service.PhpFpm74.Name, service.PhpFpm73.Name, service.PhpFpm72.Name},
		},
		Validate: survey.Required,
	},
	{
		Name: "database",
		Prompt: &survey.Select{
			Message: "Choose a database:",
			Options: []string{service.MySql57.Name, service.MySql56.Name, service.MariaDb.Name},
		},
	},
	{
		Name: "optionalServices",
		Prompt: &survey.MultiSelect{
			Message: "Optional services:",
			Options: []string{service.Redis.Name, service.ElasticSearch.Name, service.MailHog.Name},
		},
	},
	{
		Name: "apps",
		Prompt: &survey.MultiSelect{
			Message: "Tools and apps:",
			Options: []string{"wp-cli", "magerun", "magerun2", "drush"},
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
		Apps             []string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		return
	}

	// Configure required services
	installDnsMasq()
	installNginx()
	configureNginx()
	installPhpFpm(answers.PhpVersions)
	installDatabase(answers.Database)

	// Loop through the optional services and install them if they've been selected
	for _, optionalService := range answers.OptionalServices {
		if optionalService == "Elasticsearch" {
			configureElasticsearchNginxConf(answers.Domain)
		}
		if optionalService == "MailHog" {
			configureMailHogNginxConf(answers.Domain)
		}
	}

	fmt.Printf("\n\n✨ Successfully installed Sheepdog! ✅\n")
}

// Install DnsMasq
func installDnsMasq() {
	fmt.Printf("\n👉 Installing DnsMasq... ")

	if brew.FormulaIsInstalled(service.DnsMasq.Name) {
		fmt.Printf("Already installed.\n")
		return
	}

	err := service.DnsMasq.Install()

	if err != nil {
		panic(err)
	}

	fmt.Print("Done")
}

// Install Nginx
func installNginx() {
	fmt.Printf("\n👉 Installing Nginx... ")

	if brew.FormulaIsInstalled("nginx") {
		fmt.Printf("Already installed.\n")
		return
	}

	err := service.Nginx.Install()

	if err != nil {
		panic(err)
	}

	fmt.Print("Done")
}

// Configure Nginx
func configureNginx() {
	fmt.Printf("\n👉 Configuring Nginx... ")
	file, err := os.Create(service.NginxPath + "/nginx.conf")

	if err != nil {
		panic(err)
	}

	// Create other required nginx directories
	err = os.MkdirAll(service.NginxPath+"/sheepdog/apps", 0755)

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

	err = service.Nginx.Restart()

	if err != nil {
		panic(err)
	}

	fmt.Print("Done")
}

// Install PHP-FPM versions selected by the user
func installPhpFpm(phpVersions []string) {
	fmt.Printf("\n👉 Installing php-fpm version(s): " + strings.Join(phpVersions, ", ") + "... ")

	for _, phpVersion := range phpVersions {
		if brew.FormulaIsInstalled("php@" + phpVersion) {
			fmt.Printf("Php " + phpVersion + " already is installed.\n")
			continue
		}

		var err error
		switch phpVersion {
		case service.PhpFpm72.Name:
			err = service.PhpFpm72.Install()
		case service.PhpFpm73.Name:
			err = service.PhpFpm73.Install()
		case service.PhpFpm74.Name:
			err = service.PhpFpm74.Install()
		}

		if err != nil {
			panic(err)
		}
	}

	fmt.Print("Done")
}

// Install the database service selected by the user
func installDatabase(database string) {
	fmt.Printf("\n👉 Installing database (" + database + ")... ")

	if brew.FormulaIsInstalled(database) {
		fmt.Printf("Database service " + database + " already is installed.\n")
		fmt.Print("Done")
		return
	}

	if database != service.MySql56.Name &&
		brew.FormulaIsInstalled(service.MySql56.Name) {
		uninstallDatabase(service.MySql56)
	}
	if database != service.MySql57.Name &&
		brew.FormulaIsInstalled(service.MySql57.Name) {
		uninstallDatabase(service.MySql57)
	}
	if database != service.MariaDb.Name &&
		brew.FormulaIsInstalled(service.MariaDb.Name) {
		uninstallDatabase(service.MariaDb)
	}

	var err error
	switch database {
	case service.MySql56.Name:
		err = service.MySql56.Install()
	case service.MySql57.Name:
		err = service.MySql57.Install()
	case service.MariaDb.Name:
		err = service.MariaDb.Install()
	}

	if err != nil {
		panic(err)
	}

	fmt.Print("Done")
}

func uninstallDatabase(s service.IService) {
	var database = s.(*service.Service)

	err := database.Stop()
	if err != nil {
		panic(err)
	}

	err = command.Brew("uninstall", database.Name).Run()
	if err != nil {
		// Brew throws exit status 1 as warning, just go on...
		if !strings.Contains(err.Error(), "exit status 1") {
			panic(err)
		}
	}

	fmt.Printf("\nService " + database.Name + " has been uninstalled.\n")
}

// ##############################
// Configure optional services
// ##############################

// Configure Elasticsearch Virtual Host Config
func configureElasticsearchNginxConf(domain string) {
	fmt.Printf("\n👉 Configuring Elasticsearch... ")
	file, err := os.Create(service.NginxPath + "/sheepdog/apps/elasticsearch.conf")

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
	file, err := os.Create(service.NginxPath + "/sheepdog/apps/mailhog.conf")

	if err != nil {
		panic(err)
	}

	data := struct {
		Domain string
	}{
		domain,
	}

	tmpl := template.Must(template.New("mailhog-nginx-tmpl").Parse(stubs.MailHogNginxConf))
	err = tmpl.Execute(file, data)

	_ = file.Chmod(0644)

	fmt.Print("Done")
}

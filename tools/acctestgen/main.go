package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	var dryRun = flag.Bool("dry-run", false, "preview list of files to be created")
	flag.Parse()

	service := flag.Arg(0)
	service = strings.ToLower(service)
	err := validate(service)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}

	datasourcePath := fmt.Sprintf("internal/services/%s/data_source_test.go", service)
	resourcePath := fmt.Sprintf("internal/services/%s/resource_test.go", service)
	testdataDirPath := fmt.Sprintf("internal/services/%s/testdata/", service)
	resourceDataPath := fmt.Sprintf("internal/services/%s/testdata/basic.tf", service)
	datasourceDataPath := fmt.Sprintf("internal/services/%s/testdata/datasource_basic.tf", service)

	// print preview of files/directories
	fmt.Println(fmt.Sprintf(`To be created:
- file:       %[1]s
- file:       %[2]s
- directory:  %[3]s
- file:       %[4]s
- file:       %[5]s
`, datasourcePath, resourcePath, testdataDirPath, resourceDataPath, datasourceDataPath))

	if *dryRun {
		fmt.Println(fmt.Sprintf(`
Dry-run mode is enabled.

To create the files, run this again without the 'dry-run' flag.

Example: go run cmd/acctest/main.go %s
`, service))
		os.Exit(0)
	}

	// get user confirmation before creating files
	var input string
	fmt.Print("Do you confirm that the file paths above are correct? [y|n]: ")
	_, err = fmt.Scanln(&input)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}
	input = strings.ToLower(input)
	if input != "y" && input != "yes" {
		fmt.Println("\nFile creation aborted")
		os.Exit(0)
	}

	names := formattedServiceNames{
		PascalCase: toPascalCase(service),
		SnakeCase:  service,
	}

	// create test data directory
	err = os.Mkdir(testdataDirPath, 0755)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}

	// write resource test data file
	f, err := os.Create(resourceDataPath)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}
	tmpl, err := template.New("basic.tf").Parse(resourceConfigTmpl())
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}
	err = tmpl.Execute(f, names)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}

	// write data source test data file
	f, err = os.Create(datasourceDataPath)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}
	tmpl, err = template.New("datasource_basic.tf").Parse(datasourceConfigTmpl())
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}
	err = tmpl.Execute(f, names)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}

	// write resource acceptance test file
	f, err = os.Create(resourcePath)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}
	tmpl, err = template.New("resource_test.go").Parse(resourceAccTestTmpl())
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}
	err = tmpl.Execute(f, names)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}

	// write data source acceptance test file
	f, err = os.Create(datasourcePath)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}
	tmpl, err = template.New("data_source_test.go").Parse(dataSourceAccTestTmpl())
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}
	err = tmpl.Execute(f, names)
	if err != nil {
		fmt.Printf("\nERROR: %s\n\n", err)
		os.Exit(1)
	}

	printNextStep(service)
}

// validate checks that the service provided exists, which implicitly
// checks that the input is in snake case, since that is what Stainless generates.
// To prevent overwriting existing test configurations and test cases, it also checks
// none of the files and directories that would be generated already exist.
func validate(service string) error {
	if service == "" || len(service) > 128 {
		return fmt.Errorf("a valid service name must be provided")
	}

	basePath := fmt.Sprintf("./internal/services/%s/", service)
	// check if service does not exist
	if _, err := os.Stat(basePath); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("service '%s' not found", service)
	}
	// check if any of the test files/configs already exist
	var errs []error
	if _, err := os.Stat(basePath + "resource_test.go"); err == nil {
		errs = append(errs, fmt.Errorf("the service '%s' already has a 'resource_test.go' file", service))
	}
	if _, err := os.Stat(basePath + "data_source_test.go"); err == nil {
		errs = append(errs, fmt.Errorf("the service '%s' already has a 'data_source_test.go' file", service))
	}
	if _, err := os.Stat(basePath + "testdata"); err == nil {
		errs = append(errs, fmt.Errorf("the service '%s' already has a 'testdata' directory", service))
	}

	return errors.Join(errs...)
}

func printNextStep(service string) {
	fmt.Println()
	fmt.Println(fmt.Sprintf(`
The files have been created. To run the acceptance tests, you will need to first
set the necessary environment variables. See https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4.39.0/contributing/environment-variable-dictionary.md
for the list of variables. The values will vary depending on the account and zone
being used for the tests. Once the environment variables are set, you can run the
tests using the command:

TF_ACC=1 go test ./internal/services/%s -run "^TestAccCloudflare" -v -count 1s`, service))
}

type formattedServiceNames struct {
	PascalCase string
	SnakeCase  string
}

func toPascalCase(s string) string {
	words := strings.Split(s, "_")
	caser := cases.Title(language.English)
	pascal := ""
	for _, word := range words {
		pascal += caser.String(word)
	}
	return pascal
}

func resourceConfigTmpl() string {
	return `resource "cloudflare_{{.SnakeCase}}" "%[1]s" {}`
}

func datasourceConfigTmpl() string {
	return `data "cloudflare_{{.SnakeCase}}" "%[1]s" {}`
}

func resourceAccTestTmpl() string {
	return `package {{.SnakeCase}}_test

import (
	"errors"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflare{{.PascalCase}}_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_{{.SnakeCase}}." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAcc{{.PascalCase}}Config(rnd),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						return errors.New("test not implemented")
					},
					resource.TestCheckResourceAttr(name, "some_string_attribute", "string_value"),
				),
			},
		},
	})
}

func testAcc{{.PascalCase}}Config(rnd string) string {
	return acctest.LoadTestCase("basic.tf", rnd)
}
`
}

func dataSourceAccTestTmpl() string {
	return `package {{.SnakeCase}}_test

import (
	"errors"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflare{{.PascalCase}}DataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_{{.SnakeCase}}." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAcc{{.PascalCase}}DataSourceConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						return errors.New("test not implemented")
					},
					resource.TestCheckResourceAttr(name, "some_string_attribute", "string_value"),
				),
			},
		},
	})
}

func testAcc{{.PascalCase}}DataSourceConfig(rnd string) string {
	return acctest.LoadTestCase("datasource_basic.tf", rnd)
}
`
}

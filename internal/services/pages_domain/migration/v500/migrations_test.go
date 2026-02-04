package v500_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zones"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// getTestZoneDomain returns a zone domain from the test account for use in tests
func getTestZoneDomain(t *testing.T, accountID string) string {
	client := acctest.SharedClient()
	zonePage, err := client.Zones.List(context.Background(), zones.ZoneListParams{
		Account: cloudflare.F(zones.ZoneListParamsAccount{
			ID: cloudflare.F(accountID),
		}),
	})
	if err != nil || zonePage == nil || len(zonePage.Result) == 0 {
		t.Skip("No zones available in account for testing")
	}
	return zonePage.Result[0].Name
}

// TestMigratePagesDomain_Basic tests basic migration from v4 to v5
// Migrates domain field to name field
func TestMigratePagesDomain_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	projectName := fmt.Sprintf("tf-acc-test-pages-project-%s", rnd)
	zoneDomain := getTestZoneDomain(t, accountID)
	domainName := fmt.Sprintf("%s.%s", rnd, zoneDomain)
	tmpDir := t.TempDir()

	// V4 config using cloudflare_pages_domain with domain field
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s_project" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"
}

resource "cloudflare_pages_domain" "%[1]s" {
  account_id   = "%[2]s"
  project_name = cloudflare_pages_project.%[1]s_project.name
  domain       = "%[4]s"
}`, rnd, accountID, projectName, domainName)

	resourceName := fmt.Sprintf("cloudflare_pages_domain.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: internal.LegacyProviderVersion,
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify field rename (domain → name)
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, internal.LegacyProviderVersion, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("project_name"), knownvalue.StringExact(projectName)),
				// Verify domain field was renamed to name
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(domainName)),
			}),
		},
	})
}

// TestMigratePagesDomain_WithVariables tests migration with variable references
func TestMigratePagesDomain_WithVariables(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	projectName := fmt.Sprintf("tf-acc-test-pages-project-%s", rnd)
	zoneDomain := getTestZoneDomain(t, accountID)
	domainName := fmt.Sprintf("%s.%s", rnd, zoneDomain)
	tmpDir := t.TempDir()

	// V4 config using variables
	v4Config := fmt.Sprintf(`
variable "account_id" {
  default = "%[2]s"
}

variable "project_name" {
  default = "%[3]s"
}

variable "domain_name" {
  default = "%[4]s"
}

resource "cloudflare_pages_project" "%[1]s_project" {
  account_id        = var.account_id
  name              = var.project_name
  production_branch = "main"
}

resource "cloudflare_pages_domain" "%[1]s" {
  account_id   = var.account_id
  project_name = cloudflare_pages_project.%[1]s_project.name
  domain       = var.domain_name
}`, rnd, accountID, projectName, domainName)

	resourceName := fmt.Sprintf("cloudflare_pages_domain.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: internal.LegacyProviderVersion,
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify variable references preserved
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, internal.LegacyProviderVersion, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("project_name"), knownvalue.StringExact(projectName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(domainName)),
			}),
		},
	})
}

// TestMigratePagesDomain_Subdomain tests migration with subdomain
func TestMigratePagesDomain_Subdomain(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	projectName := fmt.Sprintf("tf-acc-test-pages-project-%s", rnd)
	zoneDomain := getTestZoneDomain(t, accountID)
	domainName := fmt.Sprintf("blog.%s.%s", rnd, zoneDomain)
	tmpDir := t.TempDir()

	// V4 config with subdomain
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s_project" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"
}

resource "cloudflare_pages_domain" "%[1]s" {
  account_id   = "%[2]s"
  project_name = cloudflare_pages_project.%[1]s_project.name
  domain       = "%[4]s"
}`, rnd, accountID, projectName, domainName)

	resourceName := fmt.Sprintf("cloudflare_pages_domain.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: internal.LegacyProviderVersion,
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify subdomain preserved
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, internal.LegacyProviderVersion, "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("project_name"), knownvalue.StringExact(projectName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(domainName)),
			}),
		},
	})
}

// TestMigratePagesDomain_MultipleResources tests migration with multiple domains
func TestMigratePagesDomain_MultipleResources(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	projectName := fmt.Sprintf("tf-acc-test-pages-project-%s", rnd)
	zoneDomain := getTestZoneDomain(t, accountID)
	domain1 := fmt.Sprintf("%s-1.%s", rnd, zoneDomain)
	domain2 := fmt.Sprintf("%s-2.%s", rnd, zoneDomain)
	tmpDir := t.TempDir()

	// V4 config with multiple domains
	v4Config := fmt.Sprintf(`
resource "cloudflare_pages_project" "%[1]s_project" {
  account_id        = "%[2]s"
  name              = "%[3]s"
  production_branch = "main"
}

resource "cloudflare_pages_domain" "%[1]s_domain1" {
  account_id   = "%[2]s"
  project_name = cloudflare_pages_project.%[1]s_project.name
  domain       = "%[4]s"
}

resource "cloudflare_pages_domain" "%[1]s_domain2" {
  account_id   = "%[2]s"
  project_name = cloudflare_pages_project.%[1]s_project.name
  domain       = "%[5]s"
}`, rnd, accountID, projectName, domain1, domain2)

	resource1Name := fmt.Sprintf("cloudflare_pages_domain.%s_domain1", rnd)
	resource2Name := fmt.Sprintf("cloudflare_pages_domain.%s_domain2", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: internal.LegacyProviderVersion,
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify both domains migrated
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, internal.LegacyProviderVersion, "v4", "v5", []statecheck.StateCheck{
				// Verify first domain
				statecheck.ExpectKnownValue(resource1Name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resource1Name, tfjsonpath.New("project_name"), knownvalue.StringExact(projectName)),
				statecheck.ExpectKnownValue(resource1Name, tfjsonpath.New("name"), knownvalue.StringExact(domain1)),
				// Verify second domain
				statecheck.ExpectKnownValue(resource2Name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resource2Name, tfjsonpath.New("project_name"), knownvalue.StringExact(projectName)),
				statecheck.ExpectKnownValue(resource2Name, tfjsonpath.New("name"), knownvalue.StringExact(domain2)),
			}),
		},
	})
}

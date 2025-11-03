package workers_custom_domain_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// TestMigrateWorkersCustomDomain_Basic tests basic migration from v4 to v5
// This test verifies that:
// - cloudflare_worker_domain is renamed to cloudflare_workers_custom_domain
// - all attributes are preserved
func TestMigrateWorkersCustomDomain_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	hostname := rnd + "." + zoneName
	serviceName := "test-service-" + rnd
	resourceName := "cloudflare_workers_custom_domain." + rnd
	tmpDir := t.TempDir()

	// V4 config using cloudflare_worker_domain (old name)
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_domain" "%[1]s" {
  zone_id     = "%[2]s"
  account_id  = "%[3]s"
  hostname    = "%[4]s"
  service     = "%[5]s"
  environment = "production"
}`, rnd, zoneID, accountID, hostname, serviceName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service"), knownvalue.StringExact(serviceName)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("environment"), knownvalue.StringExact("production")),
			}),
		},
	})
}

// TestMigrateWorkersCustomDomain_WithoutEnvironment tests migration without environment attribute
func TestMigrateWorkersCustomDomain_WithoutEnvironment(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	hostname := rnd + "." + zoneName
	serviceName := "test-service-" + rnd
	resourceName := "cloudflare_workers_custom_domain." + rnd
	tmpDir := t.TempDir()

	// V4 config without environment (optional attribute)
	v4Config := fmt.Sprintf(`
resource "cloudflare_worker_domain" "%[1]s" {
  zone_id    = "%[2]s"
  account_id = "%[3]s"
  hostname   = "%[4]s"
  service    = "%[5]s"
}`, rnd, zoneID, accountID, hostname, serviceName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service"), knownvalue.StringExact(serviceName)),
			}),
		},
	})
}

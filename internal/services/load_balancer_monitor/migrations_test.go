package load_balancer_monitor_test

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

// TestMigrateLoadBalancerMonitor_Basic tests basic migration from v4 to v5
// This test verifies that:
// - Basic monitor attributes are preserved
func TestMigrateLoadBalancerMonitor_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor." + rnd
	tmpDir := t.TempDir()

	// V4 config with basic monitor
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type       = "http"
  expected_codes = "200"
  method     = "GET"
  timeout    = 5
  path       = "/health"
  interval   = 60
  retries    = 2
  description = "Test monitor %[1]s"
}`, rnd, accountID)

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
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify state
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("http")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("path"), knownvalue.StringExact("/health")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(fmt.Sprintf("Test monitor %s", rnd))),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitor_WithHeader tests migration with header transformation
// This test verifies that:
// - header array format is converted to map format
// - v4: header = [{"header": "Host", "values": ["example.com"]}]
// - v5: header = {"Host": ["example.com"]}
func TestMigrateLoadBalancerMonitor_WithHeader(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor." + rnd
	tmpDir := t.TempDir()

	// V4 config with header in array format
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type       = "http"
  expected_codes = "200"
  method     = "GET"
  path       = "/health"
  interval   = 60
  retries    = 2

  header {
    header = "Host"
    values = ["example.com"]
  }
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("http")),
				// Verify header was converted from array to map format
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("header"), knownvalue.MapPartial(map[string]knownvalue.Check{
					"Host": knownvalue.SetExact([]knownvalue.Check{
						knownvalue.StringExact("example.com"),
					}),
				})),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitor_MultipleHeaders tests migration with multiple headers
func TestMigrateLoadBalancerMonitor_MultipleHeaders(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple header blocks
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type       = "http"
  expected_codes = "200"
  method     = "GET"
  path       = "/health"
  interval   = 60
  retries    = 2

  header {
    header = "Host"
    values = ["example.com", "www.example.com"]
  }

  header {
    header = "X-Custom-Header"
    values = ["test-value"]
  }
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("http")),
				// Verify both headers were converted to map format
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("header"), knownvalue.MapSizeExact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("header").AtMapKey("Host"), knownvalue.SetSizeExact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("header").AtMapKey("X-Custom-Header"), knownvalue.SetExact([]knownvalue.Check{
					knownvalue.StringExact("test-value"),
				})),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitor_HTTPS tests migration with HTTPS type
func TestMigrateLoadBalancerMonitor_HTTPS(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor." + rnd
	tmpDir := t.TempDir()

	// V4 config with HTTPS and port
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type       = "https"
  expected_codes = "200"
  method     = "GET"
  path       = "/health"
  interval   = 60
  retries    = 2
  port       = 443
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("https")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("port"), knownvalue.Int64Exact(443)),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitor_TCP tests migration with TCP type
func TestMigrateLoadBalancerMonitor_TCP(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor." + rnd
	tmpDir := t.TempDir()

	// V4 config with TCP monitor
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type       = "tcp"
  interval   = 60
  retries    = 2
  timeout    = 5
  port       = 8080
}`, rnd, accountID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("tcp")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("port"), knownvalue.Int64Exact(8080)),
			}),
		},
	})
}

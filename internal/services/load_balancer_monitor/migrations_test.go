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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateLoadBalancerMonitorBasic tests basic HTTP monitor migration from v4 to v5
func TestMigrateLoadBalancerMonitorBasic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor." + rnd
	tmpDir := t.TempDir()

	// V4 config - basic HTTP monitor
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id     = "%[2]s"
  type           = "http"
  method         = "GET"
  path           = "/health"
  expected_codes = "2xx"
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
				// Verify resource exists (no type rename)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("http")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("path"), knownvalue.StringExact("/health")),
				// Verify defaults were added
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_insecure"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("follow_redirects"), knownvalue.Bool(false)),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitorWithHeaders tests header transformation (TypeSet → MapAttribute)
func TestMigrateLoadBalancerMonitorWithHeaders(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor." + rnd
	tmpDir := t.TempDir()

	// V4 config - monitor with headers (blocks)
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id     = "%[2]s"
  type           = "https"
  method         = "GET"
  path           = "/api/health"
  expected_codes = "2xx"

  header {
    header = "Host"
    values = ["api.example.com"]
  }

  header {
    header = "X-Custom-Header"
    values = ["CustomValue"]
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("https")),
				// Verify header map exists (transformed from blocks to map)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("header"), knownvalue.NotNull()),
				// Verify header values (map structure)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("header").AtMapKey("Host"), knownvalue.ListExact([]knownvalue.Check{
					knownvalue.StringExact("api.example.com"),
				})),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("header").AtMapKey("X-Custom-Header"), knownvalue.ListExact([]knownvalue.Check{
					knownvalue.StringExact("CustomValue"),
				})),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitorNumericConversions tests int → float64 conversions
func TestMigrateLoadBalancerMonitorNumericConversions(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor." + rnd
	tmpDir := t.TempDir()

	// V4 config - test all numeric fields
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id       = "%[2]s"
  type             = "http"
  expected_codes   = "2xx"
  interval         = 60
  retries          = 3
  timeout          = 5
  port             = 8080
  consecutive_down = 5
  consecutive_up   = 2
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
			// Step 2: Run migration and verify numeric conversions
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify all numeric fields are present (converted to float64)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("interval"), knownvalue.Float64Exact(60)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("retries"), knownvalue.Float64Exact(3)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("timeout"), knownvalue.Float64Exact(5)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("port"), knownvalue.Float64Exact(8080)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consecutive_down"), knownvalue.Float64Exact(5)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consecutive_up"), knownvalue.Float64Exact(2)),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitorMaximal tests migration with all optional fields
func TestMigrateLoadBalancerMonitorMaximal(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor." + rnd
	tmpDir := t.TempDir()

	// V4 config - all optional fields
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id         = "%[2]s"
  description        = "Test monitor for migration"
  type               = "https"
  method             = "GET"
  path               = "/status"
  port               = 8443
  interval           = 120
  retries            = 5
  timeout            = 7
  expected_codes     = "2xx,3xx"
  expected_body      = "healthy"
  follow_redirects   = true
  allow_insecure     = true
  consecutive_down   = 3
  consecutive_up     = 2

  header {
    header = "Host"
    values = ["status.example.com"]
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
				// Step 1: Create with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify all fields
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Test monitor for migration")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("https")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("path"), knownvalue.StringExact("/status")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expected_codes"), knownvalue.StringExact("2xx,3xx")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expected_body"), knownvalue.StringExact("healthy")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("follow_redirects"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("allow_insecure"), knownvalue.Bool(true)),
				// Numeric fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("port"), knownvalue.Float64Exact(8443)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("interval"), knownvalue.Float64Exact(120)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("retries"), knownvalue.Float64Exact(5)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("timeout"), knownvalue.Float64Exact(7)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consecutive_down"), knownvalue.Float64Exact(3)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consecutive_up"), knownvalue.Float64Exact(2)),
				// Header map
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("header").AtMapKey("Host"), knownvalue.ListExact([]knownvalue.Check{
					knownvalue.StringExact("status.example.com"),
				})),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitorTCP tests TCP monitor type (no headers)
func TestMigrateLoadBalancerMonitorTCP(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_load_balancer_monitor." + rnd
	tmpDir := t.TempDir()

	// V4 config - TCP monitor
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type       = "tcp"
  method     = "connection_established"
  port       = 8080
  interval   = 60
  retries    = 2
  timeout    = 5
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
			// Step 2: Run migration and verify TCP-specific fields
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("tcp")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("method"), knownvalue.StringExact("connection_established")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("port"), knownvalue.Float64Exact(8080)),
			}),
		},
	})
}

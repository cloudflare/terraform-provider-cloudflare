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

// TestMigrateLoadBalancerMonitorBasicHTTP tests migration of a simple HTTP monitor from v4 to v5
func TestMigrateLoadBalancerMonitorBasicHTTP(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id     = "%[2]s"
  type           = "http"
  method         = "GET"
  path           = "/health"
  interval       = 60
  retries        = 2
  timeout        = 5
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
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("http")),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("path"), knownvalue.StringExact("/health")),
				// Numeric fields should be converted to float64
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("interval"), knownvalue.Float64Exact(60)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("retries"), knownvalue.Float64Exact(2)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("timeout"), knownvalue.Float64Exact(5)),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitorWithHeaders tests migration of a monitor with headers from v4 to v5
func TestMigrateLoadBalancerMonitorWithHeaders(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config with header blocks
	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id     = "%[2]s"
  type           = "https"
  path           = "/api/status"
  expected_codes = "2xx"

  header {
    header = "Host"
    values = ["api.example.com"]
  }

  header {
    header = "Authorization"
    values = ["Bearer token123"]
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
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("https")),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("path"), knownvalue.StringExact("/api/status")),
				// Headers should be converted from array-of-objects to map-of-arrays
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("header").AtMapKey("Host"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("api.example.com")})),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("header").AtMapKey("Authorization"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("Bearer token123")})),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitorTCP tests migration of a TCP monitor with port
func TestMigrateLoadBalancerMonitorTCP(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type       = "tcp"
  port       = 3306
  timeout    = 10
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
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("tcp")),
				// Numeric fields should be converted to float64
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("port"), knownvalue.Float64Exact(3306)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("timeout"), knownvalue.Float64Exact(10)),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitorWithConsecutive tests migration of a monitor with consecutive checks
func TestMigrateLoadBalancerMonitorWithConsecutive(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id       = "%[2]s"
  type             = "http"
  expected_codes   = "2xx"
  consecutive_up   = 2
  consecutive_down = 3
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
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("http")),
				// Numeric fields should be converted to float64
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("consecutive_up"), knownvalue.Float64Exact(2)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("consecutive_down"), knownvalue.Float64Exact(3)),
			}),
		},
	})
}

// TestMigrateLoadBalancerMonitorFull tests migration of a monitor with all fields populated
func TestMigrateLoadBalancerMonitorFull(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id       = "%[2]s"
  description      = "Production HTTPS monitor"
  type             = "https"
  method           = "GET"
  path             = "/healthz"
  interval         = 60
  retries          = 2
  timeout          = 5
  expected_codes   = "2xx"
  expected_body    = "ok"
  allow_insecure   = true
  follow_redirects = true
  consecutive_up   = 2
  consecutive_down = 3

  header {
    header = "Host"
    values = ["health.example.com"]
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
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Production HTTPS monitor")),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("https")),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("method"), knownvalue.StringExact("GET")),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("path"), knownvalue.StringExact("/healthz")),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("expected_codes"), knownvalue.StringExact("2xx")),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("expected_body"), knownvalue.StringExact("ok")),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("allow_insecure"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("follow_redirects"), knownvalue.Bool(true)),
				// Numeric fields should be converted to float64
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("interval"), knownvalue.Float64Exact(60)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("retries"), knownvalue.Float64Exact(2)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("timeout"), knownvalue.Float64Exact(5)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("consecutive_up"), knownvalue.Float64Exact(2)),
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("consecutive_down"), knownvalue.Float64Exact(3)),
				// Headers should be converted from array-of-objects to map-of-arrays
				statecheck.ExpectKnownValue("cloudflare_load_balancer_monitor."+rnd, tfjsonpath.New("header").AtMapKey("Host"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("health.example.com")})),
			}),
		},
	})
}

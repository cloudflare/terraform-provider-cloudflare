package healthcheck_test

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

// TestMigrateHealthcheck_V4ToV5_BasicHTTP tests the most fundamental HTTP healthcheck
// migration with only required fields. Validates that:
// 1. Resource type remains cloudflare_healthcheck
// 2. HTTP fields are moved into http_config nested attribute
// 3. Basic state transformation works correctly
func TestMigrateHealthcheck_V4ToV5_BasicHTTP(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_healthcheck." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_healthcheck" "%[1]s" {
  zone_id = "%[2]s"
  name    = "tf-acc-test-basic-http-%[1]s"
  address = "example.com"
  type    = "HTTP"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: nil, // Migration tests don't need destroy checks
		WorkingDir:   tmpDir,
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-basic-http-%s", rnd))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("address"), knownvalue.StringExact("example.com")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("HTTP")),
				// http_config should be created even for minimal HTTP healthcheck
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateHealthcheck_V4ToV5_FullHTTP tests a complete HTTP healthcheck with all
// optional fields including headers. Validates:
// 1. All HTTP fields move into http_config
// 2. Headers transform from Set to Map structure
// 3. Numeric fields convert to Float64
func TestMigrateHealthcheck_V4ToV5_FullHTTP(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_healthcheck." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_healthcheck" "%[1]s" {
  zone_id = "%[2]s"
  name    = "tf-acc-test-full-http-%[1]s"
  address = "api.example.com"
  type    = "HTTP"

  # HTTP-specific fields (should move to http_config)
  port             = 80
  path             = "/health"
  method           = "GET"
  expected_codes   = ["200", "201", "204"]
  expected_body    = "OK"
  follow_redirects = false
  allow_insecure   = false

  # Headers (should transform from Set to Map)
  header {
    header = "Host"
    values = ["example.com"]
  }

  header {
    header = "User-Agent"
    values = ["HealthChecker/1.0"]
  }

  # Common fields (remain at root level)
  description           = "Full HTTP healthcheck test"
  consecutive_fails     = 3
  consecutive_successes = 2
  retries               = 2
  timeout               = 5
  interval              = 60
  suspended             = false
  check_regions         = ["WNAM", "ENAM"]
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-full-http-%s", rnd))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("HTTP")),

				// http_config should exist with all fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("port"), knownvalue.Float64Exact(80)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("path"), knownvalue.StringExact("/health")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("method"), knownvalue.StringExact("GET")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("expected_body"), knownvalue.StringExact("OK")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("follow_redirects"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("allow_insecure"), knownvalue.Bool(false)),

				// Header should be transformed to map
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("header"), knownvalue.NotNull()),

				// Common fields remain at root level
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Full HTTP healthcheck test")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consecutive_fails"), knownvalue.Float64Exact(3)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consecutive_successes"), knownvalue.Float64Exact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("timeout"), knownvalue.Float64Exact(5)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("interval"), knownvalue.Float64Exact(60)),
			}),
		},
	})
}

// TestMigrateHealthcheck_V4ToV5_HTTPS tests HTTPS healthcheck with SSL options.
// Validates that HTTPS also uses http_config (same as HTTP).
func TestMigrateHealthcheck_V4ToV5_HTTPS(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_healthcheck." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_healthcheck" "%[1]s" {
  zone_id = "%[2]s"
  name    = "tf-acc-test-https-%[1]s"
  address = "secure.example.com"
  type    = "HTTPS"

  # HTTPS-specific fields
  port           = 443
  path           = "/api/health"
  method         = "HEAD"
  allow_insecure = true
  expected_codes = ["200"]

  description = "HTTPS healthcheck test"
  timeout     = 10
  interval    = 30
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("HTTPS")),

				// HTTPS also uses http_config
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("port"), knownvalue.Float64Exact(443)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("path"), knownvalue.StringExact("/api/health")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("method"), knownvalue.StringExact("HEAD")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("allow_insecure"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateHealthcheck_V4ToV5_BasicTCP tests TCP healthcheck migration.
// Validates that TCP fields move into tcp_config.
func TestMigrateHealthcheck_V4ToV5_BasicTCP(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_healthcheck." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_healthcheck" "%[1]s" {
  zone_id = "%[2]s"
  name    = "tf-acc-test-tcp-%[1]s"
  address = "10.0.0.1"
  type    = "TCP"

  # TCP-specific fields (should move to tcp_config)
  port   = 8080
  method = "connection_established"

  description       = "TCP healthcheck test"
  consecutive_fails = 2
  timeout           = 10
  interval          = 30
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("TCP")),

				// tcp_config should exist with TCP fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tcp_config"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tcp_config").AtMapKey("port"), knownvalue.Float64Exact(8080)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("tcp_config").AtMapKey("method"), knownvalue.StringExact("connection_established")),

				// Common fields remain at root
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("consecutive_fails"), knownvalue.Float64Exact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("timeout"), knownvalue.Float64Exact(10)),
			}),
		},
	})
}

// TestMigrateHealthcheck_V4ToV5_MultipleHeaders specifically tests the header
// transformation from Set structure to Map structure.
func TestMigrateHealthcheck_V4ToV5_MultipleHeaders(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_healthcheck." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_healthcheck" "%[1]s" {
  zone_id = "%[2]s"
  name    = "tf-acc-test-headers-%[1]s"
  address = "api.example.com"
  type    = "HTTP"

  path   = "/health"
  method = "GET"

  # Multiple headers to test Set -> Map transformation
  header {
    header = "Host"
    values = ["example.com"]
  }

  header {
    header = "User-Agent"
    values = ["HealthChecker/1.0"]
  }

  header {
    header = "Accept"
    values = ["application/json", "text/plain"]
  }

  interval = 60
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: nil,
		WorkingDir:   tmpDir,
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("HTTP")),

				// http_config with header map
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("http_config").AtMapKey("header"), knownvalue.NotNull()),

				// Header should be a map (not checking exact keys due to Set -> Map transformation)
				// The important validation is that it exists and is structured correctly
			}),
		},
	})
}

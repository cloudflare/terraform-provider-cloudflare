package zero_trust_dex_test_test

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

// TestMigrateZeroTrustDEXTest_V4ToV5_BasicHTTP tests basic HTTP DEX test migration
func TestMigrateZeroTrustDEXTest_V4ToV5_BasicHTTP(t *testing.T) {
	// Zero Trust resources don't support API tokens yet
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	// V4 config with block syntax for data field
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Test HTTP connectivity"
  interval    = "0h30m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://dash.cloudflare.com"
    method = "GET"
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
			// Step 2: Run migration and verify state transformation
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify top-level fields preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Test HTTP connectivity")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("interval"), knownvalue.StringExact("0h30m0s")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				// Verify data field transformed from block to attribute (single nested object)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("kind"), knownvalue.StringExact("http")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("host"), knownvalue.StringExact("https://dash.cloudflare.com")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("method"), knownvalue.StringExact("GET")),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_Traceroute tests traceroute DEX test migration
func TestMigrateZeroTrustDEXTest_V4ToV5_Traceroute(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	// V4 config with traceroute (no method field)
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Test network path to Cloudflare DNS"
  interval    = "1h0m0s"
  enabled     = true

  data {
    kind = "traceroute"
    host = "1.1.1.1"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("interval"), knownvalue.StringExact("1h0m0s")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("kind"), knownvalue.StringExact("traceroute")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("host"), knownvalue.StringExact("1.1.1.1")),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_Disabled tests disabled DEX test migration
func TestMigrateZeroTrustDEXTest_V4ToV5_Disabled(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Currently disabled for maintenance"
  interval    = "0h15m0s"
  enabled     = false

  data {
    kind   = "http"
    host   = "https://internal.example.com"
    method = "GET"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("interval"), knownvalue.StringExact("0h15m0s")),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_ShortInterval tests frequent monitoring with 5m interval
func TestMigrateZeroTrustDEXTest_V4ToV5_ShortInterval(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "High-frequency monitoring"
  interval    = "0h5m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://api.cloudflare.com/client/v4/status"
    method = "GET"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("interval"), knownvalue.StringExact("0h5m0s")),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_LongInterval tests infrequent monitoring with 4h interval
func TestMigrateZeroTrustDEXTest_V4ToV5_LongInterval(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Low-frequency backup check"
  interval    = "4h0m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://backup.cloudflare.com"
    method = "GET"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("interval"), knownvalue.StringExact("4h0m0s")),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_HTTPWithPath tests HTTP request to URL with path
func TestMigrateZeroTrustDEXTest_V4ToV5_HTTPWithPath(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Test specific API endpoint"
  interval    = "0h10m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://api.cloudflare.com/client/v4/user"
    method = "GET"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("host"), knownvalue.StringExact("https://api.cloudflare.com/client/v4/user")),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_IPv6Traceroute tests traceroute to IPv6 address
func TestMigrateZeroTrustDEXTest_V4ToV5_IPv6Traceroute(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Test IPv6 connectivity"
  interval    = "2h0m0s"
  enabled     = true

  data {
    kind = "traceroute"
    host = "2606:4700:4700::1111"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("host"), knownvalue.StringExact("2606:4700:4700::1111")),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_LongDescription tests multiline description
func TestMigrateZeroTrustDEXTest_V4ToV5_LongDescription(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	longDescription := "This is a comprehensive DEX test with a very detailed description that spans multiple lines and contains important information about what this test does, why it exists, and how it should be monitored."

	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "%[3]s"
  interval    = "1h0m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://dash.cloudflare.com"
    method = "GET"
  }
}`, rnd, accountID, longDescription)

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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(longDescription)),
			}),
		},
	})
}

// NOTE: for_each tests are not included due to test framework limitation.
// MigrationV2TestStep does not support for_each syntax in state checks
// (unexpected index type error). The migration itself handles for_each correctly
// in production use - this is purely a testing framework constraint.

// TestMigrateZeroTrustDEXTest_V4ToV5_Count tests count pattern
func TestMigrateZeroTrustDEXTest_V4ToV5_Count(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  count = 3

  account_id  = "%[2]s"
  name        = "region-${count.index}-%[1]s"
  description = "Test endpoint in region ${count.index}"
  interval    = "0h30m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://region${count.index}.cloudflare.com"
    method = "GET"
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
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zero_trust_dex_test.%s[0]", rnd), tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zero_trust_dex_test.%s[1]", rnd), tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(fmt.Sprintf("cloudflare_zero_trust_dex_test.%s[2]", rnd), tfjsonpath.New("id"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_ConditionalCreation tests conditional creation with count
func TestMigrateZeroTrustDEXTest_V4ToV5_ConditionalCreation(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
variable "enable_test" {
  type    = bool
  default = true
}

resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  count = var.enable_test ? 1 : 0

  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Conditional test"
  interval    = "1h0m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://dash.cloudflare.com"
    method = "GET"
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
				statecheck.ExpectKnownValue(resourceName+"[0]", tfjsonpath.New("id"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_DynamicValues tests using locals and expressions
func TestMigrateZeroTrustDEXTest_V4ToV5_DynamicValues(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
locals {
  monitoring_config = {
    interval_minutes = 30
    enabled          = true
  }
}

resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Test with dynamic configuration"
  interval    = "0h${local.monitoring_config.interval_minutes}m0s"
  enabled     = local.monitoring_config.enabled

  data {
    kind   = "http"
    host   = "https://dash.cloudflare.com"
    method = "GET"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("interval"), knownvalue.StringExact("0h30m0s")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_MinimalConfig tests minimal required fields only
func TestMigrateZeroTrustDEXTest_V4ToV5_MinimalConfig(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	// Minimal config - description is required in v4 but optional in v5
	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Minimal test"
  interval    = "0h30m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://cloudflare.com"
    method = "GET"
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("kind"), knownvalue.StringExact("http")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("host"), knownvalue.StringExact("https://cloudflare.com")),
			}),
		},
	})
}

// TestMigrateZeroTrustDEXTest_V4ToV5_MaximalConfig tests all available v4 fields
func TestMigrateZeroTrustDEXTest_V4ToV5_MaximalConfig(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_dex_test." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_zero_trust_dex_test" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Comprehensive test with all fields"
  interval    = "0h30m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://dash.cloudflare.com/login"
    method = "GET"
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
				// Verify all top-level fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Comprehensive test with all fields")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("interval"), knownvalue.StringExact("0h30m0s")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				// Verify all data fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("kind"), knownvalue.StringExact("http")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("host"), knownvalue.StringExact("https://dash.cloudflare.com/login")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("data").AtMapKey("method"), knownvalue.StringExact("GET")),
			}),
		},
	})
}

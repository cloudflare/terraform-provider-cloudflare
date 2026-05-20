package zero_trust_gateway_policy_test

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

// TestMigrateZeroTrustGatewayPolicy_V4ToV5_Minimal tests migration of a minimal gateway policy
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_Minimal(t *testing.T) {
	// Zero Trust resources don't support API token authentication yet
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_rule" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-test-minimal-%[1]s"
  description = "Minimal policy for migration testing"
  precedence  = 10000
  action      = "block"
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] in {\"example.com\"})"
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
			acctest.MigrationV2TestStepForGatewayPolicy(t, v4Config, tmpDir, "4.52.1", "v4", "v5", false, []statecheck.StateCheck{
				// Resource should be renamed to cloudflare_zero_trust_gateway_policy
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-test-minimal-%s", rnd))),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Minimal policy for migration testing")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("action"), knownvalue.StringExact("block")),
				// Precedence is auto-calculated by API, just verify it exists and is float64
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateZeroTrustGatewayPolicy_V4ToV5_WithRuleSettings tests migration with rule_settings and field renames
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_WithRuleSettings(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_rule" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-test-settings-%[1]s"
  description = "Policy with rule settings"
  precedence  = 10000
  action      = "block"
  enabled     = true
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] in {\"badsite.com\"})"

  rule_settings {
    block_page_enabled = true
    block_page_reason  = "Access blocked by company policy"
    ip_categories      = true
    add_headers        = {}
    override_ips       = []
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
			acctest.MigrationV2TestStepForGatewayPolicy(t, v4Config, tmpDir, "4.52.1", "v4", "v5", true, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("action"), knownvalue.StringExact("block")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				// Precedence is auto-calculated by API
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.NotNull()),
				// Rule settings should be converted from block to attribute
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("block_page_enabled"), knownvalue.Bool(true)),
				// Field rename: block_page_reason → block_reason
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("block_reason"), knownvalue.StringExact("Access blocked by company policy")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("ip_categories"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateZeroTrustGatewayPolicy_V4ToV5_WithNestedBlocks tests migration with nested blocks
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_WithNestedBlocks(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_rule" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-test-nested-%[1]s"
  description = "Policy with nested blocks"
  precedence  = 10000
  action      = "block"
  enabled     = true
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] in {\"blocked.com\"})"

  rule_settings {
    block_page_enabled = true
    add_headers        = {}
    override_ips       = []

    notification_settings {
      enabled     = true
      message     = "Connection blocked"
      support_url = "https://support.example.com/"
    }
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
			acctest.MigrationV2TestStepForGatewayPolicy(t, v4Config, tmpDir, "4.52.1", "v4", "v5", true, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("action"), knownvalue.StringExact("block")),
				// Precedence is auto-calculated by API
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.NotNull()),
				// block_page_enabled should be present
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("block_page_enabled"), knownvalue.Bool(true)),
				// Nested notification_settings block should be converted to attribute
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("notification_settings").AtMapKey("enabled"), knownvalue.Bool(true)),
				// Field rename: message → msg
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("notification_settings").AtMapKey("msg"), knownvalue.StringExact("Connection blocked")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("notification_settings").AtMapKey("support_url"), knownvalue.StringExact("https://support.example.com/")),
			}),
		},
	})
}

// TestMigrateZeroTrustGatewayPolicy_V4ToV5_ComplexSettings tests migration with multiple nested structures
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_ComplexSettings(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_rule" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-test-complex-%[1]s"
  description = "Policy with complex settings"
  precedence  = 10000
  action      = "allow"
  enabled     = true
  filters     = ["http"]
  traffic     = "http.request.uri matches \".*api.*\""

  rule_settings {
    add_headers  = {}
    override_ips = []

    check_session {
      enforce  = true
      duration = "24h0m0s"
    }

    payload_log {
      enabled = true
    }
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
			acctest.MigrationV2TestStepForGatewayPolicy(t, v4Config, tmpDir, "4.52.1", "v4", "v5", true, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("action"), knownvalue.StringExact("allow")),
				// Precedence is auto-calculated by API
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("traffic"), knownvalue.StringExact("http.request.uri matches \".*api.*\"")),
				// All nested blocks should be converted to attributes
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("check_session").AtMapKey("enforce"), knownvalue.Bool(true)),
				// Duration should be "24h0m0s"
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("check_session").AtMapKey("duration"), knownvalue.StringExact("24h0m0s")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rule_settings").AtMapKey("payload_log").AtMapKey("enabled"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateZeroTrustGatewayPolicy_V4ToV5_EmptyRuleSettings tests migration without rule_settings
func TestMigrateZeroTrustGatewayPolicy_V4ToV5_EmptyRuleSettings(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_gateway_policy." + rnd
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_teams_rule" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-test-empty-%[1]s"
  description = "Policy without rule settings"
  precedence  = 10000
  action      = "block"
  enabled     = false
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] in {\"test.com\"})"
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
			acctest.MigrationV2TestStepForGatewayPolicy(t, v4Config, tmpDir, "4.52.1", "v4", "v5", false, []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("action"), knownvalue.StringExact("block")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
				// Precedence is auto-calculated by API
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("precedence"), knownvalue.NotNull()),
			}),
		},
	})
}

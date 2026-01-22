package zero_trust_device_posture_rule_test

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

// TestMigrateDevicePostureRuleBasic tests migration of a basic posture rule with input and match transformations 
func TestMigrateDevicePostureRuleBasic(t *testing.T) {
	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-posture-%s", rnd)
	tmpDir := t.TempDir()

	// Use v4 resource name in v4 config
	v4Config := fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type        = "os_version"
  description = "Device posture rule for corporate devices."
  schedule    = "24h"
  expiration  = "24h"

  match {
    platform = "linux"
  }

  input {
    version            = "1.0.0"
    operator           = "<"
    os_distro_name     = "ubuntu"
    os_distro_revision = "1.0.0"
    os_version_extra   = "(a)"
  }
}`, rnd, accountID, name)

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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("os_version")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("description"), knownvalue.StringExact("Device posture rule for corporate devices.")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("schedule"), knownvalue.StringExact("24h")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("expiration"), knownvalue.StringExact("24h")),

				// Match should be converted to array
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("linux")),

				// Input should be object with all fields
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("1.0.0")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact("<")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("os_distro_name"), knownvalue.StringExact("ubuntu")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("os_distro_revision"), knownvalue.StringExact("1.0.0")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("os_version_extra"), knownvalue.StringExact("(a)")),
			}),
		},
	})
}

// TestMigrateDevicePostureRuleWithoutName tests optional->required transformation for name field
// This test verifies that when name is missing from v4 config (and empty in API),
// the migration automatically populates it using the Terraform resource name as a fallback.
// Since v4 didn't set the name in the API, v5 will update it during apply.
func TestMigrateDevicePostureRuleWithoutName(t *testing.T) {
	// NOTE: This test passes because TransformConfigWithState automatically generates
	// a name (using the Terraform resource name) when state/API don't have one.

	// Zero Trust resources require API_KEY + EMAIL, not API_TOKEN
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// V4 config without name field (name was optional in v4, required in v5)
	// When v4 creates this resource without a name, the API stores it with empty name.
	// After migration, we auto-generate a name, and v5 will update the API resource.
	v4Config := fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  type       = "os_version"
  schedule   = "5m"

  match {
    platform = "linux"
  }

  input {
    version  = "10.0.0"
    operator = ">="
  }
}`, rnd, accountID)

	// Build test steps
	steps := []resource.TestStep{
		{
			// Step 1: Create with v4 provider (without name in config)
			ExternalProviders: map[string]resource.ExternalProvider{
				"cloudflare": {
					Source:            "cloudflare/cloudflare",
					VersionConstraint: "4.52.1",
				},
			},
			Config: v4Config,
		},
	}

	// Step 2: Migration adds name to config/state, v5 updates API with the name
	// A plan diff is expected because the API resource needs to be updated with the name
	migrationSteps := acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
		// Migration auto-populates name using Terraform resource name
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("os_version")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("schedule"), knownvalue.StringExact("5m")),

		// Match should be converted to array
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("linux")),

		// Input should be object
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("version"), knownvalue.StringExact("10.0.0")),
		statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("operator"), knownvalue.StringExact(">=")),
	})
	steps = append(steps, migrationSteps...)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		WorkingDir: tmpDir,
		Steps:      steps,
	})
}

// TestMigrateDevicePostureRuleFirewall tests migration of a firewall rule
func TestMigrateDevicePostureRuleFirewall(t *testing.T) {
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-firewall-%s", rnd)
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type       = "firewall"
  schedule   = "5m"

  match {
    platform = "windows"
  }

  input {
    enabled = true
  }
}`, rnd, accountID, name)

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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("firewall")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("schedule"), knownvalue.StringExact("5m")),

				// Match should be converted to array
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),

				// Input should be object with enabled field
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("enabled"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateDevicePostureRuleDiskEncryption tests Set->List conversion for check_disks
func TestMigrateDevicePostureRuleDiskEncryption(t *testing.T) {
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-disk-%s", rnd)
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type       = "disk_encryption"
  schedule   = "5m"

  match {
    platform = "windows"
  }

  input {
    check_disks = ["C:", "D:"]
    require_all = true
  }
}`, rnd, accountID, name)

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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("disk_encryption")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("schedule"), knownvalue.StringExact("5m")),

				// Match should be converted to array
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),

				// check_disks should be a list (Set->List conversion)
				// Note: Set ordering may vary, so we check size and presence
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("check_disks"), knownvalue.ListSizeExact(2)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("require_all"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateDevicePostureRuleMultiplePlatforms tests multiple match blocks->array conversion
func TestMigrateDevicePostureRuleMultiplePlatforms(t *testing.T) {
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-multi-%s", rnd)
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type       = "firewall"
  schedule   = "5m"

  match {
    platform = "windows"
  }

  match {
    platform = "mac"
  }

  match {
    platform = "linux"
  }

  input {
    enabled = true
  }
}`, rnd, accountID, name)

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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("firewall")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("schedule"), knownvalue.StringExact("5m")),

				// Match blocks should be converted to array with 3 items
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(0).AtMapKey("platform"), knownvalue.StringExact("windows")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(1).AtMapKey("platform"), knownvalue.StringExact("mac")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("match").AtSliceIndex(2).AtMapKey("platform"), knownvalue.StringExact("linux")),

				// Input should be object with enabled field
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("enabled"), knownvalue.Bool(true)),
			}),
		},
	})
}

// TestMigrateDevicePostureRuleFileType tests file type with multiple input attributes
func TestMigrateDevicePostureRuleFileType(t *testing.T) {
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-file-%s", rnd)
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type       = "file"
  schedule   = "5m"

  match {
    platform = "windows"
  }

  input {
    path       = "C:\\Program Files\\app.exe"
    exists     = true
    thumbprint = "0123456789abcdef"
    sha256     = "abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"
  }
}`, rnd, accountID, name)

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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("file")),

				// Input fields for file type
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("path"), knownvalue.StringExact("C:\\Program Files\\app.exe")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("exists"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("thumbprint"), knownvalue.StringExact("0123456789abcdef")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("sha256"), knownvalue.StringExact("abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890")),
			}),
		},
	})
}

// TestMigrateDevicePostureRuleDomainJoined tests domain_joined type
func TestMigrateDevicePostureRuleDomainJoined(t *testing.T) {
	originalToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if originalToken != "" {
		os.Unsetenv("CLOUDFLARE_API_TOKEN")
		defer os.Setenv("CLOUDFLARE_API_TOKEN", originalToken)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("tf-test-domain-%s", rnd)
	tmpDir := t.TempDir()

	v4Config := fmt.Sprintf(`
resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type       = "domain_joined"
  schedule   = "5m"

  match {
    platform = "windows"
  }

  input {
    domain = "example.com"
  }
}`, rnd, accountID, name)

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
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("type"), knownvalue.StringExact("domain_joined")),
				statecheck.ExpectKnownValue("cloudflare_zero_trust_device_posture_rule."+rnd, tfjsonpath.New("input").AtMapKey("domain"), knownvalue.StringExact("example.com")),
			}),
		},
	})
}

// Note: Tests for integration and certificate types (client_certificate, client_certificate_v2, sentinelone_s2s,
// tanium_s2s, kolide, crowdstrike_s2s, intune, workspace_one)
// are not included because they require valid connection_id values from real integrations.
// These would need to be tested manually with actual integration setups.

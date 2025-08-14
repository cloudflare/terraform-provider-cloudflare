package zone_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareZone_Migration(t *testing.T) {

	// Generate a unique zone name for this test
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	resourceName := fmt.Sprintf("cloudflare_zone.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	v4Config := testAccCloudflareZoneConfigV4(rnd, zoneName, accountID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create zone with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			{
				// Step 2: Migrate to v5 provider
				PreConfig: func() {
					// Run the migration command to transform config and state
					acctest.RunMigrationCommand(t, v4Config, tmpDir)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				// Verify no changes needed after migration
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("paused"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("full")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("plan").AtMapKey("legacy_id"), knownvalue.StringExact("free")),

					// Verify meta transformed correctly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta").AtMapKey("page_rule_quota"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta").AtMapKey("custom_certificate_quota"), knownvalue.NotNull()),
				},
			},
			{
				// Step 3 - make sure importing gives the same result as migrating
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateVerify:        true,
			},
		},
	})
}

func TestAccCloudflareZone_MigrationMinimalConfig(t *testing.T) {

	// Generate a unique zone name for this test
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	resourceName := fmt.Sprintf("cloudflare_zone.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	v4Config := testAccCloudflareZoneConfigV4Minimal(rnd, zoneName, accountID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create zone with v4 provider with minimal config (only required fields)
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			{
				// Step 2: Migrate to v5 provider
				PreConfig: func() {
					// Run the migration command to transform config and state
					acctest.RunMigrationCommand(t, v4Config, tmpDir)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				// Verify no changes needed after migration
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
					// Verify default values migrated correctly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("paused"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("full")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("plan").AtMapKey("legacy_id"), knownvalue.StringExact("free")),
					// Verify meta transformed correctly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta").AtMapKey("page_rule_quota"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta").AtMapKey("custom_certificate_quota"), knownvalue.NotNull()),
				},
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateVerify:        true,
			},
		},
	})
}

func TestAccCloudflareZone_MigrationEnterpriseWithVanityNameServers(t *testing.T) {

	// Generate a unique zone name for this test
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	resourceName := fmt.Sprintf("cloudflare_zone.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	v4Config := testAccCloudflareZoneConfigV4Enterprise(rnd, zoneName, accountID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create zone with v4 provider that has meta attributes
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			{
				// Step 2: Migrate to v5 provider
				PreConfig: func() {
					// Run the migration command to transform config and state
					acctest.RunMigrationCommand(t, v4Config, tmpDir)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				// Verify no updates needed after migration
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("paused"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("full")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("plan").AtMapKey("legacy_id"), knownvalue.StringExact("enterprise")),
					// Verify vanity name servers migrated correctly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("vanity_name_servers"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact(fmt.Sprintf("ns1.%s", zoneName)), knownvalue.StringExact(fmt.Sprintf("ns2.%s", zoneName))})),
					// Verify meta transformed correctly
					// Verify meta transformed correctly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta").AtMapKey("page_rule_quota"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta").AtMapKey("custom_certificate_quota"), knownvalue.NotNull()),
				},
			},
			{
				// Step 3 - make sure importing gives the same result as migrating
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateVerify:        true,
			},
		},
	})
}

func TestAccCloudflareZone_MigrationWithUnicode(t *testing.T) {

	// Generate a unique zone name for this test with Unicode characters
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("żółw-%s.cfapi.net", rnd)
	resourceName := fmt.Sprintf("cloudflare_zone.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	v4Config := testAccCloudflareZoneConfigV4Minimal(rnd, zoneName, accountID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create zone with v4 provider with Unicode domain name
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			{
				// Step 2: Migrate to v5 provider
				PreConfig: func() {
					// Run the migration command to transform config and state
					acctest.RunMigrationCommand(t, v4Config, tmpDir)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				// Step 3 - make sure importing gives the same result as migrating
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateVerify:        true,
			},
		},
	})
}

func TestAccCloudflareZone_MigrationWithPunycode(t *testing.T) {

	// TODO: Check with the service team about Punycode to Unicode conversion behavior
	// and update the migration guide accordingly. The v5 provider may deliberately handle
	// Punycode domains differently than v4.
	t.Skip("Skipping Punycode test - need to verify expected behavior with service team")

	// Generate a unique zone name for this test with Punycode
	rnd := utils.GenerateRandomResourceName()
	punycodeZoneName := fmt.Sprintf("xn--w-uga1v8h-%s.cfapi.net", rnd) // Punycode for żółw
	unicodeZoneName := fmt.Sprintf("żółw-%s.cfapi.net", rnd)           // Expected Unicode result
	resourceName := fmt.Sprintf("cloudflare_zone.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	v4Config := testAccCloudflareZoneConfigV4Minimal(rnd, punycodeZoneName, accountID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create zone with v4 provider with Punycode domain name
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// NOTE: PlanOnly: true followed by refresh or import test seems to catch issues that make "terraform show" choke,
			// but actually applying the plan does not catch those issues
			{
				// Step 2: Migrate to v5 provider
				PreConfig: func() {
					// Run the migration command to transform config and state
					acctest.RunMigrationCommand(t, v4Config, tmpDir)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					// V5 provider should maintain Unicode representation
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(unicodeZoneName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
					// Verify default values migrated correctly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("paused"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("full")),
					// Plan should be set to its default
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("plan").AtMapKey("legacy_id"), knownvalue.StringExact("free")),
					// Verify meta transformed correctly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta").AtMapKey("page_rule_quota"), knownvalue.NotNull()),
				},
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateVerify:        true,
			},
		},
	})
}

func TestAccCloudflareZone_MigrationPartialType(t *testing.T) {

	// Generate a unique zone name for this test
	rnd := utils.GenerateRandomResourceName()
	// Can't use whatever.cfapi.net for partial zones b/c base domain is already a full zone
	zoneName := fmt.Sprintf("%s-partial.net", rnd)
	resourceName := fmt.Sprintf("cloudflare_zone.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	v4Config := testAccCloudflareZoneConfigV4Type(rnd, zoneName, accountID, "partial")
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create zone with v4 provider with partial type
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			{
				// Step 2: Migrate to v5 provider
				PreConfig: func() {
					// Run the migration command to transform config and state
					acctest.RunMigrationCommand(t, v4Config, tmpDir)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("paused"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("partial")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("plan").AtMapKey("legacy_id"), knownvalue.StringExact("enterprise")),
				},
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateVerify:        true,
			},
		},
	})
}

func TestAccCloudflareZone_MigrationSecondaryType(t *testing.T) {

	// Generate a unique zone name for this test
	rnd := utils.GenerateRandomResourceName()
	zoneName := fmt.Sprintf("%s.cfapi.net", rnd)
	resourceName := fmt.Sprintf("cloudflare_zone.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	v4Config := testAccCloudflareZoneConfigV4Type(rnd, zoneName, accountID, "secondary")
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acctest.TestAccPreCheck(t); acctest.TestAccPreCheck_AccountID(t) },
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create zone with v4 provider with secondary type
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			{
				// Step 2: Migrate to v5 provider
				PreConfig: func() {
					// Run the migration command to transform config and state
					acctest.RunMigrationCommand(t, v4Config, tmpDir)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(tmpDir),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("paused"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("secondary")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("plan").AtMapKey("legacy_id"), knownvalue.StringExact("enterprise")),
					// Verify meta transformed correctly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("meta").AtMapKey("page_rule_quota"), knownvalue.NotNull()),
				},
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				ResourceName:             resourceName,
				ImportState:              true,
				ImportStateVerify:        true,
			},
		},
	})
}

func testAccCloudflareZoneConfigV4(rnd, zoneName, accountID string) string {
	return acctest.LoadTestCase("zoneconfigv4.tf", rnd, zoneName, accountID)
}

func testAccCloudflareZoneConfigV4Minimal(rnd, zoneName, accountID string) string {
	return acctest.LoadTestCase("zoneconfigv4minimal.tf", rnd, zoneName, accountID)
}

func testAccCloudflareZoneConfigV4Enterprise(rnd, zoneName, accountID string) string {
	return acctest.LoadTestCase("zoneconfigv4enterprise.tf", rnd, zoneName, accountID)
}

func testAccCloudflareZoneConfigV4Type(rnd, zoneName, accountID, zoneType string) string {
	return acctest.LoadTestCase("zoneconfigv4type.tf", rnd, zoneName, accountID, zoneType)
}

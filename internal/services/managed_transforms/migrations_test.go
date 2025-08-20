package managed_transforms_test

import (
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func testAccCheckCloudflareManagedTransformsDestroy(s *terraform.State) error {
	// Managed transforms are zone-wide settings that can't really be "destroyed"
	// They can only be disabled. This is a no-op check.
	return nil
}

func TestAccCloudflareManagedTransforms_Migration_RequestOnly(t *testing.T) {
	zoneID := acctest.TestAccCloudflareZoneID
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_managed_transforms." + rnd
	v4Config := testAccCloudflareManagedTransformsMigrationConfigV4RequestOnly(rnd, zoneID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: testAccCheckCloudflareManagedTransformsDestroy,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create managed headers with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_managed_headers."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue("cloudflare_managed_headers."+rnd, tfjsonpath.New("managed_request_headers"), knownvalue.SetSizeExact(1)),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("add_visitor_location_headers")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.ListSizeExact(0)),
			}),
		},
	},
	)
}

func TestAccCloudflareManagedTransforms_Migration_ResponseOnly(t *testing.T) {
	zoneID := acctest.TestAccCloudflareZoneID
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_managed_transforms." + rnd
	v4Config := testAccCloudflareManagedTransformsMigrationConfigV4ResponseOnly(rnd, zoneID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: testAccCheckCloudflareManagedTransformsDestroy,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create managed headers with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_managed_headers."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue("cloudflare_managed_headers."+rnd, tfjsonpath.New("managed_response_headers"), knownvalue.SetSizeExact(1)),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1",
				[]statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("remove_x-powered-by_header")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
				},
			),
		},
	})
}

func TestAccCloudflareManagedTransforms_Migration_Both(t *testing.T) {
	zoneID := acctest.TestAccCloudflareZoneID
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_managed_transforms." + rnd
	v4Config := testAccCloudflareManagedTransformsMigrationConfigV4Both(rnd, zoneID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: testAccCheckCloudflareManagedTransformsDestroy,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create managed headers with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_managed_headers."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue("cloudflare_managed_headers."+rnd, tfjsonpath.New("managed_request_headers"), knownvalue.SetSizeExact(2)),
					statecheck.ExpectKnownValue("cloudflare_managed_headers."+rnd, tfjsonpath.New("managed_response_headers"), knownvalue.SetSizeExact(1)),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.ListSizeExact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.ListSizeExact(1)),
			}),
		}},
	)
}

func TestAccCloudflareManagedTransforms_Migration_Empty(t *testing.T) {
	t.Skip("Skipping test - API behavior incompatibility: when no transforms are specified, " +
		"the API returns all available transforms with enabled:false rather than empty lists. " +
		"This creates unavoidable plan differences after migration.")

	// Test migration when v4 config has no managed_request_headers or managed_response_headers blocks
	// The state upgrade should add empty lists for these required attributes
	zoneID := acctest.TestAccCloudflareZoneID
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_managed_transforms." + rnd
	v4Config := testAccCloudflareManagedTransformsMigrationConfigV4Empty(rnd, zoneID)
	tmpDir := t.TempDir()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		CheckDestroy: testAccCheckCloudflareManagedTransformsDestroy,
		WorkingDir:   tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create managed headers with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						VersionConstraint: "4.52.1",
						Source:            "cloudflare/cloudflare",
					},
				},
				Config: v4Config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("cloudflare_managed_headers."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				},
			},
			// Step 2: Migrate to v5 provider
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.ListSizeExact(0)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.ListSizeExact(0)),
			}),
		},
	})
}

// V4 Configuration Functions

func testAccCloudflareManagedTransformsMigrationConfigV4RequestOnly(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_request_headers {
    id      = "add_visitor_location_headers"
    enabled = true
  }
}
`, rnd, zoneID)
}

func testAccCloudflareManagedTransformsMigrationConfigV4ResponseOnly(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_response_headers {
    id      = "remove_x-powered-by_header"
    enabled = true
  }
}
`, rnd, zoneID)
}

func testAccCloudflareManagedTransformsMigrationConfigV4Both(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_request_headers {
    id      = "add_visitor_location_headers"
    enabled = true
  }
  
  managed_request_headers {
    id      = "add_bot_protection_headers"
    enabled = false
  }
  
  managed_response_headers {
    id      = "add_security_headers"
    enabled = true
  }
}
`, rnd, zoneID)
}

func testAccCloudflareManagedTransformsMigrationConfigV4Empty(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
  
  # In v4, these blocks can be omitted, but v5 requires them as empty lists
  # The migration script should add empty lists for these
}
`, rnd, zoneID)
}

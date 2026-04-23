package certificate_pack_test

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

// TestMigrateCertificatePack_V4ToV5_Basic tests basic migration with minimal required fields
func TestMigrateCertificatePack_V4ToV5_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_certificate_pack." + rnd
	tmpDir := t.TempDir()

	// V4 config with minimal required fields
	v4Config := fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[1]s" {
	zone_id               = "%[2]s"
	type                  = "advanced"
	hosts                 = ["%[3]s", "*.%[3]s"]
	validation_method     = "txt"
	validity_days         = 90
	certificate_authority = "lets_encrypt"
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
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
				// Verify resource exists with correct type (no rename)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify required fields migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("lets_encrypt")),
				// Verify hosts array preserved
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(2)),
				// Verify computed fields present
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateCertificatePack_V4ToV5_AllFields tests migration with all optional fields
func TestMigrateCertificatePack_V4ToV5_AllFields(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_certificate_pack." + rnd
	tmpDir := t.TempDir()

	// V4 config with all optional fields (use TXT validation for wildcard support)
	v4Config := fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[1]s" {
	zone_id               = "%[2]s"
	type                  = "advanced"
	hosts                 = ["%[3]s", "*.%[3]s", "www.%[3]s"]
	validation_method     = "txt"
	validity_days         = 90
	certificate_authority = "google"
	cloudflare_branding   = false
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
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
				// Verify resource exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify all fields migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validation_method"), knownvalue.StringExact("txt")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("google")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cloudflare_branding"), knownvalue.Bool(false)),
				// Verify hosts array with multiple entries
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hosts"), knownvalue.ListSizeExact(3)),
				// Verify computed fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateCertificatePack_V4ToV5_WaitForActiveStatusRemoved tests that wait_for_active_status is removed
func TestMigrateCertificatePack_V4ToV5_WaitForActiveStatusRemoved(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_certificate_pack." + rnd
	tmpDir := t.TempDir()

	// V4 config with wait_for_active_status = false (removed in v5)
	// Using false to avoid timeout - the point is to test removal, not waiting
	v4Config := fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[1]s" {
	zone_id                = "%[2]s"
	type                   = "advanced"
	hosts                  = ["%[3]s", "*.%[3]s"]
	validation_method      = "txt"
	validity_days          = 90
	certificate_authority  = "lets_encrypt"
	wait_for_active_status = false
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify wait_for_active_status is removed
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify standard fields migrated
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("type"), knownvalue.StringExact("advanced")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
				// Note: wait_for_active_status should not be in state (it's removed from HCL)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
			}),
		},
	})
}

// TestMigrateCertificatePack_V4ToV5_GoogleCA tests migration with Google certificate authority
func TestMigrateCertificatePack_V4ToV5_GoogleCA(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_certificate_pack." + rnd
	tmpDir := t.TempDir()

	// V4 config with Google CA
	v4Config := fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[1]s" {
	zone_id               = "%[2]s"
	type                  = "advanced"
	hosts                 = ["%[3]s", "*.%[3]s"]
	validation_method     = "txt"
	validity_days         = 90
	certificate_authority = "google"
	cloudflare_branding   = false
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify Google CA
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify certificate authority
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("google")),
				// Verify other fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("cloudflare_branding"), knownvalue.Bool(false)),
			}),
		},
	})
}

// TestMigrateCertificatePack_V4ToV5_SSLComCA tests migration with SSL.com certificate authority
func TestMigrateCertificatePack_V4ToV5_SSLComCA(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_certificate_pack." + rnd
	tmpDir := t.TempDir()

	// V4 config with SSL.com CA and 30-day validity
	v4Config := fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[1]s" {
	zone_id               = "%[2]s"
	type                  = "advanced"
	hosts                 = ["%[3]s", "*.%[3]s"]
	validation_method     = "txt"
	validity_days         = 30
	certificate_authority = "ssl_com"
	cloudflare_branding   = false
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify SSL.com CA
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify certificate authority
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("ssl_com")),
				// Verify validity_days conversion (Int to Int64)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(30)),
				// Verify other fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
			}),
		},
	})
}

// TestMigrateCertificatePack_V4ToV5_HttpValidation tests migration with HTTP validation
func TestMigrateCertificatePack_V4ToV5_HttpValidation(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_certificate_pack." + rnd
	tmpDir := t.TempDir()

	// V4 config with HTTP validation (no wildcard - HTTP doesn't support wildcards)
	v4Config := fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[1]s" {
	zone_id               = "%[2]s"
	type                  = "advanced"
	hosts                 = ["%[3]s"]
	validation_method     = "http"
	validity_days         = 90
	certificate_authority = "lets_encrypt"
	cloudflare_branding   = false
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify HTTP validation
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify validation method
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validation_method"), knownvalue.StringExact("http")),
				// Verify other fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(90)),
			}),
		},
	})
}

// TestMigrateCertificatePack_V4ToV5_Validity14Days tests validity_days = 14
func TestMigrateCertificatePack_V4ToV5_Validity14Days(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_certificate_pack." + rnd
	tmpDir := t.TempDir()

	// V4 config with 14-day validity (requires Google or SSL.com CA)
	v4Config := fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[1]s" {
	zone_id               = "%[2]s"
	type                  = "advanced"
	hosts                 = ["%[3]s", "*.%[3]s"]
	validation_method     = "txt"
	validity_days         = 14
	certificate_authority = "google"
	cloudflare_branding   = false
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify Int to Int64 conversion
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify validity_days conversion (Int to Int64)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(14)),
				// Verify other fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("google")),
			}),
		},
	})
}

// TestMigrateCertificatePack_V4ToV5_Validity30Days tests validity_days = 30
func TestMigrateCertificatePack_V4ToV5_Validity30Days(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_certificate_pack." + rnd
	tmpDir := t.TempDir()

	// V4 config with 30-day validity
	v4Config := fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[1]s" {
	zone_id               = "%[2]s"
	type                  = "advanced"
	hosts                 = ["%[3]s", "*.%[3]s"]
	validation_method     = "txt"
	validity_days         = 30
	certificate_authority = "google"
	cloudflare_branding   = false
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify Int to Int64 conversion
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify validity_days conversion (Int to Int64)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(30)),
				// Verify other fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
			}),
		},
	})
}

// TestMigrateCertificatePack_V4ToV5_Validity365Days tests validity_days = 365
func TestMigrateCertificatePack_V4ToV5_Validity365Days(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_certificate_pack." + rnd
	tmpDir := t.TempDir()

	// V4 config with 365-day validity
	v4Config := fmt.Sprintf(`
resource "cloudflare_certificate_pack" "%[1]s" {
	zone_id               = "%[2]s"
	type                  = "advanced"
	hosts                 = ["%[3]s", "*.%[3]s"]
	validation_method     = "txt"
	validity_days         = 365
	certificate_authority = "ssl_com"
	cloudflare_branding   = false
}`, rnd, zoneID, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create resource with v4 provider
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "4.52.1",
					},
				},
				Config: v4Config,
			},
			// Step 2: Run migration and verify Int to Int64 conversion
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
				// Verify resource exists
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				// Verify validity_days conversion (Int to Int64) for max value
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("validity_days"), knownvalue.Int64Exact(365)),
				// Verify other fields
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate_authority"), knownvalue.StringExact("ssl_com")),
			}),
		},
	})
}

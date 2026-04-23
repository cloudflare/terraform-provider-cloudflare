package api_shield_test

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

// TestMigrateAPIShield_V4ToV5_SingleHeader tests migration of a single header-based auth characteristic
func TestMigrateAPIShield_V4ToV5_SingleHeader(t *testing.T) {
	// API Shield doesn't support API tokens yet
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	// V4 config using block syntax for auth_id_characteristics
	v4Config := fmt.Sprintf(`
resource "cloudflare_api_shield" "%[1]s" {
  zone_id = "%[2]s"

  auth_id_characteristics {
    type = "header"
    name = "authorization"
  }
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
				// Verify resource exists with correct zone_id
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Verify auth_id_characteristics array has 1 element
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(1)),
				// Verify the first characteristic has correct type and name
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("authorization")),
			}),
		},
	})
}

// TestMigrateAPIShield_V4ToV5_SingleCookie tests migration of a single cookie-based auth characteristic
func TestMigrateAPIShield_V4ToV5_SingleCookie(t *testing.T) {
	// API Shield doesn't support API tokens yet
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	// V4 config using block syntax with cookie type
	v4Config := fmt.Sprintf(`
resource "cloudflare_api_shield" "%[1]s" {
  zone_id = "%[2]s"

  auth_id_characteristics {
    type = "cookie"
    name = "session_id"
  }
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("cookie")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("session_id")),
			}),
		},
	})
}

// TestMigrateAPIShield_V4ToV5_MultipleCharacteristics tests migration with multiple auth characteristics
func TestMigrateAPIShield_V4ToV5_MultipleCharacteristics(t *testing.T) {
	// API Shield doesn't support API tokens yet
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple blocks (mixed types)
	v4Config := fmt.Sprintf(`
resource "cloudflare_api_shield" "%[1]s" {
  zone_id = "%[2]s"

  auth_id_characteristics {
    type = "header"
    name = "authorization"
  }

  auth_id_characteristics {
    type = "cookie"
    name = "session_id"
  }

  auth_id_characteristics {
    type = "header"
    name = "x-api-key"
  }
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(3)),
				// Verify first characteristic
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("authorization")),
				// Verify second characteristic
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(1).AtMapKey("type"), knownvalue.StringExact("cookie")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("session_id")),
				// Verify third characteristic
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("type"), knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("name"), knownvalue.StringExact("x-api-key")),
			}),
		},
	})
}

// TestMigrateAPIShield_V4ToV5_ComplexRealWorld tests migration with real-world OAuth scenario
func TestMigrateAPIShield_V4ToV5_ComplexRealWorld(t *testing.T) {
	// API Shield doesn't support API tokens yet
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	// V4 config simulating OAuth flow with multiple characteristics
	v4Config := fmt.Sprintf(`
resource "cloudflare_api_shield" "%[1]s" {
  zone_id = "%[2]s"

  auth_id_characteristics {
    type = "header"
    name = "Authorization"
  }

  auth_id_characteristics {
    type = "header"
    name = "X-OAuth-Token"
  }

  auth_id_characteristics {
    type = "cookie"
    name = "oauth_state"
  }

  auth_id_characteristics {
    type = "header"
    name = "X-Request-ID"
  }
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(4)),
				// Verify Authorization header (with capital A)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("Authorization")),
				// Verify X-OAuth-Token header (with hyphens)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(1).AtMapKey("type"), knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("X-OAuth-Token")),
				// Verify oauth_state cookie (with underscore)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("type"), knownvalue.StringExact("cookie")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("name"), knownvalue.StringExact("oauth_state")),
				// Verify X-Request-ID header
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(3).AtMapKey("type"), knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(3).AtMapKey("name"), knownvalue.StringExact("X-Request-ID")),
			}),
		},
	})
}

// TestMigrateAPIShield_V4ToV5_SpecialCharacters tests migration with various naming patterns
func TestMigrateAPIShield_V4ToV5_SpecialCharacters(t *testing.T) {
	// API Shield doesn't support API tokens yet
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_api_shield." + rnd
	tmpDir := t.TempDir()

	// V4 config with special characters in names
	v4Config := fmt.Sprintf(`
resource "cloudflare_api_shield" "%[1]s" {
  zone_id = "%[2]s"

  auth_id_characteristics {
    type = "header"
    name = "X-API-Key"
  }

  auth_id_characteristics {
    type = "header"
    name = "X_Custom_Header"
  }

  auth_id_characteristics {
    type = "cookie"
    name = "SessionID"
  }

  auth_id_characteristics {
    type = "cookie"
    name = "user-session-token"
  }
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics"), knownvalue.ListSizeExact(4)),
				// Verify X-API-Key (hyphens)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("X-API-Key")),
				// Verify X_Custom_Header (underscores)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(1).AtMapKey("type"), knownvalue.StringExact("header")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(1).AtMapKey("name"), knownvalue.StringExact("X_Custom_Header")),
				// Verify SessionID (mixed case)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("type"), knownvalue.StringExact("cookie")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(2).AtMapKey("name"), knownvalue.StringExact("SessionID")),
				// Verify user-session-token (hyphens in cookie)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(3).AtMapKey("type"), knownvalue.StringExact("cookie")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auth_id_characteristics").AtSliceIndex(3).AtMapKey("name"), knownvalue.StringExact("user-session-token")),
			}),
		},
	})
}

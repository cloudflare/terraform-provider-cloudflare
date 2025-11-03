package api_token_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestMigrateAPIToken_Basic tests basic migration from v4 to v5
// This test verifies that:
// - policy blocks are converted to policies list attribute
func TestMigrateAPIToken_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("terraform-test-%s", rnd)
	resourceName := "cloudflare_api_token." + rnd
	tmpDir := t.TempDir()

	// V4 config using policy blocks
	v4Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[2]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b",
      "c8fed203ed3043cba015a93ad1616f1f"
    ]
    resources = {
      "com.cloudflare.api.account.*" = "*"
    }
  }
}`, rnd, name)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				// Verify policy blocks were converted to policies list
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("permission_groups"), knownvalue.SetSizeExact(2)),
			}),
		},
	})
}

// TestMigrateAPIToken_MultiplePolicies tests migration with multiple policy blocks
func TestMigrateAPIToken_MultiplePolicies(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("terraform-test-%s", rnd)
	resourceName := "cloudflare_api_token." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple policy blocks
	v4Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[2]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b"
    ]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }

  policy {
    permission_groups = [
      "c8fed203ed3043cba015a93ad1616f1f"
    ]
    resources = {
      "com.cloudflare.api.account.*" = "*"
    }
  }
}`, rnd, name)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				// Verify both policy blocks were converted to policies list
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
			}),
		},
	})
}

// TestMigrateAPIToken_WithCondition tests migration with condition block
func TestMigrateAPIToken_WithCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("terraform-test-%s", rnd)
	resourceName := "cloudflare_api_token." + rnd
	tmpDir := t.TempDir()

	// V4 config with condition block containing request_ip
	v4Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[2]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b"
    ]
    resources = {
      "com.cloudflare.api.account.*" = "*"
    }
  }

  condition {
    request_ip {
      in     = ["192.0.2.0/24", "198.51.100.0/24"]
      not_in = ["192.0.2.1/32"]
    }
  }
}`, rnd, name)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
				// Verify condition block was converted to condition object
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition"), knownvalue.ObjectPartial(map[string]knownvalue.Check{
					"request_ip": knownvalue.NotNull(),
				})),
				// Verify request_ip is an object with in and not_in
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("in"), knownvalue.SetSizeExact(2)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("condition").AtMapKey("request_ip").AtMapKey("not_in"), knownvalue.SetSizeExact(1)),
			}),
		},
	})
}

// TestMigrateAPIToken_WithExpiresOn tests migration with expires_on attribute
func TestMigrateAPIToken_WithExpiresOn(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("terraform-test-%s", rnd)
	resourceName := "cloudflare_api_token." + rnd
	tmpDir := t.TempDir()

	// V4 config with expires_on
	v4Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name       = "%[2]s"
  expires_on = "2025-01-01T00:00:00Z"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b"
    ]
    resources = {
      "com.cloudflare.api.account.*" = "*"
    }
  }
}`, rnd, name)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.StringExact("2025-01-01T00:00:00Z")),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(1)),
			}),
		},
	})
}

// TestMigrateAPIToken_ComplexResources tests migration with complex resource patterns
func TestMigrateAPIToken_ComplexResources(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("terraform-test-%s", rnd)
	resourceName := "cloudflare_api_token." + rnd
	tmpDir := t.TempDir()

	// V4 config with complex resource specifications
	v4Config := fmt.Sprintf(`
resource "cloudflare_api_token" "%[1]s" {
  name = "%[2]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b",
      "c8fed203ed3043cba015a93ad1616f1f"
    ]
    resources = {
      "com.cloudflare.api.account.zone.*"        = "*"
      "com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4" = "*"
    }
  }

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b"
    ]
    resources = {
      "com.cloudflare.api.account.*" = "*"
    }
    effect = "deny"
  }
}`, rnd, name)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
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
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies"), knownvalue.ListSizeExact(2)),
				// Verify first policy has 2 resources
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(0).AtMapKey("resources"), knownvalue.MapSizeExact(2)),
				// Verify second policy has effect = deny
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("policies").AtSliceIndex(1).AtMapKey("effect"), knownvalue.StringExact("deny")),
			}),
		},
	})
}

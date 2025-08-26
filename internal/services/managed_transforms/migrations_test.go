package managed_transforms_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMigrateManagedHeadersToManagedTransformsFromV4(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_managed_transforms." + rnd
	tmpDir := t.TempDir()

	// V4 config using managed_headers resource
	v4Config := fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_request_headers {
    id      = "add_true_client_ip_headers"
    enabled = true
  }
  
  managed_request_headers {
    id      = "add_visitor_location_headers"
    enabled = true
  }
  
  managed_response_headers {
    id      = "add_security_headers"
    enabled = true
  }
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Verify request headers migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{
					knownvalue.ObjectExact(map[string]knownvalue.Check{
						"id":      knownvalue.StringExact("add_true_client_ip_headers"),
						"enabled": knownvalue.Bool(true),
					}),
					knownvalue.ObjectExact(map[string]knownvalue.Check{
						"id":      knownvalue.StringExact("add_visitor_location_headers"),
						"enabled": knownvalue.Bool(true),
					}),
				})),
				// Verify response headers migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{
					knownvalue.ObjectExact(map[string]knownvalue.Check{
						"id":      knownvalue.StringExact("add_security_headers"),
						"enabled": knownvalue.Bool(true),
					}),
				})),
			}),
		},
	})
}

func TestMigrateManagedHeadersToManagedTransformsFromV4OnlyRequestHeaders(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_managed_transforms." + rnd
	tmpDir := t.TempDir()

	// V4 config with only request headers (no response headers)
	v4Config := fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_request_headers {
    id      = "add_true_client_ip_headers"
    enabled = true
  }
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Verify request headers migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{
					knownvalue.ObjectExact(map[string]knownvalue.Check{
						"id":      knownvalue.StringExact("add_true_client_ip_headers"),
						"enabled": knownvalue.Bool(true),
					}),
				})),
				// Verify response headers exists as empty set (required in v5)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{})),
			}),
		},
	})
}

func TestMigrateManagedHeadersToManagedTransformsFromV4OnlyResponseHeaders(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_managed_transforms." + rnd
	tmpDir := t.TempDir()

	// V4 config with only response headers (no request headers)
	v4Config := fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_response_headers {
    id      = "add_security_headers"
    enabled = true
  }
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Verify request headers exists as empty set (required in v5)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{})),
				// Verify response headers migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{
					knownvalue.ObjectExact(map[string]knownvalue.Check{
						"id":      knownvalue.StringExact("add_security_headers"),
						"enabled": knownvalue.Bool(true),
					}),
				})),
			}),
		},
	})
}

func TestMigrateManagedHeadersToManagedTransformsFromV4EmptyHeaders(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_managed_transforms." + rnd
	tmpDir := t.TempDir()

	// V4 config with no headers (edge case - both will be empty)
	v4Config := fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Verify both headers exist as empty sets (required in v5)
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{})),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{})),
			}),
		},
	})
}

func TestMigrateManagedHeadersToManagedTransformsFromV4MultipleHeaders(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_managed_transforms." + rnd
	tmpDir := t.TempDir()

	// V4 config with multiple headers of each type
	// Note: v4 provider doesn't properly store headers with enabled=false, so we only test enabled headers
	v4Config := fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_request_headers {
    id      = "add_true_client_ip_headers"
    enabled = true
  }
  
  managed_request_headers {
    id      = "add_visitor_location_headers"
    enabled = true
  }
  
  managed_response_headers {
    id      = "add_security_headers"
    enabled = true
  }
}`, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
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
			acctest.MigrationTestStep(t, v4Config, tmpDir, "4.52.1", []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
				// Verify all request headers migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{
					knownvalue.ObjectExact(map[string]knownvalue.Check{
						"id":      knownvalue.StringExact("add_true_client_ip_headers"),
						"enabled": knownvalue.Bool(true),
					}),
					knownvalue.ObjectExact(map[string]knownvalue.Check{
						"id":      knownvalue.StringExact("add_visitor_location_headers"),
						"enabled": knownvalue.Bool(true),
					}),
				})),
				// Verify all response headers migrated correctly
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{
					knownvalue.ObjectExact(map[string]knownvalue.Check{
						"id":      knownvalue.StringExact("add_security_headers"),
						"enabled": knownvalue.Bool(true),
					}),
				})),
			}),
		},
	})
}
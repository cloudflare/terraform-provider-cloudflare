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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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
			acctest.MigrationV2TestStep(t, v4Config, tmpDir, "4.52.1", "v4", "v5", []statecheck.StateCheck{
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

func TestMigrateManagedTransformsMultiVersion(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	// Test cases for different provider versions
	testCases := []struct {
		name          string
		version       string
		configFunc    func(rnd, zoneID string) string
		skipPlanCheck bool // Some v5 versions have known issues with refresh plan
	}{
		{
			name:       "from_v4_52_1",
			version:    "4.52.1",
			configFunc: managedTransformsConfigV4,
		},
		{
			name:       "from_v5_0_0",
			version:    "5.0.0",
			configFunc: managedTransformsConfigV5,
		},
		{
			name:       "from_v5_8_4", // Current stable release
			version:    "5.8.4",
			configFunc: managedTransformsConfigV5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_managed_transforms." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, zoneID)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			step1 := resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			}

			// Some v5 versions have known issues, skip plan check if needed
			if tc.skipPlanCheck {
				step1.ExpectNonEmptyPlan = true
			}

			steps = append(steps, step1)

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			// MigrationV2TestStep automatically detects v4 vs v5 and only runs migration for v4
			steps = append(steps,
				acctest.MigrationV2TestStep(t, config, tmpDir, tc.version, "v4", "v5", []statecheck.StateCheck{
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
			)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}

func TestMigrateManagedTransformsEdgeCases(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	// Test edge cases with different header combinations
	edgeCases := []struct {
		name           string
		version        string
		configFunc     func(rnd, zoneID string) string
		expectedChecks func(resourceName, zoneID string) []statecheck.StateCheck
	}{
		{
			name:       "only_request_headers_from_v4",
			version:    "4.52.1",
			configFunc: managedTransformsConfigV4RequestOnly,
			expectedChecks: func(resourceName, zoneID string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"id":      knownvalue.StringExact("add_true_client_ip_headers"),
							"enabled": knownvalue.Bool(true),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{})),
				}
			},
		},
		{
			name:       "only_response_headers_from_v4",
			version:    "4.52.1",
			configFunc: managedTransformsConfigV4ResponseOnly,
			expectedChecks: func(resourceName, zoneID string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"id":      knownvalue.StringExact("add_security_headers"),
							"enabled": knownvalue.Bool(true),
						}),
					})),
				}
			},
		},
		{
			name:       "empty_headers_from_v4",
			version:    "4.52.1",
			configFunc: managedTransformsConfigV4Empty,
			expectedChecks: func(resourceName, zoneID string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{})),
				}
			},
		},
		{
			name:       "only_request_headers_from_v5_0_0",
			version:    "5.0.0",
			configFunc: managedTransformsConfigV5RequestOnly,
			expectedChecks: func(resourceName, zoneID string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"id":      knownvalue.StringExact("add_true_client_ip_headers"),
							"enabled": knownvalue.Bool(true),
						}),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{})),
				}
			},
		},
		{
			name:       "only_response_headers_from_v5_0_0",
			version:    "5.0.0",
			configFunc: managedTransformsConfigV5ResponseOnly,
			expectedChecks: func(resourceName, zoneID string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.ObjectExact(map[string]knownvalue.Check{
							"id":      knownvalue.StringExact("add_security_headers"),
							"enabled": knownvalue.Bool(true),
						}),
					})),
				}
			},
		},
		{
			name:       "empty_headers_from_v5_7_1",
			version:    "5.7.1",
			configFunc: managedTransformsConfigV5Empty,
			expectedChecks: func(resourceName, zoneID string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{})),
				}
			},
		},
	}

	for _, tc := range edgeCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_managed_transforms." + rnd
			tmpDir := t.TempDir()

			config := tc.configFunc(rnd, zoneID)
			expectedChecks := tc.expectedChecks(resourceName, zoneID)

			// Build test steps
			steps := []resource.TestStep{}

			// Step 1: Create with specific provider version
			step1 := resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: config,
			}

			// v5.7.0 and v5.7.1 have known issues where they fetch all headers from API
			// causing perpetual diffs, so skip that check
			if tc.version == "5.7.0" || tc.version == "5.7.1" {
				step1.ExpectNonEmptyPlan = true
			}

			steps = append(steps, step1)

			// Step 2: Run migration (for v4) or just upgrade provider (for v5)
			// MigrationV2TestStep automatically detects v4 vs v5 and only runs migration for v4
			steps = append(steps,
				acctest.MigrationV2TestStep(t, config, tmpDir, tc.version, "v4", "v5", expectedChecks),
			)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps:      steps,
			})
		})
	}
}

// Config generators for different provider versions
func managedTransformsConfigV4(rnd, zoneID string) string {
	return fmt.Sprintf(`
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
}

func managedTransformsConfigV5(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_transforms" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_request_headers = [
    {
      id      = "add_true_client_ip_headers"
      enabled = true
    },
    {
      id      = "add_visitor_location_headers"
      enabled = true
    }
  ]
  
  managed_response_headers = [
    {
      id      = "add_security_headers"
      enabled = true
    }
  ]
}`, rnd, zoneID)
}

func managedTransformsConfigV4RequestOnly(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_request_headers {
    id      = "add_true_client_ip_headers"
    enabled = true
  }
}`, rnd, zoneID)
}

func managedTransformsConfigV5RequestOnly(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_transforms" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_request_headers = [
    {
      id      = "add_true_client_ip_headers"
      enabled = true
    }
  ]
  
  managed_response_headers = []
}`, rnd, zoneID)
}

func managedTransformsConfigV4ResponseOnly(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_response_headers {
    id      = "add_security_headers"
    enabled = true
  }
}`, rnd, zoneID)
}

func managedTransformsConfigV5ResponseOnly(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_transforms" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_request_headers = []
  
  managed_response_headers = [
    {
      id      = "add_security_headers"
      enabled = true
    }
  ]
}`, rnd, zoneID)
}

func managedTransformsConfigV4Empty(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_headers" "%[1]s" {
  zone_id = "%[2]s"
}`, rnd, zoneID)
}

func managedTransformsConfigV5Empty(rnd, zoneID string) string {
	return fmt.Sprintf(`
resource "cloudflare_managed_transforms" "%[1]s" {
  zone_id = "%[2]s"
  
  managed_request_headers = []
  managed_response_headers = []
}`, rnd, zoneID)
}

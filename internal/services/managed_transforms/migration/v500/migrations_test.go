package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion
)

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_request_only.tf
var v4RequestOnlyConfig string

//go:embed testdata/v5_request_only.tf
var v5RequestOnlyConfig string

//go:embed testdata/v4_response_only.tf
var v4ResponseOnlyConfig string

//go:embed testdata/v5_response_only.tf
var v5ResponseOnlyConfig string

//go:embed testdata/v4_empty.tf
var v4EmptyConfig string

//go:embed testdata/v5_empty.tf
var v5EmptyConfig string

// TestMigrateManagedHeaders_Basic tests migration with both request and response headers
func TestMigrateManagedHeaders_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_managed_transforms." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: append([]resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
				},
					acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
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
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"id":      knownvalue.StringExact("add_security_headers"),
								"enabled": knownvalue.Bool(true),
							}),
						})),
					})...),
			})
		})
	}
}

// TestMigrateManagedHeaders_RequestOnly tests with only request headers (response becomes empty set)
func TestMigrateManagedHeaders_RequestOnly(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4RequestOnlyConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5RequestOnlyConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_managed_transforms." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: append([]resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
				},
					acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"id":      knownvalue.StringExact("add_true_client_ip_headers"),
								"enabled": knownvalue.Bool(true),
							}),
						})),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{})),
					})...),
			})
		})
	}
}

// TestMigrateManagedHeaders_ResponseOnly tests with only response headers (request becomes empty set)
func TestMigrateManagedHeaders_ResponseOnly(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4ResponseOnlyConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5ResponseOnlyConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_managed_transforms." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: append([]resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
				},
					acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{})),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"id":      knownvalue.StringExact("add_security_headers"),
								"enabled": knownvalue.Bool(true),
							}),
						})),
					})...),
			})
		})
	}
}

// TestMigrateManagedHeaders_Empty tests with no headers (both become empty sets)
func TestMigrateManagedHeaders_Empty(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4EmptyConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5EmptyConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_managed_transforms." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: append([]resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
				},
					acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{})),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{})),
					})...),
			})
		})
	}
}

// TestMigrateManagedHeaders_MultiVersion tests migration from multiple provider versions
func TestMigrateManagedHeaders_MultiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:    "from_v4_52_1",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5_0_0",
			version: "5.0.0",
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, zoneID)
			},
		},
		{
			name:    "from_v5_8_4",
			version: "5.8.4",
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, zoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_managed_transforms." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: append([]resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
				},
					acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
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
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.ObjectExact(map[string]knownvalue.Check{
								"id":      knownvalue.StringExact("add_security_headers"),
								"enabled": knownvalue.Bool(true),
							}),
						})),
					})...),
			})
		})
	}
}

// TestMigrateManagedHeaders_EdgeCases tests edge cases from various versions
func TestMigrateManagedHeaders_EdgeCases(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
		checks   func(resourceName, zoneID string) []statecheck.StateCheck
	}{
		{
			name:    "request_only_from_v4",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4RequestOnlyConfig, rnd, zoneID)
			},
			checks: func(resourceName, zoneID string) []statecheck.StateCheck {
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
			name:    "response_only_from_v4",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4ResponseOnlyConfig, rnd, zoneID)
			},
			checks: func(resourceName, zoneID string) []statecheck.StateCheck {
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
			name:    "empty_from_v4",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v4EmptyConfig, rnd, zoneID)
			},
			checks: func(resourceName, zoneID string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{})),
				}
			},
		},
		{
			name:    "request_only_from_v5_0_0",
			version: "5.0.0",
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5RequestOnlyConfig, rnd, zoneID)
			},
			checks: func(resourceName, zoneID string) []statecheck.StateCheck {
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
			name:    "response_only_from_v5_0_0",
			version: "5.0.0",
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5ResponseOnlyConfig, rnd, zoneID)
			},
			checks: func(resourceName, zoneID string) []statecheck.StateCheck {
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
			name:    "empty_from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID string) string {
				return fmt.Sprintf(v5EmptyConfig, rnd, zoneID)
			},
			checks: func(resourceName, zoneID string) []statecheck.StateCheck {
				return []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_request_headers"), knownvalue.SetExact([]knownvalue.Check{})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("managed_response_headers"), knownvalue.SetExact([]knownvalue.Check{})),
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_managed_transforms." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: append([]resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
				},
					acctest.MigrationV2TestStepWithPlan(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, tc.checks(resourceName, zoneID))...),
			})
		})
	}
}

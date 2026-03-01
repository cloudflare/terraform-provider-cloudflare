package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed migration test configuration files

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_simple_rules.tf
var v4SimpleRulesConfig string

//go:embed testdata/v5_simple_rules.tf
var v5SimpleRulesConfig string

//go:embed testdata/v4_headers.tf
var v4HeadersConfig string

//go:embed testdata/v5_headers.tf
var v5HeadersConfig string

//go:embed testdata/v4_log_custom_fields.tf
var v4LogCustomFieldsConfig string

//go:embed testdata/v5_log_custom_fields.tf
var v5LogCustomFieldsConfig string

//go:embed testdata/v4_ratelimit.tf
var v4RatelimitConfig string

//go:embed testdata/v5_ratelimit.tf
var v5RatelimitConfig string

// TestMigrateRulesetBasic tests migration of a minimal ruleset (no rules).
// Covers: schema version bump, top-level scalar fields.
func TestMigrateRulesetBasic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4BasicConfig, zoneID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5BasicConfig, zoneID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("My ruleset %s", rnd))),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("phase"), knownvalue.StringExact("http_request_firewall_custom")),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("kind"), knownvalue.StringExact("zone")),
					}),
				},
			})
		})
	}
}

// TestMigrateRulesetSimpleRules tests migration of rules blocks to list attribute syntax.
// Covers: rules ListNestedBlock → ListNestedAttribute, basic rule fields.
func TestMigrateRulesetSimpleRules(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4SimpleRulesConfig, zoneID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5SimpleRulesConfig, zoneID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expression"), knownvalue.StringExact("ip.src eq 1.1.1.1")),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action"), knownvalue.StringExact("block")),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("expression"), knownvalue.StringExact("ip.src eq 2.2.2.2")),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action"), knownvalue.StringExact("log")),
					}),
				},
			})
		})
	}
}

// TestMigrateRulesetHeadersListToMap tests migration of headers from list blocks to map attribute.
// Covers: headers ListNestedBlock (with name field) → MapNestedAttribute (keyed by name).
func TestMigrateRulesetHeadersListToMap(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4HeadersConfig, zoneID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5HeadersConfig, zoneID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("headers").AtMapKey("X-Custom-Header").AtMapKey("operation"), knownvalue.StringExact("set")),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("headers").AtMapKey("X-Custom-Header").AtMapKey("value"), knownvalue.StringExact("custom-value")),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("headers").AtMapKey("Authorization").AtMapKey("operation"), knownvalue.StringExact("remove")),
					}),
				},
			})
		})
	}
}

// TestMigrateRulesetLogCustomFields tests migration of log custom field arrays.
// Covers: cookie_fields/request_fields: Set[string] → List[{name}],
// response_fields: Set[string] → List[{name, preserve_duplicates}].
func TestMigrateRulesetLogCustomFields(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4LogCustomFieldsConfig, zoneID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5LogCustomFieldsConfig, zoneID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),
						// Rule 0: cookie_fields transformed to list of objects (API may return in any order)
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("action_parameters").AtMapKey("cookie_fields"), knownvalue.ListSizeExact(3)),
						// Rule 1: request_fields transformed to list of objects (API may return in any order)
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("action_parameters").AtMapKey("request_fields"), knownvalue.ListSizeExact(2)),
						// Rule 2: response_fields transformed to list of objects (API may return in any order)
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("action_parameters").AtMapKey("response_fields"), knownvalue.ListSizeExact(2)),
					}),
				},
			})
		})
	}
}

// TestMigrateRulesetRateLimit tests migration of ratelimit blocks to attributes.
// Covers: ratelimit ListNestedBlock → SingleNestedAttribute, characteristics Set → List.
func TestMigrateRulesetRateLimit(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID string) string
	}{
		{
			name:     "from_v4_latest",
			version:  acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v4RatelimitConfig, zoneID, rnd) },
		},
		{
			name:     "from_v5",
			version:  currentProviderVersion,
			configFn: func(rnd, zoneID string) string { return fmt.Sprintf(v5RatelimitConfig, zoneID, rnd) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: testConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),
						// Rule 0: ratelimit is now an object (not array), characteristics is a list
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("characteristics"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("period"), knownvalue.Int64Exact(60)),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("requests_per_period"), knownvalue.Int64Exact(100)),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("ratelimit").AtMapKey("mitigation_timeout"), knownvalue.Int64Exact(600)),
						// Rule 1: counting_expression preserved
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("ratelimit").AtMapKey("counting_expression"), knownvalue.StringExact("http.request.method eq \"POST\"")),
						statecheck.ExpectKnownValue("cloudflare_ruleset."+rnd, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("ratelimit").AtMapKey("mitigation_timeout"), knownvalue.Int64Exact(0)),
					}),
				},
			})
		})
	}
}

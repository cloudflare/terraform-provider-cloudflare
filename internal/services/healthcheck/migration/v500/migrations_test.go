package v500_test

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// buildHealthcheckPlanChecks returns the appropriate plan checks for migration step 2.
// For from_v4 (v4→v5): allows tcp_config/http_config computed field drift.
// For from_v5 (v5→v5): requires a completely empty plan.
func buildHealthcheckPlanChecks(sourceVer string) []plancheck.PlanCheck {
	if sourceVer == "v4" {
		return []plancheck.PlanCheck{
			acctest.DebugNonEmptyPlan,
			ExpectEmptyPlanExceptHealthcheckConfigDrift,
		}
	}
	return []plancheck.PlanCheck{
		acctest.DebugNonEmptyPlan,
		plancheck.ExpectEmptyPlan(),
	}
}

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test configs - HTTP Basic
//
//go:embed testdata/v4_http_basic.tf
var v4HTTPBasicConfig string

//go:embed testdata/v5_http_basic.tf
var v5HTTPBasicConfig string

// Embed test configs - HTTP with Headers
//
//go:embed testdata/v4_http_headers.tf
var v4HTTPHeadersConfig string

//go:embed testdata/v5_http_headers.tf
var v5HTTPHeadersConfig string

// Embed test configs - TCP
//
//go:embed testdata/v4_tcp.tf
var v4TCPConfig string

//go:embed testdata/v5_tcp.tf
var v5TCPConfig string

// Embed test configs - HTTPS Full
//
//go:embed testdata/v4_https_full.tf
var v4HTTPSFullConfig string

//go:embed testdata/v5_https_full.tf
var v5HTTPSFullConfig string

// Custom plan check for healthcheck that allows only tcp_config/http_config computed drift
type expectEmptyPlanExceptHealthcheckConfigDrift struct{}

func (e expectEmptyPlanExceptHealthcheckConfigDrift) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, rc := range req.Plan.ResourceChanges {
		// Skip no-op and read actions
		if rc.Change.Actions[0] == "no-op" || rc.Change.Actions[0] == "read" {
			continue
		}

		// Check if this is an update action
		if rc.Change.Actions[0] != "update" {
			resp.Error = fmt.Errorf("expected empty plan, but %s has planned action(s): %v", rc.Address, rc.Change.Actions)
			return
		}

		// For updates, check each attribute change
		beforeMap, beforeOk := rc.Change.Before.(map[string]interface{})
		afterMap, afterOk := rc.Change.After.(map[string]interface{})

		if !beforeOk || !afterOk {
			resp.Error = fmt.Errorf("expected empty plan, but %s has non-map changes", rc.Address)
			return
		}

		// Check each attribute that's different
		for key, afterValue := range afterMap {
			beforeValue, _ := beforeMap[key]

			// Skip if values are the same
			if reflect.DeepEqual(beforeValue, afterValue) {
				continue
			}

			// Allow changes from falsey to null
			if afterValue == nil {
				if isFalseyValue(beforeValue) {
					continue // This change is allowed
				}
			}

			// Healthcheck specific: Allow tcp_config/http_config computed field drift
			// The unused config (tcp_config for HTTP type, http_config for TCP type) is marked as
			// Computed:true in the schema and will show as "(known after apply)" on refresh
			if key == "tcp_config" || key == "http_config" {
				if beforeValue == nil {
					continue // This change is allowed - computed field drift
				}
			}

			// If we get here, it's a disallowed change
			resp.Error = fmt.Errorf("expected empty plan except for healthcheck config drift, but %s.%s has change from %v to %v",
				rc.Address, key, beforeValue, afterValue)
			return
		}
	}
}

// isFalseyValue checks if a value is "falsey" (false, 0, "", empty slice/map)
func isFalseyValue(v interface{}) bool {
	if v == nil {
		return true
	}
	switch val := v.(type) {
	case bool:
		return !val
	case string:
		return val == ""
	case float64:
		return val == 0
	case int:
		return val == 0
	case []interface{}:
		return len(val) == 0
	case map[string]interface{}:
		return len(val) == 0
	default:
		return false
	}
}

var ExpectEmptyPlanExceptHealthcheckConfigDrift = expectEmptyPlanExceptHealthcheckConfigDrift{}

// TestMigrateHealthcheck_V4ToV5_HTTPBasic tests basic HTTP healthcheck migration
// with flat → nested http_config transformation
func TestMigrateHealthcheck_V4ToV5_HTTPBasic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:    "from_v4_latest", // Tests legacy v4 → current v5
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, name string) string {
				return fmt.Sprintf(v4HTTPBasicConfig, rnd, zoneID, name)
			},
		},
		{
			name:    "from_v5", // Tests within v5 (version bump)
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, name string) string {
				return fmt.Sprintf(v5HTTPBasicConfig, rnd, zoneID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// Build Step 1
			step1 := resource.TestStep{
				// Step 1: Create with specific provider version
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: testConfig,
			}

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					step1,
					// Step 2: Run migration and verify state with custom plan check.
					// For from_v4: allows tcp_config/http_config computed field drift.
					// For from_v5: requires a completely empty plan (no PostApplyPostRefresh
					// to avoid "expected a non-empty plan" framework error when plan is empty).
					{
						PreConfig: func() {
							acctest.WriteOutConfig(t, testConfig, tmpDir)
							acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: buildHealthcheckPlanChecks(sourceVer),
						},
						ConfigStateChecks: []statecheck.StateCheck{
							// Verify core fields
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("zone_id"),
								knownvalue.StringExact(zoneID),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("name"),
								knownvalue.StringExact(name),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("address"),
								knownvalue.StringExact("example.com"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("HTTP"),
							),

							// Verify HTTP config nested object exists
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config"),
								knownvalue.NotNull(),
							),

							// Verify HTTP config fields (flat → nested transformation)
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("method"),
								knownvalue.StringExact("GET"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("port"),
								knownvalue.Int64Exact(80),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("path"),
								knownvalue.StringExact("/health"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("expected_body"),
								knownvalue.StringExact("OK"),
							),

							// Note: No need to verify flat fields are null at root -
							// v5 schema doesn't have method/port/path at root level, they only exist in http_config

							// Verify TCP config doesn't exist (HTTP type)
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("tcp_config"),
								knownvalue.Null(),
							),
						},
					},
				},
			})
		})
	}
}

// TestMigrateHealthcheck_V4ToV5_HTTPHeaders tests header Set → Map transformation
func TestMigrateHealthcheck_V4ToV5_HTTPHeaders(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, name string) string {
				return fmt.Sprintf(v4HTTPHeadersConfig, rnd, zoneID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, name string) string {
				return fmt.Sprintf(v5HTTPHeadersConfig, rnd, zoneID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// Build Step 1
			step1 := resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: testConfig,
			}

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					step1,
					// Step 2: Run migration and verify state with custom plan check.
					// For from_v4: allows tcp_config/http_config computed field drift.
					// For from_v5: requires a completely empty plan (no PostApplyPostRefresh
					// to avoid "expected a non-empty plan" framework error when plan is empty).
					{
						PreConfig: func() {
							acctest.WriteOutConfig(t, testConfig, tmpDir)
							acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: buildHealthcheckPlanChecks(sourceVer),
						},
						ConfigStateChecks: []statecheck.StateCheck{
							// Verify http_config.header exists as map
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("header"),
								knownvalue.NotNull(),
							),

							// Verify header map contains transformed values
							// v4: header { header = "Host" values = ["example.com"] }
							// v5: header = { "Host" = ["example.com"] }
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("header").AtMapKey("Host"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact("example.com"),
								}),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("header").AtMapKey("User-Agent"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact("Cloudflare-Healthcheck/1.0"),
								}),
							),

							// Note: No need to verify flat header field is null at root -
							// v5 schema doesn't have header at root level, it only exists in http_config
						},
					},
				},
			})
		})
	}
}

// TestMigrateHealthcheck_V4ToV5_TCP tests TCP config transformation
func TestMigrateHealthcheck_V4ToV5_TCP(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, name string) string {
				return fmt.Sprintf(v4TCPConfig, rnd, zoneID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, name string) string {
				return fmt.Sprintf(v5TCPConfig, rnd, zoneID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// Build Step 1
			step1 := resource.TestStep{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: tc.version,
					},
				},
				Config: testConfig,
			}

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					step1,
					// Step 2: Run migration and verify state with custom plan check.
					// For from_v4: allows tcp_config/http_config computed field drift.
					// For from_v5: requires a completely empty plan (no PostApplyPostRefresh
					// to avoid "expected a non-empty plan" framework error when plan is empty).
					{
						PreConfig: func() {
							acctest.WriteOutConfig(t, testConfig, tmpDir)
							acctest.RunMigrationV2Command(t, testConfig, tmpDir, sourceVer, targetVer)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						ConfigDirectory:          config.StaticDirectory(tmpDir),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: buildHealthcheckPlanChecks(sourceVer),
						},
						ConfigStateChecks: []statecheck.StateCheck{
							// Verify type is TCP
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("TCP"),
							),

							// Verify tcp_config nested object exists
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("tcp_config"),
								knownvalue.NotNull(),
							),

							// Verify TCP config fields (flat → nested transformation)
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("tcp_config").AtMapKey("method"),
								knownvalue.StringExact("connection_established"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("tcp_config").AtMapKey("port"),
								knownvalue.Int64Exact(443),
							),

							// Note: No need to verify flat fields are null at root -
							// v5 schema doesn't have method/port at root level, they only exist in tcp_config

							// Verify HTTP config doesn't exist (TCP type)
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config"),
								knownvalue.Null(),
							),
						},
					},
				},
			})
		})
	}
}

// TestMigrateHealthcheck_V4ToV5_HTTPSFull tests full HTTPS healthcheck with all options
func TestMigrateHealthcheck_V4ToV5_HTTPSFull(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, name string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, name string) string {
				return fmt.Sprintf(v4HTTPSFullConfig, rnd, zoneID, name)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, name string) string {
				return fmt.Sprintf(v5HTTPSFullConfig, rnd, zoneID, name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			rnd := utils.GenerateRandomResourceName()
			name := "tf-test-" + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, zoneID, name)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify core fields
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("type"),
								knownvalue.StringExact("HTTPS"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("description"),
								knownvalue.StringExact("Full HTTPS healthcheck"),
							),

							// Verify check_regions (List conversion)
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("check_regions"),
								knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact("WNAM"),
									knownvalue.StringExact("ENAM"),
								}),
							),

							// Verify http_config with all HTTPS fields
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("method"),
								knownvalue.StringExact("GET"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("port"),
								knownvalue.Int64Exact(443),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("path"),
								knownvalue.StringExact("/api/health"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("expected_body"),
								knownvalue.StringExact("healthy"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("follow_redirects"),
								knownvalue.Bool(true),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("http_config").AtMapKey("allow_insecure"),
								knownvalue.Bool(true),
							),

							// Verify config values
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("interval"),
								knownvalue.Int64Exact(60),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("retries"),
								knownvalue.Int64Exact(2),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("timeout"),
								knownvalue.Int64Exact(10),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("consecutive_fails"),
								knownvalue.Int64Exact(2),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("consecutive_successes"),
								knownvalue.Int64Exact(2),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_healthcheck."+rnd,
								tfjsonpath.New("suspended"),
								knownvalue.Bool(false),
							),
						},
					),
				},
			})
		})
	}
}

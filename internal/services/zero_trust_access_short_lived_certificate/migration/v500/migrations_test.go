package v500_test

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_short_lived_certificate/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion
)

//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_zone_scoped.tf
var v4ZoneScopedConfig string

//go:embed testdata/v5_zone_scoped.tf
var v5ZoneScopedConfig string

// --------------------------------------------------------------------------
// Acceptance tests (require TF_ACC=1 and API credentials)
// --------------------------------------------------------------------------

// TestMigrateAccessCACertificateBasic tests basic migration with account_id.
// Verifies that cloudflare_access_ca_certificate (v4) state is correctly
// transformed to cloudflare_zero_trust_access_short_lived_certificate (v5),
// including the application_id -> app_id attribute rename.
func TestMigrateAccessCACertificateBasic(t *testing.T) {
	testCases := []struct {
		name       string
		version    string
		setupFn    func(rnd, accountID, domain string) string // config for Step 1 (create)
		migrateFn  func(rnd, accountID, domain string) string // config for Step 2 (post-migration plan)
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			setupFn: func(rnd, accountID, domain string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, domain)
			},
			migrateFn: func(rnd, accountID, domain string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			setupFn: func(rnd, accountID, domain string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)
			},
			migrateFn: func(rnd, accountID, domain string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, domain)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			resourceName := "cloudflare_zero_trust_access_short_lived_certificate." + rnd

			setupConfig := tc.setupFn(rnd, accountID, domain)
			migrateConfig := tc.migrateFn(rnd, accountID, domain)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   setupConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: setupConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				CheckDestroy: nil,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					migrationStep(t, migrateConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("app_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("aud"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("public_key"), knownvalue.NotNull()),
					}),
					// Clear state before post-test destroy. The Access CA delete
					// API returns 500 (error 12055); removing resources from state
					// prevents the framework's destroy from hitting that bug.
					{
						PreConfig: func() {
							clearState(t, tmpDir)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						Config:                   "# empty",
						PlanOnly:                 true,
					},
				},
			})
		})
	}
}

// TestMigrateAccessCACertificateZoneScoped tests zone-scoped resource migration.
func TestMigrateAccessCACertificateZoneScoped(t *testing.T) {
	testCases := []struct {
		name      string
		version   string
		setupFn   func(rnd, zoneID, domain string) string // config for Step 1 (create)
		migrateFn func(rnd, zoneID, domain string) string // config for Step 2 (post-migration plan)
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			setupFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v4ZoneScopedConfig, rnd, zoneID, domain)
			},
			migrateFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v5ZoneScopedConfig, rnd, zoneID, domain)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			setupFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v5ZoneScopedConfig, rnd, zoneID, domain)
			},
			migrateFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v5ZoneScopedConfig, rnd, zoneID, domain)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			rnd := utils.GenerateRandomResourceName()
			tmpDir := t.TempDir()
			resourceName := "cloudflare_zero_trust_access_short_lived_certificate." + rnd

			setupConfig := tc.setupFn(rnd, zoneID, domain)
			migrateConfig := tc.migrateFn(rnd, zoneID, domain)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   setupConfig,
				}
			} else {
				firstStep = resource.TestStep{
					ExternalProviders: map[string]resource.ExternalProvider{
						"cloudflare": {
							Source:            "cloudflare/cloudflare",
							VersionConstraint: tc.version,
						},
					},
					Config: setupConfig,
				}
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_ZoneID(t)
				},
				CheckDestroy: nil,
				WorkingDir:   tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					migrationStep(t, migrateConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("app_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("aud"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("public_key"), knownvalue.NotNull()),
					}),
					{
						PreConfig: func() {
							clearState(t, tmpDir)
						},
						ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
						Config:                   "# empty",
						PlanOnly:                 true,
					},
				},
			})
		})
	}
}

// --------------------------------------------------------------------------
// Test helpers (acceptance)
// --------------------------------------------------------------------------

// v4TypeRenames maps v4 resource type names to their v5 equivalents. After
// tf-migrate rewrites the HCL config, the state file still contains the old
// type names. The v5 provider cannot read state entries with unknown types, so
// we rename them in-place before Terraform tries to load the state.
var v4TypeRenames = map[string]string{
	"cloudflare_access_application":  "cloudflare_zero_trust_access_application",
	"cloudflare_access_ca_certificate": "cloudflare_zero_trust_access_short_lived_certificate",
}

func renameStateTypes(t *testing.T, tmpDir string, renames map[string]string) {
	t.Helper()

	matches, err := filepath.Glob(filepath.Join(tmpDir, "work*"))
	if err != nil || len(matches) == 0 {
		t.Logf("renameStateTypes: no work* directory found in %s", tmpDir)
		return
	}
	stateFile := filepath.Join(matches[0], "terraform.tfstate")

	data, err := os.ReadFile(stateFile)
	if err != nil {
		t.Fatalf("renameStateTypes: cannot read state file %s: %v", stateFile, err)
	}

	var state map[string]json.RawMessage
	if err := json.Unmarshal(data, &state); err != nil {
		t.Fatalf("renameStateTypes: cannot parse state file: %v", err)
	}

	var resources []map[string]json.RawMessage
	if err := json.Unmarshal(state["resources"], &resources); err != nil {
		t.Fatalf("renameStateTypes: cannot parse resources: %v", err)
	}

	changed := false
	for i, res := range resources {
		var typ string
		if err := json.Unmarshal(res["type"], &typ); err != nil {
			continue
		}
		if newType, ok := renames[typ]; ok {
			newTypeJSON, _ := json.Marshal(newType)
			resources[i]["type"] = newTypeJSON
			changed = true
		}
	}

	if !changed {
		return
	}

	// Increment serial so Terraform accepts the modified state.
	var serial int64
	if err := json.Unmarshal(state["serial"], &serial); err == nil {
		serialJSON, _ := json.Marshal(serial + 1)
		state["serial"] = serialJSON
	}

	resourcesJSON, _ := json.Marshal(resources)
	state["resources"] = resourcesJSON

	out, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		t.Fatalf("renameStateTypes: failed to marshal state: %v", err)
	}
	if err := os.WriteFile(stateFile, out, 0600); err != nil {
		t.Fatalf("renameStateTypes: failed to write state: %v", err)
	}
}

// clearState empties the resources array in the terraform.tfstate file so the
// framework's post-test destroy has nothing to clean up.
func clearState(t *testing.T, tmpDir string) {
	t.Helper()

	matches, err := filepath.Glob(filepath.Join(tmpDir, "work*"))
	if err != nil || len(matches) == 0 {
		return
	}
	stateFile := filepath.Join(matches[0], "terraform.tfstate")

	data, err := os.ReadFile(stateFile)
	if err != nil {
		return
	}

	var state map[string]json.RawMessage
	if err := json.Unmarshal(data, &state); err != nil {
		return
	}

	state["resources"] = json.RawMessage("[]")

	var serial int64
	if err := json.Unmarshal(state["serial"], &serial); err == nil {
		serialJSON, _ := json.Marshal(serial + 1)
		state["serial"] = serialJSON
	}

	out, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return
	}
	_ = os.WriteFile(stateFile, out, 0600)
}

// migrationStep builds a test step that runs tf-migrate, renames v4 state
// entries to their v5 types, and validates the result with the v5 provider.
func migrationStep(t *testing.T, v5Config string, tmpDir string, exactVersion string, sourceVersion string, targetVersion string, stateChecks []statecheck.StateCheck) resource.TestStep {
	t.Helper()

	var planChecks []plancheck.PlanCheck
	if sourceVersion == "v4" {
		planChecks = []plancheck.PlanCheck{
			acctest.DebugNonEmptyPlan,
			acctest.ExpectEmptyPlanExceptFalseyToNull,
		}
	} else {
		planChecks = []plancheck.PlanCheck{
			acctest.DebugNonEmptyPlan,
			plancheck.ExpectEmptyPlan(),
		}
	}

	return resource.TestStep{
		PreConfig: func() {
			acctest.WriteOutConfig(t, v5Config, tmpDir)
			acctest.RunMigrationV2Command(t, v5Config, tmpDir, sourceVersion, targetVersion)
			renameStateTypes(t, tmpDir, v4TypeRenames)

			// Remove provider.tf so the test framework uses ProtoV6ProviderFactories
			// instead of resolving to the published registry binary.
			providerTF := filepath.Join(tmpDir, "provider.tf")
			if err := os.Remove(providerTF); err != nil && !os.IsNotExist(err) {
				t.Fatalf("failed to remove provider.tf after migration: %v", err)
			}
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: planChecks,
		},
		ConfigStateChecks: stateChecks,
	}
}

// --------------------------------------------------------------------------
// Unit tests (no API credentials required)
// --------------------------------------------------------------------------

// TestMigrateAccessCACertificateTransform is a unit test for the Transform function.
// It verifies the attribute mapping without requiring API credentials.
func TestMigrateAccessCACertificateTransform(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		source v500.SourceAccessCACertificateModel
		check  func(t *testing.T, target *v500.TargetAccessShortLivedCertificateModel)
	}{
		{
			name: "basic_account_scoped",
			source: v500.SourceAccessCACertificateModel{
				ID:            types.StringValue("test-id"),
				AccountID:     types.StringValue("abc123"),
				ZoneID:        types.StringNull(),
				ApplicationID: types.StringValue("app-uuid"),
				AUD:           types.StringValue("aud-tag"),
				PublicKey:     types.StringValue("ssh-rsa AAAA..."),
			},
			check: func(t *testing.T, target *v500.TargetAccessShortLivedCertificateModel) {
				t.Helper()
				assertStringValue(t, "id", target.ID, "test-id")
				assertStringValue(t, "app_id", target.AppID, "app-uuid")
				assertStringValue(t, "account_id", target.AccountID, "abc123")
				assertStringNull(t, "zone_id", target.ZoneID)
				assertStringValue(t, "aud", target.AUD, "aud-tag")
				assertStringValue(t, "public_key", target.PublicKey, "ssh-rsa AAAA...")
			},
		},
		{
			name: "basic_zone_scoped",
			source: v500.SourceAccessCACertificateModel{
				ID:            types.StringValue("test-id"),
				AccountID:     types.StringNull(),
				ZoneID:        types.StringValue("zone-123"),
				ApplicationID: types.StringValue("app-uuid"),
				AUD:           types.StringValue("aud-tag"),
				PublicKey:     types.StringValue("ssh-rsa AAAA..."),
			},
			check: func(t *testing.T, target *v500.TargetAccessShortLivedCertificateModel) {
				t.Helper()
				assertStringValue(t, "id", target.ID, "test-id")
				assertStringValue(t, "app_id", target.AppID, "app-uuid")
				assertStringNull(t, "account_id", target.AccountID)
				assertStringValue(t, "zone_id", target.ZoneID, "zone-123")
			},
		},
		{
			name: "empty_string_account_id_becomes_null",
			source: v500.SourceAccessCACertificateModel{
				ID:            types.StringValue("test-id"),
				AccountID:     types.StringValue(""), // v4 Computed could produce empty string
				ZoneID:        types.StringValue("zone-123"),
				ApplicationID: types.StringValue("app-uuid"),
				AUD:           types.StringValue("aud-tag"),
				PublicKey:     types.StringValue("ssh-rsa AAAA..."),
			},
			check: func(t *testing.T, target *v500.TargetAccessShortLivedCertificateModel) {
				t.Helper()
				assertStringNull(t, "account_id", target.AccountID)
				assertStringValue(t, "zone_id", target.ZoneID, "zone-123")
			},
		},
		{
			name: "empty_string_zone_id_becomes_null",
			source: v500.SourceAccessCACertificateModel{
				ID:            types.StringValue("test-id"),
				AccountID:     types.StringValue("abc123"),
				ZoneID:        types.StringValue(""), // v4 Computed could produce empty string
				ApplicationID: types.StringValue("app-uuid"),
				AUD:           types.StringValue("aud-tag"),
				PublicKey:     types.StringValue("ssh-rsa AAAA..."),
			},
			check: func(t *testing.T, target *v500.TargetAccessShortLivedCertificateModel) {
				t.Helper()
				assertStringValue(t, "account_id", target.AccountID, "abc123")
				assertStringNull(t, "zone_id", target.ZoneID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()
			target, diags := v500.Transform(ctx, tc.source)
			if diags.HasError() {
				t.Fatalf("Transform returned errors: %v", diags)
			}
			if target == nil {
				t.Fatal("Transform returned nil target")
			}
			tc.check(t, target)
		})
	}
}

// --------------------------------------------------------------------------
// Test helpers
// --------------------------------------------------------------------------

func assertStringValue(t *testing.T, field string, got types.String, want string) {
	t.Helper()
	if got.IsNull() || got.IsUnknown() {
		t.Errorf("%s: expected %q, got null/unknown", field, want)
		return
	}
	if got.ValueString() != want {
		t.Errorf("%s: expected %q, got %q", field, want, got.ValueString())
	}
}

func assertStringNull(t *testing.T, field string, got types.String) {
	t.Helper()
	if !got.IsNull() {
		t.Errorf("%s: expected null, got %q", field, got.ValueString())
	}
}

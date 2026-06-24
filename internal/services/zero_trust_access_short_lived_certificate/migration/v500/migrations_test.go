package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
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
		name     string
		version  string
		configFn func(rnd, accountID, domain string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, domain string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, domain)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, domain string) string {
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

			testConfig := tc.configFn(rnd, accountID, domain)
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
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					firstStep,
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("app_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("aud"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("public_key"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateAccessCACertificateZoneScoped tests zone-scoped resource migration.
func TestMigrateAccessCACertificateZoneScoped(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, zoneID, domain string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, zoneID, domain string) string {
				return fmt.Sprintf(v4ZoneScopedConfig, rnd, zoneID, domain)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, zoneID, domain string) string {
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

			testConfig := tc.configFn(rnd, zoneID, domain)
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
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("app_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("aud"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("public_key"), knownvalue.NotNull()),
					}),
				},
			})
		})
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

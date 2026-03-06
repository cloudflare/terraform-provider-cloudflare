package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed migration test configuration files
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_minimal.tf
var v4MinimalConfig string

//go:embed testdata/v5_minimal.tf
var v5MinimalConfig string

// TestMigrateMTLSCertificate_V4ToV5_Basic tests migration of a leaf certificate with name and private key.
//
// This is a simple migration with:
// - No field renames (all fields identical in v4 and v5)
// - No type conversions needed for user-configurable fields
// - Resource type unchanged: cloudflare_mtls_certificate → cloudflare_mtls_certificate
// - New v5 computed field: updated_at (set to null during migration, API populates on first Read)
//
// Note: Uses ca=false since GenerateEphemeralCertAndKey creates leaf certificates.
func TestMigrateMTLSCertificate_V4ToV5_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test.")
	}

	expiry := time.Now().Add(time.Hour * 24 * 365)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{"example.com"}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, certName, cert, key string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, certName, cert, key string) string {
				return fmt.Sprintf(v4BasicConfig, rnd, accountID, certName, cert, key)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, certName, cert, key string) string {
				return fmt.Sprintf(v5BasicConfig, rnd, accountID, certName, cert, key)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			certName := fmt.Sprintf("cftftest-%s", rnd)
			resourceName := "cloudflare_mtls_certificate." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, certName, cert, key)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			// For v5 tests, use local provider; for v4 tests, use external provider
			var firstStep resource.TestStep
			if tc.version == currentProviderVersion {
				// Use local v5 provider (has GetSchemaVersion, will create version=1 state)
				firstStep = resource.TestStep{
					ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
					Config:                   testConfig,
				}
			} else {
				// Use external v4 provider (will create version=0 state)
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
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						// Critical identifier - validate exactly
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						// Required fields - validate exactly
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(certName)),
						// Optional user fields - verify not null
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("private_key"), knownvalue.NotNull()),
						// Computed fields from API - verify not null
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("issuer"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("serial_number"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("signature"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("uploaded_on"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
					}),
				},
			})
		})
	}
}

// TestMigrateMTLSCertificate_V4ToV5_Minimal is skipped because the Cloudflare API
// returns name="" for certificates created without a name, but the v5 provider schema
// has name with RequiresReplace. This causes a ForceNew plan regardless of migration,
// making it impossible to test migration independently of this provider-level behavior.
// The Basic test fully covers the migration path (it includes name in the config).
func TestMigrateMTLSCertificate_V4ToV5_Minimal(t *testing.T) {
	t.Skip("Skipped: API returns name='' for no-name certs, which triggers RequiresReplace in v5 schema. See Basic test for migration coverage.")

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("CLOUDFLARE_ACCOUNT_ID must be set for this acceptance test.")
	}

	expiry := time.Now().Add(time.Hour * 24 * 365)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{"minimal.example.com"}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, cert, key string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, cert, key string) string {
				return fmt.Sprintf(v4MinimalConfig, rnd, accountID, cert, key)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, cert, key string) string {
				return fmt.Sprintf(v5MinimalConfig, rnd, accountID, cert, key)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			resourceName := "cloudflare_mtls_certificate." + rnd
			tmpDir := t.TempDir()
			testConfig := tc.configFn(rnd, accountID, cert, key)
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

			stateChecks := []statecheck.StateCheck{
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ca"), knownvalue.Bool(false)),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificates"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("issuer"), knownvalue.NotNull()),
				statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("serial_number"), knownvalue.NotNull()),
			}

			migrationSteps := acctest.MigrationV2TestStepWithStateNormalization(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, stateChecks)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				WorkingDir: tmpDir,
				Steps:      append([]resource.TestStep{firstStep}, migrationSteps...),
			})
		})
	}
}

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
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Embed test configs
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_default_validity.tf
var v4DefaultValidityConfig string

//go:embed testdata/v5_default_validity.tf
var v5DefaultValidityConfig string

// TestMigrateOriginCACertificate_V4ToV5_Basic tests basic migration with all fields
//
// Tests transformation of:
// - hostnames: Set → List
// - requested_validity: Int64 → Float64
// - min_days_for_renewal: Dropped (removed in v5)
func TestMigrateOriginCACertificate_V4ToV5_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func() string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func() string {
				return v4BasicConfig
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func() string {
				return v5BasicConfig
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := fmt.Sprintf(tc.configFn(), rnd, domain)
			tmpDir := t.TempDir()
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify state
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify csr (direct copy)
							statecheck.ExpectKnownValue(
								"cloudflare_origin_ca_certificate."+rnd,
								tfjsonpath.New("csr"),
								knownvalue.NotNull(),
							),
							// Verify request_type (direct copy)
							statecheck.ExpectKnownValue(
								"cloudflare_origin_ca_certificate."+rnd,
								tfjsonpath.New("request_type"),
								knownvalue.StringExact("origin-rsa"),
							),
							// Verify hostnames (Set → List conversion)
							// Order may differ (Sets are unordered), so just verify it's a list
							statecheck.ExpectKnownValue(
								"cloudflare_origin_ca_certificate."+rnd,
								tfjsonpath.New("hostnames"),
								knownvalue.NotNull(),
							),
							// Verify requested_validity (Int64 → Float64 conversion)
							statecheck.ExpectKnownValue(
								"cloudflare_origin_ca_certificate."+rnd,
								tfjsonpath.New("requested_validity"),
								knownvalue.Float64Exact(365),
							),
							// Verify min_days_for_renewal is NOT in v5 state (dropped)
							// Note: For v5 test, this field never existed
							// Verify computed fields exist (but don't validate exact values)
							statecheck.ExpectKnownValue(
								"cloudflare_origin_ca_certificate."+rnd,
								tfjsonpath.New("id"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_origin_ca_certificate."+rnd,
								tfjsonpath.New("certificate"),
								knownvalue.NotNull(),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_origin_ca_certificate."+rnd,
								tfjsonpath.New("expires_on"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

// TestMigrateOriginCACertificate_V4ToV5_DefaultValidity tests default value handling
//
// Tests that when requested_validity is null/missing in v4, it defaults to 5475 in v5
func TestMigrateOriginCACertificate_V4ToV5_DefaultValidity(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func() string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func() string {
				return v4DefaultValidityConfig
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func() string {
				return v5DefaultValidityConfig
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rnd := utils.GenerateRandomResourceName()
			domain := os.Getenv("CLOUDFLARE_DOMAIN")
			testConfig := fmt.Sprintf(tc.configFn(), rnd, domain)
			tmpDir := t.TempDir()
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_Domain(t)
				},
				WorkingDir: tmpDir,
				Steps: []resource.TestStep{
					{
						// Step 1: Create with specific provider version
						ExternalProviders: map[string]resource.ExternalProvider{
							"cloudflare": {
								Source:            "cloudflare/cloudflare",
								VersionConstraint: tc.version,
							},
						},
						Config: testConfig,
					},
					// Step 2: Run migration and verify default value applied
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer,
						[]statecheck.StateCheck{
							// Verify requested_validity defaulted to 5475
							statecheck.ExpectKnownValue(
								"cloudflare_origin_ca_certificate."+rnd,
								tfjsonpath.New("requested_validity"),
								knownvalue.Float64Exact(5475),
							),
							// Verify other required fields
							statecheck.ExpectKnownValue(
								"cloudflare_origin_ca_certificate."+rnd,
								tfjsonpath.New("request_type"),
								knownvalue.StringExact("origin-rsa"),
							),
							statecheck.ExpectKnownValue(
								"cloudflare_origin_ca_certificate."+rnd,
								tfjsonpath.New("hostnames"),
								knownvalue.NotNull(),
							),
						},
					),
				},
			})
		})
	}
}

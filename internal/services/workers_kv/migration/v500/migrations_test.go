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

// Embed test configs
//
//go:embed testdata/v4_basic.tf
var v4BasicConfig string

//go:embed testdata/v5_basic.tf
var v5BasicConfig string

//go:embed testdata/v4_special_chars.tf
var v4SpecialCharsConfig string

//go:embed testdata/v5_special_chars.tf
var v5SpecialCharsConfig string

//go:embed testdata/v4_empty_value.tf
var v4EmptyValueConfig string

//go:embed testdata/v5_empty_value.tf
var v5EmptyValueConfig string

// TestMigrateWorkersKV_Basic tests migration of a basic Workers KV resource from v4 to v5
// Main change: key field renamed to key_name
func TestMigrateWorkersKV_Basic(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(namespaceName, accountID, kvName, keyName, value string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(namespaceName, accountID, kvName, keyName, value string) string {
				return fmt.Sprintf(v4BasicConfig, namespaceName, accountID, namespaceName, kvName, keyName, value)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(namespaceName, accountID, kvName, keyName, value string) string {
				return fmt.Sprintf(v5BasicConfig, namespaceName, accountID, namespaceName, kvName, keyName, value)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			namespaceName := fmt.Sprintf("tf-test-ns-%s", rnd)
			kvName := fmt.Sprintf("tf-test-kv-%s", rnd)
			keyName := "test_key"
			value := "test_value"
			tmpDir := t.TempDir()
			testConfig := tc.configFn(namespaceName, accountID, kvName, keyName, value)
			sourceVer, targetVer := "v4", "v5"
			if tc.version != acctest.GetLastV4Version() {
				sourceVer, targetVer = "v5", "v5"
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
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
							// Verify the key field was renamed to key_name
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("key_name"), knownvalue.StringExact(keyName)),
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							// Verify namespace_id is preserved
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("namespace_id"), knownvalue.NotNull()),
							// Verify id field is preserved
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("id"), knownvalue.NotNull()),
						},
					),
				},
			})
		})
	}
}

// TestMigrateWorkersKV_SpecialCharacters tests migration with URL-encoded special characters
func TestMigrateWorkersKV_SpecialCharacters(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(namespaceName, accountID, kvName, keyName, value string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(namespaceName, accountID, kvName, keyName, value string) string {
				return fmt.Sprintf(v4SpecialCharsConfig, namespaceName, accountID, namespaceName, kvName, keyName, value)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(namespaceName, accountID, kvName, keyName, value string) string {
				return fmt.Sprintf(v5SpecialCharsConfig, namespaceName, accountID, namespaceName, kvName, keyName, value)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			namespaceName := fmt.Sprintf("tf-test-ns-%s", rnd)
			kvName := fmt.Sprintf("tf-test-kv-%s", rnd)
			// URL-encoded key with special characters
			keyName := "api/token/key"
			value := `{"api_key": "test123", "endpoint": "https://api.example.com"}`
			tmpDir := t.TempDir()
			testConfig := tc.configFn(namespaceName, accountID, kvName, keyName, value)
			sourceVer, targetVer := "v4", "v5"
			if tc.version != acctest.GetLastV4Version() {
				sourceVer, targetVer = "v5", "v5"
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
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
							// Verify the key field was renamed to key_name and special characters preserved
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("key_name"), knownvalue.StringExact(keyName)),
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("value"), knownvalue.StringExact(value)),
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						},
					),
				},
			})
		})
	}
}

// TestMigrateWorkersKV_EmptyValue tests migration with an empty string value
func TestMigrateWorkersKV_EmptyValue(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(namespaceName, accountID, kvName, keyName, value string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(namespaceName, accountID, kvName, keyName, value string) string {
				return fmt.Sprintf(v4EmptyValueConfig, namespaceName, accountID, namespaceName, kvName, keyName, value)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(namespaceName, accountID, kvName, keyName, value string) string {
				return fmt.Sprintf(v5EmptyValueConfig, namespaceName, accountID, namespaceName, kvName, keyName, value)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			namespaceName := fmt.Sprintf("tf-test-ns-%s", rnd)
			kvName := fmt.Sprintf("tf-test-kv-%s", rnd)
			keyName := "empty_key"
			value := ""
			tmpDir := t.TempDir()
			testConfig := tc.configFn(namespaceName, accountID, kvName, keyName, value)
			sourceVer, targetVer := "v4", "v5"
			if tc.version != acctest.GetLastV4Version() {
				sourceVer, targetVer = "v5", "v5"
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
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
							// Verify empty value is preserved
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("key_name"), knownvalue.StringExact(keyName)),
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("value"), knownvalue.StringExact("")),
							statecheck.ExpectKnownValue("cloudflare_workers_kv."+kvName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						},
					),
				},
			})
		})
	}
}

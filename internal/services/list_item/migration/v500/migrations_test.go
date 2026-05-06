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

// setV4ProviderRetryEnv configures the v4 Cloudflare provider's retry and backoff
// settings via environment variables for the duration of the test. The v4 provider
// defaults to 4 retries (CLOUDFLARE_RETRIES=4) with a max backoff of 30s, which is
// insufficient when multiple migration test packages run concurrently and hammer the
// Lists API simultaneously. Bumping retries and max backoff makes the v4 provider
// wait longer and retry more before giving up with "exceeded available rate limit retries".
func setV4ProviderRetryEnv(t *testing.T) {
	t.Helper()
	if os.Getenv("CLOUDFLARE_RETRIES") == "" {
		t.Setenv("CLOUDFLARE_RETRIES", "8")
	}
	if os.Getenv("CLOUDFLARE_MAX_BACKOFF") == "" {
		t.Setenv("CLOUDFLARE_MAX_BACKOFF", "60")
	}
}

// Embed migration test configuration files
//
//go:embed testdata/v4_ip_item.tf
var v4IPItemConfig string

//go:embed testdata/v4_hostname_item.tf
var v4HostnameItemConfig string

//go:embed testdata/v4_redirect_item.tf
var v4RedirectItemConfig string

//go:embed testdata/v5_ip_item.tf
var v5IPItemConfig string

//go:embed testdata/v5_hostname_item.tf
var v5HostnameItemConfig string

//go:embed testdata/v5_redirect_item.tf
var v5RedirectItemConfig string

const listTestPrefix = "tf_test_list_"

// TestMigrateCloudflareListItemWithIP tests migration of a standalone list_item with IP.
// Tests both v4→v5 and v5→v5 (stepping stone) upgrade paths.
func TestMigrateCloudflareListItemWithIP(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, listName, description string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v4IPItemConfig, rnd, accountID, listName, description)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v5IPItemConfig, rnd, accountID, listName, description)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setV4ProviderRetryEnv(t)
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
			listResourceName := "cloudflare_list." + rnd
			itemResourceName := "cloudflare_list_item." + rnd + "_item"
			tmpDir := t.TempDir()
			description := fmt.Sprintf("Test list with IP item %s", rnd)
			testConfig := tc.configFn(rnd, accountID, listName, description)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
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
				}, acctest.MigrationV2TestStepAllowCreate(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
					statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("kind"), knownvalue.StringExact("ip")),
					statecheck.ExpectKnownValue(itemResourceName, tfjsonpath.New("ip"), knownvalue.StringExact("192.0.2.1")),
					statecheck.ExpectKnownValue(itemResourceName, tfjsonpath.New("comment"), knownvalue.StringExact("Test IP item")),
				})...),
			})
		})
	}
}

// TestMigrateCloudflareListItemWithHostname tests migration of a standalone list_item with hostname.
// Tests both v4→v5 (block → SingleNestedAttribute) and v5→v5 (stepping stone) upgrade paths.
// The from_v5 case specifically reproduces the regression described in GitHub issue #7073.
func TestMigrateCloudflareListItemWithHostname(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, listName, description string) string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v4HostnameItemConfig, rnd, accountID, listName, description)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v5HostnameItemConfig, rnd, accountID, listName, description)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setV4ProviderRetryEnv(t)
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
			listResourceName := "cloudflare_list." + rnd
			itemResourceName := "cloudflare_list_item." + rnd + "_item"
			tmpDir := t.TempDir()
			description := fmt.Sprintf("Test list with hostname item %s", rnd)
			testConfig := tc.configFn(rnd, accountID, listName, description)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
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
				}, acctest.MigrationV2TestStepAllowCreate(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
					statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("kind"), knownvalue.StringExact("hostname")),
					statecheck.ExpectKnownValue(
						itemResourceName,
						tfjsonpath.New("hostname").AtMapKey("url_hostname"),
						knownvalue.StringExact("example.com"),
					),
					statecheck.ExpectKnownValue(itemResourceName, tfjsonpath.New("comment"), knownvalue.StringExact("Test hostname item")),
				})...),
			})
		})
	}
}

// TestMigrateCloudflareListItemWithRedirect tests migration of a standalone list_item with redirect.
// Tests both v4→v5 (block → SingleNestedAttribute + string→bool) and v5→v5 (stepping stone) paths.
func TestMigrateCloudflareListItemWithRedirect(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, listName, description string) string
		// v4 redirect source_url has no trailing slash; the state upgrader appends one.
		// v5 config already has the slash.
		expectedSourceURL string
	}{
		{
			name:    "from_v4_latest",
			version: acctest.GetLastV4Version(),
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v4RedirectItemConfig, rnd, accountID, listName, description)
			},
			expectedSourceURL: "https://example.com/old",
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v5RedirectItemConfig, rnd, accountID, listName, description)
			},
			expectedSourceURL: "https://example.com/old/",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setV4ProviderRetryEnv(t)
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
			listResourceName := "cloudflare_list." + rnd
			itemResourceName := "cloudflare_list_item." + rnd + "_item"
			tmpDir := t.TempDir()
			description := fmt.Sprintf("Test list with redirect item %s", rnd)
			testConfig := tc.configFn(rnd, accountID, listName, description)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
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
				}, acctest.MigrationV2TestStepAllowCreate(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
					statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
					statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("kind"), knownvalue.StringExact("redirect")),
					statecheck.ExpectKnownValue(
						itemResourceName,
						tfjsonpath.New("redirect").AtMapKey("source_url"),
						knownvalue.StringExact(tc.expectedSourceURL),
					),
					statecheck.ExpectKnownValue(
						itemResourceName,
						tfjsonpath.New("redirect").AtMapKey("target_url"),
						knownvalue.StringExact("https://example.com/new"),
					),
					statecheck.ExpectKnownValue(
						itemResourceName,
						tfjsonpath.New("redirect").AtMapKey("status_code"),
						knownvalue.Int64Exact(301),
					),
					statecheck.ExpectKnownValue(
						itemResourceName,
						tfjsonpath.New("redirect").AtMapKey("include_subdomains"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						itemResourceName,
						tfjsonpath.New("redirect").AtMapKey("subpath_matching"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						itemResourceName,
						tfjsonpath.New("redirect").AtMapKey("preserve_query_string"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						itemResourceName,
						tfjsonpath.New("redirect").AtMapKey("preserve_path_suffix"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(itemResourceName, tfjsonpath.New("comment"), knownvalue.StringExact("Test redirect item")),
				})...),
			})
		})
	}
}

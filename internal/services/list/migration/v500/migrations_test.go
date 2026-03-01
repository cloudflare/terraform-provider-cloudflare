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

// Migration test version configuration.
// - Last stable v4 release: read from LAST_V4_VERSION env var (set in CI)
// - Current v5 release: auto-updates with releases (internal.PackageVersion)
var currentProviderVersion = internal.PackageVersion

// Embed migration test configuration files
//
//go:embed testdata/v4_ip_items.tf
var v4IPItemsConfig string

//go:embed testdata/v5_ip_items.tf
var v5IPItemsConfig string

//go:embed testdata/v4_asn_items.tf
var v4ASNItemsConfig string

//go:embed testdata/v5_asn_items.tf
var v5ASNItemsConfig string

//go:embed testdata/v4_hostname_items.tf
var v4HostnameItemsConfig string

//go:embed testdata/v5_hostname_items.tf
var v5HostnameItemsConfig string

//go:embed testdata/v4_redirect_items.tf
var v4RedirectItemsConfig string

//go:embed testdata/v5_redirect_items.tf
var v5RedirectItemsConfig string

//go:embed testdata/v4_dynamic_items.tf
var v4DynamicItemsConfig string

//go:embed testdata/v5_dynamic_items.tf
var v5DynamicItemsConfig string

//go:embed testdata/v4_separate_list_items.tf
var v4SeparateListItemsConfig string

const listTestPrefix = "tf_test_list_"

// TestMigrateCloudflareListFromV4WithIPItems tests migration with IP item blocks
func TestMigrateCloudflareListFromV4WithIPItems(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, listName, description string) string
	}{
		{
			name:    "from_v4_latest",
			version: os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v4IPItemsConfig, rnd, accountID, listName, description)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v5IPItemsConfig, rnd, accountID, listName, description)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
			resourceName := "cloudflare_list." + rnd
			tmpDir := t.TempDir()
			description := fmt.Sprintf("Test list %s", rnd)
			testConfig := tc.configFn(rnd, accountID, listName, description)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("ip")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("ip"), knownvalue.StringExact("192.0.2.1")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Test IP 1")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("ip"), knownvalue.StringExact("192.0.2.2")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("comment"), knownvalue.StringExact("Test IP 2")),
					}),
				},
			})
		})
	}
}

// TestMigrateCloudflareListFromV4WithASNItems tests migration with ASN item blocks
func TestMigrateCloudflareListFromV4WithASNItems(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, listName, description string) string
	}{
		{
			name:    "from_v4_latest",
			version: os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v4ASNItemsConfig, rnd, accountID, listName, description)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v5ASNItemsConfig, rnd, accountID, listName, description)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
			resourceName := "cloudflare_list." + rnd
			tmpDir := t.TempDir()
			description := fmt.Sprintf("Test ASN list %s", rnd)
			testConfig := tc.configFn(rnd, accountID, listName, description)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("asn")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("asn"), knownvalue.Int64Exact(12345)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Test ASN 1")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("asn"), knownvalue.Int64Exact(67890)),
					}),
				},
			})
		})
	}
}

// TestMigrateCloudflareListFromV4WithHostnameItems tests migration with hostname item blocks
func TestMigrateCloudflareListFromV4WithHostnameItems(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, listName, description string) string
	}{
		{
			name:    "from_v4_latest",
			version: os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v4HostnameItemsConfig, rnd, accountID, listName, description)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v5HostnameItemsConfig, rnd, accountID, listName, description)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
			resourceName := "cloudflare_list." + rnd
			tmpDir := t.TempDir()
			description := fmt.Sprintf("Test hostname list %s", rnd)
			testConfig := tc.configFn(rnd, accountID, listName, description)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("hostname")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("hostname").AtMapKey("url_hostname"),
							knownvalue.StringExact("example.com"),
						),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Test hostname 1")),
						statecheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("hostname").AtMapKey("url_hostname"),
							knownvalue.StringExact("test.example.com"),
						),
					}),
				},
			})
		})
	}
}

// TestMigrateCloudflareListFromV4WithRedirectItems tests migration with redirect item blocks including boolean conversion
func TestMigrateCloudflareListFromV4WithRedirectItems(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, listName, description string) string
	}{
		{
			name:    "from_v4_latest",
			version: os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v4RedirectItemsConfig, rnd, accountID, listName, description)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v5RedirectItemsConfig, rnd, accountID, listName, description)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
			resourceName := "cloudflare_list." + rnd
			tmpDir := t.TempDir()
			description := fmt.Sprintf("Test redirect list %s", rnd)
			testConfig := tc.configFn(rnd, accountID, listName, description)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("redirect")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("source_url"),
							knownvalue.StringExact("https://example.com/old"),
						),
						statecheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("target_url"),
							knownvalue.StringExact("https://example.com/new"),
						),
						statecheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("status_code"),
							knownvalue.Int64Exact(301),
						),
						// Verify boolean conversions from "enabled"/"disabled" to true/false
						statecheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("include_subdomains"),
							knownvalue.Bool(true),
						),
						statecheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("subpath_matching"),
							knownvalue.Bool(false),
						),
						statecheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("preserve_query_string"),
							knownvalue.Bool(true),
						),
						statecheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("redirect").AtMapKey("preserve_path_suffix"),
							knownvalue.Bool(false),
						),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Test redirect")),
					}),
				},
			})
		})
	}
}

// TestMigrateCloudflareListFromV4WithDynamicItems tests migration with dynamic item blocks
func TestMigrateCloudflareListFromV4WithDynamicItems(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		configFn func(rnd, accountID, listName, description string) string
	}{
		{
			name:    "from_v4_latest",
			version: os.Getenv("LAST_V4_VERSION"),
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v4DynamicItemsConfig, rnd, accountID, listName, description)
			},
		},
		{
			name:    "from_v5",
			version: currentProviderVersion,
			configFn: func(rnd, accountID, listName, description string) string {
				return fmt.Sprintf(v5DynamicItemsConfig, rnd, accountID, listName, description)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			rnd := utils.GenerateRandomResourceName()
			listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
			resourceName := "cloudflare_list." + rnd
			tmpDir := t.TempDir()
			description := fmt.Sprintf("Test dynamic list %s", rnd)
			testConfig := tc.configFn(rnd, accountID, listName, description)
			sourceVer, targetVer := acctest.InferMigrationVersions(tc.version)

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
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
					acctest.MigrationV2TestStep(t, testConfig, tmpDir, tc.version, sourceVer, targetVer, []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("kind"), knownvalue.StringExact("ip")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("ip"), knownvalue.StringExact("10.0.0.1")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("comment"), knownvalue.StringExact("Dynamic IP 1")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("ip"), knownvalue.StringExact("10.0.0.2")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("items").AtSliceIndex(1).AtMapKey("comment"), knownvalue.StringExact("Dynamic IP 2")),
					}),
				},
			})
		})
	}
}

// TestMigrateCloudflareListFromV4WithSeparateListItems tests migration with separate cloudflare_list_item resources.
// In v5, separate list_item resources remain as separate resources (they are NOT merged into the parent list).
// The provider's StateUpgraders handle the state migration for both cloudflare_list and cloudflare_list_item.
func TestMigrateCloudflareListFromV4WithSeparateListItems(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
	listResourceName := "cloudflare_list." + rnd
	item1ResourceName := "cloudflare_list_item." + rnd + "_item1"
	item2ResourceName := "cloudflare_list_item." + rnd + "_item2"
	tmpDir := t.TempDir()
	description := fmt.Sprintf("Test list with separate items %s", rnd)

	v4Config := fmt.Sprintf(v4SeparateListItemsConfig, rnd, accountID, listName, description)

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
						VersionConstraint: os.Getenv("LAST_V4_VERSION"),
					},
				},
				Config: v4Config,
			},
		}, acctest.MigrationV2TestStepAllowCreate(t, v4Config, tmpDir, os.Getenv("LAST_V4_VERSION"), "v4", "v5", []statecheck.StateCheck{
			// Verify the parent list resource
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("kind"), knownvalue.StringExact("ip")),
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("description"), knownvalue.StringExact(description)),
			// Verify separate list_item resources remain as separate resources
			statecheck.ExpectKnownValue(item1ResourceName, tfjsonpath.New("ip"), knownvalue.StringExact("172.16.0.1")),
			statecheck.ExpectKnownValue(item1ResourceName, tfjsonpath.New("comment"), knownvalue.StringExact("Separate item 1")),
			statecheck.ExpectKnownValue(item2ResourceName, tfjsonpath.New("ip"), knownvalue.StringExact("172.16.0.2")),
			statecheck.ExpectKnownValue(item2ResourceName, tfjsonpath.New("comment"), knownvalue.StringExact("Separate item 2")),
		})...),
	})
}

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

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
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

const listTestPrefix = "tf_test_list_"

// TestMigrateCloudflareListItemFromV4WithIP tests migration of a standalone list_item with IP.
// In v5, separate list_item resources remain as separate resources (they are NOT merged into the parent list).
// The provider's StateUpgraders handle the state migration for both cloudflare_list and cloudflare_list_item.
func TestMigrateCloudflareListItemFromV4WithIP(t *testing.T) {
	setV4ProviderRetryEnv(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
	listResourceName := "cloudflare_list." + rnd
	itemResourceName := "cloudflare_list_item." + rnd + "_item"
	tmpDir := t.TempDir()
	description := fmt.Sprintf("Test list with IP item %s", rnd)

	v4Config := fmt.Sprintf(v4IPItemConfig, rnd, accountID, listName, description)

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
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("kind"), knownvalue.StringExact("ip")),
			// Verify list_item remains as a separate resource
			statecheck.ExpectKnownValue(itemResourceName, tfjsonpath.New("ip"), knownvalue.StringExact("192.0.2.1")),
			statecheck.ExpectKnownValue(itemResourceName, tfjsonpath.New("comment"), knownvalue.StringExact("Test IP item")),
		})...),
	})
}

// TestMigrateCloudflareListItemFromV4WithHostname tests migration of a standalone list_item with hostname.
// In v5, separate list_item resources remain as separate resources.
// tf-migrate converts the hostname block to a hostname attribute (block → SingleNestedAttribute).
func TestMigrateCloudflareListItemFromV4WithHostname(t *testing.T) {
	setV4ProviderRetryEnv(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
	listResourceName := "cloudflare_list." + rnd
	itemResourceName := "cloudflare_list_item." + rnd + "_item"
	tmpDir := t.TempDir()
	description := fmt.Sprintf("Test list with hostname item %s", rnd)

	v4Config := fmt.Sprintf(v4HostnameItemConfig, rnd, accountID, listName, description)

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
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("kind"), knownvalue.StringExact("hostname")),
			// Verify list_item remains as a separate resource with hostname attribute
			statecheck.ExpectKnownValue(
				itemResourceName,
				tfjsonpath.New("hostname").AtMapKey("url_hostname"),
				knownvalue.StringExact("example.com"),
			),
			statecheck.ExpectKnownValue(itemResourceName, tfjsonpath.New("comment"), knownvalue.StringExact("Test hostname item")),
		})...),
	})
}

// TestMigrateCloudflareListItemFromV4WithRedirect tests migration of a standalone list_item with redirect.
// In v5, separate list_item resources remain as separate resources.
// tf-migrate converts the redirect block to a redirect attribute (block → SingleNestedAttribute)
// and converts "enabled"/"disabled" strings to true/false booleans.
// Note: The v4 test config uses true/false (not "enabled"/"disabled") because v4.52.7 already uses BoolAttribute.
func TestMigrateCloudflareListItemFromV4WithRedirect(t *testing.T) {
	setV4ProviderRetryEnv(t)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	listName := fmt.Sprintf("%s%s", listTestPrefix, rnd)
	listResourceName := "cloudflare_list." + rnd
	itemResourceName := "cloudflare_list_item." + rnd + "_item"
	tmpDir := t.TempDir()
	description := fmt.Sprintf("Test list with redirect item %s", rnd)

	v4Config := fmt.Sprintf(v4RedirectItemConfig, rnd, accountID, listName, description)

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
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("name"), knownvalue.StringExact(listName)),
			statecheck.ExpectKnownValue(listResourceName, tfjsonpath.New("kind"), knownvalue.StringExact("redirect")),
			// Verify list_item remains as a separate resource with redirect attribute
			statecheck.ExpectKnownValue(
				itemResourceName,
				tfjsonpath.New("redirect").AtMapKey("source_url"),
				knownvalue.StringExact("https://example.com/old"),
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
}

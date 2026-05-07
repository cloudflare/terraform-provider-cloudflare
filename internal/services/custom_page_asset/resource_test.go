package custom_page_asset_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareCustomPageAsset_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_custom_page_asset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPageAssetConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Basic test asset")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://example.com/error.html")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
				},
			},
			{
				Config: testAccCustomPageAssetConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Basic test asset")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://example.com/error.html")),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: acctest.LogResourceDrift,
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
		},
	})
}

func TestAccCloudflareCustomPageAsset_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_custom_page_asset." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPageAssetConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Basic test asset")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://example.com/error.html")),
				},
			},
			{
				Config: testAccCustomPageAssetUpdateConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("description"), knownvalue.StringExact("Updated test asset")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://example.com/updated-error.html")),
				},
			},
		},
	})
}

func testAccCustomPageAssetConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID)
}

func testAccCustomPageAssetUpdateConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("update.tf", rnd, accountID)
}

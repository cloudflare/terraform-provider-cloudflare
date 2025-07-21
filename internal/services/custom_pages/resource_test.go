package custom_pages_test

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


// Test account-level custom pages with basic configuration
func TestAccCloudflareCustomPages_AccountBasic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "default", "http://www.example.com/custom_page"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("basic_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("http://www.example.com/custom_page")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "default", "http://www.example.com/custom_page_2"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("basic_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("http://www.example.com/custom_page_2")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"state", "url"},
			},
		},
	})
}

// Test zone-level custom pages with basic configuration
func TestAccCloudflareCustomPages_ZoneBasic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "500_errors", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("500_errors")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "500_errors", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("500_errors")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("zones/%s/", zoneID),
				ImportStateVerifyIgnore: []string{"state", "url"},
			},
		},
	})
}

// Test account-level custom pages with multiple page types
func TestAccCloudflareCustomPages_AccountMultiplePages(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName1 := "cloudflare_custom_pages." + rnd + "_challenge"
	resourceName2 := "cloudflare_custom_pages." + rnd + "_block"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesAccountMultipleConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					// Challenge page checks
					statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_challenge")),
					statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("url"), knownvalue.StringExact("")),
					// Block page checks
					statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_block")),
					statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName2, tfjsonpath.New("url"), knownvalue.StringExact("")),
				},
			},
			{
				ResourceName:            resourceName1,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"state", "url"},
			},
			{
				ResourceName:            resourceName2,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"state", "url"},
			},
		},
	})
}


func testAccCustomPagesAccountConfig(rnd, accountID, identifier, state, url string) string {
	return acctest.LoadTestCase("account_basic.tf", rnd, accountID, identifier, state, url)
}

func testAccCustomPagesZoneConfig(rnd, zoneID, identifier, state, url string) string {
	return acctest.LoadTestCase("zone_basic.tf", rnd, zoneID, identifier, state, url)
}

func testAccCustomPagesAccountMultipleConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("account_multiple.tf", rnd, accountID)
}

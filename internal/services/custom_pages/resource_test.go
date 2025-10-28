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

// NOTE: https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/ is used for all custom page types
// that require a ::CF_WIDGET_BOX:: (including country and IP challenges), not just for waf challenges

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
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("basic_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("basic_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
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

func TestAccCloudflareCustomPages_ZoneCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "500_errors", "customized", "https://custom-pages-basic.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("500_errors")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-basic.terraform-provider-acceptance-testing.workers.dev/")),
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

// Test account-level custom pages with customized state
func TestAccCloudflareCustomPages_AccountCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "waf_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "waf_block", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
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
					statecheck.ExpectKnownValue(resourceName1, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_block")),
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

// Test zone-level custom page with ip_block identifier
func TestAccCloudflareCustomPages_ZoneIpBlock(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "ip_block", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("ip_block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "ip_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("ip_block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
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

// Test zone-level custom page with country_challenge identifier
func TestAccCloudflareCustomPages_ZoneCountryChallenge(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "country_challenge", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("country_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "country_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("country_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
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

// Test zone-level custom page with 1000_errors identifier
func TestAccCloudflareCustomPages_Zone1000Errors(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "1000_errors", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("1000_errors")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "1000_errors", "customized", "https://custom-pages-1000-errors.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("1000_errors")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-1000-errors.terraform-provider-acceptance-testing.workers.dev/")),
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

// Test zone-level custom page with managed_challenge identifier
func TestAccCloudflareCustomPages_ZoneManagedChallenge(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "managed_challenge", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("managed_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "managed_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("managed_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
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

// Test zone-level custom page with ratelimit_block identifier
func TestAccCloudflareCustomPages_ZoneRatelimitBlock(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "ratelimit_block", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("ratelimit_block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.Null()),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: acctest.LogResourceDrift,
				},
			},
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "ratelimit_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("ratelimit_block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
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

// Test account-level custom page with basic_challenge identifier (customized)
func TestAccCloudflareCustomPages_AccountBasicChallengeCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("basic_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesAccountConfig(rnd, accountID, "basic_challenge", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("basic_challenge")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("")),
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

// Test zone-level custom page with waf_block identifier (customized)
func TestAccCloudflareCustomPages_ZoneWafBlockCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "waf_block", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_block")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "waf_block", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("waf_block")),
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

// Test zone-level custom page with under_attack identifier (customized)
func TestAccCloudflareCustomPages_ZoneUnderAttackCustomized(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_custom_pages." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "under_attack", "customized", "https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("under_attack")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("customized")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact("https://custom-pages-waf-challenge.terraform-provider-acceptance-testing.workers.dev/")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.Null()),
				},
			},
			{
				Config: testAccCustomPagesZoneConfig(rnd, zoneID, "under_attack", "default", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("identifier"), knownvalue.StringExact("under_attack")),
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

func testAccCustomPagesAccountConfig(rnd, accountID, identifier, state, url string) string {
	return acctest.LoadTestCase("account_basic.tf", rnd, accountID, identifier, state, url)
}

func testAccCustomPagesZoneConfig(rnd, zoneID, identifier, state, url string) string {
	return acctest.LoadTestCase("zone_basic.tf", rnd, zoneID, identifier, state, url)
}

func testAccCustomPagesAccountMultipleConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("account_multiple.tf", rnd, accountID)
}

package r2_custom_domain_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareR2CustomDomain_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domainName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_r2_custom_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2CustomDomainConfig(rnd, accountID, zoneID, domainName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(domainName)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.Null()),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ciphers"), knownvalue.Null()),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.StringExact(domainName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ciphers"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestAccCloudflareR2CustomDomain_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domainName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_r2_custom_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2CustomDomainConfig(rnd, accountID, zoneID, domainName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.Null()),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ciphers"), knownvalue.Null()),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ciphers"), knownvalue.Null()),
				},
			},
			{
				Config: testAccR2CustomDomainUpdateConfig(rnd, accountID, zoneID, domainName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.StringExact("1.2")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ciphers"), knownvalue.ListSizeExact(2)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.StringExact("1.2")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ciphers"), knownvalue.ListSizeExact(2)),
				},
			},
		},
	})
}

func TestAccCloudflareR2CustomDomain_JurisdictionEU(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domainName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_r2_custom_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2CustomDomainJurisdictionEUConfig(rnd, accountID, zoneID, domainName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.StringExact("1.3")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.StringExact("1.3")),
				},
			},
		},
	})
}

func TestAccCloudflareR2CustomDomain_JurisdictionFedramp(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domainName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_r2_custom_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2CustomDomainJurisdictionFedrampConfig(rnd, accountID, zoneID, domainName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("fedramp")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.StringExact("1.3")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.StringExact("1.3")),
				},
			},
		},
	})
}

func TestAccCloudflareR2CustomDomain_Minimal(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domainName := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_r2_custom_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2CustomDomainMinimalConfig(rnd, accountID, zoneID, domainName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.Null()),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("ciphers"), knownvalue.Null()),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("min_tls"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("ciphers"), knownvalue.Null()),
				},
			},
		},
	})
}

func testAccR2CustomDomainConfig(rnd, accountID, zoneID, domainName string) string {
	return acctest.LoadTestCase("r2custom_basic.tf", rnd, accountID, zoneID, domainName)
}

func testAccR2CustomDomainUpdateConfig(rnd, accountID, zoneID, domainName string) string {
	return acctest.LoadTestCase("r2custom_update.tf", rnd, accountID, zoneID, domainName)
}

func testAccR2CustomDomainJurisdictionEUConfig(rnd, accountID, zoneID, domainName string) string {
	return acctest.LoadTestCase("r2custom_jurisdiction_eu.tf", rnd, accountID, zoneID, domainName)
}

func testAccR2CustomDomainMinimalConfig(rnd, accountID, zoneID, domainName string) string {
	return acctest.LoadTestCase("r2custom_minimal.tf", rnd, accountID, zoneID, domainName)
}

func testAccR2CustomDomainJurisdictionFedrampConfig(rnd, accountID, zoneID, domainName string) string {
	return acctest.LoadTestCase("r2custom_jurisdiction_fedramp.tf", rnd, accountID, zoneID, domainName)
}

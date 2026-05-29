package custom_csr_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/custom_csrs"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_custom_csr", &resource.Sweeper{
		Name: "cloudflare_custom_csr",
		F:    testSweepCustomCsrs,
	})
}

func testSweepCustomCsrs(_ string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping custom CSR sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	csrs, err := client.CustomCsrs.List(ctx, custom_csrs.CustomCsrListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list custom CSRs: %s", err))
		return nil
	}

	for _, csr := range csrs.Result {
		if !utils.ShouldSweepResource(csr.ID) {
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleting custom CSR: %s", csr.ID))
		_, err := client.CustomCsrs.Delete(ctx, csr.ID, custom_csrs.CustomCsrDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete custom CSR %s: %s", csr.ID, err))
		}
	}

	return nil
}

func TestAccCloudflareCustomCsr_ZoneWorkflow(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_custom_csr." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCustomCsrDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create zone-scoped CSR with default key_type (rsa2048)
			{
				Config: testAccCustomCsrZoneConfig(rnd, zoneID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("country"), knownvalue.StringExact("US")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("state"), knownvalue.StringExact("California")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("locality"), knownvalue.StringExact("San Francisco")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("organization"), knownvalue.StringExact("Terraform Test")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("common_name"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_type"), knownvalue.StringExact("rsa2048")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("csr"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 2: Re-apply same config — verify no drift
			{
				Config: testAccCustomCsrZoneConfig(rnd, zoneID, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 3: Import by zone
			{
				ResourceName: resourceName,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("resource %s not found", resourceName)
					}
					return fmt.Sprintf("zones/%s/%s", zoneID, rs.Primary.ID), nil
				},
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sans"},
			},
		},
	})
}

func TestAccCloudflareCustomCsr_ZoneP256(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_custom_csr." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCustomCsrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomCsrZoneP256Config(rnd, zoneID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_type"), knownvalue.StringExact("p256v1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("csr"), knownvalue.NotNull()),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccCloudflareCustomCsr_AccountWorkflow(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	resourceName := "cloudflare_custom_csr." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCustomCsrDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create account-scoped CSR with name and description
			{
				Config: testAccCustomCsrAccountConfig(rnd, accountID, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("country"), knownvalue.StringExact("US")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("common_name"), knownvalue.StringExact(domain)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("key_type"), knownvalue.StringExact("rsa2048")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("csr"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 2: Re-apply same config — verify no drift
			{
				Config: testAccCustomCsrAccountConfig(rnd, accountID, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 3: Import by account
			{
				ResourceName: resourceName,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("resource %s not found", resourceName)
					}
					return fmt.Sprintf("accounts/%s/%s", accountID, rs.Primary.ID), nil
				},
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sans", "name", "description"},
			},
		},
	})
}

func testAccCheckCustomCsrDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_csr" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]

		params := custom_csrs.CustomCsrGetParams{}
		if accountID != "" {
			params.AccountID = cloudflare.F(accountID)
		} else {
			params.ZoneID = cloudflare.F(zoneID)
		}

		_, err := client.CustomCsrs.Get(
			context.Background(),
			rs.Primary.ID,
			params,
		)
		if err == nil {
			return fmt.Errorf("custom CSR %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCustomCsrZoneConfig(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("customcsrzone.tf", rnd, zoneID, domain)
}

func testAccCustomCsrZoneP256Config(rnd, zoneID, domain string) string {
	return acctest.LoadTestCase("customcsrzonep256.tf", rnd, zoneID, domain)
}

func testAccCustomCsrAccountConfig(rnd, accountID, domain string) string {
	return acctest.LoadTestCase("customcsraccount.tf", rnd, accountID, domain)
}

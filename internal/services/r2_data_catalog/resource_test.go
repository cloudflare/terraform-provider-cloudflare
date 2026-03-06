package r2_data_catalog_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/r2_data_catalog"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareR2DataCatalog_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_data_catalog." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareR2DataCatalogDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccR2DataCatalogConfig(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     fmt.Sprintf("%s/%s", accountID, rnd),
				ImportStateVerify: true,
				// Enable (POST) returns only {id, name}; Get (GET) returns all fields.
				// These 4 computed-only fields are null after Create but populated after Import.
				ImportStateVerifyIgnore: []string{
					"bucket",
					"credential_status",
					"maintenance_config",
					"status",
				},
			},
		},
	})
}

func testAccCheckCloudflareR2DataCatalogDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_r2_data_catalog" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		bucketName := rs.Primary.Attributes["bucket_name"]
		res, err := client.R2DataCatalog.Get(
			context.Background(),
			bucketName,
			r2_data_catalog.R2DataCatalogGetParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err != nil {
			// API error means resource is gone, which is fine
			continue
		}
		if res.Status == r2_data_catalog.R2DataCatalogGetResponseStatusActive {
			return fmt.Errorf("r2 data catalog %s still active", bucketName)
		}
	}

	return nil
}

func testAccR2DataCatalogConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2dc_enable.tf", rnd, accountID)
}

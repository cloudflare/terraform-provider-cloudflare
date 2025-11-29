package cloud_connector_rules_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/cloud_connector"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_cloud_connector_rules", &resource.Sweeper{
		Name: "cloudflare_cloud_connector_rules",
		F:    testSweepCloudflareCloudConnectorRules,
	})
}

func testSweepCloudflareCloudConnectorRules(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping cloud connector rules sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	tflog.Info(ctx, fmt.Sprintf("Clearing cloud connector rules for zone: %s", zoneID))
	_, err := client.CloudConnector.Rules.Update(
		ctx,
		cloud_connector.RuleUpdateParams{
			ZoneID: cloudflare.F(zoneID),
		},
		option.WithRequestBody("application/json", []byte(`[]`)),
	)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to clear cloud connector rules: %s", err))
		return fmt.Errorf("failed to clear cloud connector rules: %w", err)
	}

	tflog.Info(ctx, "Cleared cloud connector rules")
	return nil
}

func TestAccCloudflareCloudConnectorRules(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	resourceName := "cloudflare_cloud_connector_rules." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCloudConnectorRules(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "some description 1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.provider", "aws_s3"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.parameters.host", "mystorage1.s3.ams.amazonaws.com"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "some description 2"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.provider", "aws_s3"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.parameters.host", "mystorage2.s3.ams.amazonaws.com"),

					resource.TestCheckResourceAttr(resourceName, "rules.2.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.description", "some description 3"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.provider", "aws_s3"),
					resource.TestCheckResourceAttr(resourceName, "rules.2.parameters.host", "mystorage3.s3.ams.amazonaws.com"),
				),
			},
			{
				Config: testAccCheckCloudflareCloudConnectorRulesRemovedRule(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "2"),

					resource.TestCheckResourceAttr(resourceName, "rules.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "some description 2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.provider", "aws_s3"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.parameters.host", "mystorage2.s3.ams.amazonaws.com"),

					resource.TestCheckResourceAttr(resourceName, "rules.1.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.expression", "true"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "some description 3"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.provider", "aws_s3"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.parameters.host", "mystorage3.s3.ams.amazonaws.com"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCloudConnectorRules(rnd, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_cloud_connector_rules" "%[1]s" {
	zone_id  = "%[2]s"
	rules = [{
		enabled = true
		expression = "true"
		provider = "aws_s3"
		description = "some description 1"
		parameters = {
			host = "mystorage1.s3.ams.amazonaws.com"
		}
	},
	{
		enabled = true
		expression = "true"
		provider = "aws_s3"
		description = "some description 2"
		parameters ={
			host = "mystorage2.s3.ams.amazonaws.com"
		}
	},
	{
		enabled = true
		expression = "true"
		provider = "aws_s3"
		description = "some description 3"
		parameters = {
			host = "mystorage3.s3.ams.amazonaws.com"
		}
	}]
  }`, rnd, zoneID)
}

func testAccCheckCloudflareCloudConnectorRulesRemovedRule(rnd, zoneID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_cloud_connector_rules" "%[1]s" {
	zone_id  = "%[2]s"
	rules = [{
		enabled = true
		expression = "true"
		provider = "aws_s3"
		description = "some description 2"
		parameters = {
			host = "mystorage2.s3.ams.amazonaws.com"
		}
	},
	{
		enabled = true
		expression = "true"
		provider = "aws_s3"
		description = "some description 3"
		parameters = {
			host = "mystorage3.s3.ams.amazonaws.com"
		}
	}]
  }`, rnd, zoneID)
}

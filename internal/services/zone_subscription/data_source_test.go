package zone_subscription_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareZoneSubscriptionDataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_zone_subscription.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneSubscriptionDataSourceConfig(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("id"), knownvalue.NotNull()), // ID is subscription ID, not zone ID
					// Verify computed attributes exist
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("currency"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("state"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("price"), knownvalue.NotNull()),
					// frequency should exist for most subscriptions
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("frequency"), knownvalue.NotNull()),
					// current_period_end and current_period_start might be null for enterprise plans
					// rate_plan should exist with nested attributes
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rate_plan"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rate_plan").AtMapKey("id"), knownvalue.StringExact("enterprise")), // Enterprise rate plan ID
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rate_plan").AtMapKey("currency"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rate_plan").AtMapKey("externally_managed"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rate_plan").AtMapKey("is_contract"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rate_plan").AtMapKey("public_name"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rate_plan").AtMapKey("scope"), knownvalue.NotNull()),
					// sets can be null for some plans
				},
			},
		},
	})
}

func testAccCloudflareZoneSubscriptionDataSourceConfig(rnd, zoneID string) string {
	return acctest.LoadTestCase("datasource_basic.tf", rnd, zoneID)
}

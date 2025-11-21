package argo_tiered_caching_test

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

func TestAccCloudflareArgoTieredCachingDataSource_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_argo_tiered_caching.%s", rnd)
	dataSourceName := fmt.Sprintf("data.cloudflare_argo_tiered_caching.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_AccountID(t)
			acctest.TestAccPreCheck_Credentials(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareArgoTieredCachingDataSourceBasic(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					// Check the resource attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),

					// Check the data source attributes
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("value"), knownvalue.StringExact("on")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("editable"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccCheckCloudflareArgoTieredCachingDataSourceBasic(zoneID, name string) string {
	return acctest.LoadTestCase("datasource_basic.tf", zoneID, name)
}

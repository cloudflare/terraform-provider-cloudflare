package workers_custom_domain_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareWorkersCustomDomainsDataSource_List(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	hostname := rnd + "." + zoneName
	resourceName := "cloudflare_workers_custom_domain." + rnd
	dataSourceName := "data.cloudflare_workers_custom_domains." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersCustomDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkersCustomDomainsListDataSourceConfig(rnd, accountID, hostname, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					// Check the resource was created properly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service"), knownvalue.StringExact("mute-truth-fdb1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),

					// Check the list data source returns results
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("service"), knownvalue.StringExact("mute-truth-fdb1")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("zone_id"), knownvalue.StringExact(zoneID)),
				},
			},
		},
	})
}

func testAccCloudflareWorkersCustomDomainsListDataSourceConfig(rnd, accountID, hostname, zoneID string) string {
	return acctest.LoadTestCase("list_datasource.tf", rnd, accountID, hostname, zoneID)
}

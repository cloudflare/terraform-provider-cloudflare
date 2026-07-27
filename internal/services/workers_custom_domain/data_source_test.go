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

func TestAccCloudflareWorkersCustomDomainDataSource_ByID(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	hostname := rnd + "." + zoneName
	resourceName := "cloudflare_workers_custom_domain." + rnd
	dataSourceName := "data.cloudflare_workers_custom_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersCustomDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkersCustomDomainDataSourceByIDConfig(rnd, accountID, hostname, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					// Check the resource was created properly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service"), knownvalue.StringExact("mute-truth-fdb1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),

					// Check the data source fetches the domain correctly by ID
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("service"), knownvalue.StringExact("mute-truth-fdb1")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("environment"), knownvalue.StringExact("production")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("zone_name"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("domain_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareWorkersCustomDomainDataSource_ByFilter(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	hostname := rnd + "." + zoneName
	resourceName := "cloudflare_workers_custom_domain." + rnd
	dataSourceName := "data.cloudflare_workers_custom_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersCustomDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareWorkersCustomDomainDataSourceByFilterConfig(rnd, accountID, hostname, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					// Check the resource was created properly
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("service"), knownvalue.StringExact("mute-truth-fdb1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),

					// Check the data source fetches the domain correctly by filter
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("hostname"), knownvalue.StringExact(hostname)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("service"), knownvalue.StringExact("mute-truth-fdb1")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("environment"), knownvalue.StringExact("production")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("zone_name"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("domain_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccCloudflareWorkersCustomDomainDataSourceByIDConfig(rnd, accountID, hostname, zoneID string) string {
	return acctest.LoadTestCase("datasource_by_id.tf", rnd, accountID, hostname, zoneID)
}

func testAccCloudflareWorkersCustomDomainDataSourceByFilterConfig(rnd, accountID, hostname, zoneID string) string {
	return acctest.LoadTestCase("datasource_by_filter.tf", rnd, accountID, hostname, zoneID)
}

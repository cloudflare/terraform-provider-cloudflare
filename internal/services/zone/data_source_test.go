package zone_test

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

func TestAccCloudflareZoneDataSource_ByZoneID(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_zone.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDataSourceConfig_ByZoneID(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					// Core attributes
					// Note: zone_id is a path parameter and not saved in state
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("paused"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("type"), knownvalue.NotNull()),

					// Timestamps
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),

					// Name servers
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name_servers"), knownvalue.NotNull()),

					// Account information
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account").AtMapKey("name"), knownvalue.NotNull()),

					// Meta information
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("meta"), knownvalue.NotNull()),
					// Note: cdn_only and dns_only may be null for certain zone types
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("meta").AtMapKey("page_rule_quota"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("meta").AtMapKey("phishing_detected"), knownvalue.NotNull()),

					// Owner information
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("owner"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("owner").AtMapKey("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("owner").AtMapKey("type"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareZoneDataSource_ByName(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_zone.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDataSourceConfig_ByName(rnd, zoneName),
				ConfigStateChecks: []statecheck.StateCheck{
					// Core attributes
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("status"), knownvalue.StringExact("active")),

					// Account information
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),

					// Name servers should be present for active zones
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name_servers"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareZoneDataSource_ByNameWithFilter(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_zone.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDataSourceConfig_WithFilter(rnd, zoneName, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					// Core attributes
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("status"), knownvalue.StringExact("active")),

					// Account filter verification
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

func TestAccCloudflareZoneDataSource_FilterByStatus(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	dataSourceName := fmt.Sprintf("data.cloudflare_zone.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDataSourceConfig_FilterByStatus(rnd, zoneName, "active"),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify the zone returned has the requested status
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("status"), knownvalue.StringExact("active")),
				},
			},
		},
	})
}

func TestAccCloudflareZoneDataSource_FullZoneAttributes(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_zone.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDataSourceConfig_ByZoneID(rnd, zoneID),
				ConfigStateChecks: []statecheck.StateCheck{
					// All possible attributes
					// Note: zone_id is a path parameter and not saved in state
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("paused"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("development_mode"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name_servers"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("meta"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("owner"), knownvalue.NotNull()),
					// Optional attributes that may or may not be present
					// original_dnshost, original_name_servers, original_registrar may be null
					// verification_key is only present for partial zones
					// vanity_name_servers only for Business/Enterprise
					// tenant and tenant_unit may be null for non-tenant zones
					// cname_suffix is only for tenants
				},
			},
		},
	})
}

// Helper functions to generate test configurations
func testAccCloudflareZoneDataSourceConfig_ByZoneID(rnd, zoneID string) string {
	return acctest.LoadTestCase("datasource_zone_by_id.tf", rnd, zoneID)
}

func testAccCloudflareZoneDataSourceConfig_ByName(rnd, zoneName string) string {
	return acctest.LoadTestCase("datasource_zone_by_name.tf", rnd, zoneName)
}

func testAccCloudflareZoneDataSourceConfig_WithFilter(rnd, zoneName, accountID string) string {
	return acctest.LoadTestCase("datasource_zone_with_filter.tf", rnd, zoneName, accountID)
}

func testAccCloudflareZoneDataSourceConfig_FilterByStatus(rnd, zoneName, status string) string {
	return acctest.LoadTestCase("datasource_zone_filter_status.tf", rnd, zoneName, status)
}
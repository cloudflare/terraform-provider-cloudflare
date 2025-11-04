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

func TestAccCloudflareZonesDataSource_Basic(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_zones.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesDataSourceConfig_Basic(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify that we get at least one zone
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result"), knownvalue.NotNull()),
					// Check account filter
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

func TestAccCloudflareZonesDataSource_FilterByName(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	dataSourceName := fmt.Sprintf("data.cloudflare_zones.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesDataSourceConfig_FilterByName(rnd, zoneName),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify filter is applied
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name"), knownvalue.StringExact(zoneName)),
					// Verify we get results
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result"), knownvalue.NotNull()),
					// Check that first result matches the filter
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact(zoneName)),
				},
			},
		},
	})
}

func TestAccCloudflareZonesDataSource_FilterByStatus(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_zones.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesDataSourceConfig_FilterByStatus(rnd, accountID, "active"),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify filter is applied
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("status"), knownvalue.StringExact("active")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
					// Verify we get results
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result"), knownvalue.NotNull()),
					// Check that first result has active status
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("status"), knownvalue.StringExact("active")),
				},
			},
		},
	})
}

func TestAccCloudflareZonesDataSource_FilterByNamePattern(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_zones.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesDataSourceConfig_FilterByNamePattern(rnd, accountID, "contains:."),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify filter is applied
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name"), knownvalue.StringExact("contains:.")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
					// Verify we get results (all zones contain a dot)
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareZonesDataSource_OrderAndDirection(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_zones.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesDataSourceConfig_OrderAndDirection(rnd, accountID, "name", "asc"),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify order and direction are set
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("order"), knownvalue.StringExact("name")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("direction"), knownvalue.StringExact("asc")),
					// Verify we get results
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareZonesDataSource_MaxItems(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_zones.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesDataSourceConfig_MaxItems(rnd, accountID, 1),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify max_items is set
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("max_items"), knownvalue.Int64Exact(1)),
					// Verify we get results
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareZonesDataSource_MatchAny(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	dataSourceName := fmt.Sprintf("data.cloudflare_zones.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesDataSourceConfig_MatchAny(rnd, "active", "pending"),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify match is set to any
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("match"), knownvalue.StringExact("any")),
					// Verify we get results
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareZonesDataSource_CompleteZoneAttributes(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	dataSourceName := fmt.Sprintf("data.cloudflare_zones.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZonesDataSourceConfig_FilterByName(rnd, zoneName),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify we get the zone
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result"), knownvalue.NotNull()),
					// Check first zone has all expected attributes
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact(zoneName)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("status"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("paused"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("development_mode"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("name_servers"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("modified_on"), knownvalue.NotNull()),
					// Account nested object
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("account"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("account").AtMapKey("id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("account").AtMapKey("name"), knownvalue.NotNull()),
					// Meta nested object
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("meta"), knownvalue.NotNull()),
					// Note: cdn_only and dns_only may be null for certain zone types
					// Owner nested object
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("owner"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("owner").AtMapKey("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result").AtSliceIndex(0).AtMapKey("owner").AtMapKey("type"), knownvalue.NotNull()),
				},
			},
		},
	})
}

// Helper functions to generate test configurations
func testAccCloudflareZonesDataSourceConfig_Basic(rnd, accountID string) string {
	return acctest.LoadTestCase("datasource_zones_basic.tf", rnd, accountID)
}

func testAccCloudflareZonesDataSourceConfig_FilterByName(rnd, zoneName string) string {
	return acctest.LoadTestCase("datasource_zones_filter_name.tf", rnd, zoneName)
}

func testAccCloudflareZonesDataSourceConfig_FilterByStatus(rnd, accountID, status string) string {
	return acctest.LoadTestCase("datasource_zones_filter_status.tf", rnd, accountID, status)
}

func testAccCloudflareZonesDataSourceConfig_FilterByNamePattern(rnd, accountID, namePattern string) string {
	return acctest.LoadTestCase("datasource_zones_filter_pattern.tf", rnd, accountID, namePattern)
}

func testAccCloudflareZonesDataSourceConfig_OrderAndDirection(rnd, accountID, order, direction string) string {
	return acctest.LoadTestCase("datasource_zones_order.tf", rnd, accountID, order, direction)
}

func testAccCloudflareZonesDataSourceConfig_MaxItems(rnd, accountID string, maxItems int) string {
	return acctest.LoadTestCase("datasource_zones_max_items.tf", rnd, accountID, maxItems)
}

func testAccCloudflareZonesDataSourceConfig_MatchAny(rnd, status1, status2 string) string {
	return acctest.LoadTestCase("datasource_zones_match_any.tf", rnd, status1, status2)
}
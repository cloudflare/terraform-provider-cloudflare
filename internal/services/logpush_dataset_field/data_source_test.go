package logpush_dataset_field_test

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

func TestAccCloudflareLogpushDatasetFieldDataSource_AccountID(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_logpush_dataset_field.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareLogpushDatasetFieldDataSourceConfig_AccountID(rnd, accountID, "http_requests"),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify the input parameters are set correctly
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("dataset_id"), knownvalue.StringExact("http_requests")),
					// Verify zone_id is not set (mutually exclusive)
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("zone_id"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestAccCloudflareLogpushDatasetFieldDataSource_ZoneID(t *testing.T) {
	t.Parallel()
	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	dataSourceName := fmt.Sprintf("data.cloudflare_logpush_dataset_field.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareLogpushDatasetFieldDataSourceConfig_ZoneID(rnd, zoneID, "firewall_events"),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify the input parameters are set correctly
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("dataset_id"), knownvalue.StringExact("firewall_events")),
					// Verify account_id is not set (mutually exclusive)
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account_id"), knownvalue.Null()),
				},
			},
		},
	})
}

// Helper functions to generate test configurations
func testAccCloudflareLogpushDatasetFieldDataSourceConfig_AccountID(rnd, accountID, datasetID string) string {
	return fmt.Sprintf(`
data "cloudflare_logpush_dataset_field" "%[1]s" {
  account_id = "%[2]s"
  dataset_id = "%[3]s"
}
`, rnd, accountID, datasetID)
}

func testAccCloudflareLogpushDatasetFieldDataSourceConfig_ZoneID(rnd, zoneID, datasetID string) string {
	return fmt.Sprintf(`
data "cloudflare_logpush_dataset_field" "%[1]s" {
  zone_id    = "%[2]s"
  dataset_id = "%[3]s"
}
`, rnd, zoneID, datasetID)
}

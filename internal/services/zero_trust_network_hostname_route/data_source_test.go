package zero_trust_network_hostname_route_test

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

func TestAccCloudflareZeroTrustNetworkHostnameRouteDataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_zero_trust_network_hostname_route." + rnd
	dataSourceName := "data.cloudflare_zero_trust_network_hostname_route." + rnd
	acctest.TestAccPreCheck_AccountID(t)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccZeroTrustNetworkHostnameRouteDataSourceConfig(rnd, accountID, "datasource"),
				ConfigStateChecks: []statecheck.StateCheck{
					// Check resource attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("hostname"), knownvalue.StringExact("datasource.test.example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("comment"), knownvalue.StringExact(fmt.Sprintf("Test hostname route for tf-acctest-%s", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),

					// Check data source attributes match resource attributes
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("hostname"), knownvalue.StringExact("datasource.test.example.com")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("comment"), knownvalue.StringExact(fmt.Sprintf("Test hostname route for tf-acctest-%s", rnd))),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("hostname_route_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("created_at"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func testAccZeroTrustNetworkHostnameRouteDataSourceConfig(rnd, accountID string, prefix string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID, prefix)
}

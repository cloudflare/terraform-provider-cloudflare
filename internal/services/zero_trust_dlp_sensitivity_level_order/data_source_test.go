package zero_trust_dlp_sensitivity_level_order_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

// TestAccCloudflareZeroTrustDLPSensitivityLevelOrderDataSource_Basic verifies
// that the data source can read an order configured by the resource. We apply
// the resource first, then a config that references the same group via the data
// source, and confirm the data source's level_ids match the resource's.
func TestAccCloudflareZeroTrustDLPSensitivityLevelOrderDataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resourceName := fmt.Sprintf("cloudflare_zero_trust_dlp_sensitivity_level_order.%s", rnd)
	dataSourceName := fmt.Sprintf("data.cloudflare_zero_trust_dlp_sensitivity_level_order.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("datasource_basic.tf", rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					// The data source's level_ids must match the resource's.
					statecheck.CompareValuePairs(
						resourceName, tfjsonpath.New("level_ids"),
						dataSourceName, tfjsonpath.New("level_ids"),
						compare.ValuesSame()),
				},
			},
		},
	})
}

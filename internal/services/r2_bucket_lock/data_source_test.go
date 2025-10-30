package r2_bucket_lock_test

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

func TestAccCloudflareR2BucketLockDataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := "data.cloudflare_r2_bucket_lock." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLockDataSourceConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("datasource-test-rule")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("prefix"), knownvalue.StringExact("test-data/")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("type"), knownvalue.StringExact("Age")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("max_age_seconds"), knownvalue.Int64Exact(604800)),
				},
			},
		},
	})
}

func testAccR2BucketLockDataSourceConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("datasource_basic.tf", rnd, accountID)
}

package r2_bucket_cors_test

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

func TestAccCloudflareR2BucketCorsDataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := "data.cloudflare_r2_bucket_cors." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketCorsDataSourceConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules"), knownvalue.NotNull()),
						// Validate the CORS rule structure
						statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("example-cors-rule")),
						statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.StringExact("GET"),
							knownvalue.StringExact("POST"),
						})),
						statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.StringExact("https://example.com"),
							knownvalue.StringExact("https://sub.example.com"),
						})),
						statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("headers"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.StringExact("Content-Type"),
							knownvalue.StringExact("Authorization"),
						})),
						statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expose_headers"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.StringExact("ETag"),
							knownvalue.StringExact("Content-Length"),
						})),
						statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(3600)),
					},
			},
		},
	})
}

func testAccR2BucketCorsDataSourceConfig(rnd string, accountID string) string {
	return acctest.LoadTestCase("datasource_basic.tf", rnd, accountID)
}

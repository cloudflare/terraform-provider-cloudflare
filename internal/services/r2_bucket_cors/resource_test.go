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

func TestAccCloudflareR2BucketCors_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_cors." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketCorsConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("rule1")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("GET"), knownvalue.StringExact("POST")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://example.com")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("headers"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("Content-Type")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expose_headers"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("ETag")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(3600)),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketCors_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_cors." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketCorsConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(3600)),
				},
			},
			{
				Config: testAccR2BucketCorsUpdateConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(7200)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("GET"), knownvalue.StringExact("POST"), knownvalue.StringExact("PUT")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"), knownvalue.StringExact("rule2")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("DELETE"), knownvalue.StringExact("HEAD")})),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketCors_JurisdictionEU(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_cors." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketCorsJurisdictionEUConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("eu-rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://eu.example.com")})),
				},
			},
		},
	})
}

func testAccR2BucketCorsConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2bucketcorsbasic.tf", rnd, accountID)
}

func testAccR2BucketCorsUpdateConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2bucketcorsupdate.tf", rnd, accountID)
}

func testAccR2BucketCorsJurisdictionEUConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2bucketcorsjurisdiction_eu.tf", rnd, accountID)
}

package r2_bucket_cors_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_r2_bucket_cors", &resource.Sweeper{
		Name: "cloudflare_r2_bucket_cors",
		F:    testSweepCloudflareR2BucketCors,
	})
}

func testSweepCloudflareR2BucketCors(r string) error {
	ctx := context.Background()
	// R2 Bucket CORS is a configuration setting for R2 buckets.
	// It's a singleton configuration per bucket, cleaned up when buckets are swept.
	// No sweeping required.
	tflog.Info(ctx, "R2 Bucket CORS doesn't require sweeping (bucket setting)")
	return nil
}

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
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("rule1")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("GET"), knownvalue.StringExact("POST")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://example.com")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("headers"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("Content-Type")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expose_headers"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("ETag")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(3600)),
					},
				},
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
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("rule1")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("GET"), knownvalue.StringExact("POST"), knownvalue.StringExact("PUT")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://example.com"), knownvalue.StringExact("https://test.com")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("headers"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("Content-Type"), knownvalue.StringExact("Authorization")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expose_headers"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("ETag"), knownvalue.StringExact("Content-Length")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(7200)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"), knownvalue.StringExact("rule2")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("DELETE"), knownvalue.StringExact("HEAD")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://admin.com")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(1800)),
					},
				},
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
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("eu-rule")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://eu.example.com")})),
					},
				},
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

func TestAccCloudflareR2BucketCors_JurisdictionFedramp(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_cors." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketCorsJurisdictionFedrampConfig(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("fedramp")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("fedramp-rule")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://fedramp.example.com")})),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("fedramp")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("fedramp-rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://fedramp.example.com")})),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketCors_AllHttpMethods(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_cors." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketCorsAllHttpMethodsConfig(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("all-methods-rule")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("GET"), knownvalue.StringExact("PUT"), knownvalue.StringExact("POST"), knownvalue.StringExact("DELETE"), knownvalue.StringExact("HEAD")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://example.com")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("headers"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("Content-Type"), knownvalue.StringExact("Authorization")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expose_headers"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("ETag"), knownvalue.StringExact("Content-Length")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(3600)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("all-methods-rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.StringExact("GET"),
						knownvalue.StringExact("PUT"),
						knownvalue.StringExact("POST"),
						knownvalue.StringExact("DELETE"),
						knownvalue.StringExact("HEAD"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://example.com")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("headers"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.StringExact("Content-Type"),
						knownvalue.StringExact("Authorization"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expose_headers"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.StringExact("ETag"),
						knownvalue.StringExact("Content-Length"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(3600)),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketCors_MinimalRules(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_cors." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketCorsMinimalRulesConfig(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("GET")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("*")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("headers"), knownvalue.Null()),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.Null()),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expose_headers"), knownvalue.Null()),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Null()),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("GET")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("*")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("headers"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("expose_headers"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketCors_CompleteLifecycle(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_cors." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{ // rule setup
				Config: testAccR2BucketCorsLifecycleMinimalConfig(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("GET")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://minimal.com")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.Null()),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("GET")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://minimal.com")})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.Null()),
				},
			},
			{ // rule add
				Config: testAccR2BucketCorsLifecycleFullConfig(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("full-rule")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("GET"), knownvalue.StringExact("POST"), knownvalue.StringExact("PUT")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("https://full.com"), knownvalue.StringExact("https://example.com")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(7200)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"), knownvalue.StringExact("admin-rule")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("full-rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.StringExact("GET"),
						knownvalue.StringExact("POST"),
						knownvalue.StringExact("PUT"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.SetExact([]knownvalue.Check{
						knownvalue.StringExact("https://full.com"),
						knownvalue.StringExact("https://example.com"),
					})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_age_seconds"), knownvalue.Float64Exact(7200)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"), knownvalue.StringExact("admin-rule")),
				},
			},
			{ // rule removal and edit
				Config: testAccR2BucketCorsLifecycleMinimalConfig(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("GET")})),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("origins"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("https://minimal.com")})),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("allowed").AtMapKey("methods"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("GET")})),
				},
			},
		},
	})
}

func testAccR2BucketCorsConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2cors_basic.tf", rnd, accountID)
}

func testAccR2BucketCorsUpdateConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2cors_update.tf", rnd, accountID)
}

func testAccR2BucketCorsJurisdictionEUConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2cors_jurisdiction_eu.tf", rnd, accountID)
}

func testAccR2BucketCorsJurisdictionFedrampConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2cors_jurisdiction_fedramp.tf", rnd, accountID)
}

func testAccR2BucketCorsAllHttpMethodsConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2cors_all_http_methods.tf", rnd, accountID)
}

func testAccR2BucketCorsMinimalRulesConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2cors_minimal_rules.tf", rnd, accountID)
}

func testAccR2BucketCorsLifecycleMinimalConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2cors_lifecycle_minimal.tf", rnd, accountID)
}

func testAccR2BucketCorsLifecycleFullConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2cors_lifecycle_full.tf", rnd, accountID)
}

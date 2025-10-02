package r2_bucket_lifecycle_test

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

func TestAccCloudflareR2BucketLifecycle_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lifecycle." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLifecycleConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("delete-old-objects")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("conditions").AtMapKey("prefix"), knownvalue.StringExact("logs/")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("delete_objects_transition").AtMapKey("condition").AtMapKey("type"), knownvalue.StringExact("Age")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("delete_objects_transition").AtMapKey("condition").AtMapKey("max_age"), knownvalue.Int64Exact(2592000)),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketLifecycle_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lifecycle." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLifecycleConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("delete_objects_transition").AtMapKey("condition").AtMapKey("max_age"), knownvalue.Int64Exact(2592000)),
				},
			},
			{
				Config: testAccR2BucketLifecycleUpdateConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("rules"),
						knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":      knownvalue.StringExact("delete-old-objects"),
								"enabled": knownvalue.Bool(true),
								"conditions": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"prefix": knownvalue.StringExact("logs/"),
								}),
								"delete_objects_transition": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"condition": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"type":    knownvalue.StringExact("Age"),
										"max_age": knownvalue.Int64Exact(5184000),
									}),
								}),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":      knownvalue.StringExact("archive-objects"),
								"enabled": knownvalue.Bool(true),
								"conditions": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"prefix": knownvalue.StringExact("archive/"),
								}),
								"storage_class_transitions": knownvalue.ListPartial(map[int]knownvalue.Check{
									0: knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"condition": knownvalue.ObjectPartial(map[string]knownvalue.Check{
											"type":    knownvalue.StringExact("Age"),
											"max_age": knownvalue.Int64Exact(604800),
										}),
										"storage_class": knownvalue.StringExact("InfrequentAccess"),
									}),
								}),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"id":      knownvalue.StringExact("cleanup-multipart"),
								"enabled": knownvalue.Bool(true),
								"conditions": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"prefix": knownvalue.StringExact(""),
								}),
								"abort_multipart_uploads_transition": knownvalue.ObjectPartial(map[string]knownvalue.Check{
									"condition": knownvalue.ObjectPartial(map[string]knownvalue.Check{
										"type":    knownvalue.StringExact("Age"),
										"max_age": knownvalue.Int64Exact(86400),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketLifecycle_JurisdictionEU(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lifecycle." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLifecycleJurisdictionEUConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("eu-compliance-delete")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("conditions").AtMapKey("prefix"), knownvalue.StringExact("gdpr/")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("delete_objects_transition").AtMapKey("condition").AtMapKey("max_age"), knownvalue.Int64Exact(7776000)),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketLifecycle_DateCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lifecycle." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLifecycleDateConditionConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("delete-by-date")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("delete_objects_transition").AtMapKey("condition").AtMapKey("type"), knownvalue.StringExact("Date")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("delete_objects_transition").AtMapKey("condition").AtMapKey("date"), knownvalue.StringExact("2024-12-31T23:59:59Z")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("delete_objects_transition").AtMapKey("condition").AtMapKey("max_age"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketLifecycle_MinimalRules(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lifecycle." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLifecycleMinimalRulesConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("minimal-rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("conditions").AtMapKey("prefix"), knownvalue.StringExact("")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("delete_objects_transition"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("storage_class_transitions"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("abort_multipart_uploads_transition"), knownvalue.Null()),
				},
			},
		},
	})
}

func testAccR2BucketLifecycleConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lifecycle_basic.tf", rnd, accountID)
}

func testAccR2BucketLifecycleUpdateConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lifecycle_multiple_rules.tf", rnd, accountID)
}

func testAccR2BucketLifecycleJurisdictionEUConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lifecycle_jurisdiction_eu.tf", rnd, accountID)
}

func testAccR2BucketLifecycleDateConditionConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lifecycle_date_condition.tf", rnd, accountID)
}

func testAccR2BucketLifecycleMinimalRulesConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lifecycle_minimal_rules.tf", rnd, accountID)
}

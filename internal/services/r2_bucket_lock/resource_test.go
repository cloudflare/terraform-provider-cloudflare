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

func TestAccCloudflareR2BucketLock_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lock." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLockConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("retention-rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("prefix"), knownvalue.StringExact("documents/")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("type"), knownvalue.StringExact("Age")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("max_age_seconds"), knownvalue.Int64Exact(86400)),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketLock_MultipleRules(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lock." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLockConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("max_age_seconds"), knownvalue.Int64Exact(86400)),
				},
			},
			{
				Config: testAccR2BucketLockMultipleRulesConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(3)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("archive-lock")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("max_age_seconds"), knownvalue.Int64Exact(604800)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("id"), knownvalue.StringExact("indefinite-lock")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("condition").AtMapKey("type"), knownvalue.StringExact("Indefinite")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("id"), knownvalue.StringExact("retention-rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(2).AtMapKey("condition").AtMapKey("max_age_seconds"), knownvalue.Int64Exact(172800)),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketLock_JurisdictionEU(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lock." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLockJurisdictionEUConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("eu-compliance-lock")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("prefix"), knownvalue.StringExact("gdpr/")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("max_age_seconds"), knownvalue.Int64Exact(2592000)),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketLock_DateCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lock." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLockDateConditionConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("date-lock")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("type"), knownvalue.StringExact("Date")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("date"), knownvalue.StringExact("2025-12-31T23:59:59Z")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("max_age_seconds"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketLock_IndefiniteCondition(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lock." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLockIndefiniteConditionConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("indefinite-lock")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("type"), knownvalue.StringExact("Indefinite")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("max_age_seconds"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("date"), knownvalue.Null()),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketLock_MinimalRules(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_lock." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketLockMinimalRulesConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("minimal-lock")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("prefix"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("type"), knownvalue.StringExact("Age")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("condition").AtMapKey("max_age_seconds"), knownvalue.Int64Exact(3600)),
				},
			},
		},
	})
}

func testAccR2BucketLockConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lock_basic.tf", rnd, accountID)
}

func testAccR2BucketLockMultipleRulesConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lock_multiple_rules.tf", rnd, accountID)
}

func testAccR2BucketLockJurisdictionEUConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lock_jurisdiction_eu.tf", rnd, accountID)
}

func testAccR2BucketLockDateConditionConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lock_date_condition.tf", rnd, accountID)
}

func testAccR2BucketLockIndefiniteConditionConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lock_indefinite_condition.tf", rnd, accountID)
}

func testAccR2BucketLockMinimalRulesConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2_lock_minimal_rules.tf", rnd, accountID)
}

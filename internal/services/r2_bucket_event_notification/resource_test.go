package r2_bucket_event_notification_test

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

func TestAccCloudflareR2BucketEventNotification_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_event_notification." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketEventNotificationConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_name"), knownvalue.StringExact(fmt.Sprintf("%s-queue", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("actions"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Basic event notification rule")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("prefix"), knownvalue.StringExact("documents/")),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketEventNotification_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_event_notification." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketEventNotificationConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("actions"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("prefix"), knownvalue.StringExact("documents/")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("suffix"), knownvalue.Null()),
				},
			},
			{
				Config: testAccR2BucketEventNotificationMultipleRulesConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("actions"), knownvalue.ListSizeExact(4)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("prefix"), knownvalue.StringExact("uploads/")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("suffix"), knownvalue.StringExact(".jpg")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("actions"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(1).AtMapKey("description"), knownvalue.StringExact("Lifecycle deletion notifications")),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketEventNotification_JurisdictionEU(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_event_notification." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketEventNotificationJurisdictionEUConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("queue_name"), knownvalue.StringExact(fmt.Sprintf("%s-queue-eu", rnd))),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("prefix"), knownvalue.StringExact("gdpr/")),
				},
			},
		},
	})
}

func TestAccCloudflareR2BucketEventNotification_AllActions(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket_event_notification." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketEventNotificationAllActionsConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("actions"), knownvalue.ListSizeExact(5)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("All supported R2 object actions")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("prefix"), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("suffix"), knownvalue.Null()),
				},
			},
		},
	})
}

func testAccR2BucketEventNotificationConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2notificationbasic.tf", rnd, accountID)
}

func testAccR2BucketEventNotificationMultipleRulesConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2notificationmultiple_rules.tf", rnd, accountID)
}

func testAccR2BucketEventNotificationJurisdictionEUConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2notificationjurisdiction_eu.tf", rnd, accountID)
}

func testAccR2BucketEventNotificationAllActionsConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("r2notificationall_actions.tf", rnd, accountID)
}

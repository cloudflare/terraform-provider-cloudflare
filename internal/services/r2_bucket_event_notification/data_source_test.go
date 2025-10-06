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

func TestAccCloudflareR2BucketEventNotificationDataSource_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := "data.cloudflare_r2_bucket_event_notification." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2BucketEventNotificationDataSourceConfig(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("queue_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("queue_name"), knownvalue.StringExact(fmt.Sprintf("%s-queue", rnd))),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(1)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("actions"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("description"), knownvalue.StringExact("Data source test event notifications")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("prefix"), knownvalue.StringExact("test/")),
				},
			},
		},
	})
}

func testAccR2BucketEventNotificationDataSourceConfig(rnd, accountID string) string {
	return acctest.LoadTestCase("datasource_basic.tf", rnd, accountID)
}

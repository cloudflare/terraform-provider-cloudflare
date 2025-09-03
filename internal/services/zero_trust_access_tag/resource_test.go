package zero_trust_access_tag_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_tag", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_tag",
		F:    testSweepCloudflareZeroTrustAccessTag,
	})
}

func testSweepCloudflareZeroTrustAccessTag(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Client
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return fmt.Errorf("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	// List all Access Tags
	tags, err := client.ListAccessTags(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListAccessTagsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch zero trust access tags: %s", err))
		return err
	}

	if len(tags) == 0 {
		log.Print("[DEBUG] No Cloudflare zero trust access tags to sweep")
		return nil
	}

	for _, tag := range tags {
		// Delete the tag
		err := client.DeleteAccessTag(ctx, cloudflare.AccountIdentifier(accountID), tag.Name)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete zero trust access tag %s: %s", tag.Name, err))
		} else {
			log.Printf("[INFO] Deleted zero trust access tag: %s", tag.Name)
		}
	}

	return nil
}

func TestAccCloudflareAccessTag_Basic(t *testing.T) {
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_tag.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessTag(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessTag(rnd, zoneID string) string {
	return acctest.LoadTestCase("accesstag.tf", rnd, zoneID)
}

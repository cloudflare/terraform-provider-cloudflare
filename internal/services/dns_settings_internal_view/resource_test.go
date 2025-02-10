package dns_settings_internal_view_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_dns_settings_internal_view", &resource.Sweeper{
		Name: "cloudflare_dns_settings_internal_view",
		F:    testSweepCloudflareDNSSettingsInternalView,
	})
}

func testSweepCloudflareDNSSettingsInternalView(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	// Clean up the account level views
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return errors.New("CLOUDFLARE_ACCOUNT_ID must be set")
	}

	views, err := client.DNS.Settings.Views.List(context.Background(), dns.SettingViewListParams{AccountID: cloudflare.F(accountID)})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare DNS internal views: %s", err))
	}

	if len(views.Result) == 0 {
		log.Print("[DEBUG] No Cloudflare views to sweep")
		return nil
	}

	for _, view := range views.Result {
		tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare View ID: %s", view.ID))
		//nolint:errcheck
		client.DNS.Settings.Views.Delete(context.TODO(), view.ID, dns.SettingViewDeleteParams{AccountID: cloudflare.F(accountID)})
	}

	return nil
}

func TestAccCloudflareDNSInternalView_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_dns_settings_internal_view." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDNSInternalViewConfig(rnd, accountID, rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "zones.0", zoneID),
				),
			},
		},
	})
}

func TestAccCloudflareDNSInternalView_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_dns_settings_internal_view." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDNSInternalViewConfig(rnd, accountID, rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "zones.0", zoneID),
				),
			},
			{
				Config: testDNSInternalViewConfig(rnd, accountID, rnd+"-update", zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd+"-update"),
					resource.TestCheckResourceAttr(name, "zones.0", zoneID),
				),
			},
		},
	})
}

func testDNSInternalViewConfig(resourceID, accountID, name, zone string) string {
	return acctest.LoadTestCase("view.tf", resourceID, accountID, name, zone)
}

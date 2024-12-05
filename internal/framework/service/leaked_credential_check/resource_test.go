package leaked_credential_check_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_leaked_credential_check", &resource.Sweeper{
		Name: "cloudflare_leaked_credential_check",
		F:    testSweepCloudflareLCC,
	})
}

func testSweepCloudflareLCC(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}

	status, err := client.LeakedCredentialCheckGetStatus(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.LeakedCredentialCheckGetStatusParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to GET Leaked Credential status: %s", err))
	}

	if *status.Enabled == false {
		log.Print("[DEBUG] LCC already disabled")
		return nil
	}
	_, err = client.LeakedCredentialCheckSetStatus(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.LeakCredentialCheckSetStatusParams{Enabled: cfv1.BoolPtr(false)})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to disable Leaked Credential Check: %s", err))
		return err
	}

	return nil
}

func TestAccCloudflareLeakedCredentialCheck_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_leaked_credential_check.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLCCSimple(rnd, zoneID, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
			{
				Config: testAccLCCSimple(rnd, zoneID, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
			},
		},
	})
}

func testAccLCCSimple(ID, zoneID, enabled string) string {
	return fmt.Sprintf(`
  resource "cloudflare_leaked_credential_check" "%[1]s" {
    zone_id = "%[2]s"
    enabled = "%[3]s"
  }`, ID, zoneID, enabled)
}

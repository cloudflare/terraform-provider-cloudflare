package leaked_credential_check_rule_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_leaked_credential_check_rule", &resource.Sweeper{
		Name: "cloudflare_leaked_credential_check_rule",
		F:    testSweepCloudflareLCCRules,
	})
}

func testSweepCloudflareLCCRules(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		return errors.New("CLOUDFLARE_ZONE_ID must be set")
	}
	// fetch existing rules from API
	rules, err := client.LeakedCredentialCheckListDetections(ctx, cfv1.ZoneIdentifier(zoneID), cfv1.LeakedCredentialCheckListDetectionsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error fetching Leaked Credential Check user-defined detection patterns: %s", err))
		return err
	}
	for _, rule := range rules {
		deleteParam := cfv1.LeakedCredentialCheckDeleteDetectionParams{DetectionID: rule.ID}
		_, err := client.LeakedCredentialCheckDeleteDetection(ctx, cfv1.ZoneIdentifier(zoneID), deleteParam)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error deleting a user-defined detection patter for Leaked Credential Check: %s", err))
		}
	}

	return nil
}

func TestAccCloudflareLeakedCredentialCheckRule_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_leaked_credential_check_rule.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccConfigAddHeader(rnd, zoneID, testAccLCCTwoSimpleRules(rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name+"_first", "zone_id", zoneID),
					resource.TestCheckResourceAttr(name+"_first", "username", "lookup_json_string(http.request.body.raw, \"user\")"),
					resource.TestCheckResourceAttr(name+"_first", "password", "lookup_json_string(http.request.body.raw, \"pass\")"),

					resource.TestCheckResourceAttr(name+"_second", "zone_id", zoneID),
					resource.TestCheckResourceAttr(name+"_second", "username", "lookup_json_string(http.request.body.raw, \"id\")"),
					resource.TestCheckResourceAttr(name+"_second", "password", "lookup_json_string(http.request.body.raw, \"secret\")"),
				),
			},
			{
				Config: testAccConfigAddHeader(rnd, zoneID, testAccLCCUpdateOneRule(rnd)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name+"_first", "zone_id", zoneID),
					resource.TestCheckResourceAttr(name+"_first", "username", "lookup_json_string(http.request.body.raw, \"username\")"),
					resource.TestCheckResourceAttr(name+"_first", "password", "lookup_json_string(http.request.body.raw, \"password\")"),

					resource.TestCheckResourceAttr(name+"_second", "zone_id", zoneID),
					resource.TestCheckResourceAttr(name+"_second", "username", "lookup_json_string(http.request.body.raw, \"id\")"),
					resource.TestCheckResourceAttr(name+"_second", "password", "lookup_json_string(http.request.body.raw, \"secret\")"),
				),
			},
		},
	})
}

func testAccConfigAddHeader(name, zoneID, config string) string {
	header := fmt.Sprintf(`
	resource "cloudflare_leaked_credential_check" "%[1]s" {
		zone_id = "%[2]s"
		enabled = true
	}`, name, zoneID)
	return header + "\n" + config
}

func testAccLCCTwoSimpleRules(name string) string {
	return fmt.Sprintf(`
  resource "cloudflare_leaked_credential_check_rule" "%[1]s_first" {
    zone_id = cloudflare_leaked_credential_check.%[1]s.zone_id
		username = "lookup_json_string(http.request.body.raw, \"user\")"
		password = "lookup_json_string(http.request.body.raw, \"pass\")"
  }

  resource "cloudflare_leaked_credential_check_rule" "%[1]s_second" {
    zone_id = cloudflare_leaked_credential_check.%[1]s.zone_id
    username = "lookup_json_string(http.request.body.raw, \"id\")"
		password = "lookup_json_string(http.request.body.raw, \"secret\")"
  }`, name)
}

func testAccLCCUpdateOneRule(name string) string {
	return fmt.Sprintf(`
  resource "cloudflare_leaked_credential_check_rule" "%[1]s_first" {
    zone_id = cloudflare_leaked_credential_check.%[1]s.zone_id
		username = "lookup_json_string(http.request.body.raw, \"username\")"
		password = "lookup_json_string(http.request.body.raw, \"password\")"
  }

  resource "cloudflare_leaked_credential_check_rule" "%[1]s_second" {
    zone_id = cloudflare_leaked_credential_check.%[1]s.zone_id
    username = "lookup_json_string(http.request.body.raw, \"id\")"
		password = "lookup_json_string(http.request.body.raw, \"secret\")"
  }`, name)
}

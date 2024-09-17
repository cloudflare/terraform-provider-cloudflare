package custom_ssl_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareCustomSSL_Basic(t *testing.T) {
	var customSSL cloudflare.ZoneCustomSSL
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_ssl." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomSSLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomSSLCertBasic(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareCustomSSLExists(resourceName, &customSSL),
					resource.TestCheckResourceAttr(resourceName, consts.ZoneIDSchemaKey, zoneID),
					resource.TestMatchResourceAttr(resourceName, "priority", regexp.MustCompile("^[0-9]\\d*$")),
					resource.TestCheckResourceAttr(resourceName, "status", "pending"),
					resource.TestMatchResourceAttr(resourceName, consts.ZoneIDSchemaKey, regexp.MustCompile("^[a-z0-9]{32}$")),
					resource.TestCheckResourceAttr(resourceName, "type", "legacy_custom"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomSSLCertBasic(zoneID string, rName string) string {
	return acctest.LoadTestCase("customsslcertbasic.tf", zoneID, rName)
}

func testAccCheckCloudflareCustomSSLDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_ssl" {
			continue
		}

		err := client.DeleteSSL(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cert still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareCustomSSLExists(n string, customSSL *cloudflare.ZoneCustomSSL) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cert ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundCustomSSL, err := client.SSLDetails(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundCustomSSL.ID != rs.Primary.ID {
			return fmt.Errorf("cert not found")
		}

		*customSSL = foundCustomSSL

		return nil
	}
}

func TestAccCloudflareCustomSSL_WithEmptyGeoRestrictions(t *testing.T) {
	var customSSL cloudflare.ZoneCustomSSL
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_ssl." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomSSLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareCustomSSLWithEmptyGeoRestrictions(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareCustomSSLExists(resourceName, &customSSL),
					resource.TestCheckNoResourceAttr(resourceName, "geo_restrictions.%"),
				),
			},
		},
	})
}

func testAccCheckCloudflareCustomSSLWithEmptyGeoRestrictions(zoneID string, rName string) string {
	return acctest.LoadTestCase("customsslwithemptygeorestrictions.tf", zoneID, rName)
}

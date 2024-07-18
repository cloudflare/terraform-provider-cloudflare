package certificate_pack_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_certificate_pack", &resource.Sweeper{
		Name: "cloudflare_certificate_pack",
		F:    testSweepCloudflareCertificatePack,
	})
}

func testSweepCloudflareCertificatePack(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	certificates, certErr := client.ListCertificatePacks(context.Background(), zoneID)
	if certErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch certificate packs: %s", clientErr))
	}

	if len(certificates) == 0 {
		log.Print("[DEBUG] No Cloudflare certificate packs to sweep")
		return nil
	}

	for _, certificate := range certificates {
		if err := client.DeleteCertificatePack(context.Background(), zoneID, certificate.ID); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete certificate pack %s", certificate.ID))
		}
	}

	return nil
}

func TestAccCertificatePack_AdvancedLetsEncrypt(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackAdvancedLetsEncryptConfig(zoneID, domain, "advanced", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "type", "advanced"),
					resource.TestCheckResourceAttr(name, "hosts.#", "2"),
					resource.TestCheckResourceAttr(name, "validation_method", "txt"),
					resource.TestCheckResourceAttr(name, "validity_days", "90"),
					resource.TestCheckResourceAttr(name, "certificate_authority", "lets_encrypt"),
					resource.TestCheckResourceAttr(name, "cloudflare_branding", "false"),
					resource.TestCheckResourceAttr(name, "wait_for_active_status", "false"),
				),
			},
		},
	})
}

func testAccCertificatePackAdvancedLetsEncryptConfig(zoneID, domain, certType, rnd string) string {
	return acctest.LoadTestCase("acccertificatepackadvancedletsencryptconfig.tf", zoneID, domain, rnd, certType)
}

func TestAccCertificatePack_WaitForActive(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_certificate_pack." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCertificatePackAdvancedWaitForActiveConfig(zoneID, domain, "advanced", rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "type", "advanced"),
					resource.TestCheckResourceAttr(name, "hosts.#", "2"),
					resource.TestCheckResourceAttr(name, "validation_method", "txt"),
					resource.TestCheckResourceAttr(name, "validity_days", "90"),
					resource.TestCheckResourceAttr(name, "certificate_authority", "lets_encrypt"),
					resource.TestCheckResourceAttr(name, "cloudflare_branding", "false"),
					resource.TestCheckResourceAttr(name, "wait_for_active_status", "true"),
				),
			},
		},
	})
}

func testAccCertificatePackAdvancedWaitForActiveConfig(zoneID, domain, certType, rnd string) string {
	return acctest.LoadTestCase("acccertificatepackadvancedwaitforactiveconfig.tf", zoneID, domain, rnd, certType)
}

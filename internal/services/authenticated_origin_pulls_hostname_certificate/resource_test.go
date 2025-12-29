package authenticated_origin_pulls_hostname_certificate_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/origin_tls_client_auth"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_authenticated_origin_pulls_hostname_certificate", &resource.Sweeper{
		Name: "cloudflare_authenticated_origin_pulls_hostname_certificate",
		F:    testSweepCloudflareAuthenticatedOriginPullsHostnameCertificates,
	})
}

func testSweepCloudflareAuthenticatedOriginPullsHostnameCertificates(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	if zoneID == "" {
		tflog.Info(ctx, "Skipping authenticated origin pulls hostname certificates sweep: CLOUDFLARE_ZONE_ID not set")
		return nil
	}

	// List all hostname certificates
	page, err := client.OriginTLSClientAuth.Hostnames.Certificates.List(ctx, origin_tls_client_auth.HostnameCertificateListParams{
		ZoneID: cloudflare.F(zoneID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch hostname certificates: %s", err))
		return err
	}

	if len(page.Result) == 0 {
		tflog.Info(ctx, "No hostname certificates to sweep")
		return nil
	}

	for _, cert := range page.Result {
		// Only sweep test certificates
		if !utils.ShouldSweepResource(cert.ID) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting hostname certificate: %s", cert.ID))
		_, err := client.OriginTLSClientAuth.Hostnames.Certificates.Delete(ctx, cert.ID, origin_tls_client_auth.HostnameCertificateDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete hostname certificate %s: %s", cert.ID, err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted hostname certificate: %s", cert.ID))
	}

	return nil
}

func TestAccCloudflareAuthenticatedOriginPullsHostnameCertificate_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_authenticated_origin_pulls_hostname_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAuthenticatedOriginPullsHostnameCertificateConfig(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "certificate_id"),
					resource.TestCheckResourceAttr(name, "zone_id", zoneID),
					resource.TestCheckResourceAttrSet(name, "certificate"),
					resource.TestCheckResourceAttrSet(name, "expires_on"),
					resource.TestCheckResourceAttrSet(name, "issuer"),
					resource.TestCheckResourceAttrSet(name, "status"),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key"},
			},
		},
	})
}

func testAccCheckCloudflareAuthenticatedOriginPullsHostnameCertificateConfig(zoneID, name string) string {
	return acctest.LoadTestCase("authenticatedoriginpullshostnamecertificateconfig.tf", name, zoneID)
}

package mtls_certificate_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_mtls_certificate", &resource.Sweeper{
		Name: "cloudflare_mtls_certificate",
		F:    testSweepCloudflareMTLSCertificates,
	})
}

func testSweepCloudflareMTLSCertificates(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accountIDrc := cloudflare.AccountIdentifier(accountID)
	mtlsCertificates, _, certsErr := client.ListMTLSCertificates(context.Background(), accountIDrc, cloudflare.ListMTLSCertificatesParams{})

	if certsErr != nil {
		tflog.Error(ctx, fmt.Sprintf("failed to fetch mtls certificates: %s", certsErr))
	}

	if len(mtlsCertificates) == 0 {
		tflog.Debug(ctx, "no mtls certificates to sweep")
		return nil
	}

	for _, certificate := range mtlsCertificates {
		_, err := client.DeleteMTLSCertificate(context.Background(), accountIDrc, certificate.ID)

		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("failed to delete mtls pull certificate (%s) in account ID: %s", certificate.ID, accountID))
		}
	}

	return nil
}

func TestAccCloudflareMTLSCertificate(t *testing.T) {
	var mtlsCert cloudflare.MTLSCertificate
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_mtls_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck_Credentials(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareMTLSCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareMTLSCertificateConfig(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareMTLSCertificateExists(name, &mtlsCert),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
				),
			},
		},
	})
}

func testAccCheckCloudflareMTLSCertificateExists(name string, mtlsCert *cloudflare.MTLSCertificate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No cert ID is set")
		}
		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		accountIDrc := cloudflare.AccountIdentifier(accountID)
		foundMTLSCert, err := client.GetMTLSCertificate(context.Background(), accountIDrc, rs.Primary.ID)
		if err != nil {
			return err
		}
		if foundMTLSCert.ID != rs.Primary.ID {
			return fmt.Errorf("cert not found")
		}
		*mtlsCert = foundMTLSCert
		return nil
	}
}

func testAccCheckCloudflareMTLSCertificateConfig(accountID, rnd string) string {
	return acctest.LoadTestCase("mtlscertificateconfig.tf", accountID, rnd)
}

func testAccCheckCloudflareMTLSCertificateDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}
	for _, rs := range s.RootModule().Resources {
		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		accountIDrc := cloudflare.AccountIdentifier(accountID)
		_, err := client.DeleteMTLSCertificate(context.Background(), accountIDrc, rs.Primary.ID)
		if err == nil {
			// certificate should have already been destroyed before this check function is called
			return fmt.Errorf("error deleting mTLS certificate in account %q: %w", accountID, err)
		}
	}
	return nil
}

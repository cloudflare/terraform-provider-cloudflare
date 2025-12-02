package zero_trust_access_short_lived_certificate_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	cloudflare_v6 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_short_lived_certificate", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_short_lived_certificate",
		F:    testSweepCloudflareZeroTrustAccessShortLivedCertificate,
	})
}

func testSweepCloudflareZeroTrustAccessShortLivedCertificate(r string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	// List all applications to get app names for filtering
	appsResp, err := client.ZeroTrust.Access.Applications.List(
		ctx,
		zero_trust.AccessApplicationListParams{
			AccountID: cloudflare_v6.F(accountID),
		},
	)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Zero Trust Access Applications: %s", err))
		return err
	}

	// Build map of app IDs that should be swept
	appIDsToSweep := make(map[string]bool)
	for _, app := range appsResp.Result {
		if utils.ShouldSweepResource(app.Name) {
			appIDsToSweep[app.ID] = true
		}
	}

	if len(appIDsToSweep) == 0 {
		tflog.Info(ctx, "No test applications found to sweep")
		return nil
	}

	// List all CAs (v6 API lists all CAs, not per-application)
	casResp, err := client.ZeroTrust.Access.Applications.CAs.List(
		ctx,
		zero_trust.AccessApplicationCAListParams{
			AccountID: cloudflare_v6.F(accountID),
		},
	)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Failed to fetch CAs: %s", err))
		return err
	}

	// Iterate through all CAs and delete those belonging to test applications
	for _, ca := range casResp.Result {
		// The AUD field contains the application ID
		if !appIDsToSweep[ca.AUD] {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Zero Trust Access Short Lived Certificate: %s for app %s", ca.ID, ca.AUD))
		_, err := client.ZeroTrust.Access.Applications.CAs.Delete(
			ctx,
			ca.AUD,
			zero_trust.AccessApplicationCADeleteParams{
				AccountID: cloudflare_v6.F(accountID),
			},
		)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Zero Trust Access Short Lived Certificate %s: %s", ca.ID, err))
		}
	}

	return nil
}

func TestAccCloudflareAccessCACertificate_AccountLevel(t *testing.T) {
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_short_lived_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessCACertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessCACertificateBasic(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttrSet(name, "app_id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
					resource.TestCheckResourceAttrSet(name, "public_key"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessCACertificateBasic(rnd, domain, cloudflare.AccountIdentifier(accountID)),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessCACertificate_ZoneLevel(t *testing.T) {
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_short_lived_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessCACertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessCACertificateBasic(rnd, domain, cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttrSet(name, "app_id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
					resource.TestCheckResourceAttrSet(name, "public_key"),
				),
			},
			{
				// Ensures no diff on last plan
				Config:   testAccCloudflareAccessCACertificateBasic(rnd, domain, cloudflare.ZoneIdentifier(zoneID)),
				PlanOnly: true,
			},
		},
	})
}

func testAccCloudflareAccessCACertificateBasic(resourceName, domain string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accesscacertificatebasic.tf", resourceName, domain, identifier.Type, identifier.Identifier)
}

func testAccCheckCloudflareAccessCACertificateDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_access_short_lived_certificate" {
			continue
		}

		_, err := client.GetAccessCACertificate(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Access CA certificate still exists")
		}

		_, err = client.GetAccessCACertificate(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Access CA certificate still exists")
		}
	}

	return nil
}

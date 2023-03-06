package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
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
	client, clientErr := sharedClient()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
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
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_mtls_certificate.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithoutZoneID(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareMTLSCertificateDestroy,
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
		client := testAccProvider.Meta().(*cloudflare.API)
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
	return fmt.Sprintf(`
  resource "cloudflare_mtls_certificate" "%[2]s" {
	  account_id   = "%[1]s"
	  name         = ""
	  certificates = "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoCgAwIBAgIUKTOAZNjcXVZRj4oQt0SHsl1c1vMwDQYJKoZIhvcNAQELBQAwUTELMAkGA1UEBhMCVVMxFjAUBgNVBAgMDVNhbiBGcmFuY2lzY28xEzARBgNVBAcMCkNhbGlmb3JuaWExFTATBgNVBAoMDEV4YW1wbGUgSW5jLjAgFw0yMjExMjIxNjU5NDdaGA8yMTIyMTAyOTE2NTk0N1owUTELMAkGA1UEBhMCVVMxFjAUBgNVBAgMDVNhbiBGcmFuY2lzY28xEzARBgNVBAcMCkNhbGlmb3JuaWExFTATBgNVBAoMDEV4YW1wbGUgSW5jLjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMRcORwgJFTdcG/2GKI+cFYiOBNDKjCZUXEOvXWY42BkH9wxiMT869CO+enA1w5pIrXow6kCM1sQspHHaVmJUlotEMJxyoLFfA/8Kt1EKFyobOjuZs2SwyVyJ2sStvQuUQEosULZCNGZEqoH5g6zhMPxaxm7ZLrrsDZ9maNGVqo7EWLWHrZ57Q/5MtTrbxQL+eXjUmJ9K3kS+3uEwMdqR6Z3BluU1ivanpPc1CN2GNhdO0/hSY4YkGEnuLsqJyDd3cIiB1MxuCBJ4ZaqOd2viV1WcP3oU3dxVPm4MWyfYIldMWB14FahScxLhWdRnM9YZ/i9IFcLypXsuz7DjrJPtPUCAwEAAaNmMGQwHQYDVR0OBBYEFP5JzLUawNF+c3AXsYTEWHh7z2czMB8GA1UdIwQYMBaAFP5JzLUawNF+c3AXsYTEWHh7z2czMA4GA1UdDwEB/wQEAwIBBjASBgNVHRMBAf8ECDAGAQH/AgEBMA0GCSqGSIb3DQEBCwUAA4IBAQBc+Be7NDhpE09y7hLPZGRPl1cSKBw4RI0XIv6rlbSTFs5EebpTGjhx/whNxwEZhB9HZ7111Oa1YlT8xkI9DshB78mjAHCKBAJ76moK8tkG0aqdYpJ4ZcJTVBB7l98Rvgc7zfTii7WemTy72deBbSeiEtXavm4EF0mWjHhQ5Nxpnp00Bqn5g1x8CyTDypgmugnep+xG+iFzNmTdsz7WI9T/7kDMXqB7M/FPWBORyS98OJqNDswCLF8bIZYwUBEe+bRHFomoShMzaC3tvim7WCb16noDkSTMlfKO4pnvKhpcVdSgwcruATV7y+W+Lvmz2OT/Gui4JhqeoTewsxndhDDE\n-----END CERTIFICATE-----"
	  private_key  = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDEXDkcICRU3XBv9hiiPnBWIjgTQyowmVFxDr11mONgZB/cMYjE/OvQjvnpwNcOaSK16MOpAjNbELKRx2lZiVJaLRDCccqCxXwP/CrdRChcqGzo7mbNksMlcidrErb0LlEBKLFC2QjRmRKqB+YOs4TD8WsZu2S667A2fZmjRlaqOxFi1h62ee0P+TLU628UC/nl41JifSt5Evt7hMDHakemdwZblNYr2p6T3NQjdhjYXTtP4UmOGJBhJ7i7Kicg3d3CIgdTMbggSeGWqjndr4ldVnD96FN3cVT5uDFsn2CJXTFgdeBWoUnMS4VnUZzPWGf4vSBXC8qV7Ls+w46yT7T1AgMBAAECggEAQZnp/oqCeNPOR6l5S2L+1tfx0gWjZ78hJVteUpZ0iHSK7F6kKeOxyOird7vUXV0kmo+cJq+0hp0Ke4eam640FCpwKfYoSQ4/R3vgujGWJnaihCN5tv5sMet0XeJPuz5qE7ALoKCvwI6aXLHs20aAeZIDTQJ9QbGSGnJVzOWn+JDTidIgZpN57RpXfSAwnJPTQK/PN8i5z108hsaDOdEgGmxYZ7kYqMqzX20KXmth58LDfPixs5JGtS60iiKC/wOcGzkB2/AdTSojR76oEU77cANP/3zO25NG//whUdYlW0t0d7PgXxIeJe+xgYnamDQJx3qonVyt4H77ha0ObRAj9QKBgQDicZr+VTwFMnELP3a+FXGnjehRiuS1i7MXGKxNweCD+dFlML0FplSQS8Ro2n+d8lu8BBXGx0qm6VXu8Rhn7TAUL6q+PCgfarzxfIhacb/TZCqfieIHsMlVBfhV5HCXnk+kis0tuC/PRArcWTwDHJUJXkBhvkUsNswvQzavDPI7KwKBgQDd/WgLkj7A3X5fgIHZH/GbDSBiXwzKb+rF4ZCT2XFgG/OAW7vapfcX/w+v+5lBLyrocmOAS3PGGAhM5T3HLnUCQfnK4qgps1Lqibkc9Tmnsn60LanUjuUMsYv/zSw70tozbzhJ0pioEpWfRxRZBztO2Rr8Ntm7h6Fk701EXGNAXwKBgQCD1xsjy2J3sCerIdcz0u5qXLAPkeuZW+34m4/ucdwTWwc0gEz9lhsULFj9p4G351zLuiEnq+7mAWLcDJlmIO3mQt6JhiLiL9Y0T4pgBmxmWqKKYtAsJB0EmMY+1BNN44mBRqMxZFTJu1cLdhT/xstrOeoIPqytknYNanfTMZlzIwKBgHrLXe5oq0XMP8dcMneEcAUwsaU4pr6kQd3L9EmUkl5zl7J9C+DaxWAEuwzBw/iGutlxzRB+rD/7szu14wJ29EqXbDGKRzMp+se5/yfBjm7xEZ1hVPw7PwBShfqt57X/4Ktq7lwHnmH6RcGhc+P7WBc5iO/S94YAdIp8xOT3pf9JAoGAE0QkqJUY+5Mgr+fBO0VNV72ZoPveGpW+De59uhKAOnu1zljQCUtk59m6+DXfm0tNYKtawa5n8iN71Zh+s62xXSt3pYi1Y5CCCmv8Y4BhwIcPwXKk3zEvLgSHVTpC0bayA9aSO4bbZgVXa5w+Z0w/vvfp9DWo1IS3EnQRrz6WMYA=\n-----END PRIVATE KEY-----"
	  ca           = true
  }`, accountID, rnd)
}

func testAccCheckCloudflareMTLSCertificateDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)
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

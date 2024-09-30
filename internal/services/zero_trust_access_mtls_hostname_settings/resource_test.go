package zero_trust_access_mtls_hostname_settings_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	cfv2 "github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_mtls_hostname_settings", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_mtls_hostname_settings",
		F: func(region string) error {
			ctx := context.Background()

			client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
			if clientErr != nil {
				return fmt.Errorf("Failed to create Cloudflare client: %w", clientErr)
			}

			deletedSettings := cfv1.UpdateAccessMutualTLSHostnameSettingsParams{
				Settings: []cfv1.AccessMutualTLSHostnameSettings{},
			}

			accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
			_, err := client.UpdateAccessMutualTLSHostnameSettings(ctx, cfv1.AccountIdentifier(accountID), deletedSettings)
			if err != nil {
				return fmt.Errorf("Failed to fetch Cloudflare Access Mutual TLS hostname settings: %w", err)
			}

			zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
			_, err = client.UpdateAccessMutualTLSHostnameSettings(ctx, cfv1.ZoneIdentifier(zoneID), deletedSettings)
			if err != nil {
				return fmt.Errorf("Failed to delete Cloudflare Access Mutual TLS hostname settings: %w", err)
			}

			return nil
		},
	})
}

func TestAccCloudflareAccessMutualTLSHostnameSettings_Simple(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_mtls_hostname_settings.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessMutualTLSHostnameSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessMutualTLSHostnameSettingsConfig(rnd, cfv1.AccountIdentifier(accountID), domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "settings.0.hostname", domain),
					resource.TestCheckResourceAttr(name, "settings.0.china_network", "false"),
					resource.TestCheckResourceAttr(name, "settings.0.client_certificate_forwarding", "true"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessMutualTLSHostnameSettingsDestroy(s *terraform.State) error {
	client := acctest.SharedV2Client()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_access_mtls_hostname_settings" {
			continue
		}

		for _, rs := range s.RootModule().Resources {
			certificates, _ := client.ZeroTrust.Access.Certificates.Get(context.Background(), rs.Primary.Attributes["id"], zero_trust.AccessCertificateGetParams{
				AccountID: cfv2.F(rs.Primary.Attributes[consts.AccountIDSchemaKey]),
			})

			if certificates != nil {
				return fmt.Errorf("access_mtls_hostname_settings still exists")
			}
		}
	}

	return nil
}

func testAccessMutualTLSHostnameSettingsConfig(rnd string, identifier *cfv1.ResourceContainer, domain string) string {
	return acctest.LoadTestCase("accessmutualtlshostnamesettingsconfig.tf", rnd, identifier.Type, identifier.Identifier, domain)
}

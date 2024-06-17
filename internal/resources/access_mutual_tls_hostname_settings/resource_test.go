package access_mutual_tls_hostname_settings_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_access_mutual_tls_hostname_settings", &resource.Sweeper{
		Name: "cloudflare_access_mutual_tls_hostname_settings",
		F: func(region string) error {
			ctx := context.Background()

			client, clientErr := acctest.SharedV1Client()
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
	name := fmt.Sprintf("cloudflare_access_mutual_tls_hostname_settings.%s", rnd)
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_Account(t)
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
	client, err := acctest.SharedV1Client()
	if err != nil {
		return fmt.Errorf("Failed to create Cloudflare client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_access_mutual_tls_hostname_settings" {
			continue
		}

		if rs.Primary.Attributes[consts.ZoneIDSchemaKey] != "" {
			settings, _ := client.GetAccessMutualTLSHostnameSettings(context.Background(), cfv1.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]))
			if len(settings) != 0 {
				return fmt.Errorf("AccessMutualTLSHostnameSettings still exists")
			}
		}

		if rs.Primary.Attributes[consts.AccountIDSchemaKey] != "" {
			settings, _ := client.GetAccessMutualTLSHostnameSettings(context.Background(), cfv1.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]))
			if len(settings) != 0 {
				return fmt.Errorf("AccessMutualTLSHostnameSettings still exists")
			}
		}
	}

	return nil
}

func testAccessMutualTLSHostnameSettingsConfig(rnd string, identifier *cfv1.ResourceContainer, domain string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_mutual_tls_hostname_settings" "%[1]s" {
	%[2]s_id             = "%[3]s"
	settings {
		hostname = "%[4]s"
		client_certificate_forwarding = true
		china_network = false
	}
}
`, rnd, identifier.Type, identifier.Identifier, domain)
}

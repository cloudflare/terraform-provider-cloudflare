package zero_trust_device_posture_integration_test

import (
	"context"
	"fmt"
	"os"
	"testing"

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
	resource.AddTestSweepers("cloudflare_zero_trust_device_posture_integration", &resource.Sweeper{
		Name: "cloudflare_zero_trust_device_posture_integration",
		F:    testSweepCloudflareZeroTrustDevicePostureIntegration,
	})
}

func testSweepCloudflareZeroTrustDevicePostureIntegration(r string) error {
	ctx := context.Background()
	client, clientErr := acctest.SharedV1Client()
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	integrations, _, err := client.DevicePostureIntegrations(ctx, accountID)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Zero Trust Device Posture Integrations: %s", err))
		return err
	}

	for _, integration := range integrations {
		if !utils.ShouldSweepResource(integration.Name) {
			continue
		}

		tflog.Info(ctx, fmt.Sprintf("Deleting Zero Trust Device Posture Integration: %s", integration.IntegrationID))
		err := client.DeleteDevicePostureIntegration(ctx, accountID, integration.IntegrationID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Zero Trust Device Posture Integration %s: %s", integration.IntegrationID, err))
		}
	}

	return nil
}

func TestAccCloudflareDevicePostureIntegrationCreate(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_posture_integration.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	clientID := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_CLIENT_ID")
	clientSecret := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_CLIENT_SECRET")
	apiURL := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_API_URL")
	authURL := os.Getenv("CLOUDFLARE_WORKSPACE_ONE_AUTH_URL")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_WorkspaceOne(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureIntegration(rnd, accountID, clientID, clientSecret, apiURL, authURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "workspace_one"),
					resource.TestCheckResourceAttr(name, "interval", "24h"),
					resource.TestCheckResourceAttr(name, "config.auth_url", authURL),
					resource.TestCheckResourceAttr(name, "config.api_url", apiURL),
					resource.TestCheckResourceAttr(name, "config.client_id", clientID),
				),
			},
		},
	})
}

func testAccCloudflareDevicePostureIntegration(rnd, accountID, clientID, clientSecret, apiURL, authURL string) string {
	return acctest.LoadTestCase("devicepostureintegration.tf", rnd, accountID, clientID, clientSecret, apiURL, authURL)
}

func testAccCheckCloudflareDevicePostureIntegrationDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_device_posture_integration" {
			continue
		}

		_, err := client.DevicePostureIntegration(context.Background(), rs.Primary.Attributes[consts.AccountIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Posture Integration still exists")
		}
	}

	return nil
}

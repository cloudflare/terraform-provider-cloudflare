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
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

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

	clientID := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_ID")
	clientSecret := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CLIENT_SECRET")
	apiURL := os.Getenv("CLOUDFLARE_CROWDSTRIKE_API_URL")
	customerId := os.Getenv("CLOUDFLARE_CROWDSTRIKE_CUSTOMER_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_CrowdStrike(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareDevicePostureIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareDevicePostureIntegration(rnd, accountID, clientID, clientSecret, apiURL, customerId),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("type"), knownvalue.StringExact("crowdstrike_s2s")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("interval"), knownvalue.StringExact("24h")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("api_url"), knownvalue.StringExact(apiURL)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("client_id"), knownvalue.StringExact(clientID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("config").AtMapKey("customer_id"), knownvalue.StringExact(customerId)),
				},
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("%s/", accountID),
				ImportStateVerifyIgnore: []string{"config.client_secret", "config.customer_id"},
			},
		},
	})
}

func testAccCloudflareDevicePostureIntegration(rnd, accountID, clientID, clientSecret, apiURL, customerId string) string {
	return acctest.LoadTestCase("devicepostureintegration.tf", rnd, accountID, clientID, clientSecret, apiURL, customerId)
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

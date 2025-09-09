package zero_trust_device_custom_profile_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_device_custom_profile", &resource.Sweeper{
		Name: "cloudflare_zero_trust_device_custom_profile",
		F:    testSweepCloudflareZeroTrustDeviceCustomProfile,
	})
}

func testSweepCloudflareZeroTrustDeviceCustomProfile(region string) error {
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		log.Print("[DEBUG] CLOUDFLARE_ACCOUNT_ID not set, skipping sweep")
		return nil
	}

	ctx := context.Background()

	// List all custom device profiles
	profiles, err := client.ZeroTrust.Devices.Policies.Custom.List(
		ctx,
		zero_trust.DevicePolicyCustomListParams{
			AccountID: cloudflare.F(accountID),
		},
	)
	if err != nil {
		return fmt.Errorf("error listing device custom profiles for sweep: %w", err)
	}

	log.Printf("[DEBUG] Found %d device custom profiles to sweep", len(profiles.Result))

	for _, profile := range profiles.Result {
		if profile.PolicyID == "" {
			log.Printf("[DEBUG] Skipping device custom profile with empty policy_id: %s", profile.Name)
			continue
		}

		log.Printf("[INFO] Deleting device custom profile: %s (%s)", profile.Name, profile.PolicyID)

		_, err := client.ZeroTrust.Devices.Policies.Custom.Delete(
			ctx,
			profile.PolicyID,
			zero_trust.DevicePolicyCustomDeleteParams{
				AccountID: cloudflare.F(accountID),
			},
		)
		if err != nil {
			log.Printf("[ERROR] Failed to delete device custom profile %s: %v", profile.PolicyID, err)
			continue
		}
	}

	return nil
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

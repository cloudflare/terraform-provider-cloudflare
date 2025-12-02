package zero_trust_device_default_profile_certificates_test

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
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_device_default_profile_certificates", &resource.Sweeper{
		Name: "cloudflare_zero_trust_device_default_profile_certificates",
		F:    testSweepCloudflareZeroTrustDeviceDefaultProfileCertificates,
	})
}

func testSweepCloudflareZeroTrustDeviceDefaultProfileCertificates(r string) error {
	ctx := context.Background()
	// Device Default Profile Certificates is a zone-level certificate setting for the default device profile.
	// It's a singleton setting per zone, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Zero Trust Device Default Profile Certificates doesn't require sweeping (zone setting)")
	return nil
}

func TestAccCloudflareDeviceDefaultProfileCertificates_Create(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_device_default_profile_certificates.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCloudflareDevicePolicyCertificates(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
				),
			},
		},
	})
}

func testCloudflareDevicePolicyCertificates(rnd, zoneID string, enable bool) string {
	return fmt.Sprintf(`
resource "cloudflare_zero_trust_device_default_profile_certificates" "%[1]s" {
	zone_id = "%[2]s"
	enabled = %[3]t
}
`, rnd, zoneID, enable)
}

package zero_trust_gateway_certificate_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv6 "github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_gateway_certificate", &resource.Sweeper{
		Name: "cloudflare_zero_trust_gateway_certificate",
		F:    testSweepCloudflareZeroTrustGatewayCertificate,
	})
}

// testSweepCloudflareZeroTrustGatewayCertificate removes gateway-managed certificates
// that are not actively in use by any Gateway policy.
//
// Gateway certificates do not have user-controlled names, so the standard
// test-prefix filter (cftftest) cannot be applied. Instead this sweeper targets:
//   - type == "gateway_managed" (never touches BYO/custom certificates)
//   - in_use == false (never deletes certificates actively used by a policy)
//
// The Cloudflare API requires certificates to be deactivated before deletion.
// This sweeper calls Deactivate() for any certificate with binding_status "available"
// before calling Delete().
//
// Note: Zero Trust resources require API key authentication (CLOUDFLARE_EMAIL +
// CLOUDFLARE_API_KEY). The sweeper skips if CLOUDFLARE_API_TOKEN is set.
//
// Run with: go test ./internal/services/zero_trust_gateway_certificate/ -v -sweep=all
//
// Requires:
//   - CLOUDFLARE_ACCOUNT_ID
//   - CLOUDFLARE_EMAIL + CLOUDFLARE_API_KEY (not CLOUDFLARE_API_TOKEN)
func testSweepCloudflareZeroTrustGatewayCertificate(r string) error {
	ctx := context.Background()

	// Zero Trust resources do not support API token authentication.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		tflog.Info(ctx, "Skipping zero_trust_gateway_certificate sweep: CLOUDFLARE_API_TOKEN is set (requires API key auth)")
		return nil
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		tflog.Info(ctx, "Skipping zero_trust_gateway_certificate sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	client := acctest.SharedClient()

	page, err := client.ZeroTrust.Gateway.Certificates.List(ctx, zero_trust.GatewayCertificateListParams{
		AccountID: cfv6.F(accountID),
	})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to list gateway certificates: %s", err))
		return fmt.Errorf("error listing zero_trust_gateway_certificate for sweep: %w", err)
	}

	deletedCount := 0
	failedCount := 0

	for page != nil && len(page.Result) > 0 {
		for _, cert := range page.Result {
			// Only sweep gateway-managed certificates; never touch BYO/custom certs.
			if cert.Type != zero_trust.GatewayCertificateListResponseTypeGatewayManaged {
				tflog.Debug(ctx, fmt.Sprintf("Skipping non-gateway-managed certificate: %s (type: %s)", cert.ID, cert.Type))
				continue
			}

			// Never delete a certificate that is actively used by a Gateway policy.
			if cert.InUse {
				tflog.Debug(ctx, fmt.Sprintf("Skipping in-use gateway certificate: %s", cert.ID))
				continue
			}

			// The API requires deactivation before deletion.
			if cert.BindingStatus == zero_trust.GatewayCertificateListResponseBindingStatusAvailable {
				tflog.Info(ctx, fmt.Sprintf("Deactivating gateway certificate before deletion: %s", cert.ID))
				_, err := client.ZeroTrust.Gateway.Certificates.Deactivate(
					ctx,
					cert.ID,
					zero_trust.GatewayCertificateDeactivateParams{
						AccountID: cfv6.F(accountID),
						Body:      struct{}{},
					},
				)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf("Failed to deactivate gateway certificate %s: %s", cert.ID, err))
					failedCount++
					continue
				}
				tflog.Info(ctx, fmt.Sprintf("Deactivated gateway certificate: %s", cert.ID))
			}

			tflog.Info(ctx, fmt.Sprintf("Deleting gateway certificate: %s", cert.ID))
			_, err := client.ZeroTrust.Gateway.Certificates.Delete(
				ctx,
				cert.ID,
				zero_trust.GatewayCertificateDeleteParams{
					AccountID: cfv6.F(accountID),
				},
			)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete gateway certificate %s: %s", cert.ID, err))
				failedCount++
				continue
			}

			deletedCount++
			tflog.Info(ctx, fmt.Sprintf("Deleted gateway certificate: %s", cert.ID))
		}

		page, err = page.GetNextPage()
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to fetch next page of gateway certificates: %s", err))
			break
		}
	}

	tflog.Info(ctx, fmt.Sprintf("Completed sweeping zero_trust_gateway_certificate: deleted %d, failed %d", deletedCount, failedCount))
	return nil
}

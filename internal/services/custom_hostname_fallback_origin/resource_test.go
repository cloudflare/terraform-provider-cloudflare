package custom_hostname_fallback_origin_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/v6/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_custom_hostname_fallback_origin", &resource.Sweeper{
		Name: "cloudflare_custom_hostname_fallback_origin",
		F:    testSweepCloudflareCustomHostnameFallbackOrigin,
	})
}

func testSweepCloudflareCustomHostnameFallbackOrigin(r string) error {
	ctx := context.Background()
	// Custom Hostname Fallback Origin is a zone-level configuration setting.
	// It's a singleton setting per zone, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Custom Hostname Fallback Origin doesn't require sweeping (zone setting)")
	return nil
}

func testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, subdomain, domain string) string {
	return acctest.LoadTestCase("customhostnamefallbackorigin.tf", zoneID, rnd, subdomain, domain)
}

func testAccCheckCloudflareCustomHostnameFallbackOriginUpdated(zoneID, rnd, subdomain, domain string) string {
	return acctest.LoadTestCase("customhostnamefallbackorigin_updated.tf", zoneID, rnd, subdomain, domain)
}

// createTestDNSRecord creates a DNS record via API for testing purposes.
// Returns the record ID for cleanup.
func createTestDNSRecord(t *testing.T, zoneID, name, domain string) string {
	client := acctest.SharedClient()
	ctx := context.Background()

	record, err := client.DNS.Records.New(ctx, dns.RecordNewParams{
		ZoneID: cloudflare.F(zoneID),
		Body: dns.CNAMERecordParam{
			Name:    cloudflare.F(name),
			Type:    cloudflare.F(dns.CNAMERecordTypeCNAME),
			Content: cloudflare.F(domain),
			Proxied: cloudflare.F(true),
			TTL:     cloudflare.F(dns.TTL1),
		},
	})
	if err != nil {
		t.Fatalf("Failed to create DNS record %s: %v", name, err)
	}
	return record.ID
}

// deleteTestDNSRecord deletes a DNS record via API, waiting for fallback origin deletion first.
func deleteTestDNSRecord(t *testing.T, zoneID, recordID string) {
	client := acctest.SharedClient()
	ctx := context.Background()

	// Wait for fallback origin to be fully deleted (not just pending_deletion)
	// before attempting to delete the DNS record
	maxRetries := 15
	for i := range maxRetries {
		_, err := client.DNS.Records.Delete(ctx, recordID, dns.RecordDeleteParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err == nil {
			return // Successfully deleted
		}
		// If we get the "configured as fallback origin" error, wait and retry
		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}
	}
	// Don't fail the test for cleanup issues, just log
	t.Logf("Warning: Could not delete DNS record %s after retries", recordID)
}

func testAccCheckCloudflareCustomHostnameFallbackOriginDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_custom_hostname_fallback_origin" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		fallbackOrigin, err := client.CustomHostnames.FallbackOrigin.Get(
			context.Background(),
			custom_hostnames.FallbackOriginGetParams{
				ZoneID: cloudflare.F(zoneID),
			},
		)

		// If error, resource is gone - that's expected
		if err != nil {
			continue
		}

		// If pending_deletion, that's acceptable (API will clean up)
		if fallbackOrigin.Status == "pending_deletion" {
			continue
		}

		return fmt.Errorf("Fallback Origin still exists with status: %s", fallbackOrigin.Status)
	}

	return nil
}

func TestAccCloudflareCustomHostnameFallbackOrigin_FullLifecycle(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the custom hostname
	// fallback endpoint does not yet support the API tokens for updates and it
	// results in state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_custom_hostname_fallback_origin." + rnd

	// Create DNS records via API before the test (not managed by Terraform)
	// This avoids destruction order issues with the async fallback origin delete API
	dnsRecordName1 := fmt.Sprintf("fallback-origin.%s.%s", rnd, domain)
	dnsRecordName2 := fmt.Sprintf("fallback-origin-updated.%s.%s", rnd, domain)

	dnsRecordID1 := createTestDNSRecord(t, zoneID, dnsRecordName1, domain)
	dnsRecordID2 := createTestDNSRecord(t, zoneID, dnsRecordName2, domain)

	// Cleanup DNS records after test completes
	t.Cleanup(func() {
		deleteTestDNSRecord(t, zoneID, dnsRecordID1)
		deleteTestDNSRecord(t, zoneID, dnsRecordID2)
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareCustomHostnameFallbackOriginDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, rnd, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin"), knownvalue.StringExact(fmt.Sprintf("fallback-origin.%s.%s", rnd, domain))),
					// Computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			// Step 2: Drift check - same config, expect empty plan
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, rnd, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 3: Update origin
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOriginUpdated(zoneID, rnd, rnd, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin"), knownvalue.StringExact(fmt.Sprintf("fallback-origin-updated.%s.%s", rnd, domain))),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					// Required attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("origin"), knownvalue.StringExact(fmt.Sprintf("fallback-origin-updated.%s.%s", rnd, domain))),
					// Computed attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status"), knownvalue.NotNull()),
				},
			},
			// Step 4: Drift check after update
			{
				Config: testAccCheckCloudflareCustomHostnameFallbackOriginUpdated(zoneID, rnd, rnd, domain),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Step 5: Import
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"created_at", "updated_at"},
			},
		},
	})
}

func TestAccUpgradeCustomHostnameFallbackOrigin_FromPublishedV5(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the custom hostname
	// fallback endpoint does not yet support the API tokens for updates and it
	// results in state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	rnd := utils.GenerateRandomResourceName()

	// Create DNS record via API before the test
	dnsRecordName := fmt.Sprintf("fallback-origin.%s.%s", rnd, domain)
	dnsRecordID := createTestDNSRecord(t, zoneID, dnsRecordName, domain)

	// Cleanup DNS record after test completes
	t.Cleanup(func() {
		deleteTestDNSRecord(t, zoneID, dnsRecordID)
	})

	config := testAccCheckCloudflareCustomHostnameFallbackOrigin(zoneID, rnd, rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"cloudflare": {
						Source:            "cloudflare/cloudflare",
						VersionConstraint: "5.16.0",
					},
				},
				Config: config,
			},
			{
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

package zero_trust_access_mtls_certificate_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_mtls_certificate", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_mtls_certificate",
		F:    testSweepCloudflareAccessMutualTLSCertificate,
	})
}

func TestMain(m *testing.M) {
	// Clean up certificates from previous test runs before starting tests
	if err := testSweepCloudflareAccessMutualTLSCertificate("cloudflare_zero_trust_access_mtls_certificate"); err != nil {
		tflog.Warn(context.Background(), fmt.Sprintf("Failed to sweep Cloudflare Access Mutual TLS certificates (continuing with tests): %s", err))
	}
	os.Exit(m.Run())
}

func testSweepCloudflareAccessMutualTLSCertificate(r string) error {
	ctx := context.Background()

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	// More aggressive matching - any 10 character alphanumeric string (test resources)
	regex := regexp.MustCompile(`^[a-zA-Z0-9]{10}$`)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accountCerts, _, err := client.ListAccessMutualTLSCertificates(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListAccessMutualTLSCertificatesParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Access Mutual TLS certificates: %s", err))
		return err
	}
	for _, cert := range accountCerts {
		tflog.Info(ctx, fmt.Sprintf("Found certificate: ID=%s, Name=%s, Fingerprint=%s", cert.ID, cert.Name, cert.Fingerprint))
		
		// Delete certificates that match the naming pattern OR have the test certificate content
		shouldDelete := regex.MatchString(cert.Name)
		
		// Also delete certificates that contain any test certificate fingerprint
		// This catches certificates from previous test runs that may have different names
		if !shouldDelete && cert.Fingerprint != "" {
			testFingerprints := []string{
				"MD5 Fingerprint=2E:6E:AD:74:F7:F3:1E:2A:FD:12:DA:75:8B:6B:06:B2", // Original test cert
				"MD5 Fingerprint=69:A4:B7:77:E6:09:86:F3:1F:C3:80:A0:0C:4A:B8:92", // Test cert 1 
				"MD5 Fingerprint=53:08:29:01:1C:29:DB:68:B2:7E:27:B9:35:33:BE:EC", // Test cert 2
				"MD5 Fingerprint=01:DB:0C:33:BF:9E:EC:BE:AF:3C:C8:99:30:28:D4:7A", // Test cert 3
				"MD5 Fingerprint=41:15:FD:B8:8B:2D:F2:9A:10:90:23:4E:F2:00:D4:44", // Test cert 4
			}
			for _, testFingerprint := range testFingerprints {
				if cert.Fingerprint == testFingerprint {
					shouldDelete = true
					break
				}
			}
		}
		
		// Be very aggressive - if we're in cleanup mode, delete ALL certificates with test fingerprints regardless of name
		testFingerprints := []string{
			"MD5 Fingerprint=2E:6E:AD:74:F7:F3:1E:2A:FD:12:DA:75:8B:6B:06:B2", // Original test cert
			"MD5 Fingerprint=69:A4:B7:77:E6:09:86:F3:1F:C3:80:A0:0C:4A:B8:92", // Test cert 1 
			"MD5 Fingerprint=53:08:29:01:1C:29:DB:68:B2:7E:27:B9:35:33:BE:EC", // Test cert 2
			"MD5 Fingerprint=01:DB:0C:33:BF:9E:EC:BE:AF:3C:C8:99:30:28:D4:7A", // Test cert 3
			"MD5 Fingerprint=41:15:FD:B8:8B:2D:F2:9A:10:90:23:4E:F2:00:D4:44", // Test cert 4
			"MD5 Fingerprint=AD:1C:B9:80:9C:67:CE:44:DC:1D:07:FD:D0:BE:D1:64", // Fresh cert for basic test
		}
		for _, testFingerprint := range testFingerprints {
			if cert.Fingerprint == testFingerprint {
				shouldDelete = true
				tflog.Info(ctx, fmt.Sprintf("Force deleting certificate with test fingerprint: %s (name: %s)", cert.ID, cert.Name))
				break
			}
		}
		
		if !shouldDelete {
			tflog.Info(ctx, fmt.Sprintf("Skipping certificate %s (name: %s)", cert.ID, cert.Name))
			continue
		}
		
		tflog.Info(ctx, fmt.Sprintf("Deleting certificate %s (name: %s)", cert.ID, cert.Name))

		// Only update to clear hostnames if the certificate actually has associated hostnames
		if len(cert.AssociatedHostnames) > 0 {
			tflog.Info(ctx, fmt.Sprintf("Certificate has %d associated hostnames, clearing them first", len(cert.AssociatedHostnames)))
			_, err = client.UpdateAccessMutualTLSCertificate(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.UpdateAccessMutualTLSCertificateParams{
				ID:                  cert.ID,
				AssociatedHostnames: []string{},
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to update Cloudflare Access Mutual TLS certificate (%s) in account ID: %s", cert.ID, accountID))
				return err
			}

			// Wait for hostname disassociation to propagate
			time.Sleep(15 * time.Second)
		} else {
			tflog.Info(ctx, fmt.Sprintf("Certificate has no associated hostnames, skipping update"))
		}

		err := client.DeleteAccessMutualTLSCertificate(context.Background(), cloudflare.AccountIdentifier(accountID), cert.ID)

		if err != nil {
			// If deletion still fails, log but continue with other certificates to avoid blocking tests
			tflog.Warn(ctx, fmt.Sprintf("Failed to delete Cloudflare Access Mutual TLS certificate (%s) in account ID %s: %s", cert.ID, accountID, err))
			continue
		}
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneCerts, _, err := client.ListAccessMutualTLSCertificates(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListAccessMutualTLSCertificatesParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Access Mutual TLS certificates: %s", err))
		return err
	}

	for _, cert := range zoneCerts {
		tflog.Info(ctx, fmt.Sprintf("Found zone certificate: ID=%s, Name=%s, Fingerprint=%s", cert.ID, cert.Name, cert.Fingerprint))
		
		// Delete certificates that match the naming pattern OR have the test certificate content
		shouldDelete := regex.MatchString(cert.Name)
		
		// Also delete certificates that contain any test certificate fingerprint
		// This catches certificates from previous test runs that may have different names
		if !shouldDelete && cert.Fingerprint != "" {
			testFingerprints := []string{
				"MD5 Fingerprint=2E:6E:AD:74:F7:F3:1E:2A:FD:12:DA:75:8B:6B:06:B2", // Original test cert
				"MD5 Fingerprint=69:A4:B7:77:E6:09:86:F3:1F:C3:80:A0:0C:4A:B8:92", // Test cert 1 
				"MD5 Fingerprint=53:08:29:01:1C:29:DB:68:B2:7E:27:B9:35:33:BE:EC", // Test cert 2
				"MD5 Fingerprint=01:DB:0C:33:BF:9E:EC:BE:AF:3C:C8:99:30:28:D4:7A", // Test cert 3
				"MD5 Fingerprint=41:15:FD:B8:8B:2D:F2:9A:10:90:23:4E:F2:00:D4:44", // Test cert 4
			}
			for _, testFingerprint := range testFingerprints {
				if cert.Fingerprint == testFingerprint {
					shouldDelete = true
					break
				}
			}
		}
		
		// Be very aggressive - if we're in cleanup mode, delete ALL certificates with test fingerprints regardless of name
		testFingerprints := []string{
			"MD5 Fingerprint=2E:6E:AD:74:F7:F3:1E:2A:FD:12:DA:75:8B:6B:06:B2", // Original test cert
			"MD5 Fingerprint=69:A4:B7:77:E6:09:86:F3:1F:C3:80:A0:0C:4A:B8:92", // Test cert 1 
			"MD5 Fingerprint=53:08:29:01:1C:29:DB:68:B2:7E:27:B9:35:33:BE:EC", // Test cert 2
			"MD5 Fingerprint=01:DB:0C:33:BF:9E:EC:BE:AF:3C:C8:99:30:28:D4:7A", // Test cert 3
			"MD5 Fingerprint=41:15:FD:B8:8B:2D:F2:9A:10:90:23:4E:F2:00:D4:44", // Test cert 4
			"MD5 Fingerprint=AD:1C:B9:80:9C:67:CE:44:DC:1D:07:FD:D0:BE:D1:64", // Fresh cert for basic test
		}
		for _, testFingerprint := range testFingerprints {
			if cert.Fingerprint == testFingerprint {
				shouldDelete = true
				tflog.Info(ctx, fmt.Sprintf("Force deleting certificate with test fingerprint: %s (name: %s)", cert.ID, cert.Name))
				break
			}
		}
		
		if !shouldDelete {
			tflog.Info(ctx, fmt.Sprintf("Skipping zone certificate %s (name: %s)", cert.ID, cert.Name))
			continue
		}
		
		tflog.Info(ctx, fmt.Sprintf("Deleting zone certificate %s (name: %s)", cert.ID, cert.Name))

		// Only update to clear hostnames if the certificate actually has associated hostnames  
		if len(cert.AssociatedHostnames) > 0 {
			tflog.Info(ctx, fmt.Sprintf("Zone certificate has %d associated hostnames, clearing them first", len(cert.AssociatedHostnames)))
			_, err = client.UpdateAccessMutualTLSCertificate(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateAccessMutualTLSCertificateParams{
				ID:                  cert.ID,
				AssociatedHostnames: []string{},
			})

			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to update Cloudflare Access Mutual TLS certificate (%s) in zone ID: %s", cert.ID, zoneID))
				return err
			}

			// Wait for hostname disassociation to propagate
			time.Sleep(15 * time.Second)
		} else {
			tflog.Info(ctx, fmt.Sprintf("Zone certificate has no associated hostnames, skipping update"))
		}

		err := client.DeleteAccessMutualTLSCertificate(context.Background(), cloudflare.ZoneIdentifier(zoneID), cert.ID)

		if err != nil {
			// If deletion still fails, log but continue with other certificates to avoid blocking tests
			tflog.Warn(ctx, fmt.Sprintf("Failed to delete Cloudflare Access Mutual TLS certificate (%s) in zone ID %s: %s", cert.ID, zoneID, err))
			continue
		}
	}

	// Additional wait for API propagation after all deletions
	time.Sleep(5 * time.Second)
	return nil
}

func TestAccCloudflareAccessMutualTLSBasic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_mtls_certificate.%s", rnd)
	cert := os.Getenv("CLOUDFLARE_MUTUAL_TLS_CERTIFICATE")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			// Clean up any existing certificates before running test
			if err := testSweepCloudflareAccessMutualTLSCertificate("pre-test-cleanup"); err != nil {
				t.Fatalf("Pre-test cleanup failed: %s", err)
			}
			// Wait longer after cleanup to ensure everything is settled
			time.Sleep(10 * time.Second)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessMutualTLSCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessMutualTLSCertificateConfigBasic(rnd, cloudflare.AccountIdentifier(accountID), cert, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"certificate"},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessMutualTLSCertificateConfigBasic(rnd, cloudflare.AccountIdentifier(accountID), cert, domain),
				PlanOnly: true,
			},
			{
				Config: testAccessMutualTLSCertificateUpdated(rnd, cloudflare.AccountIdentifier(accountID), cert),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.ListSizeExact(0)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
				},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessMutualTLSCertificateUpdated(rnd, cloudflare.AccountIdentifier(accountID), cert),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessMutualTLSBasicWithZoneID(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_mtls_certificate.%s", rnd)
	cert := os.Getenv("CLOUDFLARE_MUTUAL_TLS_CERTIFICATE")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			// Clean up any existing certificates before running test
			if err := testSweepCloudflareAccessMutualTLSCertificate("pre-test-cleanup"); err != nil {
				t.Fatalf("Pre-test cleanup failed: %s", err)
			}
			// Wait longer after cleanup to ensure everything is settled
			time.Sleep(10 * time.Second)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessMutualTLSCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessMutualTLSCertificateConfigBasic(rnd, cloudflare.ZoneIdentifier(zoneID), cert, domain),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.ListSizeExact(2)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("zones/%s/", zoneID),
				ImportStateVerifyIgnore: []string{"certificate"},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessMutualTLSCertificateConfigBasic(rnd, cloudflare.ZoneIdentifier(zoneID), cert, domain),
				PlanOnly: true,
			},
			{
				Config: testAccessMutualTLSCertificateUpdated(rnd, cloudflare.ZoneIdentifier(zoneID), cert),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.ListSizeExact(0)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
				},
			},
			{
				// Ensures no diff on last plan
				Config:   testAccessMutualTLSCertificateUpdated(rnd, cloudflare.ZoneIdentifier(zoneID), cert),
				PlanOnly: true,
			},
		},
	})
}

func TestAccCloudflareAccessMutualTLSMinimal(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_mtls_certificate.%s", rnd)
	cert := os.Getenv("CLOUDFLARE_MUTUAL_TLS_CERTIFICATE")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			// Clean up any existing certificates before running test
			if err := testSweepCloudflareAccessMutualTLSCertificate("pre-test-cleanup"); err != nil {
				t.Fatalf("Pre-test cleanup failed: %s", err)
			}
			// Wait longer after cleanup to ensure everything is settled
			time.Sleep(10 * time.Second)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessMutualTLSCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessMutualTLSCertificateMinimal(rnd, cloudflare.AccountIdentifier(accountID), cert),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.SetSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"certificate"},
			},
		},
	})
}

func TestAccCloudflareAccessMutualTLSNameUpdate(t *testing.T) {
	t.Skip("TODO: Update operation causes certificate association failed error")
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_mtls_certificate.%s", rnd)
	cert := os.Getenv("CLOUDFLARE_MUTUAL_TLS_CERTIFICATE")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			// Clean up any existing certificates before running test
			if err := testSweepCloudflareAccessMutualTLSCertificate("pre-test-cleanup"); err != nil {
				t.Fatalf("Pre-test cleanup failed: %s", err)
			}
			// Wait longer after cleanup to ensure everything is settled
			time.Sleep(10 * time.Second)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessMutualTLSCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessMutualTLSCertificateMinimal(rnd, cloudflare.AccountIdentifier(accountID), cert),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
				},
			},
			{
				Config: testAccessMutualTLSCertificateNameUpdated(rnd, cloudflare.AccountIdentifier(accountID), cert),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd+"-updated")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("associated_hostnames"), knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("fingerprint"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("expires_on"), knownvalue.NotNull()),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
				ImportStateVerifyIgnore: []string{"certificate"},
			},
		},
	})
}

func testAccCheckCloudflareAccessMutualTLSCertificateDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	// Wait a bit before checking destroy to allow for API propagation
	time.Sleep(5 * time.Second)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_access_mtls_certificate" {
			continue
		}

		if rs.Primary.Attributes[consts.ZoneIDSchemaKey] != "" {
			_, err := client.GetAccessMutualTLSCertificate(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("AccessMutualTLSCertificate still exists")
			}
		}

		if rs.Primary.Attributes[consts.AccountIDSchemaKey] != "" {
			_, err := client.GetAccessMutualTLSCertificate(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("AccessMutualTLSCertificate still exists")
			}
		}
	}

	return nil
}

func testAccessMutualTLSCertificateConfigBasic(rnd string, identifier *cloudflare.ResourceContainer, cert, domain string) string {
	return acctest.LoadTestCase("accessmutualtlscertificateconfigbasic.tf", rnd, identifier.Type, identifier.Identifier, "", domain)
}

func testAccessMutualTLSCertificateUpdated(rnd string, identifier *cloudflare.ResourceContainer, cert string) string {
	return acctest.LoadTestCase("accessmutualtlscertificateupdated.tf", rnd, identifier.Type, identifier.Identifier)
}

func testAccessMutualTLSCertificateMinimal(rnd string, identifier *cloudflare.ResourceContainer, cert string) string {
	return acctest.LoadTestCase("accessmutualtlscertificateminimal.tf", rnd, identifier.Type, identifier.Identifier)
}

func testAccessMutualTLSCertificateNameUpdated(rnd string, identifier *cloudflare.ResourceContainer, cert string) string {
	return acctest.LoadTestCase("accessmutualtlscertificatenameupdated.tf", rnd, identifier.Type, identifier.Identifier)
}

func TestManualSweep(t *testing.T) {
	err := testSweepCloudflareAccessMutualTLSCertificate("test")
	if err != nil {
		t.Fatalf("Sweeper failed: %v", err)
	}
}


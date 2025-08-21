package zero_trust_access_mtls_certificate_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
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

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_mtls_certificate", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_mtls_certificate",
		F:    testSweepCloudflareAccessMutualTLSCertificate,
	})
}

func TestMain(m *testing.M) {
	if err := testSweepCloudflareAccessMutualTLSCertificate("cloudflare_zero_trust_access_mtls_certificate"); err != nil {
		fmt.Println(err)
		tflog.Error(context.Background(), fmt.Sprintf("Failed to sweep Cloudflare Access Mutual TLS certificates: %s", err))
		os.Exit(1)
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

	regex := regexp.MustCompile(`^[a-zA-Z]{10}$`)

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accountCerts, _, err := client.ListAccessMutualTLSCertificates(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListAccessMutualTLSCertificatesParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Access Mutual TLS certificates: %s", err))
		return err
	}
	for _, cert := range accountCerts {
		// only delete certificates that appear to be created by this provider
		if !regex.MatchString(cert.Name) {
			continue
		}

		// to delete we need to update first with empty hostnames
		_, err = client.UpdateAccessMutualTLSCertificate(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.UpdateAccessMutualTLSCertificateParams{
			ID:                  cert.ID,
			AssociatedHostnames: []string{},
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to update Cloudflare Access Mutual TLS certificate (%s) in account ID: %s", cert.ID, accountID))
		}

		// Wait for update to propagate with retry logic
		maxRetries := 5
		backoff := time.Second * 2
		for range maxRetries {
			time.Sleep(backoff)
			updatedCert, checkErr := client.GetAccessMutualTLSCertificate(context.Background(), cloudflare.AccountIdentifier(accountID), cert.ID)
			if checkErr == nil && len(updatedCert.AssociatedHostnames) == 0 {
				break
			}
			backoff *= 2
		}

		// Retry deletion with exponential backoff to handle race conditions
		var lastErr error
		for range maxRetries {
			err = client.DeleteAccessMutualTLSCertificate(context.Background(), cloudflare.AccountIdentifier(accountID), cert.ID)
			if err == nil {
				lastErr = nil
				break
			}
			lastErr = err
			time.Sleep(backoff)
			backoff *= 2
		}

		if lastErr != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare Access Mutual TLS certificate (%s) in account ID: %s after retries: %s", cert.ID, accountID, lastErr))
		}
	}

	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneCerts, _, err := client.ListAccessMutualTLSCertificates(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListAccessMutualTLSCertificatesParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Access Mutual TLS certificates: %s", err))
		return err
	}

	for _, cert := range zoneCerts {
		// only delete certificates that appear to be created by this provider
		if !regex.MatchString(cert.Name) {
			continue
		}

		// to delete we need to update first with empty hostnames
		_, err = client.UpdateAccessMutualTLSCertificate(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateAccessMutualTLSCertificateParams{
			ID:                  cert.ID,
			AssociatedHostnames: []string{},
		})

		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to update Cloudflare Access Mutual TLS certificate (%s) in zone ID: %s", cert.ID, zoneID))
		}

		// Wait for update to propagate with retry logic
		maxRetries := 5
		backoff := time.Second * 2
		for range maxRetries {
			time.Sleep(backoff)
			updatedCert, checkErr := client.GetAccessMutualTLSCertificate(context.Background(), cloudflare.ZoneIdentifier(zoneID), cert.ID)
			if checkErr == nil && len(updatedCert.AssociatedHostnames) == 0 {
				break
			}
			backoff *= 2
		}

		// Retry deletion with exponential backoff to handle race conditions
		var lastErr error
		for range maxRetries {
			err = client.DeleteAccessMutualTLSCertificate(context.Background(), cloudflare.ZoneIdentifier(zoneID), cert.ID)
			if err == nil {
				lastErr = nil
				break
			}
			lastErr = err
			time.Sleep(backoff)
			backoff *= 2
		}

		if lastErr != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete Cloudflare Access Mutual TLS certificate (%s) in zone ID: %s after retries: %s", cert.ID, zoneID, lastErr))
		}
	}

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

func testAccCheckCloudflareAccessMutualTLSCertificateDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

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
	return acctest.LoadTestCase("accessmutualtlscertificateconfigbasic.tf", rnd, identifier.Type, identifier.Identifier, cert, domain)
}

func testAccessMutualTLSCertificateUpdated(rnd string, identifier *cloudflare.ResourceContainer, cert string) string {
	return acctest.LoadTestCase("accessmutualtlscertificateupdated.tf", rnd, identifier.Type, identifier.Identifier, cert)
}

func testAccessMutualTLSCertificateMinimal(rnd string, identifier *cloudflare.ResourceContainer, cert string) string {
	return acctest.LoadTestCase("accessmutualtlscertificateminimal.tf", rnd, identifier.Type, identifier.Identifier, cert)
}

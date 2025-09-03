package zero_trust_access_mtls_certificate_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
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
	resource.TestMain(m)
}

func testSweepCloudflareAccessMutualTLSCertificate(r string) error {
	ctx := context.Background()

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
		return clientErr
	}

	// In test environment, be more aggressive about cleanup to prevent certificate conflicts
	// This prevents "certificate already exists" errors in tests
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accountCerts, _, err := client.ListAccessMutualTLSCertificates(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListAccessMutualTLSCertificatesParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch Cloudflare Access Mutual TLS certificates: %s", err))
		return err
	}

	tflog.Info(ctx, fmt.Sprintf("Found %d certificates in account, cleaning all to prevent test conflicts", len(accountCerts)))

	for _, cert := range accountCerts {
		tflog.Info(ctx, fmt.Sprintf("Deleting certificate: Name=%s, ID=%s, Hostnames=%v", cert.Name, cert.ID, cert.AssociatedHostnames))

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

	tflog.Info(ctx, fmt.Sprintf("Found %d certificates in zone, cleaning all to prevent test conflicts", len(zoneCerts)))

	for _, cert := range zoneCerts {
		tflog.Info(ctx, fmt.Sprintf("Deleting zone certificate: Name=%s, ID=%s, Hostnames=%v", cert.Name, cert.ID, cert.AssociatedHostnames))

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
	waitBetweenTests(t, false)
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_mtls_certificate.%s", rnd)
	// Use unique certificate for basic test to avoid conflicts
	cert := `-----BEGIN CERTIFICATE-----
MIIDsTCCApmgAwIBAgIUTaqzZvhPgvlEKYHsKdHKvshta/YwDQYJKoZIhvcNAQEL
BQAwaDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQHDA1TYW4gRnJh
bmNpc2NvMRMwEQYDVQQKDApCYXNpYyBUZXN0MR8wHQYDVQQDDBZiYXNpYy10ZXN0
LmV4YW1wbGUuY29tMB4XDTI1MDgyNTEyNDUzMFoXDTI2MDgyNTEyNDUzMFowaDEL
MAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQHDA1TYW4gRnJhbmNpc2Nv
MRMwEQYDVQQKDApCYXNpYyBUZXN0MR8wHQYDVQQDDBZiYXNpYy10ZXN0LmV4YW1w
bGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkmzUoLpL0Upy
luplSqgz0ec44trHhr7dhh0Omam2BI9fZziRWqFmF8sodieIvL51Wzko36WYttqY
6dv8q3+GLoVN2X0BaF2As0hCphFiFC1I3k4AlHwUx96SUlK4cvyS46bP+nlO6zWG
JzH4i80Cp6LkFXI5j84yQcAqCwxyoEVlgQLz+e5bgs+VufMDyo9TQ9GXy9EZ1Ysi
bG6DvxNnTDAf0P5c1t5BXe1EeB0413lRN7mrpgH0kxM43SwcMWx6dx3IDLEv82oo
wpc98FPqeLK2ZN62luO0ZwBhmECBhl1ABCmOF/n6Yavx8rjLdN/MBO7/AODsyjKl
eE2HKhYrAQIDAQABo1MwUTAdBgNVHQ4EFgQUgqmZSicLOQY9zOKORrF2XACXzpMw
HwYDVR0jBBgwFoAUgqmZSicLOQY9zOKORrF2XACXzpMwDwYDVR0TAQH/BAUwAwEB
/zANBgkqhkiG9w0BAQsFAAOCAQEAJfFylVu/lVDAqUpgWOe0ZWF0djEiw+wphgC6
2EshvccoW3YLku0XPPc1+DZEjMtCMh06woaH1R8koWXnOjs2com2UA2ccDB3mkZG
Wl0sBxQLPu/Gj6jKbXnWm24morGzYWyZlLNP9178tVdhMgNMOR50qB6QsQLcDRHu
Tj/DLAsk96PEw3AZ/Lad3oJs5me0uVxPxrAcqKtrAOx2CpMxmvFNZ5ZiKr8b8uJQ
0X2o8qiA8Qd4qMTnI8lHpuuGsun/RIiBBxkzHDrc0qZVEm2sqE97tA0vBczcn7fN
jzhIPJ0iyPgZhFlsHjGxWghkaLqdCdtDOSdb6SepEZHaq32j/A==
-----END CERTIFICATE-----`
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
	waitBetweenTests(t, true)
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_mtls_certificate.%s", rnd)
	// Use unique certificate for zone test to avoid conflicts
	cert := `-----BEGIN CERTIFICATE-----
MIIDrTCCApWgAwIBAgIUU3Ss5iR+GdD6EeSfS+12wbzTvK4wDQYJKoZIhvcNAQEL
BQAwZjELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQHDA1TYW4gRnJh
bmNpc2NvMRIwEAYDVQQKDAlab25lIFRlc3QxHjAcBgNVBAMMFXpvbmUtdGVzdC5l
eGFtcGxlLmNvbTAeFw0yNTA4MjUxMjQ1MjRaFw0yNjA4MjUxMjQ1MjRaMGYxCzAJ
BgNVBAYTAlVTMQswCQYDVQQIDAJDQTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzES
MBAGA1UECgwJWm9uZSBUZXN0MR4wHAYDVQQDDBV6b25lLXRlc3QuZXhhbXBsZS5j
b20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC3uxA57+zYxNQzRo9B
O3TR2z82DJjEfkMjZm5OlutG7Mi5J5HIa5rexBkn+aOYtTpu1idfV3RESsPXclYg
kGI5Ywo8PR4DjFD7Tu/0ZgdeftaThuc1PQnxqP+KZ3Azjke3o6KfhqtNiRO/qLWL
sEaIKpHlA7RFWytqONTwG+FGYbvX3h+ox/DQA4vdzi3xHCyf1XsWr1uoPasIZn1L
1sHI5+wD2JHpOxWli9VplPXVIKGMPnUmysawn2H5kcTJAEJJoRHDTBIlfR5MnLG2
vnaUE9cKZ96pUIrl+4El1IjxGd30tbld3YQYgCrhbVt/pcQXyUmhKgrZv/QrV9o2
v2T7AgMBAAGjUzBRMB0GA1UdDgQWBBR4RNyAV3GbgaAi75iRWmJbM5mbMDAfBgNV
HSMEGDAWgBR4RNyAV3GbgaAi75iRWmJbM5mbMDAPBgNVHRMBAf8EBTADAQH/MA0G
CSqGSIb3DQEBCwUAA4IBAQB6ewMW4W0znrb7AcqVCnc6nz7mFY+uwVldJN3nywW4
iO1TAqUczEhzhEnz3Ly+27o3gkjVPmqAPSge8kLNAKFDJ43xn2G28u7UhSo4b6IN
EPV7b3GIFSBfVd0S8D7dYnlKbE4YAjx+A84+SrqwXg3NfD5ES/XogGE9VWxbN8To
LwvNJwCB23tldWpGiGXwmQVfA0ptA4ys4GoU/ss0BEg9h4BlPjnlwcw5O9cLdZTs
6Dv+537EyG/WsdyNAs/TLeHgM+I9yw4SePhaVq2Zhv/Tz4JryIuDgpp2iwVEnSB/
Ujl4YD+WK1PhWs9G3UVUeGG+93ZVJRC6Em6ZMMMFxQyL
-----END CERTIFICATE-----`
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
	waitBetweenTests(t, false)
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_access_mtls_certificate.%s", rnd)
	// Use unique certificate for minimal test to avoid conflicts
	cert := `-----BEGIN CERTIFICATE-----
MIIDuTCCAqGgAwIBAgIUOLxKkiemLStP0PSReIA36itjxpEwDQYJKoZIhvcNAQEL
BQAwbDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQHDA1TYW4gRnJh
bmNpc2NvMRUwEwYDVQQKDAxNaW5pbWFsIFRlc3QxITAfBgNVBAMMGG1pbmltYWwt
dGVzdC5leGFtcGxlLmNvbTAeFw0yNTA4MjUxMjQ1MTlaFw0yNjA4MjUxMjQ1MTla
MGwxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTEWMBQGA1UEBwwNU2FuIEZyYW5j
aXNjbzEVMBMGA1UECgwMTWluaW1hbCBUZXN0MSEwHwYDVQQDDBhtaW5pbWFsLXRl
c3QuZXhhbXBsZS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDJ
RiCi73nTalZfCy1GKHU4Rj5RWCUDkMPavkIjSQDbdgYpNqQSyMLNZBH0NciavNHC
QweQmAuDOmgxXwc8iDFx/n69UrkV5viu8siuYWnX9LLKmzDUrncUvguSLVl2iH/a
H5ZJya/qnMlu84Ks2iKX2KINbY3A/Py3TvTWDO10W5Ocmi+WQBxzLIHBY4QE/bpV
gtWJF5482v0nT8nQ6ITYOoC5csr4L/ukhrz5460j3pFdE3gRQu5LYhQlU1rygpWA
SU6zCP/jkUxEIPiRRUhZxbM8IPebnRr8oD9TVqjiZpJnbf6LtoxdWRD6YDhqORJg
CcKYO8xVKtzaOAzRbxUbAgMBAAGjUzBRMB0GA1UdDgQWBBSZU4SVPnjPhKQ5PqCV
Va55bSHpxDAfBgNVHSMEGDAWgBSZU4SVPnjPhKQ5PqCVVa55bSHpxDAPBgNVHRMB
Af8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBUG9H2YLQKkUmSe62h972mxL2Y
5LXJBAB0Ys+UyJSGX1pouTZVBYf1h7JJGHgQJT3ITQmn04ipIGLvqYQ2fzvCsAdv
DuQwvWnX70XFQGbZfb8Iu9Yq+6/zHW5iraDqUakDCHybLKgjxX+1+n3fP9xfHSFl
3/wO0yvffxsTMnTFz+4ZVnPl9R948NaeDR+hePZKnabuGozUrRqy/Al7Bcigwy6X
Gsj65OJaR1Y2l1B1gmQlULcbGYV4vQzYosy3mdpd6m8wsP1KZ9mGCPJ/SspW0tiY
SZ0xvbqc2JanR3lB6r5+QAI8KZPjiUInAi/kO0+TAQzQzGLwEgR/cmYHpWsf
-----END CERTIFICATE-----`
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
	// Convert literal \n to actual newlines for proper certificate format
	processedCert := strings.ReplaceAll(cert, "\\n", "\n")
	return acctest.LoadTestCase("accessmutualtlscertificateconfigbasic.tf", rnd, identifier.Type, identifier.Identifier, processedCert, domain)
}

func testAccessMutualTLSCertificateUpdated(rnd string, identifier *cloudflare.ResourceContainer, cert string) string {
	// Convert literal \n to actual newlines for proper certificate format
	processedCert := strings.ReplaceAll(cert, "\\n", "\n")
	return acctest.LoadTestCase("accessmutualtlscertificateupdated.tf", rnd, identifier.Type, identifier.Identifier, processedCert)
}

func testAccessMutualTLSCertificateMinimal(rnd string, identifier *cloudflare.ResourceContainer, cert string) string {
	// Convert literal \n to actual newlines for proper certificate format
	processedCert := strings.ReplaceAll(cert, "\\n", "\n")
	return acctest.LoadTestCase("accessmutualtlscertificateminimal.tf", rnd, identifier.Type, identifier.Identifier, processedCert)
}

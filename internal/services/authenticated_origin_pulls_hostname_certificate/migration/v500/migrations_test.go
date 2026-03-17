package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Migration Test Configuration
//
// IMPORTANT: This resource was renamed from v4's cloudflare_authenticated_origin_pulls_certificate
// (with type="per-hostname") to v5's cloudflare_authenticated_origin_pulls_hostname_certificate.
// Because of this rename, we cannot test v4→v5 migration using ExternalProviders (Terraform validates
// against local schema which doesn't recognize the old resource name).
//
// v4→v5 migration is tested in:
// 1. tf-migrate integration tests (config transformation)
// 2. MoveState logic (state transformation)
//
// These tests focus on v5→v5 version bumps (schema version 1 → 500), which validates that the
// state upgrader correctly handles version increases within v5.

// Embed migration test configuration files
//
//go:embed testdata/v5_hostname_cert.tf
var v5HostnameCertConfig string

//go:embed testdata/v5_hostname_cert_minimal.tf
var v5HostnameCertMinimalConfig string

// TestMigrateAuthenticatedOriginPullsHostnameCertificateBasic tests v5→v5 migration (schema version bump)
// for per-hostname AOP certificates
func TestMigrateAuthenticatedOriginPullsHostnameCertificateBasic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	if domain == "" {
		t.Skip("CLOUDFLARE_DOMAIN must be set for this test")
	}

	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// Generate ephemeral certificate for testing
	expiry := time.Now().Add(time.Hour * 24 * 365)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{domain}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	testConfig := fmt.Sprintf(v5HostnameCertConfig, rnd, zoneID, cert, key, rnd, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
			},
			// Step 2: Run migration (v5→v5 version bump)
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, currentProviderVersion, "v5", "v5", []statecheck.StateCheck{
				// Verify zone_id preserved
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_hostname_certificate."+rnd,
					tfjsonpath.New("zone_id"),
					knownvalue.StringExact(zoneID),
				),
				// Verify certificate preserved
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_hostname_certificate."+rnd,
					tfjsonpath.New("certificate"),
					knownvalue.NotNull(),
				),
				// Verify private_key preserved
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_hostname_certificate."+rnd,
					tfjsonpath.New("private_key"),
					knownvalue.NotNull(),
				),
				// Verify computed fields present
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_hostname_certificate."+rnd,
					tfjsonpath.New("issuer"),
					knownvalue.NotNull(),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_hostname_certificate."+rnd,
					tfjsonpath.New("signature"),
					knownvalue.NotNull(),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_hostname_certificate."+rnd,
					tfjsonpath.New("status"),
					knownvalue.NotNull(),
				),
			}),
		},
	})
}

// TestMigrateAuthenticatedOriginPullsHostnameCertificateMinimal tests v5→v5 migration
// for minimal per-hostname certificate configuration
func TestMigrateAuthenticatedOriginPullsHostnameCertificateMinimal(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	if domain == "" {
		t.Skip("CLOUDFLARE_DOMAIN must be set for this test")
	}

	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()

	// Generate ephemeral certificate for testing
	expiry := time.Now().Add(time.Hour * 24 * 365)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{domain}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	testConfig := fmt.Sprintf(v5HostnameCertMinimalConfig, rnd, zoneID, cert, key, rnd, rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
			acctest.TestAccPreCheck_Domain(t)
		},
		WorkingDir: tmpDir,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with v5 provider
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				Config:                   testConfig,
			},
			// Step 2: Run migration (v5→v5 version bump)
			acctest.MigrationV2TestStep(t, testConfig, tmpDir, currentProviderVersion, "v5", "v5", []statecheck.StateCheck{
				// Verify zone_id preserved
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_hostname_certificate."+rnd,
					tfjsonpath.New("zone_id"),
					knownvalue.StringExact(zoneID),
				),
				// Verify certificate preserved
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_hostname_certificate."+rnd,
					tfjsonpath.New("certificate"),
					knownvalue.NotNull(),
				),
				// Verify private_key preserved
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_hostname_certificate."+rnd,
					tfjsonpath.New("private_key"),
					knownvalue.NotNull(),
				),
			}),
		},
	})
}

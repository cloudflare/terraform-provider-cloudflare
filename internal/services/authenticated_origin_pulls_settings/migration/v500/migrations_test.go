package v500_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

var (
	currentProviderVersion = internal.PackageVersion // Current v5 release
)

// Migration Test Configuration
//
// IMPORTANT: This resource was renamed from v4's cloudflare_authenticated_origin_pulls (zone-wide mode)
// to v5's cloudflare_authenticated_origin_pulls_settings. Because of this rename, we cannot test
// v4→v5 migration using ExternalProviders (Terraform validates against local schema which doesn't
// recognize the old resource name).
//
// v4→v5 migration is tested in:
// 1. tf-migrate integration tests (config transformation)
// 2. State upgrader tests (state transformation)
//
// These tests focus on v5→v5 version bumps (schema version 1 → 500), which validates that the
// state upgrader correctly handles version increases within v5.

// Embed migration test configuration files
//
//go:embed testdata/v5_zone_wide.tf
var v5ZoneWideConfig string

//go:embed testdata/v5_disabled.tf
var v5DisabledConfig string

//go:embed testdata/v5_with_certificate_resource.tf
var v5WithCertificateResourceConfig string

// TestMigrateAuthenticatedOriginPullsSettingsBasic tests v5→v5 migration (schema version bump)
// for basic zone-wide AOP settings
func TestMigrateAuthenticatedOriginPullsSettingsBasic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v5ZoneWideConfig, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
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
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_settings."+rnd,
					tfjsonpath.New("zone_id"),
					knownvalue.StringExact(zoneID),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_settings."+rnd,
					tfjsonpath.New("id"),
					knownvalue.StringExact(zoneID),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_settings."+rnd,
					tfjsonpath.New("enabled"),
					knownvalue.Bool(true),
				),
			}),
		},
	})
}

// TestMigrateAuthenticatedOriginPullsSettingsDisabled tests v5→v5 migration (schema version bump)
// for disabled AOP settings
func TestMigrateAuthenticatedOriginPullsSettingsDisabled(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v5DisabledConfig, rnd, zoneID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_ZoneID(t)
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
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_settings."+rnd,
					tfjsonpath.New("zone_id"),
					knownvalue.StringExact(zoneID),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_settings."+rnd,
					tfjsonpath.New("id"),
					knownvalue.StringExact(zoneID),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_settings."+rnd,
					tfjsonpath.New("enabled"),
					knownvalue.Bool(false),
				),
			}),
		},
	})
}

// TestMigrateAuthenticatedOriginPullsSettingsWithCertificateResource tests v5→v5 migration
// with a certificate resource reference (validates that references are preserved across migration)
func TestMigrateAuthenticatedOriginPullsSettingsWithCertificateResource(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	if domain == "" {
		t.Skip("CLOUDFLARE_DOMAIN must be set for this test")
	}
	rnd := utils.GenerateRandomResourceName()

	// Generate ephemeral certificate for testing
	expiry := time.Now().Add(time.Hour * 24 * 365)
	cert, key, err := utils.GenerateEphemeralCertAndKey([]string{domain}, expiry)
	if err != nil {
		t.Fatalf("Failed to generate certificate: %s", err)
	}

	tmpDir := t.TempDir()
	testConfig := fmt.Sprintf(v5WithCertificateResourceConfig, rnd, zoneID, cert, key)

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
				// Verify AOP settings resource
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_settings."+rnd,
					tfjsonpath.New("zone_id"),
					knownvalue.StringExact(zoneID),
				),
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_settings."+rnd,
					tfjsonpath.New("enabled"),
					knownvalue.Bool(true),
				),
				// Verify certificate resource still exists and is preserved
				statecheck.ExpectKnownValue(
					"cloudflare_authenticated_origin_pulls_certificate."+rnd+"_cert",
					tfjsonpath.New("zone_id"),
					knownvalue.StringExact(zoneID),
				),
			}),
		},
	})
}

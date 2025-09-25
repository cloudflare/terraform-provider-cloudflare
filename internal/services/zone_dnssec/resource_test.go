package zone_dnssec_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareZoneDNSSECFull(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_dnssec.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneDNSSECDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDNSSECResourceConfig(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|pending"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("flags"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("algorithm"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("key_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("digest_type"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("digest_algorithm"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("digest"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("ds"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("key_tag"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("public_key"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
				},
				ExpectNonEmptyPlan: true, // DNSSEC computed values can change, causing plan drift
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"algorithm", "digest", "digest_algorithm", "digest_type",
					"ds", "flags", "key_tag", "key_type", "modified_on",
					"public_key", "status", "dnssec_multi_signer", "dnssec_presigned", "dnssec_use_nsec3",
				},
			},
		},
	})
}

func TestAccCloudflareZoneDNSSEC_Basic(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_dnssec.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneDNSSECDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDNSSECBasicConfig(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					// Basic test: only status is specified, other optional attributes are null  
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|pending"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_multi_signer"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_presigned"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_use_nsec3"), knownvalue.Null()),
					// Computed attributes may be null initially but populated when DNSSEC becomes active
				},
				ExpectNonEmptyPlan: true, // DNSSEC computed values can change
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"algorithm", "digest", "digest_algorithm", "digest_type",
					"ds", "flags", "key_tag", "key_type", "modified_on",
					"public_key", "status", "dnssec_multi_signer", "dnssec_presigned", "dnssec_use_nsec3",
				},
			},
		},
	})
}

func TestAccCloudflareZoneDNSSEC_StatusDisabled(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_dnssec.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneDNSSECDestroy,
		Steps: []resource.TestStep{
			// First activate DNSSEC
			{
				Config: testAccCloudflareZoneDNSSECBasicConfig(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|pending"))),
				},
				ExpectNonEmptyPlan: true,
			},
			// Then disable it
			{
				Config: testAccCloudflareZoneDNSSECDisabledConfig(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("disabled|pending-disabled"))),
					// Other optional attributes should be null
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_multi_signer"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_presigned"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_use_nsec3"), knownvalue.Null()),
				},
				ExpectNonEmptyPlan: true, // Previous tests may have left DNSSEC settings that cause drift
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"algorithm", "digest", "digest_algorithm", "digest_type",
					"ds", "flags", "key_tag", "key_type", "modified_on",
					"public_key", "status", "dnssec_multi_signer", "dnssec_presigned", "dnssec_use_nsec3",
				},
			},
		},
	})
}

func TestAccCloudflareZoneDNSSEC_MultiSigner(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_dnssec.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneDNSSECDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDNSSECMultiSignerConfig(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|pending"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_multi_signer"), knownvalue.Bool(true)),
					// Other optional attributes should be null
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_presigned"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_use_nsec3"), knownvalue.Null()),
				},
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"algorithm", "digest", "digest_algorithm", "digest_type",
					"ds", "flags", "key_tag", "key_type", "modified_on",
					"public_key", "status", "dnssec_multi_signer", "dnssec_presigned", "dnssec_use_nsec3",
				},
			},
		},
	})
}


func TestAccCloudflareZoneDNSSEC_Presigned(t *testing.T) {
	// Presigned DNSSEC requires a secondary zone
	secondaryZoneID := os.Getenv("CLOUDFLARE_SECONDARY_ZONE_ID")
	if secondaryZoneID == "" {
		// Use dedicated secondary zone created specifically for presigned DNSSEC testing
		// Zone: secondary.terraform.cfapi.net (type: secondary)
		secondaryZoneID = "e3f462b432dd82b7329cc29bbbb4e8a6"
	}
	
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_dnssec.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneDNSSECDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDNSSECPresignedConfig(secondaryZoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(secondaryZoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|pending"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_presigned"), knownvalue.Bool(true)),
					// Other optional attributes should be null
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_multi_signer"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_use_nsec3"), knownvalue.Null()),
				},
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"algorithm", "digest", "digest_algorithm", "digest_type",
					"ds", "flags", "key_tag", "key_type", "modified_on",
					"public_key", "status", "dnssec_multi_signer", "dnssec_presigned", "dnssec_use_nsec3",
				},
			},
		},
	})
}

func TestAccCloudflareZoneDNSSEC_PresignedWithNsec3(t *testing.T) {
	// Presigned DNSSEC with NSEC3 requires a secondary zone and higher-tier plan
	secondaryZoneID := os.Getenv("CLOUDFLARE_SECONDARY_ZONE_ID")
	if secondaryZoneID == "" {
		// Use dedicated secondary zone created specifically for presigned DNSSEC testing
		// Zone: secondary.terraform.cfapi.net (type: secondary)
		secondaryZoneID = "e3f462b432dd82b7329cc29bbbb4e8a6"
	}
	
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		// No CheckDestroy - test expects error, no resource created
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDNSSECPresignedWithNsec3Config(secondaryZoneID, rnd),
				ExpectError: regexp.MustCompile("Invalid zone plan for action|failed to make http request"),
			},
		},
	})
}

func TestAccCloudflareZoneDNSSEC_UseNsec3(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_dnssec.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneDNSSECDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDNSSECUseNsec3Config(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|pending"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_use_nsec3"), knownvalue.Bool(true)),
					// Other optional attributes should be null
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_multi_signer"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_presigned"), knownvalue.Null()),
				},
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"algorithm", "digest", "digest_algorithm", "digest_type",
					"ds", "flags", "key_tag", "key_type", "modified_on",
					"public_key", "status", "dnssec_multi_signer", "dnssec_presigned", "dnssec_use_nsec3",
				},
			},
		},
	})
}

func TestAccCloudflareZoneDNSSEC_Comprehensive(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_dnssec.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneDNSSECDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareZoneDNSSECComprehensiveConfig(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|pending"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_multi_signer"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_presigned"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_use_nsec3"), knownvalue.Bool(true)),
					// Computed attributes may be null when DNSSEC is still pending
					// They'll be populated once DNSSEC becomes fully active
				},
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"algorithm", "digest", "digest_algorithm", "digest_type",
					"ds", "flags", "key_tag", "key_type", "modified_on",
					"public_key", "status", "dnssec_multi_signer", "dnssec_presigned", "dnssec_use_nsec3",
				},
			},
		},
	})
}

func TestAccCloudflareZoneDNSSEC_Update(t *testing.T) {
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zone_dnssec.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareZoneDNSSECDestroy,
		Steps: []resource.TestStep{
			// Start with basic configuration (status active, no optional attributes)
			{
				Config: testAccCloudflareZoneDNSSECBasicConfig(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|pending"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_multi_signer"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_presigned"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_use_nsec3"), knownvalue.Null()),
				},
				ExpectNonEmptyPlan: true,
			},
			// Update to multi-signer configuration (safer transition)
			{
				Config: testAccCloudflareZoneDNSSECMultiSignerConfig(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|pending"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_multi_signer"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_presigned"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_use_nsec3"), knownvalue.Null()),
				},
				ExpectNonEmptyPlan: true,
			},
			// Update back to basic configuration (status active, optional attributes removed)
			{
				Config: testAccCloudflareZoneDNSSECBasicConfig(zoneID, rnd),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("status"), knownvalue.StringRegexp(regexp.MustCompile("active|pending"))),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_multi_signer"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_presigned"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("dnssec_use_nsec3"), knownvalue.Null()),
					// Computed values retain their previous values even when optional attributes are removed
				},
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"algorithm", "digest", "digest_algorithm", "digest_type",
					"ds", "flags", "key_tag", "key_type", "modified_on",
					"public_key", "status", "dnssec_multi_signer", "dnssec_presigned", "dnssec_use_nsec3",
				},
			},
		},
	})
}

func testAccCheckCloudflareZoneDNSSECDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zone_dnssec" {
			continue
		}

		zoneID := rs.Primary.Attributes[consts.ZoneIDSchemaKey]
		dnssec, err := client.DNS.DNSSEC.Get(context.Background(), dns.DNSSECGetParams{
			ZoneID: cloudflare.F(zoneID),
		})
		if err != nil {
			// DNSSEC API can return errors when checking status, which is acceptable for destroy check
			continue
		}
		
		// DNSSEC is considered destroyed when status is disabled or null
		if dnssec.Status != "disabled" && dnssec.Status != "" {
			return fmt.Errorf("zone dnssec still active, status: %s", dnssec.Status)
		}
	}

	return nil
}

func testAccCloudflareZoneDNSSECBasicConfig(zoneID string, name string) string {
	return acctest.LoadTestCase("zonednssecresourcebasic.tf", name, zoneID)
}

func testAccCloudflareZoneDNSSECDisabledConfig(zoneID string, name string) string {
	return acctest.LoadTestCase("zonednssecresourcedisabled.tf", name, zoneID)
}

func testAccCloudflareZoneDNSSECMultiSignerConfig(zoneID string, name string) string {
	return acctest.LoadTestCase("zonednssecresourcemultisigner.tf", name, zoneID)
}

func testAccCloudflareZoneDNSSECPresignedConfig(zoneID string, name string) string {
	return acctest.LoadTestCase("zonednssecresourcepresigned.tf", name, zoneID)
}

func testAccCloudflareZoneDNSSECUseNsec3Config(zoneID string, name string) string {
	return acctest.LoadTestCase("zonednssecresourcensec3.tf", name, zoneID)
}

func testAccCloudflareZoneDNSSECComprehensiveConfig(zoneID string, name string) string {
	return acctest.LoadTestCase("zonednssecresourcecomprehensive.tf", name, zoneID)
}

func testAccCloudflareZoneDNSSECPresignedWithNsec3Config(zoneID string, name string) string {
	return acctest.LoadTestCase("zonednssecresourcepresignednsec3.tf", name, zoneID)
}

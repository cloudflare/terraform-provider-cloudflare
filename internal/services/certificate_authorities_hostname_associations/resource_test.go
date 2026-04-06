package certificate_authorities_hostname_associations_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareCertificateAuthoritiesHostnameAssociations_FullLifecycle(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_certificate_authorities_hostname_associations.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")

	hostname := fmt.Sprintf("%s.%s", rnd, domain)
	hostnameUpdated := fmt.Sprintf("%s-updated.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCloudflareCertificateAuthoritiesHostnameAssociationsConfig(rnd, zoneID, hostname),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hostnames"), knownvalue.NotNull()),
				},
			},
			// Re-apply same config to verify no drift
			{
				Config: testAccCloudflareCertificateAuthoritiesHostnameAssociationsConfig(rnd, zoneID, hostname),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Update hostnames
			{
				Config: testAccCloudflareCertificateAuthoritiesHostnameAssociationsConfig(rnd, zoneID, hostnameUpdated),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hostnames"), knownvalue.ListExact(
						[]knownvalue.Check{knownvalue.StringExact(hostnameUpdated)},
					)),
				},
			},
			// Import
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"mtls_certificate_id"},
			},
		},
	})
}

func testAccCloudflareCertificateAuthoritiesHostnameAssociationsConfig(rnd, zoneID, hostname string) string {
	return acctest.LoadTestCase("certificateauthoritieshostnameassociationslifecycle.tf", rnd, zoneID, hostname)
}

func TestAccCloudflareCertificateAuthoritiesHostnameAssociations_WithMTLSCertificate(t *testing.T) {
	mtlsCertID := os.Getenv("CLOUDFLARE_MTLS_CERTIFICATE_ID")
	if mtlsCertID == "" {
		t.Skip("CLOUDFLARE_MTLS_CERTIFICATE_ID not set, skipping mTLS certificate test")
	}

	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_certificate_authorities_hostname_associations.%s", rnd)
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	hostname := fmt.Sprintf("%s.%s", rnd, domain)
	hostnameUpdated := fmt.Sprintf("%s-updated.%s", rnd, domain)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCloudflareCertificateAuthoritiesHostnameAssociationsWithMTLSConfig(rnd, zoneID, hostname, mtlsCertID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hostnames"), knownvalue.NotNull()),
				},
			},
			// Re-apply same config to verify no drift
			{
				Config: testAccCloudflareCertificateAuthoritiesHostnameAssociationsWithMTLSConfig(rnd, zoneID, hostname, mtlsCertID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			// Update hostnames
			{
				Config: testAccCloudflareCertificateAuthoritiesHostnameAssociationsWithMTLSConfig(rnd, zoneID, hostnameUpdated, mtlsCertID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("zone_id"), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("hostnames"), knownvalue.ListExact(
						[]knownvalue.Check{knownvalue.StringExact(hostnameUpdated)},
					)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("mtls_certificate_id"), knownvalue.StringExact(mtlsCertID)),
				},
			},
			// Import with <zone_id>/<mtls_certificate_id> format
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateId:     fmt.Sprintf("%s/%s", zoneID, mtlsCertID),
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudflareCertificateAuthoritiesHostnameAssociationsWithMTLSConfig(rnd, zoneID, hostname, mtlsCertID string) string {
	return acctest.LoadTestCase("certificateauthoritieshostnameassociationswithmtls.tf", rnd, zoneID, hostname, mtlsCertID)
}

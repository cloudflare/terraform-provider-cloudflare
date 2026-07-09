// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_certificate_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cfv7 "github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustGatewayCertificateModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_gateway_certificate.ZeroTrustGatewayCertificateModel)(nil)
	schema := zero_trust_gateway_certificate.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

func testAccCloudflareTeamsGatewayCert(rnd, accountID string) string {
	return acctest.LoadTestCase("teamscertificateconfigbasic.tf", rnd, accountID)
}

func testAccCloudflareTeamsGatewayCertActivate(rnd, accountID string) string {
	return acctest.LoadTestCase("teamscertificateconfigbasic-activate.tf", rnd, accountID)
}

func testAccCloudflareTeamsGatewayCertDeactivate(rnd, accountID string) string {
	return acctest.LoadTestCase("teamscertificateconfigbasic-deactivate.tf", rnd, accountID)
}

func testAccCheckCloudflareTeamsGatewayCertDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_zero_trust_gateway_certificate" {
			continue
		}

		accountId := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		_, err := client.ZeroTrust.Gateway.Certificates.Get(context.Background(), rs.Primary.ID, zero_trust.GatewayCertificateGetParams{
			AccountID: cfv7.F(accountId),
		})
		if err == nil {
			return fmt.Errorf("Teams cert still exists")
		}
	}

	return nil
}

func TestAccCloudflareTeamsCertificateLifeCycle(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Access
	// service does not yet support the API tokens and it results in
	// misleading state error messages.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_gateway_certificate.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsGatewayCertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsGatewayCert(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "issuer_org", "Cloudflare, Inc."),
					resource.TestCheckResourceAttr(resourceName, "in_use", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "activate"),
				),
			},
			{
				Config: testAccCloudflareTeamsGatewayCertActivate(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "issuer_org", "Cloudflare, Inc."),
					resource.TestCheckResourceAttr(resourceName, "in_use", "false"),
					resource.TestCheckResourceAttr(resourceName, "activate", "true"),
				),
			},
			{
				Config: testAccCloudflareTeamsGatewayCertDeactivate(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "issuer_org", "Cloudflare, Inc."),
					resource.TestCheckResourceAttr(resourceName, "in_use", "false"),
					resource.TestCheckResourceAttr(resourceName, "activate", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareTeamsCertificate_CreateWithActivate(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_gateway_certificate.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsGatewayCertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsGatewayCertActivate(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "issuer_org", "Cloudflare, Inc."),
					resource.TestCheckResourceAttr(resourceName, "in_use", "false"),
					resource.TestCheckResourceAttr(resourceName, "activate", "true"),
				),
			},
			{
				Config: testAccCloudflareTeamsGatewayCertDeactivate(rnd, accountID),
			},
		},
	})
}

func TestAccCloudflareTeamsCertificate_CreateWithDeactivate(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := fmt.Sprintf("cloudflare_zero_trust_gateway_certificate.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTeamsGatewayCertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTeamsGatewayCertDeactivate(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "issuer_org", "Cloudflare, Inc."),
					resource.TestCheckResourceAttr(resourceName, "in_use", "false"),
					resource.TestCheckResourceAttr(resourceName, "activate", "false"),
				),
			},
		},
	})
}

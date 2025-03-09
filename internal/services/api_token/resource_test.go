package api_token_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAPIToken_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceID := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithoutCondition(rnd, rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", rnd),
				),
			},
			{
				Config: testAccCloudflareAPITokenWithoutCondition(rnd, rnd+"-updated", permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceID, "name", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccAPIToken_AllowDeny(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAPITokenConfigAllowDeny(rnd, permissionID, zoneID, false),
			},
			{
				Config: testAPITokenConfigAllowDeny(rnd, permissionID, zoneID, true),
			},
			{
				Config: testAPITokenConfigAllowDeny(rnd, permissionID, zoneID, false),
			},
		},
	})
}

func TestAccAPIToken_DoesNotSetConditions(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithoutCondition(rnd, rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckNoResourceAttr(name, "condition.0.request_ip.0.in"),
					resource.TestCheckNoResourceAttr(name, "condition.0.request_ip.0.not_in"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithoutCondition(resourceName, rnd, permissionID string) string {
	return acctest.LoadTestCase("apitokenwithoutcondition.tf", resourceName, rnd, permissionID)
}

func TestAccAPIToken_SetIndividualCondition(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending service fix to address nested object syntax as strings for `conditions`.")

	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithIndividualCondition(rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "condition.0.request_ip.0.in.0", "192.0.2.1/32"),
					resource.TestCheckNoResourceAttr(name, "condition.0.request_ip.0.not_in"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithIndividualCondition(rnd string, permissionID string) string {
	return acctest.LoadTestCase("apitokenwithindividualcondition.tf", rnd, permissionID)
}

func TestAccAPIToken_SetAllCondition(t *testing.T) {
	acctest.TestAccSkipForDefaultZone(t, "Pending service fix to address nested object syntax as strings for `conditions`.")

	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithAllCondition(rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "condition.0.request_ip.0.in.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr(name, "condition.0.request_ip.0.not_in.0", "198.51.100.1/32"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithAllCondition(rnd string, permissionID string) string {
	return acctest.LoadTestCase("apitokenwithallcondition.tf", rnd, permissionID)
}

func testAPITokenConfigAllowDeny(resourceID, permissionID, zoneID string, allowAllZonesExceptOne bool) string {
	var add string
	if allowAllZonesExceptOne {
		add = acctest.LoadTestCase("apitokenconfigallowdeny.tf", permissionID, zoneID)
	}

	return acctest.LoadTestCase("apitokenconfigallowdeny.tf", resourceID, permissionID, add)
}

func TestAccAPIToken_TokenTTL(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the API token
	// endpoint does not yet support the API tokens without an explicit scope.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_api_token." + rnd
	permissionID := "82e64a83756745bbbb1c9c2701bf816b" // DNS read

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAPITokenWithTTL(rnd, permissionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "not_before", "2018-07-01T05:20:00Z"),
					resource.TestCheckResourceAttr(name, "expires_on", "2032-01-01T00:00:00Z"),
				),
			},
		},
	})
}

func testAccCloudflareAPITokenWithTTL(rnd string, permissionID string) string {
	return acctest.LoadTestCase("apitokenwithttl.tf", rnd, permissionID)
}

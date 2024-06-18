package access_identity_provider_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

func TestAccCloudflareAccessIdentityProviderDataSource_PreventZoneIdAndAccountIdConflicts(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testCloudflareAccessIdentityProviderDataSourceConfigConflictingFields(rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("only one of `account_id,zone_id` can be specified")),
			},
		},
	})
}

func testCloudflareAccessIdentityProviderDataSourceConfigConflictingFields(rnd string) string {
	return fmt.Sprintf(`
data "cloudflare_access_identity_provider" "%[1]s" {
  account_id = "123abc"
  zone_id    = "abc123"
  name       = "foo"
}
`, rnd)
}

func TestAccCloudflareAccessIdentityProviderDataSource_PreventNoInputSpecify(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testCloudflareAccessIdentityProviderDataSourceNoInput(rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("one of `account_id,zone_id` must be specified")),
			},
		},
	})
}

func testCloudflareAccessIdentityProviderDataSourceNoInput(rnd string) string {
	return fmt.Sprintf(`
data "cloudflare_access_identity_provider" "%[1]s" {
	name = "foo"
}
`, rnd)
}

func TestAccCloudflareAccessIdentityProviderDataSourceNotFound(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareAccessIdentityProviderDataSource_NotFound(accountID, rnd),
				ExpectError: regexp.MustCompile(`no Access Identity Providers found|no Access Identity Provider matching name`),
			},
		},
	})
}

func testAccCheckCloudflareAccessIdentityProviderDataSource_NotFound(accountID, name string) string {
	return fmt.Sprintf(`
data "cloudflare_access_identity_provider" "%[1]s" {
	account_id = "%[2]s"
	name = "%[1]s-abc123"
}

resource "cloudflare_access_identity_provider" "%[1]s" {
	account_id = "%[2]s"
  name = "%[1]s"
  type = "github"
  config {
    client_id = "test"
    client_secret = "secret"
	}
}
	`, name, accountID)
}

func TestAccCloudflareAccessIdentityProviderDataSource_GitHub(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()

	name := "data.cloudflare_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderDataSourceGitHub(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "github"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessIdentityProviderDataSourceGitHub(accountID, name string) string {
	return fmt.Sprintf(`
	resource "cloudflare_access_identity_provider" "%[1]s" {
		account_id = "%[2]s"
	  name = "%[1]s"
	  type = "github"
	  config {
		client_id = "test"
		client_secret = "secret"
		}
	}

	data "cloudflare_access_identity_provider" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		depends_on = [cloudflare_access_identity_provider.%[1]s]
	}
	`, name, accountID)
}

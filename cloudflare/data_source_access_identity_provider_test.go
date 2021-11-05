package cloudflare

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudflareAccessIdentityProviderDataSource_PreventZoneIdAndAccountIdConflicts(t *testing.T) {
	t.Parallel()
	rnd := generateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	t.Parallel()
	rnd := generateRandomResourceName()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareAccessIdentityProviderDataSource_NotFound(accountID, rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("no Access Identity Provider matching name")),
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

func TestAccCloudflareAccessIdentityProviderDataSourceGitHub(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := generateRandomResourceName()

	name := "data.cloudflare_access_identity_provider." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccessAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessIdentityProviderDataSourceGitHub(accountID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "type", "github"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessIdentityProviderDataSourceGitHub(accountID, name string) string {
	return fmt.Sprintf(`
data "cloudflare_access_identity_provider" "%[1]s" {
	account_id = "%[2]s"
	name = "%[1]s"
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

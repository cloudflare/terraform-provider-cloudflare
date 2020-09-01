package cloudflare

import (
	"fmt"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	accountID   = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	email       = "test@example.com"
	accessGroup cloudflare.AccessGroup
)

func TestAccCloudflareAccessGroupConfig_Basic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &accessGroup),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Exclude(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessGroupConfigExclude(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &accessGroup),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "exclude.0.email.0", email),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Require(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessGroupConfigRequire(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &accessGroup),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "require.0.email.0", email),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_FullConfig(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessGroupConfigFullConfig(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &accessGroup),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "exclude.0.email.0", email),
					resource.TestCheckResourceAttr(name, "require.0.email.0", email),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroupWithIDP(t *testing.T) {
	rnd := generateRandomResourceName()
	groupName := fmt.Sprintf("cloudflare_access_group.%s", rnd)
	githubOrg := "Terraform-Cloudflare-Provider-Test-Org"
	team := "test-team-1"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupWithIDP(accountID, rnd, githubOrg, team),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(groupName, &accessGroup),
					resource.TestCheckResourceAttr(groupName, "account_id", accountID),
					resource.TestCheckResourceAttr(groupName, "name", rnd),
					resource.TestCheckResourceAttrSet(groupName, "include.0.github.0.identity_provider_id"),
					resource.TestCheckResourceAttr(groupName, "include.0.github.0.name", githubOrg),
					resource.TestCheckResourceAttr(groupName, "include.0.github.0.teams.0", team),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Updated(t *testing.T) {
	var before, after cloudflare.AccessGroup
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &before),
				),
			},
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, accountID, "test-changed@example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &after),
					testAccCheckCloudflareAccessGroupIDUnchanged(&before, &after),
					resource.TestCheckResourceAttr(name, "include.0.email.0", "test-changed@example.com"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_CreateAfterManualDestroy(t *testing.T) {
	var before, after cloudflare.AccessGroup
	var initialID string
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &before),
					testAccManuallyDeleteAccessGroup(name, &initialID),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCloudflareAccessGroupConfigBasicWithUpdate(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &after),
					testAccCheckCloudflareAccessGroupRecreated(&before, &after),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s-updated", rnd)),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
				),
			},
		},
	})
}

func testAccCloudflareAccessGroupConfigBasic(resourceName, accountID, email string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  include {
    email = ["%[3]s"]
  }
}`, resourceName, accountID, email)
}

func testAccCloudflareAccessGroupConfigBasicWithUpdate(resourceName, accountID, email string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s-updated"

  include {
    email = ["%[3]s"]
  }
}`, resourceName, accountID, email)
}

func testAccessGroupConfigExclude(resourceName, accountID, email string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  include {
    email = ["%[3]s"]
  }

  exclude {
    email = ["%[3]s"]
  }
}`, resourceName, accountID, email)
}

func testAccessGroupConfigRequire(resourceName, accountID, email string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  include {
    email = ["%[3]s"]
  }

  require {
    email = ["%[3]s"]
  }
}`, resourceName, accountID, email)
}

func testAccessGroupConfigFullConfig(resourceName, accountID, email string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  include {
    email = ["%[3]s"]
  }

  require {
    email = ["%[3]s"]
  }

  exclude {
    email = ["%[3]s"]
  }
}`, resourceName, accountID, email)
}

func testAccCloudflareAccessGroupWithIDP(accountID, rnd, githubOrg, team string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"
  type = "github"
  config {
    client_id = "test"
    client_secret = "secret"
  }
}

resource "cloudflare_access_group" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"

  include {
    github {
      name                 = "%[3]s"
      teams                = ["%[4]s"]
      identity_provider_id = cloudflare_access_identity_provider.%[2]s.id
    }
  }
}`, accountID, rnd, githubOrg, team)
}

func testAccCheckCloudflareAccessGroupExists(n string, accessGroup *cloudflare.AccessGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AccessGroup ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundAccessGroup, err := client.AccessGroup(rs.Primary.Attributes["account_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundAccessGroup.ID != rs.Primary.ID {
			return fmt.Errorf("AccessGroup not found")
		}

		*accessGroup = foundAccessGroup

		return nil
	}
}

func testAccCheckCloudflareAccessGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudflare.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_access_group" {
			continue
		}

		_, err := client.AccessGroup(rs.Primary.Attributes["account_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("AccessGroup still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareAccessGroupIDUnchanged(before, after *cloudflare.AccessGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID != after.ID {
			return fmt.Errorf("ID should not change suring in place update, but got change %s -> %s", before.ID, after.ID)
		}
		return nil
	}
}

func testAccManuallyDeleteAccessGroup(name string, initialID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		*initialID = rs.Primary.ID
		err := client.DeleteAccessGroup(rs.Primary.Attributes["account_id"], rs.Primary.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareAccessGroupRecreated(before, after *cloudflare.AccessGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("Expected change of AccessGroup Ids, but both were %v", before.ID)
		}
		return nil
	}
}

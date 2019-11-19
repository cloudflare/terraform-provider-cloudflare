package cloudflare

import (
	"fmt"
	"os"
	"strings"
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
	resourceName := strings.Split(name, ".")[1]
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(resourceName, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &accessGroup),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Exclude(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)
	resourceName := strings.Split(name, ".")[1]
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessGroupConfigExclude(resourceName, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &accessGroup),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
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
	resourceName := strings.Split(name, ".")[1]
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessGroupConfigRequire(resourceName, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &accessGroup),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
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
	resourceName := strings.Split(name, ".")[1]
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessGroupConfigFullConfig(resourceName, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, &accessGroup),
					resource.TestCheckResourceAttr(name, "account_id", accountID),
					resource.TestCheckResourceAttr(name, "name", resourceName),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "exclude.0.email.0", email),
					resource.TestCheckResourceAttr(name, "require.0.email.0", email),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Updated(t *testing.T) {
	var before, after cloudflare.AccessGroup
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(resourceName, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists("cloudflare_access_group.test", &before),
				),
			},
			{
				Config: testAccCheckCloudflareAccessGroupConfigNewValue(name, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists("cloudflare_access_group.test", &after),
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
	resourceName := strings.Split(name, ".")[1]

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(resourceName, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists("cloudflare_access_group.test", &before),
					testAccManuallyDeleteAccessGroup("cloudflare_access_group.test", &initialID),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflareAccessGroupConfigNewValue(name, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists("cloudflare_access_group.test", &after),
					testAccCheckCloudflareAccessGroupRecreated(&before, &after),
					resource.TestCheckResourceAttr(
						"cloudflare_access_group.test", "account_id", accountID),
					resource.TestCheckResourceAttr(
						"cloudflare_access_group.test", "name", fmt.Sprintf("%s/updated", name)),
					resource.TestCheckResourceAttr(
						"cloudflare_access_group.test", "include.0.email.0", email),
					resource.TestCheckResourceAttr(
						"cloudflare_access_group.test", "exclude.0.email.0", email),
					resource.TestCheckResourceAttr(
						"cloudflare_access_group.test", "require.0.email.0", email),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_UpdatingAccountIDForcesNewResource(t *testing.T) {
	var before, after cloudflare.AccessGroup
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)
	resourceName := strings.Split(name, ".")[1]
	newAccountID := "01a7362d577a6c3019a474fd6f485634"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(resourceName, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists("cloudflare_access_group.test", &before),
					resource.TestCheckResourceAttr("cloudflare_access_group.test", "account_id", accountID),
				),
			},
			{
				Config: testAccCloudflareAccessGroupConfigBasic(resourceName, newAccountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists("cloudflare_page_rule.test", &after),
					testAccCheckCloudflareAccessGroupRecreated(&before, &after),
					resource.TestCheckResourceAttr("cloudflare_page_rule.test", "account_id", newAccountID),
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

func testAccCheckCloudflareAccessGroupConfigNewValue(resourceName, accountID, email string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_group" "test" {
  account_id = "%[2]s"
  name = "%[1]s"
  include {
    email = ["%[3]s"]
  }
}`, resourceName, accountID, email)
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

package access_group_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	resource.AddTestSweepers("cloudflare_access_group", &resource.Sweeper{
		Name: "cloudflare_access_group",
		F:    testSweepCloudflareAccessGroups,
	})
}

func testSweepCloudflareAccessGroups(r string) error {
	ctx := context.Background()

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to create Cloudflare client: %s", clientErr))
	}

	// Zone level Access Groups
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	zoneAccessGroups, _, err := client.ListAccessGroups(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListAccessGroupsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch zone level Access Groups: %s", err))
	}

	if len(zoneAccessGroups) == 0 {
		log.Print("[DEBUG] No Cloudflare zone level Access Groups to sweep")
		return nil
	}

	for _, accessGroup := range zoneAccessGroups {
		if err := client.DeleteAccessGroup(context.Background(), cloudflare.ZoneIdentifier(zoneID), accessGroup.ID); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete zone level Access Group %s", accessGroup.ID))
		}
	}

	// Account level Access Groups
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accountAccessGroups, _, err := client.ListAccessGroups(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListAccessGroupsParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch account level Access Groups: %s", err))
	}

	if len(accountAccessGroups) == 0 {
		log.Print("[DEBUG] No Cloudflare account level Access Groups to sweep")
		return nil
	}

	for _, accessGroup := range accountAccessGroups {
		if err := client.DeleteAccessGroup(context.Background(), cloudflare.AccountIdentifier(accountID), accessGroup.ID); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete account level Access Group %s", accessGroup.ID))
		}
	}

	return nil
}

var (
	zoneID      = os.Getenv("CLOUDFLARE_ZONE_ID")
	accountID   = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	email       = "test@example.com"
	accessGroup cloudflare.AccessGroup
)

func TestAccCloudflareAccessGroup_ConfigBasicZone(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, email, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "include.0.email_domain.0", "example.com"),
					resource.TestCheckResourceAttr(name, "include.0.any_valid_service_token", "true"),
					resource.TestCheckResourceAttr(name, "include.0.ip.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr(name, "include.0.ip.1", "192.0.2.2/32"),
					resource.TestCheckResourceAttr(name, "include.0.ip_list.0", "e3a0f205-c525-4e48-a293-ba5d1f00e638"),
					resource.TestCheckResourceAttr(name, "include.0.ip_list.1", "5d54cd30-ce52-46e4-9a46-a47887e1a167"),
				),
			},
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, email, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "include.0.email_domain.0", "example.com"),
					resource.TestCheckResourceAttr(name, "include.0.any_valid_service_token", "true"),
					resource.TestCheckResourceAttr(name, "include.0.ip.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr(name, "include.0.ip.1", "192.0.2.2/32"),
					resource.TestCheckResourceAttr(name, "include.0.ip_list.0", "e3a0f205-c525-4e48-a293-ba5d1f00e638"),
					resource.TestCheckResourceAttr(name, "include.0.ip_list.1", "5d54cd30-ce52-46e4-9a46-a47887e1a167"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_ConfigBasicAccount(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, email, cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.ZoneIdentifier(zoneID), &accessGroup),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "include.0.email_domain.0", "example.com"),
					resource.TestCheckResourceAttr(name, "include.0.any_valid_service_token", "true"),
					resource.TestCheckResourceAttr(name, "include.0.ip.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr(name, "include.0.ip.1", "192.0.2.2/32"),
					resource.TestCheckResourceAttr(name, "include.0.ip_list.0", "e3a0f205-c525-4e48-a293-ba5d1f00e638"),
					resource.TestCheckResourceAttr(name, "include.0.ip_list.1", "5d54cd30-ce52-46e4-9a46-a47887e1a167"),
					resource.TestCheckResourceAttr(name, "include.0.saml.0.attribute_name", "Name1"),
					resource.TestCheckResourceAttr(name, "include.0.saml.0.attribute_value", "Value1"),
					resource.TestCheckResourceAttr(name, "include.0.saml.1.attribute_name", "Name2"),
					resource.TestCheckResourceAttr(name, "include.0.saml.1.attribute_value", "Value2"),
					resource.TestCheckResourceAttr(name, "include.0.azure.0.id.0", "group1"),
					resource.TestCheckResourceAttr(name, "include.0.azure.0.identity_provider_id", "1234"),
					resource.TestCheckResourceAttr(name, "include.0.azure.1.id.0", "group2"),
					resource.TestCheckResourceAttr(name, "include.0.azure.1.identity_provider_id", "5678"),
				),
			},
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, email, cloudflare.ZoneIdentifier(zoneID)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.ZoneIdentifier(zoneID), &accessGroup),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "include.0.email_domain.0", "example.com"),
					resource.TestCheckResourceAttr(name, "include.0.any_valid_service_token", "true"),
					resource.TestCheckResourceAttr(name, "include.0.ip.0", "192.0.2.1/32"),
					resource.TestCheckResourceAttr(name, "include.0.ip.1", "192.0.2.2/32"),
					resource.TestCheckResourceAttr(name, "include.0.ip_list.0", "e3a0f205-c525-4e48-a293-ba5d1f00e638"),
					resource.TestCheckResourceAttr(name, "include.0.ip_list.1", "5d54cd30-ce52-46e4-9a46-a47887e1a167"),
					resource.TestCheckResourceAttr(name, "include.0.saml.0.attribute_name", "Name1"),
					resource.TestCheckResourceAttr(name, "include.0.saml.0.attribute_value", "Value1"),
					resource.TestCheckResourceAttr(name, "include.0.saml.1.attribute_name", "Name2"),
					resource.TestCheckResourceAttr(name, "include.0.saml.1.attribute_value", "Value2"),
					resource.TestCheckResourceAttr(name, "include.0.azure.0.id.0", "group1"),
					resource.TestCheckResourceAttr(name, "include.0.azure.0.identity_provider_id", "1234"),
					resource.TestCheckResourceAttr(name, "include.0.azure.1.id.0", "group2"),
					resource.TestCheckResourceAttr(name, "include.0.azure.1.identity_provider_id", "5678"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_ConfigEmailList(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	rnd2 := utils.GenerateRandomResourceName()
	emailListName := fmt.Sprintf("cloudflare_teams_list.%s", rnd2)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigEmailList(rnd, rnd2, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttrSet(name, "include.0.email_list.0"),

					// Check that the email list is destroyed
					resource.TestCheckResourceAttr(emailListName, "name", rnd2),
					resource.TestCheckResourceAttr(emailListName, "type", "EMAIL"),
					resource.TestCheckResourceAttr(emailListName, "items.0", "test@example.com"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Exclude(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessGroupConfigExclude(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "include.0.email_domain.0", "example.com"),
					resource.TestCheckResourceAttr(name, "exclude.0.email.0", email),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Require(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessGroupConfigRequire(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "include.0.email_domain.0", "example.com"),
					resource.TestCheckResourceAttr(name, "require.0.email.0", email),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_FullConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessGroupConfigFullConfig(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "include.0.email_domain.0", "example.com"),
					resource.TestCheckResourceAttr(name, "exclude.0.email.0", email),
					resource.TestCheckResourceAttr(name, "require.0.email.0", email),
					resource.TestCheckResourceAttr(name, "include.0.common_names.0", "common"),
					resource.TestCheckResourceAttr(name, "include.0.common_names.1", "name"),
					resource.TestCheckNoResourceAttr(name, "include.0.common_name.0"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_WithIDP(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	groupName := fmt.Sprintf("cloudflare_access_group.%s", rnd)
	githubOrg := "Terraform-Cloudflare-Provider-Test-Org"
	team := "test-team-1"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupWithIDP(accountID, rnd, githubOrg, team),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(groupName, cloudflare.AccountIdentifier(accountID), &accessGroup),
					resource.TestCheckResourceAttr(groupName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(groupName, "name", rnd),
					resource.TestCheckResourceAttrSet(groupName, "include.0.github.0.identity_provider_id"),
					resource.TestCheckResourceAttr(groupName, "include.0.github.0.name", githubOrg),
					resource.TestCheckResourceAttr(groupName, "include.0.github.0.teams.0", team),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_WithIDPAuthContext(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	groupName := fmt.Sprintf("cloudflare_access_group.%s", rnd)
	ctxID := utils.GenerateRandomResourceName()
	ctxACID := "c1"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupWithIDPAuthContext(accountID, rnd, ctxID, ctxACID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(groupName, cloudflare.AccountIdentifier(accountID), &accessGroup),
					resource.TestCheckResourceAttr(groupName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(groupName, "name", rnd),
					resource.TestCheckResourceAttrSet(groupName, "require.0.auth_context.0.identity_provider_id"),
					resource.TestCheckResourceAttr(groupName, "require.0.auth_context.0.id", ctxID),
					resource.TestCheckResourceAttr(groupName, "require.0.auth_context.0.ac_id", ctxACID),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Updated(t *testing.T) {
	var before, after cloudflare.AccessGroup
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, email, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &before),
				),
			},
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, "test-changed@example.com", cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &after),
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
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, email, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &before),
					testAccManuallyDeleteAccessGroup(name, &initialID),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCloudflareAccessGroupConfigBasicWithUpdate(rnd, accountID, email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &after),
					testAccCheckCloudflareAccessGroupRecreated(&before, &after),
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", fmt.Sprintf("%s-updated", rnd)),
					resource.TestCheckResourceAttr(name, "include.0.email.0", email),
					resource.TestCheckResourceAttr(name, "include.0.email_domain.0", "example.com"),
				),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_UpdatedFromCommonNameToCommonNames(t *testing.T) {
	var before, after cloudflare.AccessGroup
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasicWithCommonName(rnd, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &before),
				),
			},
			{
				Config: testAccCloudflareAccessGroupConfigBasicWithCommonNames(rnd, cloudflare.AccountIdentifier(accountID)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &after),
					testAccCheckCloudflareAccessGroupIDUnchanged(&before, &after),
					resource.TestCheckResourceAttr(name, "include.0.common_names.0", "common"),
					resource.TestCheckResourceAttr(name, "include.0.common_names.1", "name"),
					resource.TestCheckNoResourceAttr(name, "include.0.common_name.0"),
				),
			},
		},
	})
}

func testAccCloudflareAccessGroupConfigBasic(resourceName string, email string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accessgroupconfigbasic.tf", resourceName, email, identifier.Type, identifier.Identifier)
}

func testAccCloudflareAccessGroupConfigBasicWithUpdate(resourceName, accountID, email string) string {
	return acctest.LoadTestCase("accessgroupconfigbasicwithupdate.tf", resourceName, accountID, email)
}

func testAccessGroupConfigExclude(resourceName, accountID, email string) string {
	return acctest.LoadTestCase("accessgroupconfigexclude.tf", resourceName, accountID, email)
}

func testAccessGroupConfigRequire(resourceName, accountID, email string) string {
	return acctest.LoadTestCase("accessgroupconfigrequire.tf", resourceName, accountID, email)
}

func testAccessGroupConfigFullConfig(resourceName, accountID, email string) string {
	return acctest.LoadTestCase("accessgroupconfigfullconfig.tf", resourceName, accountID, email)
}

func testAccCloudflareAccessGroupWithIDP(accountID, rnd, githubOrg, team string) string {
	return acctest.LoadTestCase("accessgroupwithidp.tf", accountID, rnd, githubOrg, team)
}

func testAccCloudflareAccessGroupWithIDPAuthContext(accountID, rnd, authCtxID, authCtxACID string) string {
	return acctest.LoadTestCase("accessgroupwithidpauthcontext.tf", accountID, rnd, authCtxID, authCtxACID)
}

func testAccCloudflareAccessGroupConfigBasicWithCommonName(resourceName string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accessgroupconfigbasicwithcommonname.tf", resourceName, identifier.Type, identifier.Identifier)
}

func testAccCloudflareAccessGroupConfigBasicWithCommonNames(resourceName string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accessgroupconfigbasicwithcommonnames.tf", resourceName, identifier.Type, identifier.Identifier)
}

func testAccCloudflareAccessGroupConfigEmailList(resourceName string, emailListName string, identifier *cloudflare.ResourceContainer) string {
	return acctest.LoadTestCase("accessgroupconfigemaillist.tf", resourceName, emailListName, identifier.Type, identifier.Identifier)
}

func testAccCheckCloudflareAccessGroupExists(n string, accessIdentifier *cloudflare.ResourceContainer, accessGroup *cloudflare.AccessGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AccessGroup ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		var foundAccessGroup cloudflare.AccessGroup
		var err error

		if accessIdentifier.Level == cloudflare.AccountRouteLevel {
			foundAccessGroup, err = client.GetAccessGroup(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), rs.Primary.ID)
			if err != nil {
				return err
			}
		} else {
			foundAccessGroup, err = client.GetAccessGroup(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
			if err != nil {
				return err
			}
		}

		if foundAccessGroup.ID != rs.Primary.ID {
			return fmt.Errorf("AccessGroup not found")
		}

		*accessGroup = foundAccessGroup

		return nil
	}
}

func testAccCheckCloudflareAccessGroupDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_access_group" {
			continue
		}

		_, err := client.GetAccessGroup(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("AccessGroup still exists")
		}

		_, err = client.GetAccessGroup(context.Background(), cloudflare.ZoneIdentifier(rs.Primary.Attributes[consts.ZoneIDSchemaKey]), rs.Primary.ID)
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
			return fmt.Errorf("not found: %s", name)
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		*initialID = rs.Primary.ID
		err := client.DeleteAccessGroup(context.Background(), cloudflare.AccountIdentifier(rs.Primary.Attributes[consts.AccountIDSchemaKey]), rs.Primary.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareAccessGroupRecreated(before, after *cloudflare.AccessGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before.ID == after.ID {
			return fmt.Errorf("expected change of AccessGroup Ids, but both were %v", before.ID)
		}
		return nil
	}
}

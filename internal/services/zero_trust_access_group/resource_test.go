package zero_trust_access_group_test

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
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func init() {
	resource.AddTestSweepers("cloudflare_zero_trust_access_group", &resource.Sweeper{
		Name: "cloudflare_zero_trust_access_group",
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

func TestAccCloudflareAccessGroup_ConfigBasicAccount(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(2).AtMapKey("email_domain").AtMapKey("domain"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(3).AtMapKey("ip").AtMapKey("ip"), knownvalue.StringExact("192.0.2.1/32")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(4).AtMapKey("ip_list").AtMapKey("id"), knownvalue.StringExact("e3a0f205-c525-4e48-a293-ba5d1f00e638")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(5).AtMapKey("saml").AtMapKey("attribute_name"), knownvalue.StringExact("Name1")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(5).AtMapKey("saml").AtMapKey("attribute_value"), knownvalue.StringExact("Value1")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(5).AtMapKey("saml").AtMapKey("identity_provider_id"), knownvalue.StringExact("1234")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(6).AtMapKey("azure_ad").AtMapKey("id"), knownvalue.StringExact("group1")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(6).AtMapKey("azure_ad").AtMapKey("identity_provider_id"), knownvalue.StringExact("1234")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(7).AtMapKey("ip").AtMapKey("ip"), knownvalue.StringExact("192.0.2.2/32")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(8).AtMapKey("ip_list").AtMapKey("id"), knownvalue.StringExact("5d54cd30-ce52-46e4-9a46-a47887e1a167")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(9).AtMapKey("saml").AtMapKey("attribute_name"), knownvalue.StringExact("Name2")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(9).AtMapKey("saml").AtMapKey("attribute_value"), knownvalue.StringExact("Value2")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(9).AtMapKey("saml").AtMapKey("identity_provider_id"), knownvalue.StringExact("1234")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(10).AtMapKey("azure_ad").AtMapKey("id"), knownvalue.StringExact("group2")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(10).AtMapKey("azure_ad").AtMapKey("identity_provider_id"), knownvalue.StringExact("5678")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_ConfigBasicZone(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigBasic(rnd, email, cloudflare.ZoneIdentifier(zoneID)),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.ZoneIDSchemaKey), knownvalue.StringExact(zoneID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(2).AtMapKey("email_domain").AtMapKey("domain"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(3).AtMapKey("ip").AtMapKey("ip"), knownvalue.StringExact("192.0.2.1/32")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(4).AtMapKey("ip_list").AtMapKey("id"), knownvalue.StringExact("e3a0f205-c525-4e48-a293-ba5d1f00e638")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(5).AtMapKey("saml").AtMapKey("attribute_name"), knownvalue.StringExact("Name1")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(5).AtMapKey("saml").AtMapKey("attribute_value"), knownvalue.StringExact("Value1")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(5).AtMapKey("saml").AtMapKey("identity_provider_id"), knownvalue.StringExact("1234")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(6).AtMapKey("azure_ad").AtMapKey("id"), knownvalue.StringExact("group1")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(6).AtMapKey("azure_ad").AtMapKey("identity_provider_id"), knownvalue.StringExact("1234")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(7).AtMapKey("ip").AtMapKey("ip"), knownvalue.StringExact("192.0.2.2/32")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(8).AtMapKey("ip_list").AtMapKey("id"), knownvalue.StringExact("5d54cd30-ce52-46e4-9a46-a47887e1a167")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(9).AtMapKey("saml").AtMapKey("attribute_name"), knownvalue.StringExact("Name2")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(9).AtMapKey("saml").AtMapKey("attribute_value"), knownvalue.StringExact("Value2")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(9).AtMapKey("saml").AtMapKey("identity_provider_id"), knownvalue.StringExact("1234")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(10).AtMapKey("azure_ad").AtMapKey("id"), knownvalue.StringExact("group2")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(10).AtMapKey("azure_ad").AtMapKey("identity_provider_id"), knownvalue.StringExact("5678")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.ZoneIdentifier(zoneID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("zones/%s/", zoneID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_ConfigEmailList(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

	rnd2 := utils.GenerateRandomResourceName()
	emailListName := fmt.Sprintf("cloudflare_zero_trust_list.%s", rnd2)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigEmailList(rnd, rnd2, cloudflare.AccountIdentifier(accountID)),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(emailListName, tfjsonpath.New("name"), knownvalue.StringExact(rnd2)),
					statecheck.ExpectKnownValue(emailListName, tfjsonpath.New("type"), knownvalue.StringExact("EMAIL")),
					statecheck.ExpectKnownValue(emailListName, tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("value"), knownvalue.StringExact("test@example.com")),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Exclude(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("email_domain").AtMapKey("domain"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Require(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("email_domain").AtMapKey("domain"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_FullConfig(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("email_domain").AtMapKey("domain"), knownvalue.StringExact("example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(2).AtMapKey("common_name").AtMapKey("common_name"), knownvalue.StringExact("common")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(3).AtMapKey("common_name").AtMapKey("common_name"), knownvalue.StringExact("name")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact(email)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_WithIDP(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	groupName := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("github_organization").AtMapKey("name"), knownvalue.StringExact(githubOrg)),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("github_organization").AtMapKey("team"), knownvalue.StringExact(team)),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("exclude"), knownvalue.Null()),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("require"), knownvalue.Null()),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("is_default"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(groupName, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            groupName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_WithIDPAuthContext(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	groupName := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)
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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("require").AtSliceIndex(0).AtMapKey("auth_context").AtMapKey("id"), knownvalue.StringExact(ctxID)),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("require").AtSliceIndex(0).AtMapKey("auth_context").AtMapKey("ac_id"), knownvalue.StringExact(ctxACID)),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("exclude"), knownvalue.Null()),
					statecheck.ExpectKnownValue(groupName, tfjsonpath.New("is_default"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(groupName, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            groupName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_Updated(t *testing.T) {
	var before, after cloudflare.AccessGroup
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("test-changed@example.com")),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &after),
					testAccCheckCloudflareAccessGroupIDUnchanged(&before, &after),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_UpdatedFromCommonNameToCommonNames(t *testing.T) {
	var before, after cloudflare.AccessGroup
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

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
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("common_name").AtMapKey("common_name"), knownvalue.StringExact("common")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("common_name").AtMapKey("common_name"), knownvalue.StringExact("name")),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &after),
					testAccCheckCloudflareAccessGroupIDUnchanged(&before, &after),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
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

func testAccCloudflareAccessGroupConfigMinimal(resourceName, accountID string) string {
	return acctest.LoadTestCase("accessgroupconfigminimal.tf", resourceName, accountID)
}

func testAccCloudflareAccessGroupConfigUpdateRuleTypes(resourceName, accountID string) string {
	return acctest.LoadTestCase("accessgroupconfigupdateruletypes.tf", resourceName, accountID)
}

func testAccCloudflareAccessGroupConfigWithIsDefault(resourceName, accountID string, isDefault bool) string {
	return acctest.LoadTestCase("accessgroupconfigwithisdefault.tf", resourceName, accountID, isDefault)
}

func testAccCloudflareAccessGroupConfigComplexRules(resourceName, accountID string) string {
	return acctest.LoadTestCase("accessgroupconfigcomplexrules.tf", resourceName, accountID)
}

func testAccCloudflareAccessGroupConfigAllRuleTypes(resourceName, accountID string) string {
	return acctest.LoadTestCase("accessgroupconfigallruletypes.tf", resourceName, accountID)
}

func testAccCloudflareAccessGroupConfigServiceTokens(resourceName, accountID string) string {
	return acctest.LoadTestCase("accessgroupconfigservicetokens.tf", resourceName, accountID)
}

func testAccCloudflareAccessGroupConfigLoginMethod(resourceName, accountID string) string {
	return acctest.LoadTestCase("accessgroupconfigloginmethod.tf", resourceName, accountID)
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
		if rs.Type != "cloudflare_zero_trust_access_group" {
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

func TestAccCloudflareAccessGroup_MinimalConfiguration(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigMinimal(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("everyone"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_MultipleIncludeRuleTypes(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigUpdateRuleTypes(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("test@example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("everyone"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_IsDefaultAttribute(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigWithIsDefault(rnd, accountID, true),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default"), knownvalue.Bool(true)),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				Config: testAccCloudflareAccessGroupConfigWithIsDefault(rnd, accountID, false),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default"), knownvalue.Bool(false)),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_ComplexRuleCombinations(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigComplexRules(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("include@example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("ip").AtMapKey("ip"), knownvalue.StringExact("10.0.0.0/8")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("exclude@example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require").AtSliceIndex(0).AtMapKey("email_domain").AtMapKey("domain"), knownvalue.StringExact("company.com")),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}


func TestAccCloudflareAccessGroup_UpdateOptionalAttributes(t *testing.T) {
	var before, after cloudflare.AccessGroup
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigMinimal(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("is_default"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &before),
				),
			},
			{
				Config: testAccCloudflareAccessGroupConfigComplexRules(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("exclude@example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require").AtSliceIndex(0).AtMapKey("email_domain").AtMapKey("domain"), knownvalue.StringExact("company.com")),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &after),
					testAccCheckCloudflareAccessGroupIDUnchanged(&before, &after),
				),
			},
			{
				Config: testAccCloudflareAccessGroupConfigMinimal(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude"), knownvalue.Null()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require"), knownvalue.Null()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &after),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_AllRuleTypes(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigAllRuleTypes(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("test@example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude").AtSliceIndex(0).AtMapKey("geo").AtMapKey("country_code"), knownvalue.StringExact("CN")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude").AtSliceIndex(1).AtMapKey("device_posture").AtMapKey("integration_uid"), knownvalue.StringExact("test-device-posture-uid")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude").AtSliceIndex(2).AtMapKey("external_evaluation").AtMapKey("evaluate_url"), knownvalue.StringExact("https://example.com/evaluate")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude").AtSliceIndex(2).AtMapKey("external_evaluation").AtMapKey("keys_url"), knownvalue.StringExact("https://example.com/keys")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require").AtSliceIndex(0).AtMapKey("auth_method").AtMapKey("auth_method"), knownvalue.StringExact("hwk")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require").AtSliceIndex(1).AtMapKey("certificate"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require").AtSliceIndex(2).AtMapKey("any_valid_service_token"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_MultipleRuleTypesWithGeo(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigServiceTokens(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("test@example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("everyone"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude").AtSliceIndex(0).AtMapKey("geo").AtMapKey("country_code"), knownvalue.StringExact("RU")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require").AtSliceIndex(0).AtMapKey("auth_method").AtMapKey("auth_method"), knownvalue.StringExact("swk")),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}

func TestAccCloudflareAccessGroup_IPRangeRules(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_zero_trust_access_group.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareAccessGroupConfigLoginMethod(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(0).AtMapKey("email").AtMapKey("email"), knownvalue.StringExact("test@example.com")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("include").AtSliceIndex(1).AtMapKey("ip").AtMapKey("ip"), knownvalue.StringExact("192.0.2.0/24")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("exclude").AtSliceIndex(0).AtMapKey("ip").AtMapKey("ip"), knownvalue.StringExact("192.0.2.100/32")),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("require").AtSliceIndex(0).AtMapKey("email_domain").AtMapKey("domain"), knownvalue.StringExact("company.com")),
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareAccessGroupExists(name, cloudflare.AccountIdentifier(accountID), &accessGroup),
				),
			},
			{
				ResourceName:            name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_default"},
				ImportStateIdPrefix:     fmt.Sprintf("accounts/%s/", accountID),
			},
		},
	})
}
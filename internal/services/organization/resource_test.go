package organization_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6/organizations"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TestMain is the entry point for test execution

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_organization", &resource.Sweeper{
		Name: "cloudflare_organization",
		F:    testSweepCloudflareOrgs,
	})
}

func testSweepCloudflareOrgs(_ string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	orgID := os.Getenv("CLOUDFLARE_ORGANIZATION_ID")

	if orgID == "" {
		tflog.Info(ctx, "Skipping organizations sweep: CLOUDFLARE_ORGANIZATION_ID not set")
		return nil
	}

	orgs, err := client.Organizations.List(ctx, organizations.OrganizationListParams{})
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to fetch organizations: %s",err))
		return err
	}
	if len(orgs.Result) == 0 {
		tflog.Info(ctx, "No Cloudflare organizations to sweep")
		return nil
	}

	for _, org := range orgs.Result {
		if !utils.ShouldSweepResource(org.Name) {
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleting organization: %s (%s)", org.Name, org.ID))
		_, err = client.Organizations.Delete(ctx, org.ID)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete organization %s (%s): %s", org.Name, org.ID,err))
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleted organization: %s (%s)", org.Name, org.ID))
	}
	return nil
}

// TestAccCloudflareOrganization_Basic tests the basic CRUD operations for organization resource
func TestAccCloudflareOrganization_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_organization." + rnd
	orgName := rnd
	updatedOrgName := rnd + "-updated"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareOrganizationDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create - Test resource creation with all required attributes
			{
				Config: testAccOrganizationConfig(rnd, orgName),
				Check: resource.ComposeTestCheckFunc(
					// Verify required attributes
					resource.TestCheckResourceAttr(resourceName, "name", orgName),
					// Verify computed attributes are set
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "create_time"),
					// Verify meta attributes
					resource.TestCheckResourceAttrSet(resourceName, "meta.%"),
				),
			},
			// Step 2: Update - Test modifying updatable attributes
			{
				Config: testAccOrganizationConfig(rnd, updatedOrgName),
				Check: resource.ComposeTestCheckFunc(
					// Verify the name was updated
					resource.TestCheckResourceAttr(resourceName, "name", updatedOrgName),
					// Verify ID remains the same (it should be set)
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					// Verify other attributes remain consistent
					resource.TestCheckResourceAttrSet(resourceName, "create_time"),
				),
			},
			// Step 3: Import - Test import functionality with proper ID format
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// Organization import uses just the ID, no prefix needed
			},
		},
	})
}

// TestAccCloudflareOrganization_WithProfile tests organization creation with profile information
func TestAccCloudflareOrganization_WithProfile(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_organization." + rnd
	orgName := rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareOrganizationDestroy,
		Steps: []resource.TestStep{
			// Create organization with profile
			{
				Config:             testAccOrganizationConfigWithProfile(rnd, orgName),
				ExpectNonEmptyPlan: true, // Allow non-empty plan due to profile field handling
				Check: resource.ComposeTestCheckFunc(
					// Basic attribute checks
					resource.TestCheckResourceAttr(resourceName, "name", orgName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "create_time"),
					// Profile checks
					resource.TestCheckResourceAttr(resourceName, "profile.business_name", "Test Business"),
					resource.TestCheckResourceAttr(resourceName, "profile.business_email", "test@example.com"),
					resource.TestCheckResourceAttr(resourceName, "profile.business_phone", "+1234567890"),
					resource.TestCheckResourceAttr(resourceName, "profile.business_address", "123 Test St, Test City, TC 12345"),
				),
			},
			// Update profile information
			{
				Config:             testAccOrganizationConfigWithProfileUpdated(rnd, orgName),
				ExpectNonEmptyPlan: true, // Allow non-empty plan due to profile field handling
				Check: resource.ComposeTestCheckFunc(
					// Verify profile was updated
					resource.TestCheckResourceAttr(resourceName, "profile.business_name", "Updated Business"),
					resource.TestCheckResourceAttr(resourceName, "profile.business_email", "updated@example.com"),
				),
			},
			// Import test
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,

				ImportStateVerifyIgnore: []string{
					"profile", // Profile fields not populated on import
					"profile.%",
					"profile.business_name",
					"profile.business_email",
					"profile.business_phone",
					"profile.business_address",
					"profile.external_metadata",
				},
			},
		},
	})
}

// Test configuration functions that load from testdata files

func testAccOrganizationConfig(rnd, name string) string {
	return acctest.LoadTestCase("basic.tf", rnd, name)
}

func testAccOrganizationConfigWithProfile(rnd, name string) string {
	return acctest.LoadTestCase("with_profile.tf", rnd, name)
}

func testAccOrganizationConfigWithProfileUpdated(rnd, name string) string {
	return acctest.LoadTestCase("with_profile_updated.tf", rnd, name)
}

func testAccOrganizationConfigWithParent(rnd, name, parentID string) string {
	return acctest.LoadTestCase("with_parent.tf", rnd, name, parentID)
}

// testAccCheckCloudflareOrganizationDestroy verifies the organization has been destroyed
func testAccCheckCloudflareOrganizationDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_organization" {
			continue
		}

		// Try to fetch the organization
		_, err := client.Organizations.Get(context.Background(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("organization %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

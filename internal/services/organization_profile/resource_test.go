package organization_profile_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// TestMain is the entry point for test execution

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_organization_profile", &resource.Sweeper{
		Name: "cloudflare_organization_profile",
		F:    testSweepCloudflareOrganizationProfile,
	})
}

func testSweepCloudflareOrganizationProfile(r string) error {
	ctx := context.Background()
	// Organization Profile is an organization-level configuration setting.
	// It's a singleton setting per organization, not something that accumulates.
	// No sweeping required.
	tflog.Info(ctx, "Organization Profile doesn't require sweeping (organization setting)")
	return nil
}

// TestAccCloudflareOrganizationProfile_Basic tests the basic CRUD operations for organization_profile resource
func TestAccCloudflareOrganizationProfile_Basic(t *testing.T) {
	// Skip if no organization ID is provided
	orgID := os.Getenv("CLOUDFLARE_ORGANIZATION_ID")
	if orgID == "" {
		t.Skip("CLOUDFLARE_ORGANIZATION_ID not set, skipping organization profile test")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_organization_profile." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		// Note: Organization profiles cannot be destroyed via API, only updated
		// No CheckDestroy as the resource doesn't support deletion
		Steps: []resource.TestStep{
			// Step 1: Create - Test resource creation with all required attributes
			{
				Config: testAccOrganizationProfileConfig(rnd, orgID, 
					"Test Business", 
					"test@example.com", 
					"+1234567890", 
					`{\"line1\":\"123 Test St\",\"line2\":\"\",\"country\":\"US\",\"zipcode\":\"12345\",\"city\":\"Test City\",\"stateOrProvince\":\"TC\"}`,
					`{\"department\":\"IT\"}`),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify required attributes
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("organization_id"), knownvalue.StringExact(orgID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_name"), knownvalue.StringExact("Test Business")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_email"), knownvalue.StringExact("test@example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_phone"), knownvalue.StringExact("+1234567890")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_address"), knownvalue.StringExact(`{"line1":"123 Test St","line2":"","country":"US","zipcode":"12345","city":"Test City","stateOrProvince":"TC"}`)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("external_metadata"), knownvalue.StringExact(`{"department":"IT"}`)),
				},
			},
			// Step 2: Read - Verify all attributes are correctly set (implicit in state checks)
			
			// Step 3: Update - Test modifying updatable attributes
			{
				Config: testAccOrganizationProfileConfig(rnd, orgID,
					"Updated Business Name",
					"updated@example.com",
					"+9876543210",
					`{\"line1\":\"456 Updated Ave\",\"line2\":\"Suite 200\",\"country\":\"US\",\"zipcode\":\"54321\",\"city\":\"New City\",\"stateOrProvince\":\"NC\"}`,
					`{\"department\":\"Engineering\",\"team\":\"Platform\"}`),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify the attributes were updated
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("organization_id"), knownvalue.StringExact(orgID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_name"), knownvalue.StringExact("Updated Business Name")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_email"), knownvalue.StringExact("updated@example.com")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_phone"), knownvalue.StringExact("+9876543210")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_address"), knownvalue.StringExact(`{"line1":"456 Updated Ave","line2":"Suite 200","country":"US","zipcode":"54321","city":"New City","stateOrProvince":"NC"}`)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("external_metadata"), knownvalue.StringExact(`{"department":"Engineering","team":"Platform"}`)),
				},
			},
			// Note: No import test as this resource doesn't support import functionality
		},
	})
}


// TestAccCloudflareOrganizationProfile_MinimalMetadata tests with minimal external metadata
func TestAccCloudflareOrganizationProfile_MinimalMetadata(t *testing.T) {
	// Skip if no organization ID is provided
	orgID := os.Getenv("CLOUDFLARE_ORGANIZATION_ID")
	if orgID == "" {
		t.Skip("CLOUDFLARE_ORGANIZATION_ID not set, skipping organization profile test")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_organization_profile." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimal metadata
			{
				Config: testAccOrganizationProfileConfig(rnd, orgID,
					"Minimal Business",
					"minimal@example.com",
					"+1111111111",
					`{\"line1\":\"111 Minimal St\",\"line2\":\"\",\"country\":\"US\",\"zipcode\":\"11111\",\"city\":\"Minimal City\",\"stateOrProvince\":\"MC\"}`,
					"{}"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("organization_id"), knownvalue.StringExact(orgID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_name"), knownvalue.StringExact("Minimal Business")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_address"), knownvalue.StringExact(`{"line1":"111 Minimal St","line2":"","country":"US","zipcode":"11111","city":"Minimal City","stateOrProvince":"MC"}`)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("external_metadata"), knownvalue.StringExact("{}")),
				},
			},
			// Update to add more metadata
			{
				Config: testAccOrganizationProfileConfig(rnd, orgID,
					"Minimal Business",
					"minimal@example.com",
					"+1111111111",
					`{\"line1\":\"111 Minimal St\",\"line2\":\"\",\"country\":\"US\",\"zipcode\":\"11111\",\"city\":\"Minimal City\",\"stateOrProvince\":\"MC\"}`,
					`{\"status\":\"active\"}`),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("external_metadata"), knownvalue.StringExact(`{"status":"active"}`)),
				},
			},
		},
	})
}

// TestAccCloudflareOrganizationProfile_UpdateSingleField tests updating individual fields
func TestAccCloudflareOrganizationProfile_UpdateSingleField(t *testing.T) {
	// Skip if no organization ID is provided
	orgID := os.Getenv("CLOUDFLARE_ORGANIZATION_ID")
	if orgID == "" {
		t.Skip("CLOUDFLARE_ORGANIZATION_ID not set, skipping organization profile test")
	}

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_organization_profile." + rnd
	
	initialEmail := "initial@example.com"
	updatedEmail := "updated@example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Initial configuration
			{
				Config: testAccOrganizationProfileConfig(rnd, orgID,
					"Test Company",
					initialEmail,
					"+1234567890",
					`{\"line1\":\"123 Main St\",\"line2\":\"\",\"country\":\"US\",\"zipcode\":\"10001\",\"city\":\"Test City\",\"stateOrProvince\":\"NY\"}`,
					"{}"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_email"), knownvalue.StringExact(initialEmail)),
				},
			},
			// Update only the email
			{
				Config: testAccOrganizationProfileConfig(rnd, orgID,
					"Test Company",
					updatedEmail,
					"+1234567890",
					`{\"line1\":\"123 Main St\",\"line2\":\"\",\"country\":\"US\",\"zipcode\":\"10001\",\"city\":\"Test City\",\"stateOrProvince\":\"NY\"}`,
					"{}"),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify only email changed
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_email"), knownvalue.StringExact(updatedEmail)),
					// Verify other fields remain unchanged
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_name"), knownvalue.StringExact("Test Company")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_phone"), knownvalue.StringExact("+1234567890")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("business_address"), knownvalue.StringExact(`{"line1":"123 Main St","line2":"","country":"US","zipcode":"10001","city":"Test City","stateOrProvince":"NY"}`)),
				},
			},
		},
	})
}



// Test configuration functions that load from testdata files

func testAccOrganizationProfileConfig(rnd, orgID, businessName, businessEmail, businessPhone, businessAddress, externalMetadata string) string {
	return acctest.LoadTestCase("basic.tf", rnd, orgID, businessName, businessEmail, businessPhone, businessAddress, externalMetadata)
}

// Note: No CheckDestroy function as organization profiles cannot be deleted via API
// The Delete operation in the resource is a no-op that only removes from Terraform state

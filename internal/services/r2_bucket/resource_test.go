package r2_bucket_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/r2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_r2_bucket", &resource.Sweeper{
		Name: "cloudflare_r2_bucket",
		F:    testSweepCloudflareR2Bucket,
	})
}

func testSweepCloudflareR2Bucket(r string) error {
	ctx := context.Background()
	client, err := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if err != nil {
		return fmt.Errorf("error establishing client: %w", err)
	}

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	if accountID == "" {
		return nil // Skip sweep if account ID not set
	}

	buckets, err := client.ListR2Buckets(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListR2BucketsParams{})
	if err != nil {
		return fmt.Errorf("failed to fetch R2 buckets: %w", err)
	}

	for _, bucket := range buckets {
		// Use standard filtering helper
		if !utils.ShouldSweepResource(bucket.Name) {
			continue
		}

		err = client.DeleteR2Bucket(ctx, cfv1.AccountIdentifier(accountID), bucket.Name)
		if err != nil {
			// Log error but continue with other buckets
			// Buckets with objects will fail to delete, which is expected
			fmt.Printf("Warning: failed to delete R2 bucket %q: %v\n", bucket.Name, err)
			continue
		}
	}

	return nil
}

func TestAccCloudflareR2Bucket_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareR2BucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareR2BucketBasic(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("location"), knownvalue.StringExact("ENAM")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.StringExact("Standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				Config: testAccCheckCloudflareR2BucketUpdate(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.StringExact("InfrequentAccess")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("location"), knownvalue.StringExact("ENAM")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.StringExact("InfrequentAccess")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				ResourceName: resourceName,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return strings.Join([]string{accountID, rnd, "default"}, "/"), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareR2Bucket_Minimum(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareR2BucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareR2BucketMinimum(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
			{
				ResourceName: resourceName,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return strings.Join([]string{accountID, rnd, "default"}, "/"), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareR2Bucket_Jurisdiction(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareR2BucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareR2BucketJurisdiction(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
				},
			},
			{
				ResourceName: resourceName,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return strings.Join([]string{accountID, rnd, "eu"}, "/"), nil
				},
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareR2Bucket_AllLocations(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	locations := []string{"apac", "eeur", "enam", "weur", "wnam", "oc"}

	for _, location := range locations {
		t.Run(location, func(t *testing.T) {
			testRnd := rnd + location
			testResourceName := "cloudflare_r2_bucket." + testRnd

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				CheckDestroy:             testAccCheckCloudflareR2BucketDestroy,
				Steps: []resource.TestStep{
					{
						Config: testAccCheckCloudflareR2BucketLocation(testRnd, accountID, location),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testResourceName, "name", testRnd),
							resource.TestCheckResourceAttr(testResourceName, "location", location),
							resource.TestCheckResourceAttr(testResourceName, "account_id", accountID),
							resource.TestCheckResourceAttr(testResourceName, "jurisdiction", "default"),
							resource.TestCheckResourceAttr(testResourceName, "storage_class", "Standard"),
						),
					},
				},
			})
		})
	}
}

func TestAccCloudflareR2Bucket_AllJurisdictions(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	jurisdictions := []string{"default", "fedramp"}

	for _, jurisdiction := range jurisdictions {
		t.Run(jurisdiction, func(t *testing.T) {
			testRnd := rnd + jurisdiction
			testResourceName := "cloudflare_r2_bucket." + testRnd

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					acctest.TestAccPreCheck(t)
					acctest.TestAccPreCheck_AccountID(t)
				},
				ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
				CheckDestroy:             testAccCheckCloudflareR2BucketDestroy,
				Steps: []resource.TestStep{
					{
						Config: testAccCheckCloudflareR2BucketJurisdictionSpecific(testRnd, accountID, jurisdiction),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(testResourceName, tfjsonpath.New("name"), knownvalue.StringExact(testRnd)),
							statecheck.ExpectKnownValue(testResourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact(jurisdiction)),
							statecheck.ExpectKnownValue(testResourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(testResourceName, tfjsonpath.New("storage_class"), knownvalue.StringExact("Standard")),
						},
					},
				},
			})
		})
	}
}

func TestAccCloudflareR2Bucket_ComprehensiveConfiguration(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareR2BucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareR2BucketComprehensive(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("location"), knownvalue.StringExact("weur")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.StringExact("InfrequentAccess")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("creation_date"), knownvalue.NotNull()),
				},
				ExpectNonEmptyPlan: false,
			},
			{
				ResourceName: resourceName,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return strings.Join([]string{accountID, rnd, "eu"}, "/"), nil
				},
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func TestAccCloudflareR2Bucket_DefaultValues(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareR2BucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareR2BucketMinimum(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.StringExact("Standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("location"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("creation_date"), knownvalue.NotNull()),
				},
			},
		},
	})
}

func TestAccCloudflareR2Bucket_StorageClassUpdate(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareR2BucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareR2BucketStorageClass(rnd, accountID, "Standard"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.StringExact("Standard")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
				},
			},
			{
				Config: testAccCheckCloudflareR2BucketStorageClass(rnd, accountID, "InfrequentAccess"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.StringExact("InfrequentAccess")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("storage_class"), knownvalue.StringExact("InfrequentAccess")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
				},
			},
		},
	})
}

func TestAccCloudflareR2Bucket_LocationCaseInsensitive(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "weur"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "weur"),
				),
			},
			{

				// Apply with uppercase - should be treated as no change (case-insensitive)
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "WEUR"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "WEUR"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				// Apply with mixed case - should be treated as no change (case-insensitive)
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "WeUr"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "WEUR"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				// Now actually apply with uppercase to set it in state
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "WEUR"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "WEUR"),
				),
			},
			{
				// Reapply same uppercase - should not cause a plan change
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "WEUR"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "WEUR"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				// Apply with different case - should be treated as no change (case-insensitive)
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "weur"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "WEUR"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccCloudflareR2Bucket_LocationIgnoreChange(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "weur"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "weur"),
				),
			},
			{
				// Apply with EEUR - should not be changed
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "EEUR"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "weur"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				// Apply with APAC - should not be changed
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "apac"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "weur"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func testAccCheckCloudflareR2BucketMinimum(rnd, accountID string) string {
	return acctest.LoadTestCase("r2bucketminimum.tf", rnd, accountID)
}

func testAccCheckCloudflareR2BucketBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("r2bucketbasic.tf", rnd, accountID)
}

func testAccCheckCloudflareR2BucketUpdate(rnd, accountID string) string {
	return acctest.LoadTestCase("r2bucketupdate.tf", rnd, accountID)
}

func testAccCheckCloudflareR2BucketJurisdiction(rnd, accountID string) string {
	return acctest.LoadTestCase("jurisdiction.tf", rnd, accountID)
}

func testAccCheckCloudflareR2BucketLocation(rnd, accountID, location string) string {
	return acctest.LoadTestCase("location.tf", rnd, accountID, location)
}

func testAccCheckCloudflareR2BucketJurisdictionSpecific(rnd, accountID, jurisdiction string) string {
	return acctest.LoadTestCase("jurisdiction_specific.tf", rnd, accountID, jurisdiction)
}

func testAccCheckCloudflareR2BucketComprehensive(rnd, accountID string) string {
	return acctest.LoadTestCase("comprehensive.tf", rnd, accountID)
}

func testAccCheckCloudflareR2BucketStorageClass(rnd, accountID, storageClass string) string {
	return acctest.LoadTestCase("storage_class.tf", rnd, accountID, storageClass)
}

func testAccCheckCloudflareR2BucketDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_r2_bucket" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		jurisdiction := rs.Primary.Attributes["jurisdiction"]
		_, err := client.R2.Buckets.Get(
			context.Background(),
			rs.Primary.ID,
			r2.BucketGetParams{
				AccountID:    cloudflare.F(accountID),
				Jurisdiction: cloudflare.F(r2.BucketGetParamsCfR2Jurisdiction(jurisdiction)),
			},
		)
		if err == nil {
			return fmt.Errorf("r2 bucket still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, location string) string {
	return fmt.Sprintf(`
resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  location   = "%[3]s"
}`, rnd, accountID, location)
}

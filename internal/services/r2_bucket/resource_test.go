package r2_bucket_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/r2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	// TODO: fixme - auth error
	//resource.AddTestSweepers("cloudflare_r2_bucket", &resource.Sweeper{
	//	Name: "cloudflare_r2_bucket",
	//	F:    testSweepCloudflareR2Bucket,
	//})
}

func testSweepCloudflareR2Bucket(r string) error {
	client, err := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	accessKeyId := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_SECRET")

	if accessKeyId == "" {
		return errors.New("CLOUDFLARE_R2_ACCESS_KEY_ID must be set for this acceptance test")
	}

	if accessKeyId == "" {
		return errors.New("CLOUDFLARE_R2_ACCESS_KEY_SECRET must be set for this acceptance test")
	}

	if err != nil {
		return fmt.Errorf("error establishing client: %w", err)
	}

	ctx := context.Background()
	buckets, err := client.ListR2Buckets(ctx, cfv1.AccountIdentifier(accountID), cfv1.ListR2BucketsParams{})
	if err != nil {
		return fmt.Errorf("failed to fetch R2 buckets: %w", err)
	}

	for _, bucket := range buckets {
		// hard coded bucket name for Worker script acceptance tests
		// until we can break out the packages without cyclic errors.
		if bucket.Name == "bnfywlzwpt" {
			continue
		}

		r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID),
			}, nil
		})

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithEndpointResolverWithOptions(r2Resolver),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
			config.WithRegion("auto"),
		)
		if err != nil {
			return err
		}

		s3client := s3.NewFromConfig(cfg)
		listObjectsOutput, err := s3client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: &bucket.Name,
		})
		if err != nil {
			return err
		}

		for _, object := range listObjectsOutput.Contents {
			_, err = s3client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
				Bucket: &bucket.Name,
				Key:    object.Key,
			})
			if err != nil {
				return err
			}
		}

		err = client.DeleteR2Bucket(ctx, cfv1.AccountIdentifier(accountID), bucket.Name)
		if err != nil {
			return fmt.Errorf("failed to delete R2 bucket %q: %w", bucket.Name, err)
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
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(testResourceName, tfjsonpath.New("name"), knownvalue.StringExact(testRnd)),
							statecheck.ExpectKnownValue(testResourceName, tfjsonpath.New("location"), knownvalue.StringExact(location)),
							statecheck.ExpectKnownValue(testResourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
							statecheck.ExpectKnownValue(testResourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
							statecheck.ExpectKnownValue(testResourceName, tfjsonpath.New("storage_class"), knownvalue.StringExact("Standard")),
						},
						ExpectNonEmptyPlan: true,
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
				ExpectNonEmptyPlan: true,
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
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "WEUR"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "weur"),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "WEUR"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "weur"),
				),
			},
			{
				Config: testAccCheckCloudflareR2BucketLocationCase(rnd, accountID, "WeUr"),
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

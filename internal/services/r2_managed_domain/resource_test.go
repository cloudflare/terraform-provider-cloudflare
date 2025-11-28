package r2_managed_domain_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	resource.AddTestSweepers("cloudflare_r2_managed_domain", &resource.Sweeper{
		Name: "cloudflare_r2_managed_domain",
		F:    testSweepCloudflareR2ManagedDomain,
	})
	// TODO: fixme - auth error
	//resource.AddTestSweepers("cloudflare_r2_bucket", &resource.Sweeper{
	//	Name: "cloudflare_r2_bucket",
	//	F:    testSweepCloudflareR2Bucket,
	//})
}

func testSweepCloudflareR2ManagedDomain(r string) error {
	ctx := context.Background()
	// R2 Managed Domain is a bucket-level configuration.
	// When R2 buckets are swept, managed domains are cleaned up automatically.
	// No sweeping required.
	tflog.Info(ctx, "R2 Managed Domain doesn't require sweeping (bucket-level resource)")
	return nil
}

func testSweepCloudflareR2Bucket(r string) error {
	client, err := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	accessKeyId := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_SECRET")

	if accessKeyId == "" {
		return errors.New("CLOUDFLARE_R2_ACCESS_KEY_ID must be set for this acceptance test")
	}

	if accessKeySecret == "" {
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

func TestAccCloudflareR2ManagedDomain_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_managed_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2ManagedDomainConfigEnable(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.NotNull()),
				},
				Check: testCheckResourceDomainMatchesBucketID(resourceName, "bucket_id", "domain"),
			},
			{
				Config: testAccR2ManagedDomainConfigUpdate(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("default")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.NotNull()),
				},
				Check: testCheckResourceDomainMatchesBucketID(resourceName, "bucket_id", "domain"),
			},
		},
	})
}

func TestAccCloudflareR2ManagedDomain_Jurisdiction(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_managed_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2ManagedDomainConfigJurisdiction(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("eu")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.NotNull()),
				},
				Check: testCheckResourceDomainMatchesBucketID(resourceName, "bucket_id", "domain"),
			},
		},
	})
}

func TestAccCloudflareR2ManagedDomain_JurisdictionFedramp(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_managed_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccR2ManagedDomainConfigJurisdictionFedramp(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("fedramp")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jurisdiction"), knownvalue.StringExact("fedramp")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("bucket_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("domain"), knownvalue.NotNull()),
				},
				Check: testCheckResourceDomainMatchesBucketID(resourceName, "bucket_id", "domain"),
			},
		},
	})
}

func testCheckResourceDomainMatchesBucketID(name, keyBucketID, keyDomain string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}

		// retrieve attributes from the instance
		bucketID, ok := rs.Primary.Attributes[keyBucketID]
		if !ok {
			return fmt.Errorf("%s not an attribute in resource %s", keyBucketID, name)
		}
		domain, ok := rs.Primary.Attributes[keyDomain]
		if !ok {
			return fmt.Errorf("%s not an attribute in resource %s", keyDomain, name)
		}

		bucketIDToDomain := "pub-" + bucketID + ".r2.dev"

		// check domain name
		if bucketIDToDomain != domain {
			return fmt.Errorf(
				"%s: Attribute '%s' expected %#v, got %#v",
				name,
				keyBucketID,
				bucketIDToDomain,
				domain)
		}
		return nil
	}
}

func testAccR2ManagedDomainConfigEnable(rnd string, accountID string) string {
	return acctest.LoadTestCase("r2managed_enable.tf", rnd, accountID)
}

func testAccR2ManagedDomainConfigJurisdiction(rnd string, accountID string) string {
	return acctest.LoadTestCase("r2managed_jurisdiction.tf", rnd, accountID)
}

func testAccR2ManagedDomainConfigJurisdictionFedramp(rnd string, accountID string) string {
	return acctest.LoadTestCase("r2managed_jurisdiction_fedramp.tf", rnd, accountID)
}

func testAccR2ManagedDomainConfigUpdate(rnd, accountID string) string {
	return acctest.LoadTestCase("r2managed_disable.tf", rnd, accountID)
}

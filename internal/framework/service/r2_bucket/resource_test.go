package r2_bucket_test

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
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_r2_bucket", &resource.Sweeper{
		Name: "cloudflare_r2_bucket",
		F: func(region string) error {
			client, err := acctest.SharedClient()
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
			buckets, err := client.ListR2Buckets(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListR2BucketsParams{})
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

				err = client.DeleteR2Bucket(ctx, cloudflare.AccountIdentifier(accountID), bucket.Name)
				if err != nil {
					return fmt.Errorf("failed to delete R2 bucket %q: %w", bucket.Name, err)
				}
			}

			return nil
		},
	})
}

func TestAccCloudflareR2Bucket_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareR2BucketBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, "location", "ENAM"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func TestAccCloudflareR2Bucket_Minimum(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_r2_bucket." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareR2BucketMinimum(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rnd),
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
				),
			},
		},
	})
}

func testAccCheckCloudflareR2BucketMinimum(rnd, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_r2_bucket" "%[1]s" {
    account_id = "%[2]s"
    name       = "%[1]s"
  }`, rnd, accountID)
}

func testAccCheckCloudflareR2BucketBasic(rnd, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_r2_bucket" "%[1]s" {
    account_id = "%[2]s"
    name       = "%[1]s"
	location   = "ENAM"
  }`, rnd, accountID)
}

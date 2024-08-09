package workers_script_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	scriptContent1    = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`
	scriptContent2    = `addEventListener('fetch', event => {event.respondWith(new Response('test 2'))});`
	moduleContent     = `export default { fetch() { return new Response('Hello world'); }, };`
	encodedWasm       = "AGFzbQEAAAAGgYCAgAAA" // wat source: `(module)`, so literally just an empty wasm module
	compatibilityDate = "2023-03-19"
	d1DatabaseID      = "ce8b95dc-b376-4ff8-9b9e-1801ed6d745d"
)

var (
	compatibilityFlags = []string{"nodejs_compat", "web_socket_compression"}
)

func TestAccCloudflareWorkerScript_MultiScriptEnt(t *testing.T) {
	t.Parallel()

	// var script cloudflare.WorkerScript
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_script." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			testAccCheckCloudflareWorkerScriptCreateBucket(t, rnd)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptInitial(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					// testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent1),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdate(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					// testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdateBinding(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					// testAccCheckCloudflareWorkerScriptExists(name, &script, []string{"MY_KV_NAMESPACE", "MY_PLAIN_TEXT", "MY_SECRET_TEXT", "MY_WASM", "MY_SERVICE_BINDING", "MY_BUCKET", "MY_QUEUE"}),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
				),
			},
		},
	})
}

func TestAccCloudflareWorkerScript_ModuleUpload(t *testing.T) {
	t.Parallel()

	// var script cloudflare.WorkerScript
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_workers_script." + rnd
	r2AccesKeyID := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
	r2AccesKeySecret := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_SECRET")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
			testAccCheckCloudflareWorkerScriptCreateBucket(t, rnd)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptUploadModule(rnd, accountID, r2AccesKeyID, r2AccesKeySecret),
				Check: resource.ComposeTestCheckFunc(
					// testAccCheckCloudflareWorkerScriptExists(name, &script, []string{"MY_DATABASE"}),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", moduleContent),
					resource.TestCheckResourceAttr(name, "compatibility_date", compatibilityDate),
					resource.TestCheckResourceAttr(name, "compatibility_flags.#", "2"),
					resource.TestCheckResourceAttr(name, "compatibility_flags.0", compatibilityFlags[0]),
					resource.TestCheckResourceAttr(name, "logpush", "true"),
					resource.TestCheckResourceAttr(name, "placement.0.mode", "smart"),
				),
			},
		},
	})
}

// We can't currently use `cloudflare_r2_bucket` here due to not being able to
// mix V5 and V6 protocol resources without circular dependencies. In an ideal
// world, this would all be handled by the inbuilt resource.
func testAccCheckCloudflareWorkerScriptCreateBucket(t *testing.T, rnd string) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accessKeyId := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_SECRET")

	if accessKeyId == "" {
		t.Fatal("CLOUDFLARE_R2_ACCESS_KEY_ID must be set for this acceptance test")
	}

	if accessKeyId == "" {
		t.Fatal("CLOUDFLARE_R2_ACCESS_KEY_SECRET must be set for this acceptance test")
	}

	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}
	_, err := client.CreateR2Bucket(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.CreateR2BucketParameters{Name: rnd})
	if err != nil {
		t.Fatalf("unable to create test bucket named %s: %v", rnd, err)
	}

	t.Cleanup(func() {
		r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID),
			}, nil
		})

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithEndpointResolverWithOptions(r2Resolver),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
			config.WithDefaultRegion("auto"),
		)
		if err != nil {
			t.Error(err)
		}

		s3client := s3.NewFromConfig(cfg)
		listObjectsOutput, err := s3client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: &rnd,
		})
		if err != nil {
			t.Error(err)
		}

		for _, object := range listObjectsOutput.Contents {
			_, err = s3client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
				Bucket: &rnd,
				Key:    object.Key,
			})
			if err != nil {
				t.Error(err)
			}
		}

		err = client.DeleteR2Bucket(context.Background(), cloudflare.AccountIdentifier(accountID), rnd)
		if err != nil {
			t.Errorf("Failed to clean up bucket named %s: %v", rnd, err)
		}
	})
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptInitial(rnd, accountID string) string {
	return acctest.LoadTestCase("workerscriptconfigmultiscriptinitial.tf", rnd, scriptContent1, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdate(rnd, accountID string) string {
	return acctest.LoadTestCase("workerscriptconfigmultiscriptupdate.tf", rnd, scriptContent2, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdateBinding(rnd, accountID string) string {
	return acctest.LoadTestCase("workerscriptconfigmultiscriptupdatebinding.tf", rnd, scriptContent2, encodedWasm, accountID)
}

func testAccCheckCloudflareWorkerScriptUploadModule(rnd, accountID, r2AccessKeyID, r2AccessKeySecret string) string {
	return acctest.LoadTestCase("workerscriptuploadmodule.tf", rnd, moduleContent, accountID, compatibilityDate, strings.Join(compatibilityFlags, `","`), r2AccessKeyID, r2AccessKeySecret, d1DatabaseID)
}

// func testAccCheckCloudflareWorkerScriptExists(n string, script *cloudflare.WorkerScript, bindings []string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

// 		rs, ok := s.RootModule().Resources[n]
// 		if !ok {
// 			return fmt.Errorf("not found: %s", n)
// 		}

// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("No Worker Script ID is set")
// 		}

// 		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
// 		if clientErr != nil {
// 			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
// 		}

// 		r, err := client.GetWorker(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.Attributes["name"])
// 		if err != nil {
// 			return err
// 		}

// 		if r.Script == "" {
// 			return fmt.Errorf("Worker Script not found")
// 		}

// 		name := strings.Replace(n, "cloudflare_workers_script.", "", -1)
// 		foundBindings, err := getWorkerScriptBindings(context.Background(), accountID, name, nil, client)
// 		if err != nil {
// 			return fmt.Errorf("cannot list script bindings: %w", err)
// 		}

// 		for _, binding := range bindings {
// 			if _, ok := foundBindings[binding]; !ok {
// 				return fmt.Errorf("cannot find binding with name %s", binding)
// 			}
// 		}

// 		*script = r.WorkerScript
// 		return nil
// 	}
// }

func testAccCheckCloudflareWorkerScriptDestroy(s *terraform.State) error {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_script" {
			continue
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		r, _ := client.GetWorker(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.Attributes["name"])

		if r.Script != "" {
			return fmt.Errorf("worker script with id %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

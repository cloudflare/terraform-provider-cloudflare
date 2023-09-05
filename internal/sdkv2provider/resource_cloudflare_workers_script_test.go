package sdkv2provider

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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	scriptContent1    = `addEventListener('fetch', event => {event.respondWith(new Response('test 1'))});`
	scriptContent2    = `addEventListener('fetch', event => {event.respondWith(new Response('test 2'))});`
	moduleContent     = `export default { fetch() { return new Response('Hello world'); }, };`
	encodedWasm       = "AGFzbQEAAAAGgYCAgAAA" // wat source: `(module)`, so literally just an empty wasm module
	compatibilityDate = "2023-03-19"
)

var (
	compatibilityFlags = []string{"nodejs_compat", "web_socket_compression"}
)

func TestAccCloudflareWorkerScript_MultiScriptEnt(t *testing.T) {
	t.Parallel()

	var script cloudflare.WorkerScript
	rnd := generateRandomResourceName()
	name := "cloudflare_worker_script." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
			testAccCheckCloudflareWorkerScriptCreateBucket(t, rnd)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptInitial(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent1),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdate(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
				),
			},
			{
				Config: testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdateBinding(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, []string{"MY_KV_NAMESPACE", "MY_PLAIN_TEXT", "MY_SECRET_TEXT", "MY_WASM", "MY_SERVICE_BINDING", "MY_BUCKET", "MY_QUEUE"}),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", scriptContent2),
				),
			},
		},
	})
}

func TestAccCloudflareWorkerScript_ModuleUpload(t *testing.T) {
	t.Parallel()

	var script cloudflare.WorkerScript
	rnd := generateRandomResourceName()
	name := "cloudflare_worker_script." + rnd
	r2AccesKeyID := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
	r2AccesKeySecret := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_SECRET")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
			testAccCheckCloudflareWorkerScriptCreateBucket(t, rnd)
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCloudflareWorkerScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkerScriptUploadModule(rnd, accountID, r2AccesKeyID, r2AccesKeySecret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareWorkerScriptExists(name, &script, nil),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "content", moduleContent),
					resource.TestCheckResourceAttr(name, "compatibility_date", compatibilityDate),
					resource.TestCheckResourceAttr(name, "compatibility_flags.#", "2"),
					resource.TestCheckResourceAttr(name, "compatibility_flags.0", compatibilityFlags[0]),
					resource.TestCheckResourceAttr(name, "logpush", "true"),
				),
			},
		},
	})
}

// We can't currently use `cloudflare_r2_bucket` here due to not being able to
// mix V5 and V6 protocol resources without circular dependencies. In an ideal
// world, this would all be handled by the inbuilt resource.
func testAccCheckCloudflareWorkerScriptCreateBucket(t *testing.T, rnd string) {
	client := testAccProvider.Meta().(*cloudflare.API)
	_, err := client.CreateR2Bucket(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.CreateR2BucketParameters{Name: rnd})
	if err != nil {
		t.Fatalf("unable to create test bucket named %s: %v", rnd, err)
	}

	t.Cleanup(func() {
		accessKeyId := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
		accessKeySecret := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_SECRET")

		r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID),
			}, nil
		})

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithEndpointResolverWithOptions(r2Resolver),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
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
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[3]s"
  name = "%[1]s"
  content = "%[2]s"
}`, rnd, scriptContent1, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdate(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[3]s"
  name = "%[1]s"
  content = "%[2]s"
}`, rnd, scriptContent2, accountID)
}

func testAccCheckCloudflareWorkerScriptConfigMultiScriptUpdateBinding(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_workers_kv_namespace" "%[1]s" {
	account_id = "%[4]s"
	title = "%[1]s"
}

resource "cloudflare_queue" "%[1]s" {
	account_id = "%[4]s"
	name = "%[1]s"
}

resource "cloudflare_worker_script" "%[1]s-service" {
	account_id = "%[4]s"
	name    = "%[1]s-service"
	content = "%[2]s"
}

resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[4]s"
  name    = "%[1]s"
  content = "%[2]s"

  kv_namespace_binding {
    name         = "MY_KV_NAMESPACE"
    namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
  }

  plain_text_binding {
    name = "MY_PLAIN_TEXT"
    text = "%[1]s"
  }

  secret_text_binding {
    name = "MY_SECRET_TEXT"
    text = "%[1]s"
  }

  webassembly_binding {
    name = "MY_WASM"
    module = "%[3]s"
  }

  r2_bucket_binding {
	name = "MY_BUCKET"
	bucket_name = "%[1]s"
  }

  service_binding {
	name = "MY_SERVICE_BINDING"
    service = cloudflare_worker_script.%[1]s-service.name
    environment = "production"
  }

  queue_binding {
    binding         = "MY_QUEUE"
    queue = cloudflare_queue.%[1]s.name
  }

}`, rnd, scriptContent2, encodedWasm, accountID)
}

func testAccCheckCloudflareWorkerScriptUploadModule(rnd, accountID, r2AccessKeyID, r2AccessKeySecret string) string {
	return fmt.Sprintf(`
	resource "cloudflare_logpush_job" "%[1]s" {
		enabled          = true
		account_id       = "%[3]s"
		name             = "%[1]s"
		logpull_options  = "fields=Event,EventTimestampMs,Outcome,Exceptions,Logs,ScriptName"
		destination_conf = "r2://%[1]s/date={DATE}?account-id=%[3]s&access-key-id=%[6]s&secret-access-key=%[7]s"
		dataset          = "workers_trace_events"
	}

resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[3]s"
  name = "%[1]s"
  content = "%[2]s"
  module = true
  compatibility_date = "%[4]s"
  compatibility_flags = ["%[5]s"]
  logpush = true

  depends_on = [cloudflare_logpush_job.%[1]s]
}`, rnd, moduleContent, accountID, compatibilityDate, strings.Join(compatibilityFlags, `","`), r2AccessKeyID, r2AccessKeySecret)
}

func testAccCheckCloudflareWorkerScriptExists(n string, script *cloudflare.WorkerScript, bindings []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Worker Script ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)

		r, err := client.GetWorker(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.Attributes["name"])
		if err != nil {
			return err
		}

		if r.Script == "" {
			return fmt.Errorf("Worker Script not found")
		}

		name := strings.Replace(n, "cloudflare_worker_script.", "", -1)
		foundBindings, err := getWorkerScriptBindings(context.Background(), accountID, name, client)
		if err != nil {
			return fmt.Errorf("cannot list script bindings: %w", err)
		}

		for _, binding := range bindings {
			if _, ok := foundBindings[binding]; !ok {
				return fmt.Errorf("cannot find binding with name %s", binding)
			}
		}

		*script = r.WorkerScript
		return nil
	}
}

func testAccCheckCloudflareWorkerScriptDestroy(s *terraform.State) error {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_worker_script" {
			continue
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		r, _ := client.GetWorker(context.Background(), cloudflare.AccountIdentifier(accountID), rs.Primary.Attributes["name"])

		if r.Script != "" {
			return fmt.Errorf("worker script with id %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

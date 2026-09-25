package image_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/images"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// tinyPNGBase64 is a valid base64-encoded 1x1 PNG, matching filebase64() output.
const tinyPNGBase64 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAADElEQVR42mP4z8AAAAMBAQD3A0FDAAAAAElFTkSuQmCC"

// defaultTestImageURL is the upload source for tests (fetched server-side).
// Override with CLOUDFLARE_IMAGE_URL.
const defaultTestImageURL = "https://www.gstatic.com/webp/gallery/1.jpg"

func testImageURL() string {
	if v := os.Getenv("CLOUDFLARE_IMAGE_URL"); v != "" {
		return v
	}
	return defaultTestImageURL
}

// defaultTestImageURLAlt is a second source to prove url changes force replace.
// Override with CLOUDFLARE_IMAGE_URL_ALT.
const defaultTestImageURLAlt = "https://www.gstatic.com/webp/gallery/2.jpg"

func testImageURLAlt() string {
	if v := os.Getenv("CLOUDFLARE_IMAGE_URL_ALT"); v != "" {
		return v
	}
	return defaultTestImageURLAlt
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_image", &resource.Sweeper{
		Name: "cloudflare_image",
		F:    testSweepCloudflareImage,
	})
}

func testSweepCloudflareImage(_ string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping cloudflare_image sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	iter := client.Images.V1.ListAutoPaging(ctx, images.V1ListParams{
		AccountID: cloudflare.F(accountID),
	})
	for iter.Next() {
		page := iter.Current()
		for _, img := range page.Images {
			// Only delete images created by acceptance tests (test-prefixed IDs).
			if !utils.ShouldSweepResource(img.ID) {
				continue
			}
			tflog.Info(ctx, fmt.Sprintf("Deleting image: %s (account: %s)", img.ID, accountID))
			_, err := client.Images.V1.Delete(ctx, img.ID, images.V1DeleteParams{
				AccountID: cloudflare.F(accountID),
			})
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to delete image %s: %s", img.ID, err))
				continue
			}
		}
	}
	if err := iter.Err(); err != nil {
		return fmt.Errorf("failed to list images for sweep: %w", err)
	}

	return nil
}

func testAccCheckCloudflareImageDestroy(s *terraform.State) error {
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_image" {
			continue
		}

		_, err := client.Images.V1.Get(context.Background(), rs.Primary.ID, images.V1GetParams{
			AccountID: cloudflare.F(accountID),
		})
		if err == nil {
			return fmt.Errorf("image %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccImageConfigBasic(rnd, accountID, url string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID, url)
}

func testAccImageConfigUpdate(rnd, accountID, url string) string {
	return acctest.LoadTestCase("update.tf", rnd, accountID, url)
}

func testAccImageConfigFileUpload(rnd, accountID, fileBase64 string) string {
	return acctest.LoadTestCase("file_upload.tf", rnd, accountID, fileBase64)
}

func testAccImageConfigMetadata(rnd, accountID, url string) string {
	return acctest.LoadTestCase("metadata.tf", rnd, accountID, url)
}

func testAccImageConfigNoMetadata(rnd, accountID, url string) string {
	return acctest.LoadTestCase("no_metadata.tf", rnd, accountID, url)
}

// TestAccCloudflareImage_NoMetadataCreate covers create via URL with no metadata.
// Regression guard for Bug A (empty metadata part → API code 5400).
func TestAccCloudflareImage_NoMetadataCreate(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image." + rnd
	url := testImageURL()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageConfigNoMetadata(rnd, accountID, url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttrSet(resourceName, "filename"),
					resource.TestCheckResourceAttrSet(resourceName, "uploaded"),
				),
			},
		},
	})
}

// TestAccCloudflareImage_FileUpload covers create via the base64 `file` input.
// Regression guard for Bug B / issue #6176 (file sent without image
// content-type → HTTP 415 / code 5455).
func TestAccCloudflareImage_FileUpload(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageConfigFileUpload(rnd, accountID, tinyPNGBase64),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttrSet(resourceName, "filename"),
					resource.TestCheckResourceAttrSet(resourceName, "uploaded"),
					resource.TestCheckResourceAttrSet(resourceName, "variants.#"),
				),
			},
		},
	})
}

// TestAccCloudflareImage_Basic covers create (upload by URL), read of computed
// fields, and import.
func TestAccCloudflareImage_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image." + rnd
	url := testImageURL()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageConfigBasic(rnd, accountID, url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, "require_signed_urls", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "filename"),
					resource.TestCheckResourceAttrSet(resourceName, "uploaded"),
					resource.TestCheckResourceAttrSet(resourceName, "variants.#"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// account_id/image_id composite import ID.
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
				// url/file/metadata are write-only (no_refresh); uploaded rounds
				// inconsistently; variants is returned in an unstable order (Bug D).
				ImportStateVerifyIgnore: []string{"url", "file", "metadata", "uploaded", "variants"},
			},
		},
	})
}

// TestAccCloudflareImage_Update updates `metadata` in place and asserts Update,
// not replace. Regression guard for Bug C (multipart sent to JSON PATCH → code
// 5400). Note: require_signed_urls can't be toggled true — custom `id` + private
// is rejected (5410), and `id` is Required.
func TestAccCloudflareImage_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image." + rnd
	url := testImageURL()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageConfigBasic(rnd, accountID, url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "require_signed_urls", "false"),
				),
			},
			{
				Config: testAccImageConfigUpdate(rnd, accountID, url),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, "require_signed_urls", "false"),
				),
			},
		},
	})
}

// TestAccCloudflareImage_Metadata covers create with `creator` and `metadata`
// set. metadata is write-only; the stored value surfaces on computed `meta`.
func TestAccCloudflareImage_Metadata(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image." + rnd
	url := testImageURL()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageConfigMetadata(rnd, accountID, url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, "creator", "cftftest-creator"),
					resource.TestCheckResourceAttr(resourceName, "require_signed_urls", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "filename"),
					resource.TestCheckResourceAttrSet(resourceName, "meta"),
				),
			},
		},
	})
}

// TestAccCloudflareImage_RequiresReplaceOnURLChange proves changing `url` forces
// a replace, not an in-place update.
func TestAccCloudflareImage_RequiresReplaceOnURLChange(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageConfigBasic(rnd, accountID, testImageURL()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
				),
			},
			{
				Config: testAccImageConfigBasic(rnd, accountID, testImageURLAlt()),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionReplace),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, "url", testImageURLAlt()),
				),
			},
		},
	})
}

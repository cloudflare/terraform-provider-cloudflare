package image_variant_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/images"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccCloudflareImageVariant_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image_variant." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageVariantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageVariantBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "options.fit", "scale-down"),
					resource.TestCheckResourceAttr(resourceName, "options.height", "480"),
					resource.TestCheckResourceAttr(resourceName, "options.width", "640"),
					resource.TestCheckResourceAttr(resourceName, "options.metadata", "none"),
					resource.TestCheckResourceAttr(resourceName, "never_require_signed_urls", "false"),
				),
			},
		},
	})
}

func TestAccCloudflareImageVariant_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image_variant." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageVariantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageVariantUpdateInitial(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "options.fit", "scale-down"),
					resource.TestCheckResourceAttr(resourceName, "options.height", "480"),
					resource.TestCheckResourceAttr(resourceName, "options.width", "640"),
					resource.TestCheckResourceAttr(resourceName, "options.metadata", "none"),
					resource.TestCheckResourceAttr(resourceName, "never_require_signed_urls", "false"),
				),
			},
			{
				Config: testAccImageVariantUpdateModified(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "options.fit", "contain"),
					resource.TestCheckResourceAttr(resourceName, "options.height", "1080"),
					resource.TestCheckResourceAttr(resourceName, "options.width", "1920"),
					resource.TestCheckResourceAttr(resourceName, "options.metadata", "keep"),
					resource.TestCheckResourceAttr(resourceName, "never_require_signed_urls", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareImageVariant_Import(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image_variant." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageVariantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageVariantBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, rnd), nil
				},
			},
		},
	})
}

func testAccImageVariantBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID)
}

func testAccImageVariantUpdateInitial(rnd, accountID string) string {
	return acctest.LoadTestCase("update_initial.tf", rnd, accountID)
}

func testAccImageVariantUpdateModified(rnd, accountID string) string {
	return acctest.LoadTestCase("update_modified.tf", rnd, accountID)
}

func testAccCheckCloudflareImageVariantDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_image_variant" {
			continue
		}

		res := new(http.Response)
		_, err := client.Images.V1.Variants.Get(
			context.Background(),
			rs.Primary.ID,
			images.V1VariantGetParams{
				AccountID: cloudflare.F(rs.Primary.Attributes[consts.AccountIDSchemaKey]),
			},
			option.WithResponseBodyInto(&res),
		)
		// If we get a 404, the resource is properly destroyed
		if res != nil && res.StatusCode == 404 {
			continue
		}
		if err == nil {
			return fmt.Errorf("image variant %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

package image_variant_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/images"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("cloudflare_image_variant", &resource.Sweeper{
		Name: "cloudflare_image_variant",
		F:    testSweepCloudflareImageVariant,
	})
}

func testSweepCloudflareImageVariant(_ string) error {
	ctx := context.Background()
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if accountID == "" {
		tflog.Info(ctx, "Skipping cloudflare_image_variant sweep: CLOUDFLARE_ACCOUNT_ID not set")
		return nil
	}

	// The generated SDK list response only surfaces a hardcoded "hero" key, so
	// capture the raw response and decode the full variant map generically.
	res := new(http.Response)
	_, err := client.Images.V1.Variants.List(ctx, images.V1VariantListParams{
		AccountID: cloudflare.F(accountID),
	}, option.WithResponseBodyInto(&res))
	if err != nil {
		return fmt.Errorf("failed to list image variants for sweep: %w", err)
	}
	body, _ := io.ReadAll(res.Body)

	var listEnv struct {
		Result struct {
			Variants map[string]struct {
				ID string `json:"id"`
			} `json:"variants"`
		} `json:"result"`
	}
	if err := json.Unmarshal(body, &listEnv); err != nil {
		return fmt.Errorf("failed to decode image variant list for sweep: %w", err)
	}

	for key, variant := range listEnv.Result.Variants {
		id := variant.ID
		if id == "" {
			id = key
		}
		// Only delete variants created by acceptance tests (test-prefixed IDs).
		if !utils.ShouldSweepResource(id) {
			continue
		}
		tflog.Info(ctx, fmt.Sprintf("Deleting image variant: %s (account: %s)", id, accountID))
		_, err := client.Images.V1.Variants.Delete(ctx, id, images.V1VariantDeleteParams{
			AccountID: cloudflare.F(accountID),
		})
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("Failed to delete image variant %s: %s", id, err))
			continue
		}
	}

	return nil
}

func testAccCheckCloudflareImageVariantDestroy(s *terraform.State) error {
	client := acctest.SharedClient()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_image_variant" {
			continue
		}

		_, err := client.Images.V1.Variants.Get(context.Background(), rs.Primary.ID, images.V1VariantGetParams{
			AccountID: cloudflare.F(accountID),
		})
		if err == nil {
			return fmt.Errorf("image variant %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccImageVariantConfigBasic(rnd, accountID string) string {
	return acctest.LoadTestCase("basic.tf", rnd, accountID)
}

func testAccImageVariantConfigUpdate(rnd, accountID string) string {
	return acctest.LoadTestCase("update.tf", rnd, accountID)
}

func testAccImageVariantConfigAlt(rnd, accountID string) string {
	return acctest.LoadTestCase("alt.tf", rnd, accountID)
}

func testAccImageVariantConfigReplace(rnd, accountID, variantID string) string {
	return acctest.LoadTestCase("replace.tf", rnd, accountID, variantID)
}

// TestAccCloudflareImageVariant_Basic covers create, read of computed fields,
// and import.
func TestAccCloudflareImageVariant_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image_variant." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageVariantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageVariantConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, "options.fit", "scale-down"),
					resource.TestCheckResourceAttr(resourceName, "options.width", "100"),
					resource.TestCheckResourceAttr(resourceName, "options.height", "100"),
					resource.TestCheckResourceAttr(resourceName, "options.metadata", "none"),
					resource.TestCheckResourceAttr(resourceName, "never_require_signed_urls", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// account_id/variant_id composite import ID.
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflareImageVariant_Update proves options and
// never_require_signed_urls update in place (only id/account_id force replace).
func TestAccCloudflareImageVariant_Update(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image_variant." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageVariantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageVariantConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "options.fit", "scale-down"),
					resource.TestCheckResourceAttr(resourceName, "never_require_signed_urls", "false"),
				),
			},
			{
				Config: testAccImageVariantConfigUpdate(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, "options.fit", "cover"),
					resource.TestCheckResourceAttr(resourceName, "options.width", "200"),
					resource.TestCheckResourceAttr(resourceName, "options.height", "200"),
					resource.TestCheckResourceAttr(resourceName, "options.metadata", "copyright"),
					resource.TestCheckResourceAttr(resourceName, "never_require_signed_urls", "true"),
				),
			},
		},
	})
}

// TestAccCloudflareImageVariant_Idempotency is a regression guard for issue
// #5448 (refresh failed with "missing required variant_id parameter"). The
// second identical step forces refresh + plan and asserts an empty plan.
func TestAccCloudflareImageVariant_Idempotency(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image_variant." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageVariantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageVariantConfigBasic(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
				),
			},
			{
				// Re-apply the identical config: refresh (Read) must succeed and
				// the plan must be empty.
				Config: testAccImageVariantConfigBasic(rnd, accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
				),
			},
		},
	})
}

// TestAccCloudflareImageVariant_AltOptions covers create with alternate enums
// (fit="crop", metadata="keep"), never_require_signed_urls=true, plus import.
func TestAccCloudflareImageVariant_AltOptions(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image_variant." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageVariantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageVariantConfigAlt(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", rnd),
					resource.TestCheckResourceAttr(resourceName, "options.fit", "crop"),
					resource.TestCheckResourceAttr(resourceName, "options.width", "512"),
					resource.TestCheckResourceAttr(resourceName, "options.height", "384"),
					resource.TestCheckResourceAttr(resourceName, "options.metadata", "keep"),
					resource.TestCheckResourceAttr(resourceName, "never_require_signed_urls", "true"),
				),
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: fmt.Sprintf("%s/", accountID),
			},
		},
	})
}

// TestAccCloudflareImageVariant_RequiresReplaceOnIDChange proves changing `id`
// forces a replace, not an in-place update.
func TestAccCloudflareImageVariant_RequiresReplaceOnIDChange(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_image_variant." + rnd
	firstID := rnd
	secondID := rnd + "b"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareImageVariantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccImageVariantConfigReplace(rnd, accountID, firstID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", firstID),
				),
			},
			{
				Config: testAccImageVariantConfigReplace(rnd, accountID, secondID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionReplace),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", secondID),
				),
			},
		},
	})
}

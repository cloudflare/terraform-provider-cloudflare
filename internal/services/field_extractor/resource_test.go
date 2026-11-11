package field_extractor_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/field_extractors"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const fieldExtractorType = "llm_prompts"

func TestAccCloudflareFieldExtractor_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_field_extractor.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFieldExtractorDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("basic.tf", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", fieldExtractorType),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(resourceName, "extractor", fieldExtractorType),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ref", "chat-completions"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "Extract prompts from chat completion requests"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.fields.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.fields.0.name", "prompt"),
					resource.TestCheckResourceAttr(
						resourceName,
						"rules.0.fields.0.expression",
						`filter(http.request.body.json.strings.values, http.request.body.json.strings.pointers[*] matches "^/messages/\\d+/content$")`,
					),
				),
			},
			{
				Config: acctest.LoadTestCase("updated.tf", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("extractor"), knownvalue.StringExact(fieldExtractorType)),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("rules"), knownvalue.ListSizeExact(2)),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "rules.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.ref", "chat-completions"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", "Extract prompts from chat completion requests"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.fields.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.fields.0.name", "prompt"),
					resource.TestCheckResourceAttr(
						resourceName,
						"rules.0.fields.0.expression",
						`filter(http.request.body.json.strings.values, http.request.body.json.strings.pointers[*] matches "^/messages/\\d+/content$")`,
					),
					resource.TestCheckResourceAttr(resourceName, "rules.1.ref", "message-content"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.description", "Extract message content"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.fields.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.1.fields.0.name", "prompt"),
					resource.TestCheckResourceAttr(
						resourceName,
						"rules.1.fields.0.expression",
						`filter(http.request.body.json.strings.values, http.request.body.json.strings.pointers[*] matches "^/messages/[0-9]+/content$")`,
					),
				),
			},
			{
				Config: acctest.LoadTestCase("updated.tf", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     fmt.Sprintf("%s/%s", accountID, fieldExtractorType),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudflareFieldExtractor_ExternalDeletion(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_field_extractor.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFieldExtractorDestroy,
		Steps: []resource.TestStep{
			{
				Config: acctest.LoadTestCase("basic.tf", accountID),
			},
			{
				PreConfig: func() {
					testAccDeleteFieldExtractor(t, accountID)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
			},
			{
				Config: acctest.LoadTestCase("basic.tf", accountID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply:             []plancheck.PlanCheck{plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate)},
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
		},
	})
}

func testAccDeleteFieldExtractor(t *testing.T, accountID string) {
	t.Helper()

	client := testAccClient()
	_, err := client.FieldExtractors.Delete(
		context.Background(),
		fieldExtractorType,
		field_extractors.FieldExtractorDeleteParams{AccountID: cloudflare.F(accountID)},
	)
	if err != nil {
		t.Fatalf("failed to delete field extractor outside Terraform: %v", err)
	}
}

func testAccCheckFieldExtractorDestroy(s *terraform.State) error {
	client := testAccClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_field_extractor" {
			continue
		}

		_, err := client.FieldExtractors.Get(
			context.Background(),
			rs.Primary.ID,
			field_extractors.FieldExtractorGetParams{
				AccountID: cloudflare.F(rs.Primary.Attributes["account_id"]),
			},
		)
		if err == nil {
			return fmt.Errorf("field extractor %s still exists", rs.Primary.ID)
		}

		var apiErr *cloudflare.Error
		if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return fmt.Errorf("error checking if field extractor %s was destroyed: %w", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccClient() *cloudflare.Client {
	var opts []option.RequestOption
	if apiToken := os.Getenv(consts.APITokenEnvVarKey); apiToken != "" {
		opts = append(opts, option.WithAPIToken(apiToken))
	} else if serviceKey := os.Getenv(consts.APIUserServiceKeyEnvVarKey); serviceKey != "" {
		opts = append(opts, option.WithUserServiceKey(serviceKey))
	} else {
		opts = append(opts,
			option.WithAPIKey(os.Getenv(consts.APIKeyEnvVarKey)),
			option.WithAPIEmail(os.Getenv(consts.EmailEnvVarKey)),
		)
	}

	return cloudflare.NewClient(opts...)
}

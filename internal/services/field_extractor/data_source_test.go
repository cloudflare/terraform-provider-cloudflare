package field_extractor_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccCloudflareFieldExtractorDataSource_Basic(t *testing.T) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := "data.cloudflare_field_extractor.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFieldExtractorDestroy,
		Steps: []resource.TestStep{
			{
				Config:                    acctest.LoadTestCase("datasource_basic.tf", accountID),
				PreventPostDestroyRefresh: true,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "account_id", accountID),
					resource.TestCheckResourceAttr(dataSourceName, "extractor", fieldExtractorType),
					resource.TestCheckResourceAttr(dataSourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "rules.0.ref", "chat-completions"),
					resource.TestCheckResourceAttr(dataSourceName, "rules.0.description", "Extract prompts from chat completion requests"),
					resource.TestCheckResourceAttr(dataSourceName, "rules.0.fields.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "rules.0.fields.0.name", "prompt"),
					resource.TestCheckResourceAttr(
						dataSourceName,
						"rules.0.fields.0.expression",
						`filter(http.request.body.json.strings.values, http.request.body.json.strings.pointers[*] matches "^/messages/\\d+/content$")`,
					),
				),
			},
		},
	})
}

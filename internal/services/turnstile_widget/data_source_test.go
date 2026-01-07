package turnstile_widget_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccCloudflareTurnstileWidgetDataSource_BySitekey(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	resourceName := "cloudflare_turnstile_widget." + rnd
	dataSourceName := "data.cloudflare_turnstile_widget." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTurnstileWidgetDataSource_BySitekey(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("mode"), knownvalue.StringExact("managed")),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("sitekey"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("secret"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("modified_on"), knownvalue.NotNull()),
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "sitekey", resourceName, "sitekey"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "mode", resourceName, "mode"),
				),
			},
		},
	})
}

func TestAccCloudflareTurnstileWidgetsDataSource_List(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	dataSourceName := "data.cloudflare_turnstile_widgets." + rnd

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareTurnstileWidgetsDataSource_List(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("result"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New("account_id"), knownvalue.StringExact(accountID)),
				},
			},
		},
	})
}

func testAccCloudflareTurnstileWidgetDataSource_BySitekey(rnd, accountID string) string {
	return acctest.LoadTestCase("datasourceturnstilewidgetbysitekey.tf", rnd, accountID)
}

func testAccCloudflareTurnstileWidgetsDataSource_List(rnd, accountID string) string {
	return acctest.LoadTestCase("datasourceturnstilewidgetslist.tf", rnd, accountID)
}

package workers_script_subdomain_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccCloudflareWorkersScriptSubdomain_Basic(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script_subdomain." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersScriptSubdomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersScriptSubdomain(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(consts.AccountIDSchemaKey), knownvalue.StringExact(accountID)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, rnd), nil
				},
			},
		},
	})
}

func TestAccCloudflareWorkersScriptSubdomain_Update(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script_subdomain." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersScriptSubdomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersScriptSubdomain(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
				},
			},
			{
				Config: testAccCheckCloudflareWorkersScriptSubdomainDisabled(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(false)),
				},
			},
		},
	})
}

func TestAccCloudflareWorkersScriptSubdomain_WithPreviews(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script_subdomain." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersScriptSubdomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersScriptSubdomainWithPreviews(rnd, accountID),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("script_name"), knownvalue.StringExact(rnd)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("previews_enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("%s/%s", accountID, rnd), nil
				},
			},
		},
	})
}

func TestAccCloudflareWorkersScriptSubdomain_InvalidImportID(t *testing.T) {
	t.Parallel()

	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_workers_script_subdomain." + rnd
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			acctest.TestAccPreCheck_AccountID(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareWorkersScriptSubdomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWorkersScriptSubdomain(rnd, accountID),
			},
			{
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: "invalid-import-id",
				ExpectError:   regexp.MustCompile("invalid ID|expected urlencoded segments"),
			},
			{
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: accountID,
				ExpectError:   regexp.MustCompile("invalid ID|expected urlencoded segments"),
			},
		},
	})
}

func testAccCheckCloudflareWorkersScriptSubdomainDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_workers_script_subdomain" {
			continue
		}

		scriptName := rs.Primary.Attributes["script_name"]
		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]

		_, err := client.Workers.Scripts.Subdomain.Get(context.Background(), scriptName, workers.ScriptSubdomainGetParams{
			AccountID: cloudflare.F(accountID),
		})

		if err == nil {
			return fmt.Errorf("workers script subdomain still exists for script %s", scriptName)
		}
	}

	return nil
}

func testAccCheckCloudflareWorkersScriptSubdomain(rnd, accountID string) string {
	return acctest.LoadTestCase("workersscriptsubdomain.tf", rnd, accountID)
}

func testAccCheckCloudflareWorkersScriptSubdomainDisabled(rnd, accountID string) string {
	return acctest.LoadTestCase("workersscriptsubdomaindisabled.tf", rnd, accountID)
}

func testAccCheckCloudflareWorkersScriptSubdomainWithPreviews(rnd, accountID string) string {
	return acctest.LoadTestCase("workersscriptsubdomainwithpreviews.tf", rnd, accountID)
}
